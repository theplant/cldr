package ms_Latn_BN

import "github.com/ContextLogic/cldr"

var (
	symbols = cldr.Symbols{Decimal: ",", Group: ".", Negative: "", Percent: "", PerMille: ""}
	formats = cldr.NumberFormats{Decimal: "", Currency: "¤\u00a0#,##0.00", Percent: ""}
)
