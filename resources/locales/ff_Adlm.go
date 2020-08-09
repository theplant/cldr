// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_ff_Adlm() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y MMMM d, EEEE", Long: "y MMMM d", Medium: "y MMM d", Short: "y-MM-dd"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1}{0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "𞤑𞤖𞤏{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "𞤅𞤭𞥅𞤤", Feb: "𞤕𞤮𞤤", Mar: "𞤄𞤮𞥅𞤴", Apr: "𞤅𞤫𞥅𞤼", May: "𞤁𞤵𞥅𞤶", Jun: "𞤑𞤮𞤪", Jul: "𞤃𞤮𞤪", Aug: "𞤔𞤵𞤳", Sep: "𞤅𞤭𞤤", Oct: "𞤒𞤢𞤪", Nov: "𞤔𞤮𞤤", Dec: "𞤄𞤮𞤱"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "𞤅", Feb: "𞤕", Mar: "𞤄", Apr: "𞤅", May: "𞤁", Jun: "𞤑", Jul: "𞤃", Aug: "𞤔", Sep: "𞤅", Oct: "𞤒", Nov: "𞤄", Dec: "𞤄"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "𞤅𞤭𞥅𞤤𞤮", Feb: "𞤕𞤮𞤤𞤼𞤮", Mar: "𞤐𞥋𞤄𞤮𞥅𞤴𞤮", Apr: "𞤅𞤫𞥅𞤼𞤮", May: "𞤁𞤵𞥅𞤶𞤮", Jun: "𞤑𞤮𞤪𞤧𞤮", Jul: "𞤃𞤮𞤪𞤧𞤮", Aug: "𞤔𞤵𞤳𞤮", Sep: "𞤅𞤭𞤤𞤼𞤮", Oct: "𞤒𞤢𞤪𞤳𞤮", Nov: "𞤔𞤮𞤤𞤮", Dec: "𞤐𞥋𞤄𞤮𞤱𞤼𞤮"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "𞤈𞤫𞤬", Mon: "𞤀𞥄𞤩𞤵", Tue: "𞤃𞤢𞤦", Wed: "𞤔𞤫𞤧", Thu: "𞤐𞤢𞥄𞤧", Fri: "𞤃𞤢𞤣", Sat: "𞤖𞤮𞤪"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "𞤈", Mon: "𞤀𞥄", Tue: "𞤃", Wed: "𞤔", Thu: "𞤐", Fri: "𞤃", Sat: "𞤖"}, Short: cldr.CalendarDayFormatNameValue{Sun: "𞤈𞤫𞤬", Mon: "𞤀𞥄𞤩𞤵", Tue: "𞤃𞤢𞤦", Wed: "𞤔𞤫𞤧", Thu: "𞤐𞤢𞥄𞤧", Fri: "𞤃𞤢𞤣", Sat: "𞤖𞤮𞤪"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "𞤈𞤫𞤬𞤦𞤭𞤪𞥆𞤫", Mon: "𞤀𞥄𞤩𞤵𞤲𞥋𞤣𞤫", Tue: "𞤃𞤢𞤱𞤦𞤢𞥄𞤪𞤫", Wed: "𞤐𞥋𞤔𞤫𞤧𞤤𞤢𞥄𞤪𞤫", Thu: "𞤐𞤢𞥄𞤧𞤢𞥄𞤲𞥋𞤣𞤫", Fri: "𞤃𞤢𞤱𞤲𞥋𞤣𞤫", Sat: "𞤖𞤮𞤪𞤦𞤭𞤪𞥆𞤫"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "𞤀𞤎", PM: "𞤇𞤎"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "𞤀𞤎", PM: "𞤇𞤎"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "𞤀𞤎", PM: "𞤇𞤎"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_root]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "", Group: "", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a00\u00a0𞤘𞤵𞤤", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "𞤁𞤭𞤪𞤸𞤢𞤥𞤵 𞤋𞤥𞤢𞥄𞤪𞤢𞤼𞤭𞤲𞤳𞤮", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "𞤀𞤬𞤿𞤢𞤲𞤭 𞤀𞤬𞤿𞤢𞤲𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "𞤂𞤫𞤳 𞤀𞤤𞤦𞤢𞤲𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "𞤁𞤢𞤪𞤢𞤥𞤵 𞤀𞤪𞤥𞤢𞤲𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.ANG: cldr.Currency{DisplayName: "𞤊𞤵𞤤𞤮𞤪𞤭𞤲 𞤀𞤲𞤼𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "𞤑𞤵𞤱𞤢𞤲𞥁𞤢 𞤀𞤲𞤺𞤮𞤤𞤢𞤲𞤳𞤮", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "𞤆𞤫𞤧𞤮 𞤀𞤪𞤶𞤢𞤲𞤼𞤭𞤲𞤢", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "𞤁𞤢𞤤𞤢 𞤌𞤧𞤼𞤪𞤢𞤤𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "𞤊𞤵𞤤𞤮𞤪𞤭𞤲 𞤀𞤪𞤵𞤦𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "𞤃𞤢𞤲𞤢𞥄𞤼𞤵 𞤀𞥁𞤫𞤪𞤦𞤢𞤴𞤶𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "𞤃𞤢𞤪𞤳 𞤄𞤮𞤧𞤲𞤭𞤴𞤢-𞤖𞤫𞤪𞤶𞤫𞤺𞤮𞤾𞤭𞤲𞤳𞤮 𞤱𞤢𞤴𞤤𞤮𞤼𞤮𞥅𞤯𞤭", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "𞤁𞤢𞤤𞤢 𞤄𞤢𞤪𞤦𞤢𞤣𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "𞤚𞤢𞤪𞤢 𞤄𞤢𞤲𞤺𞤭𞤤𞤢𞤣𞤫𞥅𞤧𞤭𞤲𞤳𞤮", Symbol: "৳"},
				currency.BGN: cldr.Currency{DisplayName: "𞤂𞤫𞥅𞤾 𞤄𞤭𞤤𞤺𞤢𞤪𞤭𞤲𞤳𞤮", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "𞤁𞤭𞤲𞤢𞥄𞤪 𞤄𞤢𞤸𞤢𞤪𞤢𞥄𞤲𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "𞤊𞤢𞤪𞤢𞤲 𞤄𞤵𞤪𞤵𞤲𞤣𞤭𞤲𞤳𞤮", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "𞤁𞤢𞤤𞤢 𞤄𞤫𞤪𞤥𞤵𞤣𞤢𞥄𞤲", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "𞤁𞤢𞤤𞤢 𞤄𞤵𞤪𞤲𞤫𞤴𞤢𞤲𞤳𞤮", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "𞤄𞤮𞤤𞤭𞤾𞤭𞤴𞤢𞤲𞤮 𞤄𞤮𞤤𞤭𞤾𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "𞤈𞤭𞤴𞤢𞤤 𞤄𞤪𞤢𞤧𞤭𞤤𞤴𞤢𞤲𞤳𞤮", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "𞤁𞤢𞤤𞤢 𞤄𞤢𞤸𞤢𞤥𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "𞤐𞤘𞤵𞤤𞤼𞤵𞤪𞤵𞤥𞤵 𞤄𞤵𞤼𞤢𞤲𞤭𞤲𞤳𞤮", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "𞤆𞤵𞤤𞤢 𞤄𞤮𞤼𞤵𞤧𞤱𞤢𞤲𞤢𞤲𞤳𞤮", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "𞤈𞤵𞥅𞤦𞤮𞤤 𞤄𞤫𞤤𞤢𞤪𞤭𞥅𞤧𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "BYN"},
				currency.BZD: cldr.Currency{DisplayName: "𞤁𞤢𞤤𞤢 𞤄𞤫𞤤𞤭𞥅𞤧", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "𞤁𞤢𞤤𞤢 𞤑𞤢𞤲𞤢𞤣𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "𞤊𞤢𞤪𞤢𞤲 𞤑𞤮𞤲𞤺𞤮𞤲𞤳𞤮", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "𞤊𞤢𞤪𞤢𞤲 𞤅𞤵𞤱𞤭𞥅𞤧", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "𞤆𞤫𞤧𞤮 𞤕𞤭𞤤𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "$"},
				currency.CNH: cldr.Currency{DisplayName: "𞤒𞤵𞤱𞤢𞤲 𞤕𞤢𞤴𞤲𞤭𞤲𞤳𞤮 (𞤺𞤢𞥄𞤲𞤭𞤲𞤳𞤮)", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "𞤒𞤵𞤱𞤢𞤲 𞤕𞤢𞤴𞤲𞤭𞤲𞤳𞤮", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "𞤆𞤫𞤧𞤮 𞤑𞤮𞤤𞤮𞤥𞤦𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "𞤑𞤮𞤤𞤮𞥅𞤲 𞤑𞤮𞤧𞤼𞤢 𞤈𞤭𞤳𞤢𞤲", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "𞤆𞤫𞤧𞤮 𞤑𞤵𞤦𞤢𞤲𞤳𞤮 𞤏𞤢𞤴𞤤𞤮𞤼𞤮𞥅𞤲𞥋𞤺𞤮", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "𞤆𞤫𞤧𞤮 𞤑𞤵𞤦𞤢𞤲𞤳𞤮", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "𞤉𞤧𞤳𞤵𞤣𞤮 𞤑𞤢𞤨-𞤜𞤫𞥅𞤪𞤣𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "𞤑𞤮𞤪𞤵𞤲𞤢 𞤕𞤫𞥅𞤳𞤭𞤲𞤳𞤮", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "𞤊𞤢𞤪𞤢𞤲 𞤔𞤭𞤦𞤵𞤼𞤭𞤲𞤳𞤮", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "𞤑𞤮𞤪𞤲𞤫 𞤁𞤢𞤲𞤭𞥅𞤧𞤭𞤲𞤳𞤮", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "𞤆𞤫𞤧𞤮 𞤁𞤮𞤥𞤭𞤲𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "𞤁𞤭𞤲𞤢𞥄𞤪 𞤀𞤤𞤶𞤢𞤪𞤭𞤲𞤳𞤮", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "𞤆𞤢𞤱𞤲𞤣𞤵 𞤃𞤭𞤧𞤭𞤪𞤢𞤲𞤳𞤮", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "𞤐𞤢𞤳𞤬𞤢 𞤉𞤪𞤭𞤼𞤫𞤪𞤭𞤲𞤳𞤮", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "𞤄𞤭𞤪 𞤖𞤢𞤦𞤢𞤧𞤭𞤲𞤳𞤮", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "𞤒𞤵𞤪𞤮𞥅", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "𞤁𞤢𞤤𞤢 𞤊𞤭𞤶𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "𞤆𞤢𞤱𞤲𞤣𞤵 𞤅𞤵𞤪𞤭𞥅𞤶𞤫 𞤊𞤢𞤤𞤳𞤵𞤤𞤢𞤲𞤣𞤭𞤳𞤮", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "𞤆𞤢𞤱𞤲𞤣𞤵 𞤄𞤪𞤭𞤼𞤭𞥅𞤧𞤭𞤲𞤳𞤮", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "𞤂𞤢𞥄𞤪𞤭 𞤔𞤮𞤪𞤶𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "𞤅𞤭𞤣𞤭 𞤘𞤢𞤲𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "𞤆𞤢𞤱𞤲𞥋𞤣𞤵 𞤔𞤭𞤤𞤦𞤪𞤢𞤤𞤼𞤢𞤪", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "𞤁𞤢𞤤𞤢𞤧𞤭 𞤘𞤢𞤥𞤦𞤭𞤲𞤳𞤮", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "𞤊𞤢𞤪𞤢𞤲 𞤘𞤭𞤲𞤫𞤲𞤳𞤮", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "𞤑𞤫𞤼𞤵𞥁𞤢𞤤 𞤘𞤵𞤱𞤢𞤼𞤫𞤥𞤢𞤤𞤢𞤲𞤳𞤮", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "𞤁𞤢𞤤𞤢 𞤘𞤵𞤴𞤢𞤲𞤫𞥅𞤧𞤭𞤲𞤳𞤮", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "𞤁𞤢𞤤𞤢 𞤖𞤮𞤲𞤳𞤮𞤲", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "𞤂𞤫𞤥𞤨𞤭𞤪𞤢 𞤖𞤮𞤲𞤣𞤵𞤪𞤢𞤲𞤳𞤮", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "𞤑𞤵𞤲𞤢 𞤑𞤵𞤪𞤢𞥄𞤧𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "𞤘𞤵𞥅𞤪𞤣𞤫 𞤖𞤢𞤴𞤼𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "𞤊𞤮𞤪𞤭𞤲𞤼𞤵 𞤖𞤵𞤲𞤺𞤢𞤪𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "𞤈𞤵𞤨𞤭𞤴𞤢 𞤋𞤲𞤣𞤮𞤲𞤫𞤧𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "𞤡𞤫𞤳𞤫𞤤 𞤋𞤧𞤪𞤢𞥄𞤤𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "𞤈𞤵𞥅𞤨𞤭𞥅 𞤖𞤭𞤲𞤣𞤭𞤧𞤼𞤢𞤲𞤳𞤮", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "𞤁𞤭𞤲𞤢𞥄𞤪 𞤋𞤪𞤢𞥄𞤳𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "𞤈𞤭𞤴𞤢𞥄𞤤 𞤋𞤪𞤢𞤲𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "𞤑𞤮𞤪𞤮𞤲𞤢 𞤀𞤴𞤧𞤭𞤤𞤢𞤲𞤣𞤭𞤲𞤳𞤮", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "𞤁𞤢𞤤𞤢 𞤔𞤢𞤥𞤢𞤴𞤭𞤲𞤳𞤮", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "𞤁𞤭𞤲𞤢𞥄𞤪 𞤔𞤮𞤪𞤣𞤢𞤲𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "𞤒𞤫𞤲 𞤔𞤢𞤨𞤢𞤲𞤳𞤮", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "𞤅𞤭𞤤𞤭𞤲 𞤑𞤫𞤲𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "𞤅𞤮𞤥𞤵 𞤑𞤭𞤪𞤺𞤭𞤧𞤼𞤢𞤲𞤭𞤲𞤳𞤮", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "𞤈𞤭𞤴𞤢𞤤 𞤑𞤢𞤥𞤦𞤮𞤣𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "𞤊𞤢𞤪𞤢𞤲 𞤑𞤮𞤥𞤮𞤪𞤭𞤲𞤳𞤮", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "𞤏𞤮𞤲 𞤁𞤮𞤱𞤣𞤮𞤱𞤪𞤭 𞤑𞤮𞥅𞤪𞤫𞤴𞤢𞤲𞤳𞤮", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "𞤱𞤮𞤲 𞤂𞤫𞤴𞤤𞤫𞤴𞤪𞤭 𞤑𞤮𞥅𞤪𞤫𞤴𞤢𞤲𞤳𞤮", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "𞤁𞤋𞤲𞤢𞥄𞤪 𞤑𞤵𞤱𞤢𞤴𞤼𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "𞤁𞤢𞤤𞤢 𞤅𞤵𞤪𞤭𞥅𞤶𞤫 𞤑𞤢𞤴𞤥𞤢𞥄𞤲", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "𞤚𞤫𞤲𞤺𞤫 𞤑𞤢𞥁𞤢𞤳𞤭𞤧𞤼𞤢𞤲𞤭𞤲𞤳𞤮", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "𞤑𞤭𞤨𞤵 𞤂𞤢𞤱𞤮𞥅𞤧𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "𞤆𞤢𞤱𞤲𞥋𞤣𞤵 𞤂𞤭𞤦𞤢𞤲𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "𞤈𞤵𞥅𞤨𞤭𞥅 𞤅𞤭𞤪𞤭-𞤂𞤢𞤲𞤳𞤢𞤲𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "𞤁𞤢𞤤𞤢 𞤂𞤭𞤦𞤫𞤪𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "$"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "𞤁𞤭𞤲𞤢𞥄𞤪 𞤂𞤭𞤦𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "𞤁𞤭𞤪𞤸𞤢𞤥𞤵 𞤃𞤮𞤪𞤮𞤳𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "𞤂𞤭𞥅𞤱𞤮 𞤃𞤮𞤤𞤣𞤮𞤾𞤢𞤲𞤳𞤮", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "𞤀𞤪𞤭𞤴𞤢𞤪𞤭 𞤃𞤢𞤤𞤺𞤢𞤲𞤭𞤲𞤳𞤮", Symbol: "Ar"},
				currency.MKD: cldr.Currency{DisplayName: "𞤁𞤭𞤲𞤢𞥄𞤪 𞤃𞤢𞤧𞤫𞤣𞤮𞤲𞤭𞤲𞤳𞤮", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "𞤑𞤭𞤴𞤢𞤼𞤵 𞤃𞤭𞤴𞤢𞤥𞤢𞤪𞤭𞤲𞤳𞤮", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "𞤚𞤵𞤺𞤵𞤪𞤭𞤳𞤵 𞤃𞤮𞤲𞤺𞤮𞤤𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "𞤆𞤢𞤼𞤢𞤳𞤢 𞤃𞤢𞤳𞤢𞤱𞤮𞤴𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.MRO: cldr.Currency{DisplayName: "𞤓𞤺𞤭𞤴𞤢 𞤃𞤮𞤪𞤭𞤼𞤢𞤲𞤭𞤴𞤢𞤲𞤳𞤮 (𞥑𞥙𞥗𞥓 - 𞥒𞥐𞥑𞥗)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "𞤓𞤺𞤭𞤴𞤢 𞤃𞤮𞤪𞤭𞤼𞤢𞤲𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "𞤈𞤵𞤨𞤭𞥅 𞤃𞤮𞤪𞤭𞤧𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "𞤈𞤵𞤬𞤭𞤴𞤢𞥄 𞤃𞤢𞤤𞤣𞤭𞤾𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "𞤑𞤢𞤱𞤢𞤷𞤢 𞤃𞤢𞤤𞤢𞤱𞤭𞤲𞤳𞤮", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "𞤆𞤫𞤧𞤮 𞤃𞤫𞤳𞤧𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "𞤈𞤭𞤲𞤺𞤵𞤼𞤵 𞤃𞤢𞤤𞤫𞥅𞤧𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "RM"},
				currency.MZN: cldr.Currency{DisplayName: "𞤃𞤫𞤼𞤭𞤳𞤮𞤤 𞤃𞤮𞥁𞤢𞤥𞤦𞤭𞤲𞤳𞤮", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "𞤁𞤢𞤤𞤢 𞤐𞤢𞤥𞤭𞤥𞤦𞤭𞤲𞤳𞤮", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "𞤐𞤢𞤴𞤪𞤢 𞤐𞤢𞤶𞤭𞤪𞤢𞤴𞤢𞤲𞤳𞤮", Symbol: "𞤐𞤐𞤘"},
				currency.NIO: cldr.Currency{DisplayName: "𞤑𞤮𞥅𞤪𞤣𞤮𞤦𞤢 𞤐𞤭𞤳𞤢𞤪𞤢𞤺𞤵𞤱𞤢𞤲𞤳𞤮", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "𞤑𞤪𞤮𞤲𞤫 𞤐𞤮𞤪𞤱𞤫𞤶𞤭𞤲𞤳𞤮", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "𞤈𞤵𞥅𞤨𞤭𞥅 𞤐𞤫𞤨𞤢𞤤𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "𞤁𞤢𞤤𞤢 𞤐𞤫𞤱-𞤔𞤭𞤤𞤢𞤲𞤣𞤭𞤲𞤳𞤮", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "𞤈𞤭𞤴𞤢𞥄𞤤 𞤌𞤥𞤢𞤲𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "𞤄𞤢𞤤𞤦𞤮𞤱𞤢 𞤆𞤢𞤲𞤢𞤥𞤢𞤴𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "𞤅𞤮𞤤 𞤆𞤫𞤪𞤵𞤲𞤳𞤮", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "𞤑𞤭𞤲𞤢 𞤆𞤢𞤨𞤵𞤱𞤢 𞤐𞤫𞤱-𞤘𞤭𞤲𞤫𞤲𞤳𞤮", Symbol: "𞤑𞤆𞤘"},
				currency.PHP: cldr.Currency{DisplayName: "𞤆𞤭𞤧𞤮 𞤊𞤭𞤤𞤭𞤨𞥆𞤭𞤲𞤳𞤮", Symbol: "𞤆𞤆𞤖"},
				currency.PKR: cldr.Currency{DisplayName: "𞤈𞤵𞥅𞤨𞤭𞥅 𞤆𞤢𞤳𞤭𞤧𞤼𞤢𞤲𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "𞤔𞤢𞤤𞤮𞤼𞤵 𞤆𞤮𞤤𞤭𞥅𞤧𞤭𞤲𞤳𞤮", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "𞤘𞤵𞤱𞤢𞤪𞤢𞤲𞤭 𞤆𞤢𞥄𞤪𞤢𞤺𞤵𞤴𞤫𞤲𞤳𞤮", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "𞤈𞤭𞤴𞤢𞥄𞤤 𞤗𞤢𞤼𞤢𞤪𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "𞤂𞤫𞤱𞤵 𞤈𞤮𞤥𞤢𞤲𞤭𞤲𞤳𞤮", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "𞤁𞤭𞤲𞤢𞥄𞤪 𞤅𞤫𞤪𞤦𞤭𞤲𞤳𞤮", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "𞤈𞤵𞥅𞤦𞤮𞤤 𞤈𞤭𞥅𞤧𞤭𞤲𞤳𞤮", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "𞤊𞤢𞤪𞤢𞤲 𞤈𞤵𞤱𞤢𞤲𞤣𞤢𞤲𞤳𞤮", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "𞤈𞤭𞤴𞤢𞤤 𞤅𞤢𞤵𞥅𞤣𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "𞤁𞤢𞤤𞤢 𞤅𞤵𞤪𞤭𞥅𞤶𞤫 𞤅𞤵𞤤𞤫𞤴𞤥𞤢𞥄𞤲𞤭𞤲𞤳𞤮", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "𞤈𞤵𞤨𞤭𞥅 𞤅𞤫𞤴𞤧𞤭𞤤𞤭𞤲𞤳𞤮", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "𞤆𞤢𞤱𞤲𞤣𞤵 𞤅𞤵𞤣𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "𞤑𞤪𞤮𞤲𞤢 𞤅𞤵𞤱𞤫𞤣𞤭𞤲𞤳𞤮", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "𞤁𞤢𞤤𞤢 𞤅𞤭𞤲𞤺𞤢𞤨𞤮𞤪𞤫𞤲𞤳𞤮", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "𞤆𞤢𞤱𞤲𞤣𞤵 𞤅𞤫𞤲-𞤖𞤫𞤤𞤫𞤲𞤢", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "𞤂𞤫𞤴𞤮𞤲 𞤅𞤫𞤪𞤢𞤤𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "𞤅𞤭𞤤𞤭𞤲 𞤅𞤮𞤥𞤢𞤤𞤭𞤲𞤳𞤮", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "𞤁𞤢𞤤𞤢 𞤅𞤵𞤪𞤵𞤲𞤢𞤥𞤭𞤲𞤳𞤮", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "𞤆𞤢𞤱𞤲𞤣𞤵 𞤂𞤫𞤴𞤤𞤫𞤴𞤪𞤭 𞤅𞤵𞤣𞤢𞤲𞤭𞤲𞤳𞤮", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "𞤁𞤮𞤦𞤢𞤪𞤢 𞤅𞤢𞤱𞤮-𞤚𞤮𞤥𞤫 & 𞤆𞤫𞤪𞤫𞤲𞤧𞤭𞤨", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "𞤆𞤢𞤱𞤲𞥋𞤣𞤵 𞤅𞤭𞤪𞤢𞤴𞤢𞤲𞤳𞤮", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "𞤂𞤭𞤤𞤢𞤲𞤺𞤫𞤲𞤭 𞤅𞤵𞤱𞤢𞤶𞤭", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "𞤄𞤢𞤸𞤼𞤵 𞤚𞤢𞤴𞤤𞤢𞤲𞤣𞤭𞤲𞤳𞤮", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "𞤅𞤢𞤥𞤮𞥅𞤲𞤭 𞤚𞤢𞤶𞤭𞤳𞤭𞤧𞤼𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "𞤃𞤢𞤲𞤢𞤼𞤵 𞤚𞤵𞤪𞤳𞤵𞤥𞤫𞤲𞤭𞤧𞤼𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "𞤁𞤭𞤲𞤢𞥄𞤪 𞤚𞤵𞥅𞤲𞤭𞤧𞤭𞤲𞤳𞤮", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "𞤆𞤢𞤢𞤲𞤺𞤢 𞤚𞤮𞤲𞤺𞤢𞤲𞤳𞤮", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "𞤂𞤭𞤪𞤢 𞤚𞤵𞤪𞤳𞤭𞤴𞤢𞤲𞤳𞤮", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "𞤁𞤢𞤤𞤢 𞤚𞤭𞤪𞤲𞤭𞤣𞤢𞥄𞤣 & 𞤚𞤮𞤦𞤢𞤺𞤮", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "𞤁𞤢𞤤𞤢 𞤚𞤢𞤴𞤱𞤢𞥄𞤲𞤳𞤮", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "𞤅𞤭𞤤𞤭𞤲 𞤚𞤢𞤲𞥁𞤢𞤲𞤭𞤲𞤳𞤮", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "𞤖𞤵𞤪𞤢𞤾𞤫𞤲𞤭𞤴𞤢 𞤒𞤵𞤳𞤫𞤪𞤫𞥅𞤲𞤭𞤲𞤳𞤮", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "𞤅𞤭𞤤𞤭𞤲 𞤓𞤺𞤢𞤲𞤣𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "𞤁𞤢𞤤𞤢 𞤁𞤫𞤲𞤼𞤢𞤤 𞤂𞤢𞤪𞤫 𞤀𞤥𞤫𞤪𞤭𞤳", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "𞤆𞤫𞤧𞤮 𞤓𞤪𞤵𞤺𞤵𞤪𞤭𞤲𞤳𞤮", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "𞤅𞤮𞤥𞤵 𞤓𞥁𞤦𞤫𞤳𞤭𞤧𞤼𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "𞤄𞤮𞤤𞤭𞤾𞤢𞥄𞤪 𞤜𞤫𞤲𞤭𞥅𞤧𞤫𞤤𞤢𞤲𞤳𞤮 (𞥒𞥐𞥐𞥘 - 𞥒𞥐𞥑𞥘)", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "𞤄𞤮𞤤𞤭𞤾𞤢𞥄𞤪 𞤜𞤫𞤲𞤭𞥅𞤧𞤫𞤤𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "𞤁𞤮𞤲𞤺𞤵 𞤜𞤭𞤴𞤫𞤼𞤭𞤲𞤢𞤴𞤢𞤲𞤳𞤮", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "𞤜𞤢𞤼𞤵 𞤜𞤢𞤲𞤵𞤴𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "𞤚𞤢𞤤𞤢 𞤅𞤢𞤥𞤮𞤱𞤢𞤴𞤢𞤲𞤳𞤮", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "𞤊𞤢𞤪𞤢𞤲 𞤚𞤵𞤦𞤮𞥅𞤪𞤭 𞤀𞤬𞤪𞤭𞤳𞤭𞤲𞤳𞤮", Symbol: "𞤊𞤅𞤊𞤀"},
				currency.XCD: cldr.Currency{DisplayName: "𞤁𞤢𞤤𞤢 𞤊𞤵𞤯𞤲𞤢𞥄𞤲𞥋𞤺𞤫 𞤑𞤢𞤪𞤭𞤦𞤭𞤴𞤢", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "𞤊𞤢𞤪𞤢𞤲 𞤅𞤊𞤀 𞤖𞤭𞥅𞤪𞤲𞤢𞥄𞤲𞤺𞤫 𞤀𞤬𞤪𞤭𞤳𞤢", Symbol: "𞤅𞤊𞤀"},
				currency.XPF: cldr.Currency{DisplayName: "𞤊𞤢𞤪𞤢𞤲 𞤅𞤊𞤆", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "𞤐𞤄𞤵𞥅𞤯𞤭 𞤢𞤧-𞤢𞤲𞤣𞤢𞥄𞤯𞤭", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "𞤈𞤭𞤴𞤢𞥄𞤤 𞤒𞤫𞤥𞤫𞤲𞤭𞤲𞤳𞤮", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "𞤈𞤢𞤲𞤣𞤭 𞤂𞤫𞤴𞤤𞤫𞤴𞤪𞤭 𞤀𞤬𞤪𞤭𞤳𞤢𞤲𞤳𞤮", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "𞤑𞤢𞤱𞤢𞤧𞤢 𞤟𞤢𞤥𞤦𞤭𞤲𞤳𞤮", Symbol: "ZK"},
			},
		},
	}
}
