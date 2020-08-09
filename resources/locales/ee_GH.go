// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_ee_GH() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, MMMM d 'lia' y", Long: "MMMM d 'lia' y", Medium: "MMM d 'lia', y", Short: "M/d/yy"}, Time: cldr.CalendarDateFormat{Full: "a 'ga' h:mm:ss zzzz", Long: "a 'ga' h:mm:ss z", Medium: "a 'ga' h:mm:ss", Short: "a 'ga' h:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{0} {1}", Long: "{0} {1}", Medium: "{0} {1}", Short: "{0} {1}"}, GMT: "{0} GMT"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "dzv", Feb: "dzd", Mar: "ted", Apr: "afɔ", May: "dam", Jun: "mas", Jul: "sia", Aug: "dea", Sep: "any", Oct: "kel", Nov: "ade", Dec: "dzm"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "d", Feb: "d", Mar: "t", Apr: "a", May: "d", Jun: "m", Jul: "s", Aug: "d", Sep: "a", Oct: "k", Nov: "a", Dec: "d"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "dzove", Feb: "dzodze", Mar: "tedoxe", Apr: "afɔfĩe", May: "dama", Jun: "masa", Jul: "siamlɔm", Aug: "deasiamime", Sep: "anyɔnyɔ", Oct: "kele", Nov: "adeɛmekpɔxe", Dec: "dzome"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "kɔs", Mon: "dzo", Tue: "bla", Wed: "kuɖ", Thu: "yaw", Fri: "fiɖ", Sat: "mem"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "k", Mon: "d", Tue: "b", Wed: "k", Thu: "y", Fri: "f", Sat: "m"}, Short: cldr.CalendarDayFormatNameValue{Sun: "kɔs", Mon: "dzo", Tue: "bla", Wed: "kuɖ", Thu: "yaw", Fri: "fiɖ", Sat: "mem"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "kɔsiɖa", Mon: "dzoɖa", Tue: "blaɖa", Wed: "kuɖa", Thu: "yawoɖa", Fri: "fiɖa", Sat: "memleɖa"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "ŋdi", PM: "ɣetrɔ"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "ŋ", PM: "ɣ"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "ŋdi", PM: "ɣetrɔ"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_ee]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "andorraga peseta", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "united arab emiratesga dirham", Symbol: ""},
				currency.AFA: cldr.Currency{DisplayName: "afghanistanga afghani (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "afghanistanga afghani", Symbol: ""},
				currency.ALK: cldr.Currency{DisplayName: "albaniaga lek (1946–1965)", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "albaniaga lek", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "armeniaga dram", Symbol: ""},
				currency.ANG: cldr.Currency{DisplayName: "netherlands antilleaga guilder", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "angolaga kwanza", Symbol: "Kz"},
				currency.AOK: cldr.Currency{DisplayName: "angolaga kwanza (1977–1991)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "angolaga kwanza (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "angolaga kwanza xoxotɔ (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "argentinaga austral", Symbol: ""},
				currency.ARL: cldr.Currency{DisplayName: "argentinaga peso ley (1970–1983)", Symbol: ""},
				currency.ARM: cldr.Currency{DisplayName: "argentinaga peso (1881–1970)", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "argentinaga peso (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "argentinaga peso", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "ɔstriaga schilling", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "Australiaga dollar", Symbol: "AU$"},
				currency.AWG: cldr.Currency{DisplayName: "arubaga lorin", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "azerbaidzanga manat (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "azerbaidzanga manat", Symbol: ""},
				currency.BAD: cldr.Currency{DisplayName: "bosnia-herzegovinaga dinar (1992–1994)", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "bosnia-herzegovinaga convertible mark", Symbol: "BAM"},
				currency.BAN: cldr.Currency{DisplayName: "bosnia kple herzegovinaga dinar yeyètɔ (1994–1997)", Symbol: ""},
				currency.BBD: cldr.Currency{DisplayName: "barbadosga dollar", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "bangladeshga taka", Symbol: "৳"},
				currency.BEC: cldr.Currency{DisplayName: "beldziumga franc (convertible)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "beldziumga franc", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "beldziumga franc (financial)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "bɔlgariaga hard lev", Symbol: ""},
				currency.BGM: cldr.Currency{DisplayName: "bɔlgariaga socialist lev", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "bulgariaga lev", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "bulgariaga lev (1879–1952)", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "bahrainga dinar", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "burundiga franc", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "bermudaga dollar", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "bruneiga dollar", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "boliviaga boliviano", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "boliviaga boliviano (1863–1963)", Symbol: ""},
				currency.BOP: cldr.Currency{DisplayName: "boliviaga peso", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "boliviaga mvdol", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "braziliaga cruzeiro xoxotɔ (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "brazilia cruzado (1986–1989)", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "braziliaga cruzeiro xoxotɔ gbãtɔ (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "braziliaga real", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "brazilia cruzado xoxotɔ (1989–1990)", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "braziliaga cruzeiro (1993–1994)", Symbol: ""},
				currency.BRZ: cldr.Currency{DisplayName: "braziliaga cruzeiro (1942–1967)", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "bahamasga dollar", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "bhutanga ngultrum", Symbol: ""},
				currency.BUK: cldr.Currency{DisplayName: "burmaga kyat", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "botswanaga pula", Symbol: "P"},
				currency.BYB: cldr.Currency{DisplayName: "belarusiaga ruble yeytɔ (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "belarusiaga ruble", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "belarusiaga ruble (2000–2016)", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "belizega dollar", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "canadaga dollar", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "kongoga franc", Symbol: ""},
				currency.CHE: cldr.Currency{DisplayName: "WIR euro CHE", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "switzerlandga franc", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "WIR euro CHW", Symbol: ""},
				currency.CLE: cldr.Currency{DisplayName: "tsilega escudo", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "tsilegakɔnta dzidzenu UF", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "chilega peso", Symbol: "CLP"},
				currency.CNX: cldr.Currency{DisplayName: "tsainatɔwo ƒe gadzraɖoƒe dollar", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "Chinesega yuan", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "colombiaga peso", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "kolombiaga vavãtɔ", Symbol: ""},
				currency.CRC: cldr.Currency{DisplayName: "costa ricaga colón", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "serbiaga dinar (2002–2006)", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "tsɛkoslovakiaga hard koruna", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "cubaga convertible peso", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "cubaga peso", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "kape verdega escudo", Symbol: ""},
				currency.CYP: cldr.Currency{DisplayName: "saipriɔtga pound", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "czechga koruna", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "ɣedzeƒe germaniaga mark", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "germaniaga mark", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "dziboutiga franc", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "denmarkga krone", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "dominicaga peso", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "aldzeriaga dinar", Symbol: ""},
				currency.ECS: cldr.Currency{DisplayName: "ekuadɔga sucre", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "ekuadɔ dzidzenu matrɔmatrɔ", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "estoniaga kroon", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "egyptega pound", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "eritreaga nakfa", Symbol: ""},
				currency.ESA: cldr.Currency{DisplayName: "spaniaga peseta (A)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "spaniaga peseta (Convertible)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "spaniaga peseta", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "ethiopiaga birr", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "finlandga markka", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "fidziga dollar", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "falkland islands pound", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "frentsiga franc", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "britainga pound", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "dzɔdziaga kupon larit", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "dzɔdziaga lari", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "ghana siɖi (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "ghana siɖi", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "gilbratarga pound", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "gambiaga dalasi", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "giniga franc", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "giniga syli", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "ekuatorial giniga ekwele", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "grisiga drachma", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "guatemalaga quetzal", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "pɔtugaltɔwo ƒe giniga escudo", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "gini-bisau peso", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "guyanaga dollar", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Hong Kongga dollar", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "honduraga lempira", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "kroatiaga dinar", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "croatiaga kuna", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "haitiga gourde", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "hungariaga forint", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Indonesiaga rupiah", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "ireland pound", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "israelga pound", Symbol: ""},
				currency.ILR: cldr.Currency{DisplayName: "israelga sheqel (1980–1985)", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "israelga yeyetɔ sheqel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Indiaga rupee", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "irakga dinar", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "iranga rial", Symbol: ""},
				currency.ISJ: cldr.Currency{DisplayName: "aiselandga króna (1918–1981)", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "icelandga króna", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "italiaga lira", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "jamaicaga dollar", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "yɔdanga dinar", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "Japanesega yen", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "kenyaga shilling", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "kirgistanga som", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "kambodiaga riel", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "komoroga franc", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "dziehe koreaga won", Symbol: "₩"},
				currency.KRH: cldr.Currency{DisplayName: "anyiehe koreaga hwan (1953–1962)", Symbol: ""},
				currency.KRO: cldr.Currency{DisplayName: "anyiehe koreaga won (1945–1953)", Symbol: ""},
				currency.KRW: cldr.Currency{DisplayName: "South Koreaga won", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "kuwaitga dinar", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "cayman islandsga dollar", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "kazakhstanga tenge", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "laosga kip", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "lebanonga pound", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "sri lankaga rupee", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "liberiaga dollar", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "lesotoga loti", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "lithuaniaga litas", Symbol: "Lt"},
				currency.LTT: cldr.Currency{DisplayName: "lithuaniaga talonas", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "lazembɔgga convertible franc", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "lazembɔgga franc", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "lazembɔgga gadzikpɔ franc", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "latviaga lats", Symbol: "Ls"},
				currency.LVR: cldr.Currency{DisplayName: "latviaga ruble", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "libyaga dinar", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "morokoga dirham", Symbol: ""},
				currency.MAF: cldr.Currency{DisplayName: "morokoga franc", Symbol: ""},
				currency.MCF: cldr.Currency{DisplayName: "monegaskga franc", Symbol: ""},
				currency.MDC: cldr.Currency{DisplayName: "moldovaga cupon", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "moldovaga leu", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "malagasega ariary", Symbol: "Ar"},
				currency.MGF: cldr.Currency{DisplayName: "malagasega franc", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "marcedoniaga denar", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "makedoniaga denar (1992–1993)", Symbol: ""},
				currency.MLF: cldr.Currency{DisplayName: "maliga franc", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "myanmaga kyat", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "mongoliaga tugrik", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "makanesega pataca", Symbol: ""},
				currency.MRO: cldr.Currency{DisplayName: "mɔritaniaga ouguiya (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "mɔritaniaga ouguiya", Symbol: ""},
				currency.MTL: cldr.Currency{DisplayName: "maltaga lira", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "maltaga pound", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "mɔritiusga rupee", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "maldiviaga rufiyaa", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "malawiga kwacha", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "mexicoga peso", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "malaysiaga ringit", Symbol: "RM"},
				currency.MZN: cldr.Currency{DisplayName: "mozambikga metikal", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "namibiaga dollar", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "naidzeriaga naira", Symbol: "₦"},
				currency.NIC: cldr.Currency{DisplayName: "nikaraguaga córdoba (1988–1991)", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "nicaraguaga córdoba", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "hollandga guilder", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "norwayga krone", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "nepalga rupee", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "new zealanɖga dollar", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "omanga rial", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "panamaga balboa", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "peruga inti", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "peruga sol", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "peruga sol (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "papua new guineaga kina", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "filipiniga peso", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "pakistaniga rupee", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "polandga zloty", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "polanɖga zloty (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "pɔtugalga escudo", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "paraguayga guarani", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "katarga rial", Symbol: ""},
				currency.RHD: cldr.Currency{DisplayName: "rhodesiaga dollar", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "romaniaga leu (1952–2006)", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "romaniaga leu", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "serbiaga dinar", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "russiaga ruble", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "rɔtsiaga ruble (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "rwandaga franc", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Saudi Arabiaga riyal", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "solomon ƒudomekpo dukɔwo ƒe ga dollar", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "sɛtselsga rupee", Symbol: ""},
				currency.SDD: cldr.Currency{DisplayName: "sudanga dinar (1992–2007)", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "sudanga pound", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "sudanga pound (1957–1998)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "swedenga krone", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "singapɔga dollar", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "saint helenaga pound", Symbol: "£"},
				currency.SIT: cldr.Currency{DisplayName: "slovaniaga tolar", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "slovakga koruna", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "sierra leonega leone", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "somaliaga shilling", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "surinamga dollar", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "surinamega guilder", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "anyiehe sudanga pound", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "são tomé kple príncipega dobra (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "são tomé kple príncipega dobra", Symbol: "Db"},
				currency.SUR: cldr.Currency{DisplayName: "sovietga rouble", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "salvadɔga colón", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "syriaga pound", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "swaziga lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "Thailandga baht", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "tajikistanga ruble", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "tajikistanga somoni", Symbol: ""},
				currency.TMM: cldr.Currency{DisplayName: "turkmenistanga manat (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "turkmenistanga manat", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "tunisiaga dinar", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "tonagaga pa’anga", Symbol: "T$"},
				currency.TPE: cldr.Currency{DisplayName: "timɔga escudo", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "tɛkiiga lira (1922–2005)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "Turkishga lira", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "trinidad & tobagoga dollar", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Taiwanga dollar", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Tanzaniatɔwofɛgadudu", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "ukrainega hryvnia", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "ukrainega karbovanet", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "ugandaga shilling (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "ugandaga shilling", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "us ga dollar", Symbol: "US$"},
				currency.USN: cldr.Currency{DisplayName: "us ga dollar (ŋkeke si gbɔna tɔ)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "us ga dollar (ŋkeke ma ke tɔ)", Symbol: ""},
				currency.UYI: cldr.Currency{DisplayName: "uruguayga peso UYI", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "uruguayga peso (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "uruguayga peso", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "uzbekistanga som", Symbol: ""},
				currency.VEB: cldr.Currency{DisplayName: "venezuelaga bolívar (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "venezuelaga bolívar (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "venezuelaga bolívar", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "vietnamga dong", Symbol: "₫"},
				currency.VNN: cldr.Currency{DisplayName: "vietnamga dong (1978–1985)", Symbol: ""},
				currency.VUV: cldr.Currency{DisplayName: "vanuatuga vatu", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "samaoga tala", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "ɣetoɖofe afrikaga CFA franc BEAC", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "klosalo", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "sika", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "europa dzidzenu xba", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "europa gadzidzenu xbb", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "europa kɔnta dzidzenu xbc", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "europa kɔnta dzidzenu xbd", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "east caribbeanga dollar", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "gaɖuɖu ɖoɖo tɔxɛ", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "europa gaɖuɖu", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "fransemega sika franc", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "frentsi UIC-franc", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "ɣetoɖofe afrikaga CFA franc BCEAO", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "palladiumga", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "CFP ga franc", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "platinum", Symbol: ""},
				currency.XRE: cldr.Currency{DisplayName: "RINET gadodo XRE", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "gaɖuɖu dodokpɔ dzesi xts", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "gaɖuɖu manya", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "yemeniga dinar", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "yemeniga rial", Symbol: ""},
				currency.YUD: cldr.Currency{DisplayName: "yugoslaviaga hard dinar (1966–1990)", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "yugoslaviaga yeyetɔ dinar (1994–2002)", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "yugoslaviaga convertible dinar (1990–1992)", Symbol: ""},
				currency.YUR: cldr.Currency{DisplayName: "yugoslaviaga dinar (1992–1993)", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "anyiehe afrikaga rand (gadzikpɔtɔ)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "South Africaga rand", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "zambiaga kwacha (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "zambiaga kwacha", Symbol: "ZK"},
				currency.ZRN: cldr.Currency{DisplayName: "zairega yeyetɔ zaire", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "zairega zaire (1971–1993)", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "zimbabwega dollar (1980–2008)", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "zimbabwega dollar (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "zimbabwega dollar (2008)", Symbol: ""},
			},
		},
	}
}
