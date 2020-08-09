// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_sv_AX() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "sv_AX",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "y-MM-dd"}, Time: cldr.CalendarDateFormat{Full: "'kl'. HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan.", Feb: "feb.", Mar: "mars", Apr: "apr.", May: "maj", Jun: "juni", Jul: "juli", Aug: "aug.", Sep: "sep.", Oct: "okt.", Nov: "nov.", Dec: "dec."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "januari", Feb: "februari", Mar: "mars", Apr: "april", May: "maj", Jun: "juni", Jul: "juli", Aug: "augusti", Sep: "september", Oct: "oktober", Nov: "november", Dec: "december"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "sön", Mon: "mån", Tue: "tis", Wed: "ons", Thu: "tors", Fri: "fre", Sat: "lör"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "O", Thu: "T", Fri: "F", Sat: "L"}, Short: cldr.CalendarDayFormatNameValue{Sun: "sö", Mon: "må", Tue: "ti", Wed: "on", Thu: "to", Fri: "fr", Sat: "lö"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "söndag", Mon: "måndag", Tue: "tisdag", Wed: "onsdag", Thu: "torsdag", Fri: "fredag", Sat: "lördag"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "f.m.", PM: "e.m."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "fm", PM: "em"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "förmiddag", PM: "eftermiddag"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_sv]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "٫", Group: "\u00a0", Negative: "\u061c−", Percent: "٪\u061c", PerMille: "؉\u200f"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "andorransk peseta", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "Förenade Arabemiratens dirham", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "afghani (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "afghansk afghani", Symbol: "AFN"},
				currency.ALK: cldr.Currency{DisplayName: "albansk lek (1946–1965)", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "albansk lek", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "armenisk dram", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Nederländska Antillernas gulden", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "angolansk kwanza", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "angolansk kwanza (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "angolansk ny kwanza (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "angolansk kwanza reajustado (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "argentinsk austral", Symbol: ""},
				currency.ARL: cldr.Currency{DisplayName: "argentisk peso (1970–1983)", Symbol: ""},
				currency.ARM: cldr.Currency{DisplayName: "argentisk peso (1881–1969)", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "argentinsk peso (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "argentinsk peso", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "österrikisk schilling", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "australisk dollar", Symbol: "AUD"},
				currency.AWG: cldr.Currency{DisplayName: "arubansk florin", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "azerbajdzjansk manat (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "azerbajdzjansk manat", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "bosnisk-hercegovinsk dinar (1992–1994)", Symbol: "BAD"},
				currency.BAM: cldr.Currency{DisplayName: "bosnisk-hercegovinsk mark (konvertibel)", Symbol: "BAM"},
				currency.BAN: cldr.Currency{DisplayName: "bosnisk-hercegovinsk dinar (1994–1998)", Symbol: ""},
				currency.BBD: cldr.Currency{DisplayName: "Barbados-dollar", Symbol: "Bds$"},
				currency.BDT: cldr.Currency{DisplayName: "bangladeshisk taka", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "belgisk franc (konvertibel)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "belgisk franc", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "belgisk franc (finansiell)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "bulgarisk hård lev (1962–1999)", Symbol: ""},
				currency.BGM: cldr.Currency{DisplayName: "bulgarisk lev (1952–1962)", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "bulgarisk lev", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "bulgarisk lev (1881–1952)", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "bahrainsk dinar", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "burundisk franc", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Bermuda-dollar", Symbol: "BM$"},
				currency.BND: cldr.Currency{DisplayName: "bruneisk dollar", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "boliviansk boliviano", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "boliviansk boliviano (1864–1963)", Symbol: ""},
				currency.BOP: cldr.Currency{DisplayName: "boliviansk peso", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "boliviansk mvdol", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "brasiliansk cruzeiro novo (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "brasiliansk cruzado", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "brasiliansk cruzeiro (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "brasiliansk real", Symbol: "BR$"},
				currency.BRN: cldr.Currency{DisplayName: "brasiliansk cruzado novo", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "brasiliansk cruzeiro", Symbol: ""},
				currency.BRZ: cldr.Currency{DisplayName: "brasiliansk cruzeiro (1942–1967)", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "bahamansk dollar", Symbol: "BS$"},
				currency.BTN: cldr.Currency{DisplayName: "bhutanesisk ngultrum", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "burmesisk kyat", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "botswansk pula", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "vitrysk ny rubel (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "vitrysk rubel", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "vitrysk rubel (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "belizisk dollar", Symbol: "BZ$"},
				currency.CAD: cldr.Currency{DisplayName: "kanadensisk dollar", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "kongolesisk franc", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "euro (konvertibelt konto, WIR Bank, Schweiz)", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "schweizisk franc", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "franc (konvertibelt konto, WIR Bank, Schweiz)", Symbol: ""},
				currency.CLE: cldr.Currency{DisplayName: "chilensk escudo (1960–1975)", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "chilensk unidad de fomento", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "chilensk peso", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "kinesisk yuan (offshore)", Symbol: ""},
				currency.CNX: cldr.Currency{DisplayName: "kinesisk dollar", Symbol: "CNX"},
				currency.CNY: cldr.Currency{DisplayName: "kinesisk yuan", Symbol: "CNY"},
				currency.COP: cldr.Currency{DisplayName: "colombiansk peso", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "colombiansk unidad de valor real", Symbol: ""},
				currency.CRC: cldr.Currency{DisplayName: "costarikansk colón", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "serbisk dinar (2002–2006)", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "tjeckoslovakisk krona (–1993)", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "kubansk peso (konvertibel)", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "kubansk peso", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "kapverdisk escudo", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "cypriotiskt pund", Symbol: "CYP"},
				currency.CZK: cldr.Currency{DisplayName: "tjeckisk koruna", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "östtysk mark", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "tysk mark", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "djiboutisk franc", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "dansk krona", Symbol: "Dkr"},
				currency.DOP: cldr.Currency{DisplayName: "dominikansk peso", Symbol: "RD$"},
				currency.DZD: cldr.Currency{DisplayName: "algerisk dinar", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "ecuadoriansk sucre", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "ecuadoriansk unidad de valor constante", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "estnisk krona", Symbol: "Ekr"},
				currency.EGP: cldr.Currency{DisplayName: "egyptiskt pund", Symbol: "EG£"},
				currency.ERN: cldr.Currency{DisplayName: "eritreansk nakfa", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "spansk peseta (konto)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "spansk peseta (konvertibelt konto)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "spansk peseta", Symbol: "ESP"},
				currency.ETB: cldr.Currency{DisplayName: "etiopisk birr", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "finsk mark", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "Fijidollar", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Falklandspund", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "fransk franc", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "brittiskt pund", Symbol: "GBP"},
				currency.GEK: cldr.Currency{DisplayName: "georgisk kupon larit", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "georgisk lari", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "ghanansk cedi (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "ghanansk cedi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "gibraltiskt pund", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "gambisk dalasi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "guineansk franc", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "guineansk syli", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "ekvatorialguineansk ekwele", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "grekisk drachma", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "guatemalansk quetzal", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "Portugisiska Guinea-escudo", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "Guinea-Bissau-peso", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "Guyanadollar", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Hongkongdollar", Symbol: "HKD"},
				currency.HNL: cldr.Currency{DisplayName: "honduransk lempira", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "kroatisk dinar", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "kroatisk kuna", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "haitisk gourde", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "ungersk forint", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "indonesisk rupie", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "irländskt pund", Symbol: "IE£"},
				currency.ILP: cldr.Currency{DisplayName: "israeliskt pund", Symbol: ""},
				currency.ILR: cldr.Currency{DisplayName: "israelisk shekel (1980–1985)", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "israelisk ny shekel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "indisk rupie", Symbol: "INR"},
				currency.IQD: cldr.Currency{DisplayName: "irakisk dinar", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "iransk rial", Symbol: "IRR"},
				currency.ISJ: cldr.Currency{DisplayName: "isländsk gammal krona", Symbol: "ISJ"},
				currency.ISK: cldr.Currency{DisplayName: "isländsk krona", Symbol: "Ikr"},
				currency.ITL: cldr.Currency{DisplayName: "italiensk lire", Symbol: "ITL"},
				currency.JMD: cldr.Currency{DisplayName: "Jamaica-dollar", Symbol: "JM$"},
				currency.JOD: cldr.Currency{DisplayName: "jordansk dinar", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "japansk yen", Symbol: "JPY"},
				currency.KES: cldr.Currency{DisplayName: "kenyansk shilling", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "kirgizisk som", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "kambodjansk riel", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "komorisk franc", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "nordkoreansk won", Symbol: "KPW"},
				currency.KRH: cldr.Currency{DisplayName: "sydkoreansk hwan (1953–1962)", Symbol: ""},
				currency.KRO: cldr.Currency{DisplayName: "sydkoreansk won (1945–1953)", Symbol: ""},
				currency.KRW: cldr.Currency{DisplayName: "sydkoreansk won", Symbol: "KRW"},
				currency.KWD: cldr.Currency{DisplayName: "kuwaitisk dinar", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Cayman-dollar", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "kazakisk tenge", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "laotisk kip", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "libanesiskt pund", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "srilankesisk rupie", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "liberiansk dollar", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "lesothisk loti", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "litauisk litas", Symbol: "Lt"},
				currency.LTT: cldr.Currency{DisplayName: "litauisk talonas", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "luxemburgsk franc (konvertibel)", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "luxemburgsk franc", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "luxemburgsk franc (finansiell)", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "lettisk lats", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "lettisk rubel", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "libysk dinar", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "marockansk dirham", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "marockansk franc", Symbol: ""},
				currency.MCF: cldr.Currency{DisplayName: "monegaskisk franc (–2001)", Symbol: ""},
				currency.MDC: cldr.Currency{DisplayName: "moldavisk cupon (1992–1993)", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "moldavisk leu", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "madagaskisk ariary", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "madagaskisk franc", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "makedonisk denar", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "makedonisk denar (1992–1993)", Symbol: "MKN"},
				currency.MLF: cldr.Currency{DisplayName: "malisk franc", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "myanmarisk kyat", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "mongolisk tögrög", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "makanesisk pataca", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "mauretansk ouguiya (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "mauretansk ouguiya", Symbol: ""},
				currency.MTL: cldr.Currency{DisplayName: "maltesisk lire", Symbol: "MTL"},
				currency.MTP: cldr.Currency{DisplayName: "maltesiskt pund", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "mauritisk rupie", Symbol: "MUR"},
				currency.MVP: cldr.Currency{DisplayName: "maldivisk rupie", Symbol: "MVP"},
				currency.MVR: cldr.Currency{DisplayName: "maldivisk rufiyaa", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "malawisk kwacha", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "mexikansk peso", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "mexikansk silverpeso (1861–1992)", Symbol: ""},
				currency.MXV: cldr.Currency{DisplayName: "mexikansk unidad de inversion", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "malaysisk ringgit", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "moçambikisk escudo", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "gammal moçambikisk metical", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "moçambikisk metical", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "namibisk dollar", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "nigeriansk naira", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "nicaraguansk córdoba (1998–1991)", Symbol: "NIC"},
				currency.NIO: cldr.Currency{DisplayName: "nicaraguansk córdoba", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "nederländsk gulden", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "norsk krona", Symbol: "Nkr"},
				currency.NPR: cldr.Currency{DisplayName: "nepalesisk rupie", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "nyzeeländsk dollar", Symbol: "NZD"},
				currency.OMR: cldr.Currency{DisplayName: "omansk rial", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "panamansk balboa", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "peruansk inti", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "peruansk sol", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "peruansk sol (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "papuansk kina", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "filippinsk peso", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "pakistansk rupie", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "polsk zloty", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "polsk zloty (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "portugisisk escudo", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "paraguayansk guarani", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "qatarisk rial", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "rhodesisk dollar", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "rumänsk leu (1952–2005)", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "rumänsk leu", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "serbisk dinar", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "rysk rubel", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "rysk rubel (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "rwandisk franc", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "saudisk riyal", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Salomondollar", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "seychellisk rupie", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "sudansk dinar (1992–2007)", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "sudanesiskt pund", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "sudanskt pund (1916–1992)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "svensk krona", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "singaporiansk dollar", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "S:t Helena-pund", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "slovensk tolar", Symbol: "SIT"},
				currency.SKK: cldr.Currency{DisplayName: "slovakisk koruna", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "sierraleonsk leone", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "somalisk shilling", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "surinamesisk dollar", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "surinamesisk gulden", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "sydsudanesiskt pund", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "saotomeansk dobra (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "saotomeansk dobra", Symbol: "Db"},
				currency.SUR: cldr.Currency{DisplayName: "sovjetisk rubel", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "salvadoransk colón", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "syriskt pund", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "swaziländsk lilangeni", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "thailändsk baht", Symbol: "THB"},
				currency.TJR: cldr.Currency{DisplayName: "tadzjikisk rubel", Symbol: "TJR"},
				currency.TJS: cldr.Currency{DisplayName: "tadzjikisk somoni", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "turkmenistansk manat (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "turkmenistansk manat", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "tunisisk dinar", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "tongansk paʻanga", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "östtimoresisk escudo", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "turkisk lire (1922–2005)", Symbol: "TRL"},
				currency.TRY: cldr.Currency{DisplayName: "turkisk lira", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Trinidad och Tobago-dollar", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Taiwandollar", Symbol: "TWD"},
				currency.TZS: cldr.Currency{DisplayName: "tanzanisk shilling", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "ukrainsk hryvnia", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "ukrainsk karbovanetz", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "ugandisk shilling (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "ugandisk shilling", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "US-dollar", Symbol: "US$"},
				currency.USN: cldr.Currency{DisplayName: "US-dollar (nästa dag)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "US-dollar (samma dag)", Symbol: ""},
				currency.UYI: cldr.Currency{DisplayName: "uruguayansk peso en unidades indexadas", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "uruguayansk peso (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "uruguayansk peso", Symbol: "UYU"},
				currency.UYW: cldr.Currency{DisplayName: "uruguayansk indexenhet för nominell lön", Symbol: ""},
				currency.UZS: cldr.Currency{DisplayName: "uzbekisk sum", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "venezuelansk bolivar (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "venezuelansk bolívar (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "venezuelansk bolívar", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "vietnamesisk dong", Symbol: "VND"},
				currency.VNN: cldr.Currency{DisplayName: "vietnamesisk dong (1978–1985)", Symbol: ""},
				currency.VUV: cldr.Currency{DisplayName: "vanuatisk vatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "västsamoansk tala", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "centralafrikansk franc", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "silver", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "guld", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "europeisk kompositenhet", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "europeisk monetär enhet", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "europeisk kontoenhet (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "europeisk kontoenhet (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "östkaribisk dollar", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "IMF särskild dragningsrätt", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "europeisk valutaenhet", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "fransk guldfranc", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "internationella järnvägsunionens franc", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "västafrikansk franc", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "palladium", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "CFP-franc", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "platina", Symbol: ""},
				currency.XRE: cldr.Currency{DisplayName: "RINET-fond", Symbol: ""},
				currency.XSU: cldr.Currency{DisplayName: "latinamerikansk sucre", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "testvalutaenhet", Symbol: "XTS"},
				currency.XUA: cldr.Currency{DisplayName: "afrikansk kontoenhet", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "okänd valuta", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "jemenitisk dinar", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "jemenitisk rial", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "jugoslavisk dinar (1966–1990)", Symbol: "YUD"},
				currency.YUM: cldr.Currency{DisplayName: "jugoslavisk dinar (1994–2002)", Symbol: "YUM"},
				currency.YUN: cldr.Currency{DisplayName: "jugoslavisk dinar (1990–1992)", Symbol: "YUN"},
				currency.YUR: cldr.Currency{DisplayName: "jugoslavisk dinar (1992–1993)", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "sydafrikansk rand (finansiell)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "sydafrikansk rand", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "zambisk kwacha (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "zambisk kwacha", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "zairisk ny zaire", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "zairisk zaire", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "Zimbabwe-dollar", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "Zimbabwe-dollar (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "Zimbabwe-dollar (2008)", Symbol: ""},
			},
		},
	}
}
