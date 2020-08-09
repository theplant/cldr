// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_bs_Cyrl() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, dd. MMMM y.", Long: "dd. MMMM y.", Medium: "dd.MM.y.", Short: "d.M.yy."}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "јан", Feb: "феб", Mar: "мар", Apr: "апр", May: "мај", Jun: "јун", Jul: "јул", Aug: "ауг", Sep: "сеп", Oct: "окт", Nov: "нов", Dec: "дец"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "ј", Feb: "ф", Mar: "м", Apr: "а", May: "м", Jun: "ј", Jul: "ј", Aug: "а", Sep: "с", Oct: "о", Nov: "н", Dec: "д"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "јануар", Feb: "фебруар", Mar: "март", Apr: "април", May: "мај", Jun: "јуни", Jul: "јули", Aug: "аугуст", Sep: "септембар", Oct: "октобар", Nov: "новембар", Dec: "децембар"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "нед", Mon: "пон", Tue: "уто", Wed: "сри", Thu: "чет", Fri: "пет", Sat: "суб"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "н", Mon: "п", Tue: "у", Wed: "с", Thu: "ч", Fri: "п", Sat: "с"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "недјеља", Mon: "понедјељак", Tue: "уторак", Wed: "сриједа", Thu: "четвртак", Fri: "петак", Sat: "субота"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "прије подне", PM: "послије подне"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_root]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "Андорска пезета", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "Уједињени арапски емирати дирхам", Symbol: ""},
				currency.AFA: cldr.Currency{DisplayName: "Авганистански авган (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Афганистански афгани", Symbol: ""},
				currency.ALK: cldr.Currency{DisplayName: "стари албански лек", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "Албански лек", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "Арменски драм", Symbol: ""},
				currency.ANG: cldr.Currency{DisplayName: "Холандски антили гилдер", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Анголска кванза", Symbol: "Kz"},
				currency.AOK: cldr.Currency{DisplayName: "анголијска кванза (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "Анголијска нова кванза (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "Анголска кванза реађустадо (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "Аргентински аустрал", Symbol: ""},
				currency.ARL: cldr.Currency{DisplayName: "аргентински пезос леј", Symbol: ""},
				currency.ARM: cldr.Currency{DisplayName: "аргентински пезос монедо национал", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "аргентински пезо (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "Аргентински пезос", Symbol: "$"},
				currency.ATS: cldr.Currency{DisplayName: "Аустријски шилинг", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "Аустралијски долар", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Арубански флорин", Symbol: ""},
				currency.AZM: cldr.Currency{DisplayName: "Азербејџански манат (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Азербејџански манат", Symbol: ""},
				currency.BAD: cldr.Currency{DisplayName: "Босанско-Херцеговачки динар", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Конвертибилна марка", Symbol: "КМ"},
				currency.BAN: cldr.Currency{DisplayName: "босанско-херцеговачки нови динар", Symbol: ""},
				currency.BBD: cldr.Currency{DisplayName: "Барбадоски долар", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "Бангладешка така", Symbol: "৳"},
				currency.BEC: cldr.Currency{DisplayName: "Белгијски франак (конвертибилни)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "Белгијски франак", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "Белгијски франак (финансијски)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "Бугарски тврди лев", Symbol: ""},
				currency.BGM: cldr.Currency{DisplayName: "бугарски социјалистички лев", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "Бугарски лев", Symbol: ""},
				currency.BGO: cldr.Currency{DisplayName: "стари бугарски лев", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Бахреински динар", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Бурундски франак", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "Бермудски долар", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "Брунејски долар", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "Боливијски боливиано", Symbol: "Bs"},
				currency.BOL: cldr.Currency{DisplayName: "стари боливијски боливијано", Symbol: ""},
				currency.BOP: cldr.Currency{DisplayName: "Боливијски пезо", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "Боливијски мвдол", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "Бразилски нови крузеиро (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "Бразилијски крузадо", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "Бразилски крузеиро (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "Бразилски реал", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "Бразилијски нови крузадо", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "Бразилски крузеиро", Symbol: ""},
				currency.BRZ: cldr.Currency{DisplayName: "стари бразилски крузеиро", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "Бахамски долар", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "Бутански нгултрум", Symbol: ""},
				currency.BUK: cldr.Currency{DisplayName: "Бурмански кјат", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Боцванска пула", Symbol: "P"},
				currency.BYB: cldr.Currency{DisplayName: "Белоруска нова рубља (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "Белоруска рубља", Symbol: "р."},
				currency.BYR: cldr.Currency{DisplayName: "Белоруска рубља (2000–2016)", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "Белизеански долар", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Канадски долар", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Конгоански франак", Symbol: ""},
				currency.CHE: cldr.Currency{DisplayName: "WIR евро", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Швајцарски франак", Symbol: ""},
				currency.CHW: cldr.Currency{DisplayName: "WIR франак", Symbol: ""},
				currency.CLE: cldr.Currency{DisplayName: "чилеански ескудо", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "Чилеовски унидадес се фоменто", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Чилеански пезос", Symbol: "$"},
				currency.CNH: cldr.Currency{DisplayName: "Кинески јуан (острвски)", Symbol: ""},
				currency.CNX: cldr.Currency{DisplayName: "долар кинеске народне банке", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "Кинески јуан", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Колумбијски пезос", Symbol: "$"},
				currency.COU: cldr.Currency{DisplayName: "Унидад де валоршки реал", Symbol: ""},
				currency.CRC: cldr.Currency{DisplayName: "Костарикански колон", Symbol: "₡"},
				currency.CSD: cldr.Currency{DisplayName: "Стари српски динар", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "Чехословачка тврда круна", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "кубански конвертибилни пезос", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "Кубански пезос", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Зеленортски ескудо", Symbol: ""},
				currency.CYP: cldr.Currency{DisplayName: "Кипарска фунта", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Чешка круна", Symbol: "Кч"},
				currency.DDM: cldr.Currency{DisplayName: "Источно-немачка марка", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "Немачка марка", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "Џибутски франак", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "Данска круна", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "Доминикански пезос", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Алжирски динар", Symbol: ""},
				currency.ECS: cldr.Currency{DisplayName: "Еквадорски сакр", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "Еквадорски унидад де валор константе", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "Естонска кроон", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Египатска фунта", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Еритрејска накфа", Symbol: ""},
				currency.ESA: cldr.Currency{DisplayName: "Шпанска пезета (рачун)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "Шпанска пезета (конвертибилнирачун)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "Шпанска пезета", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Етиопијски бир", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Евро", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Финска марка", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "Фиџи долар", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "Фолкландска фунта", Symbol: "£"},
				currency.FRF: cldr.Currency{DisplayName: "Француски франак", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "Британска фунта", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "Грузијски купон ларит", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "Грузијски лари", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Гански цеди (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Гански цеди", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "Гибралтаска фунта", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Гамбијски даласи", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "Гвинејски франак", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Гвинејски сили", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "Екваторијално-гвинејски еквеле", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "Грчка драхма", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "Гватемалски квецал", Symbol: "Q"},
				currency.GWE: cldr.Currency{DisplayName: "португалска гвинеја ескудо", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "Гвинеја Бисао Пезо", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "Гвајански долар", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "Хонгконшки долар", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Хондурашка лемпира", Symbol: "L"},
				currency.HRD: cldr.Currency{DisplayName: "Хрватски динар", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "Хрватска куна", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "Хаићански гурд", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "Мађарска форинта", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "Индонежанска рупија", Symbol: "Rp"},
				currency.IEP: cldr.Currency{DisplayName: "Ирска фунта", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "Израелска фунта", Symbol: ""},
				currency.ILR: cldr.Currency{DisplayName: "стари израелски шекели", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "Израелски нови шекел", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Индијска рупија", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Ирачки динар", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "Ирански ријал", Symbol: ""},
				currency.ISJ: cldr.Currency{DisplayName: "стара исландска круна", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Исландска круна", Symbol: "kr"},
				currency.ITL: cldr.Currency{DisplayName: "Италијанска лира", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "Јамајски долар", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "Јордански динар", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "Јапански јен", Symbol: "¥"},
				currency.KES: cldr.Currency{DisplayName: "Кенијски шилинг", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "Киргистански сом", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "Камбоџански ријел", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Коморски франак", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "Севернокорејски вон", Symbol: "₩"},
				currency.KRH: cldr.Currency{DisplayName: "јужнокорејски хван", Symbol: ""},
				currency.KRO: cldr.Currency{DisplayName: "стари јужнокорејски вон", Symbol: ""},
				currency.KRW: cldr.Currency{DisplayName: "Јужнокорејски вон", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Кувајтски динар", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "Кајмански долар", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "Казахстански тенге", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "Лаоски кип", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "Либанска фунта", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "Шриланканска рупија", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Либеријски долар", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Лесото лоти", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Литвански литас", Symbol: "Lt"},
				currency.LTT: cldr.Currency{DisplayName: "Литвански талонас", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "Луксембуршки конвертибилни франак", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "Луксембуршки франак", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "Луксембуршки финансијски франак", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "Латвијски лати", Symbol: "Ls"},
				currency.LVR: cldr.Currency{DisplayName: "атвијска рубља", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "Либијски динар", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Марокански дирхам", Symbol: ""},
				currency.MAF: cldr.Currency{DisplayName: "Марокански франак", Symbol: ""},
				currency.MCF: cldr.Currency{DisplayName: "монегаскански франак", Symbol: ""},
				currency.MDC: cldr.Currency{DisplayName: "молдовански купон", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Молдавски леј", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Мадагаскарски аријари", Symbol: "Ar"},
				currency.MGF: cldr.Currency{DisplayName: "Малагасијски франак", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "Македонски денар", Symbol: ""},
				currency.MKN: cldr.Currency{DisplayName: "стари македонски денар", Symbol: ""},
				currency.MLF: cldr.Currency{DisplayName: "Малијански франак", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Мијанмарски кјат", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "Монголски тугрик", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "Макаоска патака", Symbol: ""},
				currency.MRO: cldr.Currency{DisplayName: "Мауританијска угвија (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Мауританска огија", Symbol: ""},
				currency.MTL: cldr.Currency{DisplayName: "Малтешка лира", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "Малтешка фунта", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Маурицијска рупија", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "Малдивска руфија", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "Малавијска квача", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "Мексички пезос", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "Мексички сребрни пезо (1861–1992)", Symbol: ""},
				currency.MXV: cldr.Currency{DisplayName: "Мексички унидад де инверсион (UDI)", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "Малезијски рингит", Symbol: "RM"},
				currency.MZE: cldr.Currency{DisplayName: "Мозамбијски ескудо", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "Стари мозамбијски метикал", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Мозамбијски метикал", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Намибијски долар", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Нигеријска наира", Symbol: "₦"},
				currency.NIC: cldr.Currency{DisplayName: "Никарагванска кордоба", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "Никарагванска златна кордоба", Symbol: "C$"},
				currency.NLG: cldr.Currency{DisplayName: "Холандски гулден", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "Норвешка круна", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "Непалска рупија", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "Новозеландски долар", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Омански ријал", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "Панамска балбоа", Symbol: ""},
				currency.PEI: cldr.Currency{DisplayName: "Перуански инти", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Перуански сол", Symbol: ""},
				currency.PES: cldr.Currency{DisplayName: "Перуански сол (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Папуанска кина", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "Филипински пезос", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "Пакистанска рупија", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "Пољски злот", Symbol: "зл"},
				currency.PLZ: cldr.Currency{DisplayName: "Пољски злоти (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "Португалски ескудо", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "Парагвајски гварани", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "Катарски ријал", Symbol: ""},
				currency.RHD: cldr.Currency{DisplayName: "Родејскидолар", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "Стари румунски љу", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Румунски леј", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "Српски динар", Symbol: "дин."},
				currency.RUB: cldr.Currency{DisplayName: "Руска рубља", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "Руска рубља (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Руандски франак", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Саудијски ријал", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "Соломонски долар", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Сејшелска рупија", Symbol: ""},
				currency.SDD: cldr.Currency{DisplayName: "Стари судански динар", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Суданска фунта", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Стара суданска фунта", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Шведска круна", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "Сингапурски долар", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Света Јелена фунта", Symbol: "£"},
				currency.SIT: cldr.Currency{DisplayName: "Словеначки толар", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "Словачка круна", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Сијералеонски леоне", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Сомалијски шилинг", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "Суринамски долар", Symbol: "$"},
				currency.SRG: cldr.Currency{DisplayName: "Суринамски гилдер", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "Јужносуданска фунта", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Сао Томе и Принципе добра (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Сао Томе и Принципе добра", Symbol: "Db"},
				currency.SUR: cldr.Currency{DisplayName: "Совјетска рубља", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "Салвадорски колон", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "Сиријска фунта", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Свазилендски лилангени", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "Тајски бахт", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "Таџихистанска рубља", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "Таџикистански сомони", Symbol: ""},
				currency.TMM: cldr.Currency{DisplayName: "Туркменистански манат (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Туркменистански манат", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "Туниски динар", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "Тонгоанска панга", Symbol: "T$"},
				currency.TPE: cldr.Currency{DisplayName: "Тиморшки ескудо", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "Стара турска лира", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "Турска лира", Symbol: "Тл"},
				currency.TTD: cldr.Currency{DisplayName: "Тринидад-тобагошки долар", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "Нови тајвански долар", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Танзанијски шилинг", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "Украјинска хривња", Symbol: "₴"},
				currency.UAK: cldr.Currency{DisplayName: "Украјински карбованети", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "Угандски шилинг (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "Угандски шилинг", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Амерички долар", Symbol: "US$"},
				currency.USN: cldr.Currency{DisplayName: "САД долар (следећи дан)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "САД долар (исти дан)", Symbol: ""},
				currency.UYI: cldr.Currency{DisplayName: "Уругвајски пезо ен унидадес индексадас", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "Уругвајски пезо (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "Уругвајски пезос", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "Узбекистански сом", Symbol: ""},
				currency.VEB: cldr.Currency{DisplayName: "Венецуелански боливар (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "Венецуелански боливар (2008–2018)", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "Венецуелански боливар", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "Вијетнамски донг", Symbol: "₫"},
				currency.VNN: cldr.Currency{DisplayName: "стари вијетнамски донг", Symbol: ""},
				currency.VUV: cldr.Currency{DisplayName: "Вануатски вату", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "Самоанска тала", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "CFA франак BEAC", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "Сребро", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "Злато", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "Европска композитна јединица", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "Европска новчана јединица", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "Европска јединица рачуна (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "Европска јединица рачуна (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "Источно-карибски долар", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "Посебна цртаћа права", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "Европска валутна јединица", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "Француски златни франак", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "Француски UIC-франак", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "CFA франак BCEAO", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "Паладијум", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "CFP франак", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "Платина", Symbol: ""},
				currency.XRE: cldr.Currency{DisplayName: "RINET фонд", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "Код тестиране валуте", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "Непозната или неважећа валута", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "Јеменски динар", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "Јеменски ријал", Symbol: ""},
				currency.YUD: cldr.Currency{DisplayName: "Југословенски тврди динар", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "Југословенски нови динар", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "Југословенски конвертибилни динар", Symbol: ""},
				currency.YUR: cldr.Currency{DisplayName: "југословенски реформирани динар", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "Јужно-афрички ранд (финансијски)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Јужноафрички ранд", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Замбијска квача (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Замбијска квача", Symbol: "ZK"},
				currency.ZRN: cldr.Currency{DisplayName: "Заирски нови заир", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "Заирски заир", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "Зимбабвејски долар", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "Зимбабвеански долар (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "Зимбабвеански долар (2008)", Symbol: ""},
			},
		},
	}
}
