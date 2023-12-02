package entity

type IssueResponse struct {
	Name          string `json:"name"`
	Pests         string `json:"pests"`
	Difficulties  string `json:"difficulties"`
	Care          string `json:"care"`
	Miscellaneous string `json:"miscellaneous"`
}
