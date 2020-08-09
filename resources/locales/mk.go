// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_mk() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "mk",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d.M.y", Short: "d.M.yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "јан.", Feb: "фев.", Mar: "мар.", Apr: "апр.", May: "мај", Jun: "јун.", Jul: "јул.", Aug: "авг.", Sep: "септ.", Oct: "окт.", Nov: "ноем.", Dec: "дек."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "ј", Feb: "ф", Mar: "м", Apr: "а", May: "м", Jun: "ј", Jul: "ј", Aug: "а", Sep: "с", Oct: "о", Nov: "н", Dec: "д"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "јануари", Feb: "февруари", Mar: "март", Apr: "април", May: "мај", Jun: "јуни", Jul: "јули", Aug: "август", Sep: "септември", Oct: "октомври", Nov: "ноември", Dec: "декември"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "нед.", Mon: "пон.", Tue: "вто.", Wed: "сре.", Thu: "чет.", Fri: "пет.", Sat: "саб."}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "н", Mon: "п", Tue: "в", Wed: "с", Thu: "ч", Fri: "п", Sat: "с"}, Short: cldr.CalendarDayFormatNameValue{Sun: "нед.", Mon: "пон.", Tue: "вто.", Wed: "сре.", Thu: "чет.", Fri: "пет.", Sat: "саб."}, Wide: cldr.CalendarDayFormatNameValue{Sun: "недела", Mon: "понеделник", Tue: "вторник", Wed: "среда", Thu: "четврток", Fri: "петок", Sat: "сабота"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "претпл.", PM: "попл."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "претпл.", PM: "попл."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "претпладне", PM: "попладне"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_mk]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "Андорска Пезета", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "Дирхам на Обединети Арапски Емирати", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "Авгани (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Авганистански авгани", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "Албански лек", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Ерменски драм", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Холандски антилски гилдер", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Анголска кванза", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "Анголска Кванза (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "Анголска нова Кванза (1990–2000)", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "Аргентински Пезос (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "Аргентински пезос", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "Австралиски Шилинг", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "Австралиски долар", Symbol: "AUD"},
				currency.AWG: cldr.Currency{DisplayName: "Арубиски флорин", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "Азербејџански манат", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "Босанско-Херцеговски Динар", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Босанско-херцеговска конвертибилна марка", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "Барбадоски долар", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Бангладешка така", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "Белгиски Франк (конвертибилен)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "Белгиски Франк", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "Белгиски Франк (финансиски)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "Бугарски цврст лев", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "Бугарски лев", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "Бахреински динар", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Бурундиски франк", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Бермудски долар", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Брунејски долар", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Боливиски боливиано", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "Бразилски реал", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Бахамски долар", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Бутански нгултрум", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "Боцванска пула", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "Белоруска нова рубља (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "Белоруска рубља", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Белоруска рубља (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "Белизиски долар", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Канадски долар", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Конголски франк", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "Швајцарски франк", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "Чилеански пезос", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Кинески јуан (офшор)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "Кинески јуан", Symbol: "CNY"},
				currency.COP: cldr.Currency{DisplayName: "Колумбиски пезос", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "Костарикански колон", Symbol: "CRC"},
				currency.CSK: cldr.Currency{DisplayName: "Чехословачка цврста корона", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "Кубански пезос (конвертибилен)", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Кубански пезос", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Зелено’ртски ескудо", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "Кипарска фунта", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Чешка круна", Symbol: "CZK"},
				currency.DEM: cldr.Currency{DisplayName: "Германска Марка", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "Џибутски франк", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Данска круна", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Доминикански пезос", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Алжирски динар", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "Египетска фунта", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Еритрејска накфа", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "Шпанска Пезета", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Етиописки бир", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Евро", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Финска марка", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "Фиџиски долар", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Фолкландска фунта", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "Француски франк", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "Британска фунта", Symbol: "GBP"},
				currency.GEL: cldr.Currency{DisplayName: "Грузиски лари", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "Ганајски Седи", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Гански седи", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Гибралтарска фунта", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Гамбиски даласи", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Гвинејски франк", Symbol: "GNF"},
				currency.GRD: cldr.Currency{DisplayName: "Грчка драхма", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "Гватемалски кветцал", Symbol: "GTQ"},
				currency.GWP: cldr.Currency{DisplayName: "Гвинејски Бисау пезос", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "Гвајански долар", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Хонгконшки долар", Symbol: "HKD"},
				currency.HNL: cldr.Currency{DisplayName: "Хондурска лемпира", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "Хрватски динар", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "Хрватска куна", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Хаитски гурд", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Унгарска форинта", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Индонезиска рупија", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "Ирска фунта", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "Изрелска фунта", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "Израелски нов шекел", Symbol: "ILS"},
				currency.INR: cldr.Currency{DisplayName: "Индиска рупија", Symbol: "INR"},
				currency.IQD: cldr.Currency{DisplayName: "Ирачки динар", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Ирански риал", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "Исландска крона", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "Италијанска лира", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "Јамајкански долар", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Јордански динар", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Јапонски јен", Symbol: "JPY"},
				currency.KES: cldr.Currency{DisplayName: "Кениски шилинг", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Киргистански сом", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Камбоџиски рел", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Коморски франк", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Северно корејски вон", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "Јужно корејски вон", Symbol: "KRW"},
				currency.KWD: cldr.Currency{DisplayName: "Кувајтски динар", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Долар на кајмански острови", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Казахстанска тенга", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Лаоски кип", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Либанска фунта", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Шриланканска рупија", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Либериски долар", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Лесотско лоти", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Литваниска лита", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "Литваниски литаз", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "Луксембуршки франк", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "Латвијски лат", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "Латвијска рубља", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "Либијски динар", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Марокански дирхам", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "Марокански франк", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Молдавски леу", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Малагасиски ариари", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "Македонски денар", Symbol: "ден."},
				currency.MLF: cldr.Currency{DisplayName: "Малски франк", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Мјанмарски киат", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Монголиски тугрик", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Макао патака", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Мавританска угија (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Мавританска угија", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "Малтешка лира", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "Малтешка фунта", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Мавританска рупија", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Малдивиска руфија", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Малависка квача", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Мексикански пезос", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "Мексикански сребрен пезос (1861–1992)", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "Малезиски рингит", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "Мозамбиско ескудо", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "Стар мозамбиски метикал", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Мозамбиски метикал", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Намибиски долар", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Нигериска наира", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "Никарагванска кордоба (1988–1991)", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "Никарагванска кордоба", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "Холандски гилдер", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "Норвешка круна", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Непалска рупија", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Новозеландски долар", Symbol: "NZD"},
				currency.OMR: cldr.Currency{DisplayName: "Омански риал", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Панамска балбоа", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "Перуански сол", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "Перуански сол (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Папуа новогвинејска кина", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Филипински пезос", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Пакистанска рупија", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Полска злота", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "Полска злота (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "Португалско ескудо", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "Парагвајска гуарана", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Катарски риал", Symbol: "QAR"},
				currency.ROL: cldr.Currency{DisplayName: "Романска леи (1952–2006)", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Романски леу", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Српски динар", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Руска рубља", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "Руска рубља (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Руандски франк", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Саудиски ријал", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Соломонски долар", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Сејшелска рупија", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "Стар судански динар", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Суданска фунта", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "Стара суданска фунта", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Шведска круна", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Сингапурски долар", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Фунта на Света Елена", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "Словенечки толар", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "Словачка круна", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Сиералеонско леоне", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Сомалијски шилинг", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Суринамски долар", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "Суринамски гилдер", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "Јужносуданска фунта", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "Добра на Саун Томе и Принсип (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "Добра на Сао Томе и Принсипе", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "Советска рубља", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "Салвадорски колон", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "Сиријска фунта", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Свазилендски лиланген", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Таи бат", Symbol: "THB"},
				currency.TJR: cldr.Currency{DisplayName: "Таџикистанска рубља", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "Таџикистански сомони", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "Туркменистански манат", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Туркменист. манат", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Тунизиски динар", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Тонганска панга", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "Тиморски ескудо", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "Турска лира (1922–2005)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "Турска лира", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Долар на Тринидад и Тобаго", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Тајвански нов долар", Symbol: "TWD"},
				currency.TZS: cldr.Currency{DisplayName: "Танзаниски шилинг", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Украинска хривнија", Symbol: "UAH"},
				currency.UGS: cldr.Currency{DisplayName: "Угандиски шилинг (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "Угандиски шилинг", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "Американски долар", Symbol: "US$"},
				currency.USN: cldr.Currency{DisplayName: "САД долар (Next day)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "САД долар (Same day)", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "Уругвајски пезос (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "Уругвајски пезос", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Узбекистански сом", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "Венецуелски боливар (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "Венецуелски боливар (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Венецуелски боливар", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Виетнамски донг", Symbol: "VND"},
				currency.VUV: cldr.Currency{DisplayName: "Ванатски вату", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Самоанска тала", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Централноафрикански франк", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Источно карипски долар", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Западноафрикански франк", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "ЦФП франк", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Непозната валута", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "Јеменски динар", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "Јеменски риал", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "Југословенски динар", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "Југословенски конвертибилен динар", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "Јужно афрички ранд(финансиски)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Јужноафрикански ранд", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Замбијска квача (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Замбијска квача", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "Заирско новозаире", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "Заирско заире", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "Зимбабвиски долар", Symbol: ""},
			},
		},
	}
}
