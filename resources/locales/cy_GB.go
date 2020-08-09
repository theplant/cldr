// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_cy_GB() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'am' {0}", Long: "{1} 'am' {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Ion", Feb: "Chw", Mar: "Maw", Apr: "Ebr", May: "Mai", Jun: "Meh", Jul: "Gor", Aug: "Awst", Sep: "Medi", Oct: "Hyd", Nov: "Tach", Dec: "Rhag"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "I", Feb: "Ch", Mar: "M", Apr: "E", May: "M", Jun: "M", Jul: "G", Aug: "A", Sep: "M", Oct: "H", Nov: "T", Dec: "Rh"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Ionawr", Feb: "Chwefror", Mar: "Mawrth", Apr: "Ebrill", May: "Mai", Jun: "Mehefin", Jul: "Gorffennaf", Aug: "Awst", Sep: "Medi", Oct: "Hydref", Nov: "Tachwedd", Dec: "Rhagfyr"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Sul", Mon: "Llun", Tue: "Maw", Wed: "Mer", Thu: "Iau", Fri: "Gwe", Sat: "Sad"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "Ll", Tue: "M", Wed: "M", Thu: "I", Fri: "G", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Su", Mon: "Ll", Tue: "Ma", Wed: "Me", Thu: "Ia", Fri: "Gw", Sat: "Sa"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Dydd Sul", Mon: "Dydd Llun", Tue: "Dydd Mawrth", Wed: "Dydd Mercher", Thu: "Dydd Iau", Fri: "Dydd Gwener", Sat: "Dydd Sadwrn"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_cy]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "", Group: "", Negative: "\u061c-", Percent: "٪\u061c", PerMille: "؉"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Dirham Yr Emiradau Arabaidd Unedig", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "Afghani Afghanistan (1927–2002)", Symbol: "AFA"},
				currency.AFN: cldr.Currency{DisplayName: "Afghani Afghanistan", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "Lek Albania", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Dram Armenia", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Guilder Antilles yr Iseldiroedd", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Kwanza Angola", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "Kwanza Angola (1977–1991)", Symbol: "AOK"},
				currency.AON: cldr.Currency{DisplayName: "Kwanza Newydd Angola (1990–2000)", Symbol: "AON"},
				currency.AOR: cldr.Currency{DisplayName: "Kwanza Ailgymhwysedig Angola (1995–1999)", Symbol: "AOR"},
				currency.ARA: cldr.Currency{DisplayName: "Austral yr Ariannin", Symbol: "ARA"},
				currency.ARL: cldr.Currency{DisplayName: "Peso Ley yr Ariannin (1970–1983)", Symbol: "ARL"},
				currency.ARM: cldr.Currency{DisplayName: "Peso yr Ariannin (1881–1970)", Symbol: "ARM"},
				currency.ARP: cldr.Currency{DisplayName: "Peso yr Ariannin (1983–1985)", Symbol: "ARP"},
				currency.ARS: cldr.Currency{DisplayName: "Peso yr Ariannin", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "Swllt Awstria", Symbol: "ATS"},
				currency.AUD: cldr.Currency{DisplayName: "Doler Awstralia", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Fflorin Aruba", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "Manat Azerbaijan (1993–2006)", Symbol: "AZM"},
				currency.AZN: cldr.Currency{DisplayName: "Manat Azerbaijan", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "Marc Trosadwy Bosnia a Hercegovina", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "Doler Barbados", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Taka Bangladesh", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "Ffranc Gwlad Belg (arnewidiol)", Symbol: "BEC"},
				currency.BEF: cldr.Currency{DisplayName: "Ffranc Gwlad Belg", Symbol: "BEF"},
				currency.BEL: cldr.Currency{DisplayName: "Ffranc Gwlad Belg (ariannol)", Symbol: "BEL"},
				currency.BGM: cldr.Currency{DisplayName: "Lev Sosialaidd Bwlgaria", Symbol: "BGM"},
				currency.BGN: cldr.Currency{DisplayName: "Lev Bwlgaria", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "Lev Bwlgaria (1879–1952)", Symbol: "BGO"},
				currency.BHD: cldr.Currency{DisplayName: "Dinar Bahrain", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Ffranc Burundi", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Doler Bermuda", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Doler Brunei", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Boliviano Bolifia", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "Boliviano Bolifia (1863–1963)", Symbol: "BOL"},
				currency.BOP: cldr.Currency{DisplayName: "Peso Bolifia", Symbol: "BOP"},
				currency.BOV: cldr.Currency{DisplayName: "Mvdol Bolifia", Symbol: "BOV"},
				currency.BRB: cldr.Currency{DisplayName: "Cruzeiro Newydd Brasil (1967–1986)", Symbol: "BRB"},
				currency.BRC: cldr.Currency{DisplayName: "Cruzado Brasil (1986–1989)", Symbol: "BRC"},
				currency.BRE: cldr.Currency{DisplayName: "Cruzeiro Brasil (1990–1993)", Symbol: "BRE"},
				currency.BRL: cldr.Currency{DisplayName: "Real Brasil", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "Cruzado Newydd Brasil (1989–1990)", Symbol: "BRN"},
				currency.BRR: cldr.Currency{DisplayName: "Cruzeiro Brasil (1993–1994)", Symbol: "BRR"},
				currency.BRZ: cldr.Currency{DisplayName: "Cruzeiro Brasil (1942–1967)", Symbol: "BRZ"},
				currency.BSD: cldr.Currency{DisplayName: "Doler y Bahamas", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Ngultrum Bhutan", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "Kyat Byrma", Symbol: "BUK"},
				currency.BWP: cldr.Currency{DisplayName: "Pula Botswana", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "Rwbl Belarws", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Rwbl Belarws (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "Doler Belize", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Doler Canada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Ffranc Congo", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "Ewro WIR", Symbol: "CHE"},
				currency.CHF: cldr.Currency{DisplayName: "Ffranc y Swistir", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "Ffranc WIR", Symbol: "CHW"},
				currency.CLE: cldr.Currency{DisplayName: "Escudo Chile", Symbol: "CLE"},
				currency.CLF: cldr.Currency{DisplayName: "Uned Cyfrifo Chile (UF)", Symbol: "CLF"},
				currency.CLP: cldr.Currency{DisplayName: "Peso Chile", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Yuan Tsieina (ar y môr)", Symbol: "CNH"},
				currency.CNX: cldr.Currency{DisplayName: "Doler Banc Pobl Tsieina", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "Yuan Tsieina", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Peso Colombia", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "Uned Gwir Werth Colombia", Symbol: "COU"},
				currency.CRC: cldr.Currency{DisplayName: "Colón Costa Rica", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "", Symbol: "CSD"},
				currency.CUC: cldr.Currency{DisplayName: "Peso Trosadwy Ciwba", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Peso Ciwba", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Esgwdo Cabo Verde", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "Punt Cyprus", Symbol: "CYP"},
				currency.CZK: cldr.Currency{DisplayName: "Koruna’r Weriniaeth Tsiec", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "Marc Dwyrain yr Almaen", Symbol: "DDM"},
				currency.DEM: cldr.Currency{DisplayName: "Marc yr Almaen", Symbol: "DEM"},
				currency.DJF: cldr.Currency{DisplayName: "Ffranc Djibouti", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Krone Denmarc", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Peso Gweriniaeth Dominica", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Dinar Algeria", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "Sucre Ecuador", Symbol: "ECS"},
				currency.ECV: cldr.Currency{DisplayName: "Uned Gwerth Gyson Ecuador", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "Kroon Estonia", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Punt Yr Aifft", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfa Eritrea", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "ESP"},
				currency.ETB: cldr.Currency{DisplayName: "Birr Ethiopia", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Ewro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Markka’r Ffindir", Symbol: "FIM"},
				currency.FJD: cldr.Currency{DisplayName: "Doler Ffiji", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Punt Ynysoedd Falkland/Malvinas", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "Ffranc Ffrainc", Symbol: "FRF"},
				currency.GBP: cldr.Currency{DisplayName: "Punt Prydain", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "Kupon Larit Georgia", Symbol: "GEK"},
				currency.GEL: cldr.Currency{DisplayName: "Lari Georgia", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "Cedi Ghana (1979–2007)", Symbol: "GHC"},
				currency.GHS: cldr.Currency{DisplayName: "Cedi Ghana", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Punt Gibraltar", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi Gambia", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Ffranc Guinée", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "Syli Guinée", Symbol: "GNS"},
				currency.GQE: cldr.Currency{DisplayName: "Ekwele Guinea Gyhydeddol", Symbol: "GQE"},
				currency.GRD: cldr.Currency{DisplayName: "", Symbol: "GRD"},
				currency.GTQ: cldr.Currency{DisplayName: "Quetzal Guatemala", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "", Symbol: "GWE"},
				currency.GWP: cldr.Currency{DisplayName: "Peso Guiné-Bissau", Symbol: "GWP"},
				currency.GYD: cldr.Currency{DisplayName: "Doler Guyana", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Doler Hong Kong", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Lempira Honduras", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "", Symbol: "HRD"},
				currency.HRK: cldr.Currency{DisplayName: "Kuna Croatia", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Gourde Haiti", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Fforint Hwngari", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Rupiah Indonesia", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "Punt Iwerddon", Symbol: "IEP"},
				currency.ILP: cldr.Currency{DisplayName: "Punt Israel", Symbol: "ILP"},
				currency.ILR: cldr.Currency{DisplayName: "Shegel Israel (1980–1985)", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "Shegel Newydd Israel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rwpî India", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Dinar Irac", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Rial Iran", Symbol: "IRR"},
				currency.ISJ: cldr.Currency{DisplayName: "Króna Gwlad yr Iâ (1918 – 1981)", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Króna Gwlad yr Iâ", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "", Symbol: "ITL"},
				currency.JMD: cldr.Currency{DisplayName: "Doler Jamaica", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Dinar Gwlad yr Iorddonen", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Yen Japan", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Swllt Kenya", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Som Kyrgyzstan", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Riel Cambodia", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Ffranc Comoros", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Won Gogledd Corea", Symbol: "KPW"},
				currency.KRH: cldr.Currency{DisplayName: "Hwan De Corea (1953–1962)", Symbol: "KRH"},
				currency.KRO: cldr.Currency{DisplayName: "Won De Corea (1945–1953)", Symbol: "KRO"},
				currency.KRW: cldr.Currency{DisplayName: "Won De Corea", Symbol: "KRW"},
				currency.KWD: cldr.Currency{DisplayName: "Dinar Kuwait", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Doler Ynysoedd Cayman", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Tenge Kazakstan", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Kip Laos", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Punt Libanus", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Rwpî Sri Lanka", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Doler Liberia", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Loti Lesotho", Symbol: "LSL"},
				currency.LTL: cldr.Currency{DisplayName: "Litas Lithwania", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "Talonas Lithwania", Symbol: "LTT"},
				currency.LUC: cldr.Currency{DisplayName: "", Symbol: "LUC"},
				currency.LUF: cldr.Currency{DisplayName: "Ffranc Lwcsembwrg", Symbol: "LUF"},
				currency.LUL: cldr.Currency{DisplayName: "", Symbol: "LUL"},
				currency.LVL: cldr.Currency{DisplayName: "Lats Latfia", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "Rwbl Latfia", Symbol: "LVR"},
				currency.LYD: cldr.Currency{DisplayName: "Dinar Libya", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Dirham Moroco", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "Ffranc Moroco", Symbol: "MAF"},
				currency.MCF: cldr.Currency{DisplayName: "Ffranc Monaco", Symbol: "MCF"},
				currency.MDL: cldr.Currency{DisplayName: "Leu Moldofa", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Ariary Madagascar", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "Ffranc Madagascar", Symbol: "MGF"},
				currency.MKD: cldr.Currency{DisplayName: "Denar Macedonia", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "", Symbol: "MKN"},
				currency.MLF: cldr.Currency{DisplayName: "Ffranc Mali", Symbol: "MLF"},
				currency.MMK: cldr.Currency{DisplayName: "Kyat Myanmar", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Tugrik Mongolia", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Pataca Macau", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Ouguiya Mauritania (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Ouguiya Mauritania", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "", Symbol: "MTL"},
				currency.MTP: cldr.Currency{DisplayName: "", Symbol: "MTP"},
				currency.MUR: cldr.Currency{DisplayName: "Rwpî Mauritius", Symbol: "MUR"},
				currency.MVP: cldr.Currency{DisplayName: "Rwpî’r Maldives (1947–1981)", Symbol: ""},
				currency.MVR: cldr.Currency{DisplayName: "Rufiyaa’r Maldives", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Kwacha Malawi", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Peso Mecsico", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "Peso Arian México (1861–1992)", Symbol: "MXP"},
				currency.MXV: cldr.Currency{DisplayName: "Uned Fuddsoddi México", Symbol: "MXV"},
				currency.MYR: cldr.Currency{DisplayName: "Ringgit Malaysia", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "Escudo Mozambique", Symbol: "MZE"},
				currency.MZM: cldr.Currency{DisplayName: "Metical Mozambique (1980–2006)", Symbol: "MZM"},
				currency.MZN: cldr.Currency{DisplayName: "Metical Mozambique", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Doler Namibia", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Naira Nigeria", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "Córdoba Nicaragua (1988–1991)", Symbol: "NIC"},
				currency.NIO: cldr.Currency{DisplayName: "Cordoba Nicaragwa", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "Guilder yr Iseldiroedd", Symbol: "NLG"},
				currency.NOK: cldr.Currency{DisplayName: "Krone Norwy", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Rwpî Nepal", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Doler Seland Newydd", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Rial Oman", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Balboa Panama", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "Inti Periw", Symbol: "PEI"},
				currency.PEN: cldr.Currency{DisplayName: "Sol Periw", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "Sol Periw (1863–1965)", Symbol: "PES"},
				currency.PGK: cldr.Currency{DisplayName: "Kina Papua Guinea Newydd", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Peso Philipinas", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Rwpî Pacistan", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Zloty Gwlad Pwyl", Symbol: "PLN"},
				currency.PTE: cldr.Currency{DisplayName: "", Symbol: "PTE"},
				currency.PYG: cldr.Currency{DisplayName: "Guarani Paraguay", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Rial Qatar", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "Doler Rhodesia", Symbol: "RHD"},
				currency.RON: cldr.Currency{DisplayName: "Leu Rwmania", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Dinar Serbia", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Rwbl Rwsia", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Ffranc Rwanda", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Riyal Saudi Arabia", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Doler Ynysoedd Solomon", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Rwpî Seychelles", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "Dinar Sudan (1992–2007)", Symbol: "SDD"},
				currency.SDG: cldr.Currency{DisplayName: "Punt Sudan", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "Punt Sudan (1957–1998)", Symbol: "SDP"},
				currency.SEK: cldr.Currency{DisplayName: "Krona Sweden", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Doler Singapore", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Punt St Helena", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "", Symbol: "SIT"},
				currency.SLL: cldr.Currency{DisplayName: "Leone Sierra Leone", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Swllt Somalia", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Doler Surinam", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "Guilder Surinam", Symbol: "SRG"},
				currency.SSP: cldr.Currency{DisplayName: "Punt De Sudan", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "Dobra São Tomé a Príncipe (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "Dobra São Tomé a Príncipe", Symbol: "STN"},
				currency.SVC: cldr.Currency{DisplayName: "Colón El Salvador", Symbol: "SVC"},
				currency.SYP: cldr.Currency{DisplayName: "Punt Syria", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni Gwlad Swazi", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Baht Gwlad Thai", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "Rwbl Tajikistan", Symbol: "TJR"},
				currency.TJS: cldr.Currency{DisplayName: "Somoni Tajikistan", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "Manat Turkmenistan (1993–2009)", Symbol: "TMM"},
				currency.TMT: cldr.Currency{DisplayName: "Manat Turkmenistan", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Dinar Tunisia", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Paʻanga Tonga", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "Escudo Timor", Symbol: "TPE"},
				currency.TRL: cldr.Currency{DisplayName: "Lira Twrci (1922–2005)", Symbol: "TRL"},
				currency.TRY: cldr.Currency{DisplayName: "Lira Twrci", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Doler Trinidad a Tobago", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Doler Newydd Taiwan", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Swllt Tanzania", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Hryvnia Wcráin", Symbol: "UAH"},
				currency.UGS: cldr.Currency{DisplayName: "Swllt Uganda (1966–1987)", Symbol: "UGS"},
				currency.UGX: cldr.Currency{DisplayName: "Swllt Uganda", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "Doler UDA", Symbol: "US$"},
				currency.USN: cldr.Currency{DisplayName: "Doler UDA (y diwrnod nesaf)", Symbol: "USN"},
				currency.USS: cldr.Currency{DisplayName: "Doler UDA (yr un diwrnod)", Symbol: "USS"},
				currency.UYI: cldr.Currency{DisplayName: "", Symbol: "UYI"},
				currency.UYP: cldr.Currency{DisplayName: "Peso Uruguay (1975–1993)", Symbol: "UYP"},
				currency.UYU: cldr.Currency{DisplayName: "Peso Uruguay", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Som Uzbekistan", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "Bolívar Venezuela (1871–2008)", Symbol: "VEB"},
				currency.VEF: cldr.Currency{DisplayName: "Bolívar Venezuela (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Bolívar Venezuela", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Dong Fietnam", Symbol: "₫"},
				currency.VNN: cldr.Currency{DisplayName: "Dong Fietnam (1978–1985)", Symbol: "VNN"},
				currency.VUV: cldr.Currency{DisplayName: "Vatu Vanuatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Tala Samoa", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Ffranc CFA Canol Affrica", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "Arian", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "Aur", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "Uned Cyfansawdd Ewropeaidd", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "Uned Ariannol Ewropeaidd", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "Doler Dwyrain y Caribî", Symbol: "EC$"},
				currency.XEU: cldr.Currency{DisplayName: "Uned Arian Cyfred Ewropeaidd", Symbol: "XEU"},
				currency.XOF: cldr.Currency{DisplayName: "Ffranc CFA Gorllewin Affrica", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "Paladiwm", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "Ffranc CFP", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "Platinwm", Symbol: ""},
				currency.XSU: cldr.Currency{DisplayName: "Sucre", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "Arian Cyfred Anhysbys", Symbol: "XXX"},
				currency.YDD: cldr.Currency{DisplayName: "Dinar Yemen", Symbol: "YDD"},
				currency.YER: cldr.Currency{DisplayName: "Rial Yemen", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "", Symbol: "YUD"},
				currency.YUM: cldr.Currency{DisplayName: "", Symbol: "YUM"},
				currency.YUN: cldr.Currency{DisplayName: "", Symbol: "YUN"},
				currency.YUR: cldr.Currency{DisplayName: "", Symbol: "YUR"},
				currency.ZAL: cldr.Currency{DisplayName: "Rand (ariannol) De Affrica", Symbol: "ZAL"},
				currency.ZAR: cldr.Currency{DisplayName: "Rand De Affrica", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwacha Zambia (1968–2012)", Symbol: "ZMK"},
				currency.ZMW: cldr.Currency{DisplayName: "Kwacha Zambia", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "Zaire Newydd Zaire (1993–1998)", Symbol: "ZRN"},
				currency.ZRZ: cldr.Currency{DisplayName: "Zaire Zaire (1971–1993)", Symbol: "ZRZ"},
				currency.ZWD: cldr.Currency{DisplayName: "Doler Zimbabwe (1980–2008)", Symbol: "ZWD"},
				currency.ZWL: cldr.Currency{DisplayName: "Doler Zimbabwe (2009)", Symbol: "ZWL"},
				currency.ZWR: cldr.Currency{DisplayName: "Doler Zimbabwe (2008)", Symbol: "ZWR"},
			},
		},
	}
}
