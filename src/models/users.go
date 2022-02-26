package models

type Users struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MsgSingleUser struct {
	Msg  string `json:"msg"`
	User Users  `json:"user"`
}

type MsgManyUsers struct {
	Msg   string  `json:"msg"`
	Users []Users `json:"users"`
}
