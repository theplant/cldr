package fr_DJ

import "github.com/ContextLogic/cldr"

var Locale = &cldr.Locale{
	Locale: "fr_DJ",
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