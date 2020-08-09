// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_ms_BN() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "dd MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/MM/yy"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mac", Apr: "Apr", May: "Mei", Jun: "Jun", Jul: "Jul", Aug: "Ogo", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Dis"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "O", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Januari", Feb: "Februari", Mar: "Mac", Apr: "April", May: "Mei", Jun: "Jun", Jul: "Julai", Aug: "Ogos", Sep: "September", Oct: "Oktober", Nov: "November", Dec: "Disember"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Ahd", Mon: "Isn", Tue: "Sel", Wed: "Rab", Thu: "Kha", Fri: "Jum", Sat: "Sab"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "A", Mon: "I", Tue: "S", Wed: "R", Thu: "K", Fri: "J", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Ah", Mon: "Is", Tue: "Se", Wed: "Ra", Thu: "Kh", Fri: "Ju", Sat: "Sa"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Ahad", Mon: "Isnin", Tue: "Selasa", Wed: "Rabu", Thu: "Khamis", Fri: "Jumaat", Sat: "Sabtu"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "PG", PM: "PTG"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "a", PM: "p"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "PG", PM: "PTG"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_ms]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a0#,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Dirham Emiriah Arab Bersatu", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "Afghani Afghanistan", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "Lek Albania", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Dram Armenia", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Guilder Antillen Belanda", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Kwanza Angola", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "Peso Argentina", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "Dolar Australia", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Florin Aruba", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "Manat Azerbaijan", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "Mark Boleh Tukar Bosnia-Herzegovina", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "Dolar Barbados", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Taka Bangladesh", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "Lev Bulgaria", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "Dinar Bahrain", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Franc Burundi", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Dolar Bermuda", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Dolar Brunei", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "Boliviano Bolivia", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "Real Brazil", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Dolar Bahamas", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Ngultrum Bhutan", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "Pula Botswana", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "Rubel Belarus baharu", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Rubel Belarus (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "Dolar Belize", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Dolar Kanada", Symbol: "CAD"},
				currency.CDF: cldr.Currency{DisplayName: "Franc Congo", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "Franc Switzerland", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "Peso Chile", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Yuan China (luar pesisir)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "Yuan Cina", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Peso Colombia", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "Colon Costa Rica", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "Peso Boleh Tukar Cuba", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Peso Cuba", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Escudo Tanjung Verde", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "Koruna Republik Czech", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "Franc Djibouti", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Krone Denmark", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Peso Dominican", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Dinar Algeria", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "Paun Mesir", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfa Eritrea", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Birr Ethiopia", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Dolar Fiji", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Paun Kepulauan Falkland", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "Paun British", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Lari Georgia", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "Cedi Ghana", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Paun Gibraltar", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi Gambia", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Franc Guinea", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "Quetzal Guatemala", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "Dolar Guyana", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Dolar Hong Kong", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Lempira Honduras", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "Kuna Croatia", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Gourde Haiti", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Forint Hungary", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Rupiah Indonesia", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "Syekel Baharu Israel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupee India", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Dinar Iraq", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Rial Iran", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "Krona Iceland", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "Dolar Jamaica", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Dinar Jordan", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Yen Jepun", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Syiling Kenya", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Som Kyrgystani", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Riel Kemboja", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Franc Comoria", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Won Korea Utara", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "Won Korea Selatan", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Dinar Kuwait", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Dolar Kepulauan Cayman", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Tenge Kazakhstan", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Kip Laos", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Paun Lubnan", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Rupee Sri Lanka", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Dolar Liberia", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Loti Lesotho", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Litas Lithuania", Symbol: "LTL"},
				currency.LVL: cldr.Currency{DisplayName: "Lats Latvia", Symbol: "LVL"},
				currency.LYD: cldr.Currency{DisplayName: "Dinar Libya", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Dirham Maghribi", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "Leu Moldova", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Ariary Malagasy", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "Franc Malagasy", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "Denar Macedonia", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "Kyat Myanma", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Tugrik Mongolia", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Pataca Macau", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Ouguiya Mauritania (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Ouguiya Mauritania", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "Rupee Mauritius", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Rufiyaa Maldives", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Kwacha Malawi", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Peso Mexico", Symbol: "MXN"},
				currency.MYR: cldr.Currency{DisplayName: "Ringgit Malaysia", Symbol: "RM"},
				currency.MZE: cldr.Currency{DisplayName: "Escudo Mozambique", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "Metical Mozambique (1980–2006)", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Metikal Mozambique", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Dolar Namibia", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Naira Nigeria", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "Cordoba Nicaragua", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "Krone Norway", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Rupee Nepal", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Dolar New Zealand", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Rial Oman", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Balboa Panama", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "Sol Peru", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "Kina Papua New Guinea", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Peso Filipina", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Rupee Pakistan", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Zloty Poland", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "Guarani Paraguay", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Rial Qatar", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "Dolar Rhodesia", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Leu Romania", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Dinar Serbia", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Rubel Rusia", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Franc Rwanda", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Riyal Saudi", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Dolar Kepulauan Solomon", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Rupee Seychelles", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "Paun Sudan", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "Krona Sweden", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Dolar Singapura", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Paun Saint Helena", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "Leone Sierra Leone", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Syiling Somali", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Dolar Surinam", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "Paun Sudan selatan", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "Dobra Sao Tome dan Principe (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "Dobra Sao Tome dan Principe", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "Paun Syria", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni Swazi", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Baht Thai", Symbol: "THB"},
				currency.TJS: cldr.Currency{DisplayName: "Somoni Tajikistan", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "Manat Turkmenistan", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Dinar Tunisia", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Pa’anga Tonga", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "Lira Turki", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Dolar Trinidad dan Tobago", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Dolar Taiwan Baru", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Syiling Tanzania", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Hryvnia Ukraine", Symbol: "UAH"},
				currency.UGS: cldr.Currency{DisplayName: "Shilling Uganda (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "Syiling Uganda", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "Dolar AS", Symbol: "USD"},
				currency.UYU: cldr.Currency{DisplayName: "Peso Uruguay", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Som Uzbekistan", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "Bolivar Venezuela (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Bolivar Venezuela", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Dong Vietnam", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Vatu Vanuatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Tala Samoa", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Franc CFA BEAC", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Dolar Caribbean Timur", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Franc CFA BCEAO", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "Franc CFP", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Mata Wang Tidak Diketahui", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Rial Yaman", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "Rand Afrika Selatan", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwacha Zambia (1968–2012)", Symbol: "ZMK"},
				currency.ZMW: cldr.Currency{DisplayName: "Kwacha Zambia", Symbol: "ZMW"},
				currency.ZWD: cldr.Currency{DisplayName: "Dolar Zimbabwe (1980–2008)", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "Dolar Zimbabwe (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "Dolar Zimbabwe (2008)", Symbol: ""},
			},
		},
	}
}
