// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_sah() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "sah",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y 'сыл' MMMM d 'күнэ', EEEE", Long: "y, MMMM d", Medium: "y, MMM d", Short: "yy/M/d"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Тохс", Feb: "Олун", Mar: "Клн", Apr: "Мсу", May: "Ыам", Jun: "Бэс", Jul: "Отй", Aug: "Атр", Sep: "Блҕ", Oct: "Алт", Nov: "Сэт", Dec: "Ахс"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "Т", Feb: "О", Mar: "К", Apr: "М", May: "Ы", Jun: "Б", Jul: "О", Aug: "А", Sep: "Б", Oct: "А", Nov: "С", Dec: "А"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "тохсунньу", Feb: "олунньу", Mar: "кулун тутар", Apr: "муус устар", May: "ыам ыйа", Jun: "бэс ыйа", Jul: "от ыйа", Aug: "атырдьых ыйа", Sep: "балаҕан ыйа", Oct: "алтынньы", Nov: "сэтинньи", Dec: "ахсынньы"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "бс", Mon: "бн", Tue: "оп", Wed: "сэ", Thu: "чп", Fri: "бэ", Sat: "сб"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Б", Mon: "Б", Tue: "О", Wed: "С", Thu: "Ч", Fri: "Б", Sat: "С"}, Short: cldr.CalendarDayFormatNameValue{Sun: "бс", Mon: "бн", Tue: "оп", Wed: "сэ", Thu: "чп", Fri: "бэ", Sat: "сб"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "баскыһыанньа", Mon: "бэнидиэнньик", Tue: "оптуорунньук", Wed: "сэрэдэ", Thu: "чэппиэр", Fri: "Бээтиҥсэ", Sat: "субуота"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "ЭИ", PM: "ЭК"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "ЭИ", PM: "ЭК"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "ЭИ", PM: "ЭК"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AOA: cldr.Currency{DisplayName: "", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "", Symbol: "AWG"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "", Symbol: "CA$"},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "CRC"},
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
				currency.GBP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "HNL"},
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
				currency.NIO: cldr.Currency{DisplayName: "", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "", Symbol: "NZ$"},
				currency.PAB: cldr.Currency{DisplayName: "", Symbol: "PAB"},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "", Symbol: "₲"},
				currency.RON: cldr.Currency{DisplayName: "", Symbol: "lei"},
				currency.RUB: cldr.Currency{DisplayName: "Арассыыйа солкуобайа", Symbol: "₽"},
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
				currency.USD: cldr.Currency{DisplayName: "АХШ дуоллара", Symbol: "$"},
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
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AB:  "Абхаастыы",
			language.AF:  "Аппырыкааныстыы",
			language.ALE: "Алеуттуу",
			language.AM:  "Амхаардыы",
			language.AR:  "Араабтыы",
			language.AST: "Астуурдуу",
			language.AV:  "Аваардыы",
			language.AZ:  "Адьырбайдьаанныы",
			language.BE:  "Бөлөрүүстүү",
			language.BG:  "Булҕаардыы",
			language.BN:  "Бенгаллыы",
			language.BO:  "Тибиэттии",
			language.BS:  "Босныйалыы",
			language.CA:  "Каталаанныы",
			language.CE:  "Чэчиэннии",
			language.CKB: "Киин куурдуу",
			language.CS:  "Чиэхтии",
			language.DA:  "Даатскайдыы",
			language.DE:  "Ниэмэстии",
			language.EL:  "Гириэктии",
			language.EN:  "Ааҥыллыы",
			language.ES:  "Ыспаанныы",
			language.ET:  "Эстиэнийэлии",
			language.FA:  "Пиэристии",
			language.FI:  "Пииннии",
			language.FIL: "Пилипииннии",
			language.FR:  "Боронсуустуу",
			language.HU:  "Бэҥгиэрдии",
			language.HY:  "Эрмээннии",
			language.IT:  "Ытаалыйалыы",
			language.JA:  "Дьоппуоннуу",
			language.KA:  "Курусууннуу",
			language.KK:  "Хаһаахтыы",
			language.KO:  "Кэриэйдии",
			language.KY:  "Кыргыстыы",
			language.LA:  "Латыынныы",
			language.MN:  "Моҕуоллуу",
			language.MS:  "Малаайдыы",
			language.NE:  "Ньыпааллыы",
			language.NOG: "Нагаайдыы",
			language.PA:  "Пандьаабтыы",
			language.PT:  "Португааллыы",
			language.RO:  "Румыынныы",
			language.RU:  "Нууччалыы",
			language.SAH: "саха тыла",
			language.SK:  "Словаактыы",
			language.SQ:  "Албаанныы",
			language.TA:  "Тамыллыы",
			language.TE:  "Төлүгүлүү",
			language.TG:  "Тадьыыктыы",
			language.TT:  "Татаардыы",
			language.UG:  "Уйгуурдуу",
			language.UK:  "Украйыыньыстыы",
			language.UZ:  "Үзбиэктии",
			language.ZH:  "Кытайдыы",
			language.ZU:  "Зуулулуу",
		},
		Territories: cldr.Territories{
			territory.V_001: "Аан дойду",
			territory.V_002: "Аапырыка",
			territory.V_003: "Хотугу Эмиэрикэ",
			territory.V_005: "Соҕуруу Эмиэрикэ",
			territory.BR:    "Бразилия",
			territory.CA:    "Канаада",
			territory.CL:    "Чиили",
			territory.CN:    "Кытай",
			territory.CU:    "Кууба",
			territory.EE:    "Эстония",
			territory.FI:    "Финляндия",
			territory.GB:    "Улуу Британия",
			territory.IE:    "Ирландия",
			territory.IM:    "Мэн арыы",
			territory.IS:    "Исландия",
			territory.JM:    "Дьамаайка",
			territory.LT:    "Литва",
			territory.LV:    "Латвия",
			territory.LY:    "Лиибийэ",
			territory.MX:    "Миэксикэ",
			territory.NO:    "Норвегия",
			territory.RU:    "Арассыыйа",
			territory.SD:    "Судаан",
			territory.SE:    "Швеция",
			territory.US:    "Америка Холбоһуктаах Штааттара",
		},
	}
}
