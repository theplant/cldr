// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_prg() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "prg",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, y 'mettas' d. MMMM", Long: "y 'mettas' d. MMMM", Medium: "dd.MM 'st'. y", Short: "dd.MM.yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "rag", Feb: "was", Mar: "pūl", Apr: "sak", May: "zal", Jun: "sīm", Jul: "līp", Aug: "dag", Sep: "sil", Oct: "spa", Nov: "lap", Dec: "sal"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "R", Feb: "W", Mar: "P", Apr: "S", May: "Z", Jun: "S", Jul: "L", Aug: "D", Sep: "S", Oct: "S", Nov: "L", Dec: "S"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "rags", Feb: "wassarins", Mar: "pūlis", Apr: "sakkis", May: "zallaws", Jun: "sīmenis", Jul: "līpa", Aug: "daggis", Sep: "sillins", Oct: "spallins", Nov: "lapkrūtis", Dec: "sallaws"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "nad", Mon: "pan", Tue: "wis", Wed: "pus", Thu: "ket", Fri: "pēn", Sat: "sab"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "N", Mon: "P", Tue: "W", Wed: "P", Thu: "K", Fri: "P", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "nadīli", Mon: "panadīli", Tue: "wisasīdis", Wed: "pussisawaiti", Thu: "ketwirtiks", Fri: "pēntniks", Sat: "sabattika"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "ankstāinan", PM: "pa pussideinan"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "", Percent: "#,##0%"},
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
				currency.BRL: cldr.Currency{DisplayName: "Brazīlijas reals", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "", Symbol: "CA$"},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Kīnas juāns", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.EGP: cldr.Currency{DisplayName: "", Symbol: "E£"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.EUR: cldr.Currency{DisplayName: "eurō", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "punds sterlings", Symbol: "£"},
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
				currency.INR: cldr.Currency{DisplayName: "Īndijas rūpija", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Japānijas jāns", Symbol: "JP¥"},
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
				currency.NOK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "", Symbol: "NZ$"},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "", Symbol: "₲"},
				currency.RON: cldr.Currency{DisplayName: "", Symbol: "lei"},
				currency.RUB: cldr.Currency{DisplayName: "Russis rūbels", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "", Symbol: "RF"},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
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
				currency.USD: cldr.Currency{DisplayName: "APW dālars", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "niwaistā walūta", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AR:      "arābiskan",
			language.DA:      "dāniskan",
			language.DE:      "miksiskan",
			language.DE_AT:   "Āustrarīkis miksiskan",
			language.DE_CH:   "Šwēicis aūktamiksiskan",
			language.EL:      "grēkiskan",
			language.EN:      "ēngliskan",
			language.EN_AU:   "Austrālijas ēngliskan",
			language.EN_CA:   "Kanādas ēngliskan",
			language.EN_GB:   "brītiskan ēngliskan",
			language.EN_US:   "APW ēngliskan",
			language.ES:      "špāniskan",
			language.ES_419:  "Lātiniskas Amērikas špāniskan",
			language.ES_ES:   "eurōpiskan špāniskan",
			language.ES_MX:   "Meksikus špāniskan",
			language.ET:      "èstiskan",
			language.FI:      "sōmiskan",
			language.FR:      "prancōziskan",
			language.FR_CA:   "Kanādas prancōziskan",
			language.FR_CH:   "Šwēicis prancōziskan",
			language.IT:      "wālkiskan",
			language.JA:      "japāniskan",
			language.LT:      "laītawiskan",
			language.LV:      "lattawiskan",
			language.NL:      "ullandiskan",
			language.PL:      "pōliskan",
			language.PRG:     "prūsiskan",
			language.PT:      "pōrtugaliskan",
			language.PT_BR:   "Brazīlijas pōrtugaliskan",
			language.PT_PT:   "eurōpiskan pōrtugaliskan",
			language.RU:      "maskōwitiskan",
			language.SV:      "šwēdiskan",
			language.TR:      "turkiskan",
			language.UND:     "niwaistā bilā",
			language.ZH:      "kīniskan",
			language.ZH_HANS: "prastintan kīniskan",
			language.ZH_HANT: "tradiciōnalin kīniskan",
		},
		Territories: cldr.Territories{
			territory.V_001: "swītai",
			territory.V_002: "Afrika",
			territory.V_003: "Zēimanamērika",
			territory.V_005: "Pussideinanamērika",
			territory.V_019: "Amērika",
			territory.V_142: "Āzija",
			territory.V_150: "Eurōpa",
			territory.AD:    "Andōra",
			territory.AG:    "Antīgwa be Barbūda",
			territory.AL:    "Albānija",
			territory.AR:    "Argentīnija",
			territory.AT:    "Āustrarīki",
			territory.AU:    "Austrālija",
			territory.BA:    "Bōsnija be Ercegōwina",
			territory.BB:    "Barbādas",
			territory.BE:    "Belgija",
			territory.BG:    "Bulgārija",
			territory.BO:    "Bōliwija",
			territory.BR:    "Brazīlija",
			territory.BS:    "Bahāmai",
			territory.BY:    "Krēiwa",
			territory.BZ:    "Belīzi",
			territory.CA:    "Kānada",
			territory.CH:    "Šwēici",
			territory.CL:    "Čīli",
			territory.CN:    "Kīna",
			territory.CO:    "Kōlumbija",
			territory.CR:    "Costa Rica",
			territory.CU:    "Kūba",
			territory.CZ:    "Čekkija",
			territory.DE:    "Mikskātauta",
			territory.DK:    "Dānanmarki",
			territory.DM:    "Dōminika",
			territory.DO:    "Dōminikas Republīki",
			territory.EC:    "Ekwadōrs",
			territory.EE:    "Estantauta",
			territory.ES:    "Špānija",
			territory.FI:    "Sōmija",
			territory.FO:    "Farēirai",
			territory.FR:    "Prankrīki",
			territory.GB:    "DB",
			territory.GD:    "Grenāda",
			territory.GF:    "Prancōziska Gujāna",
			territory.GI:    "Gibrāltars",
			territory.GL:    "Grēnlandan",
			territory.GR:    "Grēkantauta",
			territory.GT:    "Gwatemāla",
			territory.GY:    "Gujāna",
			territory.HN:    "Hōnduras",
			territory.HR:    "Kruātija",
			territory.HT:    "Haīti",
			territory.HU:    "Ungrai",
			territory.ID:    "Indōnezija",
			territory.IN:    "Īndija",
			territory.IS:    "Īslandan",
			territory.IT:    "Wālkija",
			territory.JM:    "Jamāika",
			territory.JP:    "Japānija",
			territory.KR:    "Pussideinankōreja",
			territory.LI:    "Līchtenšteinan",
			territory.LT:    "Laītawa",
			territory.LU:    "Luksemburgan",
			territory.LV:    "Lattawa",
			territory.MC:    "Mōnakō",
			territory.MD:    "Mōldawija",
			territory.ME:    "Mōntenegran",
			territory.MT:    "Mālta",
			territory.MX:    "Meksiku",
			territory.NI:    "Nikarāgwa",
			territory.NO:    "Nōrwigai",
			territory.NZ:    "Nawazēlandan",
			territory.PA:    "Panāma",
			territory.PE:    "Perū",
			territory.PL:    "Pōli",
			territory.PT:    "Pōrtugalin",
			territory.PW:    "Palau",
			territory.PY:    "Paragwājs",
			territory.RO:    "Rumānija",
			territory.RS:    "Serbija",
			territory.RU:    "Russi",
			territory.SA:    "Saūdi Arābija",
			territory.SE:    "Šwēdija",
			territory.SI:    "Slōwenija",
			territory.SK:    "Slōwakei",
			territory.SM:    "San Marinō",
			territory.SR:    "Surināms",
			territory.SV:    "El Salvadōrs",
			territory.TH:    "Tāilandan",
			territory.TR:    "Turkāja",
			territory.TT:    "Trinidāds be Tobagō",
			territory.TW:    "Taiwāns",
			territory.UA:    "Ukrāini",
			territory.US:    "PW",
			territory.UY:    "Urugwājs",
			territory.VE:    "Venezuēla",
			territory.XK:    "Kōsawa",
			territory.ZA:    "Pussideinanafrika",
			territory.ZZ:    "niwaistā regiōni",
		},
	}
}
