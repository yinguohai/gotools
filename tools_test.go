package gotools

import (
	"fmt"
	"testing"
)

func TestListToMap(t *testing.T) {
	list1 := []string{"name", "sex", "age",}

	list2 := []string{"Smitch", "boy", "20",}

	content , err := ListToMap(list1,list2)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(content)
}

func TestColumeMap(t *testing.T) {
	mapDemo := []map[string]string {
		{"name":"Smith","sex":"boy","age":"20",},
		{"name":"angularBaby","sex":"girl","age":"30"},
		{"name":"Bob","sex":"boy","age":"40"},
	}

	result := ColumeMap(mapDemo,"name")

	fmt.Println(result)
}
