package main

import (
	"fmt"
	"reflect"
)

// START OMIT

func GimmeByName(i interface{}, name string) interface{} {
	r := reflect.ValueOf(i)
	f := r.FieldByName(name)
	return f.Interface()
}

func main() {

	t := map[string]int{
		"a": 1,
	}
	fmt.Println(GimmeByName(t, "Name"))
}

// END OMIT
