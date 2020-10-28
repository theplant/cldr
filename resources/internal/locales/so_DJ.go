// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_so_DJ() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "so_DJ",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, MMMM dd, y", Long: "dd MMMM y", Medium: "dd-MMM-y", Short: "dd/MM/yy"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mar", Apr: "Abr", May: "May", Jun: "Jun", Jul: "Lul", Aug: "Ogs", Sep: "Seb", Oct: "Okt", Nov: "Nof", Dec: "Dis"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "L", Aug: "O", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Jannaayo", Feb: "Febraayo", Mar: "Maarso", Apr: "Abriil", May: "May", Jun: "Juun", Jul: "Luuliyo", Aug: "Ogost", Sep: "Sebtembar", Oct: "Oktoobar", Nov: "Nofembar", Dec: "Desembar"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Axd", Mon: "Isn", Tue: "Tldo", Wed: "Arbc", Thu: "Khms", Fri: "Jmc", Sat: "Sbti"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "A", Mon: "I", Tue: "T", Wed: "A", Thu: "Kh", Fri: "J", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Axd", Mon: "Isn", Tue: "Tldo", Wed: "Arbc", Thu: "Khms", Fri: "Jmc", Sat: "Sbti"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Axad", Mon: "Isniin", Tue: "Talaado", Wed: "Arbaco", Thu: "Khamiis", Fri: "Jimco", Sat: "Sabti"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "GH", PM: "GD"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "h", PM: "d"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "GH", PM: "GD"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "", Group: "", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "0 Kun", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Dirhamka Isutaga Imaaraatka Carabta", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Afgan Afgani", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "Lekta Albaniya", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "Daraamka Armeniya", Symbol: ""},
				currency.ANG: cldr.Currency{DisplayName: "Galdarka Nadarlaan Antiliyaan", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Kawansada Angola", Symbol: "Kz"},
				currency.ARA: cldr.Currency{DisplayName: "Argentine Austral", Symbol: ""},
				currency.ARL: cldr.Currency{DisplayName: "Beesada Ley ee Arjentiin (1970–1983)", Symbol: ""},
				currency.ARM: cldr.Currency{DisplayName: "Beesada Ley ee Arjentiin (1881–1970)", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "Beesada Ley ee Arjentiin (1883–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "Beesada Arjentiin", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Doolarka Astaraaliya", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Foloorinta Aruban", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Manaata Asarbeyjan", Symbol: ""},
				currency.BAD: cldr.Currency{DisplayName: "Diinaarka BBosnia-Hersogofina 1.00 konfatibal maakta Bosnia-Hersogofina 1 konfatibal maaka Bosnia-Hersogofina (1992–1994)", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Konfatibal Maakta Bosnia-Hersogofina", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "Doolarka Barbaadiyaan", Symbol: "DBB"},
				currency.BDT: cldr.Currency{DisplayName: "Taka Bangledesh", Symbol: "৳"},
				currency.BGN: cldr.Currency{DisplayName: "Lefta Bulgariya", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Dinaarka Baxreyn", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Faranka Burundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "Doolarka Barmuuda", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "Doolarka Buruney", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "Bolifiyanada Bolifiya", Symbol: "Bs"},
				currency.BOL: cldr.Currency{DisplayName: "Bolifiyaabka Bolifiyaano(1863–1963)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "Realka Barasil", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Doolarka Bahamaas", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "Nugultaramta Butan", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Buulada Botswana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "Rubalka Belarus", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "Doolarka Beelisa", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Doolarka Kanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Faranka Kongo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Faranka Iswiska", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Beesada Jili", Symbol: "$"},
				currency.CNH: cldr.Currency{DisplayName: "Yuwanta Shiinaha (Ofshoor)", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "Yuwanta Shiinaha", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Beesada Kolombiya", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "Kolonka Kosta Riika", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "Beesada Konfatibal ee Kuuba", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "Beesada Kuuba", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Eskudo Keyb Farde", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Korunada Jeek", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Faran Jabuuti", Symbol: "Fdj"},
				currency.DKK: cldr.Currency{DisplayName: "Koronka Danishka", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "Beesada Dominiika", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dinaarka Aljeriya", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "Kroonka Estooniya", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Bowndka Masar", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfada Eritriya", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Birta Itoobbiya", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Yuuroo", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Markkada Fiinishka ah", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "Doolarka Fiji", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "Bowndka Faalklaan Aylaanis", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Bowndka Biritishka", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Laariga Joorjiya", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "Sedida Gana", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "Bowndka Gibraltar", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasida Gambiya", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "Faranka Gini", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "Kuwestalka Guwatemala", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "Doolarka Guyanes", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "Doolarka Hoon Koon", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Lembirada Honduras", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "Kunada Korooshiya", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "Goordada Hiyati", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "Forintiska Hangari", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "Rubiah Indonesiya", Symbol: "Rp"},
				currency.IEP: cldr.Currency{DisplayName: "baawnka Ayrishka", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "Niyuu Shekelka Israaiil", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rubiga Hindiya", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Dinaarka Ciraaq", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "Riyaalka Iran", Symbol: ""},
				currency.ISJ: cldr.Currency{DisplayName: "krónaha Iceland (1918–1981)", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Koronada Eysland", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "Doolarka Jamayka", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "Dinaarka Urdun", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "Yenta Jabaan", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Shilingka Kenya", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "Somta Kiyrgiystan", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "Riyf kambodiya", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Faranka Komoros", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "Wonka Waqooyiga Kuuriya", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "Wonka Koonfur Kuuriya", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Dinaarka Kuweyt", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "Doolarka Kayman Aylaanis", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "Tengeda Kasakhstan", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "Kib Laoti", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "Bowndka Lubnaan", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "Rubiga Siri lanka", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Doolarka Liberiya", Symbol: "$"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LVR: cldr.Currency{DisplayName: "Rubalka Latfiya", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "Dinaarka Libya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Dirhamka Moroko", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Leeyuuda Moldofa", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Arayrida Madagaskar", Symbol: "Ar"},
				currency.MKD: cldr.Currency{DisplayName: "Denaarka Masedoniya", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Kayatda Mayanmaar", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "Tugrikta Mongoliya", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "Bataka Makana", Symbol: ""},
				currency.MRO: cldr.Currency{DisplayName: "Oogiya Mawritaniya (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Oogiyada Mawritaaniya", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rubiga Mowrishiya", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "Rufiyada Maldifiya", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "Kawajada Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "Beesada Meksiko", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Ringitda Malayshiya", Symbol: "RM"},
				currency.MZN: cldr.Currency{DisplayName: "Metikalka Mosambik", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Doolarka Namibiya", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Nairada Neyjeeriya", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "Kordobada Nikargow", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "Koronka Norway", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "Rubiga Nebal", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "Doolarka Niyuu Siyalaan", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Riyaalka Cumaan", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "Balbow Banama", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Solsha Beeru", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Kinada Babua Niyuu Gini", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "Biso Filibin", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "Rubiga Bakistan", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "Solotida Bolaan", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "Guranida Baraguway", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "Riyaalka Qatar", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Liyuuda Romaniya", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "Dinaarka Serbiya", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "Rubalka Ruushka", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Faranka Ruwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Riyaalka Sacuudiga", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "Doolarka Solomon Aylaanis", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Rubiga Siisalis", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Bowndka Suudaan", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Koronka Isweden", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "Doolarka Singabuur", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Bowndka St Helen", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Leonka Sira Leon", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Shilingka Soomaaliya", Symbol: "S"},
				currency.SRD: cldr.Currency{DisplayName: "Doolarka Surinamees", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "Bowndka Koonfurta Suudaan", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "Dobra Sao Tome & Birinsibi", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "Bowndka Suuriya", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeenida iswaasi", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "Baatka Taylaan", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "Somoonida Tajikistan", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Manaata Turkmenistan", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "Dinaarka Tunisiya", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "Ba’angada Tonga", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "Liirada Turkiga", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "Doolarka Tirinidad iyo Tobago", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "Doolarka Taywaanta Cusub", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Shilingka Tansaaniya", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "Hirfiniyada Yukreeyn", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Shilingka Yugandha", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Doolarka Mareeykanka", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "Beesada Urugway", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "Somta Usbekistan", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "Bolifar Fenesuala (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Bolifarada Fenesuwela", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "Dongta Fitnaam", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Fatu Fanuatu", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "Tala Samao", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "Faranka CFA ee Bartamaha Afrika", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Doolarka Iist Kaaribyan", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Faranka CFA Galbeedka Afrika", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "Faranka CFP", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Lacag aan la aqoon", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Riyaalka Yemen", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Randka Koonfur Afrika", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "Kawajada Sambiya", Symbol: "ZK"},
			},
		},
		Languages: cldr.Languages{
			language.AF:      "Afrikaanka",
			language.AGQ:     "Ageem",
			language.AK:      "Akan",
			language.AM:      "Axmaari",
			language.AR:      "Carabi",
			language.AR_001:  "Carabiga rasmiga ah",
			language.AS:      "Asaamiis",
			language.ASA:     "Asu",
			language.AST:     "Astuuriyaan",
			language.AZ:      "Aseeri",
			language.BAS:     "Basaa",
			language.BE:      "Beleruusiyaan",
			language.BEM:     "Bemba",
			language.BEZ:     "Bena",
			language.BG:      "Bulgeeriyaan",
			language.BM:      "Bambaara",
			language.BN:      "Bangladesh",
			language.BO:      "Tibeetaan",
			language.BR:      "Biriton",
			language.BRX:     "Bodo",
			language.BS:      "Bosniyaan",
			language.CA:      "Katalaan",
			language.CCP:     "Jakma",
			language.CE:      "Jejen",
			language.CEB:     "Sebuano",
			language.CGG:     "Jiga",
			language.CHR:     "Jerookee",
			language.CKB:     "Bartamaha Kurdish",
			language.CO:      "Korsikan",
			language.CS:      "Jeeg",
			language.CU:      "Kaniisadda Islaafik",
			language.CY:      "Welsh",
			language.DA:      "Dhaanish",
			language.DAV:     "Taiita",
			language.DE:      "Jarmal",
			language.DE_CH:   "Jarmal (Iswiiserlaand)",
			language.DJE:     "Sarma",
			language.DSB:     "Soorbiyaanka Hoose",
			language.DUA:     "Duaala",
			language.DYO:     "Joola-Foonyi",
			language.DZ:      "D’zongqa",
			language.EBU:     "Embuu",
			language.EE:      "Eewe",
			language.EL:      "Giriik",
			language.EN:      "Ingiriisi",
			language.EN_AU:   "Ingiriis Austaraaliyaan",
			language.EN_CA:   "Ingiriis Kanadiyaan",
			language.EN_GB:   "Ingiriiska Boqortooyada Midooday",
			language.EN_US:   "Ingiriisi (US)",
			language.EO:      "Isberaanto",
			language.ES:      "Isbaanish",
			language.ES_419:  "Isbaanishka Laatiin Ameerika",
			language.ES_ES:   "Isbaanish (Isbayn)",
			language.ET:      "Istooniyaan",
			language.EU:      "Basquu",
			language.EWO:     "Eewondho",
			language.FA:      "Faarisi",
			language.FA_AF:   "Faarsi",
			language.FF:      "Fuulah",
			language.FI:      "Finishka",
			language.FIL:     "Tagalog",
			language.FO:      "Farowsi",
			language.FR:      "Faransiis",
			language.FR_CA:   "Faransiiska Kanada",
			language.FR_CH:   "Faransiis (Iswiiserlaand)",
			language.FRC:     "Faransiiska Cajun",
			language.FUR:     "Firiyuuliyaan",
			language.FY:      "Firiisiyan Galbeed",
			language.GA:      "Ayrish",
			language.GD:      "Iskot Giilik",
			language.GL:      "Galiisiyaan",
			language.GN:      "Guraani",
			language.GSW:     "Jarmal Iswiis",
			language.GU:      "Gujaraati",
			language.GUZ:     "Guusii",
			language.GV:      "Mankis",
			language.HA:      "Hawsa",
			language.HAW:     "Hawaay",
			language.HE:      "Cibraani",
			language.HI:      "Hindi",
			language.HMN:     "Hamong",
			language.HR:      "Koro’eeshiyaan",
			language.HSB:     "Sorobiyaanka Sare",
			language.HT:      "Heeytiyaan Karawle",
			language.HU:      "Hangariyaan",
			language.HY:      "Armeeniyaan",
			language.IA:      "Interlinguwa",
			language.ID:      "Indunusiyaan",
			language.IE:      "Interlingue",
			language.IG:      "Igbo",
			language.II:      "Sijuwan Yi",
			language.IS:      "Ayslandays",
			language.IT:      "Talyaani",
			language.JA:      "Jabaaniis",
			language.JGO:     "Ingoomba",
			language.JMC:     "Chaga",
			language.JV:      "Jafaaniis",
			language.KA:      "Joorijiyaan",
			language.KAB:     "Kabayle",
			language.KAM:     "Kaamba",
			language.KDE:     "Kimakonde",
			language.KEA:     "Kabuferdiyanu",
			language.KHQ:     "Koyra Jiini",
			language.KI:      "Kikuuyu",
			language.KK:      "Kasaaq",
			language.KKJ:     "Kaako",
			language.KL:      "Kalaallisuut",
			language.KLN:     "Kalenjiin",
			language.KM:      "Kamboodhian",
			language.KN:      "Kannadays",
			language.KO:      "Kuuriyaan",
			language.KOK:     "Konkani",
			language.KS:      "Kaashmiir",
			language.KSB:     "Shambaala",
			language.KSF:     "Bafiya",
			language.KSH:     "Kologniyaan",
			language.KU:      "Kurdishka",
			language.KW:      "Kornish",
			language.KY:      "Kirgiis",
			language.LA:      "Laatiin",
			language.LAG:     "Laangi",
			language.LB:      "Luksaamboorgish",
			language.LG:      "Gandha",
			language.LKT:     "Laakoota",
			language.LN:      "Lingala",
			language.LO:      "Lao",
			language.LOU:     "Louisiana Creole",
			language.LRC:     "Koonfurta Luuri",
			language.LT:      "Lituwaanays",
			language.LU:      "Luuba-kataanga",
			language.LUO:     "Luwada",
			language.LUY:     "Luhya",
			language.LV:      "Laatfiyaan",
			language.MAI:     "Dadka Maithili",
			language.MAS:     "Masaay",
			language.MER:     "Meeru",
			language.MFE:     "Moorisayn",
			language.MG:      "Malagaasi",
			language.MGH:     "Makhuwa",
			language.MGO:     "Meetaa",
			language.MI:      "Maaoori",
			language.MK:      "Masadooniyaan",
			language.ML:      "Malayalam",
			language.MN:      "Mangooli",
			language.MR:      "Maarati",
			language.MS:      "Malaay",
			language.MT:      "Maltiis",
			language.MUA:     "Miyundhaang",
			language.MUL:     "Luuqado kala duwan",
			language.MY:      "Burmese",
			language.MZN:     "Masanderaani",
			language.NAQ:     "Nama",
			language.NB:      "Noorwijiyaan Bokma",
			language.ND:      "Indhebeele",
			language.NDS:     "Jarmal Hooseeya",
			language.NE:      "Nebaali",
			language.NL:      "Holandays",
			language.NL_BE:   "Af faleemi",
			language.NMG:     "Kuwaasiyo",
			language.NN:      "Nowrwejiyan (naynoroski)",
			language.NNH:     "Ingiyembuun",
			language.NO:      "Af Noorwiijiyaan",
			language.NUS:     "Nuweer",
			language.NY:      "Inyaanja",
			language.NYN:     "Inyankoole",
			language.OC:      "Okitaan",
			language.OM:      "Oromo",
			language.OR:      "Oodhiya",
			language.OS:      "Oseetic",
			language.PA:      "Bunjaabi",
			language.PL:      "Boolish",
			language.PRG:     "Brashiyaanki Hore",
			language.PS:      "Bashtuu",
			language.PT:      "Boortaqiis",
			language.PT_BR:   "Boortaqiiska Baraasiil",
			language.PT_PT:   "Boortaqiis (Boortuqaal)",
			language.QU:      "Quwejuwa",
			language.RM:      "Romaanis",
			language.RN:      "Rundhi",
			language.RO:      "Romanka",
			language.ROF:     "Rombo",
			language.ROOT:    "Xidid",
			language.RU:      "Ruush",
			language.RW:      "Ruwaandha",
			language.RWK:     "Raawa",
			language.SA:      "Sanskrit",
			language.SAH:     "Saaqa",
			language.SAQ:     "Sambuuru",
			language.SBP:     "Sangu",
			language.SD:      "Siindhi",
			language.SE:      "Koonfurta Saami",
			language.SEH:     "Seena",
			language.SES:     "Koyraboro Seenni",
			language.SG:      "Sango",
			language.SH:      "Serbiyaan",
			language.SHI:     "Shilha",
			language.SI:      "Sinhaleys",
			language.SK:      "Isloofaak",
			language.SL:      "Islofeeniyaan",
			language.SM:      "Samowan",
			language.SMN:     "Inaari Saami",
			language.SN:      "Shoona",
			language.SO:      "Soomaali",
			language.SQ:      "Albeeniyaan",
			language.SR:      "Seerbiyaan",
			language.ST:      "Sesooto",
			language.SU:      "Suudaaniis",
			language.SV:      "Swiidhis",
			language.SW:      "Sawaaxili",
			language.TA:      "Tamiil",
			language.TE:      "Teluugu",
			language.TEO:     "Teeso",
			language.TG:      "Taajik",
			language.TH:      "Taaylandays",
			language.TI:      "Tigrinya",
			language.TK:      "Turkumaanish",
			language.TLH:     "Kiligoon",
			language.TO:      "Toongan",
			language.TR:      "Turkish",
			language.TT:      "Taatar",
			language.TW:      "Tiwiyan",
			language.TWQ:     "Tasaawaq",
			language.TZM:     "Bartamaha Atlaas Tamasayt",
			language.UG:      "Uighur",
			language.UK:      "Yukreeniyaan",
			language.UND:     "Af aan la aqoon ama aan sax ahayn",
			language.UR:      "Urduu",
			language.UZ:      "Usbakis",
			language.VAI:     "Faayi",
			language.VI:      "Fiitnaamays",
			language.VO:      "Folabuuk",
			language.VUN:     "Fuunjo",
			language.WAE:     "Walseer",
			language.WO:      "Woolof",
			language.XH:      "Hoosta",
			language.XOG:     "Sooga",
			language.YAV:     "Yaangbeen",
			language.YI:      "Yadhish",
			language.YO:      "Yoruuba",
			language.YUE:     "Kantoneese",
			language.ZGH:     "Morokaanka Tamasayt Rasmiga",
			language.ZH:      "Shiinaha Mandarin",
			language.ZH_HANS: "Shiinaha Rasmiga ah",
			language.ZH_HANT: "Shiinahii Hore",
			language.ZU:      "Zuulu",
			language.ZXX:     "Luuqad Looma Hayo",
		},
		Territories: cldr.Territories{
			territory.V_001: "Dunida",
			territory.V_002: "Afrika",
			territory.V_003: "Waqooyi Ameerika",
			territory.V_005: "Koonfur Ameerika",
			territory.V_009: "Osheeniya",
			territory.V_011: "Galbeeka Afrika",
			territory.V_013: "Bartamaha Ameerika",
			territory.V_014: "Afrikada Bari",
			territory.V_015: "Waqooyiga Afrika",
			territory.V_017: "Afrikada Dhexe",
			territory.V_018: "Afrikada Koonfureed",
			territory.V_019: "Ameerikaas",
			territory.V_021: "Waqooyiga Ameerika",
			territory.V_029: "Karibiyaan",
			territory.V_030: "Aasiyada Bari",
			territory.V_034: "Aasiyada Koonfureed",
			territory.V_035: "Aasiyada Koonfur-galbeed",
			territory.V_039: "Yurubta Koonfureed",
			territory.V_053: "Austraalaasiya",
			territory.V_054: "Melaneesiya",
			territory.V_057: "Gobolka Aasiyada yar",
			territory.V_061: "Booliyneesiya",
			territory.V_142: "Aasiya",
			territory.V_143: "Bartamaha Aasiya",
			territory.V_145: "Aasiyada Galbeed",
			territory.V_150: "Yurub",
			territory.V_151: "Yurubta Bari",
			territory.V_154: "Yurubta Waqooyi",
			territory.V_155: "Yurubta Galbeed",
			territory.V_202: "Afrikada ka hooseysa Saxaraha",
			territory.V_419: "Laatiin Ameerika",
			territory.AC:    "Jasiiradda Asensiyoon",
			territory.AD:    "Andora",
			territory.AE:    "Imaaraadka Carabta ee Midoobay",
			territory.AF:    "Afgaanistaan",
			territory.AG:    "Antigua & Barbuuda",
			territory.AI:    "Anguula",
			territory.AL:    "Albaaniya",
			territory.AM:    "Armeeniya",
			territory.AO:    "Angoola",
			territory.AQ:    "Antaarktika",
			territory.AR:    "Arjentiina",
			territory.AS:    "Samowa Ameerika",
			territory.AT:    "Awsteriya",
			territory.AU:    "Awstaraaliya",
			territory.AW:    "Aruba",
			territory.AX:    "Jasiiradda Aland",
			territory.AZ:    "Asarbajan",
			territory.BA:    "Boosniya & Harsegofina",
			territory.BB:    "Baarbadoos",
			territory.BD:    "Bangladesh",
			territory.BE:    "Biljam",
			territory.BF:    "Burkiina Faaso",
			territory.BG:    "Bulgaariya",
			territory.BH:    "Baxreyn",
			territory.BI:    "Burundi",
			territory.BJ:    "Biniin",
			territory.BL:    "St. Baathelemiy",
			territory.BM:    "Barmuuda",
			territory.BN:    "Buruneeya",
			territory.BO:    "Boliifiya",
			territory.BQ:    "Karibiyaan Nadarlands",
			territory.BR:    "Baraasiil",
			territory.BS:    "Bahaamas",
			territory.BT:    "Buutan",
			territory.BV:    "Buufet Island",
			territory.BW:    "Botuswaana",
			territory.BY:    "Belarus",
			territory.BZ:    "Beliis",
			territory.CA:    "Kanada",
			territory.CC:    "Jasiiradda Kookoos",
			territory.CD:    "Jamhuuriyadda Dimuqaadiga Kongo",
			territory.CF:    "Jamhuuriyadda Afrikada Dhexe",
			territory.CG:    "Jamhuuriyadda Kongo",
			territory.CH:    "Swiiserlaand",
			territory.CI:    "Aayforikoost",
			territory.CK:    "Jasiiradda Kook",
			territory.CL:    "Jili",
			territory.CM:    "Kaameruun",
			territory.CN:    "Shiinaha",
			territory.CO:    "Koloombiya",
			territory.CP:    "Jasiiradda Kilibarton",
			territory.CR:    "Kosta Riika",
			territory.CU:    "Kuuba",
			territory.CV:    "Jasiiradda Kayb Faarde",
			territory.CW:    "Kurakaaw",
			territory.CX:    "Jasiiradda Kirismas",
			territory.CY:    "Qubrus",
			territory.CZ:    "Jamhuuriyadda Jek",
			territory.DE:    "Jarmal",
			territory.DG:    "Diyeego Karsiya",
			territory.DJ:    "Jabuuti",
			territory.DK:    "Denmark",
			territory.DM:    "Dominika",
			territory.DO:    "Jamhuuriyaddda Dominika",
			territory.DZ:    "Aljeeriya",
			territory.EA:    "Seyuta & Meliila",
			territory.EC:    "Ikuwadoor",
			territory.EE:    "Estooniya",
			territory.EG:    "Masar",
			territory.EH:    "Saxaraha Galbeed",
			territory.ER:    "Eritreeya",
			territory.ES:    "Isbeyn",
			territory.ET:    "Itoobiya",
			territory.EU:    "Midowga Yurub",
			territory.EZ:    "Yurusoon",
			territory.FI:    "Finland",
			territory.FJ:    "Fiji",
			territory.FK:    "Jasiiradaha Fookland",
			territory.FM:    "Mikroneesiya",
			territory.FO:    "Jasiiradda Faroo",
			territory.FR:    "Faransiis",
			territory.GA:    "Gaaboon",
			territory.GB:    "UK",
			territory.GD:    "Giriinaada",
			territory.GE:    "Joorjiya",
			territory.GF:    "Faransiis Gini",
			territory.GG:    "Guurnsey",
			territory.GH:    "Gaana",
			territory.GI:    "Gibraltar",
			territory.GL:    "Greenland",
			territory.GM:    "Gambiya",
			territory.GN:    "Gini",
			territory.GP:    "Guadeluub",
			territory.GQ:    "Ekuwatooriyal Gini",
			territory.GR:    "Giriig",
			territory.GS:    "Jasiiradda Joorjiyada Koonfureed & Sandwij",
			territory.GT:    "Guwaatamaala",
			territory.GU:    "Guaam",
			territory.GW:    "Gini-Bisaaw",
			territory.GY:    "Guyana",
			territory.HK:    "Hong Kong",
			territory.HM:    "Jasiiradda Haad & MakDonald",
			territory.HN:    "Honduras",
			territory.HR:    "Korweeshiya",
			territory.HT:    "Haiti",
			territory.HU:    "Hangari",
			territory.IC:    "Jasiiradda Kanari",
			territory.ID:    "Indoneesiya",
			territory.IE:    "Ayrlaand",
			territory.IL:    "Israaʼiil",
			territory.IM:    "Jasiiradda Isle of Man",
			territory.IN:    "Hindiya",
			territory.IO:    "Dhul xadeedka Badweynta Hindiya ee Biritishka",
			territory.IQ:    "Ciraaq",
			territory.IR:    "Iiraan",
			territory.IS:    "Ayslaand",
			territory.IT:    "Talyaani",
			territory.JE:    "Jaarsey",
			territory.JM:    "Jamaaika",
			territory.JO:    "Urdun",
			territory.JP:    "Jabaan",
			territory.KE:    "Kenya",
			territory.KG:    "Kirgistaan",
			territory.KH:    "Kamboodiya",
			territory.KI:    "Kiribati",
			territory.KM:    "Komooros",
			territory.KN:    "St. Kitts & Nefis",
			territory.KP:    "Kuuriyada Waqooyi",
			territory.KR:    "Kuuriyada Koonfureed",
			territory.KW:    "Kuwayt",
			territory.KY:    "Cayman Islands",
			territory.KZ:    "Kasaakhistaan",
			territory.LA:    "Laos",
			territory.LB:    "Lubnaan",
			territory.LC:    "St. Lusia",
			territory.LI:    "Liyjtensteyn",
			territory.LK:    "Sirilaanka",
			territory.LR:    "Laybeeriya",
			territory.LS:    "Losooto",
			territory.LT:    "Lituweeniya",
			territory.LU:    "Luksemboorg",
			territory.LV:    "Latfiya",
			territory.LY:    "Liibya",
			territory.MA:    "Morooko",
			territory.MC:    "Moonako",
			territory.MD:    "Moldofa",
			territory.ME:    "Moontenegro",
			territory.MF:    "St. Maartin",
			territory.MG:    "Madagaskar",
			territory.MH:    "Jasiiradda Maarshal",
			territory.MK:    "Masedooniya Waqooyi",
			territory.ML:    "Maali",
			territory.MM:    "Miyanmar",
			territory.MN:    "Mongooliya",
			territory.MO:    "Makaaw",
			territory.MP:    "Jasiiradda Waqooyiga Mariaana",
			territory.MQ:    "Maartinik",
			territory.MR:    "Muritaaniya",
			territory.MS:    "Montserrat",
			territory.MT:    "Maalta",
			territory.MU:    "Mawrishiyaas",
			territory.MV:    "Maaldiqeen",
			territory.MW:    "Malaawi",
			territory.MX:    "Meksiko",
			territory.MY:    "Malaysia",
			territory.MZ:    "Musambiik",
			territory.NA:    "Namiibiya",
			territory.NC:    "Jasiiradda Niyuu Kaledooniya",
			territory.NE:    "Nayjer",
			territory.NF:    "Jasiiradda Noorfolk",
			territory.NG:    "Nayjeeriya",
			territory.NI:    "Nikaraaguwa",
			territory.NL:    "Nederlaands",
			territory.NO:    "Noorweey",
			territory.NP:    "Nebaal",
			territory.NR:    "Nauru",
			territory.NU:    "Niue",
			territory.NZ:    "Niyuusiilaand",
			territory.OM:    "Cumaan",
			territory.PA:    "Baanama",
			territory.PE:    "Beeru",
			territory.PF:    "Booliyneesiya Faransiiska",
			territory.PG:    "Babua Niyuu Gini",
			territory.PH:    "Filibiin",
			territory.PK:    "Bakistaan",
			territory.PL:    "Booland",
			territory.PM:    "Saint Pierre and Miquelon",
			territory.PN:    "Bitkairn",
			territory.PR:    "Bueerto Riiko",
			territory.PS:    "Falastiin",
			territory.PT:    "Bortugaal",
			territory.PW:    "Balaaw",
			territory.PY:    "Baraguaay",
			territory.QA:    "Qadar",
			territory.QO:    "Dhulxeebeedka Osheeniya",
			territory.RE:    "Riyuuniyon",
			territory.RO:    "Rumaaniya",
			territory.RS:    "Seerbiya",
			territory.RU:    "Ruush",
			territory.RW:    "Ruwanda",
			territory.SA:    "Sacuudi Carabiya",
			territory.SB:    "Jasiiradda Solomon",
			territory.SC:    "Sishelis",
			territory.SD:    "Suudaan",
			territory.SE:    "Iswidhan",
			territory.SG:    "Singaboor",
			territory.SH:    "Saint Helena",
			territory.SI:    "Islofeeniya",
			territory.SJ:    "Jasiiradda Sfaldbaad & Jaan Mayen",
			territory.SK:    "Islofaakiya",
			territory.SL:    "Siraaliyoon",
			territory.SM:    "San Marino",
			territory.SN:    "Sinigaal",
			territory.SO:    "Soomaaliya",
			territory.SR:    "Surineym",
			territory.SS:    "Koonfur Suudaan",
			territory.ST:    "Sao Tome & Birincibal",
			territory.SV:    "El Salfadoor",
			territory.SX:    "Siint Maarteen",
			territory.SY:    "Suuriya",
			territory.SZ:    "Iswaasilaan",
			territory.TA:    "Tiristan da Kunha",
			territory.TC:    "Turks & Kaikos Island",
			territory.TD:    "Jaad",
			territory.TF:    "Dhul xadeedka Koonfureed ee Faransiiska",
			territory.TG:    "Toogo",
			territory.TH:    "Taylaand",
			territory.TJ:    "Tajikistan",
			territory.TK:    "Tokelaaw",
			territory.TL:    "Timoor Bari",
			territory.TM:    "Turkmenistan",
			territory.TN:    "Tuniisiya",
			territory.TO:    "Tonga",
			territory.TR:    "Turki",
			territory.TT:    "Tirinidaad & Tobago",
			territory.TV:    "Tufaalu",
			territory.TW:    "Taywaan",
			territory.TZ:    "Tansaaniya",
			territory.UA:    "Yukrayn",
			territory.UG:    "Ugaanda",
			territory.UM:    "Jasiiradaha ka baxsan Maraykanka",
			territory.UN:    "Qaramada Midoobay",
			territory.US:    "Maraykanka",
			territory.UY:    "Uruguwaay",
			territory.UZ:    "Uusbakistaan",
			territory.VA:    "Faatikaan",
			territory.VC:    "St. Finsent & Girenadiins",
			territory.VE:    "Fenisuweela",
			territory.VG:    "Biritish Farjin Island",
			territory.VI:    "U.S Fargin Island",
			territory.VN:    "Fiyetnaam",
			territory.VU:    "Fanuaatu",
			territory.WF:    "Walis & Futuna",
			territory.WS:    "Samoowa",
			territory.XA:    "Shigshiga",
			territory.XB:    "Pseudo-Bidi",
			territory.XK:    "Koosofo",
			territory.YE:    "Yaman",
			territory.YT:    "Mayotte",
			territory.ZA:    "Koonfur Afrika",
			territory.ZM:    "Saambiya",
			territory.ZW:    "Simbaabwe",
			territory.ZZ:    "Gobol aan la aqoonin amase aan saxnayn",
		},
	}
}
