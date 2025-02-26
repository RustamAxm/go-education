
package my_module

import "testing"
import "fmt"

func TestSimple(t *testing.T) {
	fmt.Println("in test")
	m := my_module.NewModule()
	m.Module_func()
	if got := m.Module_func(); got != "in module" {
		t.Errorf("error in module test = %q", got)
	}
}
