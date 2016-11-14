package main

import(
	"fmt"
	"net/http"
	"html/template"
	"github.com/segeschecho/go-tests/toc_template/template_helpers"
	"github.com/aymerick/raymond"
)

// Ejemplo de como se manejan templates usando el framework que viene con Go.

// Mock de la estructura usada dentro de los templates.
var model = template_helpers.Vars_struct{
	CheckoutDto: template_helpers.CheckoutDto_struct{Payment: true},
	ItemSite: "http://static.mlstatic.com/org-img/commons/logo-mercadolibre-new.png",
	LocaleString: "LocaleStringValor",
	Quantity: 2,
	UnitPrice: "$200",
	ItemTitle: "Limpiador A Vapor Mopa Steam Cleaner 5 En 1 Garantia 3 Meses",
	ItemPicture: "https://mla-s2-p.mlstatic.com/294405-MLA20863199869_082016-I.jpg",
	Item: "",
}

// funciones para manejar los requests HTTP
func default_templating_handler(w http.ResponseWriter, r *http.Request) {
	// Defer funciona como un "finally", aca se usa para manejar runtime errors
	defer func(){
		r := recover()
		if r != nil {
			fmt.Fprintf(w, "Error:", r)
		}
	}()

	tmpl_base := template.New("base_template")

	// Se dan de alta las funciones que se usan dentro de los templates
	tmpl_base.Funcs(template.FuncMap{"t9n": template_helpers.T9n})
	tmpl_base.Funcs(template.FuncMap{"getUnitReferenceValue": template_helpers.GetUnitReferenceValue})
	tmpl_base.Funcs(template.FuncMap{"getFormattedUnitReferenceValue": template_helpers.GetFormattedUnitReferenceValue})

	// Los sub templates y el main template se cachean
	tmpl_base = template.Must(tmpl_base.ParseFiles(
		"web_templates/common/CHO_PAGO_ENVIO.html",
		"web_templates/templates/commons/common_header.html",
		"web_templates/templates/commons/title/title_congrats.html",
		"web_templates/templates/commons/item_info.html",
		"web_templates/templates/commons/common_footer.html"))

	// Se renderiza el template
	err := tmpl_base.ExecuteTemplate(w, "CHO_PAGO_ENVIO.html", model)

	if err != nil {
		fmt.Fprintf(w, "Error ejecutando el template" + err.Error())
	}
}


func handlebars_templating_handler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Fprintf(w, "Error: ", r)
		}
	}()

	// Estructura a usar dentro de los templates
	type model_struct struct {
		Title string
		Body string
		List_test []string
	}


	// Instanciando la estructura con datos de prueba
	var model = model_struct{
		Title: "My New Post",
		Body:  "This is my first post!",
		List_test: []string{"1","2","3"},
	}

	tpl, err := raymond.ParseFile("web_templates/handlebars/test_main.html")

	if err != nil{
		fmt.Println(err.Error())
	}

	tpl.RegisterPartialFiles("web_templates/handlebars/test_foo_test.html")

	// Se aplica el modelo a la vista
	result, err := tpl.Exec(model)

	if err != nil{
		fmt.Println(err.Error())
	}

	// Se renderiza en la respuesta http
	fmt.Fprintf(w, result)
}

func main() {
	fmt.Println("Servicio ejecutando...")

	// Se instancia y se levanta el server
	http.HandleFunc("/default", default_templating_handler)
	http.HandleFunc("/handlebars", handlebars_templating_handler)
	http.ListenAndServe(":8080", nil)

}
