// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_pa() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/yy"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ਜਨ", Feb: "ਫ਼ਰ", Mar: "ਮਾਰਚ", Apr: "ਅਪ੍ਰੈ", May: "ਮਈ", Jun: "ਜੂਨ", Jul: "ਜੁਲਾ", Aug: "ਅਗ", Sep: "ਸਤੰ", Oct: "ਅਕਤੂ", Nov: "ਨਵੰ", Dec: "ਦਸੰ"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "ਜ", Feb: "ਫ਼", Mar: "ਮਾ", Apr: "ਅ", May: "ਮ", Jun: "ਜੂ", Jul: "ਜੁ", Aug: "ਅ", Sep: "ਸ", Oct: "ਅ", Nov: "ਨ", Dec: "ਦ"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "ਜਨਵਰੀ", Feb: "ਫ਼ਰਵਰੀ", Mar: "ਮਾਰਚ", Apr: "ਅਪ੍ਰੈਲ", May: "ਮਈ", Jun: "ਜੂਨ", Jul: "ਜੁਲਾਈ", Aug: "ਅਗਸਤ", Sep: "ਸਤੰਬਰ", Oct: "ਅਕਤੂਬਰ", Nov: "ਨਵੰਬਰ", Dec: "ਦਸੰਬਰ"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ਐਤ", Mon: "ਸੋਮ", Tue: "ਮੰਗਲ", Wed: "ਬੁੱਧ", Thu: "ਵੀਰ", Fri: "ਸ਼ੁੱਕਰ", Sat: "ਸ਼ਨਿੱਚਰ"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "ਐ", Mon: "ਸੋ", Tue: "ਮੰ", Wed: "ਬੁੱ", Thu: "ਵੀ", Fri: "ਸ਼ੁੱ", Sat: "ਸ਼"}, Short: cldr.CalendarDayFormatNameValue{Sun: "ਐਤ", Mon: "ਸੋਮ", Tue: "ਮੰਗ", Wed: "ਬੁੱਧ", Thu: "ਵੀਰ", Fri: "ਸ਼ੁੱਕ", Sat: "ਸ਼ਨਿੱ"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "ਐਤਵਾਰ", Mon: "ਸੋਮਵਾਰ", Tue: "ਮੰਗਲਵਾਰ", Wed: "ਬੁੱਧਵਾਰ", Thu: "ਵੀਰਵਾਰ", Fri: "ਸ਼ੁੱਕਰਵਾਰ", Sat: "ਸ਼ਨਿੱਚਰਵਾਰ"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "ਪੂ.ਦੁ.", PM: "ਬਾ.ਦੁ."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "ਪੂ.ਦੁ.", PM: "ਬਾ.ਦੁ."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "ਪੂ.ਦੁ.", PM: "ਬਾ.ਦੁ."}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_pa]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "٫", Group: ",", Negative: "-", Percent: "٪", PerMille: "؉"},
			Formats: cldr.NumberFormats{Decimal: "#,##,##0.###", Currency: "¤#,##,##0.00", CurrencyAccounting: "", Percent: "#,##,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "ਸੰਯੁਕਤ ਅਰਬ ਅਮੀਰਾਤ ਦਿਰਹਾਮ", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "ਅਫ਼ਗਾਨ ਅਫ਼ਗਾਨੀ", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "ਅਲਬਾਨੀਆਈ ਲੇਕ", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "ਅਰਮੀਨੀਆਈ ਦਰਾਮ", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "ਨੀਦਰਲੈਂਡਸ ਐਂਟੀਲੀਅਨ ਗਿਲਡਰ", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "ਅੰਗੋਲਾ ਕਵਾਂਜਾ", Symbol: "AOA"},
				currency.ARA: cldr.Currency{DisplayName: "ਅਰਜਨਟੀਨੀ ਅਸਟਰਾਲ", Symbol: "ARA"},
				currency.ARL: cldr.Currency{DisplayName: "ਅਰਜਨਟੀਨੀ ਪੇਸੋ ਲੇ (1970–1983)", Symbol: "ARL"},
				currency.ARM: cldr.Currency{DisplayName: "ਅਰਜਨਟੀਨੀ ਪੇਸੋ (1881–1970)", Symbol: "ARM"},
				currency.ARP: cldr.Currency{DisplayName: "ਅਰਜਨਟੀਨੀ ਪੇਸੋ (1983–1985)", Symbol: "ARP"},
				currency.ARS: cldr.Currency{DisplayName: "ਅਰਜਨਟੀਨੀ ਪੇਸੋ", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "ਆਸਟ੍ਰੇਲੀਆਈ ਡਾਲਰ", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "ਅਰੂਬਨ ਫਲੋਰਿਨ", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "ਅਜ਼ਰਬਾਈਜਾਨ ਮਾਨਤ", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "ਬੋਸਨੀਆ-ਹਰਜ਼ੇਗੋਵੀਨਾ ਬਦਲਣਯੋਗ ਮਾਰਕ", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "ਬਾਰਬਾਡੀਅਨ ਡਾਲਰ", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "ਬੰਗਲਾਦੇਸ਼ੀ ਟਕਾ", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "ਬੁਲਗਾਰੀਆਈ ਲੇਵ", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "ਬਹਿਰੀਨੀ ਦਿਨਾਰ", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "ਬੁਰੁੰਡੀਆਈ ਫ੍ਰੈਂਕ", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "ਬਰਮੂਡਾ ਡਾਲਰ", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "ਬਰੂਨੇਈ ਡਾਲਰ", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "ਬੋਲੀਵੀਅਨ ਬੋਲੀਵੀਅਨੋ", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "ਬੋਲੀਵੀਆਈ ਬੋਲੀਵੀਅਨੋ (1863–1963)", Symbol: "BOL"},
				currency.BOP: cldr.Currency{DisplayName: "ਬੋਲੀਵੀਆਈ ਪੇਸੋ", Symbol: "BOP"},
				currency.BOV: cldr.Currency{DisplayName: "ਬੋਲੀਵੀਆਈ ਮਵਡੋਲ", Symbol: "BOV"},
				currency.BRB: cldr.Currency{DisplayName: "ਬ੍ਰਾਜ਼ੀਲੀਆਈ ਨਿਊ ਕਰੁਜ਼ਿਰੋਸ (1967–1986)", Symbol: "BRB"},
				currency.BRC: cldr.Currency{DisplayName: "ਬ੍ਰਾਜ਼ੀਲੀਆਈ ਕਰੁਜ਼ਾਡੂ (1986–1989)", Symbol: "BRC"},
				currency.BRE: cldr.Currency{DisplayName: "ਬ੍ਰਾਜ਼ੀਲੀਆਈ ਕਰੁਜ਼ਿਰੋਸ (1990–1993)", Symbol: "BRE"},
				currency.BRL: cldr.Currency{DisplayName: "ਬ੍ਰਾਜ਼ੀਲੀਆਈ ਰੀਅਲ", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "ਬ੍ਰਾਜ਼ੀਲੀਆਈ ਨਿਊ ਕਰੁਜ਼ਾਡੂ (1989–1990)", Symbol: "BRN"},
				currency.BRR: cldr.Currency{DisplayName: "ਬ੍ਰਾਜ਼ੀਲੀਆਈ ਕਰੁਜ਼ਿਰੋਸ (1993–1994)", Symbol: "BRR"},
				currency.BRZ: cldr.Currency{DisplayName: "ਬ੍ਰਾਜ਼ੀਲੀਆਈ ਕਰੁਜ਼ਿਰੋਸ (1942–1967)", Symbol: "BRZ"},
				currency.BSD: cldr.Currency{DisplayName: "ਬਾਹਾਮੀਅਨ ਡਾਲਰ", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "ਭੂਟਾਨੀ ਐਂਗਲਟ੍ਰਮ", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "ਬੋਟਸਵਾਨਾ ਪੁਲਾ", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "ਬੇਲਾਰੂਸੀ ਰੂਬਲ", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "ਬੇਲਾਰੂਸੀ ਰੂਬਲ (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "ਬੇਲੀਜ਼ ਡਾਲਰ", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "ਕੇਨੇਡਿਆਈ ਡਾਲਰ", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "ਕਾਂਗੋਲੀਜ਼ ਫ੍ਰੈਂਕ", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "ਸਵਿਸ ਫ੍ਰੈਂਕ", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "ਚਿਲੀ ਪੇਸੋ", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "ਚੀਨੀ ਯੁਆਨ (ਔਫ਼ਸ਼ੋਰ)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "ਚੀਨੀ ਯੁਆਨ", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "ਕੋਲੰਬਿਆਈ ਪੇਸੋ", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "ਕੋਸਟਾ ਰੀਕਨ ਕੋਲਨ", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "ਕਿਊਬਨ ਬਦਲਣਯੋਗ ਪੇਸੋ", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "ਕਿਊਬਨ ਪੇਸੋ", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "ਕੇਪ ਵਰਡੀਅਨ ਸਕੂਡੋ", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "ਚੈਕ ਗਣਰਾਜ ਕੋਰੁਨਾ", Symbol: "CZK"},
				currency.DEM: cldr.Currency{DisplayName: "ਜਰਮਨ ਮਾਰਕ", Symbol: "DEM"},
				currency.DJF: cldr.Currency{DisplayName: "ਜ਼ੀਬੂਤੀਅਨ ਫ੍ਰੈਂਕ", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "ਡੈਨਿਸ਼ ਕਰੌਨ", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "ਡੌਮਿਨਿਕਨ ਪੇਸੋ", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "ਅਲਜੀਰਿਆਈ ਦਿਨਾਰ", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "ਮਿਸਰੀ ਪੌਂਡ", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "ਇਰੀਟ੍ਰਿਆਈ ਨਾਫ਼ਾ", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "ਇਥੋਪੀਆਈ ਬਿਰ", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "ਯੂਰੋ", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "ਫ਼ਿਜ਼ੀ ਡਾਲਰ", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "ਫ਼ਾਕਲੈਂਡ ਆਈਲੈਂਡਸ ਪੌਂਡ", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "ਬ੍ਰਿਟਿਸ਼ ਪੌਂਡ", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "ਜਾਰਜੀਆਈ ਲਾਰੀ", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "ਘਾਨਾਈ ਸੇਡੀ", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "ਜਿਬਰਾਲਟਰ ਪੌਂਡ", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "ਗੈਂਬੀਆਈ ਦਲਾਸੀ", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "ਗਿਨੀ ਫ੍ਰੈਂਕ", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "ਗੁਆਟੇਮਾਲਾ ਕੁਏਟਜ਼ਲ", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "ਗੁਆਨਾਆਈ ਡਾਲਰ", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "ਹਾਂਗ ਕਾਂਗ ਡਾਲਰ", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "ਹਾਨਡੂਰਨ ਲੇਮਪਿਰਾ", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "ਕਰੋਏਸ਼ੀਆਈ ਕੁਨਾ", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "ਹੈਤੀ ਗੌਰਡੇ", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "ਹੰਗਰੀ ਫੋਰਿੰਟ", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "ਇੰਡੋਨੇਸ਼ੀਆਈ ਰੁਪਿਆਹ", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "ਆਇਰਿਸ਼ ਪੌਂਡ", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "ਇਜ਼ਰਾਈਲੀ ਪੌਂਡ", Symbol: "ILP"},
				currency.ILS: cldr.Currency{DisplayName: "ਇਜ਼ਰਾਈਲੀ ਨਵੀਂ ਸ਼ੇਕੇਲ", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "ਭਾਰਤੀ ਰੁਪਇਆ", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "ਇਰਾਕੀ ਦਿਨਾਰ", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "ਈਰਾਨੀ ਰਿਆਲ", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "ਆਈਸਲੈਂਡੀ ਕਰੋਨਾ", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "ਜਮਾਇਕਨ ਡਾਲਰ", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "ਜਾਰਡਨ ਦਿਨਾਰ", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "ਜਪਾਨੀ ਯੇਨ", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "ਕੀਨੀਆਈ ਸ਼ਿਲਿੰਗ", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "ਕਿਰਗਿਸਤਾਨੀ ਸੋਮ", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "ਕੰਬੋਡੀਆਈ ਰੀਅਲ", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "ਕੋਮੋਰੀਅਨ ਫ੍ਰੈਂਕ", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "ਉੱਤਰੀ ਕੋਰੀਆਈ ਵੋਨ", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "ਦੱਖਣੀ ਕੋਰੀਆਈ ਵੋਨ", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "ਕੁਵੈਤੀ ਦਿਨਾਰ", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "ਕੇਮੈਨ ਆਈਲੈਂਡਸ ਡਾਲਰ", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "ਕਜ਼ਾਖਸਤਾਨੀ ਤੇਂਗੇ", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "ਲਾਓਟਿਆਈ ਕਿਪ", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "ਲੈਬਨਾਨੀ ਪੌਂਡ", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "ਸ੍ਰੀਲੰਕਾਈ ਰੁਪਇਆ", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "ਲਾਈਬੀਰੀਆਈ ਡਾਲਰ", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "ਲਿਥੁਆਨੀਆਈ ਲਿਤਾਸ", Symbol: "LTL"},
				currency.LVL: cldr.Currency{DisplayName: "ਲਾਟਵਿਆਈ ਲਾਟਸ", Symbol: "LVL"},
				currency.LYD: cldr.Currency{DisplayName: "ਲੀਬੀਆਈ ਦਿਨਾਰ", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "ਮੋਰੱਕਨ ਦਿਰਹਾਮ", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "ਮੋਲਡੋਵਨ ਲੇਉ", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "ਮਾਲਾਗਾਸੀ ਅਰਾਇਰੀ", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "ਮੈਕਡੋਨੀਆਈ ਡੇਨਾਰ", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "ਮਿਆਂਮਾਰ ਕਿਆਤ", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "ਮੰਗੋਲੀਆਈ ਤੁਗਰਿਕ", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "ਮੇਕਾਨੀ ਪਟਾਕਾ", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "ਮੋਰਿਟਾਨੀਆਈ ਊਗੀਆ (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "ਮੋਰਿਟਾਨੀਆਈ ਊਗੀਆ", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "ਮੌਰਿਸ਼ੀਆਈ ਰੁਪਇਆ", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "ਮਾਲਦੀਵੀ ਰੁਫੀਆ", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "ਮਾਲਾਵੀਆਈ ਕਵਾਚਾ", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "ਮੈਕਸੀਕਨ ਪੇਸੋ", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "ਮਲੇਸ਼ੀਆਈ ਰਿੰਗਿਟ", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "ਮੋਜ਼ਾਮਬੀਕਨ ਮੈਟੀਕਲ", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "ਨਾਮੀਬੀਆਈ ਡਾਲਰ", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "ਨਾਇਜੀਰੀਆਈ ਨਾਇਰਾ", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "ਨਿਕਾਰਾਗੁਆਈ ਕੋਰਡੋਬਾ", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "ਨਾਰਵੇਜੀਆਈ ਕਰੌਨ", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "ਨੇਪਾਲੀ ਰੁਪਇਆ", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "ਨਿਊਜ਼ੀਲੈਂਡ ਡਾਲਰ", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "ਓਮਾਨੀ ਰਿਆਲ", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "ਪਨਾਮੇਨੀਅਨ ਬਾਲਬੋਆ", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "ਪੇਰੂਵੀਅਨ ਸੋਲ", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "ਪਾਪੂਆ ਨਿਊ ਗਿਨੀਆਈ ਕੀਨਾ", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "ਫਿਲਿਪੀਨੀ ਪੇਸੋ", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "ਪਾਕਿਸਤਾਨੀ ਰੁਪਇਆ", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "ਪੋਲੈਂਡੀ ਜ਼ਲੌਟੀ", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "ਪੈਰਾਗੁਵਾਇਨ ਗੁਆਰਾਨੀ", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "ਕਤਰੀ ਰਿਆਲ", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "ਰੋਮਾਨੀਆਈ ਲੇਉ", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "ਸਰਬੀਆਈ ਦਿਨਾਰ", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "ਰੂਸੀ ਰੂਬਲ", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "ਰਵਾਂਡਨ ਫ੍ਰੈਂਕ", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "ਸਾਊਦੀ ਰਿਆਲ", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "ਸੋਲੋਮਨ ਆਈਲੈਂਡਸ ਡਾਲਰ", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "ਸੇਸ਼ਲਸ ਰੁਪਇਆ", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "ਸੂਡਾਨੀ ਪੌਂਡ", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "ਸਵੀਡਿਸ਼ ਕਰੋਨਾ", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "ਸਿੰਗਾਪੁਰ ਡਾਲਰ", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "ਸੇਂਟ ਹੇਲੇਨਾ ਪੌਂਡ", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "ਸਿਏਰਾ ਲਿਓਨੀਅਨ ਲਿਓਨ", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "ਸੋਮਾਲੀ ਸ਼ਿਲਿੰਗ", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "ਸੂਰੀਨਾਮੀ ਡਾਲਰ", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "ਦੱਖਣੀ ਸੂਡਾਨੀ ਪੌਂਡ", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "ਸਾਉ ਟੋਮੀ ਐਂਡ ਪ੍ਰਿੰਸਪੀ ਡੋਬਰਾ (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "ਸਾਉ ਟੋਮੀ ਐਂਡ ਪ੍ਰਿੰਸਪੀ ਡੋਬਰਾ", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "ਸੋਵੀਅਤ ਰੂਬਲ", Symbol: "SUR"},
				currency.SYP: cldr.Currency{DisplayName: "ਸੀਰੀਆਈ ਪੌਂਡ", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "ਸਵਾਜ਼ੀ ਲਾਇਲੈਂਗਨੀ", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "ਥਾਈ ਬਾਹਤ", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "ਤਾਜਿਕਿਸਤਾਨੀ ਸੋਮੋਨੀ", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "ਤੁਰਕਮੇਨਿਸਤਾਨੀ ਮਾਨਤ", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "ਟਿਉਨੀਸ਼ੀਆਈ ਦਿਨਾਰ", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "ਟੌਂਗਨ ਪੈਂਗਾ", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "ਤੁਰਕੀ ਲੀਰਾ", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "ਟ੍ਰਿਨੀਡਾਡ ਅਤੇ ਟੋਬਾਗੋ ਡਾਲਰ", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "ਨਵਾਂ ਤਾਇਵਾਨ ਡਾਲਰ", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "ਤਨਜ਼ਾਨੀਆਈ ਸ਼ਿਲਿੰਗ", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "ਯੂਕਰੇਨੀਆਈ ਰਿਵਨਿਆ", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "ਯੂਗਾਂਡੀਆਈ ਸ਼ਿਲਿੰਗ", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "ਯੂ.ਐਸ. ਡਾਲਰ", Symbol: "US$"},
				currency.UYI: cldr.Currency{DisplayName: "", Symbol: "UYI"},
				currency.UYP: cldr.Currency{DisplayName: "ਉਰੂਗੁਵਾਇਨ ਪੇਸੋ (1975–1993)", Symbol: "UYP"},
				currency.UYU: cldr.Currency{DisplayName: "ਉਰੂਗੁਵਾਇਨ ਪੇਸੋ", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "ਉਜ਼ਬੇਕਿਸਤਾਨ ਸੋਮ", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "ਵੇਨੇਜ਼ੂਏਲਨ ਬੋਲੀਵਰ (1871–2008)", Symbol: "VEB"},
				currency.VEF: cldr.Currency{DisplayName: "ਵੇਨੇਜ਼ੂਏਲਨ ਬੋਲੀਵਰ (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "ਵੇਨੇਜ਼ੂਏਲਨ ਬੋਲੀਵਰ", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "ਵੀਅਤਨਾਮੀ ਡੋਂਗ", Symbol: "₫"},
				currency.VNN: cldr.Currency{DisplayName: "ਵੀਅਤਨਾਮੀ ਡੋਂਗ (1978–1985)", Symbol: "VNN"},
				currency.VUV: cldr.Currency{DisplayName: "ਵਾਨੂਆਟੂ ਵਾਟੂ", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "ਸਾਮੋਆਈ ਤਾਲਾ", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "ਕੇਂਦਰੀ ਅਫ਼ਰੀਕੀ [CFA] ਫ੍ਰੈਂਕ", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "ਚਾਂਦੀ", Symbol: "XAG"},
				currency.XAU: cldr.Currency{DisplayName: "ਸੋਨਾ", Symbol: "XAU"},
				currency.XBA: cldr.Currency{DisplayName: "", Symbol: "XBA"},
				currency.XBB: cldr.Currency{DisplayName: "ਯੂਰਪੀ ਵਿੱਤੀ ਇਕਾਈ", Symbol: "XBB"},
				currency.XBC: cldr.Currency{DisplayName: "", Symbol: "XBC"},
				currency.XBD: cldr.Currency{DisplayName: "", Symbol: "XBD"},
				currency.XCD: cldr.Currency{DisplayName: "ਪੂਰਬੀ ਕੈਰੇਬੀਅਨ ਡਾਲਰ", Symbol: "EC$"},
				currency.XEU: cldr.Currency{DisplayName: "ਯੂਰਪੀ ਮੁਦਰਾ ਇਕਾਈ", Symbol: "XEU"},
				currency.XOF: cldr.Currency{DisplayName: "ਪੱਛਮੀ ਅਫ਼ਰੀਕੀ (CFA) ਫ੍ਰੈਂਕ", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "ਫ੍ਰੈਂਕ (CFP)", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "ਅਗਿਆਤ ਮੁਦਰਾ", Symbol: "XXX"},
				currency.YER: cldr.Currency{DisplayName: "ਯਮਨੀ ਰਿਆਲ", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "ਦੱਖਣੀ ਅਫਰੀਕੀ ਰੈਂਡ", Symbol: "ZAR"},
				currency.ZMW: cldr.Currency{DisplayName: "ਜ਼ਾਮਬੀਆਈ ਕਵਾਚਾ", Symbol: "ZMW"},
			},
		},
	}
}
