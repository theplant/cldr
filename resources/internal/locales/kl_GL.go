// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_kl_GL() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "kl_GL",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE dd MMMM y", Long: "dd MMMM y", Medium: "dd MMM y", Short: "y-MM-dd"}, Time: cldr.CalendarDateFormat{Full: "HH.mm.ss zzzz", Long: "HH.mm.ss z", Medium: "HH.mm.ss", Short: "HH.mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan", Feb: "febr", Mar: "mar", Apr: "apr", May: "maj", Jun: "jun", Jul: "jul", Aug: "aug", Sep: "sept", Oct: "okt", Nov: "nov", Dec: "dec"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "januaari", Feb: "februaari", Mar: "marsi", Apr: "apriili", May: "maaji", Jun: "juuni", Jul: "juuli", Aug: "aggusti", Sep: "septembari", Oct: "oktobari", Nov: "novembari", Dec: "decembari"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "sap", Mon: "ata", Tue: "mar", Wed: "pin", Thu: "sis", Fri: "tal", Sat: "arf"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "A", Tue: "M", Wed: "P", Thu: "S", Fri: "T", Sat: "A"}, Short: cldr.CalendarDayFormatNameValue{Sun: "sap", Mon: "ata", Tue: "mar", Wed: "pin", Thu: "sis", Fri: "tal", Sat: "arf"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "sapaat", Mon: "ataasinngorneq", Tue: "marlunngorneq", Wed: "pingasunngorneq", Thu: "sisamanngorneq", Fri: "tallimanngorneq", Sat: "arfininngorneq"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "u.t.", PM: "u.k."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "ulloqeqqata-tungaa", PM: "ulloqeqqata-kingorna"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "−", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00;¤-#,##0.00", CurrencyAccounting: "", Percent: "#,##0\u00a0%"},
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
				currency.DKK: cldr.Currency{DisplayName: "danmarkimut koruuni", Symbol: "kr."},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.EGP: cldr.Currency{DisplayName: "", Symbol: "E£"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "", Symbol: "£"},
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
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
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
				currency.NOK: cldr.Currency{DisplayName: "norskit koruuni", Symbol: "Nkr"},
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
				currency.SEK: cldr.Currency{DisplayName: "svenskit koruuni", Symbol: "Skr"},
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
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Languages: cldr.Languages{
			language.AR:  "arabiamiusut",
			language.AZ:  "aserbajdsjaniskisut",
			language.BN:  "bengalimiutut",
			language.CS:  "tjekkiamut",
			language.DA:  "qallunaatut",
			language.DE:  "tyskisut",
			language.EN:  "tuluttut",
			language.EO:  "esperanto",
			language.ES:  "spanskisut",
			language.ET:  "estlandimiutut",
			language.FA:  "persiskisut",
			language.FI:  "finlandimiutut",
			language.FO:  "savalimmiutut",
			language.FR:  "franskisut",
			language.GA:  "irlandimiutut",
			language.HE:  "hebraimiutut",
			language.HI:  "hindimiutut",
			language.ID:  "indonesiamiutut",
			language.IS:  "islandimiusut",
			language.IT:  "italiamiutut",
			language.JA:  "japanimiusut",
			language.KL:  "kalaallisut",
			language.KO:  "koreamiusut",
			language.KU:  "kurdiskisut",
			language.LA:  "latiinerisut",
			language.LT:  "litauenimiutut",
			language.LV:  "letlandimiutut",
			language.MG:  "malagassiskisut",
			language.MI:  "maorimiutut",
			language.NL:  "hollandimiutut",
			language.PL:  "polenimiutut",
			language.PS:  "pashtomiutut",
			language.PT:  "portugalimiutut",
			language.RO:  "rumænimiutut",
			language.RU:  "russisut",
			language.SK:  "slovakimiusut",
			language.SV:  "svenskisut",
			language.SW:  "swahilimiutut",
			language.TH:  "thailandimiutut",
			language.TR:  "tyrkiskisut",
			language.UK:  "ukrainimiusut",
			language.UND: "(atorsinnaanngitsoq oqaatsit)",
			language.UR:  "urdumiutut",
			language.VI:  "vietnamimiusut",
			language.ZH:  "kineserisut",
		},
		Territories: cldr.Territories{
			territory.V_001: "silarsuaq",
			territory.V_002: "Afrika",
			territory.V_003: "Amerika Avannarleq",
			territory.V_005: "Amerika Kujalleq",
			territory.V_009: "Oceania",
			territory.V_011: "Afrika Killiit",
			territory.V_013: "America Qitiusumik",
			territory.V_014: "Afrika Kangilliit",
			territory.V_015: "Afrika Avannarleq",
			territory.V_017: "Afrika Qitiusumik",
			territory.V_018: "Afrika Kujalleq",
			territory.V_019: "Amerika",
			territory.V_030: "Asia Kangilliit",
			territory.V_034: "Asia Kujalleq",
			territory.V_039: "Europa Kujalleq",
			territory.V_053: "Australia aamma Nutaaq Zeelandi",
			territory.V_054: "Melanesia",
			territory.V_061: "Polynesia",
			territory.V_142: "Asia",
			territory.V_143: "Asia Qitiusumik",
			territory.V_145: "Asia Killiit",
			territory.V_150: "Europa",
			territory.V_151: "Europa Kangilliit",
			territory.V_154: "Europa Avannarleq",
			territory.V_155: "Europa Killiit",
			territory.V_419: "America Latin aamma Karibia",
			territory.AC:    "Ascension qeqertaq",
			territory.AD:    "Andorra",
			territory.AF:    "Afghanistani",
			territory.AG:    "Antigua aamma Barbuda",
			territory.AI:    "Anguilla",
			territory.AL:    "Albania",
			territory.AM:    "Armenia",
			territory.AO:    "Angola",
			territory.AQ:    "Qalasersuaq Kujalleq",
			territory.AR:    "Argentina",
			territory.AT:    "Østrigi",
			territory.AU:    "Australia",
			territory.AW:    "Aruba",
			territory.AX:    "Ålandi",
			territory.BA:    "Bosnia aamma Herzegovina",
			territory.BB:    "Barbados",
			territory.BD:    "Bangladesh",
			territory.BE:    "Belgia",
			territory.BF:    "Burkina Faso",
			territory.BG:    "Bulgaria",
			territory.BH:    "Bahrain",
			territory.BI:    "Burundi",
			territory.BJ:    "Benin",
			territory.BL:    "Saint Barthélemy",
			territory.BM:    "Bermuda",
			territory.BN:    "Brunei",
			territory.BO:    "Bolivia",
			territory.BR:    "Brazil",
			territory.BS:    "Bahamas",
			territory.BT:    "Bhutan",
			territory.BV:    "Bouvet qeqertaq",
			territory.BW:    "Botswana",
			territory.BY:    "Hvideruslandi",
			territory.BZ:    "Belize",
			territory.CA:    "Canada",
			territory.CC:    "Cocos qeqertaq",
			territory.CD:    "Kongo-Kinshasa",
			territory.CG:    "Kongo-Brazzaville",
			territory.CH:    "Schweizi",
			territory.CK:    "Cook qeqertaq",
			territory.CL:    "Chile",
			territory.CM:    "Kamerun",
			territory.CN:    "Kina",
			territory.CO:    "Colombia",
			territory.CP:    "Clipperton qeqertaq",
			territory.CR:    "Costa Rica",
			territory.CU:    "Kuba",
			territory.CV:    "Cap Verde",
			territory.CW:    "Curaçao",
			territory.CX:    "Jul-qeqertaq",
			territory.CY:    "Cypern",
			territory.CZ:    "Tjekkia",
			territory.DE:    "Tysklandi",
			territory.DG:    "Diego Garcia",
			territory.DJ:    "Djibouti",
			territory.DK:    "Danmarki",
			territory.DM:    "Dominica",
			territory.DZ:    "Algeriet",
			territory.EA:    "Ceuta aamma Melilla",
			territory.EC:    "Ecuador",
			territory.EE:    "Estlandi",
			territory.EG:    "Egypten",
			territory.EH:    "Sahara Killiit",
			territory.ER:    "Eritrea",
			territory.ES:    "Spania",
			territory.ET:    "Ethiopia",
			territory.EU:    "Europami nunat kattusimaffiat",
			territory.FI:    "Finlandi",
			territory.FJ:    "Fiji",
			territory.FK:    "Falklandi qeqertaq",
			territory.FM:    "Micronesia",
			territory.FO:    "Savalimmiut",
			territory.FR:    "Frankrigi",
			territory.GA:    "Gabon",
			territory.GB:    "Tuluit Nunaat",
			territory.GD:    "Grenada",
			territory.GE:    "Georgia",
			territory.GG:    "Guernsey",
			territory.GH:    "Ghana",
			territory.GI:    "Gibraltar",
			territory.GL:    "Kalaallit Nunaat",
			territory.GM:    "Gambia",
			territory.GN:    "Guinea",
			territory.GP:    "Guadeloupe",
			territory.GR:    "Grækenlandi",
			territory.GT:    "Guatemala",
			territory.GU:    "Guam",
			territory.GW:    "Guinea-Bissau",
			territory.GY:    "Guyana",
			territory.HK:    "Hongkong",
			territory.HN:    "Honduras",
			territory.HR:    "Kroatia",
			territory.HT:    "Haiti",
			territory.HU:    "Ungarni",
			territory.IC:    "Kanaria qeqertaq",
			territory.ID:    "Indonesia",
			territory.IE:    "Irlandi",
			territory.IL:    "Israel",
			territory.IM:    "Isle of Man",
			territory.IN:    "India",
			territory.IQ:    "Irak",
			territory.IR:    "Iran",
			territory.IS:    "Islandi",
			territory.IT:    "Italia",
			territory.JE:    "Jersey",
			territory.JM:    "Jamaica",
			territory.JO:    "Jordani",
			territory.JP:    "Japani",
			territory.KE:    "Kenya",
			territory.KH:    "Kambodia",
			territory.KI:    "Kiribati",
			territory.KM:    "Comoros",
			territory.KN:    "Saint Kitts aamma Nevis",
			territory.KP:    "Korea Avannarleq",
			territory.KR:    "Korea Kujalleq",
			territory.KW:    "Kuwait",
			territory.KY:    "Cayman qeqertaq",
			territory.KZ:    "Kasakhstani",
			territory.LA:    "Laos",
			territory.LB:    "Libanon",
			territory.LC:    "Saint Lucia",
			territory.LI:    "Liechtensteini",
			territory.LK:    "Sri Lanka",
			territory.LR:    "Liberia",
			territory.LS:    "Lesotho",
			territory.LT:    "Litaueni",
			territory.LU:    "Luxembourg",
			territory.LV:    "Letlandi",
			territory.LY:    "Libya",
			territory.MA:    "Marocko",
			territory.MC:    "Monaco",
			territory.MD:    "Moldova",
			territory.ME:    "Montenegro",
			territory.MF:    "Frankrigi Saint Martin",
			territory.MG:    "Madagaskar",
			territory.ML:    "Mali",
			territory.MM:    "Burma",
			territory.MO:    "Macao",
			territory.MQ:    "Martinique",
			territory.MR:    "Mauritania",
			territory.MS:    "Montserrat",
			territory.MT:    "Malta",
			territory.MU:    "Mauritius",
			territory.MW:    "Malawi",
			territory.MX:    "Mexiko",
			territory.MY:    "Malaysia",
			territory.MZ:    "Moçambique",
			territory.NA:    "Namibia",
			territory.NC:    "Nutaaq Caledonia",
			territory.NE:    "Niger",
			territory.NG:    "Nigeria",
			territory.NI:    "Nicaragua",
			territory.NL:    "Hollandi",
			territory.NO:    "Norge",
			territory.NP:    "Nepal",
			territory.NR:    "Nauru",
			territory.NU:    "Niue",
			territory.NZ:    "Nutaaq Zeelandi",
			territory.OM:    "Oman",
			territory.PA:    "Panama",
			territory.PE:    "Peru",
			territory.PG:    "Papua Nutaaq Guinea",
			territory.PK:    "Pakistani",
			territory.PL:    "Poleni",
			territory.PM:    "Saint Pierre aamma Miquelon",
			territory.PR:    "Puerto Rico",
			territory.PT:    "Portugali",
			territory.PW:    "Palau",
			territory.PY:    "Paraguay",
			territory.QA:    "Quatar",
			territory.RE:    "Réunion",
			territory.RO:    "Rumænia",
			territory.RS:    "Serbia",
			territory.RU:    "Ruslandi",
			territory.RW:    "Rwanda",
			territory.SA:    "Saudi Arabia",
			territory.SD:    "Avannarleqsudan",
			territory.SE:    "Sverige",
			territory.SG:    "Singapore",
			territory.SH:    "Saint Helena",
			territory.SI:    "Slovenia",
			territory.SJ:    "Svalbard aamma Jan Mayen",
			territory.SK:    "Slovakia",
			territory.SL:    "Sierra Leone",
			territory.SM:    "San Marino",
			territory.SN:    "Senegal",
			territory.SO:    "Somalia",
			territory.SR:    "Suriname",
			territory.SS:    "Kujalleqsudan",
			territory.ST:    "São Tomé aamma Príncipe",
			territory.SV:    "El Salvador",
			territory.SY:    "Syria",
			territory.SZ:    "Swazilandi",
			territory.TA:    "Tristan da Cunha",
			territory.TD:    "Chad",
			territory.TG:    "Togo",
			territory.TH:    "Thailandi",
			territory.TJ:    "Tajikistani",
			territory.TK:    "Tokelau",
			territory.TL:    "Timor Kangilliit",
			territory.TM:    "Turkmenistani",
			territory.TN:    "Tunisia",
			territory.TO:    "Tonga",
			territory.TR:    "Tyrkia",
			territory.TT:    "Trinidad aamma Tobago",
			territory.TV:    "Tuvalu",
			territory.TW:    "Taiwan",
			territory.TZ:    "Tanzania",
			territory.UA:    "Ukraina",
			territory.UG:    "Uganda",
			territory.US:    "USA",
			territory.UY:    "Uruguay",
			territory.UZ:    "Uzbekistani",
			territory.VA:    "Vatikani",
			territory.VE:    "Venezuela",
			territory.VN:    "Vietnam",
			territory.VU:    "Vanuatu",
			territory.WF:    "Wallis aamma Futuna",
			territory.WS:    "Samoa",
			territory.XK:    "Kosovo",
			territory.YE:    "Jemen",
			territory.YT:    "Mayotte",
			territory.ZA:    "Kujalleqafrika",
			territory.ZM:    "Zambia",
			territory.ZW:    "Zimbabwe",
			territory.ZZ:    "(atorsinnaanngitsoq nunap imartaa nunataalu)",
		},
	}
}
