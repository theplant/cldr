package en_BE

import "github.com/ContextLogic/cldr"

var (
	symbols = cldr.Symbols{Decimal: ",", Group: ".", Negative: "", Percent: "", PerMille: ""}
	formats = cldr.NumberFormats{Decimal: "", Currency: "#,##0.00\u00a0¤", Percent: ""}
)