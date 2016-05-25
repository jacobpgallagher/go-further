package main

import (
	"fmt"
	"reflect"
)

// START OMIT
type Thing struct {
	Name string `db:"db_name" json:"json_name" description:"stuff and stuff"`
}

func GimmeTagByName(i interface{}, fieldName string, tagName string) string {
	r := reflect.TypeOf(i)
	if f, exists := r.FieldByName(fieldName); exists {
		return f.Tag.Get(tagName)
	}
	return ""
}

func main() {
	t := Thing{"Name"}
	fmt.Println(GimmeTagByName(t, "Name", "db"))
	fmt.Println(GimmeTagByName(t, "Name", "json"))
	fmt.Println(GimmeTagByName(t, "Name", "description"))
}

//END OMIT
