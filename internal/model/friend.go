package model

type Friend struct {
	BaseModel

	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	URL    string `json:"url"`
	Desc   string `json:"desc"`
}
