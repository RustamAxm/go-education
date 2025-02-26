/*
docs tets for module 



!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!!!!!!!!!!!!!!!!!
*/



package my_module

import "fmt"


type Module struct {
	Name string
}

func (m *Module) Module_func() string {
	fmt.Println("in module")
	return "in module"
}

func NewModule() *Module {
	return &Module{Name: "MyModule"}
}
