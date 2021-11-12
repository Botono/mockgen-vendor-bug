package main

import (
	"fmt"
)

type TypeA struct{}

func (a *TypeA) Foo() {}

type TypeB struct{}

func (b *TypeB) Bar(manyTypeAs ...TypeA) {
	fmt.Printf("All the types: %+v", manyTypeAs)
}

func main() {

}
