package entity

type Pagination struct {
	Offset int `json:"offset" yaml:"offset"`
	Limit  int `json:"limit" yaml:"limit"`
}
