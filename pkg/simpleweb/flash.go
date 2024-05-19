package simpleweb

import (
	"html/template"
)

var infoFlash, warningFlash, errorFlash string

func Info(s string) {
	infoFlash = s
}

func Error(s string) {
	errorFlash = s
}

func Warning(s string) {
	warningFlash = s
}

var flashHelper = template.FuncMap{
	"info": func() template.HTML {
		s := infoFlash
		infoFlash = ""
		return template.HTML(s)
	},
	"hasInfo": func() bool {
		if len(infoFlash) > 0 {
			return true
		}
		return false
	},
	"hasWarning": func() bool {
		if len(warningFlash) > 0 {
			return true
		}
		return false
	},
	"warning": func() template.HTML {
		s := warningFlash
		warningFlash = ""
		return template.HTML(s)
	},
	"hasError": func() bool {
		if len(errorFlash) > 0 {
			return true
		}
		return false
	},
	"error": func() template.HTML {
		s := errorFlash
		errorFlash = ""
		return template.HTML(s)
	},
}
