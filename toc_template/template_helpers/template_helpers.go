package template_helpers


func T9n(text string, locale string, sourceLocale string, enc string) string {
	return text
}

func GetUnitReferenceValue(item string) string {
	return item
}

func GetFormattedUnitReferenceValue(unit_ref string) string {
	return unit_ref
}

// Simulacion de la estructura usada en los templates gsp
type CheckoutDto_struct struct {
	Payment bool
}

type Vars_struct struct {
	CheckoutDto CheckoutDto_struct
	ItemSite string
	LocaleString string
	Quantity int
	UnitPrice string
	ItemTitle string
	ItemPicture string
	Item string
}

func (c Vars_struct) Caca(s string) string {
	return "hola"
}