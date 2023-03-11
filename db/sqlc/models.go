// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import ()

type Project struct {
	ProjectName     string      `json:"project_name"`
	Symbol          string      `json:"symbol"`
	ContractAddress string      `json:"contract_address"`
	Owner           string      `json:"owner"`
	CreatedAt       interface{} `json:"created_at"`
}

type User struct {
	Username     string      `json:"username"`
	HashedPwd    string      `json:"hashed_pwd"`
	Email        string      `json:"email"`
	PwdChangedAt interface{} `json:"pwd_changed_at"`
	CreatedAt    interface{} `json:"created_at"`
}
