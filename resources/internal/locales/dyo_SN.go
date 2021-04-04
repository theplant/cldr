// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_dyo_SN() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "dyo_SN",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Sa", Feb: "Fe", Mar: "Ma", Apr: "Ab", May: "Me", Jun: "Su", Jul: "Sú", Aug: "Ut", Sep: "Se", Oct: "Ok", Nov: "No", Dec: "De"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "S", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "S", Jul: "S", Aug: "U", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Sanvie", Feb: "Fébirie", Mar: "Mars", Apr: "Aburil", May: "Mee", Jun: "Sueŋ", Jul: "Súuyee", Aug: "Ut", Sep: "Settembar", Oct: "Oktobar", Nov: "Novembar", Dec: "Disambar"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Dim", Mon: "Ten", Tue: "Tal", Wed: "Ala", Thu: "Ara", Fri: "Arj", Sat: "Sib"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "T", Tue: "T", Wed: "A", Thu: "A", Fri: "A", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Dimas", Mon: "Teneŋ", Tue: "Talata", Wed: "Alarbay", Thu: "Aramisay", Fri: "Arjuma", Sat: "Sibiti"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AOA: cldr.Currency{DisplayName: "kwanza yati Angola", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "dolaar yati Ostraalia", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "dinaar yati Bahrayn", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "fraaŋ yati Burundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "pula yati Boswana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "dolaar yati Kanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "fraaŋ yati Kongo", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "yuan yati Siin", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "eskuudo yati Kap Ver", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "fraaŋ yati Jibuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "dinaar yati Alseri", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "liiverey yati Esípt", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "nafka yati Eritree", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "birr yati Ecoopi", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "cedi yati Gaana", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "dalasi yati Gambi", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "sili yati Giné", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "rupii yati End", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "yen yati Sapoŋ", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "silliŋ yati Keniya", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "fraaŋ yati Komor", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "dolaar yati Liberia", Symbol: "$"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "dinaar yati Libia", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "ariari yati Madagaskaar", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "ugiiya yati Mooritanii (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "ugiiya yati Mooritanii", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "kwacha yati Malawi", Symbol: ""},
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
				currency.XAF: cldr.Currency{DisplayName: "seefa BEAC", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "seefa yati BCEAO", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK:  "akan",
			language.AM:  "amharik",
			language.AR:  "arab",
			language.BE:  "belarus",
			language.BG:  "bulgaari",
			language.BN:  "bengali",
			language.CS:  "sek",
			language.DE:  "alman",
			language.DYO: "joola",
			language.EL:  "greek",
			language.EN:  "angle",
			language.ES:  "español",
			language.FA:  "persan",
			language.FR:  "franse",
			language.HA:  "hausa",
			language.HI:  "endu",
			language.HU:  "ongrua",
			language.ID:  "indoneesi",
			language.IG:  "igbo",
			language.IT:  "italien",
			language.JA:  "saponee",
			language.JV:  "savanee",
			language.KM:  "kmeer",
			language.KO:  "koree",
			language.MS:  "maleesi",
			language.MY:  "birmani",
			language.NE:  "nepalees",
			language.NL:  "neerlande",
			language.PA:  "penjabi",
			language.PL:  "polonees",
			language.PT:  "portugees",
			language.RO:  "rumeen",
			language.RU:  "rus",
			language.RW:  "ruanda",
			language.SO:  "somali",
			language.SV:  "suedi",
			language.TA:  "tamil",
			language.TH:  "tay",
			language.TR:  "turki",
			language.UK:  "ukrain",
			language.UR:  "urdu",
			language.VI:  "vietnam",
			language.YO:  "yoruba",
			language.ZH:  "sinua",
			language.ZU:  "sulu",
		},
		Territories: cldr.Territories{
			territory.AD: "Andorra",
			territory.AF: "Afganistan",
			territory.AG: "Antigua di Barbuda",
			territory.AI: "Angiiya",
			territory.AL: "Albani",
			territory.AM: "Armeni",
			territory.AO: "Angola",
			territory.AR: "Arsantin",
			territory.AS: "Samoa yati Amerik",
			territory.AT: "Otris",
			territory.AU: "Ostraalia",
			territory.AW: "Aruba",
			territory.AZ: "Aserbaysan",
			territory.BA: "Bosni di Hersegovin",
			territory.BB: "Barbad",
			territory.BD: "Banglades",
			territory.BE: "Belsik",
			territory.BF: "Burukiina Faso",
			territory.BG: "Bulgari",
			territory.BH: "Bahrayn",
			territory.BI: "Burundi",
			territory.BJ: "Bene",
			territory.BM: "Bermud",
			territory.BN: "Buruney",
			territory.BO: "Boliivi",
			territory.BR: "Bresil",
			territory.BS: "Bahama",
			territory.BT: "Butan",
			territory.BW: "Boswana",
			territory.BY: "Belarus",
			territory.BZ: "Beliis",
			territory.CA: "Kanada",
			territory.CD: "Mofam demokratik mati Kongo",
			territory.CG: "Kongo",
			territory.CI: "Koddiwar",
			territory.CL: "Cili",
			territory.CM: "Kamerun",
			territory.CN: "Siin",
			territory.CO: "Kolombi",
			territory.CR: "Kosta Rika",
			territory.CU: "Kuba",
			territory.CV: "Kap Ver",
			territory.CY: "Siipr",
			territory.CZ: "Mofam mati Cek",
			territory.DE: "Almaañ",
			territory.DJ: "Jibuti",
			territory.DK: "Danmark",
			territory.DM: "Dominika",
			territory.DO: "Mofam mati Dominik",
			territory.DZ: "Alseri",
			territory.EC: "Ekuador",
			territory.EE: "Estoni",
			territory.EG: "Esípt",
			territory.ER: "Eritree",
			territory.ES: "Espaañ",
			territory.ET: "Ecoopi",
			territory.FI: "Finland",
			territory.FJ: "Fiji",
			territory.FR: "Frans",
			territory.GA: "Gabon",
			territory.GD: "Grenada",
			territory.GE: "Seorsi",
			territory.GH: "Gaana",
			territory.GI: "Sipraltaar",
			territory.GL: "Greenland",
			territory.GM: "Gambi",
			territory.GN: "Giné",
			territory.GP: "Guwadalup",
			territory.GR: "Gres",
			territory.GT: "Guatemala",
			territory.GU: "Guam",
			territory.GW: "Giné Bisaau",
			territory.GY: "Giyan",
			territory.HN: "Onduras",
			territory.HR: "Kroasi",
			territory.HT: "Ayti",
			territory.HU: "Oŋri",
			territory.ID: "Endonesi",
			territory.IE: "Irland",
			territory.IL: "Israel",
			territory.IN: "End",
			territory.IQ: "Irak",
			territory.IR: "Iran",
			territory.IS: "Iisland",
			territory.IT: "Itali",
			territory.JM: "Samaik",
			territory.JP: "Sapoŋ",
			territory.KE: "Keniya",
			territory.KH: "Kamboj",
			territory.KM: "Komor",
			territory.LC: "Saŋ Lusia",
			territory.LK: "Siri Lanka",
			territory.LR: "Liberia",
			territory.MG: "Madagaskaar",
			territory.ML: "Mali",
			territory.NF: "Ecinkey yati Noorfok",
			territory.SA: "Abari Saudi",
			territory.SD: "Sudan",
			territory.SG: "Singapur",
			territory.SI: "Sloveni",
			territory.SK: "Slovaki",
			territory.SL: "Serra Leon",
			territory.SN: "Senegal",
			territory.SO: "Somali",
			territory.SV: "Salvadoor",
			territory.TD: "Cad",
			territory.TG: "Togo",
			territory.TH: "Tailand",
		},
	}
}
