package main

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

func main() {
	// START OMIT
	i := map[string]interface{}{
		"l1": map[string]interface{}{
			"l2": map[string]interface{}{
				"l3": true,
			},
		},
		"x": map[string]interface{}{
			"array": []string{"one", "two"},
		},
	}

	g := gabs.New()
	g.SetP(true, "l1.l2.l3")
	g.Array("x", "array")
	g.ArrayAppend("one", "x", "array")
	g.ArrayAppend("two", "x", "array")
	fmt.Println(i)
	fmt.Println(g.Data())
	// END OMIT
}
