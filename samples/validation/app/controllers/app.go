package controllers

import "github.com/ubik86/revel"

type Application struct {
	*revel.Controller
}

func (c Application) Index() revel.Result {
	return c.Render()
}
