// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_ewo() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "ewo",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ngo", Feb: "ngb", Mar: "ngl", Apr: "ngn", May: "ngt", Jun: "ngs", Jul: "ngz", Aug: "ngm", Sep: "nge", Oct: "nga", Nov: "ngad", Dec: "ngab"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "o", Feb: "b", Mar: "l", Apr: "n", May: "t", Jun: "s", Jul: "z", Aug: "m", Sep: "e", Oct: "a", Nov: "d", Dec: "b"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "ngɔn osú", Feb: "ngɔn bɛ̌", Mar: "ngɔn lála", Apr: "ngɔn nyina", May: "ngɔn tána", Jun: "ngɔn saməna", Jul: "ngɔn zamgbála", Aug: "ngɔn mwom", Sep: "ngɔn ebulú", Oct: "ngɔn awóm", Nov: "ngɔn awóm ai dziá", Dec: "ngɔn awóm ai bɛ̌"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "sɔ́n", Mon: "mɔ́n", Tue: "smb", Wed: "sml", Thu: "smn", Fri: "fúl", Sat: "sér"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "s", Mon: "m", Tue: "s", Wed: "s", Thu: "s", Fri: "f", Sat: "s"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "sɔ́ndɔ", Mon: "mɔ́ndi", Tue: "sɔ́ndɔ məlú mə́bɛ̌", Wed: "sɔ́ndɔ məlú mə́lɛ́", Thu: "sɔ́ndɔ məlú mə́nyi", Fri: "fúladé", Sat: "séradé"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "kíkíríg", PM: "ngəgógəle"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "kíkíríg", PM: "ngəgógəle"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_root]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Dirám yá Emirá Aráb Uní", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Kwánəza yá Angolá", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Dolár yá Osətəralí", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "Dinár yá Bahərɛ́n", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Fəláŋ yá Burundí", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Púlá yá Botswána", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Dolár yá Kanáda", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Fəláŋ yá Kongo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Fəláŋ yá Suís", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Yuán Renəminəbí yá Tsainís", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Esəkúdo yá Kápə́vɛ́rə", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Fəláŋ yá dzibutí", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dinár yá Alehérí", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Lívə́lə yá Ehíbətía", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Náfəka yá Eritelé", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Bír yá Etsiópia", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "əró", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Lívə́lə Sətərəlíŋ", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Tzedí yá Ganá", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasí yá Gámbía", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Síli yá Giné", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupí yá ɛ́ndía", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Yɛ́n yá Hapɔ́n", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Silíŋ yá Keniá", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Fəláŋ yá Komória", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Dolár yá Libéria", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Lotí yá Lesotó", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Dinár yá Libí", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Dirám yá Maróg", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Ariari yá Maləgás", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Ugiya yá Moritaní (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Ugiya yá Moritaní", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupí yá Morís", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Kwatsa yá Malawí", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Metikal yá Mozambíg", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Dolár yá Namibí", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Náíra yá Nihéria", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "Fəláŋ yá Ruwandá", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Riál yá Arabí Saudí", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Rupí yá Sɛsɛ́l", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Lívələ yá Sudán", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Lívələ yá Sudán (1956–2007)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Lívələ yá Ǹfúfúb Elɛ́n", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Leóne yá Sierá-leónə", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Silíŋ yá Somalí", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Dóbə́ra yá Saó Tomé ai Pəlinəsípe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Dóbə́ra yá Saó Tomé ai Pəlinəsípe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni yá Swazí", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Dinár yá Tunisí", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Silíŋ yá Tanazaní", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Silíŋ yá Ugandá (1966–1987)", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Dolár yá amɛ́rəkə", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "Fəláŋ CFA (BEAC)", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Fəláŋ CFA (BCEAO)", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Ránədə yá Afiríka", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwatsa yá Zambí (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kwatsa yá Zambí", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Dolár yá Zimbabwé", Symbol: ""},
			},
		},
	}
}
