package es_DO

import "github.com/ContextLogic/cldr"

var Locale = &cldr.Locale{
	Locale: "es_DO",
	Number: cldr.Number{
		Symbols:    symbols,
		Formats:    formats,
		Currencies: currencies,
	},
	Calendar:   calendar,
	PluralRule: pluralRule,
}

func init() {
	cldr.RegisterLocale(Locale)
}
