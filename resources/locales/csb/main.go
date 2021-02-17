package csb

import "github.com/ContextLogic/cldr"

var Locale = &cldr.Locale{
	Locale: "csb",
	PluralRule: pluralRule,
}

func init() {
	cldr.RegisterLocale(Locale)
}
