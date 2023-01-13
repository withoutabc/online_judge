package dao

import (
	"fmt"
	"online_judge/model"
	"strings"
)

func InsertSubmission(s model.Submission) (err error) {
	_, err = DB.Exec("insert into submission (pid,uid,code,language) values (?,?,?,?)", s.Pid, s.Uid, s.Code, s.Language)
	return
}

func QueryResult(s model.Submission) (submissions []model.Submission, err error) {
	var sql strings.Builder
	var arg []interface{}
	var b bool
	sql.WriteString("select * from submission")
	if s.Pid != "" {
		if !b {
			sql.WriteString(" where")
			b = true
		} else {
			sql.WriteString(" and")
		}
		sql.WriteString(" pid=?")
		arg = append(arg, s.Pid)
	}
	if s.Uid != "" {
		if !b {
			sql.WriteString(" where")
			b = true
		} else {
			sql.WriteString(" and")
		}
		sql.WriteString(" uid=?")
		arg = append(arg, s.Uid)
	}
	if s.Code != "" {
		if !b {
			sql.WriteString(" where")
			b = true
		} else {
			sql.WriteString(" and")
		}
		sql.WriteString(" code=?")
		arg = append(arg, s.Code)
	}
	if s.Language != "" {
		if !b {
			sql.WriteString(" where")
			b = true
		} else {
			sql.WriteString(" and")
		}
		sql.WriteString(" language=?")
		arg = append(arg, s.Language)
	}
	if s.Status != "" {
		if !b {
			sql.WriteString(" where")
			b = true
		} else {
			sql.WriteString(" and")
		}
		sql.WriteString(" status=?")
		arg = append(arg, s.Status)
	}
	rows, err := DB.Query(sql.String(), arg...)
	fmt.Println(sql.String())
	if err != nil {
		return nil, err
	}
	//处理查询结果
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
	return submissions, nil
}
