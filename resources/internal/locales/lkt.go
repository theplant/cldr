// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_lkt() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "lkt",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Wiótheȟika Wí", Feb: "Thiyóȟeyuŋka Wí", Mar: "Ištáwičhayazaŋ Wí", Apr: "Pȟežítȟo Wí", May: "Čhaŋwápetȟo Wí", Jun: "Wípazukȟa-wašté Wí", Jul: "Čhaŋpȟásapa Wí", Aug: "Wasútȟuŋ Wí", Sep: "Čhaŋwápeǧi Wí", Oct: "Čhaŋwápe-kasná Wí", Nov: "Waníyetu Wí", Dec: "Tȟahékapšuŋ Wí"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "A", Mon: "W", Tue: "N", Wed: "Y", Thu: "T", Fri: "Z", Sat: "O"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Aŋpétuwakȟaŋ", Mon: "Aŋpétuwaŋži", Tue: "Aŋpétunuŋpa", Wed: "Aŋpétuyamni", Thu: "Aŋpétutopa", Fri: "Aŋpétuzaptaŋ", Sat: "Owáŋgyužažapi"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "", Currency: "¤\u00a00K", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AOA: cldr.Currency{DisplayName: "", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "", Symbol: "CA$"},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.EGP: cldr.Currency{DisplayName: "", Symbol: "E£"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.EUR: cldr.Currency{DisplayName: "", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "", Symbol: "JP¥"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.MGA: cldr.Currency{DisplayName: "", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MUR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.NAD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "", Symbol: "RF"},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.USD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Languages: cldr.Languages{
			language.AB:      "Abkhaz Iyápi",
			language.ADY:     "Adyghe Iyápi",
			language.AE:      "Avestan Iyápi",
			language.AF:      "Afrikaans Iyápi",
			language.ALT:     "Itóǧata Altai Iyápi",
			language.AM:      "Amharic Iyápi",
			language.AR:      "Arab Iyápi",
			language.ARP:     "Maȟpíya Tȟó Iyápi",
			language.AS:      "Assamese Iyápi",
			language.AV:      "Avaric Iyápi",
			language.AZ:      "Azerbaijani Iyápi",
			language.BA:      "Bashkir Iyápi",
			language.BAL:     "Baluchi Iyápi",
			language.BAX:     "Bamun Iyápi",
			language.BE:      "Belarus Iyápi",
			language.BEJ:     "Beja Iyápi",
			language.BG:      "Bulgar Iyápi",
			language.BN:      "Bengali Iyápi",
			language.BO:      "Tibetan Iyápi",
			language.BS:      "Bosnia Iyápi",
			language.BUA:     "Buriat Iyápi",
			language.CA:      "Catalan Iyápi",
			language.CE:      "Chechen Iyápi",
			language.CHM:     "Mari Iyápi",
			language.CHR:     "Cherokee Iyápi",
			language.CHY:     "Šahíyela Iyápi",
			language.COP:     "Coptic Iyápi",
			language.CR:      "Maštíŋča Oyáte Iyápi",
			language.CRH:     "Crimean Turkish Iyápi",
			language.CS:      "Czech Iyápi",
			language.CV:      "Chuvash Iyápi",
			language.CY:      "Wales Iyápi",
			language.DA:      "Dane Iyápi",
			language.DAK:     "Dakȟótiyapi",
			language.DAR:     "Dargwa Iyápi",
			language.DE:      "Iyášiča Iyápi",
			language.DOI:     "Dogri Iyápi",
			language.EL:      "Greece Iyápi",
			language.EN:      "Wašíčuiyapi",
			language.EN_GB:   "Šagláša Wašíčuiyapi",
			language.EN_US:   "Mílahaŋska Wašíčuiyapi",
			language.EO:      "Esperanto Iyápi",
			language.ES:      "Spayóla Iyápi",
			language.ES_419:  "Wiyóȟpeyata Spayóla Iyápi",
			language.ES_ES:   "Spayólaȟča Iyápi",
			language.ET:      "Estonia Iyápi",
			language.EU:      "Basque Iyápi",
			language.FA:      "Persian Iyápi",
			language.FI:      "Finnish Iyápi",
			language.FIL:     "Filipino Iyápi",
			language.FJ:      "Fiji Iyápi",
			language.FO:      "Faroese Iyápi",
			language.FR:      "Wašíču Ikčéka Iyápi",
			language.GA:      "Irish Iyápi",
			language.GBA:     "Gbaya Iyápi",
			language.GL:      "Galician Iyápi",
			language.GN:      "Guarani Iyápi",
			language.GU:      "Gujarati Iyápi",
			language.HA:      "Hausa Iyápi",
			language.HAW:     "Hawaiian Iyápi",
			language.HE:      "Hebrew Iyápi",
			language.HI:      "Hindi Iyápi",
			language.HR:      "Croatian Iyápi",
			language.HT:      "Haiti Iyápi",
			language.HU:      "Hungary Iyápi",
			language.HY:      "Armenia Iyápi",
			language.ID:      "Indonesia Iyápi",
			language.IG:      "Igbo Iyápi",
			language.INH:     "Ingush Iyápi",
			language.IS:      "Iceland Iyápi",
			language.IT:      "Italia Iyápi",
			language.JA:      "Kisúŋla Iyápi",
			language.JV:      "Java Iyápi",
			language.KA:      "Georia Iyápi",
			language.KAA:     "Kara-Kalpak Iyápi",
			language.KBD:     "Kabardian Iyápi",
			language.KK:      "Kazakh Iyápi",
			language.KM:      "Khmer Iyápi",
			language.KN:      "Kannada Iyápi",
			language.KO:      "Korea Iyápi",
			language.KS:      "Kashmir Iyápi",
			language.KU:      "Kurd Iyápi",
			language.KY:      "Kirghiz Iyápi",
			language.LA:      "Latin Iyápi",
			language.LAH:     "Lahnda Iyápi",
			language.LB:      "Luxembourg Iyápi",
			language.LKT:     "Lakȟólʼiyapi",
			language.LO:      "Lao Iyápi",
			language.LT:      "Lithuania Iyápilt",
			language.LUS:     "Mizo Iyápi",
			language.LV:      "Latvia Iyápi",
			language.MG:      "Malagasy Iyápi",
			language.MI:      "Maori Iyápi",
			language.MK:      "Macedonia Iyápi",
			language.ML:      "Malayalam Iyápi",
			language.MNI:     "Namipuri Iyápi",
			language.MR:      "Marathi Iyápi",
			language.MS:      "Malay Iyápi",
			language.MT:      "Maltese Iyápi",
			language.MY:      "Burmese Iyápi",
			language.NE:      "Nepal Iyápi",
			language.NL:      "Dutch Iyápi",
			language.NL_BE:   "Flemish Iyápi",
			language.NV:      "Šináglegleǧa Iyápi",
			language.OJ:      "Ȟaȟátȟuŋwaŋ Iyápi",
			language.OR:      "Oriya Iyápi",
			language.PA:      "Punjabi Iyápi",
			language.PL:      "Polish Iyápi",
			language.PS:      "Pashto Iyápi",
			language.PT:      "Portuguese Iyápi",
			language.QU:      "Quechua Iyápi",
			language.RM:      "Romansh Iyápi",
			language.RO:      "Romanian Iyápi",
			language.RU:      "Russia Iyápi",
			language.SA:      "Sanskrit Iyápi",
			language.SD:      "Sindhi Iyápi",
			language.SI:      "Sinhala Iyápi",
			language.SK:      "Slovak Iyápi",
			language.SL:      "Slovenian Iyápi",
			language.SO:      "Somali Iyápi",
			language.SQ:      "Albanian Iyápi",
			language.SR:      "Serbia Iyápi",
			language.SU:      "Sundanese Iyápi",
			language.SV:      "Swedish Iyápi",
			language.SW:      "Swahili Iyápi",
			language.SWB:     "Comonian Iyápi",
			language.TA:      "Tamil Iyápi",
			language.TE:      "Telugu Iyápi",
			language.TG:      "Tajik Iyápi",
			language.TH:      "Thai Iyápi",
			language.TI:      "Tigrinya Iyápi",
			language.TK:      "Turkmen Iyápi",
			language.TO:      "Tongan Iyápi",
			language.TR:      "Turkish Iyápi",
			language.TT:      "Tatar Iyápi",
			language.UG:      "Uyghur Iyápi",
			language.UK:      "Ukrain Iyápi",
			language.UND:     "Tukté iyápi tȟaŋíŋ šni",
			language.UR:      "Urdu Iyápi",
			language.UZ:      "Uzbek Iyápi",
			language.VI:      "Vietnamese Iyápi",
			language.WO:      "Wolof Iyápi",
			language.XH:      "Xhosa Iyápi",
			language.YO:      "Yoruba Iyápi",
			language.ZH:      "Pȟečhókaŋ Háŋska Iyápi",
			language.ZH_HANS: "Pȟečhókaŋ Háŋska Iyápi Ikčéka",
			language.ZH_HANT: "Pȟečhókaŋ Háŋska Iyápi Ȟče",
			language.ZU:      "Zulu Iyápi",
			language.ZZA:     "Zaza Iyápi",
		},
		Territories: cldr.Territories{
			territory.V_001: "Makȟásitomni",
			territory.V_002: "Hásapa Makȟáwita",
			territory.V_019: "Khéya Wíta",
			territory.V_142: "Hazíla Makȟáwita",
			territory.V_150: "Wašíču Makȟáwita",
			territory.CA:    "Uŋčíyapi Makȟóčhe",
			territory.CN:    "Pȟečhókaŋhaŋska Makȟóčhe",
			territory.DE:    "Iyášiča Makȟóčhe",
			territory.ES:    "Spayólaȟče Makȟóčhe",
			territory.JP:    "Kisúŋla Makȟóčhe",
			territory.MX:    "Spayóla Makȟóčhe",
			territory.US:    "Mílahaŋska Tȟamákȟočhe",
		},
	}
}
