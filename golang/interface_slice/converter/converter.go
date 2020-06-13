package converter

import (
	"github.com/somen440/zatsu/golang/interface_slice/migration"
	"github.com/somen440/zatsu/golang/interface_slice/structure"
)

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

func (c *Converter) ToMigrate(list []*structure.Structure) *migration.Migration {
	rs := []migration.StructureInterface{}
	for _, v := range list {
		rs = append(rs, v)
	}
	// cannot use list (variable of type []*structure.Structure) as []migration.StructureInterface value in argument to append
	// rs = append(rs, list...)
	return migration.NewMigration(rs)
}
