package trans

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"
)

type MyTextTemplate struct {
	templates     map[string]*template.Template
	templateFuncs template.FuncMap
}

func NewMyTextTemplate() *MyTextTemplate {
	return &MyTextTemplate{
		templates:     make(map[string]*template.Template),
		templateFuncs: make(template.FuncMap),
	}
}

func simplePrint(in string) string {
	return fmt.Sprintf("in=%s", in)
}

func (this *MyTextTemplate) AddTemplate(name, str string) (err error) {
	this.templates[name], err = template.New(name).Funcs(this.templateFuncs).Parse(str)
	return
}

func (this *MyTextTemplate) AddTempFunc(name string, f interface{}) {
	this.templateFuncs[name] = f
}

func (this *MyTextTemplate) String(name string, data interface{}) (str string, err error) {
	temp, ok := this.templates[name]
	if !ok {
		err = errors.New("no have this template name")
		return
	}

	a.Reset()
	err = temp.Execute(a, data)
	if err != nil {
		return
	}
	return a.String(), err
}

var (
	a *bytes.Buffer
)

func init() {
	a = bytes.NewBuffer([]byte{})
}
