package main

import (
	_ "embed"
	"text/template"
)

type Basic struct{}

//go:embed basic.tmd
var TplBasic string

func (Basic) Template() *template.Template {
	t, _ := template.New("templatePFI").Parse(TplBasic)
	return t
}
