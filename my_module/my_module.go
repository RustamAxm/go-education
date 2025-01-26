package my_module

import "fmt"


type Module struct {
	Name string
}

func (m *Module) Module_func() {
	fmt.Println("in module")
}

func NewModule() *Module {
	return &Module{Name: "MyModule"}
}
