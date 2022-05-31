package utilstest

import (
	"github.com/lffranca/bento/pkg/domain/entity"
)

func GetDummyTagOption() entity.TagOption {
	return entity.TagOption{
		Tags:         []string{"tag1", "tag2"},
		MatchAllTags: true,
	}
}
