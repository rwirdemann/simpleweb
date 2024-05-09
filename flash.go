package simpleweb

import (
	"fmt"
	"html/template"
)

var flash *Flash

func init() {
	flash = new(Flash)
}

type Flash struct {
	info    string
	error   string
	warning string
}

func (f *Flash) Info() string {
	return f.info
}

func (f *Flash) Error() string {
	return f.error
}

func (f *Flash) Warning() string {
	return f.warning
}

func (f *Flash) ClearInfo() {
	f.info = ""
}

func (f *Flash) ClearWarning() {
	f.warning = ""
}

func (f *Flash) ClearError() {
	f.error = ""
}

func Info(s string) {
	flash.info = s
}

func Error(s string) {
	flash.error = s
}

func Warning(s string) {
	flash.warning = s
}

var flashHelper = template.FuncMap{
	"info": func() template.HTML {
		s := ""
		if len(flash.Info()) > 0 {
			s = fmt.Sprintf("<p style=\"color: green\">%s</p>", flash.Info())
			flash.ClearInfo()
		}
		return template.HTML(s)
	},
	"hasInfo": func() bool {
		if len(flash.Info()) > 0 {
			return true
		}
		return false
	},
	"warning": func() template.HTML {
		s := ""
		if len(flash.Warning()) > 0 {
			s = fmt.Sprintf("<p style=\"color: orange\">%s</p>", flash.Warning())
			flash.ClearWarning()
		}
		return template.HTML(s)
	},
	"error": func() template.HTML {
		s := ""
		if len(flash.Error()) > 0 {
			s = fmt.Sprintf("<p style=\"color: red\">%s</p>", flash.Error())
			flash.ClearError()
		}
		return template.HTML(s)
	}}
