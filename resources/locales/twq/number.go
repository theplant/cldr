package twq

import "github.com/ContextLogic/cldr"

var (
	symbols = cldr.Symbols{Decimal: ".", Group: "\u00a0", Negative: "", Percent: "", PerMille: ""}
	formats = cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00¤", Percent: "#,##0%"}
)
