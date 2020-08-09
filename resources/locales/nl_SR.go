// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_nl_SR() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "nl_SR",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd-MM-y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'om' {0}", Long: "{1} 'om' {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan.", Feb: "feb.", Mar: "mrt.", Apr: "apr.", May: "mei", Jun: "jun.", Jul: "jul.", Aug: "aug.", Sep: "sep.", Oct: "okt.", Nov: "nov.", Dec: "dec."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "januari", Feb: "februari", Mar: "maart", Apr: "april", May: "mei", Jun: "juni", Jul: "juli", Aug: "augustus", Sep: "september", Oct: "oktober", Nov: "november", Dec: "december"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "zo", Mon: "ma", Tue: "di", Wed: "wo", Thu: "do", Fri: "vr", Sat: "za"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Z", Mon: "M", Tue: "D", Wed: "W", Thu: "D", Fri: "V", Sat: "Z"}, Short: cldr.CalendarDayFormatNameValue{Sun: "zo", Mon: "ma", Tue: "di", Wed: "wo", Thu: "do", Fri: "vr", Sat: "za"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "zondag", Mon: "maandag", Tue: "dinsdag", Wed: "woensdag", Thu: "donderdag", Fri: "vrijdag", Sat: "zaterdag"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_nl]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "٫", Group: "٬", Negative: "\u061c-", Percent: "٪\u061c", PerMille: "؉"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "Andorrese peseta", Symbol: "ADP"},
				currency.AED: cldr.Currency{DisplayName: "Verenigde Arabische Emiraten-dirham", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "Afghani (1927–2002)", Symbol: "AFA"},
				currency.AFN: cldr.Currency{DisplayName: "Afghaanse afghani", Symbol: "AFN"},
				currency.ALK: cldr.Currency{DisplayName: "Albanese lek (1946–1965)", Symbol: "ALK"},
				currency.ALL: cldr.Currency{DisplayName: "Albanese lek", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Armeense dram", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Nederlands-Antilliaanse gulden", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Angolese kwanza", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "Angolese kwanza (1977–1990)", Symbol: "AOK"},
				currency.AON: cldr.Currency{DisplayName: "Angolese nieuwe kwanza (1990–2000)", Symbol: "AON"},
				currency.AOR: cldr.Currency{DisplayName: "Angolese kwanza reajustado (1995–1999)", Symbol: "AOR"},
				currency.ARA: cldr.Currency{DisplayName: "Argentijnse austral", Symbol: "ARA"},
				currency.ARL: cldr.Currency{DisplayName: "Argentijnse peso ley (1970–1983)", Symbol: "ARL"},
				currency.ARM: cldr.Currency{DisplayName: "Argentijnse peso (1881–1970)", Symbol: "ARM"},
				currency.ARP: cldr.Currency{DisplayName: "Argentijnse peso (1983–1985)", Symbol: "ARP"},
				currency.ARS: cldr.Currency{DisplayName: "Argentijnse peso", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "Oostenrijkse schilling", Symbol: "ATS"},
				currency.AUD: cldr.Currency{DisplayName: "Australische dollar", Symbol: "AU$"},
				currency.AWG: cldr.Currency{DisplayName: "Arubaanse gulden", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "Azerbeidzjaanse manat (1993–2006)", Symbol: "AZM"},
				currency.AZN: cldr.Currency{DisplayName: "Azerbeidzjaanse manat", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "Bosnische dinar", Symbol: "BAD"},
				currency.BAM: cldr.Currency{DisplayName: "Bosnische convertibele mark", Symbol: "BAM"},
				currency.BAN: cldr.Currency{DisplayName: "Nieuwe Bosnische dinar (1994–1997)", Symbol: "BAN"},
				currency.BBD: cldr.Currency{DisplayName: "Barbadaanse dollar", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Bengalese taka", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "Belgische frank (convertibel)", Symbol: "BEC"},
				currency.BEF: cldr.Currency{DisplayName: "Belgische frank", Symbol: "BEF"},
				currency.BEL: cldr.Currency{DisplayName: "Belgische frank (financieel)", Symbol: "BEL"},
				currency.BGL: cldr.Currency{DisplayName: "Bulgaarse harde lev", Symbol: "BGL"},
				currency.BGM: cldr.Currency{DisplayName: "Bulgaarse socialistische lev", Symbol: "BGM"},
				currency.BGN: cldr.Currency{DisplayName: "Bulgaarse lev", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "Bulgaarse lev (1879–1952)", Symbol: "BGO"},
				currency.BHD: cldr.Currency{DisplayName: "Bahreinse dinar", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Burundese frank", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Bermuda-dollar", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Bruneise dollar", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Boliviaanse boliviano", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "Boliviaanse boliviano (1863–1963)", Symbol: "BOL"},
				currency.BOP: cldr.Currency{DisplayName: "Boliviaanse peso", Symbol: "BOP"},
				currency.BOV: cldr.Currency{DisplayName: "Boliviaanse mvdol", Symbol: "BOV"},
				currency.BRB: cldr.Currency{DisplayName: "Braziliaanse cruzeiro novo (1967–1986)", Symbol: "BRB"},
				currency.BRC: cldr.Currency{DisplayName: "Braziliaanse cruzado", Symbol: "BRC"},
				currency.BRE: cldr.Currency{DisplayName: "Braziliaanse cruzeiro (1990–1993)", Symbol: "BRE"},
				currency.BRL: cldr.Currency{DisplayName: "Braziliaanse real", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "Braziliaanse nieuwe cruzado (1989–1990)", Symbol: "BRN"},
				currency.BRR: cldr.Currency{DisplayName: "Braziliaanse cruzeiro", Symbol: "BRR"},
				currency.BRZ: cldr.Currency{DisplayName: "Braziliaanse cruzeiro (1942–1967)", Symbol: "BRZ"},
				currency.BSD: cldr.Currency{DisplayName: "Bahamaanse dollar", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Bhutaanse ngultrum", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "Birmese kyat", Symbol: "BUK"},
				currency.BWP: cldr.Currency{DisplayName: "Botswaanse pula", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "Wit-Russische nieuwe roebel (1994–1999)", Symbol: "BYB"},
				currency.BYN: cldr.Currency{DisplayName: "Wit-Russische roebel", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Wit-Russische roebel (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "Belizaanse dollar", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Canadese dollar", Symbol: "C$"},
				currency.CDF: cldr.Currency{DisplayName: "Congolese frank", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "WIR euro", Symbol: "CHE"},
				currency.CHF: cldr.Currency{DisplayName: "Zwitserse frank", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "WIR franc", Symbol: "CHW"},
				currency.CLE: cldr.Currency{DisplayName: "Chileense escudo", Symbol: "CLE"},
				currency.CLF: cldr.Currency{DisplayName: "Chileense unidades de fomento", Symbol: "CLF"},
				currency.CLP: cldr.Currency{DisplayName: "Chileense peso", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Chinese yuan (offshore)", Symbol: "CNH"},
				currency.CNX: cldr.Currency{DisplayName: "dollar van de Chinese Volksbank", Symbol: "CNX"},
				currency.CNY: cldr.Currency{DisplayName: "Chinese yuan", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Colombiaanse peso", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "Unidad de Valor Real", Symbol: "COU"},
				currency.CRC: cldr.Currency{DisplayName: "Costa Ricaanse colon", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "Oude Servische dinar", Symbol: "CSD"},
				currency.CSK: cldr.Currency{DisplayName: "Tsjechoslowaakse harde koruna", Symbol: "CSK"},
				currency.CUC: cldr.Currency{DisplayName: "Cubaanse convertibele peso", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Cubaanse peso", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Kaapverdische escudo", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "Cyprisch pond", Symbol: "CYP"},
				currency.CZK: cldr.Currency{DisplayName: "Tsjechische kroon", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "Oost-Duitse ostmark", Symbol: "DDM"},
				currency.DEM: cldr.Currency{DisplayName: "Duitse mark", Symbol: "DEM"},
				currency.DJF: cldr.Currency{DisplayName: "Djiboutiaanse frank", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Deense kroon", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Dominicaanse peso", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Algerijnse dinar", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "Ecuadoraanse sucre", Symbol: "ECS"},
				currency.ECV: cldr.Currency{DisplayName: "Ecuadoraanse unidad de valor constante (UVC)", Symbol: "ECV"},
				currency.EEK: cldr.Currency{DisplayName: "Estlandse kroon", Symbol: "EEK"},
				currency.EGP: cldr.Currency{DisplayName: "Egyptisch pond", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Eritrese nakfa", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "Spaanse peseta (account A)", Symbol: "ESA"},
				currency.ESB: cldr.Currency{DisplayName: "Spaanse peseta (convertibele account)", Symbol: "ESB"},
				currency.ESP: cldr.Currency{DisplayName: "Spaanse peseta", Symbol: "ESP"},
				currency.ETB: cldr.Currency{DisplayName: "Ethiopische birr", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Finse markka", Symbol: "FIM"},
				currency.FJD: cldr.Currency{DisplayName: "Fiji-dollar", Symbol: "FJ$"},
				currency.FKP: cldr.Currency{DisplayName: "Falklandeilands pond", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "Franse franc", Symbol: "FRF"},
				currency.GBP: cldr.Currency{DisplayName: "Brits pond", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "Georgische kupon larit", Symbol: "GEK"},
				currency.GEL: cldr.Currency{DisplayName: "Georgische lari", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "Ghanese cedi (1979–2007)", Symbol: "GHC"},
				currency.GHS: cldr.Currency{DisplayName: "Ghanese cedi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Gibraltarees pond", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Gambiaanse dalasi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Guinese frank", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "Guinese syli", Symbol: "GNS"},
				currency.GQE: cldr.Currency{DisplayName: "Equatoriaal-Guinese ekwele guineana", Symbol: "GQE"},
				currency.GRD: cldr.Currency{DisplayName: "Griekse drachme", Symbol: "GRD"},
				currency.GTQ: cldr.Currency{DisplayName: "Guatemalteekse quetzal", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "Portugees-Guinese escudo", Symbol: "GWE"},
				currency.GWP: cldr.Currency{DisplayName: "Guinee-Bissause peso", Symbol: "GWP"},
				currency.GYD: cldr.Currency{DisplayName: "Guyaanse dollar", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Hongkongse dollar", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Hondurese lempira", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "Kroatische dinar", Symbol: "HRD"},
				currency.HRK: cldr.Currency{DisplayName: "Kroatische kuna", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Haïtiaanse gourde", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Hongaarse forint", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Indonesische roepia", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "Iers pond", Symbol: "IEP"},
				currency.ILP: cldr.Currency{DisplayName: "Israëlisch pond", Symbol: "ILP"},
				currency.ILR: cldr.Currency{DisplayName: "Israëlische sjekel (1980–1985)", Symbol: "ILR"},
				currency.ILS: cldr.Currency{DisplayName: "Israëlische nieuwe shekel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Indiase roepie", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Iraakse dinar", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Iraanse rial", Symbol: "IRR"},
				currency.ISJ: cldr.Currency{DisplayName: "IJslandse kroon (1918–1981)", Symbol: "ISJ"},
				currency.ISK: cldr.Currency{DisplayName: "IJslandse kroon", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "Italiaanse lire", Symbol: "ITL"},
				currency.JMD: cldr.Currency{DisplayName: "Jamaicaanse dollar", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Jordaanse dinar", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Japanse yen", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Keniaanse shilling", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Kirgizische som", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Cambodjaanse riel", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Comorese frank", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Noord-Koreaanse won", Symbol: "KPW"},
				currency.KRH: cldr.Currency{DisplayName: "Zuid-Koreaanse hwan (1953–1962)", Symbol: "KRH"},
				currency.KRO: cldr.Currency{DisplayName: "Oude Zuid-Koreaanse won (1945–1953)", Symbol: "KRO"},
				currency.KRW: cldr.Currency{DisplayName: "Zuid-Koreaanse won", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Koeweitse dinar", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Kaaimaneilandse dollar", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Kazachse tenge", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Laotiaanse kip", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Libanees pond", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Sri Lankaanse roepie", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Liberiaanse dollar", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Lesothaanse loti", Symbol: "LSL"},
				currency.LTL: cldr.Currency{DisplayName: "Litouwse litas", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "Litouwse talonas", Symbol: "LTT"},
				currency.LUC: cldr.Currency{DisplayName: "Luxemburgse convertibele franc", Symbol: "LUC"},
				currency.LUF: cldr.Currency{DisplayName: "Luxemburgse frank", Symbol: "LUF"},
				currency.LUL: cldr.Currency{DisplayName: "Luxemburgse financiële franc", Symbol: "LUL"},
				currency.LVL: cldr.Currency{DisplayName: "Letse lats", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "Letse roebel", Symbol: "LVR"},
				currency.LYD: cldr.Currency{DisplayName: "Libische dinar", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Marokkaanse dirham", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "Marokkaanse franc", Symbol: "MAF"},
				currency.MCF: cldr.Currency{DisplayName: "Monegaskische frank", Symbol: "MCF"},
				currency.MDC: cldr.Currency{DisplayName: "Moldavische cupon", Symbol: "MDC"},
				currency.MDL: cldr.Currency{DisplayName: "Moldavische leu", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Malagassische ariary", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "Malagassische franc", Symbol: "MGF"},
				currency.MKD: cldr.Currency{DisplayName: "Macedonische denar", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "Macedonische denar (1992–1993)", Symbol: "MKN"},
				currency.MLF: cldr.Currency{DisplayName: "Malinese franc", Symbol: "MLF"},
				currency.MMK: cldr.Currency{DisplayName: "Myanmarese kyat", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Mongoolse tugrik", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Macause pataca", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Mauritaanse ouguiya (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Mauritaanse ouguiya", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "Maltese lire", Symbol: "MTL"},
				currency.MTP: cldr.Currency{DisplayName: "Maltees pond", Symbol: "MTP"},
				currency.MUR: cldr.Currency{DisplayName: "Mauritiaanse roepie", Symbol: "MUR"},
				currency.MVP: cldr.Currency{DisplayName: "Maldivische roepie", Symbol: "MVP"},
				currency.MVR: cldr.Currency{DisplayName: "Maldivische rufiyaa", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Malawische kwacha", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Mexicaanse peso", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "Mexicaanse zilveren peso (1861–1992)", Symbol: "MXP"},
				currency.MXV: cldr.Currency{DisplayName: "Mexicaanse unidad de inversion (UDI)", Symbol: "MXV"},
				currency.MYR: cldr.Currency{DisplayName: "Maleisische ringgit", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "Mozambikaanse escudo", Symbol: "MZE"},
				currency.MZM: cldr.Currency{DisplayName: "Oude Mozambikaanse metical", Symbol: "MZM"},
				currency.MZN: cldr.Currency{DisplayName: "Mozambikaanse metical", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Namibische dollar", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Nigeriaanse naira", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "Nicaraguaanse córdoba (1988–1991)", Symbol: "NIC"},
				currency.NIO: cldr.Currency{DisplayName: "Nicaraguaanse córdoba", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "Nederlandse gulden", Symbol: "NLG"},
				currency.NOK: cldr.Currency{DisplayName: "Noorse kroon", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Nepalese roepie", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Nieuw-Zeelandse dollar", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Omaanse rial", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Panamese balboa", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "Peruaanse inti", Symbol: "PEI"},
				currency.PEN: cldr.Currency{DisplayName: "Peruaanse sol", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "Peruaanse sol (1863–1965)", Symbol: "PES"},
				currency.PGK: cldr.Currency{DisplayName: "Papoea-Nieuw-Guinese kina", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Filipijnse peso", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Pakistaanse roepie", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Poolse zloty", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "Poolse zloty (1950–1995)", Symbol: "PLZ"},
				currency.PTE: cldr.Currency{DisplayName: "Portugese escudo", Symbol: "PTE"},
				currency.PYG: cldr.Currency{DisplayName: "Paraguayaanse guarani", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Qatarese rial", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "Rhodesische dollar", Symbol: "RHD"},
				currency.ROL: cldr.Currency{DisplayName: "Oude Roemeense leu", Symbol: "ROL"},
				currency.RON: cldr.Currency{DisplayName: "Roemeense leu", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Servische dinar", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Russische roebel", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "Russische roebel (1991–1998)", Symbol: "RUR"},
				currency.RWF: cldr.Currency{DisplayName: "Rwandese frank", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Saoedi-Arabische riyal", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Salomon-dollar", Symbol: "SI$"},
				currency.SCR: cldr.Currency{DisplayName: "Seychelse roepie", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "Soedanese dinar", Symbol: "SDD"},
				currency.SDG: cldr.Currency{DisplayName: "Soedanees pond", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "Soedanees pond (1957–1998)", Symbol: "SDP"},
				currency.SEK: cldr.Currency{DisplayName: "Zweedse kroon", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Singaporese dollar", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Sint-Heleens pond", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "Sloveense tolar", Symbol: "SIT"},
				currency.SKK: cldr.Currency{DisplayName: "Slowaakse koruna", Symbol: "SKK"},
				currency.SLL: cldr.Currency{DisplayName: "Sierraleoonse leone", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Somalische shilling", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Surinaamse dollar", Symbol: "$"},
				currency.SRG: cldr.Currency{DisplayName: "Surinaamse gulden", Symbol: "SRG"},
				currency.SSP: cldr.Currency{DisplayName: "Zuid-Soedanees pond", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "Santomese dobra (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "Santomese dobra", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "Sovjet-roebel", Symbol: "SUR"},
				currency.SVC: cldr.Currency{DisplayName: "Salvadoraanse colón", Symbol: "SVC"},
				currency.SYP: cldr.Currency{DisplayName: "Syrisch pond", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Swazische lilangeni", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Thaise baht", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "Tadzjikistaanse roebel", Symbol: "TJR"},
				currency.TJS: cldr.Currency{DisplayName: "Tadzjiekse somoni", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "Turkmeense manat (1993–2009)", Symbol: "TMM"},
				currency.TMT: cldr.Currency{DisplayName: "Turkmeense manat", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Tunesische dinar", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Tongaanse paʻanga", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "Timorese escudo", Symbol: "TPE"},
				currency.TRL: cldr.Currency{DisplayName: "Turkse lire", Symbol: "TRL"},
				currency.TRY: cldr.Currency{DisplayName: "Turkse lira", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Trinidad en Tobago-dollar", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Nieuwe Taiwanese dollar", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Tanzaniaanse shilling", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Oekraïense hryvnia", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "Oekraïense karbovanetz", Symbol: "UAK"},
				currency.UGS: cldr.Currency{DisplayName: "Oegandese shilling (1966–1987)", Symbol: "UGS"},
				currency.UGX: cldr.Currency{DisplayName: "Oegandese shilling", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "Amerikaanse dollar", Symbol: "US$"},
				currency.USN: cldr.Currency{DisplayName: "Amerikaanse dollar (volgende dag)", Symbol: "USN"},
				currency.USS: cldr.Currency{DisplayName: "Amerikaanse dollar (zelfde dag)", Symbol: "USS"},
				currency.UYI: cldr.Currency{DisplayName: "Uruguayaanse peso en geïndexeerde eenheden", Symbol: "UYI"},
				currency.UYP: cldr.Currency{DisplayName: "Uruguayaanse peso (1975–1993)", Symbol: "UYP"},
				currency.UYU: cldr.Currency{DisplayName: "Uruguayaanse peso", Symbol: "UYU"},
				currency.UYW: cldr.Currency{DisplayName: "Uruguayaanse nominale salarisindexeenheid", Symbol: ""},
				currency.UZS: cldr.Currency{DisplayName: "Oezbeekse sum", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "Venezolaanse bolivar (1871–2008)", Symbol: "VEB"},
				currency.VEF: cldr.Currency{DisplayName: "Venezolaanse bolivar (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Venezolaanse bolivar", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Vietnamese dong", Symbol: "₫"},
				currency.VNN: cldr.Currency{DisplayName: "Vietnamese dong (1978–1985)", Symbol: "VNN"},
				currency.VUV: cldr.Currency{DisplayName: "Vanuatuaanse vatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Samoaanse tala", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "CFA-frank", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "Zilver", Symbol: "XAG"},
				currency.XAU: cldr.Currency{DisplayName: "Goud", Symbol: "XAU"},
				currency.XBA: cldr.Currency{DisplayName: "Europese samengestelde eenheid", Symbol: "XBA"},
				currency.XBB: cldr.Currency{DisplayName: "Europese monetaire eenheid", Symbol: "XBB"},
				currency.XBC: cldr.Currency{DisplayName: "Europese rekeneenheid (XBC)", Symbol: "XBC"},
				currency.XBD: cldr.Currency{DisplayName: "Europese rekeneenheid (XBD)", Symbol: "XBD"},
				currency.XCD: cldr.Currency{DisplayName: "Oost-Caribische dollar", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "Special Drawing Rights", Symbol: "XDR"},
				currency.XEU: cldr.Currency{DisplayName: "European Currency Unit", Symbol: "XEU"},
				currency.XFO: cldr.Currency{DisplayName: "Franse gouden franc", Symbol: "XFO"},
				currency.XFU: cldr.Currency{DisplayName: "Franse UIC-franc", Symbol: "XFU"},
				currency.XOF: cldr.Currency{DisplayName: "CFA-franc BCEAO", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "Palladium", Symbol: "XPD"},
				currency.XPF: cldr.Currency{DisplayName: "CFP-frank", Symbol: "XPF"},
				currency.XPT: cldr.Currency{DisplayName: "Platina", Symbol: "XPT"},
				currency.XRE: cldr.Currency{DisplayName: "RINET-fondsen", Symbol: "XRE"},
				currency.XSU: cldr.Currency{DisplayName: "Sucre", Symbol: "XSU"},
				currency.XTS: cldr.Currency{DisplayName: "Valutacode voor testdoeleinden", Symbol: "XTS"},
				currency.XUA: cldr.Currency{DisplayName: "ADB-rekeneenheid", Symbol: "XUA"},
				currency.XXX: cldr.Currency{DisplayName: "onbekende munteenheid", Symbol: "XXX"},
				currency.YDD: cldr.Currency{DisplayName: "Jemenitische dinar", Symbol: "YDD"},
				currency.YER: cldr.Currency{DisplayName: "Jemenitische rial", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "Joegoslavische harde dinar", Symbol: "YUD"},
				currency.YUM: cldr.Currency{DisplayName: "Joegoslavische noviy-dinar", Symbol: "YUM"},
				currency.YUN: cldr.Currency{DisplayName: "Joegoslavische convertibele dinar", Symbol: "YUN"},
				currency.YUR: cldr.Currency{DisplayName: "Joegoslavische hervormde dinar (1992–1993)", Symbol: "YUR"},
				currency.ZAL: cldr.Currency{DisplayName: "Zuid-Afrikaanse rand (financieel)", Symbol: "ZAL"},
				currency.ZAR: cldr.Currency{DisplayName: "Zuid-Afrikaanse rand", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Zambiaanse kwacha (1968–2012)", Symbol: "ZMK"},
				currency.ZMW: cldr.Currency{DisplayName: "Zambiaanse kwacha", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "Zaïrese nieuwe zaïre", Symbol: "ZRN"},
				currency.ZRZ: cldr.Currency{DisplayName: "Zaïrese zaïre", Symbol: "ZRZ"},
				currency.ZWD: cldr.Currency{DisplayName: "Zimbabwaanse dollar", Symbol: "ZWD"},
				currency.ZWL: cldr.Currency{DisplayName: "Zimbabwaanse dollar (2009)", Symbol: "ZWL"},
				currency.ZWR: cldr.Currency{DisplayName: "Zimbabwaanse dollar (2008)", Symbol: "ZWR"},
			},
		},
	}
}
