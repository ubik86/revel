package app

import (
	"fmt"
	"github.com/ubik86/revel"
)

func init() {
	revel.OnAppStart(func() {
		fmt.Println("Go to /@tests to run the tests.")
	})
}
