// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_sg() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "sg",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Nye", Feb: "Ful", Mar: "Mbä", Apr: "Ngu", May: "Bêl", Jun: "Fön", Jul: "Len", Aug: "Kük", Sep: "Mvu", Oct: "Ngb", Nov: "Nab", Dec: "Kak"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "N", Feb: "F", Mar: "M", Apr: "N", May: "B", Jun: "F", Jul: "L", Aug: "K", Sep: "M", Oct: "N", Nov: "N", Dec: "K"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Nyenye", Feb: "Fulundïgi", Mar: "Mbängü", Apr: "Ngubùe", May: "Bêläwü", Jun: "Föndo", Jul: "Lengua", Aug: "Kükürü", Sep: "Mvuka", Oct: "Ngberere", Nov: "Nabändüru", Dec: "Kakauka"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Bk1", Mon: "Bk2", Tue: "Bk3", Wed: "Bk4", Thu: "Bk5", Fri: "Lâp", Sat: "Lây"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "K", Mon: "S", Tue: "T", Wed: "S", Thu: "K", Fri: "P", Sat: "Y"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Bikua-ôko", Mon: "Bïkua-ûse", Tue: "Bïkua-ptâ", Wed: "Bïkua-usïö", Thu: "Bïkua-okü", Fri: "Lâpôsö", Sat: "Lâyenga"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "ND", PM: "LK"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "ND", PM: "LK"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_sg]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "", Currency: "¤#,##0.00;¤-#,##0.00", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "dirâm tî âEmirâti tî Arâbo Ôko", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "kwânza tî Angoläa", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "dolära tî Ostralïi", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "dolùara tî Bahrâina", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "farânga tî Burundïi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "pûla tî Botswana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "dolära tî kanadäa", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "farânga tî Kongöo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "farânga tî Sûîsi", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "yuan renminbi tî Shîni", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "eskûêdo tî Kâpo-Vêre", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "farânga tî Dibutïi", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "dinäri tî Alzerïi", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "pôndo tî Kâmitâ", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "nakafa tî Eritrëe", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "bir tî Etiopïi", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "zoröo", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "pôndo tî Anglëe", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "sêdi tî Ganäa", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "dalasi tî gambïi", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "sili tî Ginëe", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "rupïi tî Ênnde", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "yêni tî Zapön", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "shilîngi tî Kenyäa", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "farânga tî Kömôro", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "dolära tî Liberïa", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "loti tî Lesôtho", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "dinäar tî Libïi", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "dirâm tî Marôko", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "ariâri tî Madagasikära", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "ugîya tî Moritanïi (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "ugîya tî Moritanïi", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "rupïi tî Mörîsi", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "kwâtia tî Malawïi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "metikala tî Mozambîka", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "dolära tî Namibïi", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "nâîra tî Nizerïa", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "farânga tî Ruandäa", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "riâli tî Saûdi Arabïi", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "rupïi tî Sëyshêle", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "pôndo tî Sudäan", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "pôndo tî Zûâ Sênt-Helêna", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "leône tî Sierâ-Leône", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "shilîngi tî Somalïi", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "dôbra tî Sâô Tomë na Prinsîpe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "dôbra tî Sâô Tomë na Prinsîpe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "lilangùeni tî Swazïlânde", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "dinära tî Tunizïi", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "shilîngi tî Tanzanïi", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "shilîngi tî Ugandäa", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "dol$ara ttî äLetäa-Ôko tî Amerîka", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "farânga CFA (BEAC)", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "farânga CFA (BCEAO)", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "rânde tî Mbongo-Afrîka", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "kwâtia tî Zambïi (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "kwâtia tî Zambïi", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "dolära tî Zimbäbwe", Symbol: ""},
			},
		},
	}
}
