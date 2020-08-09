// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_is() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM y", Short: "d.M.y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'kl'. {0}", Long: "{1} 'kl'. {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan.", Feb: "feb.", Mar: "mar.", Apr: "apr.", May: "maí", Jun: "jún.", Jul: "júl.", Aug: "ágú.", Sep: "sep.", Oct: "okt.", Nov: "nóv.", Dec: "des."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "Á", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "janúar", Feb: "febrúar", Mar: "mars", Apr: "apríl", May: "maí", Jun: "júní", Jul: "júlí", Aug: "ágúst", Sep: "september", Oct: "október", Nov: "nóvember", Dec: "desember"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "sun.", Mon: "mán.", Tue: "þri.", Wed: "mið.", Thu: "fim.", Fri: "fös.", Sat: "lau."}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "Þ", Wed: "M", Thu: "F", Fri: "F", Sat: "L"}, Short: cldr.CalendarDayFormatNameValue{Sun: "su.", Mon: "má.", Tue: "þr.", Wed: "mi.", Thu: "fi.", Fri: "fö.", Sat: "la."}, Wide: cldr.CalendarDayFormatNameValue{Sun: "sunnudagur", Mon: "mánudagur", Tue: "þriðjudagur", Wed: "miðvikudagur", Thu: "fimmtudagur", Fri: "föstudagur", Sat: "laugardagur"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "f.h.", PM: "e.h."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "f.h.", PM: "e.h."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "f.h.", PM: "e.h."}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_is]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "", Group: "", Negative: "\u061c-", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "Andorrskur peseti", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "arabískt dírham", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "afgani", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "albanskt lek", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "armenskt dramm", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "hollenskt Antillugyllini", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "angólsk kvansa", Symbol: "AOA"},
				currency.ARA: cldr.Currency{DisplayName: "Argentine Austral", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "Argentískur pesi (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "argentínskur pesi", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "Austurrískur skildingur", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "ástralskur dalur", Symbol: "AUD"},
				currency.AWG: cldr.Currency{DisplayName: "arúbönsk flórína", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "aserskt manat", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "skiptanlegt Bosníu og Hersegóvínu-mark", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "barbadoskur dalur", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "bangladessk taka", Symbol: "BDT"},
				currency.BEF: cldr.Currency{DisplayName: "Belgískur franki", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "Lef", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "búlgarskt lef", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "bareinskur denari", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "búrúndískur franki", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Bermúdadalur", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "brúneiskur dalur", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "bólivíani", Symbol: "BOB"},
				currency.BOP: cldr.Currency{DisplayName: "Bólivískur pesi", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "Bolivian Mvdol", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "brasilískt ríal", Symbol: "BRL"},
				currency.BSD: cldr.Currency{DisplayName: "Bahamadalur", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "bútanskt núltrum", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "Búrmverskt kjat", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "botsvönsk púla", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "hvítrússnesk rúbla", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "hvítrússnesk rúbla (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "belískur dalur", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Kanadadalur", Symbol: "CAD"},
				currency.CDF: cldr.Currency{DisplayName: "kongóskur franki", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "svissneskur franki", Symbol: "CHF"},
				currency.CLF: cldr.Currency{DisplayName: "Chilean Unidades de Fomento", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "síleskur pesi", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "kínverskt júan (utan heimalands)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "kínverskt júan", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "kólumbískur pesi", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "kostarískt kólon", Symbol: "CRC"},
				currency.CSK: cldr.Currency{DisplayName: "Tékknesk króna, eldri", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "kúbverskur skiptanlegur pesi", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "kúbverskur pesi", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "grænhöfðeyskur skúti", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "Kýpverskt pund", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "tékknesk króna", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "Austurþýskt mark", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "Þýskt mark", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "djíbútískur franki", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "dönsk króna", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "dóminískur pesi", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "alsírskur denari", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "Ecuador Sucre", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "Eistnesk króna", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "egypskt pund", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "erítresk nakfa", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "Spænskur peseti", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "eþíópískt birr", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "evra", Symbol: "EUR"},
				currency.FIM: cldr.Currency{DisplayName: "Finnskt mark", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "fidjeyskur dalur", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "falklenskt pund", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "Franskur franki", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "sterlingspund", Symbol: "GBP"},
				currency.GEL: cldr.Currency{DisplayName: "georgískur lari", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "ganverskur sedi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Gíbraltarspund", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "gambískur dalasi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Gíneufranki", Symbol: "GNF"},
				currency.GRD: cldr.Currency{DisplayName: "Drakma", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "gvatemalskt kvesal", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "Portúgalskur, gíneskur skúti", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "gvæjanskur dalur", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Hong Kong-dalur", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "hondúrsk lempíra", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "króatísk kúna", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "haítískur gúrdi", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "ungversk fórinta", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "indónesísk rúpía", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "Írskt pund", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "Ísraelskt pund", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "nýr ísraelskur sikill", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "indversk rúpía", Symbol: "INR"},
				currency.IQD: cldr.Currency{DisplayName: "írakskur denari", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "íranskt ríal", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "íslensk króna", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "Ítölsk líra", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "jamaískur dalur", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "jórdanskur denari", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "japanskt jen", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "kenískur skildingur", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "kirgiskt som", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "kambódískt ríal", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "kómoreyskur franki", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "norðurkóreskt vonn", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "suðurkóreskt vonn", Symbol: "KRW"},
				currency.KWD: cldr.Currency{DisplayName: "kúveiskur denari", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "caymaneyskur dalur", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "kasakst tengi", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "laoskt kip", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "líbanskt pund", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "srílönsk rúpía", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "líberískur dalur", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Lesotho Loti", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Litháískt lít", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "Lithuanian Talonas", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "Lúxemborgarfranki", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "Lettneskt lat", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "Lettnesk rúbla", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "líbískur denari", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "marokkóskt dírham", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "Marokkóskur franki", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "moldavískt lei", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Madagaskararjari", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "Madagaskur franki", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "makedónskur denari", Symbol: "MKD"},
				currency.MLF: cldr.Currency{DisplayName: "Malískur franki", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "mjanmarskt kjat", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "mongólskur túríkur", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "makaósk pataka", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "márítönsk úgía (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "márítönsk úgía", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "Meltnesk líra", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "Maltneskt pund", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "máritísk rúpía", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "maldíveysk rúpía", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "malavísk kvaka", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "mexíkóskur pesi", Symbol: "MXN"},
				currency.MXP: cldr.Currency{DisplayName: "Mexíkóskur silfurpesi (1861–1992)", Symbol: ""},
				currency.MXV: cldr.Currency{DisplayName: "Mexíkóskur pesi, UDI", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "malasískt ringit", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "Mósambískur skúti", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "mósambískt metikal", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "namibískur dalur", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "nígerísk næra", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "Níkarögsk kordóva (1988–1991)", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "níkaraögsk kordóva", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "Hollenskt gyllini", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "norsk króna", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "nepölsk rúpía", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "nýsjálenskur dalur", Symbol: "NZD"},
				currency.OMR: cldr.Currency{DisplayName: "ómanskt ríal", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "balbói", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "perúskt sól", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "papúsk kína", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "filippseyskur pesi", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "pakistönsk rúpía", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "pólskt slot", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "Slot", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "Portúgalskur skúti", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "paragvæskt gvaraní", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "katarskt ríal", Symbol: "QAR"},
				currency.ROL: cldr.Currency{DisplayName: "Rúmenskt lei (1952–2006)", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "rúmenskt lei", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "serbneskur denari", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "rússnesk rúbla", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "Rússnesk rúbla (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "rúandskur franki", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "sádíarabískt ríal", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "salómonseyskur dalur", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Seychellesrúpía", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "Súdanskur denari", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "súdanskt pund", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "Súdanskt pund (1957–1998)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "sænsk króna", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "singapúrskur dalur", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "helenskt pund", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "Slóvenskur dalur", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "Slóvakísk króna", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "síerraleónsk ljóna", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "sómalískur skildingur", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Súrínamdalur", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "Suriname Guilder", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "suðursúdanskt pund", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "Saó Tóme og Prinsípe-dóbra (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "Saó Tóme og Prinsípe-dóbra", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "Soviet Rouble", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "El Salvador Colon", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "sýrlenskt pund", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "svasílenskur lílangeni", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "taílenskt bat", Symbol: "THB"},
				currency.TJR: cldr.Currency{DisplayName: "Tadsjiksk rúbla", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "tadsjikskur sómóni", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "Túrkmenskt manat (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "túrkmenskt manat", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "túnískur denari", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Tongapanga", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "Tímorskur skúti", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "Tyrknesk líra (1922–2005)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "tyrknesk líra", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Trínidad og Tóbagó-dalur", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "taívanskur dalur", Symbol: "TWD"},
				currency.TZS: cldr.Currency{DisplayName: "tansanískur skildingur", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "úkraínsk hrinja", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "Ukrainian Karbovanetz", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "úgandskur skildingur", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "Bandaríkjadalur", Symbol: "USD"},
				currency.USN: cldr.Currency{DisplayName: "Bandaríkjadalur (næsta dag)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "Bandaríkjadalur (sama dag)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "úrúgvæskur pesi", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "úsbekskt súm", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "Bolívar í Venesúela (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "venesúelskur bólívari (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "venesúelskur bólívari", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "víetnamskt dong", Symbol: "VND"},
				currency.VUV: cldr.Currency{DisplayName: "vanúatúskt vatú", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Samóatala", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "miðafrískur franki", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "unse silfur", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "unse gull", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "austurkarabískur dalur", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "Sérstök dráttarréttindi", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "Franskur gullfranki", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "Franskur franki, UIC", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "vesturafrískur franki", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "unse palladín", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "pólinesískur franki", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "unse platína", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "óþekktur gjaldmiðill", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "Jemenskur denari", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "jemenskt ríal", Symbol: "YER"},
				currency.YUM: cldr.Currency{DisplayName: "Júgóslavneskur denari", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "Rand (viðskipta)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "suðurafrískt rand", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Zambian Kwacha (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "sambísk kvaka", Symbol: "ZMW"},
				currency.ZWD: cldr.Currency{DisplayName: "Simbabveskur dalur", Symbol: ""},
			},
		},
	}
}
