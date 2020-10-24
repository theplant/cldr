// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
    "{{.CLDRPackage}}"
    "{{.CLDRPackage}}/resources/currency"
)

func Get_{{.LocaleCode}}() *cldr.Locale {
    return &cldr.Locale {
        Locale: "{{.LocaleCode}}",
        Calendar: {{printf "%#v" .Calendar}},
        Number: cldr.Number{
            Symbols: {{printf "%#v" .Symbols}},
            Formats: {{printf "%#v" .NumberFormats}},
            Currencies: cldr.Currencies{
                {{range $curCode, $currency := .Currencies -}}
                currency.{{$curCode}}: {{printf "%#v" $currency}},
                {{end}}
            },
        },
    }
}