package model

type User struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
	Password []byte `json:"password"`
	Salt     []byte `json:"salt"`
}

type User1 struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
}
