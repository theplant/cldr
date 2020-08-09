// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_nn_NO() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "nn_NO",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM y", Short: "dd.MM.y"}, Time: cldr.CalendarDateFormat{Full: "'kl'. HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} 'kl'. {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan", Feb: "feb", Mar: "mar", Apr: "apr", May: "mai", Jun: "jun", Jul: "jul", Aug: "aug", Sep: "sep", Oct: "okt", Nov: "nov", Dec: "des"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "januar", Feb: "februar", Mar: "mars", Apr: "april", May: "mai", Jun: "juni", Jul: "juli", Aug: "august", Sep: "september", Oct: "oktober", Nov: "november", Dec: "desember"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "søn", Mon: "mån", Tue: "tys", Wed: "ons", Thu: "tor", Fri: "fre", Sat: "lau"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "O", Thu: "T", Fri: "F", Sat: "L"}, Short: cldr.CalendarDayFormatNameValue{Sun: "sø.", Mon: "må.", Tue: "ty.", Wed: "on.", Thu: "to.", Fri: "fr.", Sat: "la."}, Wide: cldr.CalendarDayFormatNameValue{Sun: "søndag", Mon: "måndag", Tue: "tysdag", Wed: "onsdag", Thu: "torsdag", Fri: "fredag", Sat: "laurdag"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "f.m.", PM: "e.m."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "f.m.", PM: "e.m."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "f.m.", PM: "e.m."}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_nn]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "−", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "andorranske peseta", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "emiratarabiske dirham", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "afghani (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "afghanske afghani", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "albanske lek", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "armenske dram", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "nederlandske antillegylden", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "angolanske kwanza", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "angolske kwanza (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "angolske nye kwanza (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "angolske kwanza reajustado (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "argentiske austral", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "argentinske peso (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "argentinske pesos", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "austerrikske schilling", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "australske dollar", Symbol: "AUD"},
				currency.AWG: cldr.Currency{DisplayName: "arubiske floriner", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "aserbaijanske manat", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "aserbajdsjanske manat", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "bosnisk-hercegovinske dinarar", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "bosnisk-hercegovinske konvertible mark", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "barbadiske dollar", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "bangladeshiske taka", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "belgiske franc (konvertibel)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "belgiske franc", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "belgiske franc (finansiell)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "bulgarsk hard lev", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "bulgarske lev", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "bahrainske dinarar", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "burundiske franc", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "bermudiske dollar", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "bruneiske dollar", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "bolivianske boliviano", Symbol: "BOB"},
				currency.BOP: cldr.Currency{DisplayName: "boliviske peso", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "boliviske mvdol", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "brasiliansk cruzeiro novo (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "brasilianske cruzado", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "brasilianske cruzeiro (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "brasilianske real", Symbol: "BRL"},
				currency.BRN: cldr.Currency{DisplayName: "brasilianske cruzado novo", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "brasilianske cruzeiro", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "bahamanske dollar", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "bhutanske ngultrum", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "burmesisk kyat", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "botswanske pula", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "kviterussiske nye rublar (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "kviterussiske rublar", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "kviterussiske rublar (2000–2016)", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "beliziske dollar", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "kanadiske dollar", Symbol: "CAD"},
				currency.CDF: cldr.Currency{DisplayName: "kongolesiske franc", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "WIR euro", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "sveitsiske franc", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "WIR franc", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "chilenske unidades de fomento", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "chilenske pesos", Symbol: "CLP"},
				currency.CNY: cldr.Currency{DisplayName: "kinesiske yuan", Symbol: "CNY"},
				currency.COP: cldr.Currency{DisplayName: "kolombianske pesos", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "unidad de valor real", Symbol: ""},
				currency.CRC: cldr.Currency{DisplayName: "kostarikanske colón", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "gamle serbiske dinarer", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "tsjekkoslovakiske koruna (hard)", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "kubanske konvertible pesos", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "kubanske pesos", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "kappverdiske escudo", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "kypriotiske pund", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "tsjekkiske koruna", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "austtyske mark", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "tyske mark", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "djiboutiske franc", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "danske kroner", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "dominikanske pesos", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "algeriske dinarar", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "ecuadorianske sucre", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "ecuadorianske unidad de valor constante (UVC)", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "estiske kroon", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "egyptiske pund", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "eritreiske nakfa", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "spanske peseta (A–konto)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "spanske peseta (konvertibel konto)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "spanske peseta", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "etiopiske birr", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "finske mark", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "fijianske dollar", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "falklandspund", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "franske franc", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "britiske pund", Symbol: "GBP"},
				currency.GEK: cldr.Currency{DisplayName: "georgiske kupon larit", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "georgiske lari", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "ghanesiske cedi (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "ghanesiske cedi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "gibraltarske pund", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "gambiske dalasi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "guineanske franc", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "guineanske syli", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "ekvatorialguineanske ekwele guineana", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "greske drakme", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "guatemalanske quetzal", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "portugisiske guinea escudo", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "Guinea-Bissau-peso", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "guyanske dollar", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Hongkong-dollar", Symbol: "HKD"},
				currency.HNL: cldr.Currency{DisplayName: "honduranske lempira", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "kroatiske dinar", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "kroatiske kuna", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "haitiske gourde", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "ungarske forintar", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "indonesiske rupiar", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "irske pund", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "israelske pund", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "israelske nye sheklar", Symbol: "ILS"},
				currency.INR: cldr.Currency{DisplayName: "indiske rupiar", Symbol: "INR"},
				currency.IQD: cldr.Currency{DisplayName: "irakiske dinarar", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "iranske rial", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "islandske kroner", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "italienske lire", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "jamaikanske dollar", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "jordanske dinarar", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "japanske yen", Symbol: "JPY"},
				currency.KES: cldr.Currency{DisplayName: "kenyanske shilling", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "kirgisiske som", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "kambodsjanske riel", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "komoriske franc", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "nordkoreanske won", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "sørkoreanske won", Symbol: "KRW"},
				currency.KWD: cldr.Currency{DisplayName: "kuwaitiske dinarar", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "caymanske dollar", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "kasakhstanske tenge", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "laotiske kip", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "libanesiske pund", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "srilankiske rupiar", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "liberiske dollar", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "lesothiske loti", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "litauiske lita", Symbol: "Lt"},
				currency.LTT: cldr.Currency{DisplayName: "litauiske talona", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "luxemburgske konvertibel franc", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "luxemburgske franc", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "luxemburgske finansielle franc", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "latviske lat", Symbol: "Ls"},
				currency.LVR: cldr.Currency{DisplayName: "latviske rublar", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "libyske dinarar", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "marokkanske dirham", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "marokkanske franc", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "moldovske leuar", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "madagassiske ariary", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "madagassiske franc", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "makedonske denarar", Symbol: "MKD"},
				currency.MLF: cldr.Currency{DisplayName: "maliske franc", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "myanmarske kyat", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "mongolske tugrik", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "makaoiske pataca", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "mauritanske ouguiya (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "mauritanske ouguiya", Symbol: ""},
				currency.MTL: cldr.Currency{DisplayName: "maltesiske lira", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "maltesiske pund", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "mauritanske rupiar", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "maldiviske rufiyaa", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "malawiske kwacha", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "meksikanske pesos", Symbol: "MXN"},
				currency.MXP: cldr.Currency{DisplayName: "meksikanske sølvpeso (1861–1992)", Symbol: ""},
				currency.MXV: cldr.Currency{DisplayName: "meksikanske unidad de inversion (UDI)", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "malaysiske ringgit", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "mosambikiske escudo", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "gamle mosambikiske metical", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "mosambikiske metical", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "namibiske dollar", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "nigerianske naira", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "nicaraguanske cordoba", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "nicaraguanske córdoba", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "nederlandske gylden", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "norske kroner", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "nepalske rupiar", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "nyzealandske dollar", Symbol: "NZD"},
				currency.OMR: cldr.Currency{DisplayName: "omanske rial", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "panamanske balboa", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "peruanske inti", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "peruanske sol", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "peruanske sol (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "papuanske kina", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "filippinske pesos", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "pakistanske rupiar", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "polske zloty", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "polske zloty (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "portugisiske escudo", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "paraguayanske guaraní", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "qatarske rial", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "rhodesiske dollar", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "gamle rumenske leu", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "rumenske leuar", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "serbiske dinarar", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "russiske rublar", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "russiske rublar (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "rwandiske franc", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "saudiarabiske rial", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "salomonske dollar", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "seychelliske rupiar", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "gamle sudanske dinarer", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "sudanske pund", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "gamle sudanske pund", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "svenske kroner", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "singaporske dollar", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "sankthelenske pund", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "slovenske tolar", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "slovakiske koruna", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "sierraleonske leone", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "somaliske shilling", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "surinamske dollar", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "surinamske gylden", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "sørsudanske pund", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "saotomesiske dobra (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "saotomesiske dobra", Symbol: "Db"},
				currency.SUR: cldr.Currency{DisplayName: "sovjetiske rublar", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "salvadoranske colon", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "syriske pund", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "swazilandske lilangeni", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "thailandske baht", Symbol: "THB"},
				currency.TJR: cldr.Currency{DisplayName: "tadsjikiske rublar", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "tadsjikiske somoni", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "turkmensk manat (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "turkmenske manat", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "tunisiske dinarar", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "tonganske paʻanga", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "timoresiske escudo", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "gamle tyrkiske lire", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "tyrkiske lire", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "trinidadiske dollar", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "nye taiwanske dollar", Symbol: "TWD"},
				currency.TZS: cldr.Currency{DisplayName: "tanzanianske shilling", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "ukrainske hryvnia", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "ukrainske karbovanetz", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "ugandiske shilling (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "ugandiske shilling", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "amerikanske dollar", Symbol: "USD"},
				currency.USN: cldr.Currency{DisplayName: "amerikanske dollar (neste dag)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "amerikanske dollar (same dag)", Symbol: ""},
				currency.UYI: cldr.Currency{DisplayName: "uruguayanske peso en unidades indexadas", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "uruguayanske peso (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "uruguayanske pesos", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "usbekiske sum", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "venezuelanske bolivar (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "venezuelanske bolivar (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "venezuelanske bolivar", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "vietnamesiske dong", Symbol: "VND"},
				currency.VUV: cldr.Currency{DisplayName: "vanuatuiske vatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "samoanske tala", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "sentralafrikanske CFA-franc", Symbol: "XAF"},
				currency.XAG: cldr.Currency{DisplayName: "sølv", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "gull", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "europeiske samansette einingar", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "europeiske monetære einingar", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "europeiske kontoeiningar (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "europeiske kontoeiningar (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "austkaribiske dollar", Symbol: "XCD"},
				currency.XDR: cldr.Currency{DisplayName: "spesielle trekkrettar", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "europeiske valutaeiningar", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "franske gullfranc", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "franske UIC-franc", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "vestafrikanske CFA-franc", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "palladium", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "CFP-franc", Symbol: "XPF"},
				currency.XPT: cldr.Currency{DisplayName: "platina", Symbol: ""},
				currency.XRE: cldr.Currency{DisplayName: "RINET-fond", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "testvalutakode", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "ukjend valuta", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "jemenittiske dinarar", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "jemenittiske rial", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "jugoslaviske dinarar (hard)", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "jugoslaviske noviy-dinarar", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "jugoslaviske konvertibel dinarar", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "sørafrikanske rand (finansiell)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "sørafrikanske rand", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "zambiske kwacha (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "zambiske kwacha", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "zairisk ny zaire", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "zairisk zaire", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "zimbabwisk dollar", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "Zimbabwe-dollar (2009)", Symbol: ""},
			},
		},
	}
}
