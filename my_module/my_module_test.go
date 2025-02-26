
package my_module

import "testing"
import "fmt"

func TestSimple(t *testing.T) {
	fmt.Println("in test")
	m := NewModule()
	if got := m.Module_func(); got != "in module" {
		t.Errorf("error in module test = %q", got)
	}
}

func TestSimpleErr(t *testing.T) {
        fmt.Println("in test")
        m := NewModule()
        m.Module_func()
        if got := m.Module_func(); got != "in module 2 " {
                t.Errorf("error in module test = %q", got)
        }
}

