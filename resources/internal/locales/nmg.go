// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_nmg() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "nmg",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ng1", Feb: "ng2", Mar: "ng3", Apr: "ng4", May: "ng5", Jun: "ng6", Jul: "ng7", Aug: "ng8", Sep: "ng9", Oct: "ng10", Nov: "ng11", Dec: "kris"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "ngwɛn matáhra", Feb: "ngwɛn ńmba", Mar: "ngwɛn ńlal", Apr: "ngwɛn ńna", May: "ngwɛn ńtan", Jun: "ngwɛn ńtuó", Jul: "ngwɛn hɛmbuɛrí", Aug: "ngwɛn lɔmbi", Sep: "ngwɛn rɛbvuâ", Oct: "ngwɛn wum", Nov: "ngwɛn wum navǔr", Dec: "krísimin"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "sɔ́n", Mon: "mɔ́n", Tue: "smb", Wed: "sml", Thu: "smn", Fri: "mbs", Sat: "sas"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "s", Mon: "m", Tue: "s", Wed: "s", Thu: "s", Fri: "m", Sat: "s"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "sɔ́ndɔ", Mon: "mɔ́ndɔ", Tue: "sɔ́ndɔ mafú mába", Wed: "sɔ́ndɔ mafú málal", Thu: "sɔ́ndɔ mafú mána", Fri: "mabágá má sukul", Sat: "sásadi"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "maná", PM: "kugú"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "maná", PM: "kugú"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Mɔn B ´Arabe", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Mɔn Angola", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Dɔ́llɔ Ɔstralia", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "Mɔn Bahrein", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Fraŋ Burundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Mɔn Botswana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Dɔ́llɔ Kanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Fraŋ bó Kongolɛ̌", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Fraŋ Suisse", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Mɔn bó Chinois", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Mɔn Kapvɛrt", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Fraŋ Jibuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Mɔn Algeria", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Mɔn Ägyptɛn", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Mɔn Erytré", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Mɔn Ethiopia", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Mɔn Ngɛ̄lɛ̄n", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Mɔn Gana", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Mɔn Gambia", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Fraŋ Guiné", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Mɔn India", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Mɔn Japɔn", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Mɔn Kɛnya", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Fraŋ bó Kɔmɔr", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Dɔ́llɔ Liberia", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Mɔn Lesoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Mɔn Libya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Mɔn Marɔk", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Mɔn Madagaskar", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Mɔn Moritania (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Mɔn Moritania", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Mɔn Moriss", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Mɔn Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Mɔn Mozambik", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Dɔ́llɔ Namibia", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Naïra Nigeria", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "Fraŋ Rwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Mɔn Saudi Arabia", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Mɔn Seychɛlle", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Mɔn Sudan", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Mɔn Sudan (1957–1998)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Mɔn má Saint Lina", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Mɔn Leɔne", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Mɔn Somalía", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Mɔn Sao tomé na prinship (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Mɔn Sao tomé na prinship", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Mɔn Ligangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Mɔn Tunisia", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Mɔn Tanzania", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Mɔn Uganda", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Dɔ́llɔ Amɛŕka", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "Fraŋ CFA BEAC", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Fraŋ CFA BCEAO", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Mɔn Afrik yí sí", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Mɔn Zambia (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Mɔn Zambia", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Dɔ́llɔ Zimbabwǝ (1980–2008)", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK:  "Kiɛl akan",
			language.AM:  "Kiɛl amaria",
			language.AR:  "Kiɛl b’árabe",
			language.BE:  "Kiɛl belarussie",
			language.BG:  "Kiɛl bulgaria",
			language.BN:  "Kiɛl bengalia",
			language.CS:  "Kiɛl bó tchɛk",
			language.DE:  "Jáman",
			language.EL:  "Kiɛl bó grɛk",
			language.EN:  "Ngɛ̄lɛ̄n",
			language.ES:  "Paŋá",
			language.FA:  "Kiɛl pɛrsia",
			language.FR:  "Fala",
			language.HA:  "Kiɛl máwúsá",
			language.HI:  "Kiɛl b’indien",
			language.HU:  "Kiɛl b’ɔ́ngrois",
			language.ID:  "Kiɛl indonesie",
			language.IG:  "Kiɛl ikbo",
			language.IT:  "Kiɛl italia",
			language.JA:  "Kiɛl bó japonɛ̌",
			language.JV:  "Kiɛl bó javanɛ̌",
			language.KM:  "Kiɛl bó mɛr",
			language.KO:  "Kiɛl koré",
			language.MS:  "Kiɛl Malɛ̌siā",
			language.MY:  "Kiɛl birmania",
			language.NE:  "Kiɛl nepal",
			language.NL:  "Kiɛl bóllandais",
			language.NMG: "Kwasio",
			language.PA:  "Kiɛl pɛndjabi",
			language.PL:  "Kiɛl pɔlɔŋe",
			language.PT:  "Kiɛl bó pɔ̄rtugɛ̂",
			language.RO:  "Kiɛl bó rumɛ̂n",
			language.RU:  "Kiɛl russia",
			language.RW:  "Kiɛl rwandā",
			language.SO:  "Kiɛl somaliā",
			language.SV:  "Kiɛl bó suedois",
			language.TA:  "Kiɛl tamul",
			language.TH:  "Kiɛl thaï",
			language.TR:  "Kiɛl bó turk",
			language.UK:  "Kiɛl b’ukrɛ̄nien",
			language.UR:  "Kiɛl úrdu",
			language.VI:  "Kiɛl viɛtnam",
			language.YO:  "Yorúbâ",
			language.ZH:  "Kiɛl bó chinois",
			language.ZU:  "Zulu",
		},
		Territories: cldr.Territories{
			territory.AD: "Andɔ́ra",
			territory.AE: "Minlambɔ́ Nsaŋ́nsa mí Arabia",
			territory.AF: "Afganistaŋ",
			territory.AG: "Antíga bá Barbúda",
			territory.AI: "Anguílla",
			territory.AL: "Albania",
			territory.AM: "Arménia",
			territory.AO: "Angola",
			territory.AR: "Argentína",
			territory.AS: "Samoa m ́Amɛ́rka",
			territory.AT: "Ötrish",
			territory.AU: "Östraliá",
			territory.AW: "Árúba",
			territory.AZ: "Azerbaïjaŋ",
			territory.BA: "Bosnia na Ɛrzegovina",
			territory.BB: "Barbado",
			territory.BD: "Bɛŋgladɛsh",
			territory.BE: "Bɛlgik",
			territory.BF: "Burkina Faso",
			territory.BG: "Bulgaria",
			territory.BH: "Bahrain",
			territory.BI: "Burundi",
			territory.BJ: "Benin",
			territory.BM: "Bɛrmuda",
			territory.BN: "Brunɛi",
			territory.BO: "Bolivia",
			territory.BR: "Brésil",
			territory.BS: "Bahamas",
			territory.BT: "Butaŋ",
			territory.BW: "Botswana",
			territory.BY: "Belarus",
			territory.BZ: "Bɛliz",
			territory.CA: "Kanada",
			territory.CD: "Kongó Zaïre",
			territory.CF: "Sentrafríka",
			territory.CG: "Kongo",
			territory.CH: "Switzɛrland",
			territory.CI: "Kote d´Ivoire",
			territory.CK: "Maŋ́ má Kook",
			territory.CL: "Tshili",
			territory.CM: "Kamerun",
			territory.CN: "Shine",
			territory.CO: "Kɔlɔ́mbia",
			territory.CR: "Kosta Ríka",
			territory.CU: "Kuba",
			territory.CV: "Maŋ́ má Kapvɛr",
			territory.CY: "Sipria",
			territory.CZ: "Nlambɔ́ bó tschɛk",
			territory.DE: "Jaman",
			territory.DJ: "Jibúti",
			territory.DK: "Danemark",
			territory.DM: "Dominíka",
			territory.DO: "Nlambɔ́ Dominíka",
			territory.DZ: "Algeria",
			territory.EC: "Ekuateur",
			territory.EE: "Ɛstonia",
			territory.EG: "Ägyptɛn",
			territory.ER: "Erytrea",
			territory.ES: "Paŋá",
			territory.ET: "Ethiopiá",
			territory.FI: "Finlande",
			territory.FJ: "Fijiá",
			territory.FK: "Maŋ má Falkland",
			territory.FM: "Mikronesia",
			territory.FR: "Fala",
			territory.GA: "Gabɔŋ",
			territory.GB: "Nlambɔ́ Ngɛlɛn",
			territory.GD: "Grenada",
			territory.GE: "Jɔrgia",
			territory.GF: "Guyane Fala",
			territory.GH: "Gána",
			territory.GI: "Gilbratar",
			territory.GL: "Greenland",
			territory.GM: "Gambia",
			territory.GN: "Guine",
			territory.GP: "Guadeloup",
			territory.GQ: "Guine Ekuatorial",
			territory.GR: "Grɛce",
			territory.GT: "Guatemala",
			territory.GU: "Guam",
			territory.GW: "Guine Bisso",
			territory.GY: "Guyana",
			territory.HN: "Ɔndúras",
			territory.HR: "Kroasia",
			territory.HT: "Haïti",
			territory.HU: "Ɔngría",
			territory.ID: "Indonesia",
			territory.IE: "Irland",
			territory.IL: "Äsrɛl",
			territory.IN: "India",
			territory.IO: "Nlambɔ́ ngɛlɛn ma yí maŋ ntsiɛh",
			territory.IQ: "Irak",
			territory.IR: "Iran",
			territory.IS: "Island",
			territory.IT: "Italia",
			territory.JM: "Jamaika",
			territory.JO: "Jɔrdania",
			territory.JP: "Japɔn",
			territory.KE: "Kɛnya",
			territory.KG: "Kyrgystaŋ",
			territory.KH: "Kambodia",
			territory.KI: "Kiribati",
			territory.KM: "Kɔmɔr",
			territory.KN: "Saint Kitts na Nevis",
			territory.KP: "Koré yí bvuɔ",
			territory.KR: "Koré yí sí",
			territory.KW: "Kowɛit",
			territory.KY: "Maŋ́ má kumbi",
			territory.KZ: "Kazakstaŋ",
			territory.LA: "Laos",
			territory.LB: "Libaŋ",
			territory.LC: "Saint Lucia",
			territory.LI: "Lishenstein",
			territory.LK: "Sri Lanka",
			territory.LR: "Liberia",
			territory.LS: "Lesoto",
			territory.LT: "Lituaniá",
			territory.LU: "Luxembourg",
			territory.LV: "Latvia",
			territory.LY: "Libya",
			territory.MA: "Marɔk",
			territory.MC: "Monako",
			territory.MD: "Mɔldavia",
			territory.MG: "Madagaskar",
			territory.MH: "Maŋ́ má Marshall",
			territory.ML: "Mali",
			territory.MM: "Myanmar",
			territory.MN: "Mɔngolia",
			territory.MP: "Maŋ́ Mariá",
			territory.MQ: "Martinika",
			territory.MR: "Moritania",
			territory.MS: "Mɔnserrat",
			territory.MT: "Malta",
			territory.MU: "Morisse",
			territory.MV: "Maldivia",
			territory.MW: "Malawi",
			territory.MX: "Mɛxik",
			territory.MY: "Malaysia",
			territory.MZ: "Mozambik",
			territory.NA: "Namibia",
			territory.NC: "Kaledoni nwanah",
			territory.NE: "Niger",
			territory.NF: "Maŋ́ má Nɔrfɔrk",
			territory.NG: "Nigeria",
			territory.NI: "Nikaragua",
			territory.NL: "Nedɛrland",
			territory.NO: "Nɔrvɛg",
			territory.NP: "Nepal",
			territory.NR: "Noru",
			territory.NU: "Niuɛ",
			territory.NZ: "Zeland nwanah",
			territory.OM: "Oman",
			territory.PA: "Panama",
			territory.PE: "Peru",
			territory.PF: "Polynesia Fala",
			territory.PG: "Guine Papuasi",
			territory.PH: "Filipin",
			territory.PK: "Pakistan",
			territory.PL: "Pɔlɔŋ",
			territory.PM: "Saint Peter ba Mikelɔn",
			territory.PN: "Pitkairn",
			territory.PR: "Puɛrto Riko",
			territory.PS: "Palɛstin",
			territory.PT: "Pɔrtugal",
			territory.PW: "Palo",
			territory.PY: "Paraguay",
			territory.QA: "Katar",
			territory.RE: "Réuniɔn",
			territory.RO: "Roumania",
			territory.RU: "Russi",
			territory.RW: "Rwanda",
			territory.SA: "Saudi Arabia",
			territory.SB: "Maŋ́ má Salomɔn",
			territory.SC: "Seychɛlle",
			territory.SD: "Sudaŋ",
			territory.SE: "Suɛd",
			territory.SG: "Singapur",
			territory.SH: "Saint Lina",
			territory.SI: "Slovenia",
			territory.SK: "Slovakia",
			territory.SL: "Sierra Leɔn",
			territory.SM: "San Marino",
			territory.SN: "Senegal",
			territory.SO: "Somália",
			territory.SR: "Surinam",
			territory.ST: "Sao Tomé ba Prinship",
			territory.SV: "Salvadɔr",
			territory.SY: "Syria",
			territory.SZ: "Swaziland",
			territory.TC: "Maŋ́ má Turk na Kaiko",
			territory.TD: "Tshad",
			territory.TG: "Togo",
			territory.TH: "Taïland",
			territory.TJ: "Tajikistaŋ",
			territory.TK: "Tokelo",
			territory.TL: "Timɔr tsindikēh",
			territory.TM: "Turkmɛnistaŋ",
			territory.TN: "Tunisiá",
			territory.TO: "Tɔnga",
			territory.TR: "Turki",
			territory.TT: "Trinidad ba Tobágó",
			territory.TV: "Tuvalú",
			territory.TW: "Taïwan",
			territory.TZ: "Tanzánía",
			territory.UA: "Ukrɛn",
			territory.UG: "Uganda",
			territory.US: "Amɛŕka",
			territory.UY: "Uruguay",
			territory.UZ: "Usbǝkistaŋ",
			territory.VA: "Vatikaŋ",
			territory.VC: "Saint Vincent ba Grenadines",
			territory.VE: "Vǝnǝzuela",
			territory.VG: "Minsilɛ́ mímaŋ mí ngɛ̄lɛ̄n",
			territory.VI: "Minsilɛ mí maŋ́ m´Amɛrka",
			territory.VN: "Viɛtnam",
			territory.VU: "Vanuatu",
			territory.WF: "Wallis ba Futuna",
			territory.WS: "Samoa",
			territory.YE: "Yǝmɛn",
			territory.YT: "Mayɔt",
			territory.ZA: "Afríka yí sí",
			territory.ZM: "Zambia",
			territory.ZW: "Zimbabwǝ",
		},
	}
}
