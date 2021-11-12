package main

type TypeAInterface interface {
	Foo()
}

type TypeBInterface interface {
	Bar(manyTypeAs ...TypeA)
}
