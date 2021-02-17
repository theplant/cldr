package en_TV

import "github.com/ContextLogic/cldr"

var Locale = &cldr.Locale{
	Locale: "en_TV",
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
