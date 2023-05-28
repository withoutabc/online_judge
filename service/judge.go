package service

//func Judge() {
//	//查找所有pending的submission
//	submissions, err := SearchPendingCode()
//	if err != nil {
//		fmt.Printf("search pending code err:%v", err)
//		return
//	}
//	//如果没有就直接结束
//	if submissions == nil {
//		fmt.Println("none")
//		return
//	}
//	//遍历每个submission，用submission对应的pid去找到所有的testcase,运行并输入
//	var submission model.Submission
//	for _, submission = range submissions {
//		count := 0
//		//把code写入code.go
//		err = ioutil.WriteFile("code.go", []byte(submission.Code), 0644)
//		if err != nil {
//			fmt.Printf("write code err:%v", err)
//			return
//		}
//		//查询testcase
//		var testcases []model.Testcase
//		testcases, err = SearchTestcasesByPid(submission.Pid)
//		if err != nil {
//			fmt.Printf("search testcase err:%v", err)
//			return
//		}
//		//没有测试数据直接结束
//		if testcases == nil {
//			fmt.Println("none")
//			return
//		}
//		//遍历，依次传入input运行
//		for _, testcase := range testcases {
//			fmt.Printf("input:%s\n", testcase.Input)
//			//写入input
//			err = ioutil.WriteFile("input-go.txt", []byte(testcase.Input), 0644)
//			if err != nil {
//				fmt.Printf("write input err:%v", err)
//				return
//			}
//			//docker步骤
//			//1.构建镜像
//			cmd1 := exec.Command("docker", "build", "-t", "go", ".")
//
//			cmd1.Stdout = &out
//			err = cmd1.Run()
//			if err != nil {
//				fmt.Printf("build image err:%v\n", err)
//				return
//			}
//			fmt.Println("build image success")
//			//2.复制input.txt
//			cmd1 = exec.Command("docker", "cp", "input-go.txt", "Online_judge:go/src/app")
//			err = cmd1.Run()
//			if err != nil {
//				fmt.Printf("copy input-go.txt err:%v\n", err)
//				return
//			}
//			fmt.Println("copy input-go.txt success")
//			//3.复制code.go
//			cmd1 = exec.Command("docker", "cp", "code.go", "Online_judge:/go/src/app")
//			err = cmd1.Run()
//			if err != nil {
//				fmt.Println(err)
//				return
//			}
//			fmt.Println("copy code.go success")
//			//4.编译、运行
//			cmd1 = exec.Command("docker", "exec", "Online_judge", "sh", "-c", "go build -o /go/src/app/exert /go/src/app/code.go && /go/src/app/exert < input-go.txt")
//			cmd1.Stdout = &out
//			err = cmd1.Run()
//			if err != nil {
//				fmt.Printf("run err:%v\n", err)
//				err = UpdateStatus("Compile error", strconv.Itoa(submission.Sid))
//				if err != nil {
//					fmt.Printf("update to Compile Error err:%v", err)
//					return
//				}
//				break
//			}
//			fmt.Println("run success")
//			//5.处理输出结果
//			if strings.TrimSpace(out.String()) != testcase.Output {
//				fmt.Println("wrong")
//				err = UpdateStatus("Wrong Answer", strconv.Itoa(submission.Sid))
//				if err != nil {
//					fmt.Printf("update to wrong answer err:%v", err)
//					return
//				}
//				count++
//				continue
//			}
//			if count == 0 {
//				fmt.Println("correct")
//				err = UpdateStatus("Accepted", strconv.Itoa(submission.Sid))
//				if err != nil {
//					fmt.Printf("update to accepted err:%v", err)
//					return
//				}
//			}
//		}
//	}
//}
