package service

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"online_judge/model"
	"os/exec"
	"strconv"
	"time"
)

func Judge() {
	//每2分钟一次
	ticker := time.NewTicker(time.Minute * 2)
	for range ticker.C {
		//查找所有pending的submission
		submissions, err := SearchPendingCode()
		if err != nil {
			fmt.Printf("search pending code err:%v", err)
			return
		}
		//如果没有就直接结束
		if submissions == nil {
			fmt.Println("none")
			return
		}
		//遍历每个submission，用submission对应的pid去找到所有的testcase,运行并输入
		var submission model.Submission
		for _, submission = range submissions {
			//把code写入code.go
			err = ioutil.WriteFile("code.go", []byte(submission.Code), 0644)
			if err != nil {
				fmt.Printf("write code err:%v", err)
				return
			}
			//查询testcase
			var testcases []model.Testcase
			testcases, err = SearchTestcasesByPid(submission.Pid)
			if err != nil {
				fmt.Printf("search testcase err:%v", err)
				return
			}
			fmt.Println(testcases)
			//遍历，依次传入input运行
			for k, testcase := range testcases {
				fmt.Printf("input:%s\n", testcase.Input)
				fmt.Printf("output:%s\n", testcase.Output)
				//写入input
				err = ioutil.WriteFile("input.txt", []byte(testcase.Input), 0644)
				if err != nil {
					fmt.Printf("write input err:%v", err)
					return
				}
				//docker步骤
				//1.构建镜像
				cmd := exec.Command("docker", "build", "-t", "go", ".")
				var out bytes.Buffer
				cmd.Stdout = &out
				err = cmd.Run()
				if err != nil {
					fmt.Printf("build image err:%v\n", err)
					return
				}
				fmt.Println("build image success")
				//2.复制input.txt
				cmd = exec.Command("docker", "cp", "input.txt", "Online_judge:go/src/app")
				err = cmd.Run()
				if err != nil {
					fmt.Printf("copy input.txt err:%v\n", err)
					return
				}
				fmt.Println("copy input.txt success")
				//3.复制code.go
				cmd = exec.Command("docker", "cp", "code.go", "Online_judge:/go/src/app")
				err = cmd.Run()
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("copy code.go success")
				//4.运行
				cmd = exec.Command("docker", "exec", "Online_judge", "sh", "-c", "go build -o /go/src/app/main /go/src/app/code.go && /go/src/app/main < input.txt")
				cmd.Stdout = &out
				err = cmd.Run()
				if err != nil {
					fmt.Printf("run err:%v\n", err)
					err = UpdateStatus("Compile error", strconv.Itoa(submission.Sid))
					if err != nil {
						fmt.Printf("update to Compile Error err:%v", err)
						return
					}
					break
				}
				fmt.Println("run success")
				//5.处理输出结果
				result := out.String()
				if result != testcase.Output {
					err = UpdateStatus("Wrong Answer", strconv.Itoa(submission.Sid))
					if err != nil {
						fmt.Printf("update to wrong answer err:%v", err)
						break
					}
					var count int
					count++
					if count == k+1 {
						err = UpdateStatus("Accepted", strconv.Itoa(submission.Sid))
						if err != nil {
							fmt.Printf("update to accepted err:%v", err)
							return
						}
					}
				}
			}
		}
	}
}
