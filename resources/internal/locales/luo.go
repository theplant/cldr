// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_luo() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "luo",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "DAC", Feb: "DAR", Mar: "DAD", Apr: "DAN", May: "DAH", Jun: "DAU", Jul: "DAO", Aug: "DAB", Sep: "DOC", Oct: "DAP", Nov: "DGI", Dec: "DAG"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "C", Feb: "R", Mar: "D", Apr: "N", May: "B", Jun: "U", Jul: "B", Aug: "B", Sep: "C", Oct: "P", Nov: "C", Dec: "P"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Dwe mar Achiel", Feb: "Dwe mar Ariyo", Mar: "Dwe mar Adek", Apr: "Dwe mar Ang’wen", May: "Dwe mar Abich", Jun: "Dwe mar Auchiel", Jul: "Dwe mar Abiriyo", Aug: "Dwe mar Aboro", Sep: "Dwe mar Ochiko", Oct: "Dwe mar Apar", Nov: "Dwe mar gi achiel", Dec: "Dwe mar Apar gi ariyo"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "JMP", Mon: "WUT", Tue: "TAR", Wed: "TAD", Thu: "TAN", Fri: "TAB", Sat: "NGS"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "J", Mon: "W", Tue: "T", Wed: "T", Thu: "T", Fri: "T", Sat: "N"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Jumapil", Mon: "Wuok Tich", Tue: "Tich Ariyo", Wed: "Tich Adek", Thu: "Tich Ang’wen", Fri: "Tich Abich", Sat: "Ngeso"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "OD", PM: "OT"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "OD", PM: "OT"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "", Group: "", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "", Currency: "#,##0.00¤", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Dirham ya Falme za Kiarabu", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Kwanza ya Angola", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Dola ya Australia", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "Dinari ya Bahareni", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Faranga ya Burundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Pula mar Botswana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Dola mar Kanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Faranga ya Kongo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Faranga ya Uswisi", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Yuan Renminbi ya China", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Eskudo ya Kepuvede", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Faranga ya Jibuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dinari ya Aljeria", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Paund mar Misri", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfa ya Eritrea", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Birr mar Ethiopia", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Yuro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Pauni mar Uingereza", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Sedi mar Ghana", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi ya Gambia", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Faranga ya Gine", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupia ya India", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Yen mar Japan", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Siling mar Kenya", Symbol: "Ksh"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Faranga ya Komoro", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Dola mar Liberia", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Loti ya Lesoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Dinari ya Libya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Dirham ya Moroko", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Ariary ya Bukini", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Ugwiya ya Moritania (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Ugwiya ya Moritania", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupia ya Morisi", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Kwacha ya Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Metikali ya Msumbiji", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Dola ya Namibia", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Naira ya Nijeria", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "", Symbol: "NZ$"},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "", Symbol: "₲"},
				currency.RON: cldr.Currency{DisplayName: "", Symbol: "lei"},
				currency.RUB: cldr.Currency{DisplayName: "", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Faranga ya Rwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Riyal ya Saudia", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Rupia ya Shelisheli", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Pauni ya Sudani", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Pauni ya Santahelena", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Leoni", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Shilingi ya Somalia", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Dobra ya Sao Tome na Principe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Dobra ya Sao Tome na Principe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Dinari ya Tunisia", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Shilingi ya Tanzania", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Shilingi ya Uganda", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Dola", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "Faranga CFA BEAC", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Faranga CFA BCEAO", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Randi ya Afrika Kusini", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwacha ya Zambia (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kwacha ya Zambia", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Dola ya Zimbabwe", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK:  "Kiakan",
			language.AM:  "Kiamhari",
			language.AR:  "Kiarabu",
			language.BE:  "Kibelarusi",
			language.BG:  "Kibulgaria",
			language.BN:  "Kibangla",
			language.CS:  "Kichecki",
			language.DE:  "Kijerumani",
			language.EL:  "Kigiriki",
			language.EN:  "Kingereza",
			language.ES:  "Kihispania",
			language.FA:  "Kiajemi",
			language.FR:  "Kifaransa",
			language.HA:  "Kihausa",
			language.HI:  "Kihindi",
			language.HU:  "Kihungari",
			language.ID:  "Kiindonesia",
			language.IG:  "Kiigbo",
			language.IT:  "Kiitaliano",
			language.JA:  "Kijapani",
			language.JV:  "Kijava",
			language.KM:  "Kikambodia",
			language.KO:  "Kikorea",
			language.LUO: "Dholuo",
			language.MS:  "Kimalesia",
			language.MY:  "Kiburma",
			language.NE:  "Kinepali",
			language.NL:  "Kiholanzi",
			language.PA:  "Kipunjabi",
			language.PL:  "Kipolandi",
			language.PT:  "Kireno",
			language.RO:  "Kiromania",
			language.RU:  "Kirusi",
			language.RW:  "Kinyarwanda",
			language.SO:  "Kisomali",
			language.SV:  "Kiswidi",
			language.TA:  "Kitamil",
			language.TH:  "Kitailandi",
			language.TR:  "Kituruki",
			language.UK:  "Kiukrania",
			language.UR:  "Kiurdu",
			language.VI:  "Kivietinamu",
			language.YO:  "Kiyoruba",
			language.ZH:  "Kichina",
			language.ZU:  "Kizulu",
		},
		Territories: cldr.Territories{
			territory.AD: "Andorra",
			territory.AE: "United Arab Emirates",
			territory.AF: "Afghanistan",
			territory.AG: "Antigua gi Barbuda",
			territory.AI: "Anguilla",
			territory.AL: "Albania",
			territory.AM: "Armenia",
			territory.AO: "Angola",
			territory.AR: "Argentina",
			territory.AS: "American Samoa",
			territory.AT: "Austria",
			territory.AU: "Australia",
			territory.AW: "Aruba",
			territory.AZ: "Azerbaijan",
			territory.BA: "Bosnia gi Herzegovina",
			territory.BB: "Barbados",
			territory.BD: "Bangladesh",
			territory.BE: "Belgium",
			territory.BF: "Burkina Faso",
			territory.BG: "Bulgaria",
			territory.BH: "Bahrain",
			territory.BI: "Burundi",
			territory.BJ: "Benin",
			territory.BM: "Bermuda",
			territory.BN: "Brunei",
			territory.BO: "Bolivia",
			territory.BR: "Brazil",
			territory.BS: "Bahamas",
			territory.BT: "Bhutan",
			territory.BW: "Botswana",
			territory.BY: "Belarus",
			territory.BZ: "Belize",
			territory.CA: "Canada",
			territory.CD: "Democratic Republic of the Congo",
			territory.CF: "Central African Republic",
			territory.CG: "Congo",
			territory.CH: "Switzerland",
			territory.CI: "Côte d",
			territory.CK: "Cook Islands",
			territory.CL: "Chile",
			territory.CM: "Cameroon",
			territory.CN: "China",
			territory.CO: "Colombia",
			territory.CR: "Costa Rica",
			territory.CU: "Cuba",
			territory.CV: "Cape Verde Islands",
			territory.CY: "Cyprus",
			territory.CZ: "Czech Republic",
			territory.DE: "Germany",
			territory.DJ: "Djibouti",
			territory.DK: "Denmark",
			territory.DM: "Dominica",
			territory.DO: "Dominican Republic",
			territory.DZ: "Algeria",
			territory.EC: "Ecuador",
			territory.EE: "Estonia",
			territory.EG: "Egypt",
			territory.ER: "Eritrea",
			territory.ES: "Spain",
			territory.ET: "Ethiopia",
			territory.FI: "Finland",
			territory.FJ: "Fiji",
			territory.FK: "Chuia mar Falkland",
			territory.FM: "Micronesia",
			territory.FR: "France",
			territory.GA: "Gabon",
			territory.GB: "United Kingdom",
			territory.GD: "Grenada",
			territory.GE: "Georgia",
			territory.GF: "French Guiana",
			territory.GH: "Ghana",
			territory.GI: "Gibraltar",
			territory.GL: "Greenland",
			territory.GM: "Gambia",
			territory.GN: "Guinea",
			territory.GP: "Guadeloupe",
			territory.GQ: "Equatorial Guinea",
			territory.GR: "Greece",
			territory.GT: "Guatemala",
			territory.GU: "Guam",
			territory.GW: "Guinea-Bissau",
			territory.GY: "Guyana",
			territory.HN: "Honduras",
			territory.HR: "Croatia",
			territory.HT: "Haiti",
			territory.HU: "Hungary",
			territory.ID: "Indonesia",
			territory.IE: "Ireland",
			territory.IL: "Israel",
			territory.IN: "India",
			territory.IO: "British Indian Ocean Territory",
			territory.IQ: "Iraq",
			territory.IR: "Iran",
			territory.IS: "Iceland",
			territory.IT: "Italy",
			territory.JM: "Jamaica",
			territory.JO: "Jordan",
			territory.JP: "Japan",
			territory.KE: "Kenya",
			territory.KG: "Kyrgyzstan",
			territory.KH: "Cambodia",
			territory.KI: "Kiribati",
			territory.KM: "Comoros",
			territory.KN: "Saint Kitts gi Nevis",
			territory.KP: "Korea Masawa",
			territory.KR: "Korea Milambo",
			territory.KW: "Kuwait",
			territory.KY: "Cayman Islands",
			territory.KZ: "Kazakhstan",
			territory.LA: "Laos",
			territory.LB: "Lebanon",
			territory.LC: "Saint Lucia",
			territory.LI: "Liechtenstein",
			territory.LK: "Sri Lanka",
			territory.LR: "Liberia",
			territory.LS: "Lesotho",
			territory.LT: "Lithuania",
			territory.LU: "Luxembourg",
			territory.LV: "Latvia",
			territory.LY: "Libya",
			territory.MA: "Morocco",
			territory.MC: "Monaco",
			territory.MD: "Moldova",
			territory.MG: "Madagascar",
			territory.MH: "Chuia mar Marshall",
			territory.ML: "Mali",
			territory.MM: "Myanmar",
			territory.MN: "Mongolia",
			territory.MP: "Northern Mariana Islands",
			territory.MQ: "Martinique",
			territory.MR: "Mauritania",
			territory.MS: "Montserrat",
			territory.MT: "Malta",
			territory.MU: "Mauritius",
			territory.MV: "Maldives",
			territory.MW: "Malawi",
			territory.MX: "Mexico",
			territory.MY: "Malaysia",
			territory.MZ: "Mozambique",
			territory.NA: "Namibia",
			territory.NC: "New Caledonia",
			territory.NE: "Niger",
			territory.NF: "Chuia mar Norfolk",
			territory.NG: "Nigeria",
			territory.NI: "Nicaragua",
			territory.NL: "Netherlands",
			territory.NO: "Norway",
			territory.NP: "Nepal",
			territory.NR: "Nauru",
			territory.NU: "Niue",
			territory.NZ: "New Zealand",
			territory.OM: "Oman",
			territory.PA: "Panama",
			territory.PE: "Peru",
			territory.PF: "French Polynesia",
			territory.PG: "Papua New Guinea",
			territory.PH: "Philippines",
			territory.PK: "Pakistan",
			territory.PL: "Poland",
			territory.PM: "Saint Pierre gi Miquelon",
			territory.PN: "Pitcairn",
			territory.PR: "Puerto Rico",
			territory.PS: "Palestinian West Bank gi Gaza",
			territory.PT: "Portugal",
			territory.PW: "Palau",
			territory.PY: "Paraguay",
			territory.QA: "Qatar",
			territory.RE: "Réunion",
			territory.RO: "Romania",
			territory.RU: "Russia",
			territory.RW: "Rwanda",
			territory.SA: "Saudi Arabia",
			territory.SB: "Solomon Islands",
			territory.SC: "Seychelles",
			territory.SD: "Sudan",
			territory.SE: "Sweden",
			territory.SG: "Singapore",
			territory.SH: "Saint Helena",
			territory.SI: "Slovenia",
			territory.SK: "Slovakia",
			territory.SL: "Sierra Leone",
			territory.SM: "San Marino",
			territory.SN: "Senegal",
			territory.SO: "Somalia",
			territory.SR: "Suriname",
			territory.ST: "São Tomé gi Príncipe",
			territory.SV: "El Salvador",
			territory.SY: "Syria",
			territory.SZ: "Swaziland",
			territory.TC: "Turks gi Caicos Islands",
			territory.TD: "Chad",
			territory.TG: "Togo",
			territory.TH: "Thailand",
			territory.TJ: "Tajikistan",
			territory.TK: "Tokelau",
			territory.TL: "East Timor",
			territory.TM: "Turkmenistan",
			territory.TN: "Tunisia",
			territory.TO: "Tonga",
			territory.TR: "Turkey",
			territory.TT: "Trinidad gi Tobago",
			territory.TV: "Tuvalu",
			territory.TW: "Taiwan",
			territory.TZ: "Tanzania",
			territory.UA: "Ukraine",
			territory.UG: "Uganda",
			territory.US: "USA",
			territory.UY: "Uruguay",
			territory.UZ: "Uzbekistan",
			territory.VA: "Vatican State",
			territory.VC: "Saint Vincent gi Grenadines",
			territory.VE: "Venezuela",
			territory.VG: "British Virgin Islands",
			territory.VI: "U.S. Virgin Islands",
			territory.VN: "Vietnam",
			territory.VU: "Vanuatu",
			territory.WF: "Wallis gi Futuna",
			territory.WS: "Samoa",
			territory.YE: "Yemen",
			territory.YT: "Mayotte",
			territory.ZA: "South Africa",
			territory.ZM: "Zambia",
			territory.ZW: "Zimbabwe",
		},
	}
}