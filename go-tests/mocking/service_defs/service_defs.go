package service_defs

import (
	"bytes"
	"html/template"
)

type data_struct struct{
	Name string
	Surname string
}

var tmpl_def = `
	<html>
	<body>
	<p><h1>Nombre: {{.Name}}</h1></p>
	<p><h1>Apellido: {{.Surname}}</h1></p>
	</body>
	</html>
`

func Render_service(name string, surname string) string {

	tmpl := template.New("service_template")
	tmpl.Parse(tmpl_def)

	data := data_struct{Name: name, Surname: surname}
	var out bytes.Buffer

	tmpl.Execute(&out, data)

	return out.String()
}
