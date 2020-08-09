// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_it() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "gen", Feb: "feb", Mar: "mar", Apr: "apr", May: "mag", Jun: "giu", Jul: "lug", Aug: "ago", Sep: "set", Oct: "ott", Nov: "nov", Dec: "dic"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "G", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "G", Jul: "L", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "gennaio", Feb: "febbraio", Mar: "marzo", Apr: "aprile", May: "maggio", Jun: "giugno", Jul: "luglio", Aug: "agosto", Sep: "settembre", Oct: "ottobre", Nov: "novembre", Dec: "dicembre"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "dom", Mon: "lun", Tue: "mar", Wed: "mer", Thu: "gio", Fri: "ven", Sat: "sab"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "M", Thu: "G", Fri: "V", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "dom", Mon: "lun", Tue: "mar", Wed: "mer", Thu: "gio", Fri: "ven", Sat: "sab"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "domenica", Mon: "lunedì", Tue: "martedì", Wed: "mercoledì", Thu: "giovedì", Fri: "venerdì", Sat: "sabato"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "m.", PM: "p."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_it]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "peseta andorrana", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "dirham degli Emirati Arabi Uniti", Symbol: ""},
				currency.AFA: cldr.Currency{DisplayName: "afgani (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "afghani", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "lek albanese", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "dram armeno", Symbol: ""},
				currency.ANG: cldr.Currency{DisplayName: "fiorino delle Antille olandesi", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "kwanza angolano", Symbol: "Kz"},
				currency.AOK: cldr.Currency{DisplayName: "kwanza angolano (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "nuovo kwanza angolano (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "kwanza reajustado angolano (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "austral argentino", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "peso argentino (vecchio Cod.)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "peso argentino", Symbol: "$"},
				currency.ATS: cldr.Currency{DisplayName: "scellino austriaco", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "dollaro australiano", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "fiorino di Aruba", Symbol: ""},
				currency.AZM: cldr.Currency{DisplayName: "manat azero (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "manat azero", Symbol: ""},
				currency.BAD: cldr.Currency{DisplayName: "dinar Bosnia-Herzegovina", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "marco convertibile della Bosnia-Herzegovina", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "dollaro di Barbados", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "taka bangladese", Symbol: "৳"},
				currency.BEC: cldr.Currency{DisplayName: "franco belga (convertibile)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "franco belga", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "franco belga (finanziario)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "lev bulgaro (1962–1999)", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "lev bulgaro", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "dinaro del Bahrein", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "franco del Burundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "dollaro delle Bermuda", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "dollaro del Brunei", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "boliviano", Symbol: "Bs"},
				currency.BOP: cldr.Currency{DisplayName: "peso boliviano", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "mvdol boliviano", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "cruzeiro novo brasiliano (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "cruzado brasiliano", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "cruzeiro brasiliano (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "real brasiliano", Symbol: "BRL"},
				currency.BRN: cldr.Currency{DisplayName: "cruzado novo brasiliano", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "cruzeiro brasiliano", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "dollaro delle Bahamas", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "ngultrum bhutanese", Symbol: ""},
				currency.BUK: cldr.Currency{DisplayName: "kyat birmano", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "pula del Botswana", Symbol: "P"},
				currency.BYB: cldr.Currency{DisplayName: "nuovo rublo bielorusso (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "rublo bielorusso", Symbol: "Br"},
				currency.BYR: cldr.Currency{DisplayName: "rublo bielorusso (2000–2016)", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "dollaro del Belize", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "dollaro canadese", Symbol: "$"},
				currency.CDF: cldr.Currency{DisplayName: "franco congolese", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "franco svizzero", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "unidades de fomento chilene", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "peso cileno", Symbol: "$"},
				currency.CNH: cldr.Currency{DisplayName: "renmimbi cinese offshore", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "renminbi cinese", Symbol: "¥"},
				currency.COP: cldr.Currency{DisplayName: "peso colombiano", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "colón costaricano", Symbol: "₡"},
				currency.CSD: cldr.Currency{DisplayName: "antico dinaro serbo", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "corona forte cecoslovacca", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "peso cubano convertibile", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "peso cubano", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "escudo capoverdiano", Symbol: ""},
				currency.CYP: cldr.Currency{DisplayName: "sterlina cipriota", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "corona ceca", Symbol: "Kč"},
				currency.DDM: cldr.Currency{DisplayName: "ostmark della Germania Orientale", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "marco tedesco", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "franco di Gibuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "corona danese", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "peso dominicano", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "dinaro algerino", Symbol: ""},
				currency.ECS: cldr.Currency{DisplayName: "sucre dell’Ecuador", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "unidad de valor constante (UVC) dell’Ecuador", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "corona dell’Estonia", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "sterlina egiziana", Symbol: "£E"},
				currency.ERN: cldr.Currency{DisplayName: "nakfa eritreo", Symbol: ""},
				currency.ESA: cldr.Currency{DisplayName: "peseta spagnola account", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "peseta spagnola account convertibile", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "peseta spagnola", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "birr etiope", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "markka finlandese", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "dollaro delle Figi", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "sterlina delle Falkland", Symbol: "£"},
				currency.FRF: cldr.Currency{DisplayName: "franco francese", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "sterlina britannica", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "kupon larit georgiano", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "lari georgiano", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "cedi del Ghana", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "cedi ghanese", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "sterlina di Gibilterra", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "dalasi gambiano", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "franco della Guinea", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "syli della Guinea", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "ekwele della Guinea Equatoriale", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "dracma greca", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "quetzal guatemalteco", Symbol: "Q"},
				currency.GWE: cldr.Currency{DisplayName: "escudo della Guinea portoghese", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "peso della Guinea-Bissau", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "dollaro della Guyana", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "dollaro di Hong Kong", Symbol: "HKD"},
				currency.HNL: cldr.Currency{DisplayName: "lempira honduregna", Symbol: "L"},
				currency.HRD: cldr.Currency{DisplayName: "dinaro croato", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "kuna croata", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "gourde haitiano", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "fiorino ungherese", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "rupia indonesiana", Symbol: "Rp"},
				currency.IEP: cldr.Currency{DisplayName: "sterlina irlandese", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "sterlina israeliana", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "nuovo siclo israeliano", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "rupia indiana", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "dinaro iracheno", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "rial iraniano", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "corona islandese", Symbol: "kr"},
				currency.ITL: cldr.Currency{DisplayName: "lira italiana", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "dollaro giamaicano", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "dinaro giordano", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "yen giapponese", Symbol: "JPY"},
				currency.KES: cldr.Currency{DisplayName: "scellino keniota", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "som kirghiso", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "riel cambogiano", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "franco comoriano", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "won nordcoreano", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "won sudcoreano", Symbol: "KRW"},
				currency.KWD: cldr.Currency{DisplayName: "dinaro kuwaitiano", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "dollaro delle Isole Cayman", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "tenge kazako", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "kip laotiano", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "lira libanese", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "rupia di Sri Lanka", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "dollaro liberiano", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "loti del Lesotho", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "litas lituano", Symbol: "Lt"},
				currency.LTT: cldr.Currency{DisplayName: "talonas lituani", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "franco convertibile del Lussemburgo", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "franco del Lussemburgo", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "franco finanziario del Lussemburgo", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "lats lettone", Symbol: "Ls"},
				currency.LVR: cldr.Currency{DisplayName: "rublo lettone", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "dinaro libico", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "dirham marocchino", Symbol: ""},
				currency.MAF: cldr.Currency{DisplayName: "franco marocchino", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "leu moldavo", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "ariary malgascio", Symbol: "Ar"},
				currency.MGF: cldr.Currency{DisplayName: "franco malgascio", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "denar macedone", Symbol: ""},
				currency.MLF: cldr.Currency{DisplayName: "franco di Mali", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "kyat di Myanmar", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "tugrik mongolo", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "pataca di Macao", Symbol: ""},
				currency.MRO: cldr.Currency{DisplayName: "ouguiya della Mauritania (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "ouguiya della Mauritania", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "lira maltese", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "sterlina maltese", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "rupia mauriziana", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "rufiyaa delle Maldive", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "kwacha malawiano", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "peso messicano", Symbol: "MXN"},
				currency.MXP: cldr.Currency{DisplayName: "peso messicano d’argento (1861–1992)", Symbol: ""},
				currency.MXV: cldr.Currency{DisplayName: "unidad de inversion (UDI) messicana", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "ringgit malese", Symbol: "RM"},
				currency.MZE: cldr.Currency{DisplayName: "escudo del Mozambico", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "metical mozambicano", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "dollaro namibiano", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "naira nigeriana", Symbol: "₦"},
				currency.NIC: cldr.Currency{DisplayName: "cordoba nicaraguense", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "córdoba nicaraguense", Symbol: "C$"},
				currency.NLG: cldr.Currency{DisplayName: "fiorino olandese", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "corona norvegese", Symbol: "NKr"},
				currency.NPR: cldr.Currency{DisplayName: "rupia nepalese", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "dollaro neozelandese", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "rial omanita", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "balboa panamense", Symbol: ""},
				currency.PEI: cldr.Currency{DisplayName: "inti peruviano", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "sol peruviano", Symbol: ""},
				currency.PES: cldr.Currency{DisplayName: "sol peruviano (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "kina papuana", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "peso filippino", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "rupia pakistana", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "złoty polacco", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "złoty Polacco (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "escudo portoghese", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "guaraní paraguayano", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "rial qatariano", Symbol: ""},
				currency.RHD: cldr.Currency{DisplayName: "dollaro della Rhodesia", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "leu della Romania", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "leu rumeno", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "dinaro serbo", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "rublo russo", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "rublo della CSI", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "franco ruandese", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "riyal saudita", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "dollaro delle Isole Salomone", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "rupia delle Seychelles", Symbol: ""},
				currency.SDD: cldr.Currency{DisplayName: "dinaro sudanese", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "sterlina sudanese", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "corona svedese", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "dollaro di Singapore", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "sterlina di Sant’Elena", Symbol: "£"},
				currency.SIT: cldr.Currency{DisplayName: "tallero sloveno", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "corona slovacca", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "leone della Sierra Leone", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "scellino somalo", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "dollaro del Suriname", Symbol: "$"},
				currency.SRG: cldr.Currency{DisplayName: "fiorino del Suriname", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "sterlina sud-sudanese", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "dobra di Sao Tomé e Principe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "dobra di Sao Tomé e Príncipe", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "rublo sovietico", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "colón salvadoregno", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "lira siriana", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "lilangeni dello Swaziland", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "baht thailandese", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "rublo del Tajikistan", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "somoni tagiko", Symbol: ""},
				currency.TMM: cldr.Currency{DisplayName: "manat turkmeno (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "manat turkmeno", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "dinaro tunisino", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "paʻanga tongano", Symbol: "T$"},
				currency.TPE: cldr.Currency{DisplayName: "escudo di Timor", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "lira turca (1922–2005)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "lira turca", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "dollaro di Trinidad e Tobago", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "nuovo dollaro taiwanese", Symbol: "TWD"},
				currency.TZS: cldr.Currency{DisplayName: "scellino della Tanzania", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "grivnia ucraina", Symbol: "₴"},
				currency.UAK: cldr.Currency{DisplayName: "karbovanetz ucraino", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "scellino ugandese (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "scellino ugandese", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "dollaro statunitense", Symbol: "USD"},
				currency.USN: cldr.Currency{DisplayName: "dollaro statunitense (next day)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "dollaro statunitense (same day)", Symbol: ""},
				currency.UYI: cldr.Currency{DisplayName: "peso uruguaiano in unità indicizzate", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "peso uruguaiano (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "peso uruguayano", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "sum uzbeco", Symbol: ""},
				currency.VEB: cldr.Currency{DisplayName: "bolivar venezuelano (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "bolívar venezuelano (2008–2018)", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "bolívar venezuelano", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "dong vietnamita", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "vatu di Vanuatu", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "tala samoano", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "franco CFA BEAC", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "argento", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "oro", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "unità composita europea", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "unità monetaria europea", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "unità di acconto europea (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "unità di acconto europea (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "dollaro dei Caraibi orientali", Symbol: "$"},
				currency.XDR: cldr.Currency{DisplayName: "diritti speciali di incasso", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "franco oro francese", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "franco UIC francese", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "franco CFA BCEAO", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "palladio", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "franco CFP", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "platino", Symbol: ""},
				currency.XRE: cldr.Currency{DisplayName: "fondi RINET", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "codice di verifica della valuta", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "valuta sconosciuta", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "dinaro dello Yemen", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "riyal yemenita", Symbol: ""},
				currency.YUD: cldr.Currency{DisplayName: "dinaro forte yugoslavo", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "dinaro noviy yugoslavo", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "dinaro convertibile yugoslavo", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "rand sudafricano (finanziario)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "rand sudafricano", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "kwacha dello Zambia (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "kwacha dello Zambia", Symbol: "ZK"},
				currency.ZRN: cldr.Currency{DisplayName: "nuovo zaire dello Zaire", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "zaire dello Zaire", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "dollaro dello Zimbabwe", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "dollaro zimbabwiano (2009)", Symbol: ""},
			},
		},
	}
}