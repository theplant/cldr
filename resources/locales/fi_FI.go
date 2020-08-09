// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_fi_FI() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "cccc d. MMMM y", Long: "d. MMMM y", Medium: "d.M.y", Short: "d.M.y"}, Time: cldr.CalendarDateFormat{Full: "H.mm.ss zzzz", Long: "H.mm.ss z", Medium: "H.mm.ss", Short: "H.mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'klo' {0}", Long: "{1} 'klo' {0}", Medium: "{1} 'klo' {0}", Short: "{1} {0}"}, GMT: "UTC{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "tammi", Feb: "helmi", Mar: "maalis", Apr: "huhti", May: "touko", Jun: "kesä", Jul: "heinä", Aug: "elo", Sep: "syys", Oct: "loka", Nov: "marras", Dec: "joulu"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "T", Feb: "H", Mar: "M", Apr: "H", May: "T", Jun: "K", Jul: "H", Aug: "E", Sep: "S", Oct: "L", Nov: "M", Dec: "J"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "tammikuu", Feb: "helmikuu", Mar: "maaliskuu", Apr: "huhtikuu", May: "toukokuu", Jun: "kesäkuu", Jul: "heinäkuu", Aug: "elokuu", Sep: "syyskuu", Oct: "lokakuu", Nov: "marraskuu", Dec: "joulukuu"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "su", Mon: "ma", Tue: "ti", Wed: "ke", Thu: "to", Fri: "pe", Sat: "la"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "K", Thu: "T", Fri: "P", Sat: "L"}, Short: cldr.CalendarDayFormatNameValue{Sun: "su", Mon: "ma", Tue: "ti", Wed: "ke", Thu: "to", Fri: "pe", Sat: "la"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "sunnuntai", Mon: "maanantai", Tue: "tiistai", Wed: "keskiviikko", Thu: "torstai", Fri: "perjantai", Sat: "lauantai"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "ap.", PM: "ip."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "ap.", PM: "ip."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "ap.", PM: "ip."}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_fi]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "−", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "Andorran peseta", Symbol: "ADP"},
				currency.AED: cldr.Currency{DisplayName: "Arabiemiirikuntien dirhami", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "Afganistanin afgaani (1927–2002)", Symbol: "AFA"},
				currency.AFN: cldr.Currency{DisplayName: "Afganistanin afgaani", Symbol: "AFN"},
				currency.ALK: cldr.Currency{DisplayName: "Albanian lek (1946–1965)", Symbol: "ALK"},
				currency.ALL: cldr.Currency{DisplayName: "Albanian lek", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Armenian dram", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Alankomaiden Antillien guldeni", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Angolan kwanza", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "Angolan kwanza (1977–1991)", Symbol: "AOK"},
				currency.AON: cldr.Currency{DisplayName: "Angolan uusi kwanza (1990–2000)", Symbol: "AON"},
				currency.AOR: cldr.Currency{DisplayName: "Angolan kwanza reajustado (1995–1999)", Symbol: "AOR"},
				currency.ARA: cldr.Currency{DisplayName: "Argentiinan austral", Symbol: "ARA"},
				currency.ARL: cldr.Currency{DisplayName: "Argentiinan ley-peso (1970–1983)", Symbol: "ARL"},
				currency.ARM: cldr.Currency{DisplayName: "Argentiinan peso (1881–1970)", Symbol: "ARM"},
				currency.ARP: cldr.Currency{DisplayName: "Argentiinan peso (1983–1985)", Symbol: "ARP"},
				currency.ARS: cldr.Currency{DisplayName: "Argentiinan peso", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "Itävallan šillinki", Symbol: "ATS"},
				currency.AUD: cldr.Currency{DisplayName: "Australian dollari", Symbol: "AUD"},
				currency.AWG: cldr.Currency{DisplayName: "Aruban floriini", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "Azerbaidžanin manat (1993–2006)", Symbol: "AZM"},
				currency.AZN: cldr.Currency{DisplayName: "Azerbaidžanin manat", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "Bosnia-Hertsegovinan dinaari (1992–1994)", Symbol: "BAD"},
				currency.BAM: cldr.Currency{DisplayName: "Bosnia-Hertsegovinan vaihdettava markka", Symbol: "BAM"},
				currency.BAN: cldr.Currency{DisplayName: "Bosnia-Hertsegovinan uusi dinaari (1994–1997)", Symbol: "BAN"},
				currency.BBD: cldr.Currency{DisplayName: "Barbadosin dollari", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Bangladeshin taka", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "Belgian vaihdettava frangi", Symbol: "BEC"},
				currency.BEF: cldr.Currency{DisplayName: "Belgian frangi", Symbol: "BEF"},
				currency.BEL: cldr.Currency{DisplayName: "Belgian rahoitusfrangi", Symbol: "BEL"},
				currency.BGL: cldr.Currency{DisplayName: "Bulgarian kova lev", Symbol: "BGL"},
				currency.BGM: cldr.Currency{DisplayName: "Bulgarian sosialistinen lev", Symbol: "BGM"},
				currency.BGN: cldr.Currency{DisplayName: "Bulgarian lev", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "Bulgarian lev (1879–1952)", Symbol: "BGO"},
				currency.BHD: cldr.Currency{DisplayName: "Bahrainin dinaari", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Burundin frangi", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Bermudan dollari", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Brunein dollari", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Bolivian boliviano", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "Bolivian boliviano (1863–1963)", Symbol: "BOL"},
				currency.BOP: cldr.Currency{DisplayName: "Bolivian peso", Symbol: "BOP"},
				currency.BOV: cldr.Currency{DisplayName: "Bolivian mvdol", Symbol: "BOV"},
				currency.BRB: cldr.Currency{DisplayName: "Brasilian uusi cruzeiro (1967–1986)", Symbol: "BRB"},
				currency.BRC: cldr.Currency{DisplayName: "Brasilian cruzado (1986–1989)", Symbol: "BRC"},
				currency.BRE: cldr.Currency{DisplayName: "Brasilian cruzeiro (1990–1993)", Symbol: "BRE"},
				currency.BRL: cldr.Currency{DisplayName: "Brasilian real", Symbol: "BRL"},
				currency.BRN: cldr.Currency{DisplayName: "Brasilian uusi cruzado (1989–1990)", Symbol: "BRN"},
				currency.BRR: cldr.Currency{DisplayName: "Brasilian cruzeiro (1993–1994)", Symbol: "BRR"},
				currency.BRZ: cldr.Currency{DisplayName: "Brasilian cruzeiro (1942–1967)", Symbol: "BRZ"},
				currency.BSD: cldr.Currency{DisplayName: "Bahaman dollari", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Bhutanin ngultrum", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "Burman kyat", Symbol: "BUK"},
				currency.BWP: cldr.Currency{DisplayName: "Botswanan pula", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "Valko-Venäjän uusi rupla (1994–1999)", Symbol: "BYB"},
				currency.BYN: cldr.Currency{DisplayName: "Valko-Venäjän rupla", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Valko-Venäjän rupla (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "Belizen dollari", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Kanadan dollari", Symbol: "CAD"},
				currency.CDF: cldr.Currency{DisplayName: "Kongon frangi", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "Sveitsin WIR-euro", Symbol: "CHE"},
				currency.CHF: cldr.Currency{DisplayName: "Sveitsin frangi", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "Sveitsin WIR-frangi", Symbol: "CHW"},
				currency.CLE: cldr.Currency{DisplayName: "Chilen escudo", Symbol: "CLE"},
				currency.CLF: cldr.Currency{DisplayName: "Chilen unidades de fomento", Symbol: "CLF"},
				currency.CLP: cldr.Currency{DisplayName: "Chilen peso", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Kiinan juan (offshore)", Symbol: "CNH"},
				currency.CNX: cldr.Currency{DisplayName: "Kiinan kansanpankin dollari", Symbol: "CNX"},
				currency.CNY: cldr.Currency{DisplayName: "Kiinan juan", Symbol: "CNY"},
				currency.COP: cldr.Currency{DisplayName: "Kolumbian peso", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "Kolumbian unidad de valor real", Symbol: "COU"},
				currency.CRC: cldr.Currency{DisplayName: "Costa Rican colón", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "Serbian dinaari (2002–2006)", Symbol: "CSD"},
				currency.CSK: cldr.Currency{DisplayName: "Tšekkoslovakian kova koruna", Symbol: "CSK"},
				currency.CUC: cldr.Currency{DisplayName: "Kuuban vaihdettava peso", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Kuuban peso", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Kap Verden escudo", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "Kyproksen punta", Symbol: "CYP"},
				currency.CZK: cldr.Currency{DisplayName: "Tšekin koruna", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "Itä-Saksan markka", Symbol: "DDM"},
				currency.DEM: cldr.Currency{DisplayName: "Saksan markka", Symbol: "DEM"},
				currency.DJF: cldr.Currency{DisplayName: "Djiboutin frangi", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Tanskan kruunu", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Dominikaanisen tasavallan peso", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Algerian dinaari", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "Ecuadorin sucre", Symbol: "ECS"},
				currency.ECV: cldr.Currency{DisplayName: "Ecuadorin UVC", Symbol: "ECV"},
				currency.EEK: cldr.Currency{DisplayName: "Viron kruunu", Symbol: "EEK"},
				currency.EGP: cldr.Currency{DisplayName: "Egyptin punta", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Eritrean nakfa", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "Espanjan peseta (A-tili)", Symbol: "ESA"},
				currency.ESB: cldr.Currency{DisplayName: "Espanjan peseta (vaihdettava tili)", Symbol: "ESB"},
				currency.ESP: cldr.Currency{DisplayName: "Espanjan peseta", Symbol: "ESP"},
				currency.ETB: cldr.Currency{DisplayName: "Etiopian birr", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Suomen markka", Symbol: "mk"},
				currency.FJD: cldr.Currency{DisplayName: "Fidžin dollari", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Falklandinsaarten punta", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "Ranskan frangi", Symbol: "FRF"},
				currency.GBP: cldr.Currency{DisplayName: "Englannin punta", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "Georgian kuponkilari", Symbol: "GEK"},
				currency.GEL: cldr.Currency{DisplayName: "Georgian lari", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "Ghanan cedi (1979–2007)", Symbol: "GHC"},
				currency.GHS: cldr.Currency{DisplayName: "Ghanan cedi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Gibraltarin punta", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Gambian dalasi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Guinean frangi", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "Guinean syli", Symbol: "GNS"},
				currency.GQE: cldr.Currency{DisplayName: "Päiväntasaajan Guinean ekwele", Symbol: "GQE"},
				currency.GRD: cldr.Currency{DisplayName: "Kreikan drakma", Symbol: "GRD"},
				currency.GTQ: cldr.Currency{DisplayName: "Guatemalan quetzal", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "Portugalin Guinean escudo", Symbol: "GWE"},
				currency.GWP: cldr.Currency{DisplayName: "Guinea-Bissaun peso", Symbol: "GWP"},
				currency.GYD: cldr.Currency{DisplayName: "Guyanan dollari", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Hongkongin dollari", Symbol: "HKD"},
				currency.HNL: cldr.Currency{DisplayName: "Hondurasin lempira", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "Kroatian dinaari", Symbol: "HRD"},
				currency.HRK: cldr.Currency{DisplayName: "Kroatian kuna", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Haitin gourde", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Unkarin forintti", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Indonesian rupia", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "Irlannin punta", Symbol: "IEP"},
				currency.ILP: cldr.Currency{DisplayName: "Israelin punta", Symbol: "ILP"},
				currency.ILR: cldr.Currency{DisplayName: "Israelin sekeli (1980–1985)", Symbol: "ILR"},
				currency.ILS: cldr.Currency{DisplayName: "Israelin uusi sekeli", Symbol: "ILS"},
				currency.INR: cldr.Currency{DisplayName: "Intian rupia", Symbol: "INR"},
				currency.IQD: cldr.Currency{DisplayName: "Irakin dinaari", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Iranin rial", Symbol: "IRR"},
				currency.ISJ: cldr.Currency{DisplayName: "Islannin kruunu (1918–1981)", Symbol: "ISJ"},
				currency.ISK: cldr.Currency{DisplayName: "Islannin kruunu", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "Italian liira", Symbol: "ITL"},
				currency.JMD: cldr.Currency{DisplayName: "Jamaikan dollari", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Jordanian dinaari", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Japanin jeni", Symbol: "¥"},
				currency.KES: cldr.Currency{DisplayName: "Kenian šillinki", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Kirgisian som", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Kambodžan riel", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Komorien frangi", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Pohjois-Korean won", Symbol: "KPW"},
				currency.KRH: cldr.Currency{DisplayName: "Etelä-Korean hwan (1953–1962)", Symbol: "KRH"},
				currency.KRO: cldr.Currency{DisplayName: "Etelä-Korean won (1945–1953)", Symbol: "KRO"},
				currency.KRW: cldr.Currency{DisplayName: "Etelä-Korean won", Symbol: "KRW"},
				currency.KWD: cldr.Currency{DisplayName: "Kuwaitin dinaari", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Caymansaarten dollari", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Kazakstanin tenge", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Laosin kip", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Libanonin punta", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Sri Lankan rupia", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Liberian dollari", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Lesothon loti", Symbol: "LSL"},
				currency.LTL: cldr.Currency{DisplayName: "Liettuan liti", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "Liettuan talonas", Symbol: "LTT"},
				currency.LUC: cldr.Currency{DisplayName: "Luxemburgin vaihdettava frangi", Symbol: "LUC"},
				currency.LUF: cldr.Currency{DisplayName: "Luxemburgin frangi", Symbol: "LUF"},
				currency.LUL: cldr.Currency{DisplayName: "Luxemburgin rahoitusfrangi", Symbol: "LUL"},
				currency.LVL: cldr.Currency{DisplayName: "Latvian lati", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "Latvian rupla", Symbol: "LVR"},
				currency.LYD: cldr.Currency{DisplayName: "Libyan dinaari", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Marokon dirhami", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "Marokon frangi", Symbol: "MAF"},
				currency.MCF: cldr.Currency{DisplayName: "Monacon frangi", Symbol: "MCF"},
				currency.MDC: cldr.Currency{DisplayName: "Moldovan kuponkileu", Symbol: "MDC"},
				currency.MDL: cldr.Currency{DisplayName: "Moldovan leu", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Madagaskarin ariary", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "Madagaskarin frangi", Symbol: "MGF"},
				currency.MKD: cldr.Currency{DisplayName: "Makedonian denaari", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "Makedonian dinaari (1992–1993)", Symbol: "MKN"},
				currency.MLF: cldr.Currency{DisplayName: "Malin frangi", Symbol: "MLF"},
				currency.MMK: cldr.Currency{DisplayName: "Myanmarin kyat", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Mongolian tugrik", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Macaon pataca", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Mauritanian ouguiya (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Mauritanian ouguiya", Symbol: ""},
				currency.MTL: cldr.Currency{DisplayName: "Maltan liira", Symbol: "MTL"},
				currency.MTP: cldr.Currency{DisplayName: "Maltan punta", Symbol: "MTP"},
				currency.MUR: cldr.Currency{DisplayName: "Mauritiuksen rupia", Symbol: "MUR"},
				currency.MVP: cldr.Currency{DisplayName: "Malediivien rupia (1947–1981)", Symbol: "MVP"},
				currency.MVR: cldr.Currency{DisplayName: "Malediivien rufiyaa", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Malawin kwacha", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Meksikon peso", Symbol: "MXN"},
				currency.MXP: cldr.Currency{DisplayName: "Meksikon hopeapeso (1861–1992)", Symbol: "MXP"},
				currency.MXV: cldr.Currency{DisplayName: "Meksikon UDI", Symbol: "MXV"},
				currency.MYR: cldr.Currency{DisplayName: "Malesian ringgit", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "Mosambikin escudo", Symbol: "MZE"},
				currency.MZM: cldr.Currency{DisplayName: "Mosambikin metical (1980–2006)", Symbol: "MZM"},
				currency.MZN: cldr.Currency{DisplayName: "Mosambikin metical", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Namibian dollari", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Nigerian naira", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "Nicaraguan córdoba (1988–1991)", Symbol: "NIC"},
				currency.NIO: cldr.Currency{DisplayName: "Nicaraguan córdoba", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "Alankomaiden guldeni", Symbol: "NLG"},
				currency.NOK: cldr.Currency{DisplayName: "Norjan kruunu", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Nepalin rupia", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Uuden-Seelannin dollari", Symbol: "NZD"},
				currency.OMR: cldr.Currency{DisplayName: "Omanin rial", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Panaman balboa", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "Perun inti", Symbol: "PEI"},
				currency.PEN: cldr.Currency{DisplayName: "Perun sol", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "Perun sol (1863–1965)", Symbol: "PES"},
				currency.PGK: cldr.Currency{DisplayName: "Papua-Uuden-Guinean kina", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Filippiinien peso", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Pakistanin rupia", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Puolan złoty", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "Puolan złoty (1950–1995)", Symbol: "PLZ"},
				currency.PTE: cldr.Currency{DisplayName: "Portugalin escudo", Symbol: "PTE"},
				currency.PYG: cldr.Currency{DisplayName: "Paraguayn guarani", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Qatarin rial", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "Rhodesian dollari", Symbol: "RHD"},
				currency.ROL: cldr.Currency{DisplayName: "Romanian leu (1952–2006)", Symbol: "ROL"},
				currency.RON: cldr.Currency{DisplayName: "Romanian leu", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Serbian dinaari", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Venäjän rupla", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "Venäjän rupla (1991–1998)", Symbol: "RUR"},
				currency.RWF: cldr.Currency{DisplayName: "Ruandan frangi", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Saudi-Arabian rial", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Salomonsaarten dollari", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Seychellien rupia", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "Sudanin dinaari (1992–2007)", Symbol: "SDD"},
				currency.SDG: cldr.Currency{DisplayName: "Sudanin punta", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "Sudanin punta (1957–1998)", Symbol: "SDP"},
				currency.SEK: cldr.Currency{DisplayName: "Ruotsin kruunu", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Singaporen dollari", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Saint Helenan punta", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "Slovenian tolar", Symbol: "SIT"},
				currency.SKK: cldr.Currency{DisplayName: "Slovakian koruna", Symbol: "SKK"},
				currency.SLL: cldr.Currency{DisplayName: "Sierra Leonen leone", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Somalian šillinki", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Surinamen dollari", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "Surinamen guldeni", Symbol: "SRG"},
				currency.SSP: cldr.Currency{DisplayName: "Etelä-Sudanin punta", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "São Tomén ja Príncipen dobra (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "São Tomén ja Príncipen dobra", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "Neuvostoliiton rupla", Symbol: "SUR"},
				currency.SVC: cldr.Currency{DisplayName: "El Salvadorin colón", Symbol: "SVC"},
				currency.SYP: cldr.Currency{DisplayName: "Syyrian punta", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Swazimaan lilangeni", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Thaimaan baht", Symbol: "THB"},
				currency.TJR: cldr.Currency{DisplayName: "Tadžikistanin rupla", Symbol: "TJR"},
				currency.TJS: cldr.Currency{DisplayName: "Tadžikistanin somoni", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "Turkmenistanin manat (1993–2009)", Symbol: "TMM"},
				currency.TMT: cldr.Currency{DisplayName: "Turkmenistanin manat", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Tunisian dinaari", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Tongan pa’anga", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "Timorin escudo", Symbol: "TPE"},
				currency.TRL: cldr.Currency{DisplayName: "Turkin liira (1922–2005)", Symbol: "TRL"},
				currency.TRY: cldr.Currency{DisplayName: "Turkin liira", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Trinidadin ja Tobagon dollari", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Taiwanin uusi dollari", Symbol: "TWD"},
				currency.TZS: cldr.Currency{DisplayName: "Tansanian šillinki", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Ukrainan hryvnia", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "Ukrainan karbovanetz", Symbol: "UAK"},
				currency.UGS: cldr.Currency{DisplayName: "Ugandan šillinki (1966–1987)", Symbol: "UGS"},
				currency.UGX: cldr.Currency{DisplayName: "Ugandan šillinki", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "Yhdysvaltain dollari", Symbol: "$"},
				currency.USN: cldr.Currency{DisplayName: "Yhdysvaltain dollari (seuraava päivä)", Symbol: "USN"},
				currency.USS: cldr.Currency{DisplayName: "Yhdysvaltain dollari (sama päivä)", Symbol: "USS"},
				currency.UYI: cldr.Currency{DisplayName: "Uruguayn peso en unidades indexadas", Symbol: "UYI"},
				currency.UYP: cldr.Currency{DisplayName: "Uruguayn peso (1975–1993)", Symbol: "UYP"},
				currency.UYU: cldr.Currency{DisplayName: "Uruguayn peso", Symbol: "UYU"},
				currency.UYW: cldr.Currency{DisplayName: "Uruguayn nimellinen palkkaindeksiyksikkö", Symbol: ""},
				currency.UZS: cldr.Currency{DisplayName: "Uzbekistanin som", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "Venezuelan bolívar (1871–2008)", Symbol: "VEB"},
				currency.VEF: cldr.Currency{DisplayName: "Venezuelan bolívar (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Venezuelan suvereeni bolívar", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Vietnamin dong", Symbol: "VND"},
				currency.VNN: cldr.Currency{DisplayName: "Vietnamin dong (1978–1985)", Symbol: "VNN"},
				currency.VUV: cldr.Currency{DisplayName: "Vanuatun vatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Samoan tala", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "CFA-frangi BEAC", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "hopea", Symbol: "XAG"},
				currency.XAU: cldr.Currency{DisplayName: "kulta", Symbol: "XAU"},
				currency.XBA: cldr.Currency{DisplayName: "EURCO", Symbol: "XBA"},
				currency.XBB: cldr.Currency{DisplayName: "Euroopan rahayksikkö (EMU)", Symbol: "XBB"},
				currency.XBC: cldr.Currency{DisplayName: "EUA (XBC)", Symbol: "XBC"},
				currency.XBD: cldr.Currency{DisplayName: "EUA (XBD)", Symbol: "XBD"},
				currency.XCD: cldr.Currency{DisplayName: "Itä-Karibian dollari", Symbol: "XCD"},
				currency.XDR: cldr.Currency{DisplayName: "erityisnosto-oikeus (SDR)", Symbol: "XDR"},
				currency.XEU: cldr.Currency{DisplayName: "Euroopan valuuttayksikkö (ECU)", Symbol: "XEU"},
				currency.XFO: cldr.Currency{DisplayName: "Ranskan kultafrangi", Symbol: "XFO"},
				currency.XFU: cldr.Currency{DisplayName: "Ranskan UIC-frangi", Symbol: "XFU"},
				currency.XOF: cldr.Currency{DisplayName: "CFA-frangi BCEAO", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "palladium", Symbol: "XPD"},
				currency.XPF: cldr.Currency{DisplayName: "CFP-frangi", Symbol: "XPF"},
				currency.XPT: cldr.Currency{DisplayName: "platina", Symbol: "XPT"},
				currency.XRE: cldr.Currency{DisplayName: "RINET-rahastot", Symbol: "XRE"},
				currency.XSU: cldr.Currency{DisplayName: "etelä-amerikkalaisen ALBA:n laskentayksikkö sucre", Symbol: "XSU"},
				currency.XTS: cldr.Currency{DisplayName: "testaustarkoitukseen varattu valuuttakoodi", Symbol: "XTS"},
				currency.XUA: cldr.Currency{DisplayName: "afrikkalainen AfDB-laskentayksikkö", Symbol: "XUA"},
				currency.XXX: cldr.Currency{DisplayName: "tuntematon rahayksikkö", Symbol: "XXX"},
				currency.YDD: cldr.Currency{DisplayName: "Jemenin dinaari", Symbol: "YDD"},
				currency.YER: cldr.Currency{DisplayName: "Jemenin rial", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "Jugoslavian kova dinaari (1966–1990)", Symbol: "YUD"},
				currency.YUM: cldr.Currency{DisplayName: "Jugoslavian uusi dinaari (1994–2002)", Symbol: "YUM"},
				currency.YUN: cldr.Currency{DisplayName: "Jugoslavian vaihdettava dinaari (1990–1992)", Symbol: "YUN"},
				currency.YUR: cldr.Currency{DisplayName: "Jugoslavian uudistettu dinaari (1992–1993)", Symbol: "YUR"},
				currency.ZAL: cldr.Currency{DisplayName: "Etelä-Afrikan rahoitusrandi", Symbol: "ZAL"},
				currency.ZAR: cldr.Currency{DisplayName: "Etelä-Afrikan randi", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Sambian kwacha (1968–2012)", Symbol: "ZMK"},
				currency.ZMW: cldr.Currency{DisplayName: "Sambian kwacha", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "Zairen uusi zaire (1993–1998)", Symbol: "ZRN"},
				currency.ZRZ: cldr.Currency{DisplayName: "Zairen zaire (1971–1993)", Symbol: "ZRZ"},
				currency.ZWD: cldr.Currency{DisplayName: "Zimbabwen dollari (1980–2008)", Symbol: "ZWD"},
				currency.ZWL: cldr.Currency{DisplayName: "Zimbabwen dollari (2009)", Symbol: "ZWL"},
				currency.ZWR: cldr.Currency{DisplayName: "Zimbabwen dollari (2008)", Symbol: "ZWR"},
			},
		},
	}
}
