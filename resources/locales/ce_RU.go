// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_ce_RU() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y MMMM d, EEEE", Long: "y MMMM d", Medium: "y MMM d", Short: "y-MM-dd"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "янв", Feb: "фев", Mar: "мар", Apr: "апр", May: "май", Jun: "июн", Jul: "июл", Aug: "авг", Sep: "сен", Oct: "окт", Nov: "ноя", Dec: "дек"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "Я", Feb: "Ф", Mar: "М", Apr: "А", May: "М", Jun: "И", Jul: "И", Aug: "А", Sep: "С", Oct: "О", Nov: "Н", Dec: "Д"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "январь", Feb: "февраль", Mar: "март", Apr: "апрель", May: "май", Jun: "июнь", Jul: "июль", Aug: "август", Sep: "сентябрь", Oct: "октябрь", Nov: "ноябрь", Dec: "декабрь"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "кӀи", Mon: "ор", Tue: "ши", Wed: "кха", Thu: "еа", Fri: "пӀе", Sat: "шуо"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "кӀ", Mon: "о", Tue: "ш", Wed: "кх", Thu: "е", Fri: "пӀ", Sat: "ш"}, Short: cldr.CalendarDayFormatNameValue{Sun: "кӀи", Mon: "ор", Tue: "ши", Wed: "кха", Thu: "еа", Fri: "пӀе", Sat: "шуо"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "кӀира", Mon: "оршот", Tue: "шинара", Wed: "кхаара", Thu: "еара", Fri: "пӀераска", Sat: "шуот"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_ce]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Дирхам ӀЦЭ", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "ОвхӀан-пачхьалкхан афгани", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "Албанин лек", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Эрмалойчоьнан драм", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Нидерландин Антилин гульден", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Анголан кванза", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "Аргентинан песо", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "Австралин доллар", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Арубан флорин", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "Азербайджанан манат", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "Боснин а, Герцеговинан а хийцалун марка", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "Барбадосан доллар", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Бангладешан така", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "Болгарин лев", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "Бахрейнан динар", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Бурундин франк", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Бермудан доллар", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Брунейн доллар", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Боливин боливиано", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "Бразилин реал", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Багаман доллар", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Бутанан нгултрум", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "Ботсванан пула", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "Белоруссин сом", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Белоруссин сом (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "Белизин доллар", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Канадан доллар", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Конголезин франк", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "Швейцарин франк", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "Чилин песо", Symbol: "CLP"},
				currency.CNY: cldr.Currency{DisplayName: "Китайн юань", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Колумбин песо", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "Костарикан колон", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "Кубан хийцалун песо", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Кубан песо", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Кабо-Верден эскудо", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "Чехин крона", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "Джибутин франк", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Данин крона", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Доминикан песо", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Алжиран динар", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "Мисаран фунт", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Эритрейн накфа", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Эфиопин быр", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Евро", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Фиджин доллар", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Фолклендан гӀайренийн фунт", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "Англин фунт", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Гуьржийчоьнан лари", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "Ганан седи", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Гибралтаран фунт", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Гамбин даласи", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Гвинейн франк", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "Гватемалан кетсаль", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "Гайанан доллар", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Гонконган доллар", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Гондурасан лемпира", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "Хорватин куна", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Гаитин гурд", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Венгрин форинт", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Индонезин рупи", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "Израилан керла шекель", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Индин рупи", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Ӏиракъан динар", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "ГӀажарийчоьнан риал", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "Исландин крона", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "Ямайн доллар", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Урданан динар", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Японин иена", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Кенин шиллинг", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Киргизин сом", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Камбоджан риель", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Коморийн гӀайренийн франк", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Къилбаседа Корейн вона", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "Къилба Корейн вона", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Кувейтан динар", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Кайманийн гӀайренийн доллар", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Кхазакхстанан тенге", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Лаосан кип", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Ливанан фунт", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Шри-Ланкан рупи", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Либерин доллар", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Ливин динар", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Мароккон дирхам", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "Молдавин лей", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Малагасийн ариари", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "Македонин динар", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "Мьянман кьят", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Монголин тугрик", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Макаон патака", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Мавританин уги (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Мавританин уги", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Маврикин рупи", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Мальдивийн руфи", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Малавин квача", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Мексикан песо", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Малайзин ринггит", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "Мозамбикан метикал", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Намибин доллар", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Нигерин найра", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "Никарагуан кордоба", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "Норвегин крона", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Непалан рупи", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Керла Зеландин доллар", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Оманан риал", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Панаман бальбоа", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "Перун соль", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "Папуа — Керла Гвинейн кина", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Филиппинийн песо", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Пакистанан рупи", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Польшан злотый", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "Парагвайн гуарани", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Катаран риал", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "Румынин лей", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Сербин динар", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Российн сом", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Руандан франк", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "СаӀудийн Ӏаьрбийчоьнан риал", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Соломонан гӀайренийн доллар", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Сейшелан рупи", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "Суданан фунт", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "Швецин крона", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Сингапуран доллар", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Сийлахьчу Еленин гӀайрен фунт", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "Леоне", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Сомалин шиллинг", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Суринаман доллар", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "Къилба Суданан фунт", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "Сан-Томен а, Принсипин а добра (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "Сан-Томен а, Принсипин а добра", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "Шеман фунт", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Свазилендан лилангени", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Таиландан бат", Symbol: "THB"},
				currency.TJS: cldr.Currency{DisplayName: "Таджикистанан сомони", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "Туркменин керла манат", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Тунисан динар", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Тонганан паанга", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "Туркойчоьнан лира", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Тринидадан а, Тобагон а доллар", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Тайванан керла доллар", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Танзанин шиллинг", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Украинан гривна", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "Угандан шиллинг", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "АЦШн доллар", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "Уругвайн песо", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Узбекистанан сом", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "Венесуэлан боливар (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Венесуэлан боливар", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "Вьетнаман донг", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Вануатун вату", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Самоанан тала", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Юккъерчу Африкан КФА франк", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Малхбален Карибийн доллар", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Малхбузен Африкан КФА франк", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "Французийн Тийна океанан франк", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "девзаш доцу я лелаш доцу ахча", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Йеменан риал", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "Къилба-Африкин рэнд", Symbol: "ZAR"},
				currency.ZMW: cldr.Currency{DisplayName: "Замбин квача", Symbol: "ZMW"},
			},
		},
	}
}
