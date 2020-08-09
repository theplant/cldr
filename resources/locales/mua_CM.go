// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_mua_CM() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "FLO", Feb: "CLA", Mar: "CKI", Apr: "FMF", May: "MAD", Jun: "MBI", Jul: "MLI", Aug: "MAM", Sep: "FDE", Oct: "FMU", Nov: "FGW", Dec: "FYU"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "O", Feb: "A", Mar: "I", Apr: "F", May: "D", Jun: "B", Jul: "L", Aug: "M", Sep: "E", Oct: "U", Nov: "W", Dec: "Y"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Fĩi Loo", Feb: "Cokcwaklaŋne", Mar: "Cokcwaklii", Apr: "Fĩi Marfoo", May: "Madǝǝuutǝbijaŋ", Jun: "Mamǝŋgwãafahbii", Jul: "Mamǝŋgwãalii", Aug: "Madǝmbii", Sep: "Fĩi Dǝɓlii", Oct: "Fĩi Mundaŋ", Nov: "Fĩi Gwahlle", Dec: "Fĩi Yuru"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Cya", Mon: "Cla", Tue: "Czi", Wed: "Cko", Thu: "Cka", Fri: "Cga", Sat: "Cze"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Y", Mon: "L", Tue: "Z", Wed: "O", Thu: "A", Fri: "G", Sat: "E"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Com’yakke", Mon: "Comlaaɗii", Tue: "Comzyiiɗii", Wed: "Comkolle", Thu: "Comkaldǝɓlii", Fri: "Comgaisuu", Sat: "Comzyeɓsuu"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "comme", PM: "lilli"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "comme", PM: "lilli"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_root]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Solai Arabiya", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "solai Angola", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "solai Australya", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "solai Barenya", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "solai Burundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "solai Botswana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "solai Kanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "solai Kongo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Solai Swiss", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "solai Syiŋ", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "solai Kapverdiya", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "solai Djibouti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "solai Algerya", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "solai Egypt", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "solai Eritre", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "solai Etiopia", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "solai Euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "solai Britaniya", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "solai Gana", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "solai Gambiya", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "solai Guine", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "solai India", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "solai Japoŋ", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "solai Kenia", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "solai Komorya", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "solai Liberiya", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "solai Lesotho", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "solai Libya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Solai Marok", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Solai Malagasya", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Solai Mauritaniya (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Solai Mauritaniya", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Solai Mauricǝ", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Solai Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Solai Mozambika", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Solai Namibiya", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Solai Nigeriya", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "", Symbol: "NZ$"},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "", Symbol: "₲"},
				currency.RON: cldr.Currency{DisplayName: "", Symbol: "lei"},
				currency.RUB: cldr.Currency{DisplayName: "", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Solai Rwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Solai Saudiya", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Solai Saichel", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Solai Sudaŋ ma dii ne dinar", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Solai Sudaŋ ma dii ne livre", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Solai Helena", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "solai Sierra leonǝ", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Solai Somaliya", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Solai Sao Tome (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Solai Sao Tome", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "solai Swaziland", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Solai Tunisiya", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Solai Tanzaniya", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Solai Uganda", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Solai Amerika", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "solai BEAC", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "solai BCEAO", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Solai Africa nekǝsǝŋ", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Solai Zambiya (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Solai Zambiya", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Solai Zimbabwe", Symbol: ""},
			},
		},
	}
}
