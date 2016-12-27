package store

import (
	"github.com/animalmatsuzawa/ledisdb/store/driver"
)

type Slice interface {
	driver.ISlice
}
