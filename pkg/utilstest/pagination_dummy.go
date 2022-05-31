package utilstest

import (
	"github.com/lffranca/bento/pkg/domain/entity"
)

func GetDummyPagination() entity.Pagination {
	return entity.Pagination{
		Offset: 0,
		Limit:  100,
	}
}
