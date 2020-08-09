// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_sr_Cyrl_RS() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, dd. MMMM y.", Long: "dd. MMMM y.", Medium: "dd.MM.y.", Short: "d.M.yy."}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "јан", Feb: "феб", Mar: "мар", Apr: "апр", May: "мај", Jun: "јун", Jul: "јул", Aug: "авг", Sep: "сеп", Oct: "окт", Nov: "нов", Dec: "дец"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "ј", Feb: "ф", Mar: "м", Apr: "а", May: "м", Jun: "ј", Jul: "ј", Aug: "а", Sep: "с", Oct: "о", Nov: "н", Dec: "д"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "јануар", Feb: "фебруар", Mar: "март", Apr: "април", May: "мај", Jun: "јун", Jul: "јул", Aug: "август", Sep: "септембар", Oct: "октобар", Nov: "новембар", Dec: "децембар"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "нед", Mon: "пон", Tue: "уто", Wed: "сре", Thu: "чет", Fri: "пет", Sat: "суб"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "н", Mon: "п", Tue: "у", Wed: "с", Thu: "ч", Fri: "п", Sat: "с"}, Short: cldr.CalendarDayFormatNameValue{Sun: "не", Mon: "по", Tue: "ут", Wed: "ср", Thu: "че", Fri: "пе", Sat: "су"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "недеља", Mon: "понедељак", Tue: "уторак", Wed: "среда", Thu: "четвртак", Fri: "петак", Sat: "субота"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "пре подне", PM: "по подне"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "пре подне", PM: "по подне"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "пре подне", PM: "по подне"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_sr]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "Андорска пезета", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "УАЕ дирхам", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "Авганистански авгани (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Авганистански авгани", Symbol: "AFN"},
				currency.ALK: cldr.Currency{DisplayName: "стари албански лек", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "Албански лек", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Јерменски драм", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Холандскоантилски гулден", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Анголска кванза", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "Анголијска кванза (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "Анголијска нова кванза (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "Анголијска кванза реађустадо (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "Аргентински аустрал", Symbol: ""},
				currency.ARL: cldr.Currency{DisplayName: "Аргентински пезос леј", Symbol: ""},
				currency.ARM: cldr.Currency{DisplayName: "Аргентински пезос монедо национал", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "Аргентински пезо (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "Аргентински пезос", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "Аустријски шилинг", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "Аустралијски долар", Symbol: "AUD"},
				currency.AWG: cldr.Currency{DisplayName: "Арубански флорин", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "Азербејџански манат (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Азербејџански манат", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "Босанско-Херцеговачки динар", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Босанско-херцеговачка конвертибилна марка", Symbol: "КМ"},
				currency.BAN: cldr.Currency{DisplayName: "Босанско-херцеговачки нови динар", Symbol: ""},
				currency.BBD: cldr.Currency{DisplayName: "Барбадошки долар", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Бангладешка така", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "Белгијски франак (конвертибилни)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "Белгијски франак", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "Белгијски франак (финансијски)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "Бугарски тврди лев", Symbol: ""},
				currency.BGM: cldr.Currency{DisplayName: "Бугарски социјалистички лев", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "Бугарски лев", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "Стари бугарски лев", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Бахреински динар", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Бурундски франак", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Бермудски долар", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Брунејски долар", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Боливијски боливијано", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "Стари боливијски боливијано", Symbol: ""},
				currency.BOP: cldr.Currency{DisplayName: "Боливијски пезо", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "Боливијски мвдол", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "Бразилски нови крузеиро (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "Бразилијски крузадо", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "Бразилски крузеиро (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "Бразилски реал", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "Бразилијски нови крузадо", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "Бразилски крузеиро", Symbol: ""},
				currency.BRZ: cldr.Currency{DisplayName: "Стари бразилски крузеиро", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "Бахамски долар", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Бутански нгултрум", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "Бурмански кјат", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Боцванска пула", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "Белоруска нова рубља (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "Белоруска рубља", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Белоруска рубља (2000–2016)", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "Белиски долар", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Канадски долар", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Конгоански франак", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "WIR евро", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Швајцарски франак", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "WIR франак", Symbol: ""},
				currency.CLE: cldr.Currency{DisplayName: "Чилеански ескудо", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "Чилеовски унидадес се фоменто", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Чилеански пезос", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Кинески јуан (острвски)", Symbol: "CNH"},
				currency.CNX: cldr.Currency{DisplayName: "Долар кинеске народне банке", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "Кинески јуан", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Колумбијски пезос", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "Унидад де валоршки реал", Symbol: ""},
				currency.CRC: cldr.Currency{DisplayName: "Костарикански колон", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "Стари српски динар", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "Чехословачка тврда круна", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "Кубански конвертибилни пезос", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Кубански пезос", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Зеленортски ескудо", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "Кипарска фунта", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Чешка круна", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "Источно-немачка марка", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "Немачка марка", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "Џибутски франак", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Данска круна", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Доминикански пезос", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Алжирски динар", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "Еквадорски сакр", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "Еквадорски унидад де валор константе", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "Естонска кроон", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Египатска фунта", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Eритрејска накфa", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "Шпанска пезета (рачун)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "Шпанска пезета (конвертибилнирачун)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "Шпанска пезета", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Етиопски бир", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Евро", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Финска марка", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "Фиџијски долар", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Фокландска фунта", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "Француски франак", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "Британска фунта", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "Грузијски купон ларит", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "Грузијски лари", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "Гански цеди (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Гански седи", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Гибралтарска фунта", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Гамбијски даласи", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Гвинејски франак", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "Гвинејски сили", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "Екваторијално-гвинејски еквеле", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "Грчка драхма", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "Гватемалски кецал", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "Португалска гвинеја ескудо", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "Гвинеја Бисао Пезо", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "Гвајански долар", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Хонгконшки долар", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Хондурашка лемпира", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "Хрватски динар", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "Хрватска куна", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Хаићански гурд", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Мађарска форинта", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Индонежанска рупија", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "Ирска фунта", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "Израелска фунта", Symbol: ""},
				currency.ILR: cldr.Currency{DisplayName: "Стари израелски шекели", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "Израелски нови шекел", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Индијска рупија", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Ирачки динар", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Ирански риjал", Symbol: "IRR"},
				currency.ISJ: cldr.Currency{DisplayName: "Стара исландска круна", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Исландска круна", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "Италијанска лира", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "Јамајчански долар", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Јордански динар", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Јапански јен", Symbol: "¥"},
				currency.KES: cldr.Currency{DisplayName: "Кенијски шилинг", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Киргистански сом", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Kамбоџански ријел", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Коморски франак", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Севернокорејски вон", Symbol: "KPW"},
				currency.KRH: cldr.Currency{DisplayName: "Јужнокорејски хван", Symbol: ""},
				currency.KRO: cldr.Currency{DisplayName: "Стари јужнокорејски вон", Symbol: ""},
				currency.KRW: cldr.Currency{DisplayName: "Јужнокорејски вон", Symbol: "KRW"},
				currency.KWD: cldr.Currency{DisplayName: "Кувајтски динар", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Кајмански долар", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Казахстански тенге", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Лаошки кип", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Либанска фунта", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Шриланканскa рупиja", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Либеријски долар", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Лесото лоти", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Литвански литас", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "Литвански талонас", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "Луксембуршки конвертибилни франак", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "Луксембуршки франак", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "Луксембуршки финансијски франак", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "Латвијски лати", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "Латвијска рубља", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "Либијски динар", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Марокански дирхам", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "Марокански франак", Symbol: ""},
				currency.MCF: cldr.Currency{DisplayName: "Монегаскански франак", Symbol: ""},
				currency.MDC: cldr.Currency{DisplayName: "Молдовански купон", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Молдавски леј", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Мадагаскарски ариари", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "Малагасијски франак", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "Македонски денар", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "Стари македонски денар", Symbol: ""},
				currency.MLF: cldr.Currency{DisplayName: "Малијански франак", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Мјанмарски кјат", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Монголски тугрик", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Макаоска патака", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Мауританијска oгија (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Мауританска огија", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "Малтешка лира", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "Малтешка фунта", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Маурицијска рупија", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Малдивска руфија", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Малавијска квача", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Мексички пезос", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "Мексички сребрни пезо (1861–1992)", Symbol: ""},
				currency.MXV: cldr.Currency{DisplayName: "Мексички унидад де инверсион (UDI)", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "Малезијски рингит", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "Мозамбијски ескудо", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "Стари мозамбијски метикал", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Мозамбички метикал", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Намибијски долар", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Нигеријска наира", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "Никарагванска кордоба", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "Никарагванска златна кордоба", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "Холандски гулден", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "Норвешка круна", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Непалскa рупиja", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Новозеландски долар", Symbol: "NZD"},
				currency.OMR: cldr.Currency{DisplayName: "Омански ријал", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Панамска балбоа", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "Перуански инти", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Перуански сол", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "Перуански сол (1863–1965)", Symbol: "PES"},
				currency.PGK: cldr.Currency{DisplayName: "Папуанска кина", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Филипински пезос", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Пакистанскa рупиja", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Пољски злот", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "Пољски злоти (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "Португалски ескудо", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "Парагвајски гварани", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Катарски ријал", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "Родејскидолар", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "Румунски леј (1952–2006)", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Румунски леј", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Српски динар", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Руска рубља", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "Руска рубља (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Руандски франак", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Саудијски ријал", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Соломонски долар", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Сејшелска рупија", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "Стари судански динар", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Суданска фунта", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "Стара суданска фунта", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Шведска круна", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Сингапурски долар", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Свете Јелене фунта", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "Словеначки толар", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "Словачка круна", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Сијералеонски леоне", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Сомалијски шилинг", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Суринамски долар", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "Суринамски гилдер", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "Јужносуданска фунта", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "Саотомска добра (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "Саотомска добра", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "Совјетска рубља", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "Салвадорски колон", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "Сиријска фунта", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Свазилендски лилангени", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Тајландски бат", Symbol: "THB"},
				currency.TJR: cldr.Currency{DisplayName: "Таџихистанска рубља", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "Таџикистански сомон", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "Туркменистански манат (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Туркменистански манат", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Туниски динар", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Тонганска панга", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "Тиморшки ескудо", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "Турска лира (1922–2005)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "Турска лира", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Тринидад-тобагошки долар", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Нови тајвански долар", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Танзанијски шилинг", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Украјинска хривња", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "Украјински карбованети", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "Угандски шилинг (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "Угандски шилинг", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "Амерички долар", Symbol: "US$"},
				currency.USN: cldr.Currency{DisplayName: "САД долар (следећи дан)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "САД долар (исти дан)", Symbol: ""},
				currency.UYI: cldr.Currency{DisplayName: "Уругвајски пезо ен унидадес индексадас", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "Уругвајски пезо (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "Уругвајски пезос", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Узбекистански сом", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "Венецуелански боливар (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "Венецуелански боливар (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Венецуелански боливар", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Вијетнамски донг", Symbol: "VND"},
				currency.VNN: cldr.Currency{DisplayName: "Вијетнамски донг (1978–1985)", Symbol: ""},
				currency.VUV: cldr.Currency{DisplayName: "Вануатски вату", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Самоанска тала", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "CFA франак BEAC", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "Сребро", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "Злато", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "Европска композитна јединица", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "Европска новчана јединица", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "Европска јединица рачуна (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "Европска јединица рачуна (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "Источнокарипски долар", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "Посебна цртаћа права", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "Европска валутна јединица", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "Француски златни франак", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "Француски UIC-франак", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "ЦФА франак БЦЕАО", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "Паладијум", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "CFP франак", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "Платина", Symbol: ""},
				currency.XRE: cldr.Currency{DisplayName: "RINET фонд", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "Код тестиране валуте", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "Непозната валута", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "Јеменски динар", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "Јеменски риjал", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "Југословенски тврди динар", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "Југословенски нови динар", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "Југословенски конвертибилни динар", Symbol: ""},
				currency.YUR: cldr.Currency{DisplayName: "Југословенски реформирани динар", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "Јужно-афрички ранд (финансијски)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Јужноафрички ранд", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Замбијска квача (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Замбијска квача", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "Заирски нови заир", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "Заирски заир", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "Зимбабвеански долар (1980–2008)", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "Зимбабвеански долар (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "Зимбабвеански долар (2008)", Symbol: ""},
			},
		},
	}
}
