package main

import (
	"fmt"
	"reflect"
)

// START OMIT
type Thing struct {
	Name   string
	Count  int64
	Exists bool
}

func GimmeByName(i interface{}, name string) interface{} {
	r := reflect.ValueOf(i)
	f := r.FieldByName(name)
	return f.Interface()
}

func main() {
	t := Thing{"A Thing", 1, true}
	fmt.Println(GimmeByName(t, "Name"))
	fmt.Println(GimmeByName(t, "Count"))
	fmt.Println(GimmeByName(t, "Exists"))
}

// END OMIT
