package models

type root_json struct {
	Status string `json:"status"`
}

var RootDataGetALL root_json = root_json{
	Status: "get all",
}

var RootDataPut root_json = root_json{
	Status: "put",
}
