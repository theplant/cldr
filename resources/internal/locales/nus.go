// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_nus() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "nus",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/MM/y"}, Time: cldr.CalendarDateFormat{Full: "zzzz h:mm:ss a", Long: "z h:mm:ss a", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Tiop", Feb: "Pɛt", Mar: "Duɔ̱ɔ̱", Apr: "Guak", May: "Duä", Jun: "Kor", Jul: "Pay", Aug: "Thoo", Sep: "Tɛɛ", Oct: "Laa", Nov: "Kur", Dec: "Tid"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "T", Feb: "P", Mar: "D", Apr: "G", May: "D", Jun: "K", Jul: "P", Aug: "T", Sep: "T", Oct: "L", Nov: "K", Dec: "T"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Tiop thar pɛt", Feb: "Pɛt", Mar: "Duɔ̱ɔ̱ŋ", Apr: "Guak", May: "Duät", Jun: "Kornyoot", Jul: "Pay yie̱tni", Aug: "Tho̱o̱r", Sep: "Tɛɛr", Oct: "Laath", Nov: "Kur", Dec: "Tio̱p in di̱i̱t"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Cäŋ", Mon: "Jiec", Tue: "Rɛw", Wed: "Diɔ̱k", Thu: "Ŋuaan", Fri: "Dhieec", Sat: "Bäkɛl"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "C", Mon: "J", Tue: "R", Wed: "D", Thu: "Ŋ", Fri: "D", Sat: "B"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Cäŋ kuɔth", Mon: "Jiec la̱t", Tue: "Rɛw lätni", Wed: "Diɔ̱k lätni", Thu: "Ŋuaan lätni", Fri: "Dhieec lätni", Sat: "Bäkɛl lätni"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "RW", PM: "TŊ"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "RW", PM: "TŊ"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: "#,##0%"},
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
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.EGP: cldr.Currency{DisplayName: "", Symbol: "E£"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.EUR: cldr.Currency{DisplayName: "", Symbol: "€"},
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
			language.AK:  "Thok aka̱ni",
			language.AM:  "Thok bunyni",
			language.AR:  "Thok Jalabni",
			language.BE:  "Thok bälärutha",
			language.BG:  "Thok bälga̱a̱riani",
			language.BN:  "Thok bängali",
			language.CS:  "Thok cik",
			language.DE:  "Thok jarmani",
			language.EL:  "Thok girikni",
			language.EN:  "Thok liŋli̱thni",
			language.ES:  "Thok i̱thpaaniani",
			language.FA:  "Thok perthiani",
			language.FR:  "Thok pɔrɔthani",
			language.HA:  "Thok ɣowthani",
			language.HI:  "Thok ɣändini",
			language.HU:  "Thok ɣänga̱a̱riɛni",
			language.ID:  "Thok indunithiani",
			language.IG:  "Thok i̱gboni",
			language.IT:  "Thok i̱taliani",
			language.JA:  "Thok japanni",
			language.JV:  "Thok jabanithni",
			language.KM:  "Thok kameeri",
			language.KO:  "Thok kuriani",
			language.MS:  "Thok mayɛyni",
			language.MY:  "Thok bormi̱thni",
			language.NE:  "Thok napalni",
			language.NL:  "Thok da̱c",
			language.NUS: "Thok Nath",
			language.PA:  "Thok puɔnjabani",
			language.PL:  "Thok pölicni",
			language.PT:  "Thok puɔtigali",
			language.RO:  "Thok ji̱ röm",
			language.RU:  "Thok ra̱ciaani",
			language.RW:  "Thok ruaandani",
			language.SO:  "Thok thomaaliani",
			language.SV:  "Thok i̱thwidicni",
			language.TA:  "Thok tamilni",
			language.TH:  "Thok tayni",
			language.TR:  "Thok turkicni",
			language.UK:  "Thok ukeraanini",
			language.UR:  "Thok udoni",
			language.VI:  "Thok betnaamni",
			language.YO:  "Thok yurubani",
			language.ZH:  "Thok cayna",
			language.ZU:  "Thok dhuluni",
		},
		Territories: cldr.Territories{
			territory.AD: "Andora",
			territory.AF: "Abganithtan",
			territory.AG: "Antiguaa kɛnɛ Barbuda",
			territory.AI: "Aŋguɛla",
			territory.AL: "Albänia",
			territory.AM: "Aɛrmänia",
			territory.AO: "Aŋgola",
			territory.AR: "Aɛrgentin",
			territory.AS: "Amerika thamow",
			territory.AT: "Athtɛria",
			territory.AU: "Athɔra̱lia",
			territory.AW: "Aruba",
			territory.AZ: "Adhe̱rbe̱ja̱n",
			territory.BA: "Bothnia kɛnɛ ɣärgobinia",
			territory.BB: "Bärbadoth",
			territory.BD: "Bengeladiec",
			territory.BE: "Be̱lgim",
			territory.BF: "Burkinɛ pa̱thu",
			territory.BG: "Bulga̱a̱ria",
			territory.BH: "Ba̱reen",
			territory.BI: "Burundi",
			territory.BJ: "Be̱ni̱n",
			territory.BM: "Be̱rmudaa",
			territory.BN: "Burunɛy",
			territory.BO: "Bulibia",
			territory.BR: "Bäraadhiil",
			territory.BS: "Bämuɔth",
			territory.BT: "Buta̱n",
			territory.BW: "Bothiwaana",
			territory.BY: "Be̱lɛruth",
			territory.BZ: "Bilidha",
			territory.CA: "Känɛda",
			territory.CF: "Cɛntrɔl aprika repuɔblic",
			territory.CG: "Kɔŋgɔ",
			territory.CI: "Kodibo̱o̱",
			territory.CK: "Kuk ɣa̱ylɛn",
			territory.CL: "Cili̱",
			territory.CM: "Kɛmɛrun",
			territory.CN: "Cayna",
			territory.CO: "Kolombia",
			territory.CR: "Kothtirika",
			territory.CV: "Kɛp bedi ɣa̱ylɛn",
			territory.DZ: "Algeria",
			territory.HR: "Korwaatia",
			territory.IO: "Burutic ɣe̱ndian oce̱n",
			territory.KH: "Kombodia",
			territory.KM: "Komruth",
			territory.KY: "Kaymɛn ɣa̱ylɛn",
			territory.SD: "Sudan",
			territory.TD: "Ca̱d",
			territory.VG: "Burutic dhuɔ̱ɔ̱l be̱rgin",
		},
	}
}
