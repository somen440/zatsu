package migration

import "bytes"

type StructureInterface interface {
	GetUp() string
}

type Migration struct {
	UpList []string
}

func NewMigration(list []StructureInterface) *Migration {
	rs := &Migration{}
	for _, v := range list {
		rs.UpList = append(rs.UpList, v.GetUp())
	}
	return rs
}

func (m *Migration) String() string {
	var out bytes.Buffer

	out.WriteString("*************************** Up ***************************" + "\n")
	for _, up := range m.UpList {
		out.WriteString(up + "\n")
	}

	return out.String()
}
