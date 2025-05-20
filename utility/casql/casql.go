package casql

import (
	"math"

	"github.com/gogf/gf/v2/database/gdb"
)

func MakePage(db *gdb.Model, size uint) (out *BaseListOut, err error) {

	itemCount, err := db.Count()
	if err != nil {
		return
	}

	pageCount := uint(math.Ceil(float64(itemCount) / float64(size)))

	out = &BaseListOut{
		ItemCount: uint(itemCount),
		PageCount: pageCount,
	}

	return
}
