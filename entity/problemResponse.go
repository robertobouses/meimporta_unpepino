package entity

type ProblemResponse struct {
	Name          string `json:"name"`
	Pests         string `json:"pests       "`
	Difficulties  string `json:"difficulties "`
	Care          string `json:"care"`
	Miscellaneous string `json:"miscellaneous"`
}
