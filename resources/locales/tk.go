// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_tk() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "d MMMM y EEEE", Long: "d MMMM y", Medium: "d MMM y", Short: "dd.MM.y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Ýan", Feb: "Few", Mar: "Mar", Apr: "Apr", May: "Maý", Jun: "Iýun", Jul: "Iýul", Aug: "Awg", Sep: "Sen", Oct: "Okt", Nov: "Noý", Dec: "Dek"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "Ý", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "I", Jul: "I", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Ýanwar", Feb: "Fewral", Mar: "Mart", Apr: "Aprel", May: "Maý", Jun: "Iýun", Jul: "Iýul", Aug: "Awgust", Sep: "Sentýabr", Oct: "Oktýabr", Nov: "Noýabr", Dec: "Dekabr"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Ýek", Mon: "Duş", Tue: "Siş", Wed: "Çar", Thu: "Pen", Fri: "Ann", Sat: "Şen"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Ý", Mon: "D", Tue: "S", Wed: "Ç", Thu: "P", Fri: "A", Sat: "Ş"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Ýb", Mon: "Db", Tue: "Sb", Wed: "Çb", Thu: "Pb", Fri: "An", Sat: "Şb"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Ýekşenbe", Mon: "Duşenbe", Tue: "Sişenbe", Wed: "Çarşenbe", Thu: "Penşenbe", Fri: "Anna", Sat: "Şenbe"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "g.öň", PM: "g.soň"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "öň", PM: "soň"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "günortadan öň", PM: "günortadan soň"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_tk]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "BAE dirhemi", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "Owgan afganisi", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "Alban leki", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Ermeni dramy", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Niderland antil guldeni", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Angola kwanzasy", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "Argentin pesosy", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "Awstraliýa dollary", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Aruba florini", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "Azerbaýjan manady", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "Konwertirlenýän Bosniýa we Gersegowina markasy", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "Barbados dollary", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Bangladeş takasy", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "Bolgar lewi", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "Bahreýn dinary", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Burundi franky", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Bermuda dollary", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Bruneý dollary", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Boliwiýa boliwianosy", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "Brazil realy", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Bagama dollary", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Butan ngultrumy", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "Botswana pulasy", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "Belarus rubly", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Belorus rubly (2000–2016)", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "Beliz dollary", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Kanada dollary", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Kongo franky", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "Şweýsar franky", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "Çili pesosy", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Hytaý ýuany (ofşor)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "Hytaý ýuany", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Kolumbiýa pesosy", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "Kosta-Rika kolony", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "Konwertirlenýän kuba pesosy", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Kuba pesosy", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Kabo-Werde eskudosy", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "Çeh kronasy", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "Jibuti franky", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Daniýa kronasy", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Dominikan pesosy", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Alžir dinary", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "Müsür funty", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Eritreýa nakfasy", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Efiopiýa byry", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Ýewro", Symbol: "EUR"},
				currency.FJD: cldr.Currency{DisplayName: "Fiji dollary", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Folklend adalarynyň funty", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "Britan funt sterlingi", Symbol: "GBP"},
				currency.GEL: cldr.Currency{DisplayName: "Gruzin larisi", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "Gano sedisi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Gibraltar funty", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Gambiýa dalasisi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Gwineý franky", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "Gwatemala ketsaly", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "Gaýana dollary", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Gonkong dollary", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Gonduras lempirasy", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "Horwat kunasy", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Gaiti gurdy", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Wenger forinti", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Indoneziýa rupiýasy", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "Täze Ysraýyl şekeli", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Hindi rupiýasy", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Yrak dinary", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Eýran rialy", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "Islandiýa kronasy", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "Ýamaýka dollary", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Iordan dinary", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Ýapon ýeni", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Keniýa şillingi", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Gyrgyz somy", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Kamboja riýeli", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Komor adalarynyň franky", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Demirgazyk Koreý wony", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "Günorta Koreý wony", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Kuweýt dinary", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Kaýman adalarynyň dollary", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Gazak teňňesi", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Laos kipi", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Liwan funty", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Şri-Lanka rupiýasy", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Liberiýa dollary", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Liwiýa dinary", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Marokko dirhamy", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "Moldaw leýi", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Malagasiý ariarisi", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "Makedon dinary", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "Mýanma kýaty", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Mongol tugrigi", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Makao patakasy", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Mawritan ugiýasy (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Mawritan ugiýasy", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "Mawrikiý rupiýasy", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Maldiw rufiýasy", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Malawi kwaçasy", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Meksikan pesosy", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Malaýziýa ringgiti", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "Mozambik metikaly", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Namibiýa dollary", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Nigeriýa naýrasy", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "Nikaragua kordobasy", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "Norwegiýa kronasy", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Nepal rupiýasy", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Täze Zelandiýa dollary", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Oman rialy", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Panama balboasy", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "Peru soly", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "Papua - Täze Gwineýa kinasy", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Filippin pesosy", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Päkistan rupiýasy", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Polýak zlotysy", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "Paragwaý guaranisi", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Katar rialy", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "Rumyn leýi", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Serb dinary", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Rus rubly", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Ruanda franky", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Saud rialy", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Solomon adalarynyň dollary", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Seýşel rupiýasy", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "Sudan funty", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "Şwed kronasy", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Singapur dollary", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Keramatly Ýelena adasynyň funty", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "Sýerra-Leone leony", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Somali şillingi", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Surinam dollary", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "Günorta Sudan funty", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "San-Tome we Prinsipi dobrasy (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "San-Tome we Prinsipi dobrasy", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "Siriýa funty", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Swazi lilangeni", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Taýland baty", Symbol: "THB"},
				currency.TJS: cldr.Currency{DisplayName: "Täjik somonisi", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "Türkmen manady", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Tunis dinary", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Tonga paangasy", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "Türk lirasy", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Trininad we Tobago dollary", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Täze Taýwan dollary", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Tanzaniýa şillingi", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Ukrain griwnasy", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "Uganda şillingi", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "ABŞ dollary", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "Urugwaý pesosy", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Özbek somy", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "Wenesuela boliwary (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Wenesuela boliwary", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Wýetnam dongy", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Wanuatu watusy", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Samoa talasy", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "KFA BEAC franky", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Gündogar karib dollary", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "KFA BCEAO franky", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "Fransuz ýuwaş umman franky", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Näbelli pul birligi", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Ýemen rialy", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "Günorta Afrika rendi", Symbol: "ZAR"},
				currency.ZMW: cldr.Currency{DisplayName: "Zambiýa kwaçasy", Symbol: "ZMW"},
			},
		},
	}
}
