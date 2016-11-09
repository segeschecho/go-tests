package template_helpers

import (
	"testing"
)


func Test(t *testing.T){
	type t9n_struct struct {
		text, locale, sourceLocale, enc, want string
	}

	var t9n_cases = []t9n_struct{
		{"text1", "a", "b", "c", "text1"},
		{"text2", "a", "b", "c", "text2"},
		{"text3", "a", "b", "c", "text3"},
	}

	for _, c := range t9n_cases{
		var res = T9n(c.text, c.locale, c.sourceLocale, c.enc)

		if res != c.want{
			t.Errorf("Error ejecutando la funcion t9n: %s", c)
		}
	}
}


func Test2(t *testing.T) {
	type get_unit_struct struct {
		item, want string
	}

	var get_unit_cases = []get_unit_struct{
		{"item1", "item1"},
		{"item2", "item2"},
	}

	for _, c := range get_unit_cases{
		var res = GetUnitReferenceValue(c.item)

		if res != c.want{
			t.Error("Error ejecutando la función GetUnitReferenceValue: %s", c)
		}
	}
}