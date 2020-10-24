// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func Get_hr() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "hr",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d. MMMM y.", Long: "d. MMMM y.", Medium: "d. MMM y.", Short: "dd. MM. y."}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss (zzzz)", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'u' {0}", Long: "{1} 'u' {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "sij", Feb: "velj", Mar: "ožu", Apr: "tra", May: "svi", Jun: "lip", Jul: "srp", Aug: "kol", Sep: "ruj", Oct: "lis", Nov: "stu", Dec: "pro"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1.", Feb: "2.", Mar: "3.", Apr: "4.", May: "5.", Jun: "6.", Jul: "7.", Aug: "8.", Sep: "9.", Oct: "10.", Nov: "11.", Dec: "12."}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "siječanj", Feb: "veljača", Mar: "ožujak", Apr: "travanj", May: "svibanj", Jun: "lipanj", Jul: "srpanj", Aug: "kolovoz", Sep: "rujan", Oct: "listopad", Nov: "studeni", Dec: "prosinac"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ned", Mon: "pon", Tue: "uto", Wed: "sri", Thu: "čet", Fri: "pet", Sat: "sub"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "n", Mon: "p", Tue: "u", Wed: "s", Thu: "č", Fri: "p", Sat: "s"}, Short: cldr.CalendarDayFormatNameValue{Sun: "ned", Mon: "pon", Tue: "uto", Wed: "sri", Thu: "čet", Fri: "pet", Sat: "sub"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "nedjelja", Mon: "ponedjeljak", Tue: "utorak", Wed: "srijeda", Thu: "četvrtak", Fri: "petak", Sat: "subota"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "andorska pezeta", Symbol: "ADP"},
				currency.AED: cldr.Currency{DisplayName: "UAE dirham", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "afganistanski afgani (1927.–2002.)", Symbol: "AFA"},
				currency.AFN: cldr.Currency{DisplayName: "afganistanski afgani", Symbol: "AFN"},
				currency.ALK: cldr.Currency{DisplayName: "stari albanski lek", Symbol: "ALK"},
				currency.ALL: cldr.Currency{DisplayName: "albanski lek", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "armenski dram", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "nizozemskoantilski gulden", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "angolska kvanza", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "angolska kvanza (1977.–1990.)", Symbol: "AOK"},
				currency.AON: cldr.Currency{DisplayName: "angolska nova kvanza (1990.–2000.)", Symbol: "AON"},
				currency.AOR: cldr.Currency{DisplayName: "angolska kvanza (1995.–1999.)", Symbol: "AOR"},
				currency.ARA: cldr.Currency{DisplayName: "argentinski austral", Symbol: "ARA"},
				currency.ARL: cldr.Currency{DisplayName: "argentinski pezo lej (1970.–1983.)", Symbol: "ARL"},
				currency.ARM: cldr.Currency{DisplayName: "argentinski pezo (1881.–1970.)", Symbol: "ARM"},
				currency.ARP: cldr.Currency{DisplayName: "argentinski pezo (1983.–1985.)", Symbol: "ARP"},
				currency.ARS: cldr.Currency{DisplayName: "argentinski pezo", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "austrijski šiling", Symbol: "ATS"},
				currency.AUD: cldr.Currency{DisplayName: "australski dolar", Symbol: "AUD"},
				currency.AWG: cldr.Currency{DisplayName: "arupski florin", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "azerbajdžanski manat (1993.–2006.)", Symbol: "AZM"},
				currency.AZN: cldr.Currency{DisplayName: "azerbajdžanski manat", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "bosansko-hercegovački dinar", Symbol: "BAD"},
				currency.BAM: cldr.Currency{DisplayName: "konvertibilna marka", Symbol: "BAM"},
				currency.BAN: cldr.Currency{DisplayName: "bosansko-hercegovački novi dinar", Symbol: "BAN"},
				currency.BBD: cldr.Currency{DisplayName: "barbadoski dolar", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "bangladeška taka", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "belgijski franak (konvertibilan)", Symbol: "BEC"},
				currency.BEF: cldr.Currency{DisplayName: "belgijski franak", Symbol: "BEF"},
				currency.BEL: cldr.Currency{DisplayName: "belgijski franak (financijski)", Symbol: "BEL"},
				currency.BGL: cldr.Currency{DisplayName: "bugarski čvrsti lev", Symbol: "BGL"},
				currency.BGM: cldr.Currency{DisplayName: "bugarski socijalistički lev", Symbol: "BGM"},
				currency.BGN: cldr.Currency{DisplayName: "bugarski lev", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "stari bugarski lev", Symbol: "BGO"},
				currency.BHD: cldr.Currency{DisplayName: "bahreinski dinar", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "burundski franak", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "bermudski dolar", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "brunejski dolar", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "bolivijski bolivijano", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "stari bolivijski bolivijano", Symbol: "BOL"},
				currency.BOP: cldr.Currency{DisplayName: "bolivijski pezo", Symbol: "BOP"},
				currency.BOV: cldr.Currency{DisplayName: "bolivijski mvdol", Symbol: "BOV"},
				currency.BRB: cldr.Currency{DisplayName: "brazilski novi cruzeiro (1967.–1986.)", Symbol: "BRB"},
				currency.BRC: cldr.Currency{DisplayName: "brazilski cruzado", Symbol: "BRC"},
				currency.BRE: cldr.Currency{DisplayName: "brazilski cruzeiro (1990.–1993.)", Symbol: "BRE"},
				currency.BRL: cldr.Currency{DisplayName: "brazilski real", Symbol: "BRL"},
				currency.BRN: cldr.Currency{DisplayName: "brazilski novi cruzado", Symbol: "BRN"},
				currency.BRR: cldr.Currency{DisplayName: "brazilski cruzeiro", Symbol: "BRR"},
				currency.BRZ: cldr.Currency{DisplayName: "stari brazilski kruzeiro", Symbol: "BRZ"},
				currency.BSD: cldr.Currency{DisplayName: "bahamski dolar", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "butanski ngultrum", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "burmanski kyat", Symbol: "BUK"},
				currency.BWP: cldr.Currency{DisplayName: "bocvanska pula", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "bjeloruska nova rublja (1994–1999)", Symbol: "BYB"},
				currency.BYN: cldr.Currency{DisplayName: "bjeloruski rubalj", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "bjeloruska rublja (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "belizeanski dolar", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "kanadski dolar", Symbol: "CAD"},
				currency.CDF: cldr.Currency{DisplayName: "kongoanski franak", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "WIR euro", Symbol: "CHE"},
				currency.CHF: cldr.Currency{DisplayName: "švicarski franak", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "WIR franak", Symbol: "CHW"},
				currency.CLE: cldr.Currency{DisplayName: "čileanski eskudo", Symbol: "CLE"},
				currency.CLF: cldr.Currency{DisplayName: "čileanski unidades de fomentos", Symbol: "CLF"},
				currency.CLP: cldr.Currency{DisplayName: "čileanski pezo", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "kineski juan (offshore)", Symbol: "CNH"},
				currency.CNX: cldr.Currency{DisplayName: "kineski narodni dolar", Symbol: "CNX"},
				currency.CNY: cldr.Currency{DisplayName: "kineski yuan", Symbol: "CNY"},
				currency.COP: cldr.Currency{DisplayName: "kolumbijski pezo", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "unidad de valor real", Symbol: "COU"},
				currency.CRC: cldr.Currency{DisplayName: "kostarikanski kolon", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "stari srpski dinar", Symbol: "CSD"},
				currency.CSK: cldr.Currency{DisplayName: "čehoslovačka kruna", Symbol: "CSK"},
				currency.CUC: cldr.Currency{DisplayName: "kubanski konvertibilni pezo", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "kubanski pezo", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "zelenortski eskudo", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "ciparska funta", Symbol: "CYP"},
				currency.CZK: cldr.Currency{DisplayName: "češka kruna", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "istočnonjemačka marka", Symbol: "DDM"},
				currency.DEM: cldr.Currency{DisplayName: "njemačka marka", Symbol: "DEM"},
				currency.DJF: cldr.Currency{DisplayName: "džibutski franak", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "danska kruna", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "dominikanski pezo", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "alžirski dinar", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "ekvatorska sukra", Symbol: "ECS"},
				currency.ECV: cldr.Currency{DisplayName: "ekvatorski unidad de valor constante (UVC)", Symbol: "ECV"},
				currency.EEK: cldr.Currency{DisplayName: "estonska kruna", Symbol: "EEK"},
				currency.EGP: cldr.Currency{DisplayName: "egipatska funta", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "eritrejska nakfa", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "španjolska pezeta (A račun)", Symbol: "ESA"},
				currency.ESB: cldr.Currency{DisplayName: "španjolska pezeta (konvertibilni račun)", Symbol: "ESB"},
				currency.ESP: cldr.Currency{DisplayName: "španjolska pezeta", Symbol: "ESP"},
				currency.ETB: cldr.Currency{DisplayName: "etiopski bir", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "EUR"},
				currency.FIM: cldr.Currency{DisplayName: "finska marka", Symbol: "FIM"},
				currency.FJD: cldr.Currency{DisplayName: "fidžijski dolar", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "falklandska funta", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "francuski franak", Symbol: "FRF"},
				currency.GBP: cldr.Currency{DisplayName: "britanska funta", Symbol: "GBP"},
				currency.GEK: cldr.Currency{DisplayName: "gruzijski kupon larit", Symbol: "GEK"},
				currency.GEL: cldr.Currency{DisplayName: "gruzijski lari", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "ganski cedi (1979.–2007.)", Symbol: "GHC"},
				currency.GHS: cldr.Currency{DisplayName: "ganski cedi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "gibraltarska funta", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "gambijski dalas", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "gvinejski franak", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "gvinejski syli", Symbol: "GNS"},
				currency.GQE: cldr.Currency{DisplayName: "ekvatorski gvinejski ekwele", Symbol: "GQE"},
				currency.GRD: cldr.Currency{DisplayName: "grčka drahma", Symbol: "GRD"},
				currency.GTQ: cldr.Currency{DisplayName: "gvatemalski kvecal", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "portugalski gvinejski eskudo", Symbol: "GWE"},
				currency.GWP: cldr.Currency{DisplayName: "gvinejskobisauski pezo", Symbol: "GWP"},
				currency.GYD: cldr.Currency{DisplayName: "gvajanski dolar", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "hongkonški dolar", Symbol: "HKD"},
				currency.HNL: cldr.Currency{DisplayName: "honduraška lempira", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "hrvatski dinar", Symbol: "HRD"},
				currency.HRK: cldr.Currency{DisplayName: "hrvatska kuna", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "haićanski gourd", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "mađarska forinta", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "indonezijska rupija", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "irska funta", Symbol: "IEP"},
				currency.ILP: cldr.Currency{DisplayName: "izraelska funta", Symbol: "ILP"},
				currency.ILR: cldr.Currency{DisplayName: "stari izraelski šekel", Symbol: "ILR"},
				currency.ILS: cldr.Currency{DisplayName: "novi izraelski šekel", Symbol: "ILS"},
				currency.INR: cldr.Currency{DisplayName: "indijska rupija", Symbol: "INR"},
				currency.IQD: cldr.Currency{DisplayName: "irački dinar", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "iranski rijal", Symbol: "IRR"},
				currency.ISJ: cldr.Currency{DisplayName: "stara islandska kruna", Symbol: "ISJ"},
				currency.ISK: cldr.Currency{DisplayName: "islandska kruna", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "talijanska lira", Symbol: "ITL"},
				currency.JMD: cldr.Currency{DisplayName: "jamajčanski dolar", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "jordanski dinar", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "japanski jen", Symbol: "JPY"},
				currency.KES: cldr.Currency{DisplayName: "kenijski šiling", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "kirgiski som", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "kambodžanski rijal", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "komorski franak", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "sjevernokorejski won", Symbol: "KPW"},
				currency.KRH: cldr.Currency{DisplayName: "južnokorejski hvan", Symbol: "KRH"},
				currency.KRO: cldr.Currency{DisplayName: "stari južnokorejski von", Symbol: "KRO"},
				currency.KRW: cldr.Currency{DisplayName: "južnokorejski won", Symbol: "KRW"},
				currency.KWD: cldr.Currency{DisplayName: "kuvajtski dinar", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "kajmanski dolar", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "kazahstanski tenge", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "laoski kip", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "libanonska funta", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "šrilankanska rupija", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "liberijski dolar", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "lesoto loti", Symbol: "LSL"},
				currency.LTL: cldr.Currency{DisplayName: "litavski litas", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "litavski talonas", Symbol: "LTT"},
				currency.LUC: cldr.Currency{DisplayName: "luksemburški konvertibilni franak", Symbol: "LUC"},
				currency.LUF: cldr.Currency{DisplayName: "luksemburški franak", Symbol: "LUF"},
				currency.LUL: cldr.Currency{DisplayName: "luksemburški financijski franak", Symbol: "LUL"},
				currency.LVL: cldr.Currency{DisplayName: "letonski lats", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "letonska rublja", Symbol: "LVR"},
				currency.LYD: cldr.Currency{DisplayName: "libijski dinar", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "marokanski dirham", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "marokanski franak", Symbol: "MAF"},
				currency.MCF: cldr.Currency{DisplayName: "monegaški franak", Symbol: "MCF"},
				currency.MDC: cldr.Currency{DisplayName: "moldavski kupon", Symbol: "MDC"},
				currency.MDL: cldr.Currency{DisplayName: "moldavski lej", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "malgaški arijari", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "madagaskarski franak", Symbol: "MGF"},
				currency.MKD: cldr.Currency{DisplayName: "makedonski denar", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "stari makedonski denar", Symbol: "MKN"},
				currency.MLF: cldr.Currency{DisplayName: "malijski franak", Symbol: "MLF"},
				currency.MMK: cldr.Currency{DisplayName: "mjanmarski kjat", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "mongolski tugrik", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "makaoška pataka", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "mauritanijska ouguja (1973. – 2017.)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "mauritanijska ouguja", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "malteška lira", Symbol: "MTL"},
				currency.MTP: cldr.Currency{DisplayName: "malteška funta", Symbol: "MTP"},
				currency.MUR: cldr.Currency{DisplayName: "mauricijska rupija", Symbol: "MUR"},
				currency.MVP: cldr.Currency{DisplayName: "maldivijska rupija", Symbol: "MVP"},
				currency.MVR: cldr.Currency{DisplayName: "maldivijska rufija", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "malavijska kvača", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "meksički pezo", Symbol: "MXN"},
				currency.MXP: cldr.Currency{DisplayName: "meksički srebrni pezo (1861–1992)", Symbol: "MXP"},
				currency.MXV: cldr.Currency{DisplayName: "meksički unidad de inversion (UDI)", Symbol: "MXV"},
				currency.MYR: cldr.Currency{DisplayName: "malezijski ringit", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "mozambijski eskudo", Symbol: "MZE"},
				currency.MZM: cldr.Currency{DisplayName: "stari mozambijski metikal", Symbol: "MZM"},
				currency.MZN: cldr.Currency{DisplayName: "mozambički metikal", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "namibijski dolar", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "nigerijska naira", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "nikaragvanska kordoba", Symbol: "NIC"},
				currency.NIO: cldr.Currency{DisplayName: "nikaragvanska zlatna kordoba", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "nizozemski gulden", Symbol: "NLG"},
				currency.NOK: cldr.Currency{DisplayName: "norveška kruna", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "nepalska rupija", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "novozelandski dolar", Symbol: "NZD"},
				currency.OMR: cldr.Currency{DisplayName: "omanski rijal", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "panamska balboa", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "peruanski inti", Symbol: "PEI"},
				currency.PEN: cldr.Currency{DisplayName: "peruanski sol", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "peruanski sol (1863–1965)", Symbol: "PES"},
				currency.PGK: cldr.Currency{DisplayName: "kina Papue Nove Gvineje", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "filipinski pezo", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "pakistanska rupija", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "poljska zlota", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "poljska zlota (1950.–1995.)", Symbol: "PLZ"},
				currency.PTE: cldr.Currency{DisplayName: "portugalski eskudo", Symbol: "PTE"},
				currency.PYG: cldr.Currency{DisplayName: "paragvajski gvarani", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "katarski rial", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "rodezijski dolar", Symbol: "RHD"},
				currency.ROL: cldr.Currency{DisplayName: "starorumunjski lek", Symbol: "ROL"},
				currency.RON: cldr.Currency{DisplayName: "rumunjski lej", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "srpski dinar", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "ruski rubalj", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "ruska rublja (1991.–1998.)", Symbol: "RUR"},
				currency.RWF: cldr.Currency{DisplayName: "ruandski franak", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "saudijski rijal", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "solmonskootočni dolar", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "sejšelska rupija", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "sudanski dinar", Symbol: "SDD"},
				currency.SDG: cldr.Currency{DisplayName: "sudanska funta", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "stara sudanska funta", Symbol: "SDP"},
				currency.SEK: cldr.Currency{DisplayName: "švedska kruna", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "singapurski dolar", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "svetohelenska funta", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "slovenski tolar", Symbol: "SIT"},
				currency.SKK: cldr.Currency{DisplayName: "slovačka kruna", Symbol: "SKK"},
				currency.SLL: cldr.Currency{DisplayName: "sijeraleonski leone", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "somalijski šiling", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "surinamski dolar", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "surinamski gulden", Symbol: "SRG"},
				currency.SSP: cldr.Currency{DisplayName: "južnosudanska funta", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "dobra Svetog Tome i Principa (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "dobra Svetog Tome i Principa", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "sovjetska rublja", Symbol: "SUR"},
				currency.SVC: cldr.Currency{DisplayName: "salvadorski kolon", Symbol: "SVC"},
				currency.SYP: cldr.Currency{DisplayName: "sirijska funta", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "svazi lilangeni", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "tajlandski baht", Symbol: "THB"},
				currency.TJR: cldr.Currency{DisplayName: "tajikistanska rublja", Symbol: "TJR"},
				currency.TJS: cldr.Currency{DisplayName: "tadžikistanski somoni", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "turkmenistanski manat (1993.–2009.)", Symbol: "TMM"},
				currency.TMT: cldr.Currency{DisplayName: "turkmenistanski manat", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "tuniski dinar", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "tongaška pa’anga", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "timorski eskudo", Symbol: "TPE"},
				currency.TRL: cldr.Currency{DisplayName: "turska lira (1922.–2005.)", Symbol: "TRL"},
				currency.TRY: cldr.Currency{DisplayName: "turska lira", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "trininadtobaški dolar", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "novotajvanski dolar", Symbol: "TWD"},
				currency.TZS: cldr.Currency{DisplayName: "tanzanijski šiling", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "ukrajinska hrivnja", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "ukrajinski karbovanet", Symbol: "UAK"},
				currency.UGS: cldr.Currency{DisplayName: "ugandski šiling (1966.–1987.)", Symbol: "UGS"},
				currency.UGX: cldr.Currency{DisplayName: "ugandski šiling", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "američki dolar", Symbol: "USD"},
				currency.USN: cldr.Currency{DisplayName: "američki dolar (sljedeći dan)", Symbol: "USN"},
				currency.USS: cldr.Currency{DisplayName: "američki dolar (isti dan)", Symbol: "USS"},
				currency.UYI: cldr.Currency{DisplayName: "urugvajski pezo en unidades indexadas", Symbol: "UYI"},
				currency.UYP: cldr.Currency{DisplayName: "urugvajski pezo (1975.–1993.)", Symbol: "UYP"},
				currency.UYU: cldr.Currency{DisplayName: "urugvajski pezo", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "uzbekistanski som", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "venezuelanski bolivar (1871.–2008.)", Symbol: "VEB"},
				currency.VEF: cldr.Currency{DisplayName: "venezuelanski bolivar (2008. – 2018.)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "venezuelanski bolivar", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "vijetnamski dong", Symbol: "VND"},
				currency.VNN: cldr.Currency{DisplayName: "vijetnamski dong (1978.–1985.)", Symbol: "VNN"},
				currency.VUV: cldr.Currency{DisplayName: "vanuatski vatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "samoanska tala", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "CFA franak BEAC", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "srebro", Symbol: "XAG"},
				currency.XAU: cldr.Currency{DisplayName: "zlato", Symbol: "XAU"},
				currency.XBA: cldr.Currency{DisplayName: "Europska složena jedinica", Symbol: "XBA"},
				currency.XBB: cldr.Currency{DisplayName: "Europska monetarna jedinica", Symbol: "XBB"},
				currency.XBC: cldr.Currency{DisplayName: "europska obračunska jedinica (XBC)", Symbol: "XBC"},
				currency.XBD: cldr.Currency{DisplayName: "europska obračunska jedinica (XBD)", Symbol: "XBD"},
				currency.XCD: cldr.Currency{DisplayName: "istočnokaripski dolar", Symbol: "XCD"},
				currency.XDR: cldr.Currency{DisplayName: "posebna crtaća prava", Symbol: "XDR"},
				currency.XEU: cldr.Currency{DisplayName: "europska monetarna jedinica (ECU)", Symbol: "XEU"},
				currency.XFO: cldr.Currency{DisplayName: "francuski zlatni franak", Symbol: "XFO"},
				currency.XFU: cldr.Currency{DisplayName: "francuski UIC-franak", Symbol: "XFU"},
				currency.XOF: cldr.Currency{DisplayName: "CFA franak BCEAO", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "paladij", Symbol: "XPD"},
				currency.XPF: cldr.Currency{DisplayName: "CFP franak", Symbol: "XPF"},
				currency.XPT: cldr.Currency{DisplayName: "platina", Symbol: "XPT"},
				currency.XRE: cldr.Currency{DisplayName: "RINET fondovi", Symbol: "XRE"},
				currency.XSU: cldr.Currency{DisplayName: "sukre", Symbol: "XSU"},
				currency.XTS: cldr.Currency{DisplayName: "ispitni kod valute", Symbol: "XTS"},
				currency.XUA: cldr.Currency{DisplayName: "obračunska jedinica ADB", Symbol: "XUA"},
				currency.XXX: cldr.Currency{DisplayName: "nepoznata valuta", Symbol: "XXX"},
				currency.YDD: cldr.Currency{DisplayName: "jemenski dinar", Symbol: "YDD"},
				currency.YER: cldr.Currency{DisplayName: "jemenski rijal", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "jugoslavenski čvrsti dinar", Symbol: "YUD"},
				currency.YUM: cldr.Currency{DisplayName: "jugoslavenski novi dinar", Symbol: "YUM"},
				currency.YUN: cldr.Currency{DisplayName: "jugoslavenski konvertibilni dinar", Symbol: "YUN"},
				currency.YUR: cldr.Currency{DisplayName: "jugoslavenski reformirani dinar", Symbol: "YUR"},
				currency.ZAL: cldr.Currency{DisplayName: "južnoafrički rand (financijski)", Symbol: "ZAL"},
				currency.ZAR: cldr.Currency{DisplayName: "južnoafrički rand", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "zambijska kvača (1968–2012)", Symbol: "ZMK"},
				currency.ZMW: cldr.Currency{DisplayName: "zambijska kvača", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "zairski novi zair", Symbol: "ZRN"},
				currency.ZRZ: cldr.Currency{DisplayName: "zairski zair", Symbol: "ZRZ"},
				currency.ZWD: cldr.Currency{DisplayName: "zimbabveanski dolar (1980.–2008.)", Symbol: "ZWD"},
				currency.ZWL: cldr.Currency{DisplayName: "zimbabveanski dolar (2009)", Symbol: "ZWL"},
				currency.ZWR: cldr.Currency{DisplayName: "zimbabveanski dolar (2008)", Symbol: "ZWR"},
			},
		},
	}
}