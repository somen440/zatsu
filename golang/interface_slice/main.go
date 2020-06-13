package main

import (
	"fmt"

	"github.com/somen440/zatsu/golang/interface_slice/converter"
	"github.com/somen440/zatsu/golang/interface_slice/migration"
	"github.com/somen440/zatsu/golang/interface_slice/structure"
)

var (
	c *converter.Converter
	_ migration.StructureInterface = &structure.Structure{}
)

func init() {
	c = converter.NewConverter()
}

func main() {
	sl := []*structure.Structure{
		structure.NewStructure("up1"),
		structure.NewStructure("up2"),
		structure.NewStructure("up3"),
	}
	m := c.ToMigrate(sl)
	fmt.Println(m)
}
