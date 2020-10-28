// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_ha_NG() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "ha_NG",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "d/M/yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Fab", Mar: "Mar", Apr: "Afi", May: "May", Jun: "Yun", Jul: "Yul", Aug: "Agu", Sep: "Sat", Oct: "Okt", Nov: "Nuw", Dec: "Dis"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "Y", Jul: "Y", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Janairu", Feb: "Faburairu", Mar: "Maris", Apr: "Afirilu", May: "Mayu", Jun: "Yuni", Jul: "Yuli", Aug: "Agusta", Sep: "Satumba", Oct: "Oktoba", Nov: "Nuwamba", Dec: "Disamba"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Lah", Mon: "Lit", Tue: "Tal", Wed: "Lar", Thu: "Alh", Fri: "Jum", Sat: "Asa"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "L", Mon: "L", Tue: "T", Wed: "L", Thu: "A", Fri: "J", Sat: "A"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Lh", Mon: "Li", Tue: "Ta", Wed: "Lr", Thu: "Al", Fri: "Ju", Sat: "As"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Lahadi", Mon: "Litinin", Tue: "Talata", Wed: "Laraba", Thu: "Alhamis", Fri: "Jummaʼa", Sat: "Asabar"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
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
		Languages: cldr.Languages{
			language.AF:      "Afirkanci",
			language.AGQ:     "Aghem",
			language.AK:      "Akan",
			language.AM:      "Amharik",
			language.AR:      "Larabci",
			language.AR_001:  "Larabci Asali Na Zamani",
			language.AS:      "Asamisanci",
			language.ASA:     "Asu",
			language.AST:     "Asturian",
			language.AZ:      "Azeri",
			language.BAS:     "Basaa",
			language.BE:      "Belarusanci",
			language.BEM:     "Bemba",
			language.BEZ:     "Bena",
			language.BG:      "Bulgaranci",
			language.BM:      "Bambara",
			language.BN:      "Bengali",
			language.BO:      "Tibetan",
			language.BR:      "Buretananci",
			language.BRX:     "Bodo",
			language.BS:      "Bosniyanci",
			language.CA:      "Kataloniyanci",
			language.CCP:     "Chakma",
			language.CE:      "Chechen",
			language.CEB:     "Cebuano",
			language.CGG:     "Chiga",
			language.CHR:     "Cherokee",
			language.CKB:     "Kurdish na Tsaka",
			language.CO:      "Corsican",
			language.CS:      "Harshen Cak",
			language.CU:      "Church Slavic",
			language.CY:      "Kabilar Welsh",
			language.DA:      "Danish",
			language.DAV:     "Taita",
			language.DE:      "Jamusanci",
			language.DE_AT:   "Jamusanci Ostiriya",
			language.DE_CH:   "Jamusanci Suwizalan",
			language.DJE:     "Zarma",
			language.DSB:     "Lower Sorbian",
			language.DUA:     "Duala",
			language.DYO:     "Jola-Fonyi",
			language.DZ:      "Dzongkha",
			language.EBU:     "Embu",
			language.EE:      "Ewe",
			language.EL:      "Girkanci",
			language.EN:      "Turanci",
			language.EN_AU:   "Turanci Ostareliya",
			language.EN_CA:   "Turanci Kanada",
			language.EN_GB:   "Turanci Ingila",
			language.EN_US:   "Amirka Turanci",
			language.EO:      "Dʼan/ʼYar Kabilar Andalus",
			language.ES:      "Sifaniyanci",
			language.ES_419:  "Sifaniyancin Latin Amirka",
			language.ES_ES:   "Sifaniyanci Turai",
			language.ES_MX:   "Sifaniyanci Mesiko",
			language.ET:      "Istoniyanci",
			language.EU:      "Basque",
			language.EWO:     "Ewondo",
			language.FA:      "Parisanci",
			language.FF:      "Fulah",
			language.FI:      "Yaren mutanen Finland",
			language.FIL:     "Dan Filifin",
			language.FO:      "Faroese",
			language.FR:      "Faransanci",
			language.FR_CA:   "Farasanci Kanada",
			language.FR_CH:   "Farasanci Suwizalan",
			language.FUR:     "Friulian",
			language.FY:      "Kʼabilan Firsi",
			language.GA:      "Dan Ailan",
			language.GD:      "Kʼabilan Scots Gaelic",
			language.GL:      "Bagalike",
			language.GN:      "Guwaraniyanci",
			language.GSW:     "Jamusanci Swiss",
			language.GU:      "Gujarati",
			language.GUZ:     "Gusii",
			language.GV:      "Manx",
			language.HA:      "Hausa",
			language.HAW:     "Hawaiian",
			language.HE:      "Ibrananci",
			language.HI:      "Harshen Hindi",
			language.HMN:     "Hmong",
			language.HR:      "Kuroshiyan",
			language.HSB:     "Sorbianci ta Sama",
			language.HT:      "Haitian Creole",
			language.HU:      "Harshen Hungari",
			language.HY:      "Armeniyanci",
			language.IA:      "Yare Tsakanin Kasashe",
			language.ID:      "Harshen Indunusiya",
			language.IE:      "Intagulanci",
			language.IG:      "Inyamuranci",
			language.II:      "Sichuan Yi",
			language.IS:      "Yaren mutanen Iceland",
			language.IT:      "Italiyanci",
			language.JA:      "Japananci",
			language.JGO:     "Ngomba",
			language.JMC:     "Machame",
			language.JV:      "Jabananci",
			language.KA:      "Jojiyanci",
			language.KAB:     "Kabyle",
			language.KAM:     "Kamba",
			language.KDE:     "Makonde",
			language.KEA:     "Kabuverdianu",
			language.KHQ:     "Koyra Chiini",
			language.KI:      "Kikuyu",
			language.KK:      "Kazakh",
			language.KKJ:     "Kako",
			language.KL:      "Kalaallisut",
			language.KLN:     "Kalenjin",
			language.KM:      "Harshen Kimar",
			language.KN:      "Kannada",
			language.KO:      "Harshen Koreya",
			language.KOK:     "Konkani",
			language.KS:      "Kashmiri",
			language.KSB:     "Shambala",
			language.KSF:     "Bafia",
			language.KSH:     "Colognian",
			language.KU:      "Kurdanci",
			language.KW:      "Cornish",
			language.KY:      "Kirgizanci",
			language.LA:      "Dan Kabilar Latin",
			language.LAG:     "Langi",
			language.LB:      "Luxembourgish",
			language.LG:      "Ganda",
			language.LKT:     "Lakota",
			language.LN:      "Lingala",
			language.LO:      "Laothian",
			language.LRC:     "Northern Luri",
			language.LT:      "Lituweniyanci",
			language.LU:      "Luba-Katanga",
			language.LUO:     "Luo",
			language.LUY:     "Luyia",
			language.LV:      "Latbiyanci",
			language.MAS:     "Harshen Masai",
			language.MER:     "Meru",
			language.MFE:     "Morisyen",
			language.MG:      "Malagasy",
			language.MGH:     "Makhuwa-Meetto",
			language.MGO:     "Metaʼ",
			language.MI:      "Maori",
			language.MK:      "Dan Masedoniya",
			language.ML:      "Kabilar Maleyalam",
			language.MN:      "Mongolian",
			language.MR:      "Kʼabilan Marathi",
			language.MS:      "Harshen Malai",
			language.MT:      "Harshen Maltis",
			language.MUA:     "Mundang",
			language.MUL:     "Harsuna masu yawa",
			language.MY:      "Burmanci",
			language.MZN:     "Mazanderani",
			language.NAQ:     "Nama",
			language.NB:      "Norwegian Bokmål",
			language.ND:      "North Ndebele",
			language.NDS:     "Low German",
			language.NE:      "Nepali",
			language.NL:      "Holanci",
			language.NMG:     "Kwasio",
			language.NN:      "Norwegian Nynorsk",
			language.NNH:     "Ngiemboon",
			language.NO:      "Yaren mutanen Norway",
			language.NUS:     "Nuer",
			language.NY:      "Nyanja",
			language.NYN:     "Nyankole",
			language.OC:      "Ositanci",
			language.OM:      "Oromo",
			language.OR:      "Oriyanci",
			language.OS:      "Ossetic",
			language.PA:      "Punjabi",
			language.PL:      "Harshen Polan",
			language.PRG:     "Ferusawa",
			language.PS:      "Pashtanci",
			language.PT:      "Harshen Fotugis",
			language.PT_PT:   "Fotugis kasashen Turai",
			language.QU:      "Quechua",
			language.RM:      "Romansh",
			language.RN:      "Rundi",
			language.RO:      "Romaniyanci",
			language.ROF:     "Rombo",
			language.RU:      "Rashanci",
			language.RW:      "Kiniyaruwanda",
			language.RWK:     "yaren Rwa",
			language.SA:      "Sanskrit",
			language.SAH:     "Sakha",
			language.SAQ:     "Samburu",
			language.SBP:     "Sangu",
			language.SD:      "Sindiyanci",
			language.SE:      "Northern Sami",
			language.SEH:     "Sena",
			language.SES:     "Koyraboro Senni",
			language.SG:      "Sango",
			language.SH:      "Kuroweshiyancin-Sabiya",
			language.SHI:     "Tachelhit",
			language.SI:      "Sinhalanci",
			language.SK:      "Basulke",
			language.SL:      "Basulabe",
			language.SM:      "Samoan",
			language.SMN:     "Inari Sami",
			language.SN:      "Shona",
			language.SO:      "Somalianci",
			language.SQ:      "Albanian",
			language.SR:      "Sabiyan",
			language.ST:      "Sesotanci",
			language.SU:      "Sudananci",
			language.SV:      "Harshen Suwedan",
			language.SW:      "Harshen Suwahili",
			language.TA:      "Tamil",
			language.TE:      "Dʼan/ʼYar Kabilar Telug",
			language.TEO:     "Teso",
			language.TG:      "Tajik",
			language.TH:      "Thai",
			language.TI:      "Tigriyanci",
			language.TK:      "Tukmenistanci",
			language.TLH:     "Klingon",
			language.TO:      "Tongan",
			language.TR:      "Harshen Turkiyya",
			language.TT:      "Tatar",
			language.TW:      "Tiwiniyanci",
			language.TWQ:     "Tasawaq",
			language.TZM:     "Tamazight na Atlas Tsaka",
			language.UG:      "Ugiranci",
			language.UK:      "Harshen Yukuren",
			language.UND:     "Harshen da ba a sani ba",
			language.UR:      "Urdawa",
			language.UZ:      "Uzbek",
			language.VAI:     "Vai",
			language.VI:      "Harshen Biyetinam",
			language.VO:      "Volapük",
			language.VUN:     "Vunjo",
			language.WAE:     "Walser",
			language.WO:      "Wolof",
			language.XH:      "Bazosa",
			language.XOG:     "Soga",
			language.YAV:     "Yangben",
			language.YI:      "Yiddish",
			language.YO:      "Yarbanci",
			language.YUE:     "Sinanci",
			language.ZGH:     "Standard Moroccan Tamazight",
			language.ZH:      "Harshen, Sinanci",
			language.ZH_HANS: "Sauƙaƙaƙƙen Sinanci",
			language.ZH_HANT: "Sinanci na gargajiya",
			language.ZU:      "Harshen Zulu",
			language.ZXX:     "Babu abun-ciki na yare",
		},
		Territories: cldr.Territories{
			territory.V_001: "Duniya",
			territory.V_002: "Afirka",
			territory.V_003: "North America",
			territory.V_005: "South America",
			territory.V_009: "Oceania",
			territory.V_011: "Afirka ta Yamma",
			territory.V_013: "Central America",
			territory.V_014: "Afirka ta Gabas",
			territory.V_015: "Arewacin Africa",
			territory.V_017: "Afirka ta Tsakiya",
			territory.V_018: "Kudancin Afirka",
			territory.V_019: "nahiyoyin Amurka",
			territory.V_021: "Arewacin Amurka",
			territory.V_029: "Caribbean",
			territory.V_030: "Gabashin Asiya",
			territory.V_034: "kudancin Asiya",
			territory.V_035: "Kudu Maso Gabashin Asiya",
			territory.V_039: "Kudancin Turai",
			territory.V_053: "Tsibarai na Austraila da New Zealand da maƙota",
			territory.V_054: "Tsibirai na New Guinea da Fiji da maƙota",
			territory.V_057: "Yankin Micronesia",
			territory.V_061: "tsibiri",
			territory.V_142: "Asiya",
			territory.V_143: "Asiya ta Tsakiya",
			territory.V_145: "Yammacin Asiya",
			territory.V_150: "Turai",
			territory.V_151: "Gabashin Turai",
			territory.V_154: "Arewacin Turai",
			territory.V_155: "Yammacin Turai",
			territory.V_202: "Sub-Saharan Africa",
			territory.V_419: "Latin America",
			territory.AC:    "Tsibirin Ascension",
			territory.AD:    "Andora",
			territory.AE:    "Haɗaɗɗiyar Daular Larabawa",
			territory.AF:    "Afaganistan",
			territory.AG:    "Antigwa da Barbuba",
			territory.AI:    "Angila",
			territory.AL:    "Albaniya",
			territory.AM:    "Armeniya",
			territory.AO:    "Angola",
			territory.AQ:    "Antatika",
			territory.AR:    "Arjantiniya",
			territory.AS:    "Samowa Ta Amurka",
			territory.AT:    "Ostiriya",
			territory.AU:    "Ostareliya",
			territory.AW:    "Aruba",
			territory.AX:    "Tsibirai na Åland",
			territory.AZ:    "Azarbaijan",
			territory.BA:    "Bosniya Harzagobina",
			territory.BB:    "Barbadas",
			territory.BD:    "Bangiladas",
			territory.BE:    "Belgiyom",
			territory.BF:    "Burkina Faso",
			territory.BG:    "Bulgariya",
			territory.BH:    "Baharan",
			territory.BI:    "Burundi",
			territory.BJ:    "Binin",
			territory.BL:    "St. Barthélemy",
			territory.BM:    "Barmuda",
			territory.BN:    "Burune",
			territory.BO:    "Bolibiya",
			territory.BQ:    "Caribbean Netherlands",
			territory.BR:    "Birazil",
			territory.BS:    "Bahamas",
			territory.BT:    "Butan",
			territory.BV:    "Tsibirin Bouvet",
			territory.BW:    "Baswana",
			territory.BY:    "Belarus",
			territory.BZ:    "Beliz",
			territory.CA:    "Kanada",
			territory.CC:    "Tsibirai Cocos (Keeling)",
			territory.CD:    "Kongo (DRC)",
			territory.CF:    "Jamhuriyar Afirka Ta Tsakiya",
			territory.CG:    "Jamhuriyar Kongo",
			territory.CH:    "Suwizalan",
			territory.CI:    "Aibari Kwas",
			territory.CK:    "Tsibiran Kuku",
			territory.CL:    "Cayile",
			territory.CM:    "Kamaru",
			territory.CN:    "Sin",
			territory.CO:    "Kolambiya",
			territory.CP:    "Tsibirin Clipperton",
			territory.CR:    "Kwasta Rika",
			territory.CU:    "Kyuba",
			territory.CV:    "Tsibiran Kap Barde",
			territory.CW:    "Kasar Curaçao",
			territory.CX:    "Tsibirin Kirsmati",
			territory.CY:    "Sifurus",
			territory.CZ:    "Jamhuriyar Cak",
			territory.DE:    "Jamus",
			territory.DG:    "Tsibirn Diego Garcia",
			territory.DJ:    "Jibuti",
			territory.DK:    "Danmark",
			territory.DM:    "Dominika",
			territory.DO:    "Jamhuriyar Dominika",
			territory.DZ:    "Aljeriya",
			territory.EA:    "Ceuta & Melilla",
			territory.EC:    "Ekwador",
			territory.EE:    "Estoniya",
			territory.EG:    "Misira",
			territory.EH:    "Yammacin Sahara",
			territory.ER:    "Eritireya",
			territory.ES:    "Sipen",
			territory.ET:    "Habasha",
			territory.EU:    "Tarayyar Turai",
			territory.EZ:    "Sashin Turai",
			territory.FI:    "Finlan",
			territory.FJ:    "Fiji",
			territory.FK:    "Tsibiran Falkilan",
			territory.FM:    "Mikuronesiya",
			territory.FO:    "Tsibirai na Faroe",
			territory.FR:    "Faransa",
			territory.GA:    "Gabon",
			territory.GB:    "Biritaniya",
			territory.GD:    "Girnada",
			territory.GE:    "Jiwarjiya",
			territory.GF:    "Gini Ta Faransa",
			territory.GG:    "Yankin Guernsey",
			territory.GH:    "Gana",
			territory.GI:    "Jibaraltar",
			territory.GL:    "Grinlan",
			territory.GM:    "Gambiya",
			territory.GN:    "Gini",
			territory.GP:    "Gwadaluf",
			territory.GQ:    "Gini Ta Ikwaita",
			territory.GR:    "Girka",
			territory.GS:    "Kudancin Geogia da Kudancin Tsibirin Sandiwic",
			territory.GT:    "Gwatamala",
			territory.GU:    "Gwam",
			territory.GW:    "Gini Bisau",
			territory.GY:    "Guyana",
			territory.HK:    "Hong Kong Babban Birnin Kasar Chana",
			territory.HM:    "Tsibirin Heard da McDonald",
			territory.HN:    "Honduras",
			territory.HR:    "Kurowaishiya",
			territory.HT:    "Haiti",
			territory.HU:    "Hungari",
			territory.IC:    "Canary Islands",
			territory.ID:    "Indunusiya",
			territory.IE:    "Ayalan",
			territory.IL:    "Iziraʼila",
			territory.IM:    "Isle na Mutum",
			territory.IN:    "Indiya",
			territory.IO:    "Yankin Birtaniya Na Tekun Indiya",
			territory.IQ:    "Iraƙi",
			territory.IR:    "Iran",
			territory.IS:    "Aisalan",
			territory.IT:    "Italiya",
			territory.JE:    "Kasar Jersey",
			territory.JM:    "Jamaika",
			territory.JO:    "Jordan",
			territory.JP:    "Jàpân",
			territory.KE:    "Kenya",
			territory.KG:    "Kirgizistan",
			territory.KH:    "Kambodiya",
			territory.KI:    "Kiribati",
			territory.KM:    "Kwamoras",
			territory.KN:    "San Kiti Da Nebis",
			territory.KP:    "Koriya Ta Arewa",
			territory.KR:    "Koriya Ta Kudu",
			territory.KW:    "Kwiyat",
			territory.KY:    "Tsibiran Kaiman",
			territory.KZ:    "Kazakistan",
			territory.LA:    "Lawas",
			territory.LB:    "Labanan",
			territory.LC:    "San Lusiya",
			territory.LI:    "Licansitan",
			territory.LK:    "Siri Lanka",
			territory.LR:    "Laberiya",
			territory.LS:    "Lesoto",
			territory.LT:    "Lituweniya",
			territory.LU:    "Lukusambur",
			territory.LV:    "latibiya",
			territory.LY:    "Libiya",
			territory.MA:    "Maroko",
			territory.MC:    "Monako",
			territory.MD:    "Maldoba",
			territory.ME:    "Mantanegara",
			territory.MF:    "St. Martin",
			territory.MG:    "Madagaskar",
			territory.MH:    "Tsibiran Marshal",
			territory.MK:    "Macedonia ta Arewa",
			territory.ML:    "Mali",
			territory.MM:    "Burma, Miyamar",
			territory.MN:    "Mangoliya",
			territory.MO:    "Babban Birnin Mulki na Chana",
			territory.MP:    "Tsibiran Mariyana Na Arewa",
			territory.MQ:    "Martinik",
			territory.MR:    "Moritaniya",
			territory.MS:    "Manserati",
			territory.MT:    "Malta",
			territory.MU:    "Moritus",
			territory.MV:    "Maldibi",
			territory.MW:    "Malawi",
			territory.MX:    "Makasiko",
			territory.MY:    "Malaisiya",
			territory.MZ:    "Mozambik",
			territory.NA:    "Namibiya",
			territory.NC:    "Kaledoniya Sabuwa",
			territory.NE:    "Nijar",
			territory.NF:    "Tsibirin Narfalk",
			territory.NG:    "Najeriya",
			territory.NI:    "Nikaraguwa",
			territory.NL:    "Holan",
			territory.NO:    "Norwe",
			territory.NP:    "Nefal",
			territory.NR:    "Nauru",
			territory.NU:    "Niyu",
			territory.NZ:    "Nuzilan",
			territory.OM:    "Oman",
			territory.PA:    "Panama",
			territory.PE:    "Feru",
			territory.PF:    "Folinesiya Ta Faransa",
			territory.PG:    "Papuwa Nugini",
			territory.PH:    "Filipin",
			territory.PK:    "Pakistan",
			territory.PL:    "Polan",
			territory.PM:    "San Piyar Da Mikelan",
			territory.PN:    "Pitakarin",
			territory.PR:    "Porto Riko",
			territory.PS:    "Palasɗinu",
			territory.PT:    "Portugal",
			territory.PW:    "Palau",
			territory.PY:    "Faragwai",
			territory.QA:    "Katar",
			territory.QO:    "Bakin Teku",
			territory.RE:    "Rawuniyan",
			territory.RO:    "Romaniya",
			territory.RS:    "Sabiya",
			territory.RU:    "Rasha",
			territory.RW:    "Ruwanda",
			territory.SA:    "Saudiyya",
			territory.SB:    "Tsibiran Salaman",
			territory.SC:    "Seychelles",
			territory.SD:    "Sudan",
			territory.SE:    "Suwedan",
			territory.SG:    "Singapur",
			territory.SH:    "San Helena",
			territory.SI:    "Sulobeniya",
			territory.SJ:    "Svalbard da Jan Mayen",
			territory.SK:    "Sulobakiya",
			territory.SL:    "Salewo",
			territory.SM:    "San Marino",
			territory.SN:    "Sanigal",
			territory.SO:    "Somaliya",
			territory.SR:    "Suriname",
			territory.SS:    "Sudan ta kudu",
			territory.ST:    "Sawo Tome Da Paransip",
			territory.SV:    "El Salbador",
			territory.SX:    "Sint Maarten",
			territory.SY:    "Sham, Siriya",
			territory.SZ:    "Eswatini",
			territory.TA:    "Tritan da Kunha",
			territory.TC:    "Turkis Da Tsibiran Kaikwas",
			territory.TD:    "Cadi",
			territory.TF:    "Yankin Faransi ta Kudu",
			territory.TG:    "Togo",
			territory.TH:    "Tailan",
			territory.TJ:    "Tajikistan",
			territory.TK:    "Takelau",
			territory.TL:    "Timor Ta Gabas",
			territory.TM:    "Turkumenistan",
			territory.TN:    "Tunisiya",
			territory.TO:    "Tonga",
			territory.TR:    "Turkiyya",
			territory.TT:    "Tirinidad Da Tobago",
			territory.TV:    "Tubalu",
			territory.TW:    "Taiwan",
			territory.TZ:    "Tanzaniya",
			territory.UA:    "Yukaran",
			territory.UG:    "Yuganda",
			territory.UM:    "Rukunin Tsibirin U.S",
			territory.UN:    "Majalisar Dinkin Duniya",
			territory.US:    "Amurka",
			territory.UY:    "Yurigwai",
			territory.UZ:    "Uzubekistan",
			territory.VA:    "Batikan",
			territory.VC:    "San Binsan Da Girnadin",
			territory.VE:    "Benezuwela",
			territory.VG:    "Tsibirin Birjin Na Birtaniya",
			territory.VI:    "Tsibiran Birjin Ta Amurka",
			territory.VN:    "Biyetinam",
			territory.VU:    "Banuwatu",
			territory.WF:    "Walis Da Futuna",
			territory.WS:    "Samoa",
			territory.XA:    "Gogewar Kwalwa",
			territory.XB:    "Gano wani abu ta hanyar amfani da fasaha",
			territory.XK:    "Kasar Kosovo",
			territory.YE:    "Yamal",
			territory.YT:    "Mayoti",
			territory.ZA:    "Afirka Ta Kudu",
			territory.ZM:    "Zambiya",
			territory.ZW:    "Zimbabuwe",
			territory.ZZ:    "Yanki da ba a sani ba",
		},
	}
}
