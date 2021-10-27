package model

type Depart struct {
	*BaseModel
	Name   string `json:"name"`
	Leader string `json:"leader"`
	Email  string `json:"email"`
}
