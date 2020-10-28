// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_jv() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "jv",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd-MM-y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mar", Apr: "Apr", May: "Mei", Jun: "Jun", Jul: "Jul", Aug: "Agt", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Des"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Januari", Feb: "Februari", Mar: "Maret", Apr: "April", May: "Mei", Jun: "Juni", Jul: "Juli", Aug: "Agustus", Sep: "September", Oct: "Oktober", Nov: "November", Dec: "Desember"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Ahad", Mon: "Sen", Tue: "Sel", Wed: "Rab", Thu: "Kam", Fri: "Jum", Sat: "Sab"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "A", Mon: "S", Tue: "S", Wed: "R", Thu: "K", Fri: "J", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Ahad", Mon: "Sen", Tue: "Sel", Wed: "Rab", Thu: "Kam", Fri: "Jum", Sat: "Sab"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Ahad", Mon: "Senin", Tue: "Selasa", Wed: "Rabu", Thu: "Kamis", Fri: "Jumat", Sat: "Sabtu"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "Isuk", PM: "Wengi"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "Isuk", PM: "Wengi"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "Isuk", PM: "Wengi"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "0 èwu", Currency: "¤\u00a0#,##0.00", CurrencyAccounting: "¤\u00a0#,##0.00", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Dirham Uni Emirat Arab", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Afghani Afganistan", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "Lek Albania", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "Dram Armenia", Symbol: ""},
				currency.ANG: cldr.Currency{DisplayName: "Guilder Antilla Walanda", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Kwanza Angola", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "Peso Argentina", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Dolar Australia", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Florin Aruban", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Manat Azerbaijan", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Mark Konvertibel Bosnia-Herzegovina", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "Dolar Barbadian", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "Taka Bangladesh", Symbol: "৳"},
				currency.BGN: cldr.Currency{DisplayName: "Lev Bulgaria", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Bahrain Dinar", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Franc Burundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "Dolar Bermuda", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "Dolar Brunai", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "Boliviano Bolivia", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "Real Brasil", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Dolar Bahamian", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "Ngultrum Bhutan", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Pula Botswana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "Ruble Belarusia", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "Dolar Belise", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Dolar Kanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Franc Kongo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Franc Swiss", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Peso Chili", Symbol: "$"},
				currency.CNH: cldr.Currency{DisplayName: "Yuan Tyongkok (Jaban Rangkah)", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "Yuan Tyongkok", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Peso Kolumbia", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "Colon Kosta Rika", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "Peso Konvertibel Kuba", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "Peso Kuba", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Escudo Tanjung Verde", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Koruna Czech", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Franc Djibouti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "Krone Denmark", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "Peso Dominika", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dinar Algeria", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Pound Mesir", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfa Eritrea", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Birr Ethiopia", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Dolar Fiji", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "Pound Kepuloan Falkland", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Pound Inggris", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Lari Georgia", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "Cedi Ghana", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "Pound Gibraltar", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi Gambia", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "Franc Guinea", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "Quetzal Guatemala", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "Dolar Guyana", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "Dolar Hong Kong", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Lempira Honduras", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "Kuna Kroasia", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "Gourde Haiti", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "Forint Hungaria", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "Rupiah Indonesia", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "Shekel Anyar Israel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupee India", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Dinar Irak", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "Rial Iran", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Krona Islandia", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "Dolar Jamaika", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "Dinar Yordania", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "Yen Jepang", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Shilling Kenya", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "Som Kirgistan", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "Riel Kamboja", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Franc Komoro", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "Won Korea Lor", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "Won Korea Kidul", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Dinar Kuwait", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "Dolar Kepuloan Caiman", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "Tenge Kasakhstan", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "Kip Laos", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "Pound Libanon", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "Rupee Sri Lanka", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Dolar Liberia", Symbol: "$"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Dinar Libya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Dirham Maroko", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Leu Moldova", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Ariary Malagasi", Symbol: "Ar"},
				currency.MKD: cldr.Currency{DisplayName: "Denar Masedonia", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Kyat Myanmar", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "Tugrik Mongol", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "Pataca Macau", Symbol: ""},
				currency.MRO: cldr.Currency{DisplayName: "Ouguiya Mauritania (1973 - 2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Ouguiya Mauritania", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupee Mauritius", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "Rufiyaa Maladewa", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "Kwacha Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "Peso Meksiko", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Ringgit Malaysia", Symbol: "RM"},
				currency.MZN: cldr.Currency{DisplayName: "Metical Mosambik", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Dolar Namibia", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Naira Nigeria", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "Cordoba Nikaragua", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "Krone Norwegia", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "Rupee Nepal", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "Dolar Selandia Anyar", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Rial Oman", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "Balboa Panama", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Sol Peru", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Kina Papua Nugini", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "Piso Filipina", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "Rupee Pakistan", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "Zloty Polandia", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "Guarani Paraguay", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "Rial Qatar", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Leu Rumania", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "Dinar Serbia", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "Rubel Rusia", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Franc Rwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Riyal Saudi", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "Dolar Kepuloan Solomon", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Rupee Seichelles", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Pound Sudan", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Krona Swedia", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "Dolar Singapura", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Pound Santa Helena", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Leone Sierra Leone", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Shilling Somalia", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "Dolar Suriname", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "Pound Sudan Kidul", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "Dobra Sao Tome lan Principe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "Pound Siria", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni Swasi", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "Baht Thai", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "Somoni Tajikistan", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Manat Turmenistan", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "Dinar Tunisia", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "Paʻanga Tonga", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "Lira Turki", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "Dolar Trinidad lan Tobago", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "Dolar Anyar Taiwan", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Shilling Tansania", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "Hryvnia Ukrania", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Shilling Uganda", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Dolar Amerika Serikat", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "Peso Uruguay", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "Som Usbekistan", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "Bolivar Venezuela (2008 - 2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Bolivar Venezuela", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "Dong Vietnam", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Vatu Vanuatu", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "Tala Samoa", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "CFA Franc Afrika Tengah", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Dolar Karibia Wetan", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "CFA Franc Afrika Kulon", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "Franc CFP", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Dhuwit Ora Dikenali", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Rial Yaman", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Rand Afrika Kidul", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "Kwacha Sambia", Symbol: "ZK"},
			},
		},
		Languages: cldr.Languages{
			language.AF:      "Afrika",
			language.AGQ:     "Aghem",
			language.AK:      "Akan",
			language.AM:      "Amharik",
			language.AR:      "Arab",
			language.AR_001:  "Arab Standar Anyar",
			language.AS:      "Assam",
			language.ASA:     "Asu",
			language.AST:     "Asturia",
			language.AZ:      "Azeri",
			language.BAS:     "Basaa",
			language.BE:      "Belarus",
			language.BEM:     "Bemba",
			language.BEZ:     "Bena",
			language.BG:      "Bulgaria",
			language.BM:      "Bambara",
			language.BN:      "Bengali",
			language.BO:      "Tibet",
			language.BR:      "Breton",
			language.BRX:     "Bodo",
			language.BS:      "Bosnia lan Hercegovina",
			language.CA:      "Katala",
			language.CCP:     "Chakma",
			language.CE:      "Chechen",
			language.CEB:     "Cebuano",
			language.CGG:     "Chiga",
			language.CHR:     "Cherokee",
			language.CKB:     "Kurdi Tengah",
			language.CO:      "Korsika",
			language.CS:      "Ceska",
			language.CU:      "Slavia Gerejani",
			language.CY:      "Welsh",
			language.DA:      "Dansk",
			language.DAV:     "Taita",
			language.DE:      "Jérman",
			language.DJE:     "Zarma",
			language.DSB:     "Sorbia Non Standar",
			language.DUA:     "Duala",
			language.DYO:     "Jola-Fonyi",
			language.DZ:      "Dzongkha",
			language.EBU:     "Embu",
			language.EE:      "Ewe",
			language.EL:      "Yunani",
			language.EN:      "Inggris",
			language.EN_GB:   "Inggris (Britania)",
			language.EO:      "Esperanto",
			language.ES:      "Spanyol",
			language.ES_419:  "Spanyol (Amerika Latin)",
			language.ES_ES:   "Spanyol (Eropah)",
			language.ES_MX:   "Spanyol (Meksiko)",
			language.ET:      "Estonia",
			language.EU:      "Basque",
			language.EWO:     "Ewondo",
			language.FA:      "Persia",
			language.FF:      "Fulah",
			language.FI:      "Suomi",
			language.FIL:     "Tagalog",
			language.FO:      "Faroe",
			language.FR:      "Prancis",
			language.FUR:     "Friulian",
			language.FY:      "Frisia Sisih Kulon",
			language.GA:      "Irlandia",
			language.GD:      "Gaulia",
			language.GL:      "Galisia",
			language.GSW:     "Jerman Swiss",
			language.GU:      "Gujarat",
			language.GUZ:     "Gusii",
			language.GV:      "Manx",
			language.HA:      "Hausa",
			language.HAW:     "Hawaii",
			language.HE:      "Ibrani",
			language.HI:      "India",
			language.HMN:     "Hmong",
			language.HR:      "Kroasia",
			language.HSB:     "Sorbia Standar",
			language.HT:      "Kreol Haiti",
			language.HU:      "Hungaria",
			language.HY:      "Armenia",
			language.IA:      "Interlingua",
			language.ID:      "Indonesia",
			language.IG:      "Iqbo",
			language.II:      "Sichuan Yi",
			language.IS:      "Islandia",
			language.IT:      "Italia",
			language.JA:      "Jepang",
			language.JGO:     "Ngomba",
			language.JMC:     "Machame",
			language.JV:      "Jawa",
			language.KA:      "Georgia",
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
			language.KM:      "Khmer",
			language.KN:      "Kannada",
			language.KO:      "Korea",
			language.KOK:     "Konkani",
			language.KS:      "Kashmiri",
			language.KSB:     "Shambala",
			language.KSF:     "Bafia",
			language.KSH:     "Colonia",
			language.KU:      "Kurdis",
			language.KW:      "Kernowek",
			language.KY:      "Kirgis",
			language.LA:      "Latin",
			language.LAG:     "Langi",
			language.LB:      "Luksemburg",
			language.LG:      "Ganda",
			language.LKT:     "Lakota",
			language.LN:      "Lingala",
			language.LO:      "Laos",
			language.LRC:     "Luri Sisih Lor",
			language.LT:      "Lithuania",
			language.LU:      "Luba-Katanga",
			language.LUO:     "Luo",
			language.LUY:     "Luyia",
			language.LV:      "Latvia",
			language.MAS:     "Masai",
			language.MER:     "Meru",
			language.MFE:     "Morisyen",
			language.MG:      "Malagasi",
			language.MGH:     "Makhuwa-Meeto",
			language.MGO:     "Meta’",
			language.MI:      "Maori",
			language.MK:      "Makedonia",
			language.ML:      "Malayalam",
			language.MN:      "Mongolia",
			language.MR:      "Marathi",
			language.MS:      "Melayu",
			language.MT:      "Malta",
			language.MUA:     "Mundang",
			language.MUL:     "Basa Multilingua",
			language.MY:      "Myanmar",
			language.MZN:     "Mazanderani",
			language.NAQ:     "Nama",
			language.NB:      "Bokmål Norwegia",
			language.ND:      "Ndebele Lor",
			language.NDS:     "Jerman Non Standar",
			language.NE:      "Nepal",
			language.NL:      "Walanda",
			language.NL_BE:   "Flemis",
			language.NMG:     "Kwasio",
			language.NN:      "Nynorsk Norwegia",
			language.NNH:     "Ngiemboon",
			language.NUS:     "Nuer",
			language.NY:      "Nyanja",
			language.NYN:     "Nyankole",
			language.OM:      "Oromo",
			language.OR:      "Odia",
			language.OS:      "Ossetia",
			language.PA:      "Punjab",
			language.PL:      "Polandia",
			language.PRG:     "Prusia",
			language.PS:      "Pashto",
			language.PT:      "Portugis",
			language.QU:      "Quechua",
			language.RM:      "Roman",
			language.RN:      "Rundi",
			language.RO:      "Rumania",
			language.ROF:     "Rombo",
			language.RU:      "Rusia",
			language.RW:      "Kinyarwanda",
			language.RWK:     "Rwa",
			language.SA:      "Sanskerta",
			language.SAH:     "Sakha",
			language.SAQ:     "Samburu",
			language.SBP:     "Sangu",
			language.SD:      "Sindhi",
			language.SE:      "Sami Sisih Lor",
			language.SEH:     "Sena",
			language.SES:     "Koyraboro Senni",
			language.SG:      "Sango",
			language.SHI:     "Tachelhit",
			language.SI:      "Sinhala",
			language.SK:      "Slowakia",
			language.SL:      "Slovenia",
			language.SM:      "Samoa",
			language.SMN:     "Inari Sami",
			language.SN:      "Shona",
			language.SO:      "Somalia",
			language.SQ:      "Albania",
			language.SR:      "Serbia",
			language.ST:      "Sotho Sisih Kidul",
			language.SU:      "Sunda",
			language.SV:      "Swedia",
			language.SW:      "Swahili",
			language.TA:      "Tamil",
			language.TE:      "Telugu",
			language.TEO:     "Teso",
			language.TG:      "Tajik",
			language.TH:      "Thailand",
			language.TI:      "Tigrinya",
			language.TK:      "Turkmen",
			language.TO:      "Tonga",
			language.TR:      "Turki",
			language.TT:      "Tatar",
			language.TWQ:     "Tasawaq",
			language.TZM:     "Tamazight Atlas Tengah",
			language.UG:      "Uighur",
			language.UK:      "Ukraina",
			language.UND:     "Basa Ora Dikenali",
			language.UR:      "Urdu",
			language.UZ:      "Uzbek",
			language.VAI:     "Vai",
			language.VI:      "Vietnam",
			language.VO:      "Volapuk",
			language.VUN:     "Vunjo",
			language.WAE:     "Walser",
			language.WO:      "Wolof",
			language.XH:      "Xhosa",
			language.XOG:     "Soga",
			language.YAV:     "Yangben",
			language.YI:      "Yiddish",
			language.YO:      "Yoruba",
			language.YUE:     "Tyonghwa, Kanton",
			language.ZGH:     "Tamazight Moroko Standar",
			language.ZH:      "Tyonghwa, Mandarin",
			language.ZH_HANS: "Tyonghwa Mandarin (Ringkes)",
			language.ZH_HANT: "Tyonghwa Mandarin (Tradisional)",
			language.ZU:      "Zulu",
			language.ZXX:     "Konten tanpa linguistik",
		},
		Territories: cldr.Territories{
			territory.V_001: "Donya",
			territory.V_002: "Afrika",
			territory.V_003: "Amérika Lèr",
			territory.V_005: "Amérika Kidul",
			territory.V_009: "Oséania",
			territory.V_011: "Afrika Kulon",
			territory.V_013: "Amérika Tengah",
			territory.V_014: "Afrika Wétan",
			territory.V_015: "Afrika Lèr",
			territory.V_017: "Afrika Sisih Tengah",
			territory.V_018: "Afrika Sisih Kidul",
			territory.V_019: "Amérika",
			territory.V_021: "Amérika Sisih Lor",
			territory.V_029: "Karibia",
			territory.V_030: "Asia Wétan",
			territory.V_034: "Asia Kidul",
			territory.V_035: "Asia Kidul-wétan",
			territory.V_039: "Éropah Kidul",
			territory.V_053: "Australasia",
			territory.V_054: "Melanesia",
			territory.V_057: "Daerah Mikronesia",
			territory.V_061: "Polinesia",
			territory.V_142: "Asia",
			territory.V_143: "Asia Tengah",
			territory.V_145: "Asia Kulon",
			territory.V_150: "Éropah",
			territory.V_151: "Éropah Wétan",
			territory.V_154: "Éropah Lèr",
			territory.V_155: "Éropah Kulon",
			territory.V_202: "Afrika Kidule Sahara",
			territory.V_419: "Amérika Latin",
			territory.AC:    "Pulo Ascension",
			territory.AD:    "Andora",
			territory.AE:    "Uni Émirat Arab",
			territory.AF:    "Afganistan",
			territory.AG:    "Antigua lan Barbuda",
			territory.AI:    "Anguilla",
			territory.AL:    "Albani",
			territory.AM:    "Arménia",
			territory.AO:    "Angola",
			territory.AQ:    "Antartika",
			territory.AR:    "Argèntina",
			territory.AS:    "Samoa Amerika",
			territory.AT:    "Ostenrik",
			territory.AU:    "Ostrali",
			territory.AW:    "Aruba",
			territory.AX:    "Kapuloan Alan",
			territory.AZ:    "Azerbaijan",
			territory.BA:    "Bosnia lan Hèrségovina",
			territory.BB:    "Barbadhos",
			territory.BD:    "Banggaladésa",
			territory.BE:    "Bèlgi",
			territory.BF:    "Burkina Faso",
			territory.BG:    "Bulgari",
			territory.BH:    "Bahrain",
			territory.BI:    "Burundi",
			territory.BJ:    "Bénin",
			territory.BL:    "Saint Barthélémi",
			territory.BM:    "Bermuda",
			territory.BN:    "Brunéi",
			territory.BO:    "Bolivia",
			territory.BQ:    "Karibia Walanda",
			territory.BR:    "Brasil",
			territory.BS:    "Bahamas",
			territory.BT:    "Bhutan",
			territory.BV:    "Pulo Bovèt",
			territory.BW:    "Botswana",
			territory.BY:    "Bélarus",
			territory.BZ:    "Bélisé",
			territory.CA:    "Kanada",
			territory.CC:    "Kapuloan Cocos (Keeling)",
			territory.CD:    "Républik Dhémokratik Kongo",
			territory.CF:    "Républik Afrika Tengah",
			territory.CG:    "Républik Kongo",
			territory.CH:    "Switserlan",
			territory.CI:    "Pasisir Gadhing",
			territory.CK:    "Kapuloan Cook",
			territory.CL:    "Cilé",
			territory.CM:    "Kamerun",
			territory.CN:    "Tyongkok",
			territory.CO:    "Kolombia",
			territory.CP:    "Pulo Clipperton",
			territory.CR:    "Kosta Rika",
			territory.CU:    "Kuba",
			territory.CV:    "Pongol Verdé",
			territory.CW:    "Kurasao",
			territory.CX:    "Pulo Natal",
			territory.CY:    "Siprus",
			territory.CZ:    "Républik Céko",
			territory.DE:    "Jérman",
			territory.DG:    "Diego Garcia",
			territory.DJ:    "Jibuti",
			territory.DK:    "Dhènemarken",
			territory.DM:    "Dominika",
			territory.DO:    "Républik Dominika",
			territory.DZ:    "Aljasair",
			territory.EA:    "Séuta lan Melila",
			territory.EC:    "Ékuadhor",
			territory.EE:    "Éstonia",
			territory.EG:    "Mesir",
			territory.EH:    "Sahara Kulon",
			territory.ER:    "Éritréa",
			territory.ES:    "Sepanyol",
			territory.ET:    "Étiopia",
			territory.EU:    "Uni Éropah",
			territory.EZ:    "Zona Éuro",
			territory.FI:    "Finlan",
			territory.FJ:    "Fiji",
			territory.FK:    "Kapuloan Falkland (Islas Malvinas)",
			territory.FM:    "Féderasi Mikronésia",
			territory.FO:    "Kapuloan Faro",
			territory.FR:    "Prancis",
			territory.GA:    "Gabon",
			territory.GB:    "KM",
			territory.GD:    "Grénada",
			territory.GE:    "Géorgia",
			territory.GF:    "Guyana Prancis",
			territory.GG:    "Guernsei",
			territory.GH:    "Ghana",
			territory.GI:    "Gibraltar",
			territory.GL:    "Greenland",
			territory.GM:    "Gambia",
			territory.GN:    "Guinea",
			territory.GP:    "Guadélup",
			territory.GQ:    "Guinéa Katulistiwa",
			territory.GR:    "Grikenlan",
			territory.GS:    "Georgia Kidul lan Kapuloan Sandwich Kidul",
			territory.GT:    "Guatémala",
			territory.GU:    "Guam",
			territory.GW:    "Guinea-Bissau",
			territory.GY:    "Guyana",
			territory.HK:    "Hong Kong",
			territory.HM:    "Kapuloan Heard lan McDonald",
			territory.HN:    "Honduras",
			territory.HR:    "Kroasia",
			territory.HT:    "Haiti",
			territory.HU:    "Honggari",
			territory.IC:    "Kapuloan Kanari",
			territory.ID:    "Indonésia",
			territory.IE:    "Républik Irlan",
			territory.IL:    "Israèl",
			territory.IM:    "Pulo Man",
			territory.IN:    "Indhi",
			territory.IO:    "Wilayah Inggris nang Segoro Hindia",
			territory.IQ:    "Irak",
			territory.IR:    "Iran",
			territory.IS:    "Èslan",
			territory.IT:    "Itali",
			territory.JE:    "Jersey",
			territory.JM:    "Jamaika",
			territory.JO:    "Yordania",
			territory.JP:    "Jepang",
			territory.KE:    "Kénya",
			territory.KG:    "Kirgistan",
			territory.KH:    "Kamboja",
			territory.KI:    "Kiribati",
			territory.KM:    "Komoro",
			territory.KN:    "Saint Kits lan Nèvis",
			territory.KP:    "Koréa Lèr",
			territory.KR:    "Koréa Kidul",
			territory.KW:    "Kuwait",
			territory.KY:    "Kapuloan Kéman",
			territory.KZ:    "Kasakstan",
			territory.LA:    "Laos",
			territory.LB:    "Libanon",
			territory.LC:    "Santa Lusia",
			territory.LI:    "Liktenstén",
			territory.LK:    "Sri Lanka",
			territory.LR:    "Libèria",
			territory.LS:    "Lésotho",
			territory.LT:    "Litowen",
			territory.LU:    "Luksemburg",
			territory.LV:    "Latvia",
			territory.LY:    "Libya",
			territory.MA:    "Maroko",
			territory.MC:    "Monako",
			territory.MD:    "Moldova",
			territory.ME:    "Montenégro",
			territory.MF:    "Santa Martin",
			territory.MG:    "Madagaskar",
			territory.MH:    "Kapuloan Marshall",
			territory.MK:    "Makédonia",
			territory.ML:    "Mali",
			territory.MM:    "Myanmar (Burma)",
			territory.MN:    "Mongolia",
			territory.MO:    "Macau",
			territory.MP:    "Kapuloan Mariana Lor",
			territory.MQ:    "Martinik",
			territory.MR:    "Mauritania",
			territory.MS:    "Monsérat",
			territory.MT:    "Malta",
			territory.MU:    "Mauritius",
			territory.MV:    "Maladéwa",
			territory.MW:    "Malawi",
			territory.MX:    "Mèksiko",
			territory.MY:    "Malaysia",
			territory.MZ:    "Mosambik",
			territory.NA:    "Namibia",
			territory.NC:    "Kalédonia Anyar",
			territory.NE:    "Nigér",
			territory.NF:    "Pulo Norfolk",
			territory.NG:    "Nigéria",
			territory.NI:    "Nikaragua",
			territory.NL:    "Walanda",
			territory.NO:    "Nurwègen",
			territory.NP:    "Népal",
			territory.NR:    "Nauru",
			territory.NU:    "Niue",
			territory.NZ:    "Niu Sélan",
			territory.OM:    "Oman",
			territory.PA:    "Panama",
			territory.PE:    "Peru",
			territory.PF:    "Polinesia Prancis",
			territory.PG:    "Papua Nugini",
			territory.PH:    "Pilipina",
			territory.PK:    "Pakistan",
			territory.PL:    "Polen",
			territory.PM:    "Saint Pièr lan Mikuélon",
			territory.PN:    "Kapuloan Pitcairn",
			territory.PR:    "Puèrto Riko",
			territory.PS:    "Palèstina",
			territory.PT:    "Portugal",
			territory.PW:    "Palau",
			territory.PY:    "Paraguay",
			territory.QA:    "Katar",
			territory.QO:    "Oseania Paling Njaba",
			territory.RE:    "Réunion",
			territory.RO:    "Ruméni",
			territory.RS:    "Sèrbi",
			territory.RU:    "Rusia",
			territory.RW:    "Rwanda",
			territory.SA:    "Arab Saudi",
			territory.SB:    "Kapuloan Suleman",
			territory.SC:    "Sésèl",
			territory.SD:    "Sudan",
			territory.SE:    "Swèdhen",
			territory.SG:    "Singapura",
			territory.SH:    "Saint Héléna",
			territory.SI:    "Slovénia",
			territory.SJ:    "Svalbard lan Jan Mayen",
			territory.SK:    "Slowak",
			territory.SL:    "Siéra Léoné",
			territory.SM:    "San Marino",
			territory.SN:    "Sénégal",
			territory.SO:    "Somalia",
			territory.SR:    "Suriname",
			territory.SS:    "Sudan Kidul",
			territory.ST:    "Sao Tomé lan Principé",
			territory.SV:    "Èl Salvador",
			territory.SX:    "Sint Martén",
			territory.SY:    "Suriah",
			territory.SZ:    "(Swasiland)",
			territory.TA:    "Tristan da Cunha",
			territory.TC:    "Turks lan Kapuloan Kaikos",
			territory.TD:    "Chad",
			territory.TF:    "Wilayah Prancis nang Kutub Kidul",
			territory.TG:    "Togo",
			territory.TH:    "Tanah Thai",
			territory.TJ:    "Tajikistan",
			territory.TK:    "Tokelau",
			territory.TL:    "Timor Wétan",
			territory.TM:    "Turkménistan",
			territory.TN:    "Tunisia",
			territory.TO:    "Tonga",
			territory.TR:    "Turki",
			territory.TT:    "Trinidad lan Tobago",
			territory.TV:    "Tuvalu",
			territory.TW:    "Taiwan",
			territory.TZ:    "Tansania",
			territory.UA:    "Ukrania",
			territory.UG:    "Uganda",
			territory.UM:    "Kapuloan AS Paling Njaba",
			territory.UN:    "Pasarékatan Bangsa-Bangsa",
			territory.US:    "AS",
			territory.UY:    "Uruguay",
			territory.UZ:    "Usbèkistan",
			territory.VA:    "Kutha Vatikan",
			territory.VC:    "Saint Vinsen lan Grénadin",
			territory.VE:    "Vénésuéla",
			territory.VG:    "Kapuloan Virgin Britania",
			territory.VI:    "Kapuloan Virgin Amérika",
			territory.VN:    "Viètnam",
			territory.VU:    "Vanuatu",
			territory.WF:    "Wallis lan Futuna",
			territory.WS:    "Samoa",
			territory.XA:    "Logat Semu",
			territory.XB:    "Rong Arah Semu",
			territory.XK:    "Kosovo",
			territory.YE:    "Yaman",
			territory.YT:    "Mayotte",
			territory.ZA:    "Afrika Kidul",
			territory.ZM:    "Sambia",
			territory.ZW:    "Simbabwe",
			territory.ZZ:    "Daerah Ora Dikenali",
		},
	}
}
