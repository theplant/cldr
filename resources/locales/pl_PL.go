// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_pl_PL() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd.MM.y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "sty", Feb: "lut", Mar: "mar", Apr: "kwi", May: "maj", Jun: "cze", Jul: "lip", Aug: "sie", Sep: "wrz", Oct: "paź", Nov: "lis", Dec: "gru"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "S", Feb: "L", Mar: "M", Apr: "K", May: "M", Jun: "C", Jul: "L", Aug: "S", Sep: "W", Oct: "P", Nov: "L", Dec: "G"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "styczeń", Feb: "luty", Mar: "marzec", Apr: "kwiecień", May: "maj", Jun: "czerwiec", Jul: "lipiec", Aug: "sierpień", Sep: "wrzesień", Oct: "październik", Nov: "listopad", Dec: "grudzień"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "niedz.", Mon: "pon.", Tue: "wt.", Wed: "śr.", Thu: "czw.", Fri: "pt.", Sat: "sob."}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "N", Mon: "P", Tue: "W", Wed: "Ś", Thu: "C", Fri: "P", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "nie", Mon: "pon", Tue: "wto", Wed: "śro", Thu: "czw", Fri: "pią", Sat: "sob"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "niedziela", Mon: "poniedziałek", Tue: "wtorek", Wed: "środa", Thu: "czwartek", Fri: "piątek", Sat: "sobota"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "a", PM: "p"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_pl]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "peseta andorska", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "dirham ZEA", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "afgani (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "afgani afgańskie", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "lek albański", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "dram armeński", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "gulden antylski", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "kwanza angolska", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "kwanza angolańska (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "nowa kwanza angolańska (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "kwanza angolańska Reajustado (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "austral argentyński", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "peso argentyńskie (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "peso argentyńskie", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "szyling austriacki", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "dolar australijski", Symbol: "AUD"},
				currency.AWG: cldr.Currency{DisplayName: "florin arubański", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "manat azerbejdżański", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "manat azerski", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "dinar Bośni i Hercegowiny", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "marka zamienna Bośni i Hercegowiny", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "dolar Barbadosu", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "taka bengalska", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "frank belgijski (zamienny)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "frank belgijski", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "frank belgijski (finansowy)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "lew bułgarski wymienny", Symbol: ""},
				currency.BGM: cldr.Currency{DisplayName: "lew bułgarski socjalistyczny", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "lew bułgarski", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "lew bułgarski (1879–1952)", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "dinar bahrański", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "frank burundyjski", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "dolar bermudzki", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "dolar brunejski", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "boliviano boliwijskie", Symbol: "BOB"},
				currency.BOP: cldr.Currency{DisplayName: "peso boliwijskie", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "mvdol boliwijski", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "cruzeiro novo brazylijskie (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "cruzado brazylijskie", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "cruzeiro brazylijskie (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "real brazylijski", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "nowe cruzado brazylijskie", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "cruzeiro brazylijskie", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "dolar bahamski", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "ngultrum bhutański", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "kyat birmański", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "pula botswańska", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "rubel białoruski (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "rubel białoruski", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "rubel białoruski (2000–2016)", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "dolar belizeński", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "dolar kanadyjski", Symbol: "CAD"},
				currency.CDF: cldr.Currency{DisplayName: "frank kongijski", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "frank szwajcarski", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "peso chilijskie", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "juan chiński (rynek zewnętrzny)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "juan chiński", Symbol: "CNY"},
				currency.COP: cldr.Currency{DisplayName: "peso kolumbijskie", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "colon kostarykański", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "stary dinar serbski", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "korona czechosłowacka", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "peso kubańskie wymienialne", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "peso kubańskie", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "escudo zielonoprzylądkowe", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "funt cypryjski", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "korona czeska", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "wschodnia marka wschodnioniemiecka", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "marka niemiecka", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "frank dżibutyjski", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "korona duńska", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "peso dominikańskie", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "dinar algierski", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "sucre ekwadorski", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "korona estońska", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "funt egipski", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "nakfa erytrejska", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "peseta hiszpańska (Konto A)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "peseta hiszpańska (konto wymienne)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "peseta hiszpańska", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "birr etiopski", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "marka fińska", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "dolar fidżyjski", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "funt falklandzki", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "frank francuski", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "funt szterling", Symbol: "GBP"},
				currency.GEK: cldr.Currency{DisplayName: "kupon gruziński larit", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "lari gruzińskie", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "cedi ghańskie (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "cedi ghańskie", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "funt gibraltarski", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "dalasi gambijskie", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "frank gwinejski", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "syli gwinejskie", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "ekwele gwinejskie Gwinei Równikowej", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "drachma grecka", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "quetzal gwatemalski", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "escudo Gwinea Portugalska", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "peso Guinea-Bissau", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "dolar gujański", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "dolar hongkoński", Symbol: "HKD"},
				currency.HNL: cldr.Currency{DisplayName: "lempira honduraska", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "dinar chorwacki", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "kuna chorwacka", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "gourde haitański", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "forint węgierski", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "rupia indonezyjska", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "funt irlandzki", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "funt izraelski", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "nowy szekel izraelski", Symbol: "ILS"},
				currency.INR: cldr.Currency{DisplayName: "rupia indyjska", Symbol: "INR"},
				currency.IQD: cldr.Currency{DisplayName: "dinar iracki", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "rial irański", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "korona islandzka", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "lir włoski", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "dolar jamajski", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "dinar jordański", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "jen japoński", Symbol: "JPY"},
				currency.KES: cldr.Currency{DisplayName: "szyling kenijski", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "som kirgiski", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "riel kambodżański", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "frank komoryjski", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "won północnokoreański", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "won południowokoreański", Symbol: "KRW"},
				currency.KWD: cldr.Currency{DisplayName: "dinar kuwejcki", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "dolar kajmański", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "tenge kazachskie", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "kip laotański", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "funt libański", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "rupia lankijska", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "dolar liberyjski", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "loti Lesoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "lit litewski", Symbol: "Lt"},
				currency.LTT: cldr.Currency{DisplayName: "talon litewski", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "frank luksemburski", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "łat łotewski", Symbol: "Ls"},
				currency.LVR: cldr.Currency{DisplayName: "rubel łotewski", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "dinar libijski", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "dirham marokański", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "frank marokański", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "lej mołdawski", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "ariary malgaski", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "frank malgaski", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "denar macedoński", Symbol: "MKD"},
				currency.MLF: cldr.Currency{DisplayName: "frank malijski", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "kiat birmański", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "tugrik mongolski", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "pataca Makau", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "ouguiya mauretańska (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "ugija mauretańska", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "lira maltańska", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "funt maltański", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "rupia maurytyjska", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "rupia malediwska", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "kwacha malawijska", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "peso meksykańskie", Symbol: "MXN"},
				currency.MXP: cldr.Currency{DisplayName: "peso srebrne meksykańskie (1861–1992)", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "ringgit malezyjski", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "escudo mozambickie", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "metical Mozambik", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "metical mozambicki", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "dolar namibijski", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "naira nigeryjska", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "cordoba nikaraguańska (1988–1991)", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "cordoba nikaraguańska", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "gulden holenderski", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "korona norweska", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "rupia nepalska", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "dolar nowozelandzki", Symbol: "NZD"},
				currency.OMR: cldr.Currency{DisplayName: "rial omański", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "balboa panamski", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "inti peruwiański", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "sol peruwiański", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "sol peruwiański (1863–1965)", Symbol: "PES"},
				currency.PGK: cldr.Currency{DisplayName: "kina papuańska", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "peso filipińskie", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "rupia pakistańska", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "złoty polski", Symbol: "zł"},
				currency.PLZ: cldr.Currency{DisplayName: "złoty polski (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "escudo portugalskie", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "guarani paragwajskie", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "rial katarski", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "dolar rodezyjski", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "lej rumuński (1952–2006)", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "lej rumuński", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "dinar serbski", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "rubel rosyjski", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "rubel rosyjski (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "frank ruandyjski", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "rial saudyjski", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "dolar Wysp Salomona", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "rupia seszelska", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "dinar sudański", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "funt sudański", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "funt sudański (1957–1998)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "korona szwedzka", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "dolar singapurski", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "funt Świętej Heleny", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "tolar słoweński", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "korona słowacka", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "leone sierraleoński", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "szyling somalijski", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "dolar surinamski", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "gulden surinamski", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "funt południowosudański", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "dobra Wysp Świętego Tomasza i Książęcej (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "dobra Wysp Świętego Tomasza i Książęcej", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "rubel radziecki", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "colon salwadorski", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "funt syryjski", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "lilangeni Suazi", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "baht tajski", Symbol: "THB"},
				currency.TJR: cldr.Currency{DisplayName: "rubel tadżycki", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "somoni tadżyckie", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "manat turkmeński (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "manat turkmeński", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "dinar tunezyjski", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "pa’anga tongijska", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "escudo timorskie", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "lira turecka (1922–2005)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "lira turecka", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "dolar Trynidadu i Tobago", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "nowy dolar tajwański", Symbol: "TWD"},
				currency.TZS: cldr.Currency{DisplayName: "szyling tanzański", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "hrywna ukraińska", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "karbowaniec ukraiński", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "szyling ugandyjski (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "szyling ugandyjski", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "dolar amerykański", Symbol: "USD"},
				currency.UYP: cldr.Currency{DisplayName: "peso urugwajskie (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "peso urugwajskie", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "som uzbecki", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "boliwar wenezuelski (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "boliwar wenezuelski (2008–2018)", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "boliwar wenezuelski", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "dong wietnamski", Symbol: "VND"},
				currency.VUV: cldr.Currency{DisplayName: "vatu wanuackie", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "tala samoańskie", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "frank CFA BEAC", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "srebro", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "złoto", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "jednostka emisji euroobligacji", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "europejska jednostka monetarna", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "europejska jednostka rozrachunkowa (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "europejska jednostka rozrachunkowa (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "dolar wschodniokaraibski", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "specjalne prawa ciągnienia", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "ECU", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "frank złoty francuski", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "UIC-frank francuski", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "frank CFA", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "pallad", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "frank CFP", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "platyna", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "testowy kod waluty", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "nieznana waluta", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "dinar jemeński", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "rial jemeński", Symbol: "YER"},
				currency.YUM: cldr.Currency{DisplayName: "nowy dinar jugosławiański", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "dinar jugosławiański wymienny", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "rand południowoafrykański (finansowy)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "rand południowoafrykański", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "kwacha zambijska (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "kwacha zambijska", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "nowy zair zairski", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "zair zairski", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "dolar Zimbabwe (1980–2008)", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "dolar Zimbabwe (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "dolar Zimbabwe (2008)", Symbol: ""},
			},
		},
	}
}
