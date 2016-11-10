package main

import "os"
import "html/template"
import "fmt"

// Ejemplo de como renderizar un template dentro de un archivo

type Inventory struct {
	Material string
	Count    uint
}

func main(){
	// Instancia del inventario
	sweaters := Inventory{"Sweater Color rojo", 17}

	// Cargan en memoria el template
	tmpl, err := template.ParseFiles("test.html", "aux.tpl")
	if err != nil { panic(err) }

	// Destino del render
	fmt.Println("Creando archivo")
	f, error := os.Create("render.html")

	if error != nil {
		fmt.Println("Hubo un error al crear el archivo", error)
	}

	// Renderizando y guardando el template
	err = tmpl.Execute(f, sweaters)
	if err != nil { panic(err) }
}


