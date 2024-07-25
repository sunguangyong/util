package structutil

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/mitchellh/copystructure"
	"testing"
)

type StructB struct {
	Field1 string
	Field2 int
}

type StructA struct {
	Field1 string
	Field2 int
}

func StructCopy() {
	a := StructA{"value1", 42}
	var b StructB
	copier.Copy(&b, &a)
	fmt.Println(b.Field1)
	fmt.Println(b.Field2)
}

func TestStruct(t testing.T) {
	a := StructA{"value1", 42}
	b, err := copystructure.Copy(&a)
	a.Field2 = 100
	fmt.Println(b, err)
	fmt.Println(a)
}
