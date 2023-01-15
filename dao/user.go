package dao

import "online_judge/model"

func SearchUserByUsername(username string) (u model.User, err error) {
	row := DB.QueryRow("select * from user where username=?", username)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.Uid, &u.Username, &u.Password, &u.Salt)
	return
}

func SearchUserByUid(uid string) (u model.User, err error) {
	row := DB.QueryRow("select * from user where uid=?", uid)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.Uid, &u.Username, &u.Password, &u.Salt)
	return
}

func InsertUser(u model.User) (err error) {
	_, err = DB.Exec("insert into user(username,password,salt) values (?,?,?)", u.Username, u.Password, u.Salt)
	return
}

func UpdatePassword(newPassword []byte, username string, salt []byte) (err error) {
	_, err = DB.Exec("update user set password=?,salt=? where username=?", newPassword, salt, username)
	return
}
