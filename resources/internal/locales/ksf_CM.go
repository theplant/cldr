// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_ksf_CM() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "ksf_CM",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ŋ1", Feb: "ŋ2", Mar: "ŋ3", Apr: "ŋ4", May: "ŋ5", Jun: "ŋ6", Jul: "ŋ7", Aug: "ŋ8", Sep: "ŋ9", Oct: "ŋ10", Nov: "ŋ11", Dec: "ŋ12"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "ŋwíí a ntɔ́ntɔ", Feb: "ŋwíí akǝ bɛ́ɛ", Mar: "ŋwíí akǝ ráá", Apr: "ŋwíí akǝ nin", May: "ŋwíí akǝ táan", Jun: "ŋwíí akǝ táafɔk", Jul: "ŋwíí akǝ táabɛɛ", Aug: "ŋwíí akǝ táaraa", Sep: "ŋwíí akǝ táanin", Oct: "ŋwíí akǝ ntɛk", Nov: "ŋwíí akǝ ntɛk di bɔ́k", Dec: "ŋwíí akǝ ntɛk di bɛ́ɛ"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "sɔ́n", Mon: "lǝn", Tue: "maa", Wed: "mɛk", Thu: "jǝǝ", Fri: "júm", Sat: "sam"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "s", Mon: "l", Tue: "m", Wed: "m", Thu: "j", Fri: "j", Sat: "s"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "sɔ́ndǝ", Mon: "lǝndí", Tue: "maadí", Wed: "mɛkrɛdí", Thu: "jǝǝdí", Fri: "júmbá", Sat: "samdí"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "sárúwá", PM: "cɛɛ́nko"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "sárúwá", PM: "cɛɛ́nko"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "mɔni mǝ á bǝlɔŋ bǝ kaksa bɛ táatáaŋzǝn", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "mɔni mǝ á angóla", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "mɔni mǝ á ɔstralí", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "mɔni mǝ á barǝ́n", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "mɔni mǝ á burundí", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "mɔni mǝ á botswana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "mɔni mǝ á kanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "mɔni mǝ á kɔngó", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "mɔni mǝ á swís", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "mɔni mǝ á cín", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "mɔni mǝ á kapvɛr", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "mɔni mǝ á dyibutí", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "mɔni mǝ á aljɛrí", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "mɔni mǝ á ɛjípt", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "mɔni mǝ á ɛritrɛ́", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "mɔni mǝ á ɛtyɔpí", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "mɔni mǝ á pɛrɛsǝ́", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "mɔni mǝ á ingɛrís", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "mɔni mǝ á gána", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "mɔni mǝ á gambí", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "mɔni mǝ á ginɛ́", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "mɔni mǝ á indí", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "mɔni mǝ á japɔ́ŋ", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "mɔni mǝ á kɛnya", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "mɔni mǝ á komɔr", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "mɔni mǝ á libɛrya", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "mɔni mǝ á lǝsóto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "mɔni mǝ á libí", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "mɔni mǝ á marɔk", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "mɔni mǝ á madagaska", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "mɔni mǝ á mwaritaní (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "mɔni mǝ á mwaritaní", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "mɔni mǝ á mwarís", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "mɔni mǝ á malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "mɔni mǝ á mosambík", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "mɔni mǝ á namibí", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "mɔni mǝ á nijɛ́rya", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "mɔni mǝ á rwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "mɔni mǝ á arabí saodí", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "mɔni mǝ á sɛcɛl", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "mɔni mǝ á sudan", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "mɔni mǝ á sɛntɛ́len", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "mɔni mǝ á syɛraleon", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "mɔni mǝ á somalí", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "mɔni mǝ á saotomɛ́ ri priŋsib (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "mɔni mǝ á saotomɛ́ ri priŋsib", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "mɔni mǝ á swazilan", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "mɔni mǝ á tunɛsí", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "mɔni mǝ á tanzaní", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "mɔni mǝ á uganda", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "mɔni mǝ á amɛrika", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "fráŋ", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "mɔni mǝ á afríka aná wɛs", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "mɔni mǝ á afrik anǝ a sud", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "mɔni mǝ á zambí (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "mɔni mǝ á zambí", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "mɔni mǝ á zimbabwɛ́", Symbol: ""},
			},
		},
		Languages: cldr.Languages{
			language.AK:  "riakan",
			language.AM:  "riamarik",
			language.AR:  "riarab",
			language.BE:  "ribɛlɔrís",
			language.BG:  "ribulgarí",
			language.BN:  "ribɛngáli",
			language.CS:  "ricɛ́k",
			language.DE:  "ridjɛrman",
			language.EL:  "rigrɛ́k",
			language.EN:  "riingɛrís",
			language.ES:  "rikpanyá",
			language.FA:  "ripɛrsán",
			language.FR:  "ripɛrɛsǝ́",
			language.HA:  "rikaksa",
			language.HI:  "riíndí",
			language.HU:  "riɔngrɔá",
			language.ID:  "riindonɛsí",
			language.IG:  "riigbo",
			language.IT:  "riitalyɛ́n",
			language.JA:  "rijapɔ́ŋ",
			language.JV:  "rijawanɛ́",
			language.KM:  "rikmɛr",
			language.KO:  "rikɔrɛɛ́",
			language.KSF: "rikpa",
			language.MS:  "rimalaí",
			language.MY:  "ribirmán",
			language.NE:  "rinepalɛ́",
			language.NL:  "riɔlándɛ́",
			language.PA:  "ripɛnjabí",
			language.PL:  "ripɔlɔ́n",
			language.PT:  "ripɔrtugɛ́",
			language.RO:  "rirɔmán",
			language.RU:  "rirís",
			language.RW:  "rirwanda",
			language.SO:  "risomalí",
			language.SV:  "riswɛ́dǝ",
			language.TA:  "ritamúl",
			language.TH:  "ritaí",
			language.TR:  "riturk",
			language.UK:  "riukrɛ́n",
			language.UR:  "riurdú",
			language.VI:  "riwyɛtnám",
			language.YO:  "riyúuba",
			language.ZH:  "ricinɔá",
			language.ZU:  "rizúlu",
		},
		Territories: cldr.Territories{
			territory.AD: "andɔrǝ",
			territory.AE: "bǝlɔŋ bǝ kaksa bɛ táatáaŋzǝn",
			territory.AF: "afganistáŋ",
			territory.AG: "antiga ri barbúda",
			territory.AI: "angiya",
			territory.AL: "albaní",
			territory.AM: "armɛní",
			territory.AO: "angóla",
			territory.AR: "arjǝntín",
			territory.AS: "samɔa a amɛrika",
			territory.AT: "otric",
			territory.AU: "ɔstralí",
			territory.AW: "aruba",
			territory.AZ: "azabecán",
			territory.BA: "bɔsnyɛ ri hɛrsǝgɔvín",
			territory.BB: "baabaadǝ",
			territory.BD: "baŋladɛ́c",
			territory.BE: "bɛljík",
			territory.BF: "bukína fǝ́ asɔ",
			territory.BG: "bulgarí",
			territory.BH: "barǝ́n",
			territory.BI: "burundí",
			territory.BJ: "bɛnǝ́n",
			territory.BM: "bɛɛmúdǝ",
			territory.BN: "brunǝ́",
			territory.BO: "bɔɔlíví",
			territory.BR: "brɛsíl",
			territory.BS: "baamás",
			territory.BT: "bután",
			territory.BW: "botswana",
			territory.BY: "bɛlaris",
			territory.BZ: "bɛliz",
			territory.CA: "kanada",
			territory.CD: "kɔngó anyɔ́n",
			territory.CF: "santrafrík",
			territory.CG: "kɔngó",
			territory.CH: "swís",
			territory.CI: "kɔtiwuár",
			territory.CK: "zɛ i kúk",
			territory.CL: "cíli",
			territory.CM: "kamɛrún",
			territory.CN: "cín",
			territory.CO: "kolɔmbí",
			territory.CR: "kɔstaríka",
			territory.CU: "kuba",
			territory.CV: "kapvɛr",
			territory.CY: "cíprɛ",
			territory.CZ: "cɛ́k",
			territory.DE: "djɛrman",
			territory.DJ: "dyibutí",
			territory.DK: "danmak",
			territory.DM: "dɔminik",
			territory.DO: "dɔminik rɛpublík",
			territory.DZ: "aljɛrí",
			territory.EC: "ɛkwatɛǝ́",
			territory.EE: "ɛstoní",
			territory.EG: "ɛjípt",
			territory.ER: "ɛritrɛ́",
			territory.ES: "kpanyá",
			territory.ET: "ɛtyɔpí",
			territory.FI: "fínlan",
			territory.FJ: "fíji",
			territory.FK: "zǝ maalwín",
			territory.FM: "mikronɛ́si",
			territory.FR: "pɛrɛsǝ́",
			territory.GA: "gabɔŋ",
			territory.GB: "kǝlɔŋ kǝ kǝtáatáaŋzǝn",
			territory.GD: "grɛnadǝ",
			territory.GE: "jɔrjí",
			territory.GF: "guyán i pɛrɛsǝ́",
			territory.GH: "gána",
			territory.GI: "jibraltá",
			territory.GL: "grínlan",
			territory.GM: "gambí",
			territory.GN: "ginɛ́",
			territory.GP: "gwadɛlúp",
			territory.GQ: "ginɛ́ ɛkwatɔrial",
			territory.GR: "grɛ́k",
			territory.GT: "gwátǝmala",
			territory.GU: "gwám",
			territory.GW: "ginɛ́ bisɔ́",
			territory.GY: "guyán",
			territory.HN: "ɔnduras",
			territory.HR: "krwasí",
			territory.HT: "ayiti",
			territory.HU: "ɔngrí",
			territory.ID: "indonɛsí",
			territory.IE: "ilán",
			territory.IL: "israɛ́l",
			territory.IN: "indí",
			territory.IO: "zǝ ingɛrís ncɔ́m wa indi",
			territory.IQ: "irák",
			territory.IR: "iráŋ",
			territory.IS: "zǝ i glás",
			territory.IT: "italí",
			territory.JM: "jamaík",
			territory.JO: "jɔrdán",
			territory.JP: "japɔ́ŋ",
			territory.KE: "kɛnya",
			territory.KG: "kigistáŋ",
			territory.KH: "kambodj",
			territory.KI: "kiribáti",
			territory.KM: "komɔr",
			territory.KN: "sɛnkrǝstɔ́f ri nyɛ́vǝ",
			territory.KP: "korɛanɔ́r",
			territory.KR: "korɛasud",
			territory.KW: "kuwɛit",
			territory.KY: "zǝ i gan",
			territory.KZ: "kazakstáŋ",
			territory.LA: "laɔs",
			territory.LB: "libáŋ",
			territory.LC: "sɛntlísí",
			territory.LI: "lictɛnstɛ́n",
			territory.LK: "srílaŋka",
			territory.LR: "libɛrya",
			territory.LS: "lǝsóto",
			territory.LT: "litwaní",
			territory.LU: "luksɛmbúr",
			territory.LV: "lɛtoní",
			territory.LY: "libí",
			territory.MA: "marɔk",
			territory.MC: "monako",
			territory.MD: "mɔldaví",
			territory.MG: "madagaska",
			territory.MH: "zǝ i marcál",
			territory.ML: "mali",
			territory.MM: "myanmár",
			territory.MN: "mɔŋolí",
			territory.MP: "zǝ maryánnɔ́r",
			territory.MQ: "matiník",
			territory.MR: "mwaritaní",
			territory.MS: "mɔnsɛrat",
			territory.MT: "maltǝ",
			territory.MU: "mwarís",
			territory.MV: "maldivǝ",
			territory.MW: "malawi",
			territory.MX: "mɛksík",
			territory.MY: "malɛsí",
			territory.MZ: "mosambík",
			territory.NA: "namibí",
			territory.NC: "kalɛdoní anyɔ́n",
			territory.NE: "nijɛ́r",
			territory.NF: "zɛ nɔ́fɔlk",
			territory.NG: "nijɛ́rya",
			territory.NI: "níkarágwa",
			territory.NL: "kǝlɔŋ kǝ ázǝ",
			territory.NO: "nɔrvɛjǝ",
			territory.NP: "nɛpal",
			territory.NR: "nwarú",
			territory.NU: "niwɛ́",
			territory.NZ: "zɛlan anyɔ́n",
			territory.OM: "oman",
			territory.PA: "panama",
			territory.PE: "pɛrú",
			territory.PF: "pɔlinɛsí a pɛrɛsǝ́",
			territory.PG: "papwazí ginɛ́ anyɔ́n",
			territory.PH: "filipǝ́n",
			territory.PK: "pakistáŋ",
			territory.PL: "polɔ́n",
			territory.PM: "sɛnpyɛr ri mikɛlɔŋ",
			territory.PN: "pitkɛ́n",
			territory.PR: "pɔtoríko",
			territory.PS: "zǝ palɛstínǝ",
			territory.PT: "portugál",
			territory.PW: "palwa",
			territory.PY: "paragwɛ́",
			territory.QA: "katá",
			territory.RE: "rɛunyɔŋ",
			territory.RO: "rɔmaní",
			territory.RU: "risí",
			territory.RW: "rwanda",
			territory.SA: "arabí saodí",
			territory.SB: "zǝ salomɔ́n",
			territory.SC: "sɛcɛl",
			territory.SD: "sudan",
			territory.SE: "swɛdǝ",
			territory.SG: "siŋapó",
			territory.SH: "sɛntɛ́len",
			territory.SI: "slovɛní",
			territory.SK: "slovakí",
			territory.SL: "syɛraleon",
			territory.SM: "sɛnmarǝn",
			territory.SN: "sɛnɛgal",
			territory.SO: "somalí",
			territory.SR: "surinam",
			territory.ST: "saotomɛ́ ri priŋsib",
			territory.SV: "salvadɔr",
			territory.SY: "sirí",
			territory.SZ: "swazilan",
			territory.TC: "zǝ tirk ri kakɔs",
			territory.TD: "caád",
			territory.TG: "togo",
			territory.TH: "tɛlan",
			territory.TJ: "tadjikistaŋ",
			territory.TK: "tokǝlao",
			territory.TL: "timor anǝ á ɛst",
			territory.TM: "tirkmɛnistaŋ",
			territory.TN: "tunɛsí",
			territory.TO: "tɔŋa",
			territory.TR: "tirkí",
			territory.TT: "tɛrinitɛ ri tobago",
			territory.TV: "tuwalu",
			territory.TW: "tɛwán",
			territory.TZ: "tanzaní",
			territory.UA: "ukrain",
			territory.UG: "uganda",
			territory.US: "amɛrika",
			territory.UY: "urugwɛ́",
			territory.UZ: "usbɛkistaŋ",
			territory.VA: "watikáŋ",
			territory.VC: "sɛnvǝnsǝŋ ri grɛnadín",
			territory.VE: "wɛnǝzwɛla",
			territory.VG: "zǝ bɛ gɔn inɛ a ingɛrís",
			territory.VI: "zǝ bɛ gɔn inɛ á amɛrika",
			territory.VN: "wyɛtnám",
			territory.VU: "wanwatu",
			territory.WF: "walis ri futuna",
			territory.WS: "samɔa",
			territory.YE: "yɛmɛn",
			territory.YT: "mayɔ́t",
			territory.ZA: "afrik anǝ a sud",
			territory.ZM: "zambí",
			territory.ZW: "zimbabwɛ́",
		},
	}
}
