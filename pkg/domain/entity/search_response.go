package entity

// ListOpt aids in paginating through list endpoints
type ListOpt struct {
	Size         int      `json:"size,omitempty" yaml:"size,omitempty"`
	Offset       string   `json:"offset,omitempty" yaml:"offset,omitempty"`
	Tags         []string `json:"tags,omitempty" yaml:"tags,omitempty"`
	MatchAllTags bool     `json:"match_all_tags,omitempty" yaml:"match_all_tags,omitempty"`
}

type SearchResponse[T any] struct {
	Items   []T     `json:"items" yaml:"items"`
	Options ListOpt `json:"options" yaml:"options"`
}
