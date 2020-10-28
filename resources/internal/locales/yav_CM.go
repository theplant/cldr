// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_yav_CM() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "yav_CM",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "o.1", Feb: "o.2", Mar: "o.3", Apr: "o.4", May: "o.5", Jun: "o.6", Jul: "o.7", Aug: "o.8", Sep: "o.9", Oct: "o.10", Nov: "o.11", Dec: "o.12"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "pikítíkítie, oólí ú kutúan", Feb: "siɛyɛ́, oóli ú kándíɛ", Mar: "ɔnsúmbɔl, oóli ú kátátúɛ", Apr: "mesiŋ, oóli ú kénie", May: "ensil, oóli ú kátánuɛ", Jun: "ɔsɔn", Jul: "efute", Aug: "pisuyú", Sep: "imɛŋ i puɔs", Oct: "imɛŋ i putúk,oóli ú kátíɛ", Nov: "makandikɛ", Dec: "pilɔndɔ́"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "sd", Mon: "md", Tue: "mw", Wed: "et", Thu: "kl", Fri: "fl", Sat: "ss"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "s", Mon: "m", Tue: "m", Wed: "e", Thu: "k", Fri: "f", Sat: "s"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "sɔ́ndiɛ", Mon: "móndie", Tue: "muányáŋmóndie", Wed: "metúkpíápɛ", Thu: "kúpélimetúkpiapɛ", Fri: "feléte", Sat: "séselé"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "kiɛmɛ́ɛm", PM: "kisɛ́ndɛ"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "kiɛmɛ́ɛm", PM: "kisɛ́ndɛ"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤;(#,##0.00\u00a0¤)", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AOA: cldr.Currency{DisplayName: "kuansa wu angolá", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "toláal wu ostalalí", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "tináal wu paaléen", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "faláŋɛ u pulundí", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "pula pu posuána", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "toláal u kanáta", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "́faláŋɛ u kongó", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "yuan ɛlɛnmimbí", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "ɛskúdo u kápfɛ́ɛl", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "́faláŋɛ u síputí", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "tináal wu alselí", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "lífilɛ wu isípit", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "náfka wu elitilée", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "píil wu etiopí", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "olóo", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "lífilɛ sitelelíiŋ", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "setí", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "talasí u kaambí", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "silí u kiiné", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "ulupí", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "yɛ́ɛn u sapɔ́ɔŋ", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "síliŋ u kénia", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "́faláŋɛ u kɔmɔ́ɔl", Symbol: "CF"},
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
				currency.RUB: cldr.Currency{DisplayName: "", Symbol: "₽"},
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
			language.AK:  "akánɛ",
			language.AM:  "amalíke",
			language.AR:  "́pakas",
			language.BE:  "pielúse",
			language.BG:  "bulgálɛ",
			language.BN:  "pengálɛ́ɛ",
			language.CS:  "cɛ́kɛ́ɛ",
			language.DE:  "ŋndiáman",
			language.EL:  "yavánɛ",
			language.EN:  "íŋgilísé",
			language.ES:  "nuɛspanyɔ́lɛ",
			language.FA:  "nupɛ́lisɛ",
			language.FR:  "feleŋsí",
			language.HA:  "pakas",
			language.HI:  "índí",
			language.HU:  "ɔ́ŋgɛ",
			language.ID:  "índonísiɛ",
			language.IG:  "íbo",
			language.IT:  "itáliɛ",
			language.JA:  "ndiáman",
			language.JV:  "yávanɛ",
			language.KM:  "kímɛɛ",
			language.KO:  "kolíe",
			language.MS:  "máliɛ",
			language.MY:  "bímanɛ",
			language.NE:  "nunipálɛ",
			language.NL:  "nilándɛ",
			language.PA:  "nupunsapíɛ́",
			language.PL:  "nupolonɛ́ɛ",
			language.PT:  "nupɔlitukɛ́ɛ",
			language.RO:  "nulumɛ́ŋɛ",
			language.RU:  "nulúse",
			language.RW:  "nuluándɛ́ɛ",
			language.SO:  "nusomalíɛ",
			language.SV:  "nusuetua",
			language.TA:  "nutámule",
			language.TH:  "nutáyɛ",
			language.TR:  "nutúluke",
			language.UK:  "nukeleniɛ́ŋɛ",
			language.UR:  "nulutú",
			language.VI:  "nufiɛtnamíɛŋ",
			language.YAV: "nuasue",
			language.YO:  "nuyolúpa",
			language.ZH:  "sinúɛ",
			language.ZU:  "nusulú",
		},
		Territories: cldr.Territories{
			territory.AD: "Aŋtúla",
			territory.AE: "imiláat i paaláap",
			territory.AF: "Afkanistáŋ",
			territory.AG: "Aŋtíka na Palpúta",
			territory.AI: "Aŋkíla",
			territory.AL: "Alpaní",
			territory.AM: "Almanía",
			territory.AO: "Aŋkúla",
			territory.AR: "Alsaŋtín",
			territory.AS: "Sámua u Amelíka",
			territory.AT: "Otilís",
			territory.AU: "Otalalí",
			territory.AW: "Alúpa",
			territory.AZ: "Asɛlpaisáŋ",
			territory.BA: "Pusiní-ɛlkofína",
			territory.BB: "Palpatós",
			territory.BD: "Paŋkalatɛs",
			territory.BE: "Pɛlsíik",
			territory.BF: "Pulikínafásó",
			territory.BG: "Pulukalíi",
			territory.BH: "Palɛŋ",
			territory.BI: "Púlúndí",
			territory.BJ: "Penɛŋ",
			territory.BM: "Pɛlmúta",
			territory.BN: "Pulunéy",
			territory.BO: "Polífia",
			territory.BR: "Pilesíl",
			territory.BS: "Pahámas",
			territory.BT: "Putaŋ",
			territory.BW: "Posuána",
			territory.BY: "Pelalús",
			territory.BZ: "Pelíse",
			territory.CA: "Kánáta",
			territory.CD: "kitɔŋ kí kongó",
			territory.CF: "Santalafilíik",
			territory.CG: "Kongó",
			territory.CH: "suwíis",
			territory.CI: "Kótifualɛ",
			territory.CK: "Kúuke",
			territory.CL: "Silí",
			territory.CM: "Kemelún",
			territory.CN: "Síine",
			territory.CO: "Kɔlɔ́mbía",
			territory.CR: "Kóstálíka",
			territory.CU: "kúpa",
			territory.CV: "Kápfɛl",
			territory.CY: "síplɛ",
			territory.CZ: "kitɔŋ kí cɛ́k",
			territory.DE: "nsáman",
			territory.DJ: "síputí",
			territory.DK: "tanemálk",
			territory.DM: "túmúnéke",
			territory.DO: "kitɔŋ kí tumunikɛ́ŋ",
			territory.DZ: "Alselí",
			territory.EC: "ekuatɛ́l",
			territory.EE: "ɛstoni",
			territory.EG: "isípit",
			territory.ER: "elitée",
			territory.ES: "panyá",
			territory.ET: "etiopí",
			territory.FI: "fɛnlánd",
			territory.FJ: "físi",
			territory.FK: "maluwín",
			territory.FM: "mikolonesí",
			territory.FR: "felensí",
			territory.GA: "kapɔ́ŋ",
			territory.GB: "ingilíís",
			territory.GD: "kelenáat",
			territory.GE: "sɔlsíi",
			territory.GF: "kuyáan u felensí",
			territory.GH: "kaná",
			territory.GI: "sílpalatáal",
			territory.GL: "kuluɛnlánd",
			territory.GM: "kambíi",
			territory.GN: "kiiné",
			territory.GP: "kuatelúup",
			territory.GQ: "kinéekuatolial",
			territory.GR: "kilɛ́ɛk",
			territory.GT: "kuatemalá",
			territory.GU: "kuamiɛ",
			territory.GW: "kiinépisaó",
			territory.GY: "kuyáan",
			territory.HN: "ɔndúlas",
			territory.HR: "Kolowasíi",
			territory.HT: "ayíti",
			territory.HU: "ɔngilí",
			territory.ID: "ɛndonesí",
			territory.IE: "ililánd",
			territory.IL: "ísilayɛ́l",
			territory.IN: "ɛ́ɛnd",
			territory.IO: "Kɔɔ́m kí ndián yi ngilís",
			territory.IQ: "ilák",
			territory.IR: "iláŋ",
			territory.IS: "isláand",
			territory.IT: "italí",
			territory.JM: "samayíik",
			territory.JO: "sɔltaní",
			territory.JP: "sapɔ́ɔŋ",
			territory.KE: "kénia",
			territory.KG: "kilikisistáŋ",
			territory.KH: "Kámbóse",
			territory.KI: "kilipatí",
			territory.KM: "Kɔmɔ́ɔl",
			territory.KN: "sɛ́ŋkilistɔ́f eniɛ́f",
			territory.KP: "kɔlé u muɛnɛ́",
			territory.KR: "kɔlé wu mbát",
			territory.KW: "kowéet",
			territory.KY: "Káyímanɛ",
			territory.KZ: "kasaksitáŋ",
			territory.LA: "lawós",
			territory.LB: "lipáŋ",
			territory.LC: "sɛ́ŋtɛ́lusí",
			territory.LI: "lístɛ́nsitáyin",
			territory.LK: "silíláŋka",
			territory.LR: "lipélia",
			territory.LS: "lesotó",
			territory.LT: "litiyaní",
			territory.LU: "liksambúul",
			territory.LV: "letoní",
			territory.LY: "lipíi",
			territory.MA: "malóok",
			territory.MC: "monakó",
			territory.MD: "moltafí",
			territory.MG: "matakaskáal",
			territory.MH: "ílmalasáal",
			territory.ML: "malí",
			territory.MM: "miaŋmáal",
			territory.MN: "mongolí",
			territory.MP: "il maliyanɛ u muɛnɛ́",
			territory.MQ: "maltiníik",
			territory.MR: "molitaní",
			territory.MS: "mɔŋsilá",
			territory.MT: "málɛ́t",
			territory.MU: "molís",
			territory.MV: "maletíif",
			territory.MW: "malawí",
			territory.MX: "mɛksíik",
			territory.MY: "malesí",
			territory.MZ: "mosambík",
			territory.NA: "namipí",
			territory.NC: "nufɛ́l kaletoní",
			territory.NE: "nisɛ́ɛl",
			territory.NF: "il nɔ́lfɔ́lɔk",
			territory.NG: "nisélia",
			territory.NI: "nikalaká",
			territory.NL: "nitililáand",
			territory.NO: "nɔlfɛ́ɛs",
			territory.NP: "nepáal",
			territory.NR: "nawulú",
			territory.NU: "niyuwé",
			territory.NZ: "nufɛ́l seláand",
			territory.OM: "omáŋ",
			territory.PA: "panamá",
			territory.PE: "pelú",
			territory.PF: "polinesí u felensí",
			territory.PG: "papuasí nufɛ́l kiiné",
			territory.PH: "filipíin",
			territory.PK: "pakistáŋ",
			territory.PL: "pɔlɔ́ɔny",
			territory.PM: "sɛ́ŋpiɛ́l e mikelɔ́ŋ",
			territory.PN: "pitikɛ́ɛlínɛ́",
			territory.PR: "pólótolíko",
			territory.PS: "kitɔŋ ki palɛstíin",
			territory.PT: "pɔltukáal",
			territory.PW: "palawú",
			territory.PY: "palakúé",
			territory.QA: "katáal",
			territory.RE: "elewuniɔ́ŋ",
			territory.RO: "ulumaní",
			territory.RU: "ulusí",
			territory.RW: "uluándá",
			territory.SA: "alapísawutíit",
			territory.SB: "il salomɔ́ŋ",
			territory.SC: "sesɛ́ɛl",
			territory.SD: "sutáaŋ",
			territory.SE: "suɛ́t",
			territory.SG: "singapúul",
			territory.SH: "sɛ́ŋtɛ́ elɛ́ɛnɛ",
			territory.SI: "silofení",
			territory.SK: "silofakí",
			territory.SL: "sieláleyɔ́ɔn",
			territory.SM: "san malíno",
			territory.SN: "senekáal",
			territory.SO: "somalí",
			territory.SR: "sulináam",
			territory.ST: "sáwó tomé e pelensípe",
			territory.SV: "salfatɔ́ɔl",
			territory.SZ: "suasiláand",
			territory.TC: "túluk na káyiik",
			territory.TD: "Sáat",
			territory.TG: "tokó",
			territory.TH: "tayiláand",
			territory.TJ: "tasikistáaŋ",
			territory.TK: "tokeló",
			territory.TL: "timɔ́ɔl u nipálɛ́n",
			territory.TM: "tulukmenisitáaŋ",
			territory.TN: "tunusí",
			territory.TO: "tɔ́ŋka",
			territory.TR: "tulukíi",
			territory.TT: "tilinitáat na tupákɔ",
			territory.TV: "tufalú",
			territory.TW: "tayiwáan",
			territory.TZ: "taŋsaní",
			territory.UA: "ukilɛ́ɛn",
			territory.UG: "ukánda",
			territory.US: "amálíka",
			territory.UY: "ulukuéy",
			territory.UZ: "usupekistáaŋ",
			territory.VA: "fatikáaŋ",
			territory.VC: "sɛ́ŋ fɛŋsáŋ elekelenatíin",
			territory.VE: "fenesuwelá",
			territory.VG: "Filisíin ungilís",
			territory.VI: "pindisúlɛ́ pi amálíka",
			territory.VN: "fiɛtnáam",
			territory.VU: "fanuatú",
			territory.WF: "walíis na futúna",
			territory.WS: "samowá",
			territory.YE: "yémɛn",
			territory.YT: "mayɔ́ɔt",
			territory.ZA: "afilí mbátɛ́",
			territory.ZM: "saambíi",
			territory.ZW: "simbapuwé",
		},
	}
}
