package en_NF

import "github.com/ContextLogic/cldr"

var Locale = &cldr.Locale{
	Locale: "en_NF",
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
