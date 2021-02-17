package lo

import "github.com/ContextLogic/cldr"

var (
	symbols = cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"}
	formats = cldr.NumberFormats{Decimal: "0ພັນ", Currency: "¤#,##0.00;¤-#,##0.00", Percent: "#,##0%"}
)
