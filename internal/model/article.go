package model

type Article struct {
	BaseModel

	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
	Draft   *bool  `json:"draft"`
}
