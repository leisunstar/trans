package trans

import (
	"fmt"
	"testing"
)

func TestT(t *testing.T) {
	type Inventory struct {
		Material string
		Count    uint
	}

	sweaters := Inventory{"wool", 17}
	tt := NewMyTextTemplate()
	tt.AddTempFunc("testin", simplePrint)

	tt.AddTemplate("test", "{{.Count}} items are {{testin \"output\"}} made of {{.Material}}")
	tmp, err := tt.String("test", sweaters)
	if err != nil {

	}
	fmt.Println(tmp)
}
