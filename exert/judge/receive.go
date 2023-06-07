package judge

import (
	"bytes"
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"io/ioutil"
	"log"
	"online_judge/model"
	"online_judge/redis"
	"online_judge/util"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var langMap = map[string]string{
	"Go":     "go",
	"Python": "py",
	"C++":    "cpp",
	"C":      "c",
	"Java":   "java",
}

func Consume(ch *amqp.Channel, queueName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto ack
		false,     // exclusive
		false,     // no local
		false,     // no wait
		nil,       // args
	)
	if err != nil {
		panic(err)
	}
	var forever chan struct{}

	for d := range msgs {
		var submissionId int64
		submissionId, err = strconv.ParseInt(string(d.Body), 10, 64)
		if err != nil {
			util.Log(err)
			continue
		}
		//根据id查询submission的code
		var submission model.Submission
		submission, err = NewJudImpl().SubmissionDao.SearchSubmissionById(submissionId)
		if err != nil {
			//没查询到跳过
			util.Log(err)
			continue
		}
		//	把code写入code.
		filename := fmt.Sprintf("code.%s", langMap[submission.Language])
		err = ioutil.WriteFile(filename, []byte(submission.Code), 0644)
		if err != nil {
			util.Log(err)
			continue
		}
		//放到容器里面跑，得到输出结果
		Judge(ctx, submission)
		//删除redis缓存
		redis.DeleteSubmissionId(context.Background(), submissionId)
		//确认
		d.Ack(false)
	}
	<-forever
}

func Judge(ctx context.Context, submission model.Submission) {
	//寻找测试用例
	testcases, err := NewJudImpl().TestcaseDao.SearchTestcase(submission.ProblemId)
	log.Println(testcases)
	if err != nil {
		util.Log(err)
		panic(err)
	}
	if len(testcases) == 0 {
		return
	}
	switch submission.Language {
	case "Go":
		Go(submission, testcases)
		return
	case "C":
		C(submission, testcases)
		return
	case "C++":
		Cpp(submission, testcases)
		return
	case "Python":
		Python(submission, testcases)
		return
	case "Java":
		Java(submission, testcases)
		return
	}

}

func Go(submission model.Submission, testcases []model.Testcase) {
	log.Println("go")
	var count = 0
	for _, testcase := range testcases {
		fmt.Printf("input:%s\n", testcase.Input)
		//1.写入input
		err := ioutil.WriteFile("input-go.txt", []byte(testcase.Input), 0644)
		if err != nil {
			util.Log(err)
			continue
		}
		//2.复制input.txt
		cmd := exec.Command("docker", "cp", "input-go.txt", "golang-judge:go/src/app")
		err = cmd.Run()
		if err != nil {
			util.Log(err)
			if err.Error() == "exit status 1" {
				return
			}
			continue
		}
		//3.复制code.go
		cmd = exec.Command("docker", "cp", "code.go", "golang-judge:/go/src/app")
		err = cmd.Run()
		if err != nil {
			util.Log(err)
			continue
		}
		//4.编译、运行
		cmd = exec.Command("docker", "exec", "golang-judge", "sh", "-c", "go build -o /go/src/app/exert /go/src/app/code.go && timeout 2s /go/src/app/exert < input-go.txt")
		var out bytes.Buffer
		cmd.Stdout = &out
		err = cmd.Run()
		log.Println(strings.TrimSpace(out.String()))
		if err != nil {
			if err.Error() == "exit status 124" {
				log.Println("超时")
				err = NewJudImpl().SubmissionDao.UpdateStatus(submission.SubmissionId, "运行超时")
				if err != nil {
					util.Log(err)
					panic(err)
				}
				return
			} else {
				log.Println("CE")
				err = NewJudImpl().SubmissionDao.UpdateStatus(submission.SubmissionId, "编译错误")
				if err != nil {
					util.Log(err)
					panic(err)
				}
				util.Log(err)
				return
			}
		}
		//5.处理输出结果
		if strings.TrimSpace(out.String()) != testcase.Output {
			fmt.Println("wrong")
			err = NewJudImpl().SubmissionDao.UpdateStatus(submission.SubmissionId, "答案错误")
			if err != nil {
				util.Log(err)
				panic(err)
			}
			return
		} else {
			count++
			log.Printf("正确%d次", count)
			if count == len(testcases) {
				log.Println("start add score and correct")
				DealWithCorrection(submission)
			}
		}
	}
}

func C(submission model.Submission, testcases []model.Testcase) {
	var count = 0
	for _, testcase := range testcases {
		fmt.Printf("input:%s\n", testcase.Input)
		//1.写入input
		err := ioutil.WriteFile("input-c.txt", []byte(testcase.Input), 0644)
		if err != nil {
			log.Printf("1 %v", err)
			util.Log(err)
			continue
		}
		//2.复制input.txt
		cmd := exec.Command("docker", "cp", "input-c.txt", "c-judge:c/src/app")
		err = cmd.Run()
		if err != nil {
			util.Log(err)
			log.Printf("2 %v", err)
			if err.Error() == "exit status 1" {
				return
			}
			continue
		}
		//3.复制code.go
		cmd = exec.Command("docker", "cp", "code.c", "c-judge:c/src/app")
		err = cmd.Run()
		if err != nil {
			log.Printf("3 %v", err)
			util.Log(err)
			continue
		}
		//4.编译、运行
		cmd = exec.Command("docker", "exec", "c-judge", "sh", "-c", "gcc -o /c/src/app/exert /c/src/app/code.c && timeout 2s /c/src/app/exert < input-c.txt")
		var out bytes.Buffer
		cmd.Stdout = &out
		err = cmd.Run()
		log.Println(strings.TrimSpace(out.String()))
		if err != nil {
			log.Printf("4 %v", err)
			if err.Error() == "exit status 124" {
				log.Println("超时")
				err = NewJudImpl().SubmissionDao.UpdateStatus(submission.SubmissionId, "运行超时")
				if err != nil {
					util.Log(err)
					panic(err)
				}
				return
			} else {
				log.Println("CE")
				err = NewJudImpl().SubmissionDao.UpdateStatus(submission.SubmissionId, "编译错误")
				if err != nil {
					util.Log(err)
					panic(err)
				}
				util.Log(err)
				return
			}
		}
		//5.处理输出结果
		if strings.TrimSpace(out.String()) != testcase.Output {
			fmt.Println("wrong")
			err = NewJudImpl().SubmissionDao.UpdateStatus(submission.SubmissionId, "答案错误")
			if err != nil {
				util.Log(err)
				panic(err)
			}
			return
		} else {
			count++
			log.Println(count)
			if count == len(testcases) {
				DealWithCorrection(submission)
			}
		}

	}
}

func Cpp(submission model.Submission, testcases []model.Testcase) {
	var count = 0
	for _, testcase := range testcases {
		fmt.Printf("input:%s\n", testcase.Input)
		//1.写入input
		err := ioutil.WriteFile("input-cpp.txt", []byte(testcase.Input), 0644)
		if err != nil {
			util.Log(err)
			continue
		}
		//2.复制input.txt
		cmd := exec.Command("docker", "cp", "input-cpp.txt", "cpp-judge:cpp/src/app")
		err = cmd.Run()
		if err != nil {
			util.Log(err)
			if err.Error() == "exit status 1" {
				return
			}
			continue
		}
		//3.复制code.go
		cmd = exec.Command("docker", "cp", "code.cpp", "cpp-judge:/cpp/src/app")
		err = cmd.Run()
		if err != nil {
			util.Log(err)
			continue
		}
		//4.编译、运行
		cmd = exec.Command("docker", "exec", "cpp-judge", "sh", "-c", "g++ -o /cpp/src/app/exert /cpp/src/app/code.cpp && timeout 2s /cpp/src/app/exert < input-cpp.txt")
		var out bytes.Buffer
		cmd.Stdout = &out
		err = cmd.Run()
		if err != nil {
			if err.Error() == "exit status 124" {
				log.Println("超时")
				err = NewJudImpl().SubmissionDao.UpdateStatus(submission.SubmissionId, "运行超时")
				if err != nil {
					util.Log(err)
					panic(err)
				}
				return
			} else {
				log.Println("CE")
				err = NewJudImpl().SubmissionDao.UpdateStatus(submission.SubmissionId, "编译错误")
				if err != nil {
					util.Log(err)
					panic(err)
				}
				util.Log(err)
				return
			}
		}
		log.Println(strings.TrimSpace(out.String()))
		//5.处理输出结果
		if strings.TrimSpace(out.String()) != testcase.Output {
			fmt.Println("wrong")
			err = NewJudImpl().SubmissionDao.UpdateStatus(submission.SubmissionId, "答案错误")
			if err != nil {
				util.Log(err)
				panic(err)
			}
			return
		} else {
			count++
			log.Println(count)
			if count == len(testcases) {
				DealWithCorrection(submission)
			}
		}
	}
}

func Java(submission model.Submission, testcases []model.Testcase) {
	var count = 0
	log.Println(submission)
	for _, testcase := range testcases {
		fmt.Printf("input:%s\n", testcase.Input)
		//1.写入input
		err := ioutil.WriteFile("input-java.txt", []byte(testcase.Input), 0644)
		if err != nil {
			util.Log(err)
			continue
		}
		//2.复制input.txt
		cmd := exec.Command("docker", "cp", "input-java.txt", "java-judge:java/src/app")
		err = cmd.Run()
		if err != nil {
			util.Log(err)
			if err.Error() == "exit status 1" {
				return
			}
			continue
		}
		//3.复制code.go
		cmd = exec.Command("docker", "cp", "code.java", "java-judge:/java/src/app")
		err = cmd.Run()
		if err != nil {
			util.Log(err)
			continue
		}
		//4.编译、运行
		cmd = exec.Command("docker", "exec", "java-judge", "sh", "-c", "javac /java/src/app/code.java && timeout 2s java Main < input-java.txt")
		var out bytes.Buffer
		cmd.Stdout = &out
		err = cmd.Run()
		if err != nil {
			if err.Error() == "exit status 124" {
				log.Println("超时")
				err = NewJudImpl().SubmissionDao.UpdateStatus(submission.SubmissionId, "运行超时")

				if err != nil {
					util.Log(err)
					panic(err)
				}
				return
			} else {
				log.Println(err)
				log.Println("CE")
				err = NewJudImpl().SubmissionDao.UpdateStatus(submission.SubmissionId, "编译错误")

				if err != nil {
					util.Log(err)
					panic(err)
				}
				util.Log(err)
				return
			}
		}
		log.Println(strings.TrimSpace(out.String()))
		//5.处理输出结果
		if strings.TrimSpace(out.String()) != testcase.Output {
			fmt.Println("wrong")
			err = NewJudImpl().SubmissionDao.UpdateStatus(submission.SubmissionId, "答案错误")
			if err != nil {
				util.Log(err)
				panic(err)
			}
			return
		} else {
			count++
			log.Println(count)
			if count == len(testcases) {
				DealWithCorrection(submission)
			}
		}

	}
}

func Python(submission model.Submission, testcases []model.Testcase) {
	log.Println("py")
	var count = 0
	for _, testcase := range testcases {
		fmt.Printf("input:%s\n", testcase.Input)
		//1.写入input
		err := ioutil.WriteFile("input-python.txt", []byte(testcase.Input), 0644)
		if err != nil {
			util.Log(err)
			continue
		}
		//2.复制input.txt
		cmd := exec.Command("docker", "cp", "input-python.txt", "python-judge:python/src/app")
		err = cmd.Run()
		if err != nil {
			util.Log(err)
			if err.Error() == "exit status 1" {
				return
			}
			continue
		}
		//3.复制code.go
		cmd = exec.Command("docker", "cp", "code.py", "python-judge:/python/src/app")
		err = cmd.Run()
		if err != nil {
			util.Log(err)
			continue
		}
		//4.编译、运行
		cmd = exec.Command("docker", "exec", "python-judge", "sh", "-c", "timeout 2s python /python/src/app/code.py < input-python.txt")
		var out bytes.Buffer
		cmd.Stdout = &out
		err = cmd.Run()
		log.Println(cmd.Err)
		log.Println(err)
		if err != nil {
			if err.Error() == "exit status 124" {
				log.Println("超时")
				err = NewJudImpl().SubmissionDao.UpdateStatus(submission.SubmissionId, "运行超时")
				if err != nil {
					util.Log(err)
					panic(err)
				}
				return
			} else {
				log.Println("CE")
				err = NewJudImpl().SubmissionDao.UpdateStatus(submission.SubmissionId, "编译错误")
				if err != nil {
					util.Log(err)
					panic(err)
				}
				util.Log(err)
				return
			}
		}
		log.Println(strings.TrimSpace(out.String()))
		//5.处理输出结果
		if strings.TrimSpace(out.String()) != testcase.Output {
			fmt.Println("wrong")
			err = NewJudImpl().SubmissionDao.UpdateStatus(submission.SubmissionId, "答案错误")
			if err != nil {
				util.Log(err)
				panic(err)
			}
			return
		} else {
			count++
			log.Println(count)
			if count == len(testcases) {
				DealWithCorrection(submission)
			}
		}
	}

}

func DealWithCorrection(submission model.Submission) {
	// '极易' '容易' '中等' '困难' '极难'
	var score = map[string]int{
		"极易": 5,
		"容易": 8,
		"中等": 10,
		"困难": 13,
		"极难": 16,
	}
	problems, err := NewJudImpl().ProblemDao.SearchProblem(model.ReqSearchProblem{ProblemId: submission.ProblemId})
	if err != nil {
		util.Log(err)
		panic(err)
	}
	if len(problems) == 0 {
		return
	}
	if err = NewJudImpl().SubmissionDao.UpdateStatus(submission.SubmissionId, "正确"); err != nil {
		util.Log(err)
		panic(err)
	}
	if err = NewJudImpl().ProblemDao.AddCorrect(submission.ProblemId); err != nil {
		util.Log(err)
		panic(err)
	}
	if err = NewJudImpl().InfoDao.AddCorrect(submission.UserId); err != nil {
		util.Log(err)
		panic(err)
	}
	if err = NewJudImpl().InfoDao.AddScore(submission.UserId, score[problems[0].Level]); err != nil {
		util.Log(err)
		panic(err)
	}
}
