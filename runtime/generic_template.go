package runtime

import (
	"html/template"
	"io/fs"
	"strings"
)

func TemplateParse[T any](content fs.FS, templateFile string, obj T) (output string, err error) {
	tmpl, err := template.ParseFS(content, templateFile)
	if err != nil {
		return output, err
	}
	buf := new(strings.Builder)
	err = tmpl.Execute(buf, obj)
	if err != nil {
		return output, err
	}
	output = buf.String()
	return output, nil
}
