package dao

import (
	"database/sql"
	"fmt"
	"online_judge/model"
	"strings"
)

func InsertTestcase(t model.Testcase) (err error) {
	_, err = DB.Exec("insert into testcase (pid,uid,input,output) values (?,?,?,?)", t.Pid, t.Uid, t.Input, t.Output)
	return
}

func SearchTestcase(uid string, pid string) (testcases []model.Testcase, err error) {
	var rows *sql.Rows
	if pid == "" {
		rows, err = DB.Query("select * from testcase where uid=?", uid)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var testcase model.Testcase
			if err = rows.Scan(&testcase.Tid, &testcase.Pid, &testcase.Uid, &testcase.Input, &testcase.Output); err != nil {
				return nil, err
			}
			testcases = append(testcases, testcase)
		}
		if err = rows.Err(); err != nil {
			return nil, err
		}
	} else {
		if pid != "" {
			rows, err = DB.Query("select * from testcase where uid=? and pid=?", uid, pid)
			if err != nil {
				return nil, err
			}
			for rows.Next() {
				var testcase model.Testcase
				if err = rows.Scan(&testcase.Tid, &testcase.Pid, &testcase.Uid, &testcase.Input, &testcase.Output); err != nil {
					return nil, err
				}
				testcases = append(testcases, testcase)
			}
			if err = rows.Err(); err != nil {
				return nil, err
			}
		}
	}
	return
}

func UpdateTestcase(t model.Testcase) (err error) {
	var Sql strings.Builder
	var arg []interface{}
	Sql.WriteString("update testcase set")
	if t.Input != "" {
		if len(arg) > 0 {
			Sql.WriteString(",")
		}
		Sql.WriteString(" input=?")
		arg = append(arg, t.Input)
	}
	if t.Output != "" {
		if len(arg) > 0 {
			Sql.WriteString(",")
		}
		Sql.WriteString(" output=?")
		arg = append(arg, t.Output)
	}
	Sql.WriteString(" where pid=? and uid=? and tid=?")
	arg = append(arg, t.Pid)
	arg = append(arg, t.Uid)
	arg = append(arg, t.Tid)
	fmt.Println(Sql.String())
	_, err = DB.Exec(Sql.String(), arg...)
	return
}

func DeleteTestcase(uid string, tid string) (err error) {
	_, err = DB.Exec("delete from testcase where uid=? and tid =?", uid, tid)
	return
}

func GetPendingCode() (submissions []model.Submission, err error) {
	var rows *sql.Rows
	rows, err = DB.Query("select * from submission where status='Pending'")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var submission model.Submission
		if err = rows.Scan(&submission.Sid, &submission.Pid, &submission.Uid, &submission.Status, &submission.Code, &submission.Language); err != nil {
			return nil, err
		}
		submissions = append(submissions, submission)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return
}

func GetTestcasesByPid(pid string) (testcases []model.Testcase, err error) {
	var rows *sql.Rows
	rows, err = DB.Query("select * from testcase where pid=?", pid)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var testcase model.Testcase
		if err = rows.Scan(&testcase.Tid, &testcase.Pid, &testcase.Uid, &testcase.Input, &testcase.Output); err != nil {
			return nil, err
		}
		testcases = append(testcases, testcase)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return
}

func UpdateStatus(status string, sid string) (err error) {
	_, err = DB.Exec("update submission set status=? where sid=?", status, sid)
	return
}
