package ytemplate

import (
	"html/template"
	"sir"
)

var ThePool Pool

type TemplateCache map[string]*template.Template

type Pool struct {
	Pools TemplateCache
}

func render(filenames ...string) *template.Template {
	t := template.New("layout")
	t.Delims("//", "//")

	t, err := t.ParseFiles(filenames...)
	sir.CheckError(err)

	return t
}

func (p *Pool) Fill(key string, filenames ...string) {
	if p.Pools == nil {
		p.Pools = make(TemplateCache)
	}

	p.Pools[key] = render(filenames...)
}
