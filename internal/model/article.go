package model

type Article struct {
	BaseModel

	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Content string `json:"content"`
}
