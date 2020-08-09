// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_sr_Latn() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, dd. MMMM y.", Long: "dd. MMMM y.", Medium: "dd.MM.y.", Short: "d.M.yy."}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan", Feb: "feb", Mar: "mar", Apr: "apr", May: "maj", Jun: "jun", Jul: "jul", Aug: "avg", Sep: "sep", Oct: "okt", Nov: "nov", Dec: "dec"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "j", Feb: "f", Mar: "m", Apr: "a", May: "m", Jun: "j", Jul: "j", Aug: "a", Sep: "s", Oct: "o", Nov: "n", Dec: "d"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "januar", Feb: "februar", Mar: "mart", Apr: "april", May: "maj", Jun: "jun", Jul: "jul", Aug: "avgust", Sep: "septembar", Oct: "oktobar", Nov: "novembar", Dec: "decembar"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ned", Mon: "pon", Tue: "uto", Wed: "sre", Thu: "čet", Fri: "pet", Sat: "sub"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "n", Mon: "p", Tue: "u", Wed: "s", Thu: "č", Fri: "p", Sat: "s"}, Short: cldr.CalendarDayFormatNameValue{Sun: "ne", Mon: "po", Tue: "ut", Wed: "sr", Thu: "če", Fri: "pe", Sat: "su"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "nedelja", Mon: "ponedeljak", Tue: "utorak", Wed: "sreda", Thu: "četvrtak", Fri: "petak", Sat: "subota"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "pre podne", PM: "po podne"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "pre podne", PM: "po podne"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "pre podne", PM: "po podne"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_root]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "Andorska pezeta", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "UAE dirham", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "Avganistanski avgani (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Avganistanski avgani", Symbol: "AFN"},
				currency.ALK: cldr.Currency{DisplayName: "stari albanski lek", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "Albanski lek", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Jermenski dram", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Holandskoantilski gulden", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Angolska kvanza", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "Angolijska kvanza (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "Angolijska nova kvanza (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "Angolijska kvanza reađustado (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "Argentinski austral", Symbol: ""},
				currency.ARL: cldr.Currency{DisplayName: "Argentinski pezos lej", Symbol: ""},
				currency.ARM: cldr.Currency{DisplayName: "Argentinski pezos monedo nacional", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "Argentinski pezo (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "Argentinski pezos", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "Austrijski šiling", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "Australijski dolar", Symbol: "AUD"},
				currency.AWG: cldr.Currency{DisplayName: "Arubanski florin", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "Azerbejdžanski manat (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Azerbejdžanski manat", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "Bosansko-Hercegovački dinar", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Bosansko-hercegovačka konvertibilna marka", Symbol: "KM"},
				currency.BAN: cldr.Currency{DisplayName: "Bosansko-hercegovački novi dinar", Symbol: ""},
				currency.BBD: cldr.Currency{DisplayName: "Barbadoški dolar", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Bangladeška taka", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "Belgijski franak (konvertibilni)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "Belgijski franak", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "Belgijski franak (finansijski)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "Bugarski tvrdi lev", Symbol: ""},
				currency.BGM: cldr.Currency{DisplayName: "Bugarski socijalistički lev", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "Bugarski lev", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "Stari bugarski lev", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Bahreinski dinar", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Burundski franak", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Bermudski dolar", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Brunejski dolar", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Bolivijski bolivijano", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "Stari bolivijski bolivijano", Symbol: ""},
				currency.BOP: cldr.Currency{DisplayName: "Bolivijski pezo", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "Bolivijski mvdol", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "Brazilski novi kruzeiro (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "Brazilijski kruzado", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "Brazilski kruzeiro (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "Brazilski real", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "Brazilijski novi kruzado", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "Brazilski kruzeiro", Symbol: ""},
				currency.BRZ: cldr.Currency{DisplayName: "Stari brazilski kruzeiro", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "Bahamski dolar", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Butanski ngultrum", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "Burmanski kjat", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Bocvanska pula", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "Beloruska nova rublja (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "Beloruska rublja", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Beloruska rublja (2000–2016)", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "Beliski dolar", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Kanadski dolar", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Kongoanski franak", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "WIR evro", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Švajcarski franak", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "WIR franak", Symbol: ""},
				currency.CLE: cldr.Currency{DisplayName: "Čileanski eskudo", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "Čileovski unidades se fomento", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Čileanski pezos", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Kineski juan (ostrvski)", Symbol: "CNH"},
				currency.CNX: cldr.Currency{DisplayName: "Dolar kineske narodne banke", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "Kineski juan", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Kolumbijski pezos", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "Unidad de valorški real", Symbol: ""},
				currency.CRC: cldr.Currency{DisplayName: "Kostarikanski kolon", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "Stari srpski dinar", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "Čehoslovačka tvrda kruna", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "Kubanski konvertibilni pezos", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Kubanski pezos", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Zelenortski eskudo", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "Kiparska funta", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Češka kruna", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "Istočno-nemačka marka", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "Nemačka marka", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "Džibutski franak", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Danska kruna", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Dominikanski pezos", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Alžirski dinar", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "Ekvadorski sakr", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "Ekvadorski unidad de valor konstante", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "Estonska kroon", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Egipatska funta", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Eritrejska nakfa", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "Španska pezeta (račun)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "Španska pezeta (konvertibilniračun)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "Španska pezeta", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Etiopski bir", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Evro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Finska marka", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "Fidžijski dolar", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Foklandska funta", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "Francuski franak", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "Britanska funta", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "Gruzijski kupon larit", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "Gruzijski lari", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "Ganski cedi (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Ganski sedi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Gibraltarska funta", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Gambijski dalasi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Gvinejski franak", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "Gvinejski sili", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "Ekvatorijalno-gvinejski ekvele", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "Grčka drahma", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "Gvatemalski kecal", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "Portugalska gvineja eskudo", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "Gvineja Bisao Pezo", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "Gvajanski dolar", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Hongkonški dolar", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Honduraška lempira", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "Hrvatski dinar", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "Hrvatska kuna", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Haićanski gurd", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Mađarska forinta", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Indonežanska rupija", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "Irska funta", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "Izraelska funta", Symbol: ""},
				currency.ILR: cldr.Currency{DisplayName: "Stari izraelski šekeli", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "Izraelski novi šekel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Indijska rupija", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Irački dinar", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Iranski rijal", Symbol: "IRR"},
				currency.ISJ: cldr.Currency{DisplayName: "Stara islandska kruna", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Islandska kruna", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "Italijanska lira", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "Jamajčanski dolar", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Jordanski dinar", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Japanski jen", Symbol: "¥"},
				currency.KES: cldr.Currency{DisplayName: "Kenijski šiling", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Kirgistanski som", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Kambodžanski rijel", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Komorski franak", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Severnokorejski von", Symbol: "KPW"},
				currency.KRH: cldr.Currency{DisplayName: "Južnokorejski hvan", Symbol: ""},
				currency.KRO: cldr.Currency{DisplayName: "Stari južnokorejski von", Symbol: ""},
				currency.KRW: cldr.Currency{DisplayName: "Južnokorejski von", Symbol: "KRW"},
				currency.KWD: cldr.Currency{DisplayName: "Kuvajtski dinar", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Kajmanski dolar", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Kazahstanski tenge", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Laoški kip", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Libanska funta", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Šrilankanska rupija", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Liberijski dolar", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Lesoto loti", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Litvanski litas", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "Litvanski talonas", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "Luksemburški konvertibilni franak", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "Luksemburški franak", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "Luksemburški finansijski franak", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "Latvijski lati", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "Latvijska rublja", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "Libijski dinar", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Marokanski dirham", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "Marokanski franak", Symbol: ""},
				currency.MCF: cldr.Currency{DisplayName: "Monegaskanski franak", Symbol: ""},
				currency.MDC: cldr.Currency{DisplayName: "Moldovanski kupon", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Moldavski lej", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Madagaskarski ariari", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "Malagasijski franak", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "Makedonski denar", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "Stari makedonski denar", Symbol: ""},
				currency.MLF: cldr.Currency{DisplayName: "Malijanski franak", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Mjanmarski kjat", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Mongolski tugrik", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Makaoska pataka", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Mauritanijska ogija (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Mauritanska ogija", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "Malteška lira", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "Malteška funta", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Mauricijska rupija", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Maldivska rufija", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Malavijska kvača", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Meksički pezos", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "Meksički srebrni pezo (1861–1992)", Symbol: ""},
				currency.MXV: cldr.Currency{DisplayName: "Meksički unidad de inversion (UDI)", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "Malezijski ringit", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "Mozambijski eskudo", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "Stari mozambijski metikal", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Mozambički metikal", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Namibijski dolar", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Nigerijska naira", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "Nikaragvanska kordoba", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "Nikaragvanska zlatna kordoba", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "Holandski gulden", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "Norveška kruna", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Nepalska rupija", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Novozelandski dolar", Symbol: "NZD"},
				currency.OMR: cldr.Currency{DisplayName: "Omanski rijal", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Panamska balboa", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "Peruanski inti", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Peruanski sol", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "Peruanski sol (1863–1965)", Symbol: "PES"},
				currency.PGK: cldr.Currency{DisplayName: "Papuanska kina", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Filipinski pezos", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Pakistanska rupija", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Poljski zlot", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "Poljski zloti (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "Portugalski eskudo", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "Paragvajski gvarani", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Katarski rijal", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "Rodejskidolar", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "Rumunski lej (1952–2006)", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Rumunski lej", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Srpski dinar", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Ruska rublja", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "Ruska rublja (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Ruandski franak", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Saudijski rijal", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Solomonski dolar", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Sejšelska rupija", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "Stari sudanski dinar", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Sudanska funta", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "Stara sudanska funta", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Švedska kruna", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Singapurski dolar", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Svete Jelene funta", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "Slovenački tolar", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "Slovačka kruna", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Sijeraleonski leone", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Somalijski šiling", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Surinamski dolar", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "Surinamski gilder", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "Južnosudanska funta", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "Saotomska dobra (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "Saotomska dobra", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "Sovjetska rublja", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "Salvadorski kolon", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "Sirijska funta", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Svazilendski lilangeni", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Tajlandski bat", Symbol: "THB"},
				currency.TJR: cldr.Currency{DisplayName: "Tadžihistanska rublja", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "Tadžikistanski somon", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "Turkmenistanski manat (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Turkmenistanski manat", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Tuniski dinar", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Tonganska panga", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "Timorški eskudo", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "Turska lira (1922–2005)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "Turska lira", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Trinidad-tobagoški dolar", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Novi tajvanski dolar", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Tanzanijski šiling", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Ukrajinska hrivnja", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "Ukrajinski karbovaneti", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "Ugandski šiling (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "Ugandski šiling", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "Američki dolar", Symbol: "US$"},
				currency.USN: cldr.Currency{DisplayName: "SAD dolar (sledeći dan)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "SAD dolar (isti dan)", Symbol: ""},
				currency.UYI: cldr.Currency{DisplayName: "Urugvajski pezo en unidades indeksadas", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "Urugvajski pezo (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "Urugvajski pezos", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Uzbekistanski som", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "Venecuelanski bolivar (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "Venecuelanski bolivar (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Venecuelanski bolivar", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Vijetnamski dong", Symbol: "VND"},
				currency.VNN: cldr.Currency{DisplayName: "Vijetnamski dong (1978–1985)", Symbol: ""},
				currency.VUV: cldr.Currency{DisplayName: "Vanuatski vatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Samoanska tala", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "CFA franak BEAC", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "Srebro", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "Zlato", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "Evropska kompozitna jedinica", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "Evropska novčana jedinica", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "Evropska jedinica računa (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "Evropska jedinica računa (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "Istočnokaripski dolar", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "Posebna crtaća prava", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "Evropska valutna jedinica", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "Francuski zlatni franak", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "Francuski UIC-franak", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "CFA franak BCEAO", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "Paladijum", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "CFP franak", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "Platina", Symbol: ""},
				currency.XRE: cldr.Currency{DisplayName: "RINET fond", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "Kod testirane valute", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "Nepoznata valuta", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "Jemenski dinar", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "Jemenski rijal", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "Jugoslovenski tvrdi dinar", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "Jugoslovenski novi dinar", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "Jugoslovenski konvertibilni dinar", Symbol: ""},
				currency.YUR: cldr.Currency{DisplayName: "Jugoslovenski reformirani dinar", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "Južno-afrički rand (finansijski)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Južnoafrički rand", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Zambijska kvača (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Zambijska kvača", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "Zairski novi zair", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "Zairski zair", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "Zimbabveanski dolar (1980–2008)", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "Zimbabveanski dolar (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "Zimbabveanski dolar (2008)", Symbol: ""},
			},
		},
	}
}