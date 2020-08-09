// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_bg() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y 'г'.", Long: "d MMMM y 'г'.", Medium: "d.MM.y 'г'.", Short: "d.MM.yy 'г'."}, Time: cldr.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "Гринуич{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "яну", Feb: "фев", Mar: "март", Apr: "апр", May: "май", Jun: "юни", Jul: "юли", Aug: "авг", Sep: "сеп", Oct: "окт", Nov: "ное", Dec: "дек"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "я", Feb: "ф", Mar: "м", Apr: "а", May: "м", Jun: "ю", Jul: "ю", Aug: "а", Sep: "с", Oct: "о", Nov: "н", Dec: "д"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "януари", Feb: "февруари", Mar: "март", Apr: "април", May: "май", Jun: "юни", Jul: "юли", Aug: "август", Sep: "септември", Oct: "октомври", Nov: "ноември", Dec: "декември"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "нд", Mon: "пн", Tue: "вт", Wed: "ср", Thu: "чт", Fri: "пт", Sat: "сб"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "н", Mon: "п", Tue: "в", Wed: "с", Thu: "ч", Fri: "п", Sat: "с"}, Short: cldr.CalendarDayFormatNameValue{Sun: "нд", Mon: "пн", Tue: "вт", Wed: "ср", Thu: "чт", Fri: "пт", Sat: "сб"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "неделя", Mon: "понеделник", Tue: "вторник", Wed: "сряда", Thu: "четвъртък", Fri: "петък", Sat: "събота"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "am", PM: "pm"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "am", PM: "pm"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "am", PM: "pm"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_bg]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "0.00\u00a0¤", CurrencyAccounting: "0.00\u00a0¤;(0.00\u00a0¤)", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "Андорска песета", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "Дирхам на Обединените арабски емирства", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "Афганистански афган (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Афганистански афган", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "Албански лек", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Арменски драм", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Антилски гулден", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Анголска кванза", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "Анголска кванца (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "Анголска нова кванца (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "Анголска нова кванца (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "Аржентински австрал", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "Аржентинско песо (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "Аржентинско песо", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "Австрийски шилинг", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "Австралийски долар", Symbol: "AUD"},
				currency.AWG: cldr.Currency{DisplayName: "Арубски флорин", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "Азербайджански манат (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Азербайджански манат", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "Босна и Херцеговина-динар", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Босненска конвертируема марка", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "Барбадоски долар", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Бангладешка така", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "Белгийски франк (конвертируем)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "Белгийски франк", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "Белгийски франк (финансов)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "Български конвертируем лев (1962–1999)", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "Български лев", Symbol: "лв."},
				currency.BHD: cldr.Currency{DisplayName: "Бахрейнски динар", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Бурундийски франк", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Бермудски долар", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Брунейски долар", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Боливийско боливиано", Symbol: "BOB"},
				currency.BOP: cldr.Currency{DisplayName: "Боливийско песо", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "Боливийски мвдол", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "Бразилско ново крузейро (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "Бразилско крозадо", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "Бразилско крузейро (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "Бразилски реал", Symbol: "BRL"},
				currency.BRN: cldr.Currency{DisplayName: "Бразилско ново крозадо", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "Бразилско крузейро", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "Бахамски долар", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Бутански нгултрум", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "Бирмански киат", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Ботсванска пула", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "Беларуска нова рубла (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "Беларуска рубла", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Беларуска рубла (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "Белизийски долар", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Канадски долар", Symbol: "CAD"},
				currency.CDF: cldr.Currency{DisplayName: "Конгоански франк", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "WIR евро", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Швейцарски франк", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "WIR франк", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "Условна разчетна единица на Чили", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Чилийско песо", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Китайски ренминби юан (offshore)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "Китайски юан", Symbol: "CNY"},
				currency.COP: cldr.Currency{DisplayName: "Колумбийско песо", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "Колумбийска единица на реалната стойност", Symbol: ""},
				currency.CRC: cldr.Currency{DisplayName: "Костарикански колон", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "Стар сръбски динар", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "Чехословашка конвертируема крона", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "Кубинско конвертируемо песо", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Кубинско песо", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Ескудо на Кабо Верде", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "Кипърска лира", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Чешка крона", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "Източногерманска марка", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "Германска марка", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "Джибутски франк", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Датска крона", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Доминиканско песо", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Алжирски динар", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "Еквадорско сукре", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "Еквадорска банкова единица", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "Естонска крона", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Египетска лира", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Еритрейска накфа", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "Испанска песета", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Етиопски бир", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Евро", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Финландска марка", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "Фиджийски долар", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Фолклендска лира", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "Френски франк", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "Британска лира", Symbol: "GBP"},
				currency.GEK: cldr.Currency{DisplayName: "Грузински купон", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "Грузински лари", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "Ганайско седи (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Ганайско седи", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Гибралтарска лира", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Гамбийско даласи", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Гвинейски франк", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "Гвинейска сили", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "Екваториално гвинейско еквеле", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "Гръцка драхма", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "Гватемалски кетцал", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "Ескудо от Португалска Гвинея", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "Гвинея-Бисау песо", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "Гаянски долар", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Хонконгски долар", Symbol: "HKD"},
				currency.HNL: cldr.Currency{DisplayName: "Хондураска лемпира", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "Хърватски динар", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "Хърватска куна", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Хаитски гурд", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Унгарски форинт", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Индонезийска рупия", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "Ирландска лира", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "Израелска лира", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "Израелски нов шекел", Symbol: "ILS"},
				currency.INR: cldr.Currency{DisplayName: "Индийска рупия", Symbol: "INR"},
				currency.IQD: cldr.Currency{DisplayName: "Иракски динар", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Ирански риал", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "Исландска крона", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "Италианска лира", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "Ямайски долар", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Йордански динар", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Японска йена", Symbol: "JPY"},
				currency.KES: cldr.Currency{DisplayName: "Кенийски шилинг", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Киргизстански сом", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Камбоджански риел", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Коморски франк", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Севернокорейски вон", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "Южнокорейски вон", Symbol: "KRW"},
				currency.KWD: cldr.Currency{DisplayName: "Кувейтски динар", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Кайманов долар", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Казахстанско тенге", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Лаоски кип", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Ливанска лира", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Шриланкска рупия", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Либерийски долар", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Лесотско лоти", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Литовски литас", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "Литовски талон", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "Люксембургски франк", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "Латвийски лат", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "Латвийска рубла", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "Либийски динар", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Марокански дирхам", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "Марокански франк", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Молдовско леу", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Малгашко ариари", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "Малгашки франк - Мадагаскар", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "Македонски денар", Symbol: "MKD"},
				currency.MLF: cldr.Currency{DisplayName: "Малийски франк", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Мианмарски кият", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Монголски тугрик", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Патака на Макао", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Мавританска угия (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Мавританска угия", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "Малтийска лира", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Маврицийска рупия", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Малдивска руфия", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Малавийска квача", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Мексиканско песо", Symbol: "MXN"},
				currency.MXP: cldr.Currency{DisplayName: "Мексиканско сребърно песо (1861–1992)", Symbol: "MXP"},
				currency.MXV: cldr.Currency{DisplayName: "Мексиканска конвертируема единица (UDI)", Symbol: "MXV"},
				currency.MYR: cldr.Currency{DisplayName: "Малайзийски рингит", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "Мозамбикско ескудо", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "Мозамбикски метикал (1980–2006)", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Мозамбикски метикал", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Намибийски долар", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Нигерийска найра", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "Никарагуанска кордоба (1988–1991)", Symbol: "NIC"},
				currency.NIO: cldr.Currency{DisplayName: "Никарагуанска кордоба", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "Холандски гулден", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "Норвежка крона", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Непалска рупия", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Новозеландски долар", Symbol: "NZD"},
				currency.OMR: cldr.Currency{DisplayName: "Омански риал", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Панамска балбоа", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "Перуанско инти", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Перуански сол", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "Перуански сол (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Папуа-новогвинейска кина", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Филипинско песо", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Пакистанска рупия", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Полска злота", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "Полска злота (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "Португалско ескудо", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "Парагвайско гуарани", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Катарски риал", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "Родезийски долар", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "Стара румънска лея", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Румънска лея", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Сръбски динар", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Руска рубла", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "Руска рубла (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Руандски франк", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "саудитски риал", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Долар на Соломоновите острови", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Сейшелска рупия", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "Судански динар", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Суданска лира", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "Шведска крона", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Сингапурски долар", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Лира на Света Елена", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "Словенски толар", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "Словашка крона", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Сиералеонско леоне", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Сомалийски шилинг", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Суринамски долар", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "Суринамски гилдер", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "Южносуданска лира", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "Добра на Сао Томе и Принсипи (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "Добра на Сао Томе и Принсипи", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "Съветска рубла", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "Салвадорски колон", Symbol: "SVC"},
				currency.SYP: cldr.Currency{DisplayName: "Сирийска лира", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Свазилендски лилангени", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Тайландски бат", Symbol: "THB"},
				currency.TJR: cldr.Currency{DisplayName: "Таджикистанска рубла", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "Таджикистански сомони", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "Туркменистански манат", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Туркменски манат", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Тунизийски динар", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Тонганска паанга", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "Тиморско ескудо", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "Турска лира (1922–2005)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "Турска лира", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Долар на Тринидад и Тобаго", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Тайвански долар", Symbol: "TWD"},
				currency.TZS: cldr.Currency{DisplayName: "Танзанийски шилинг", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Украинска хривня", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "Украински карбованец", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "Угандийски шилинг (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "Угандски шилинг", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "Щатски долар", Symbol: "щ.д."},
				currency.USN: cldr.Currency{DisplayName: "", Symbol: "USN"},
				currency.USS: cldr.Currency{DisplayName: "", Symbol: "USS"},
				currency.UYI: cldr.Currency{DisplayName: "Уругвайско песо (индекс на инфлацията)", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "Уругвайско песо (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "Уругвайско песо", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Узбекски сум", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "Венецуелски боливар (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "Венецуелски боливар", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Венецуелски боливар (VES)", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Виетнамски донг", Symbol: "VND"},
				currency.VUV: cldr.Currency{DisplayName: "Вануатско вату", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Самоанска тала", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Централноафрикански франк", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "Сребро", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "Злато", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "Европейска съставна единица", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "Европейска валутна единица", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "Европейска единица по сметка (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "Европейска единица по сметка (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "Източнокарибски долар", Symbol: "XCD"},
				currency.XDR: cldr.Currency{DisplayName: "Специални права на тираж", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "Еку на ЕИО", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "Френски златен франк", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "Западноафрикански франк", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "Паладий", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "CFP франк", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "Платина", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "Код резервиран за целите на тестване", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "Непозната валута", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "Йеменски динар", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "Йеменски риал", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "Югославски твърд динар", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "Югославски динар", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "Югославски конвертируем динар", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "Южноафрикански ранд (финансов)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Южноафрикански ранд", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Замбийска квача (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Замбийска куача", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "Заирско ново зайре", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "Заирско зайре", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "Зимбабвийски долар", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "Зимбабвийски долар (2009)", Symbol: ""},
			},
		},
	}
}
