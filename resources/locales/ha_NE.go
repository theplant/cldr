// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_ha_NE() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "ha_NE",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "d/M/yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Fab", Mar: "Mar", Apr: "Afi", May: "May", Jun: "Yun", Jul: "Yul", Aug: "Agu", Sep: "Sat", Oct: "Okt", Nov: "Nuw", Dec: "Dis"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "Y", Jul: "Y", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Janairu", Feb: "Faburairu", Mar: "Maris", Apr: "Afirilu", May: "Mayu", Jun: "Yuni", Jul: "Yuli", Aug: "Agusta", Sep: "Satumba", Oct: "Oktoba", Nov: "Nuwamba", Dec: "Disamba"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Lah", Mon: "Lit", Tue: "Tal", Wed: "Lar", Thu: "Alh", Fri: "Jum", Sat: "Asa"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "L", Mon: "L", Tue: "T", Wed: "L", Thu: "A", Fri: "J", Sat: "A"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Lh", Mon: "Li", Tue: "Ta", Wed: "Lr", Thu: "Al", Fri: "Ju", Sat: "As"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Lahadi", Mon: "Litinin", Tue: "Talata", Wed: "Laraba", Thu: "Alhamis", Fri: "Jummaʼa", Sat: "Asabar"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_ha]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "", Group: "", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "Dubu 0", Currency: "¤\u00a00D", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Kuɗin Haɗaɗɗiyar Daular Larabawa", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Afghani na kasar Afghanistan", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "Kudin Albanian", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "Kudin Armenian", Symbol: ""},
				currency.ANG: cldr.Currency{DisplayName: "Antillean Guilder na ƙasar Netherlands", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Kuɗin Angola", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "Peso na ƙasar Argentina", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Dalar Ostareliya", Symbol: "$"},
				currency.AWG: cldr.Currency{DisplayName: "Florin na yankin Aruba", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Kudin Azerbaijani", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Kudaden Bosnia da Harzegovina", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "Dalar ƙasar Barbados", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "Taka na kasar Bangladesh", Symbol: "৳"},
				currency.BGN: cldr.Currency{DisplayName: "Kudin Bulgeria", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Kuɗin Baharan", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Kuɗin Burundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "Dalar ƙasar Bermuda", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "Kudin Brunei", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "Boloviano na ƙasar Bolivia", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "Ril Kudin Birazil", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Dalar ƙasar Bahamas", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "Ngultrum na kasar Bhutan", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Kuɗin Baswana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "Kudin Belarusian", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "Dalar ƙasar Belize", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Dalar Kanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Kuɗin Kongo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Kuɗin Suwizalan", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Peso na ƙasar Chile", Symbol: "$"},
				currency.CNH: cldr.Currency{DisplayName: "Yuan na ƙasar Sin (na wajen ƙasa)", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "Yuwan kasar Sin", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Peso na ƙasar Columbia", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "Colón na kasar Costa Rica", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "Peso mai fuska biyu na ƙasar Cuba", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "Peso na ƙasar Cuba", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Kuɗin Tsibiran Kap Barde", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Kudin Czech", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Kuɗin Jibuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "Krone na ƙasar Denmark", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "Peso na jamhuriyar Dominica", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Kuɗin Aljeriya", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Fam kin Masar", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Kuɗin Eritireya", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Kuɗin Habasha", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Yuro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Dalar Fiji", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "Fam na ƙasar Tsibirai na Falkland", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Fam na Ingila", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Kudin Georgian", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Cedi", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Kudin Ghana", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "Kudin Gibraltar", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Kuɗin Gambiya", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "Kudin Guinean", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Kuɗin Gini", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "Quetzal na ƙasar Guatemala", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "Dalar Guyana", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "Dalar Hong Kong", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Lempira na ƙasar Honduras", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "Kudin Croatian", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "Gourde na ƙasar Haiti", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "Kudin Hungarian", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "Rupiah na ƙasar Indonesia", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "Sabbin Kudin Shekel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Kuɗin Indiya", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Dinarin Iraqi", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "Riyal na kasar Iran", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Króna na ƙasar Iceland", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "Dalar Jamaica", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "Dinarin Jordanian", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "Yen kasar Japan", Symbol: "¥"},
				currency.KES: cldr.Currency{DisplayName: "Sulen Kenya", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "Som na ƙasar Kyrgystani", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "Riel na ƙasar Cambodia", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Kuɗin Kwamoras", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "Won na ƙasar Koreya ta Arewa", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "Won na Koreya ta Kudu", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Dinarin Kuwaiti", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "Dalar ƙasar Tsibirai na Cayman", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "Tenge na ƙasar Kazkhstan", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "Kudin Laotian", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "Kudin Lebanese", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "Rupee na kasar Sri Lanka", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Dalar Laberiya", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Kuɗin Lesoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Kuɗin Libiya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Kuɗin Maroko", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "kudaden Moldova", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Kuɗin Madagaskar", Symbol: "Ar"},
				currency.MKD: cldr.Currency{DisplayName: "Dinarin Macedonian", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Kudin Myanmar", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "Tugrik na Mongolia", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "Pataca na ƙasar Macao", Symbol: ""},
				currency.MRO: cldr.Currency{DisplayName: "Kuɗin Moritaniya (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Kuɗin Moritaniya", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Kuɗin Moritus", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "Rufiyaa na kasar Maldives", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "Kuɗin Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "Peso na ƙasar Mexico", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Kudin Malaysian", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Kuɗin Mozambik", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Metical na ƙasar Mozambique", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Dalar Namibiya", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Nairar Najeriya", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "Córdoba na ƙasar Nicaragua", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "Krone na ƙasar Norway", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "Rupee na Nepal", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "Dalar New Zealand", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Riyal din Omani", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "Balboa na ƙasar Panama", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Sol na ƙasar Peru", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Kina na ƙasar Papua Sabon Guinea", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "Kudin Philippine", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "Rupee na kasar Pakistan", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "Kudaden Polish", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "Guarani na ƙasar Paraguay", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "Riyal din Qatari", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Kudin Romanian", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "Dinar Serbian", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "Ruble kasar Rasha", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Kuɗin Ruwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Riyal", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "Dalar Tsibirai na Solomon", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Kuɗin Saishal", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Fam kin Sudan", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Krona na ƙasar Sweden", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "Dilar Singapore", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Fam kin San Helena", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Kuɗin Salewo", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Sulen Somaliya", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "Dalar ƙasar Suriname", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "Fam na Kudancin Sudan", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Kuɗin Sawo Tome da Paransip (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Kuɗin Sawo Tome da Paransip", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "Kudin Syrian", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Kuɗin Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "Baht na ƙasar Thailand", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "Somoni na ƙasar Tajikistan", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Manat na ƙasar Turkmenistan", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "Kuɗin Tunisiya", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "Paʼanga na ƙasar Tonga", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "Kudin Turkish", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "Dalar ƙasar Trinidad da Tobago", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "Sabuwar Dalar Taiwan", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Sulen Tanzaniya", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "Kudin Ukrainian", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Sule Yuganda", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Dalar Amirka", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "Peso na ƙasar Uruguay", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "Som na ƙasar Uzbekistan", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "Bolívar na ƙasar Venezuela", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "Kudin Vietnamese", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Vatu da ƙasar Vanuatu", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "Tala na ƙasar Samoa", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "Kuɗin Sefa na Afirka Ta Tsakiya", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Dalar Gabashin Caribbean", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Kuɗin Sefa na Afirka Ta Yamma", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "kudin CFP Franc", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Kudin da ba a sani ba", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Riyal din Yemeni", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Kuɗin Afirka Ta Kudu", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Kuɗin Zambiya (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kuɗin Zambiya", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Dalar zimbabuwe", Symbol: ""},
			},
		},
	}
}
