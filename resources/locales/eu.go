// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_eu() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y('e')'ko' MMMM'ren' d('a'), EEEE", Long: "y('e')'ko' MMMM'ren' d('a')", Medium: "y('e')'ko' MMM d('a')", Short: "yy/M/d"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss (zzzz)", Long: "HH:mm:ss (z)", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "urt.", Feb: "ots.", Mar: "mar.", Apr: "api.", May: "mai.", Jun: "eka.", Jul: "uzt.", Aug: "abu.", Sep: "ira.", Oct: "urr.", Nov: "aza.", Dec: "abe."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "U", Feb: "O", Mar: "M", Apr: "A", May: "M", Jun: "E", Jul: "U", Aug: "A", Sep: "I", Oct: "U", Nov: "A", Dec: "A"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "urtarrila", Feb: "otsaila", Mar: "martxoa", Apr: "apirila", May: "maiatza", Jun: "ekaina", Jul: "uztaila", Aug: "abuztua", Sep: "iraila", Oct: "urria", Nov: "azaroa", Dec: "abendua"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ig.", Mon: "al.", Tue: "ar.", Wed: "az.", Thu: "og.", Fri: "or.", Sat: "lr."}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "I", Mon: "A", Tue: "A", Wed: "A", Thu: "O", Fri: "O", Sat: "L"}, Short: cldr.CalendarDayFormatNameValue{Sun: "ig.", Mon: "al.", Tue: "ar.", Wed: "az.", Thu: "og.", Fri: "or.", Sat: "lr."}, Wide: cldr.CalendarDayFormatNameValue{Sun: "igandea", Mon: "astelehena", Tue: "asteartea", Wed: "asteazkena", Thu: "osteguna", Fri: "ostirala", Sat: "larunbata"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_eu]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "−", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", Percent: "%\u00a0#,##0"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Arabiar Emirerri Batuetako dirhama", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "afgani afganiarra", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "lek albaniarra", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "dram armeniarra", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Holandarren Antilletako florina", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "kwanza angolarra", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "peso argentinarra", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "dolar australiarra", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "florin arubarra", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "manat azerbaijandarra", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "marko bihurgarri bosniarra", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "dolar barbadostarra", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "taka bangladeshtarra", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "lev bulgariarra", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "dinar bahraindarra", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "franko burundiarra", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "dolar bermudarra", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "dolar bruneitarra", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "boliviano boliviarra", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "erreal brasildarra", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "dolar bahamarra", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "ngultrum bhutandarra", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "pula botswanarra", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "errublo bielorrusiarra", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Bielorrusiako errubloa (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "dolar belizetarra", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "dolar kanadarra", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "franko kongoarra", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "franko suitzarra", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "peso txiletarra", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "yuan txinatarra (itsasoz haraindikoa)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "yuan txinatarra", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "peso kolonbiarra", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "Costa Ricako colona", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "peso bihurgarri kubatarra", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "peso kubatarra", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "ezkutu caboverdetarra", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "koroa txekiarra", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "franko djibutiarra", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "koroa danimarkarra", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "peso dominikarra", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "dinar aljeriarra", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "libera egiptoarra", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "nakfa eritrearra", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "birr etiopiarra", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "euroa", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "dolar fijiarra", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "libera falklandarra", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "libera esterlina", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "lari georgiarra", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "cedi ghanatarra", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "libera gibraltartarra", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "dalasi ganbiarra", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "franko ginearra", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "ketzal guatemalarra", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "dolar guyanarra", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "dolar hongkongtarra", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "lempira hodurastarra", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "kuna kroaziarra", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "gourde haitiarra", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "florin hungariarra", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "errupia indonesiarra", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "shekel israeldar berria", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "errupia indiarra", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "dinar irakiarra", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "rial irandarra", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "koroa islandiarra", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "Jamaikako dolarra", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "dinar jordaniarra", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "yen japoniarra", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "txelin kenyarra", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "som kirgizistandarra", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "riel kanbodiarra", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "franko komoretarra", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "won iparkorearra", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "won hegokorearra", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "dinar kuwaitarra", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "dolar kaimandarra", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "tenge kazakhstandarra", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "kip laostarra", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "libera libanoarra", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "errupia srilankarra", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "dolar liberiarra", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Lesothoko lotia", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Lituaniako litasa", Symbol: "LTL"},
				currency.LVL: cldr.Currency{DisplayName: "Letoniako latsa", Symbol: "LVL"},
				currency.LYD: cldr.Currency{DisplayName: "dinar libiarra", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "dirham marokoarra", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "leu moldaviarra", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "ariary madagaskartarra", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "dinar mazedoniarra", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "kyat myanmartarra", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "tugrik mongoliarra", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "pataca macauarra", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Mauritaniako ouguiya (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "uguiya mauritaniarra", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "errupia mauriziarra", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "errupia maldivarra", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "kwacha malawiarra", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "peso mexikarra", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "ringgit malaysiarra", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "metical mozambiketarra", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "dolar namibiarra", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "naira nigeriarra", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "córdoba nikaraguarra", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "koroa norvegiarra", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "errupia nepaldarra", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "dolar zeelandaberritarra", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "rial omandarra", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "balboa panamarra", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "sol perutarra", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "kina gineaberriarra", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "peso filipinarra", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "errupia pakistandarra", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "zloty poloniarra", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "guarani paraguaitarra", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "riyal qatartarra", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "leu errumaniarra", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "dinar serbiarra", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "errublo errusiarra", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "franko ruandarra", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "riyal saudiarabiarra", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "dolar salomondarra", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "errupia seychelletarra", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "libera sudandarra", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "koroa suediarra", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "dolar singapurtarra", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Santa Helenako libera", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "leone sierraleonarra", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "txelin somaliarra", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "dolar surinamdarra", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "libera hegosudandarra", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "Sao Tomeko eta Principeko dobra (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "dobra saotometarra", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "libera siriarra", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "lilangeni swazilandiarra", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "baht thailandiarra", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "somoni tajikistandarra", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "manat turkmenistandarra", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "dinar tunisiarra", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Tongako Paʻanga", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "lira turkiarra", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Trinidad eta Tobagoko dolarra", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "dolar taiwandar berria", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "txelin tanzaniarra", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Ukrainako hryvnia", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "txelin ugandarra", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "dolar estatubatuarra", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "peso uruguaitarra", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "sum uzbekistandarra", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "Venezuelako bolivarra (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "bolivar venezuelarra", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "dong vietnamdarra", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "vatu vanuatuarra", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "tala samoarra", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Afrika erdialdeko CFA frankoa", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Karibe ekialdeko dolarra", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Afrika mendebaldeko CFA frankoa", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "CFP frankoa", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "moneta ezezaguna", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "rial yemendarra", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "rand hegoafrikarra", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Zambiako kwacha (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "kwacha zambiarra", Symbol: "ZMW"},
			},
		},
	}
}