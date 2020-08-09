// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_ru() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y 'г'.", Long: "d MMMM y 'г'.", Medium: "d MMM y 'г'.", Short: "dd.MM.y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "янв.", Feb: "февр.", Mar: "март", Apr: "апр.", May: "май", Jun: "июнь", Jul: "июль", Aug: "авг.", Sep: "сент.", Oct: "окт.", Nov: "нояб.", Dec: "дек."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "Я", Feb: "Ф", Mar: "М", Apr: "А", May: "М", Jun: "И", Jul: "И", Aug: "А", Sep: "С", Oct: "О", Nov: "Н", Dec: "Д"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "январь", Feb: "февраль", Mar: "март", Apr: "апрель", May: "май", Jun: "июнь", Jul: "июль", Aug: "август", Sep: "сентябрь", Oct: "октябрь", Nov: "ноябрь", Dec: "декабрь"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "вс", Mon: "пн", Tue: "вт", Wed: "ср", Thu: "чт", Fri: "пт", Sat: "сб"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "В", Mon: "П", Tue: "В", Wed: "С", Thu: "Ч", Fri: "П", Sat: "С"}, Short: cldr.CalendarDayFormatNameValue{Sun: "вс", Mon: "пн", Tue: "вт", Wed: "ср", Thu: "чт", Fri: "пт", Sat: "сб"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "воскресенье", Mon: "понедельник", Tue: "вторник", Wed: "среда", Thu: "четверг", Fri: "пятница", Sat: "суббота"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_ru]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "Андоррская песета", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "дирхам ОАЭ", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "Афгани (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "афгани", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "албанский лек", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "армянский драм", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "нидерландский антильский гульден", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "ангольская кванза", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "Ангольская кванза (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "Ангольская новая кванза (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "Ангольская кванза реюстадо (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "Аргентинский аустрал", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "Аргентинское песо (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "аргентинский песо", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "Австрийский шиллинг", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "австралийский доллар", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "арубанский флорин", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "Старый азербайджанский манат", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "азербайджанский манат", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "Динар Боснии и Герцеговины", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "конвертируемая марка Боснии и Герцеговины", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "барбадосский доллар", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "бангладешская така", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "Бельгийский франк (конвертируемый)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "Бельгийский франк", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "Бельгийский франк (финансовый)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "Лев", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "болгарский лев", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "бахрейнский динар", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "бурундийский франк", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "бермудский доллар", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "брунейский доллар", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "боливийский боливиано", Symbol: "BOB"},
				currency.BOP: cldr.Currency{DisplayName: "Боливийское песо", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "Боливийский мвдол", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "Бразильский новый крузейро (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "Бразильское крузадо", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "Бразильский крузейро (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "бразильский реал", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "Бразильское новое крузадо", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "Бразильский крузейро", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "багамский доллар", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "бутанский нгултрум", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "Джа", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "ботсванская пула", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "Белорусский рубль (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "белорусский рубль", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Белорусский рубль (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "белизский доллар", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "канадский доллар", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "конголезский франк", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "WIR евро", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "швейцарский франк", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "WIR франк", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "Условная расчетная единица Чили", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "чилийский песо", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "китайский офшорный юань", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "китайский юань", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "колумбийский песо", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "Единица реальной стоимости Колумбии", Symbol: ""},
				currency.CRC: cldr.Currency{DisplayName: "костариканский колон", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "Старый Сербский динар", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "Чехословацкая твердая крона", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "кубинский конвертируемый песо", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "кубинский песо", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "эскудо Кабо-Верде", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "Кипрский фунт", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "чешская крона", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "Восточногерманская марка", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "Немецкая марка", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "франк Джибути", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "датская крона", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "доминиканский песо", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "алжирский динар", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "Эквадорский сукре", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "Постоянная единица стоимости Эквадора", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "Эстонская крона", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "египетский фунт", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "эритрейская накфа", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "Испанская песета (А)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "Испанская песета (конвертируемая)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "Испанская песета", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "эфиопский быр", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "евро", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Финская марка", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "доллар Фиджи", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "фунт Фолклендских островов", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "Французский франк", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "британский фунт стерлингов", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "Грузинский купон", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "грузинский лари", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "Ганский седи (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "ганский седи", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "гибралтарский фунт", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "гамбийский даласи", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "гвинейский франк", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "Гвинейская сили", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "Эквеле экваториальной Гвинеи", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "Греческая драхма", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "гватемальский кетсаль", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "Эскудо Португальской Гвинеи", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "Песо Гвинеи-Бисау", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "гайанский доллар", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "гонконгский доллар", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "гондурасская лемпира", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "Хорватский динар", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "хорватская куна", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "гаитянский гурд", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "венгерский форинт", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "индонезийская рупия", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "Ирландский фунт", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "Израильский фунт", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "новый израильский шекель", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "индийская рупия", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "иракский динар", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "иранский риал", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "исландская крона", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "Итальянская лира", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "ямайский доллар", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "иорданский динар", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "японская иена", Symbol: "¥"},
				currency.KES: cldr.Currency{DisplayName: "кенийский шиллинг", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "киргизский сом", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "камбоджийский риель", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "коморский франк", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "северокорейская вона", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "южнокорейская вона", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "кувейтский динар", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "доллар Островов Кайман", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "казахский тенге", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "лаосский кип", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "ливанский фунт", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "шри-ланкийская рупия", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "либерийский доллар", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Лоти", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Литовский лит", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "Литовский талон", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "Конвертируемый франк Люксембурга", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "Люксембургский франк", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "Финансовый франк Люксембурга", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "Латвийский лат", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "Латвийский рубль", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "ливийский динар", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "марокканский дирхам", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "Марокканский франк", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "молдавский лей", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "малагасийский ариари", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "Малагасийский франк", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "македонский денар", Symbol: "MKD"},
				currency.MLF: cldr.Currency{DisplayName: "Малийский франк", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "мьянманский кьят", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "монгольский тугрик", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "патака Макао", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "мавританская угия (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "мавританская угия", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "Мальтийская лира", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "Мальтийский фунт", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "маврикийская рупия", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "мальдивская руфия", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "малавийская квача", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "мексиканский песо", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "Мексиканское серебряное песо (1861–1992)", Symbol: ""},
				currency.MXV: cldr.Currency{DisplayName: "Мексиканская пересчетная единица (UDI)", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "малайзийский ринггит", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "Мозамбикское эскудо", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "Старый мозамбикский метикал", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "мозамбикский метикал", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "доллар Намибии", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "нигерийская найра", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "Никарагуанская кордоба (1988–1991)", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "никарагуанская кордоба", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "Нидерландский гульден", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "норвежская крона", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "непальская рупия", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "новозеландский доллар", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "оманский риал", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "панамский бальбоа", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "Перуанское инти", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "перуанский соль", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "Перуанский соль (1863–1965)", Symbol: "PES"},
				currency.PGK: cldr.Currency{DisplayName: "кина Папуа – Новой Гвинеи", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "филиппинский песо", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "пакистанская рупия", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "польский злотый", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "Злотый", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "Португальское эскудо", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "парагвайский гуарани", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "катарский риал", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "Родезийский доллар", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "Старый Румынский лей", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "румынский лей", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "сербский динар", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "российский рубль", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "Российский рубль (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "франк Руанды", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "саудовский риял", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "доллар Соломоновых Островов", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "сейшельская рупия", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "Суданский динар", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "суданский фунт", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "Старый суданский фунт", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "шведская крона", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "сингапурский доллар", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "фунт острова Святой Елены", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "Словенский толар", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "Словацкая крона", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "леоне", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "сомалийский шиллинг", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "суринамский доллар", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "Суринамский гульден", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "южносуданский фунт", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "добра Сан-Томе и Принсипи (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "добра Сан-Томе и Принсипи", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "Рубль СССР", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "Сальвадорский колон", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "сирийский фунт", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "свазилендский лилангени", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "таиландский бат", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "Таджикский рубль", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "таджикский сомони", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "Туркменский манат", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "новый туркменский манат", Symbol: "ТМТ"},
				currency.TND: cldr.Currency{DisplayName: "тунисский динар", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "тонганская паанга", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "Тиморское эскудо", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "Турецкая лира (1922–2005)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "турецкая лира", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "доллар Тринидада и Тобаго", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "новый тайваньский доллар", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "танзанийский шиллинг", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "украинская гривна", Symbol: "₴"},
				currency.UAK: cldr.Currency{DisplayName: "Карбованец (украинский)", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "Старый угандийский шиллинг", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "угандийский шиллинг", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "доллар США", Symbol: "$"},
				currency.USN: cldr.Currency{DisplayName: "Доллар США следующего дня", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "Доллар США текущего дня", Symbol: ""},
				currency.UYI: cldr.Currency{DisplayName: "Уругвайский песо (индекс инфляции)", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "Уругвайское старое песо (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "уругвайский песо", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "узбекский сум", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "Венесуэльский боливар (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "венесуэльский боливар (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "венесуэльский боливар", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "вьетнамский донг", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "вату Вануату", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "самоанская тала", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "франк КФА BEAC", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "Серебро", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "Золото", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "Европейская составная единица", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "Европейская денежная единица", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "расчетная единица европейского валютного соглашения (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "расчетная единица европейского валютного соглашения (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "восточно-карибский доллар", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "СДР (специальные права заимствования)", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "ЭКЮ (единица европейской валюты)", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "Французский золотой франк", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "Французский UIC-франк", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "франк КФА ВСЕАО", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "Палладий", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "французский тихоокеанский франк", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "Платина", Symbol: ""},
				currency.XRE: cldr.Currency{DisplayName: "единица RINET-фондов", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "тестовый валютный код", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "неизвестная валюта", Symbol: "XXXX"},
				currency.YDD: cldr.Currency{DisplayName: "Йеменский динар", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "йеменский риал", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "Югославский твердый динар", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "Югославский новый динар", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "Югославский динар", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "Южноафриканский рэнд (финансовый)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "южноафриканский рэнд", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Квача (замбийская) (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "замбийская квача", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "Новый заир", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "Заир", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "Доллар Зимбабве", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "Доллар Зимбабве (2009)", Symbol: ""},
			},
		},
	}
}
