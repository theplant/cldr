// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_uk_UA() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "uk_UA",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y 'р'.", Long: "d MMMM y 'р'.", Medium: "d MMM y 'р'.", Short: "dd.MM.yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'о' {0}", Long: "{1} 'о' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "січ", Feb: "лют", Mar: "бер", Apr: "кві", May: "тра", Jun: "чер", Jul: "лип", Aug: "сер", Sep: "вер", Oct: "жов", Nov: "лис", Dec: "гру"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "С", Feb: "Л", Mar: "Б", Apr: "К", May: "Т", Jun: "Ч", Jul: "Л", Aug: "С", Sep: "В", Oct: "Ж", Nov: "Л", Dec: "Г"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "січень", Feb: "лютий", Mar: "березень", Apr: "квітень", May: "травень", Jun: "червень", Jul: "липень", Aug: "серпень", Sep: "вересень", Oct: "жовтень", Nov: "листопад", Dec: "грудень"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "нд", Mon: "пн", Tue: "вт", Wed: "ср", Thu: "чт", Fri: "пт", Sat: "сб"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Н", Mon: "П", Tue: "В", Wed: "С", Thu: "Ч", Fri: "П", Sat: "С"}, Short: cldr.CalendarDayFormatNameValue{Sun: "нд", Mon: "пн", Tue: "вт", Wed: "ср", Thu: "чт", Fri: "пт", Sat: "сб"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "неділя", Mon: "понеділок", Tue: "вівторок", Wed: "середа", Thu: "четвер", Fri: "пʼятниця", Sat: "субота"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "дп", PM: "пп"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "дп", PM: "пп"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "дп", PM: "пп"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_uk]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "андоррська песета", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "дирхам ОАЕ", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "афгані (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "афганський афгані", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "албанський лек", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "вірменський драм", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "нідерландський антильський гульден", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "ангольська кванза", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "ангольська кванза (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "ангольська нова кванза (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "ангольська кванза реаджастадо (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "аргентинський австрал", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "аргентинський песо (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "аргентинський песо", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "австрійський шилінг", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "австралійський долар", Symbol: "AUD"},
				currency.AWG: cldr.Currency{DisplayName: "арубський флорин", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "азербайджанський манат (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "азербайджанський манат", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "динар (Боснія і Герцеговина)", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "конвертована марка Боснії і Герцеговини", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "барбадоський долар", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "бангладеська така", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "бельгійський франк (конвертований)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "бельгійський франк", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "бельгійський франк (фінансовий)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "болгарський твердий лев", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "болгарський лев", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "бахрейнський динар", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "бурундійський франк", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "бермудський долар", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "брунейський долар", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "болівійський болівіано", Symbol: "BOB"},
				currency.BOP: cldr.Currency{DisplayName: "болівійське песо", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "болівійський мвдол", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "бразильське нове крузейро (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "бразильське крузадо", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "бразильське крузейро (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "бразильський реал", Symbol: "BRL"},
				currency.BRN: cldr.Currency{DisplayName: "бразильське нове крузадо", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "бразильське крузейро", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "багамський долар", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "бутанський нгултрум", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "бірманський кіат", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "ботсванська пула", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "білоруський новий рубль (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "білоруський рубль", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "білоруський рубль (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "белізький долар", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "канадський долар", Symbol: "CAD"},
				currency.CDF: cldr.Currency{DisplayName: "конголезький франк", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "євро WIR", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "швейцарський франк", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "франк WIR", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "чилійський юнідадес де фоменто", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "чилійський песо", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "китайський офшорний юань", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "китайський юань", Symbol: "CNY"},
				currency.COP: cldr.Currency{DisplayName: "колумбійський песо", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "одиниця реальної вартості", Symbol: ""},
				currency.CRC: cldr.Currency{DisplayName: "костариканський колон", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "старий сербський динар", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "чехословацька тверда крона", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "кубинський конвертований песо", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "кубинський песо", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "ескудо Кабо-Верде", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "кіпрський фунт", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "чеська крона", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "марка НДР", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "німецька марка", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "джибутійський франк", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "данська крона", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "домініканський песо", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "алжирський динар", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "еквадорський сукре", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "еквадорський юнідад де валор константе", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "естонська крона", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "єгипетський фунт", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "еритрейська накфа", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "іспанська песета (\"А\" рахунок)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "іспанська песета (конвертовані рахунки)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "іспанська песета", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "ефіопський бир", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "євро", Symbol: "EUR"},
				currency.FIM: cldr.Currency{DisplayName: "фінляндська марка", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "фіджійський долар", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "фунт Фолклендських островів", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "французький франк", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "англійський фунт", Symbol: "GBP"},
				currency.GEK: cldr.Currency{DisplayName: "грузинський купон", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "грузинський ларі", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "ганський седі (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "ганський седі", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "гібралтарський фунт", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "гамбійський даласі", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "гвінейський франк", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "гвінейське сілі", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "еквеле (Екваторіальна Ґвінея)", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "грецька драхма", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "гватемальський кетсаль", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "ескудо португальської гвінеї", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "песо Гвінеї-Бісау", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "гаянський долар", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "гонконгський долар", Symbol: "HKD"},
				currency.HNL: cldr.Currency{DisplayName: "гондураська лемпіра", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "хорватський динар", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "хорватська куна", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "гаїтянський гурд", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "угорський форинт", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "індонезійська рупія", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "ірландський фунт", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "ізраїльський фунт", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "ізраїльський новий шекель", Symbol: "ILS"},
				currency.INR: cldr.Currency{DisplayName: "індійська рупія", Symbol: "INR"},
				currency.IQD: cldr.Currency{DisplayName: "іракський динар", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "іранський ріал", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "ісландська крона", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "італійська ліра", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "ямайський долар", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "йорданський динар", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "японська єна", Symbol: "¥"},
				currency.KES: cldr.Currency{DisplayName: "кенійський шилінг", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "киргизький сом", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "камбоджійський рієль", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "коморський франк", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "північнокорейський вон", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "південнокорейський вон", Symbol: "KRW"},
				currency.KWD: cldr.Currency{DisplayName: "кувейтський динар", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "долар Кайманових островів", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "казахстанський тенге", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "лаоський кіп", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "ліванський фунт", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "шрі-ланкійська рупія", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "ліберійський долар", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "лесотський лоті", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "литовський літ", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "литовський талон", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "люксембурґський франк (конвертований)", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "люксембурзький франк", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "люксембурґський франк (фінансовий)", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "латвійський лат", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "латвійський рубль", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "лівійський динар", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "марокканський дирхам", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "марокканський франк", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "молдовський лей", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "малагасійський аріарі", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "мадагаскарський франк", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "македонський денар", Symbol: "MKD"},
				currency.MLF: cldr.Currency{DisplayName: "малійський франк", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "кʼят Мʼянми", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "монгольський тугрик", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "патака Макао", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "мавританська угія (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "мавританська угія", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "мальтійська ліра", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "мальтійський фунт", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "маврикійська рупія", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "мальдівська руфія", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "малавійська квача", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "мексиканський песо", Symbol: "MXN"},
				currency.MXP: cldr.Currency{DisplayName: "мексиканське срібне песо (1861–1992)", Symbol: ""},
				currency.MXV: cldr.Currency{DisplayName: "мексиканський юнідад де інверсіон", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "малайзійський рингіт", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "мозамбіцький ескудо", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "старий мозамбіцький метикал", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "мозамбіцький метикал", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "намібійський долар", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "нігерійська найра", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "нікарагуанська кордоба (1988–1991)", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "нікарагуанська кордоба", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "нідерландський гульден", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "норвезька крона", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "непальська рупія", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "новозеландський долар", Symbol: "NZD"},
				currency.OMR: cldr.Currency{DisplayName: "оманський ріал", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "панамське бальбоа", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "перуанський інті", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "перуанський новий сол", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "перуанський сол (1863–1965)", Symbol: "PES"},
				currency.PGK: cldr.Currency{DisplayName: "кіна Папуа-Нової Ґвінеї", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "філіппінський песо", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "пакистанська рупія", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "польський злотий", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "польський злотий (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "португальський ескудо", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "парагвайський гуарані", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "катарський ріал", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "родезійський долар", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "старий румунський лей", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "румунський лей", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "сербський динар", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "російський рубль", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "російський рубль (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "руандійський франк", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "саудівський ріал", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "долар Соломонових Островів", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "сейшельська рупія", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "суданський динар", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "суданський фунт", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "старий суданський фунт", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "шведська крона", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "сінгапурський долар", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "фунт острова Святої Єлени", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "словенський толар", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "словацька крона", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "леоне Сьєрра-Леоне", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "сомалійський шилінг", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "суринамський долар", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "суринамський гульден", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "південносуданський фунт", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "добра Сан-Томе і Прінсіпі (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "добра Сан-Томе і Прінсіпі", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "радянський рубль", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "сальвадорський колон", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "сирійський фунт", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "свазілендський лілангені", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "таїландський бат", Symbol: "THB"},
				currency.TJR: cldr.Currency{DisplayName: "таджицький рубль", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "таджицький сомоні", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "туркменський манат (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "туркменський манат", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "туніський динар", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "тонґанська паанга", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "тіморський ескудо", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "турецька ліра (1922–2005)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "турецька ліра", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "долар Трінідаду і Тобаґо", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "новий тайванський долар", Symbol: "TWD"},
				currency.TZS: cldr.Currency{DisplayName: "танзанійський шилінг", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "українська гривня", Symbol: "₴"},
				currency.UAK: cldr.Currency{DisplayName: "український карбованець", Symbol: "крб."},
				currency.UGS: cldr.Currency{DisplayName: "угандійський шилінг (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "угандійський шилінг", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "долар США", Symbol: "USD"},
				currency.USN: cldr.Currency{DisplayName: "долар США (наступного дня)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "долар США (цього дня)", Symbol: ""},
				currency.UYI: cldr.Currency{DisplayName: "уругвайський песо в індексованих одиницях", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "уругвайське песо (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "уругвайський песо", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "узбецький сум", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "венесуельський болівар (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "венесуельський болівар (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "венесуельський болівар", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "вʼєтнамський донг", Symbol: "VND"},
				currency.VUV: cldr.Currency{DisplayName: "вануатський вату", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "самоанська тала", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "центральноафриканський франк", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "срібло", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "золото", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "європейська складена валютна одиниця", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "одиниця європейського валютного фонду", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "європейська розрахункова одиниця XBC", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "європейська розрахункова одиниця XBD", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "східнокарибський долар", Symbol: "XCD"},
				currency.XDR: cldr.Currency{DisplayName: "спеціальні права запозичення", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "європейська валютна одиниця", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "французький золотий франк", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "французький франк UIC", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "західноафриканський франк", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "паладій", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "французький тихоокеанський франк", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "платина", Symbol: ""},
				currency.XRE: cldr.Currency{DisplayName: "фонди RINET", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "код тестування валюти", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "невідома грошова одиниця", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "єменський динар", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "єменський ріал", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "югославський твердий динар", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "югославський новий динар", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "югославський конвертований динар", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "південноафриканський фінансовий ранд", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "південноафриканський ранд", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "замбійська квача (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "замбійська квача", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "заїрський новий заїр", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "заїрський заїр", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "зімбабвійський долар", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "зімбабвійський долар (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "зімбабвійський долар (2008)", Symbol: ""},
			},
		},
	}
}
