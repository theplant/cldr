// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_sw_UG() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "sw_UG",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT {0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mac", Apr: "Apr", May: "Mei", Jun: "Jun", Jul: "Jul", Aug: "Ago", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Des"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Januari", Feb: "Februari", Mar: "Machi", Apr: "Aprili", May: "Mei", Jun: "Juni", Jul: "Julai", Aug: "Agosti", Sep: "Septemba", Oct: "Oktoba", Nov: "Novemba", Dec: "Desemba"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Jumapili", Mon: "Jumatatu", Tue: "Jumanne", Wed: "Jumatano", Thu: "Alhamisi", Fri: "Ijumaa", Sat: "Jumamosi"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "W", Thu: "T", Fri: "F", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Jumapili", Mon: "Jumatatu", Tue: "Jumanne", Wed: "Jumatano", Thu: "Alhamisi", Fri: "Ijumaa", Sat: "Jumamosi"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Jumapili", Mon: "Jumatatu", Tue: "Jumanne", Wed: "Jumatano", Thu: "Alhamisi", Fri: "Ijumaa", Sat: "Jumamosi"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_sw]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a0#,##0.00", CurrencyAccounting: "¤\u00a0#,##0.00", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Dirham ya Falme za Kiarabu", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "Afghani ya Afghanistan", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "Lek ya Albania", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Dram ya Armenia", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Guilder ya Antili za Kiholanzi", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Kwanza ya Angola", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "Peso ya Argentina", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "Dola ya Australia", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Florin ya Aruba", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "Manat ya Azerbaijan", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "Convertible Mark ya Bosnia na Hezegovina", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "Dola ya Barbados", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Taka ya Bangladesh", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "Lev ya Bulgaria", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "Dinar ya Bahrain", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Faranga ya Burundi", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Dola ya Bermuda", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Dola ya Brunei", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Boliviano ya Bolivia", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "Real ya Brazil", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Dola ya Bahamas", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Ngultrum ya Bhutan", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "Pula ya Botswana", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "Ruble ya Belarus", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Ruble ya Belarusi (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "Dola ya Belize", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Dola ya Canada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Faranga ya Kongo", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "Faranga ya Uswisi", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "Peso ya Chile", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Yuan ya Uchina (huru)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "Yuan ya Uchina", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Peso ya Colombia", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "Colon ya Costa Rica", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "Peso ya Cuba Inayoweza Kubadilishwa", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Peso ya Cuba", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Eskudo ya Cape Verde", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "koruna ya Jamhuri ya Czech", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "Faranga ya Djibouti", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Krone ya Denmark", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Peso ya Dominica", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Dinar ya Aljeria", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "Pauni ya Misri", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfa ya Eritrea", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Birr ya Uhabeshi", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Yuro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Dola ya Fiji", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Pauni ya Visiwa vya Falkland", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "Pauni ya Uingereza", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Lari ya Georgia", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "Sedi ya Ghana", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Cedi ya Ghana", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Pauni ya Gibraltar", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi ya Gambia", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Faranga ya Guinea", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "Faranga ya Gine", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "Quetzal ya Guatemala", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "Dola ya Guyana", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Dola ya Hong Kong", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Lempira ya Hondurasi", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "Kuna ya Croatia", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Gourde ya Haiti", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Forint ya Hungaria", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Rupiah ya Indonesia", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "Shekeli Mpya ya Israel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupia ya India", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Dinar ya Iraq", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Rial ya Iran", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "Krona ya Aisilandi", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "Dola ya Jamaica", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Dinar ya Jordan", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Yen ya Ujapani", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Shilingi ya Kenya", Symbol: "Ksh"},
				currency.KGS: cldr.Currency{DisplayName: "Som ya Kyrgystan", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Riel ya Cambodia", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Faranga ya Comoros", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Won ya Korea Kaskazini", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "Won ya Korea Kusini", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Dinar ya Kuwait", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Dola ya Visiwa vya Cayman", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Tenge ya Kazakhstan", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Kip ya Laosi", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Pauni ya Lebanon", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Rupia ya Sri Lanka", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Dola ya Liberia", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Loti ya Lesoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Litas ya Lithuania", Symbol: "LTL"},
				currency.LVL: cldr.Currency{DisplayName: "Lats ya Lativia", Symbol: "LVL"},
				currency.LYD: cldr.Currency{DisplayName: "Dinari ya Libya", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Dirham ya Morocco", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "Leu ya Moldova", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Ariari ya Madagascar", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "Denar ya Macedonia", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "Kyat ya Myanmar", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Tugrik ya Mongolia", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Pataca ya Macau", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Ouguiya ya Mauritania (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Ouguiya ya Mauritania", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "Rupia ya Mauritius", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Rufiyaa ya Maldives", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Kwacha ya Malawi", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Peso ya Mexico", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Ringgit ya Malaysia", Symbol: "MYR"},
				currency.MZM: cldr.Currency{DisplayName: "Metikali ya Msumbiji (1980–2006)", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Metikali ya Msumbiji", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Dola ya Namibia", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Naira ya Nigeria", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "Cordoba ya Nicaragua", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "Krone ya Norway", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Rupia ya Nepal", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Dola ya New Zealand", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Rial ya Omani", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Balboa ya Panama", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "Sol ya Peru", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "Kina ya Papua New Guinea", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Peso ya Ufilipino", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Rupia ya Pakistan", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Zloty ya Poland", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "Guarani ya Paraguay", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Rial ya Qatar", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "Leu ya Romania", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Dinar ya Serbia", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Ruble ya Urusi", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Faranga ya Rwanda", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Riyal ya Saudia", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Dola ya Visiwa vya Solomon", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Rupia ya Ushelisheli", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "Pauni ya Sudan", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "Pauni ya Sudani (1957–1998)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Krona ya Uswidi", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Dola ya Singapore", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Pauni ya St. Helena", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "Leone", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Shilingi ya Somalia", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Dola ya Suriname", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "Pauni ya Sudan Kusini", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "Dobra ya Sao Tome na Principe (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "Dobra ya Sao Tome na Principe", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "Pauni ya Syria", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Baht ya Tailandi", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "Somoni ya Tajikistan", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "Manat ya Turkmenistan", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Dinari ya Tunisia", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Paʻanga ya Tonga", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "Lira ya Uturuki", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Dola ya Trinidad na Tobago", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Dola ya Taiwan", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Shilingi ya Tanzania", Symbol: "TSh"},
				currency.UAH: cldr.Currency{DisplayName: "Hryvnia ya Ukraine", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "Shilingi ya Uganda", Symbol: "USh"},
				currency.USD: cldr.Currency{DisplayName: "Dola ya Marekani", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "Peso ya Uruguay", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Som ya Uzbekistan", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "Bolivar ya Venezuela (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Bolivar ya Venezuela", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Dong ya Vietnam", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Vatu ya Vanuatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Tala ya Samoa", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Faranga ya Afrika ya Kati CFA", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Dola ya Caribbean Mashariki", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Faranga ya Afrika Magharibi CFA", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "Faranga ya CFP", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Sarafu isiyojulikana", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Rial ya Yemen", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "Randi ya Afrika Kusini", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwacha ya Zambia (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kwacha ya Zambia", Symbol: "ZMW"},
				currency.ZWD: cldr.Currency{DisplayName: "Dola ya Zimbabwe", Symbol: ""},
			},
		},
	}
}
