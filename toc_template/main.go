package main

import(
	"fmt"
	"net/http"
	"html/template"
	"github.com/segeschecho/toc_template/template_helpers"
	"github.com/aymerick/raymond"
)

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
	tpl := `
		<div class="entry">
  		<h1>{{title}}</h1>
  		<div class="body">
    			{{body}}
  		</div>
		</div>
	`

	ctx := map[string]string{
		"title": "My New Post",
		"body":  "This is my first post!",
	}

	result, err := raymond.Render(tpl, ctx)
	if err != nil {
		panic("Please fill a bug :)")
	}

	fmt.Fprintf(w, result)
}

func main() {
	fmt.Println("Servicio ejecutando...")

	// Se instancia y se levanta el server
	http.HandleFunc("/default", default_templating_handler)
	http.HandleFunc("/handlebars", handlebars_templating_handler)
	http.ListenAndServe(":8080", nil)

}
