package dao

import (
	"fmt"
	"online_judge/model"
	"strings"
)

func InsertTestcase(t model.Testcase) (err error) {
	_, err = DB.Exec("insert into testcase (pid,uid,input,output) values (?,?,?,?)", t.Pid, t.Uid, t.Input, t.Output)
	return
}

func SearchTestcase(uid string, pid string) (testcases []model.Testcase, err error) {
	if pid == "" {
		rows, err := DB.Query("select * from testcase where uid=?", uid)
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
			rows, err := DB.Query("select * from testcase where uid=? and pid=?", uid, pid)
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
	var sql strings.Builder
	var arg []interface{}
	sql.WriteString("update product set")
	if t.Input != "" {
		if len(arg) > 0 {
			sql.WriteString(",")
		}
		sql.WriteString(" input=?")
		arg = append(arg, t.Input)
	}
	if t.Output != "" {
		if len(arg) > 0 {
			sql.WriteString(",")
		}
		sql.WriteString(" output=?")
		arg = append(arg, t.Output)
	}
	sql.WriteString(" where pid=? and uid=?")
	arg = append(arg, t.Pid)
	arg = append(arg, t.Uid)
	fmt.Println(sql.String())
	_, err = DB.Exec(sql.String(), arg...)
	return
}

func DeleteTestcase(uid string, tid string) (err error) {
	_, err = DB.Exec("delete from testcase where uid=? and tid =?", uid, tid)
	return
}
