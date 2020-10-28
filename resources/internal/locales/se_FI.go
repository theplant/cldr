// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_se_FI() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "se_FI",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd.MM.y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "{0} GMT"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ođđj", Feb: "guov", Mar: "njuk", Apr: "cuoŋ", May: "mies", Jun: "geas", Jul: "suoi", Aug: "borg", Sep: "čakč", Oct: "golg", Nov: "skáb", Dec: "juov"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "O", Feb: "G", Mar: "N", Apr: "C", May: "M", Jun: "G", Jul: "S", Aug: "B", Sep: "Č", Oct: "G", Nov: "S", Dec: "J"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "ođđajagemánnu", Feb: "guovvamánnu", Mar: "njukčamánnu", Apr: "cuoŋománnu", May: "miessemánnu", Jun: "geassemánnu", Jul: "suoidnemánnu", Aug: "borgemánnu", Sep: "čakčamánnu", Oct: "golggotmánnu", Nov: "skábmamánnu", Dec: "juovlamánnu"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "so", Mon: "má", Tue: "di", Wed: "ga", Thu: "du", Fri: "be", Sat: "lá"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "D", Wed: "G", Thu: "D", Fri: "B", Sat: "L"}, Short: cldr.CalendarDayFormatNameValue{Sun: "so", Mon: "má", Tue: "di", Wed: "ga", Thu: "du", Fri: "be", Sat: "lá"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "sotnabeaivi", Mon: "mánnodat", Tue: "disdat", Wed: "gaskavahkku", Thu: "duorastat", Fri: "bearjadat", Sat: "lávvordat"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "ib", PM: "eb"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "ib", PM: "eb"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "ib", PM: "eb"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "−", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "0 dt", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.AOA: cldr.Currency{DisplayName: "", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "", Symbol: "CA$"},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "Dkr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.EGP: cldr.Currency{DisplayName: "", Symbol: "E£"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "suoma márkki", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "", Symbol: "GB£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "Ikr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "", Symbol: "JP¥"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.MGA: cldr.Currency{DisplayName: "", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MUR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.NAD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "norgga kruvdno", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "", Symbol: "NZ$"},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "", Symbol: "₲"},
				currency.RON: cldr.Currency{DisplayName: "", Symbol: "lei"},
				currency.RUB: cldr.Currency{DisplayName: "", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "", Symbol: "RF"},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SEK: cldr.Currency{DisplayName: "ruoŧŧa kruvdno", Symbol: "Skr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.USD: cldr.Currency{DisplayName: "", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "uns silba", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "uns golli", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Languages: cldr.Languages{
			language.ACE:     "ačehgiella",
			language.AF:      "afrikánsagiella",
			language.AN:      "aragoniagiella",
			language.ANG:     "boares eaŋgalasgiella",
			language.AR:      "arábagiella",
			language.AR_001:  "standárda arábagiella",
			language.AST:     "asturiagiella",
			language.BE:      "vilgesruoššagiella",
			language.BG:      "bulgáriagiella",
			language.BN:      "bengalagiella",
			language.BO:      "tibetagiella",
			language.BR:      "bretonagiella",
			language.BS:      "bosniagiella",
			language.CA:      "katalánagiella",
			language.CHM:     "marigiella",
			language.CO:      "corsicagiella",
			language.CS:      "čeahkagiella",
			language.CY:      "kymragiella",
			language.DA:      "dánskkagiella",
			language.DE:      "duiskkagiella",
			language.DE_AT:   "nuortariikkalaš duiskkagiella",
			language.DE_CH:   "šveicalaš duiskkagiella",
			language.DV:      "divehigiella",
			language.DZ:      "dzongkhagiella",
			language.EL:      "greikkagiella",
			language.EN:      "eaŋgalsgiella",
			language.EN_AU:   "austrálialaš eaŋgalsgiella",
			language.EN_CA:   "kanádalaš eaŋgalsgiella",
			language.EN_GB:   "brihttalaš eaŋgalsgiella",
			language.EN_US:   "amerihkálaš eaŋgalsgiella",
			language.ES:      "spánskkagiella",
			language.ES_419:  "latiinna-amerihkalaš spánskkagiella",
			language.ES_ES:   "espánjalaš spánskkagiella",
			language.ES_MX:   "meksikolaš spánskkagiella",
			language.ET:      "esttegiella",
			language.FA:      "persijagiella",
			language.FI:      "suomagiella",
			language.FIL:     "filippiinnagiella",
			language.FJ:      "fižigiella",
			language.FO:      "fearagiella",
			language.FR:      "fránskkagiella",
			language.FR_CA:   "kanádalaš fránskkagiella",
			language.FR_CH:   "šveicalaš fránskkagiella",
			language.FY:      "oarjifriisagiella",
			language.GA:      "iirragiella",
			language.GU:      "gujaratagiella",
			language.GV:      "manksgiella",
			language.HA:      "haussagiella",
			language.HAW:     "hawaiigiella",
			language.HI:      "hindigiella",
			language.HR:      "kroátiagiella",
			language.HT:      "haitigiella",
			language.HU:      "ungárgiella",
			language.HY:      "armenagiella",
			language.ID:      "indonesiagiella",
			language.IS:      "islánddagiella",
			language.IT:      "itáliagiella",
			language.JA:      "japánagiella",
			language.JV:      "javagiella",
			language.KA:      "georgiagiella",
			language.KK:      "kazakhgiella",
			language.KM:      "kambožagiella",
			language.KO:      "koreagiella",
			language.KRL:     "gárjilgiella",
			language.KU:      "kurdigiella",
			language.KV:      "komigiella",
			language.KW:      "kornagiella",
			language.LA:      "láhtengiella",
			language.LB:      "luxemburggagiella",
			language.LO:      "laogiella",
			language.LT:      "liettuvagiella",
			language.LV:      "látviagiella",
			language.MDF:     "mokšagiella",
			language.MI:      "maorigiella",
			language.MK:      "makedoniagiella",
			language.MN:      "mongoliagiella",
			language.MT:      "maltagiella",
			language.MY:      "burmagiella",
			language.MYV:     "ersagiella",
			language.NB:      "girjedárogiella",
			language.NE:      "nepalagiella",
			language.NL:      "hollánddagiella",
			language.NL_BE:   "belgialaš hollánddagiella",
			language.NN:      "ođđadárogiella",
			language.NO:      "dárogiella",
			language.OC:      "oksitánagiella",
			language.PA:      "panjabagiella",
			language.PL:      "polskkagiella",
			language.PT:      "portugálagiella",
			language.PT_BR:   "brasilialaš portugálagiella",
			language.PT_PT:   "portugálalaš portugálagiella",
			language.RM:      "romanšgiella",
			language.RO:      "romániagiella",
			language.RO_MD:   "moldávialaš romániagiella",
			language.RU:      "ruoššagiella",
			language.SC:      "sardigiella",
			language.SCN:     "sisiliagiella",
			language.SE:      "davvisámegiella",
			language.SEL:     "selkupagiella",
			language.SH:      "serbokroatiagiella",
			language.SK:      "slovákiagiella",
			language.SL:      "slovenagiella",
			language.SM:      "samoagiella",
			language.SMA:     "lullisámegiella",
			language.SMJ:     "julevsámegiella",
			language.SMN:     "anárašgiella",
			language.SMS:     "nuortalašgiella",
			language.SQ:      "albánagiella",
			language.SR:      "serbiagiella",
			language.SV:      "ruoŧagiella",
			language.SWB:     "komoragiella",
			language.TH:      "thaigiella",
			language.TR:      "durkagiella",
			language.TY:      "tahitigiella",
			language.UDM:     "udmurtagiella",
			language.UK:      "ukrainagiella",
			language.UND:     "dovdameahttun giella",
			language.UR:      "urdugiella",
			language.VI:      "vietnamagiella",
			language.WA:      "vallonagiella",
			language.YUE:     "kantongiella",
			language.ZH:      "kiinnágiella",
			language.ZH_HANS: "álkes kiinnágiella",
			language.ZH_HANT: "árbevirolaš kiinnágiella",
		},
		Territories: cldr.Territories{
			territory.V_001: "Máilbmi",
			territory.V_002: "Afrihka",
			territory.V_003: "Davvi-Amerihká ja Gaska-Amerihká",
			territory.V_005: "Lulli-Amerihká",
			territory.V_009: "Oseania",
			territory.V_011: "Oarje-Afrihká",
			territory.V_013: "Gaska-Amerihká",
			territory.V_014: "Nuorta-Afrihká",
			territory.V_015: "Davvi-Afrihká",
			territory.V_017: "Gaska-Afrihká",
			territory.V_018: "Lulli-Afrihká",
			territory.V_019: "Amerihka",
			territory.V_021: "Davvi-Amerihká",
			territory.V_029: "Karibia",
			territory.V_030: "nuorta-Ásia",
			territory.V_034: "mátta-Ásia",
			territory.V_035: "mátta-nuorta-Ásia",
			territory.V_039: "mátta-Eurohpá",
			territory.V_053: "Austrália ja Ođđa-Selánda",
			territory.V_054: "Melanesia",
			territory.V_057: "Mikronesia guovlu",
			territory.V_061: "Polynesia",
			territory.V_142: "Ásia",
			territory.V_143: "gaska-Ásia",
			territory.V_145: "oarji-Ásia",
			territory.V_150: "Eurohpá",
			territory.V_151: "nuorta-Eurohpá",
			territory.V_154: "davvi-Eurohpá",
			territory.V_155: "oarji-Eurohpá",
			territory.V_419: "Latiinnalaš Amerihká",
			territory.AC:    "Ascension",
			territory.AD:    "Andorra",
			territory.AE:    "Ovttastuvvan Arábaemiráhtat",
			territory.AF:    "Afghanistan",
			territory.AG:    "Antigua ja Barbuda",
			territory.AI:    "Anguilla",
			territory.AL:    "Albánia",
			territory.AM:    "Armenia",
			territory.AO:    "Angola",
			territory.AQ:    "Antárktis",
			territory.AR:    "Argentina",
			territory.AS:    "Amerihká Samoa",
			territory.AT:    "Nuortariika",
			territory.AU:    "Austrália",
			territory.AW:    "Aruba",
			territory.AX:    "Ålánda",
			territory.AZ:    "Aserbaižan",
			territory.BA:    "Bosnia ja Hercegovina",
			territory.BB:    "Barbados",
			territory.BD:    "Bangladesh",
			territory.BE:    "Belgia",
			territory.BF:    "Burkina Faso",
			territory.BG:    "Bulgária",
			territory.BH:    "Bahrain",
			territory.BI:    "Burundi",
			territory.BJ:    "Benin",
			territory.BL:    "Saint Barthélemy",
			territory.BM:    "Bermuda",
			territory.BN:    "Brunei",
			territory.BO:    "Bolivia",
			territory.BQ:    "Vuolleeatnamat Karibe",
			territory.BR:    "Brasil",
			territory.BS:    "Bahamas",
			territory.BT:    "Bhutan",
			territory.BV:    "Bouvet-sullot",
			territory.BW:    "Botswana",
			territory.BY:    "Vilges-Ruošša",
			territory.BZ:    "Belize",
			territory.CA:    "Kanáda",
			territory.CC:    "Cocos-sullot",
			territory.CD:    "Kongo-Kinshasa",
			territory.CF:    "Gaska-Afrihká dásseváldi",
			territory.CG:    "Kongo-Brazzaville",
			territory.CH:    "Šveica",
			territory.CI:    "Côte d’Ivoire",
			territory.CK:    "Cook-sullot",
			territory.CL:    "Čiile",
			territory.CM:    "Kamerun",
			territory.CN:    "Kiinná",
			territory.CO:    "Kolombia",
			territory.CP:    "Clipperton-sullot",
			territory.CR:    "Costa Rica",
			territory.CU:    "Kuba",
			territory.CV:    "Kap Verde",
			territory.CW:    "Curaçao",
			territory.CX:    "Juovllat-sullot",
			territory.CY:    "Kypros",
			territory.CZ:    "Čeahkka",
			territory.DE:    "Duiska",
			territory.DG:    "Diego Garcia",
			territory.DJ:    "Djibouti",
			territory.DK:    "Dánmárku",
			territory.DM:    "Dominica",
			territory.DO:    "Dominikána dásseváldi",
			territory.DZ:    "Algeria",
			territory.EA:    "Ceuta ja Melilla",
			territory.EC:    "Ecuador",
			territory.EE:    "Estlánda",
			territory.EG:    "Egypta",
			territory.EH:    "Oarje-Sahára",
			territory.ER:    "Eritrea",
			territory.ES:    "Spánia",
			territory.ET:    "Etiopia",
			territory.EU:    "Eurohpa Uniovdna",
			territory.EZ:    "Euroavádat",
			territory.FI:    "Suopma",
			territory.FJ:    "Fijisullot",
			territory.FK:    "Falklandsullot",
			territory.FM:    "Mikronesia",
			territory.FO:    "Fearsullot",
			territory.FR:    "Frankriika",
			territory.GA:    "Gabon",
			territory.GB:    "Stuorra-Británnia",
			territory.GD:    "Grenada",
			territory.GE:    "Georgia",
			territory.GF:    "Frankriikka Guayana",
			territory.GG:    "Guernsey",
			territory.GH:    "Ghana",
			territory.GI:    "Gibraltar",
			territory.GL:    "Kalaallit Nunaat",
			territory.GM:    "Gámbia",
			territory.GN:    "Guinea",
			territory.GP:    "Guadeloupe",
			territory.GQ:    "Ekvatoriála Guinea",
			territory.GR:    "Greika",
			territory.GS:    "Lulli Georgia ja Lulli Sandwich-sullot",
			territory.GT:    "Guatemala",
			territory.GU:    "Guam",
			territory.GW:    "Guinea-Bissau",
			territory.GY:    "Guyana",
			territory.HK:    "Hongkong",
			territory.HM:    "Heard- ja McDonald-sullot",
			territory.HN:    "Honduras",
			territory.HR:    "Kroátia",
			territory.HT:    "Haiti",
			territory.HU:    "Ungár",
			territory.IC:    "Kanáriasullot",
			territory.ID:    "Indonesia",
			territory.IE:    "Irlánda",
			territory.IL:    "Israel",
			territory.IM:    "Mann-sullot",
			territory.IN:    "India",
			territory.IQ:    "Irak",
			territory.IR:    "Iran",
			territory.IS:    "Islánda",
			territory.IT:    "Itália",
			territory.JE:    "Jersey",
			territory.JM:    "Jamaica",
			territory.JO:    "Jordánia",
			territory.JP:    "Japána",
			territory.KE:    "Kenia",
			territory.KG:    "Kirgisistan",
			territory.KH:    "Kamboža",
			territory.KI:    "Kiribati",
			territory.KM:    "Komoros",
			territory.KN:    "Saint Kitts ja Nevis",
			territory.KP:    "Davvi-Korea",
			territory.KR:    "Mátta-Korea",
			territory.KW:    "Kuwait",
			territory.KY:    "Cayman-sullot",
			territory.KZ:    "Kasakstan",
			territory.LA:    "Laos",
			territory.LB:    "Libanon",
			territory.LC:    "Saint Lucia",
			territory.LI:    "Liechtenstein",
			territory.LK:    "Sri Lanka",
			territory.LR:    "Liberia",
			territory.LS:    "Lesotho",
			territory.LT:    "Lietuva",
			territory.LU:    "Luxembourg",
			territory.LV:    "Látvia",
			territory.LY:    "Libya",
			territory.MA:    "Marokko",
			territory.MC:    "Monaco",
			territory.MD:    "Moldávia",
			territory.ME:    "Montenegro",
			territory.MF:    "Frankriikka Saint Martin",
			territory.MG:    "Madagaskar",
			territory.MH:    "Marshallsullot",
			territory.ML:    "Mali",
			territory.MM:    "Burma",
			territory.MN:    "Mongolia",
			territory.MO:    "Makáo",
			territory.MP:    "Davvi-Mariánat",
			territory.MQ:    "Martinique",
			territory.MR:    "Mauretánia",
			territory.MS:    "Montserrat",
			territory.MT:    "Málta",
			territory.MU:    "Mauritius",
			territory.MV:    "Malediivvat",
			territory.MW:    "Malawi",
			territory.MX:    "Meksiko",
			territory.MY:    "Malesia",
			territory.MZ:    "Mosambik",
			territory.NA:    "Namibia",
			territory.NC:    "Ođđa-Kaledonia",
			territory.NE:    "Niger",
			territory.NF:    "Norfolksullot",
			territory.NG:    "Nigeria",
			territory.NI:    "Nicaragua",
			territory.NL:    "Vuolleeatnamat",
			territory.NO:    "Norga",
			territory.NP:    "Nepal",
			territory.NR:    "Nauru",
			territory.NU:    "Niue",
			territory.NZ:    "Ođđa-Selánda",
			territory.OM:    "Oman",
			territory.PA:    "Panama",
			territory.PE:    "Peru",
			territory.PF:    "Frankriikka Polynesia",
			territory.PG:    "Papua-Ođđa-Guinea",
			territory.PH:    "Filippiinnat",
			territory.PK:    "Pakistan",
			territory.PL:    "Polen",
			territory.PM:    "Saint Pierre ja Miquelon",
			territory.PN:    "Pitcairn",
			territory.PR:    "Puerto Rico",
			territory.PS:    "Palestina",
			territory.PT:    "Portugála",
			territory.PW:    "Palau",
			territory.PY:    "Paraguay",
			territory.QA:    "Qatar",
			territory.RE:    "Réunion",
			territory.RO:    "Románia",
			territory.RS:    "Serbia",
			territory.RU:    "Ruošša",
			territory.RW:    "Rwanda",
			territory.SA:    "Saudi-Arábia",
			territory.SB:    "Salomon-sullot",
			territory.SC:    "Seychellsullot",
			territory.SD:    "Sudan",
			territory.SE:    "Ruoŧŧa",
			territory.SG:    "Singapore",
			territory.SH:    "Saint Helena",
			territory.SI:    "Slovenia",
			territory.SJ:    "Svalbárda ja Jan Mayen",
			territory.SK:    "Slovákia",
			territory.SL:    "Sierra Leone",
			territory.SM:    "San Marino",
			territory.SN:    "Senegal",
			territory.SO:    "Somália",
			territory.SR:    "Surinam",
			territory.SS:    "Máttasudan",
			territory.ST:    "São Tomé ja Príncipe",
			territory.SV:    "El Salvador",
			territory.SX:    "Vuolleeatnamat Saint Martin",
			territory.SY:    "Syria",
			territory.SZ:    "Svazieana",
			territory.TA:    "Tristan da Cunha",
			territory.TC:    "Turks ja Caicos-sullot",
			territory.TD:    "Chad",
			territory.TG:    "Togo",
			territory.TH:    "Thaieana",
			territory.TJ:    "Tažikistan",
			territory.TK:    "Tokelau",
			territory.TL:    "Nuorta-Timor",
			territory.TM:    "Turkmenistan",
			territory.TN:    "Tunisia",
			territory.TO:    "Tonga",
			territory.TR:    "Durka",
			territory.TT:    "Trinidad ja Tobago",
			territory.TV:    "Tuvalu",
			territory.TW:    "Taiwan",
			territory.TZ:    "Tanzánia",
			territory.UA:    "Ukraina",
			territory.UG:    "Uganda",
			territory.UN:    "Ovttastuvvan Našuvnnat",
			territory.US:    "USA",
			territory.UY:    "Uruguay",
			territory.UZ:    "Usbekistan",
			territory.VA:    "Vatikána",
			territory.VC:    "Saint Vincent ja Grenadine",
			territory.VE:    "Venezuela",
			territory.VG:    "Brittania Virgin-sullot",
			territory.VI:    "AOS Virgin-sullot",
			territory.VN:    "Vietnam",
			territory.VU:    "Vanuatu",
			territory.WF:    "Wallis ja Futuna",
			territory.WS:    "Samoa",
			territory.XK:    "Kosovo",
			territory.YE:    "Jemen",
			territory.YT:    "Mayotte",
			territory.ZA:    "Mátta-Afrihká",
			territory.ZM:    "Zambia",
			territory.ZW:    "Zimbabwe",
			territory.ZZ:    "dovdameahttun guovlu",
		},
	}
}
