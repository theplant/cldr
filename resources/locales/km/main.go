package km

import "github.com/ContextLogic/cldr"

var Locale = &cldr.Locale{
	Locale: "km",
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