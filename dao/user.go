package dao

import "online_judge/model"

func SearchUserByUsername(username string) (u model.User, err error) {
	row := DB.QueryRow("select * from user where username=?", username)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.Uid, &u.Username, &u.Password)
	return
}

func SearchUserByUid(uid string) (u model.User, err error) {
	row := DB.QueryRow("select * from user where uid=?", uid)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.Uid, &u.Username, &u.Password)
	return
}

func InsertUser(u model.User) (err error) {
	_, err = DB.Exec("insert into user(username,password) values (?,?)", u.Username, u.Password)
	return
}

func UpdatePassword(newPassword string, username string) (err error) {
	_, err = DB.Exec("update user set password=? where username=?", newPassword, username)
	return
}
