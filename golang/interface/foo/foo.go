package foo

import (
	"fmt"

	"github.com/somen440/zatsu/golang/interface/di"
)

func Exec(h di.Hoge) {
	fmt.Println("\t" + h.HogeMethod())
	fmt.Println("\t" + h.Bar().BarMethod())
}
