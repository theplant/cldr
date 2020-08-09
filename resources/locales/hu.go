// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_hu() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y. MMMM d., EEEE", Long: "y. MMMM d.", Medium: "y. MMM d.", Short: "y. MM. dd."}, Time: cldr.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan.", Feb: "febr.", Mar: "márc.", Apr: "ápr.", May: "máj.", Jun: "jún.", Jul: "júl.", Aug: "aug.", Sep: "szept.", Oct: "okt.", Nov: "nov.", Dec: "dec."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "Á", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "Sz", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "január", Feb: "február", Mar: "március", Apr: "április", May: "május", Jun: "június", Jul: "július", Aug: "augusztus", Sep: "szeptember", Oct: "október", Nov: "november", Dec: "december"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "V", Mon: "H", Tue: "K", Wed: "Sze", Thu: "Cs", Fri: "P", Sat: "Szo"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "V", Mon: "H", Tue: "K", Wed: "Sz", Thu: "Cs", Fri: "P", Sat: "Sz"}, Short: cldr.CalendarDayFormatNameValue{Sun: "V", Mon: "H", Tue: "K", Wed: "Sze", Thu: "Cs", Fri: "P", Sat: "Szo"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "vasárnap", Mon: "hétfő", Tue: "kedd", Wed: "szerda", Thu: "csütörtök", Fri: "péntek", Sat: "szombat"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "de.", PM: "du."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "de.", PM: "du."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "de.", PM: "du."}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_hu]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "Andorrai peseta", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "EAE-dirham", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "afgán afghani (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "afgán afghani", Symbol: "AFN"},
				currency.ALK: cldr.Currency{DisplayName: "albán lek (1946–1965)", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "albán lek", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "örmény dram", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "holland antilláki forint", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "angolai kwanza", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "Angolai kwanza (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "Angolai új kwanza (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "Angolai kwanza reajustado (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "Argentín austral", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "Argentín peso (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "argentin peso", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "Osztrák schilling", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "ausztrál dollár", Symbol: "AUD"},
				currency.AWG: cldr.Currency{DisplayName: "arubai florin", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "azerbajdzsáni manat (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "azerbajdzsáni manat", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "Bosznia-hercegovinai dínár (1992–1994)", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "bosznia-hercegovinai konvertibilis márka", Symbol: "BAM"},
				currency.BAN: cldr.Currency{DisplayName: "bosznia-hercegovinai új dínár (1994–1997)", Symbol: ""},
				currency.BBD: cldr.Currency{DisplayName: "barbadosi dollár", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "bangladesi taka", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "Belga frank (konvertibilis)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "Belga frank", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "Belga frank (pénzügyi)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "Bolgár kemény leva", Symbol: ""},
				currency.BGM: cldr.Currency{DisplayName: "bolgár szocialista leva", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "bolgár új leva", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "bolgár leva (1879–1952)", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "bahreini dinár", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "burundi frank", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "bermudai dollár", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "brunei dollár", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "bolíviai boliviano", Symbol: "BOB"},
				currency.BOP: cldr.Currency{DisplayName: "Bolíviai peso", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "Bolíviai mvdol", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "Brazi cruzeiro novo (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "Brazi cruzado (1986–1989)", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "Brazil cruzeiro (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "brazil real", Symbol: "BRL"},
				currency.BRN: cldr.Currency{DisplayName: "Brazil cruzado novo (1989–1990)", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "Brazil cruzeiro (1993–1994)", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "bahamai dollár", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "bhutáni ngultrum", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "Burmai kyat", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "botswanai pula", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "Fehérorosz új rubel (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "belarusz rubel", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "fehérorosz rubel (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "belize-i dollár", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "kanadai dollár", Symbol: "CAD"},
				currency.CDF: cldr.Currency{DisplayName: "kongói frank", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "WIR euro", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "svájci frank", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "WIR frank", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "Chilei unidades de fomento", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "chilei peso", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "kínai jüan (offshore)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "kínai jüan", Symbol: "CNY"},
				currency.COP: cldr.Currency{DisplayName: "kolumbiai peso", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "Unidad de Valor Real", Symbol: ""},
				currency.CRC: cldr.Currency{DisplayName: "Costa Rica-i colon", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "szerb dinár", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "Csehszlovák kemény korona", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "kubai konvertibilis peso", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "kubai peso", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Cape Verde-i escudo", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "Ciprusi font", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "cseh korona", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "Kelet-Német márka", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "Német márka", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "dzsibuti frank", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "dán korona", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "dominikai peso", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "algériai dínár", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "Ecuadori sucre", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "Ecuadori Unidad de Valor Constante (UVC)", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "Észt korona", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "egyiptomi font", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "eritreai nakfa", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "spanyol peseta (A–kontó)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "spanyol peseta (konvertibilis kontó)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "Spanyol peseta", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "etiópiai birr", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "euró", Symbol: "EUR"},
				currency.FIM: cldr.Currency{DisplayName: "Finn markka", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "fidzsi dollár", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "falkland-szigeteki font", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "Francia frank", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "angol font", Symbol: "GBP"},
				currency.GEK: cldr.Currency{DisplayName: "Grúz kupon larit", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "grúz lari", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "Ghánai cedi (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "ghánai cedi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "gibraltári font", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "gambiai dalasi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "guineai frank", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "Guineai syli", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "Egyenlítői-guineai ekwele guineana", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "Görög drachma", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "guatemalai quetzal", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "Portugál guinea escudo", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "Guinea-Bissaui peso", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "guyanai dollár", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "hongkongi dollár", Symbol: "HKD"},
				currency.HNL: cldr.Currency{DisplayName: "hodurasi lempira", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "Horvát dínár", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "horvát kuna", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "haiti gourde", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "magyar forint", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "indonéz rúpia", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "Ír font", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "Izraeli font", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "izraeli új sékel", Symbol: "ILS"},
				currency.INR: cldr.Currency{DisplayName: "indiai rúpia", Symbol: "INR"},
				currency.IQD: cldr.Currency{DisplayName: "iraki dínár", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "iráni rial", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "izlandi korona", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "Olasz líra", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "jamaicai dollár", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "jordániai dínár", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "japán jen", Symbol: "¥"},
				currency.KES: cldr.Currency{DisplayName: "kenyai shilling", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "kirgizisztáni szom", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "kambodzsai riel", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "comorei frank", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "észak-koreai won", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "dél-koreai won", Symbol: "KRW"},
				currency.KWD: cldr.Currency{DisplayName: "kuvaiti dínár", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "kajmán-szigeteki dollár", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "kazahsztáni tenge", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "laoszi kip", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "libanoni font", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Srí Lanka-i rúpia", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "libériai dollár", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Lesothoi loti", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "litvániai litas", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "Litvániai talonas", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "luxemburgi konvertibilis frank", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "Luxemburgi frank", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "luxemburgi pénzügyi frank", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "lett lats", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "Lett rubel", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "líbiai dínár", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "marokkói dirham", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "Marokkói frank", Symbol: ""},
				currency.MDC: cldr.Currency{DisplayName: "moldáv kupon", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "moldován lei", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "madagaszkári ariary", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "Madagaszkári frank", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "macedon dínár", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "macedón dénár (1992–1993)", Symbol: ""},
				currency.MLF: cldr.Currency{DisplayName: "Mali frank", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "mianmari kyat", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "mongóliai tugrik", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "makaói pataca", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "mauritániai ouguiya (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "mauritániai ouguiya", Symbol: ""},
				currency.MTL: cldr.Currency{DisplayName: "Máltai líra", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "Máltai font", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "mauritiusi rúpia", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "maldív-szigeteki rufiyaa", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "malawi kwacha", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "mexikói peso", Symbol: "MXN"},
				currency.MXP: cldr.Currency{DisplayName: "Mexikói ezüst peso (1861–1992)", Symbol: ""},
				currency.MXV: cldr.Currency{DisplayName: "Mexikói Unidad de Inversion (UDI)", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "malajziai ringgit", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "Mozambik escudo", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "Mozambik metical", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "mozambiki metikális", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "namíbiai dollár", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "nigériai naira", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "Nikaraguai cordoba", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "nicaraguai córdoba", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "Holland forint", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "norvég korona", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "nepáli rúpia", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "új-zélandi dollár", Symbol: "NZD"},
				currency.OMR: cldr.Currency{DisplayName: "ománi rial", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "panamai balboa", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "perui inti", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "perui sol", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "perui sol (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "pápua új-guineai kina", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "fülöp-szigeteki peso", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "pakisztáni rúpia", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "lengyel zloty", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "Lengyel zloty (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "Portugál escudo", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "paraguayi guarani", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "katari rial", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "rhodéziai dollár", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "román lej (1952–2006)", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "román lej", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "szerb dínár", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "orosz rubel", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "orosz rubel (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "ruandai frank", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "szaúdi riyal", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "salamon-szigeteki dollár", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "seychelle-szigeteki rúpia", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "Szudáni dínár (1992–2007)", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "szudáni font", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "Szudáni font (1957–1998)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "svéd korona", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "szingapúri dollár", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Szent Ilona-i font", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "Szlovén tolar", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "Szlovák korona", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Sierra Leone-i leone", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "szomáli shilling", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "suriname-i dollár", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "Suriname-i gulden", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "dél-szudáni font", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "São Tomé és Príncipe-i dobra (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "São Tomé és Príncipe-i dobra", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "Szovjet rubel", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "Salvadori colón", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "szíriai font", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "szvázi lilangeni", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "thai baht", Symbol: "THB"},
				currency.TJR: cldr.Currency{DisplayName: "Tádzsikisztáni rubel", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "tádzsikisztáni somoni", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "türkmenisztáni manat (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "türkmenisztáni manat", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "tunéziai dínár", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "tongai paanga", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "Timori escudo", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "török líra (1922–2005)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "török líra", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Trinidad és Tobago-i dollár", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "tajvani új dollár", Symbol: "TWD"},
				currency.TZS: cldr.Currency{DisplayName: "tanzániai shilling", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "ukrán hrivnya", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "Ukrán karbovanec", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "Ugandai shilling (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "ugandai shilling", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "USA-dollár", Symbol: "USD"},
				currency.USN: cldr.Currency{DisplayName: "USA dollár (következő napi)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "USA dollár (aznapi)", Symbol: ""},
				currency.UYI: cldr.Currency{DisplayName: "Uruguayi peso en unidades indexadas", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "Uruguay-i peso (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "uruguay-i peso", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "üzbegisztáni szum", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "Venezuelai bolivar (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "venezuelai bolivar (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "venezuelai bolivar", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "vietnámi dong", Symbol: "VND"},
				currency.VNN: cldr.Currency{DisplayName: "vietnámi dong (1978–1985)", Symbol: ""},
				currency.VUV: cldr.Currency{DisplayName: "vanuatui vatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "nyugat-szamoai tala", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "CFA frank BEAC", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "Ezüst", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "Arany", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "Európai kompozit egység", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "Európai monetáris egység", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "Európai kontó egység (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "Európai kontó egység (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "kelet-karibi dollár", Symbol: "XCD"},
				currency.XDR: cldr.Currency{DisplayName: "Special Drawing Rights", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "európai pénznemegység", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "Francia arany frank", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "Francia UIC-frank", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "CFA frank BCEAO", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "palládium", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "csendes-óceáni valutaközösségi frank", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "platina", Symbol: ""},
				currency.XRE: cldr.Currency{DisplayName: "RINET tőke", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "Tesztelési pénznemkód", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "ismeretlen pénznem", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "Jemeni dínár", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "jemeni rial", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "Jugoszláv kemény dínár", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "Jugoszláv új dínár", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "Jugoszláv konvertibilis dínár", Symbol: ""},
				currency.YUR: cldr.Currency{DisplayName: "jugoszláv reformált dinár (1992–1993)", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "Dél-afrikai rand (pénzügyi)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "dél-afrikai rand", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Zambiai kwacha (1968–2012)", Symbol: "ZMK"},
				currency.ZMW: cldr.Currency{DisplayName: "zambiai kwacha", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "Zairei új zaire", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "Zairei zaire", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "Zimbabwei dollár (1980–2008)", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "Zimbabwei dollár (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "Zimbabwei dollár (2008)", Symbol: ""},
			},
		},
	}
}
