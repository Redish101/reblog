package model

type Article struct {
	BaseModel

	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Desc    string `json:"desc"`
	Cover   *string `json:"cover"`
	Content string `json:"content"`
	Draft   *bool  `json:"draft"`
	AiSummary *string `json:"ai_summary"`
}
