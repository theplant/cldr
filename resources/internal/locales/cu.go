// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_cu() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "cu",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM 'л'. y.", Long: "y MMMM d", Medium: "y MMM d", Short: "y.MM.dd"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "і҆аⷩ҇", Feb: "феⷡ҇", Mar: "маⷬ҇", Apr: "а҆пⷬ҇", May: "маꙵ", Jun: "і҆ꙋⷩ҇", Jul: "і҆ꙋⷧ҇", Aug: "а҆́ѵⷢ҇", Sep: "сеⷫ҇", Oct: "ѻ҆кⷮ", Nov: "ноеⷨ", Dec: "деⷦ҇"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "І҆", Feb: "Ф", Mar: "М", Apr: "А҆", May: "М", Jun: "І҆", Jul: "І҆", Aug: "А҆", Sep: "С", Oct: "Ѻ҆", Nov: "Н", Dec: "Д"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "і҆аннꙋа́рїй", Feb: "феврꙋа́рїй", Mar: "ма́ртъ", Apr: "а҆прі́ллїй", May: "ма́їй", Jun: "і҆ꙋ́нїй", Jul: "і҆ꙋ́лїй", Aug: "а҆́ѵгꙋстъ", Sep: "септе́мврїй", Oct: "ѻ҆ктѡ́врїй", Nov: "ное́мврїй", Dec: "деке́мврїй"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ндⷧ҇ѧ", Mon: "пнⷣе", Tue: "втоⷬ҇", Wed: "срⷣе", Thu: "чеⷦ҇", Fri: "пѧⷦ҇", Sat: "сꙋⷠ҇"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Н", Mon: "П", Tue: "В", Wed: "С", Thu: "Ч", Fri: "П", Sat: "С"}, Short: cldr.CalendarDayFormatNameValue{Sun: "ндⷧ҇ѧ", Mon: "пнⷣе", Tue: "втоⷬ҇", Wed: "срⷣе", Thu: "чеⷦ҇", Fri: "пѧⷦ҇", Sat: "сꙋⷠ҇"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "недѣ́лѧ", Mon: "понедѣ́льникъ", Tue: "вто́рникъ", Wed: "среда̀", Thu: "четверто́къ", Fri: "пѧто́къ", Sat: "сꙋббѡ́та"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "ДП", PM: "ПП"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "ДП", PM: "ПП"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "ДП", PM: "ПП"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "", Percent: "#,##0\u00a0%"},
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
				currency.BRL: cldr.Currency{DisplayName: "бразі́льскїй реа́лъ", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "бѣлорꙋ́сскїй рꙋ́бль", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "бѣлорꙋ́сскїй рꙋ́бль (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "", Symbol: "CA$"},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "хи́нскїй ю҆а́нь", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.EGP: cldr.Currency{DisplayName: "", Symbol: "E£"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.EUR: cldr.Currency{DisplayName: "є҆́ѵрѡ", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "а҆нглі́йскїй фꙋ́нтъ сте́рлингѡвъ", Symbol: "£"},
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
				currency.INR: cldr.Currency{DisplayName: "і҆нді́йскаѧ рꙋ́пїѧ", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "ꙗ҆пѡ́нскаѧ і҆е́на", Symbol: "JP¥"},
				currency.KGS: cldr.Currency{DisplayName: "кирги́зскїй сꙋ́мъ", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "каза́хскаѧ деньга̀", Symbol: "₸"},
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
				currency.RUB: cldr.Currency{DisplayName: "рѡссі́йскїй рꙋ́бль", Symbol: "₽"},
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
				currency.UAH: cldr.Currency{DisplayName: "ᲂу҆краи́нскаѧ гри́вна", Symbol: "₴"},
				currency.USD: cldr.Currency{DisplayName: "а҆мерїка́нскїй до́лларъ", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "невѣ́домое пла́тное сре́дство", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Languages: cldr.Languages{
			language.AB:      "а҆бха́зскїй",
			language.AR:      "а҆ра́вскїй",
			language.AZ:      "а҆зербайджа́нскїй",
			language.BE:      "бѣлорꙋ́сскїй",
			language.BG:      "бо́лгарскїй",
			language.CU:      "церковнослове́нскїй",
			language.DE:      "нѣме́цкїй",
			language.DE_AT:   "а҆ѵстрі́йскїй нѣме́цкїй",
			language.DE_CH:   "є҆лветі́йскїй нѣме́цкїй",
			language.EL:      "є҆́ллинскїй",
			language.EN:      "а҆нглі́йскїй",
			language.EN_AU:   "а҆ѵстралі́йскїй а҆нглі́йскїй",
			language.EN_CA:   "кана́дскїй а҆нглі́йскїй",
			language.EN_GB:   "а҆нглі́йскїй (вели́каѧ брїта́нїа)",
			language.EN_US:   "а҆нглі́йскїй (асд)",
			language.ES:      "і҆спа́нскїй",
			language.ES_419:  "латїноамерїка́нскїй і҆спа́нскїй",
			language.ES_ES:   "є҆ѵрѡпе́йскїй і҆спа́нскїй",
			language.ES_MX:   "і҆спанскїй (ме́ѯїка)",
			language.ET:      "є҆сто́нскїй",
			language.FI:      "фі́нскїй",
			language.FR:      "францꙋ́зскїй",
			language.FR_CA:   "кана́дскїй францꙋ́зскїй",
			language.FR_CH:   "є҆лветі́йскїй францꙋ́зскїй",
			language.HE:      "є҆вре́йскїй",
			language.HY:      "а҆рме́нскїй",
			language.IT:      "і҆талїа́нскїй",
			language.JA:      "ꙗ҆пѡ́нскїй",
			language.KA:      "і҆́верскїй",
			language.KK:      "каза́хскїй",
			language.LA:      "латі́нскїй",
			language.LT:      "лїто́вскїй",
			language.LV:      "латві́йскїй",
			language.PT:      "портога́льскїй",
			language.PT_BR:   "бразі́льскїй портога́льскїй",
			language.PT_PT:   "є҆ѵрѡпе́йскїй портога́льскїй",
			language.RO:      "дакорꙋмы́нскїй",
			language.RO_MD:   "молда́вскїй",
			language.RU:      "рꙋ́сскїй",
			language.SR:      "се́рбскїй",
			language.UK:      "ᲂу҆краи́нскїй",
			language.UND:     "невѣ́домый ѧ҆зы́къ",
			language.ZH:      "хи́нскїй",
			language.ZH_HANS: "ᲂу҆проще́нный хи́нскїй",
			language.ZH_HANT: "традїцїо́нный хи́нскїй",
		},
		Territories: cldr.Territories{
			territory.AU: "А҆ѵстралі́ѧ",
			territory.BR: "бразі́лїа",
			territory.BY: "бѣ́лаѧ рꙋ́сь",
			territory.CA: "Кана́да",
			territory.CN: "хи́нскаѧ страна̀",
			territory.DE: "герма́нїа",
			territory.DK: "Дані́ѧ",
			territory.FR: "га́ллїа",
			territory.GB: "Вели́каѧ брїта́нїа",
			territory.IN: "і҆́ндїа",
			territory.IT: "і҆та́лїа",
			territory.JP: "ꙗ҆пѡ́нїа",
			territory.KG: "кирги́зїа",
			territory.KZ: "казахста́нъ",
			territory.MX: "Ме́ѯїко",
			territory.RU: "рѡссі́а",
			territory.UA: "ᲂу҆краи́на",
			territory.US: "а҆мерїка̑нскїѧ соединє́нныѧ держа̑вы",
			territory.ZZ: "невѣ́домаѧ страна̀",
		},
	}
}
