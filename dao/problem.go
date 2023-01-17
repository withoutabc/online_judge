package dao

import (
	"database/sql"
	"fmt"
	"online_judge/model"
	"strings"
)

func InsertProblem(p model.Problem) (err error) {
	_, err = DB.Exec("insert into problem (title,description,description_input,description_output,sample_input,sample_output,time_limit,memory_limit,uid) values (?,?,?,?,?,?,?,?,?)", p.Title, p.Description, p.DescriptionInput, p.DescriptionOutput, p.SampleInput, p.SampleOutput, p.TimeLimit, p.MemoryLimit, p.Uid)
	return

}

func SearchProblems(pid string) (problems []model.Problem, err error) {
	var rows *sql.Rows
	if pid == "" {
		rows, err = DB.Query("select * from problem ")
	} else {
		rows, err = DB.Query("select * from problem where pid=?", pid)
	}
	if err != nil {
		return nil, err
	}
	//处理查询结果
	for rows.Next() {
		var p model.Problem
		if err = rows.Scan(&p.Pid, &p.Title, &p.Description, &p.DescriptionInput, &p.DescriptionOutput, &p.SampleInput, &p.SampleOutput, &p.TimeLimit, &p.MemoryLimit, &p.Uid); err != nil {
			return nil, err
		}
		problems = append(problems, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return problems, nil
}

func UpdateProblem(p model.Problem) (err error) {
	var sql strings.Builder
	var arg []interface{}
	sql.WriteString("update problem set")
	if p.Title != "" {
		if len(arg) > 0 {
			sql.WriteString(",")
		}
		sql.WriteString(" title=?")
		arg = append(arg, p.Title)
	}
	if p.Description != "" {
		if len(arg) > 0 {
			sql.WriteString(",")
		}
		sql.WriteString(" description=?")
		arg = append(arg, p.Description)
	}
	if p.DescriptionInput != "" {
		if len(arg) > 0 {
			sql.WriteString(",")
		}
		sql.WriteString(" description_input=?")
		arg = append(arg, p.DescriptionInput)
	}
	if p.DescriptionOutput != "" {
		if len(arg) > 0 {
			sql.WriteString(",")
		}
		sql.WriteString(" description_output=?")
		arg = append(arg, p.DescriptionOutput)
	}
	if p.SampleInput != "" {
		if len(arg) > 0 {
			sql.WriteString(",")
		}
		sql.WriteString(" sample_input=?")
		arg = append(arg, p.SampleInput)
	}
	if p.SampleOutput != "" {
		if len(arg) > 0 {
			sql.WriteString(",")
		}
		sql.WriteString(" sample_output=?")
		arg = append(arg, p.SampleOutput)
	}

	if p.TimeLimit != 0 {
		if len(arg) > 0 {
			sql.WriteString(",")
		}
		sql.WriteString(" time_limit=?")
		arg = append(arg, p.TimeLimit)
	}
	if p.MemoryLimit != 0 {
		if len(arg) > 0 {
			sql.WriteString(",")
		}
		sql.WriteString(" memory_limit=?")
		arg = append(arg, p.MemoryLimit)
	}
	sql.WriteString(" where pid=?")
	arg = append(arg, p.Pid)
	fmt.Println(sql.String())
	_, err = DB.Exec(sql.String(), arg...)
	return
}
