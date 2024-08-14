package model

type Site struct {
	BaseModel

	Name string `json:"name"`
	Url  string `json:"url"`
	Desc string `json:"desc"`
	Icon string `json:"icon"`
}
