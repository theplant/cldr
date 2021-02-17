package ak

import "github.com/ContextLogic/cldr"

var Locale = &cldr.Locale{
	Locale: "ak",
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
