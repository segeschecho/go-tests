package service_defs

import (
	"testing"
)

var expect_render = `
	<html>
	<body>
	<p><h1>Nombre: Sergio</h1></p>
	<p><h1>Apellido: Gonzalez</h1></p>
	</body>
	</html>
`

func Test_render_service(t *testing.T){
	name := "Sergio"
	surname := "Gonzalez"

	render := Render_service(name, surname)

	if expect_render != render {
		t.Error("El render es distinto al esperado")
	}
}