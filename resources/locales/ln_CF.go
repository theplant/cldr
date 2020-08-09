// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_ln_CF() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "ln_CF",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "yan", Feb: "fbl", Mar: "msi", Apr: "apl", May: "mai", Jun: "yun", Jul: "yul", Aug: "agt", Sep: "stb", Oct: "ɔtb", Nov: "nvb", Dec: "dsb"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "y", Feb: "f", Mar: "m", Apr: "a", May: "m", Jun: "y", Jul: "y", Aug: "a", Sep: "s", Oct: "ɔ", Nov: "n", Dec: "d"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "sánzá ya yambo", Feb: "sánzá ya míbalé", Mar: "sánzá ya mísáto", Apr: "sánzá ya mínei", May: "sánzá ya mítáno", Jun: "sánzá ya motóbá", Jul: "sánzá ya nsambo", Aug: "sánzá ya mwambe", Sep: "sánzá ya libwa", Oct: "sánzá ya zómi", Nov: "sánzá ya zómi na mɔ̌kɔ́", Dec: "sánzá ya zómi na míbalé"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "eye", Mon: "ybo", Tue: "mbl", Wed: "mst", Thu: "min", Fri: "mtn", Sat: "mps"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "e", Mon: "y", Tue: "m", Wed: "m", Thu: "m", Fri: "m", Sat: "p"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "eyenga", Mon: "mokɔlɔ mwa yambo", Tue: "mokɔlɔ mwa míbalé", Wed: "mokɔlɔ mwa mísáto", Thu: "mokɔlɔ ya mínéi", Fri: "mokɔlɔ ya mítáno", Sat: "mpɔ́sɔ"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "ntɔ́ngɔ́", PM: "mpókwa"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "ntɔ́ngɔ́", PM: "mpókwa"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_ln]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Dirihamɛ ya Lémila alabo", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "Lek", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Kwanza ya Angóla", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "Peso y’Argentina", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Dolarɛ ya Ositali", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Guldeni y’ Aruba", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Mark ya kobóngwama", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "Dolále ya Barbados", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BGN: cldr.Currency{DisplayName: "Lev ya Bulgaria", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Dinarɛ ya Bahrɛnɛ", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Falánga ya Burundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "Boliviano", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "Real ya Brazil", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Dolále ya Bahamas", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Pula ya Botswana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "Rubelé ya Bielorusí", Symbol: "р."},
				currency.BYR: cldr.Currency{DisplayName: "Rubelé ya Bielorusí (2000–2016)", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "Dolále ya Belíze", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Dolarɛ ya Kanadá", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Falánga ya Kongó", Symbol: "FC"},
				currency.CHF: cldr.Currency{DisplayName: "Falánga ya Swisɛ", Symbol: "Fr."},
				currency.CLP: cldr.Currency{DisplayName: "Peso ya Shili", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Yuanɛ Renminbi ya Sinɛ", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Peso ya Kolombi", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "Colon ya Kosta Rika", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "Peso ya Kuba", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Esikudo ya Kapevɛrɛ", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Motolé Sheki", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "Falánga ya Dzibuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "Motolé ya Danemark", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Peso Dominikani", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dinarɛ ya Alizeri", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Paunɛ ya Ezípitɛ", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfa ya Elitlɛ", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Birɛ ya Etsiópi", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Ɛlɔ́", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Dolále ya Fiji", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Paunɛ ya Angɛlɛtɛ́lɛ", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Sedi ya Gana", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Cedi", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "Bojito ya Gibraltar", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi ya Gambi", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "Falánga ya Gine", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Falánga ya Ginɛ", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "Quetzal", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Lempira", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "Kuna ya Kroasia", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Gurde", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "Folinte", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupi ya Índɛ", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "Motolé ya Islandi", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "Dolále ya Jamaïke", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Yeni ya Zapɔ", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Shilingɛ ya Kenya", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Falánga ya Komoro", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "Dolále ya Bisanga bya Kayman", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Dolarɛ ya Liberya", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Loti ya Lesóto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Litas ya Litwani", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "Lats ya Letoni", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Dinarɛ ya Libí", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Dirihame ya Marokɛ", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Falánga ya Madagasikarɛ", Symbol: "Ar"},
				currency.MKD: cldr.Currency{DisplayName: "Denalé", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Ugwiya ya Moritani (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Ugwiya ya Moritani", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupi ya Morisi", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Kwasha ya Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "Peso ya Mexiko", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Metikali ya Mozambiki", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Métikal", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Dolarɛ ya Namibi", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Naira ya Nizerya", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "Motolé ya Norvej", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "Dolále ya Zeland ya Sika", Symbol: "NZ$"},
				currency.PAB: cldr.Currency{DisplayName: "Balboa", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Sol Sika", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Kina", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "Sloty", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "Guarani", Symbol: "₲"},
				currency.RON: cldr.Currency{DisplayName: "Leu Sika", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "Dinalé ya Serbia", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Rubelé ya Rusí", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Falánga ya Rwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Riyalɛ ya Alabi Sawuditɛ", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "Dolále ya Bisanga Solomoni", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Rupi ya Sɛshɛlɛ", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Dinarɛ ya Sudá", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Paunɛ ya Sudá", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Motolé ya Swédi", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Paunɛ ya Sántu elena", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Leonɛ", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Shilingɛ ya Somali", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "Bojito ya Sudaní ya Súdi", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Dobra ya Sao Tomé mpé Presipe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Dobra ya Sao Tomé mpé Presipe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Dinarɛ ya Tinizi", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "Pa’Anga", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "Dolále ya Trinidad mpé Tobago", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Shilingɛ ya Tanzani", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "Griwná", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "Shilingɛ ya Uganda", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Dolarɛ ya Ameriki", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Vatu", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "Tala", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "Falánga CFA BEAC", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Dolále ya Kalibí Monyɛlɛ", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Falánga CFA BCEAO", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "Falánga CFP", Symbol: "F CFP"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Randɛ ya Afríka Súdi", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwasha ya Zambi (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kwasha ya Zambi", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Dolarɛ ya Zimbabwɛ", Symbol: ""},
			},
		},
	}
}
