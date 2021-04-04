// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_pcm() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "pcm",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "H:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'fọ' {0}", Long: "{1} 'fọ' {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jén", Feb: "Fẹ́b", Mar: "Mach", Apr: "Épr", May: "Mee", Jun: "Jun", Jul: "Jul", Aug: "Ọ́gọ", Sep: "Sẹp", Oct: "Ọkt", Nov: "Nọv", Dec: "Dis"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Jénúári", Feb: "Fẹ́búári", Mar: "Mach", Apr: "Éprel", May: "Mee", Jun: "Jun", Jul: "Julai", Aug: "Ọgọst", Sep: "Sẹptẹ́mba", Oct: "Ọktóba", Nov: "Nọvẹ́mba", Dec: "Disẹ́mba"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Sọ́n", Mon: "Mọ́n", Tue: "Tiú", Wed: "Wẹ́n", Thu: "Tọ́z", Fri: "Fraí", Sat: "Sát"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "W", Thu: "T", Fri: "F", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Sọ́n", Mon: "Mọ́n", Tue: "Tiú", Wed: "Wẹ́n", Thu: "Tọ́z", Fri: "Fraí", Sat: "Sát"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Sọ́ndè", Mon: "Mọ́ndè", Tue: "Tiúzdè", Wed: "Wẹ́nẹ́zdè", Thu: "Tọ́zdè", Fri: "Fraídè", Sat: "Sátọdè"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "Fọ mọ́nin", PM: "Fọ ívnin"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Yunaítẹ́d Áráb Ẹ́míréts Dírham", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "Afgán Afgáni", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "Albéniá Lẹk", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Armẹ́niá Dram", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Nẹ́dalánds Antílián Gílda", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Angólá Kwánza", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "Ajẹntína Pẹ́so", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "Ọstréliá Dọ́la", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Arúba Flọ́rin", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "Azẹrbaiján Mánat", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "Bọ́sniá an Hẹzẹgovína Mak Wé Pẹ́sin Fít Chenj", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "Barbédọs Dọ́la", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Bangladẹ́sh Táka", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "Bọlgériá Lẹv", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "Baréin Dínar", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Burúndí Frank", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Bẹmiúda Dọ́la", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Brunẹí Dọ́la", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Bolíviá Boliviáno", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "Brazíl Rẹal", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Bahámas Dọ́la", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Bután Ngúltrum", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "Botswáná Púla", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "Bẹlarús Rúbul", Symbol: "BYN"},
				currency.BZD: cldr.Currency{DisplayName: "Bẹliz Dọ́la", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Kánádá Dọ́la", Symbol: "KA$"},
				currency.CDF: cldr.Currency{DisplayName: "Kóngó Frank", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "Swís Frank", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "Chílí Pẹ́so", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Chaíná Yuan (ples-dẹm aúsaíd chaína)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "Chaíná Yuan", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Kolómbiá Pẹ́so", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "Kósta Ríka Kólọn", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "Kiúbá Pẹ́so Wé Pẹ́sin Fít Chenj", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Kiúbá Pẹ́so", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Kép Vẹ́d Ẹskúdo", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "Chẹ́k Kórúna", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "Jibútí Frank", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Dẹ́nmák Króna", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Dọmíníkan Pẹ́so", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Aljíria Dínar", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "Íjípt Paund", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Ẹritrẹá Nákfa", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Ẹtiópiá Berr", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Yúro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Fíjí Dọ́la", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Fọlkland Aílands Paund", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "Brítísh Páund", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Jọ́jiá Lári", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "Ganá Sídi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Jibrọ́lta Páund", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Gámbiá Dalási", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Gíní Frank", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "Guátẹmála Kwuẹ́tzal", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "Giyána Dọ́la", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Họng Kọ́ng Dọ́la", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Họndúrán Lẹmpíra", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "Kroéshia Kúna", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Haíti Gourd", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Họngériá Fọ́rint", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Indoníshiá Rupia", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "Ízrẹ́l Niú Shẹ́kẹl", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Índiá Rúpi", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Irák Dínar", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Irán Rial", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "Aíslánd Króna", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "Jamaíka Dọla", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Jọ́dán Dínar", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Japán Yẹn", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Kẹ́nyá Shílin", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Kẹjístan Som", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Kambódiá Riẹl", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Kọ́mọ́ros Frank", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Nọ́t Koriá Wọn", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "Saút Koriá Wọn", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Kuwét Dínar", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Kéman Aílands Dọla", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Kazakstan Tẹ́nj", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Laós Kip", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Lẹ́bánọ́n Paund", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Sri Lánká Rúpi", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Laibẹ́riá Dọ́la", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Líbia Dínar", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Morọko Dírham", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "Mọldóva Lu", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Malagásí Ariári", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "Masẹdónia Dínar", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "Miánmá Kiat", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Mọngóliá Túgrik", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Makáo Pátáka", Symbol: "MOP"},
				currency.MRU: cldr.Currency{DisplayName: "Mọriténiá Uguíya", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "Mọríshọ́s Rúpi", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Mọ́ldívs Rúfíya", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Maláwi ́Kwácha", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Mẹ́ksíko Pẹ́so", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Maléshiá Ríngit", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "Mozámbík Métíkal", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Namíbiá Dọ́la", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Naijíriá Naíra", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "Nikarágwua Kordóba", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "Nọ́wé Króna", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Nẹ́pál Rúpi", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Niú Zílánd Dọ́las", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Omán Rial", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Pánáma Balbóa", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "Pẹrúvián Sol", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "Pápuá Niú Gíni Kína", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Fílípíns Píso", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Pakístán Rúpi", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Pólánd Zílọ́ti", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "Páragwuá Guaráni", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Kata Ríal", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "Roméniá Lu", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Sẹrbia Dínar", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Rọ́shiá Rúbul", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Ruwándá Frank", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Saúdí Arébiá Riyal", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Sólómọ́n Aílands Dọ́la", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Sẹ́chẹ́ls Rúpi", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "Sudan Paund", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "Swídẹ́n Króna", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Singapọ́ Dọ́la", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Sent Hẹlẹ́ná Paund", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "Siẹ́ra Líoniá Liọn", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Sọmáliá Shílin", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Súrínám Dọla", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "Saút Sudán Paund", Symbol: "SSP"},
				currency.STN: cldr.Currency{DisplayName: "Sao Tómẹ & Prínsípẹ Dóbra", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "Síriá Paund", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Swází Lilánjẹ́ni", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Taílánd Baht", Symbol: "THB"},
				currency.TJS: cldr.Currency{DisplayName: "Tajíkstan Sómóni", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "Tọkmẹ́nístán Mánat", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Tuníshia Dínar", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Tónga Pánga", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "Tọ́kí Líra", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Trínídad & Tobágo Dọ́la", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Niú Taiwán Dọ́la", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Tanzániá Shílin", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Yukrẹín Rívnia", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "Yugándá Shílin", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "US Dọ́la", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "Yurugwaí Pẹ́so", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Uzbẹ́kistan Som", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "Vẹnẹzuẹlá Bolívar", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Viẹ́tnám Dọng", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Vanuátú Vátu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Samóa Tála", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Sẹ́ntrál Áfríká Frank", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Íst Karíbián Dọla", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Wẹ́st Afríká Sẹ́fa Frank", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "Frẹ́nch Poliníshiá Frank", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Mọní Wé Pípol Nọ No", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Yẹ́mẹ́n Rial", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "Saút Áfríká Rand", Symbol: "ZAR"},
				currency.ZMW: cldr.Currency{DisplayName: "Zámbiá Kwácha", Symbol: "ZMW"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AF:      "Áfríkaans Lángwej",
			language.AGQ:     "Aghẹ́m Lángwej",
			language.AK:      "Akan Lángwej",
			language.AM:      "Amhárík Lángwej",
			language.AR:      "Arábík Lángwej",
			language.AR_001:  "Gẹ́nárál Arábík Lángwej fọ dís taim",
			language.AS:      "Asamíz Lángwej",
			language.ASA:     "Asu Lángwej",
			language.AST:     "Astúriá Lángwej",
			language.AZ:      "Azẹrbaijáni Lángwej",
			language.BAS:     "Básaa Lángwej",
			language.BE:      "Bẹlarúsiá Lángwej",
			language.BEM:     "Bẹ́mba Lángwej",
			language.BEZ:     "Bẹ́na Lángwej",
			language.BG:      "Bọlgériá Lángwej",
			language.BM:      "Bambára Lángwej",
			language.BN:      "Bángla Lángwej",
			language.BO:      "Tibẹ́tan",
			language.BR:      "Brẹ́tọn Lángwej",
			language.BRX:     "Bódo Lángwej",
			language.BS:      "Bọ́sniá Lángwej",
			language.CA:      "Kátála Lángwej",
			language.CCP:     "Chákma Lángwej",
			language.CE:      "Chẹ́chẹn Lángwej",
			language.CEB:     "Sẹbuáno Lángwej",
			language.CGG:     "Chíga Lángwej",
			language.CHR:     "Chẹ́rókii Lángwej",
			language.CKB:     "Mídúl Kọ́dish Lángwej",
			language.CO:      "Kọsíkan Lángwej",
			language.CS:      "Chẹ́k Lángwej",
			language.CU:      "Chọ́ch Slávik",
			language.CY:      "Wẹlsh",
			language.DA:      "Dénísh Lángwej",
			language.DAV:     "Taíta",
			language.DE:      "Jámán Lángwej",
			language.DE_AT:   "Ọ́stria Jámán",
			language.DE_CH:   "Swítzaland Haí Jámán",
			language.DJE:     "Zármá",
			language.DSB:     "Lówá Sorbiá",
			language.DUA:     "Duála Lángwej",
			language.DYO:     "Jóla-Fónyi Lángwej",
			language.DZ:      "Zọ́ngka Lángwej",
			language.EBU:     "Ẹmbu Lángwej",
			language.EE:      "Ẹ́wẹ́ Lángwej",
			language.EL:      "Grík Lángwej",
			language.EN:      "Ínglish",
			language.EN_AU:   "Ọstréliá Ínglish",
			language.EN_CA:   "Kánáda Ínglish",
			language.EN_GB:   "UK Ínglish",
			language.EN_US:   "US Ínglish",
			language.EO:      "Ẹsperánto Lángwej",
			language.ES:      "Spánish Lángwej",
			language.ES_419:  "Látín Amẹ́ríka Spánish",
			language.ES_ES:   "Yúrop Spánish",
			language.ES_MX:   "Mẹ́ksiko Spánish",
			language.ET:      "Ẹstóniá Lángwej",
			language.EU:      "Básk Lángwej",
			language.EWO:     "Ẹwondo Lángwej",
			language.FA:      "Pẹ́shiá Lángwej",
			language.FA_AF:   "Dári",
			language.FF:      "Fúlaní Lángwej",
			language.FI:      "Fínísh Lángwej",
			language.FIL:     "Filipínó Lángwej",
			language.FO:      "Fáróís Lángwej",
			language.FR:      "Frẹ́nch Lángwej",
			language.FR_CA:   "Kánádá Frẹnch",
			language.FR_CH:   "Swítzalánd Frẹnch",
			language.FUR:     "Friúlián Lángwej",
			language.FY:      "Wẹ́stán Frísiá Lángwej",
			language.GA:      "Aírísh Lángwej",
			language.GD:      "Gaelík Lángwej ọf Gael Pípol fọ Skọ́tland",
			language.GL:      "Galísiá Lángwej",
			language.GSW:     "Jámán Swis",
			language.GU:      "Gujarátí Lángwej",
			language.GUZ:     "Gusí Lángwej",
			language.GV:      "Mánks Lángwej",
			language.HA:      "Háusá Lángwej",
			language.HAW:     "Hawaii Lángwej",
			language.HE:      "Híbru Lángwej",
			language.HI:      "Híndi Lángwej",
			language.HMN:     "Mọ́ng Lángwej",
			language.HR:      "Kroéshia Lángwej",
			language.HSB:     "Sóbiá Lángwej di ọ́p-ọ́p wan",
			language.HT:      "Haítí Kriol",
			language.HU:      "Họngári Lángwej",
			language.HY:      "Armẹ́niá Lángwej",
			language.IA:      "Intalíngwuá Lángwej",
			language.ID:      "Indoníshia Lángwej",
			language.IG:      "Igbo Lángwej",
			language.II:      "Síchuan Yi",
			language.IS:      "Aíslánd Lángwej",
			language.IT:      "Ítáli Lángwej",
			language.JA:      "Japan Lángwej",
			language.JGO:     "Ngómbá Lángwej",
			language.JMC:     "Machámẹ́ Lángwej",
			language.JV:      "Javáníz Lángwej",
			language.KA:      "Jọ́jiá Lángwej",
			language.KAB:     "Kabail Lángwej",
			language.KAM:     "Kámbá Lángwej",
			language.KDE:     "Makọ́ndẹ́ Lángwej",
			language.KEA:     "Kábúvẹrdiánu Lángwej",
			language.KHQ:     "Koyra Chíní Lángwej",
			language.KI:      "Kikúyú Lángwej",
			language.KK:      "Kazák Lángwej",
			language.KKJ:     "Kákó Lángwej",
			language.KL:      "Kalálísút Lángwej",
			language.KLN:     "Kálẹ́njín Lángwej",
			language.KM:      "Kmaí Lángwej",
			language.KN:      "Kánnáda Lángwej",
			language.KO:      "Koriá Lángwej",
			language.KOK:     "Kónkéní Lángwej",
			language.KS:      "Kashmírí Lángwej",
			language.KSB:     "Shambala",
			language.KSF:     "Bafiá Lángwej",
			language.KSH:     "Kọlónián Lángwej",
			language.KU:      "Kọ́dísh Lángwej",
			language.KW:      "Kọ́nish Lángwej",
			language.KY:      "Kiẹ́gíz Lángwej",
			language.LA:      "Látín Lángwej",
			language.LAG:     "Langi Lángwej",
			language.LB:      "Lọ́ksémbọ́g Lángwej",
			language.LG:      "Gánda Lángwej",
			language.LKT:     "Lakótá Lángwej",
			language.LN:      "Lingálá Lángwej",
			language.LO:      "Láo Lángwej",
			language.LRC:     "Nọ́tán Lúrí Lángwej",
			language.LT:      "Lituéniá Lángwej",
			language.LU:      "Lúbá-Katángá Lángwej",
			language.LUO:     "Luó Lángwej",
			language.LUY:     "Luyia Lángwej",
			language.LV:      "Látvián Lángwej",
			language.MAS:     "Masaí Lángwej",
			language.MER:     "Mẹ́rú Lángwej",
			language.MFE:     "Morísiẹ́n Lángwej",
			language.MG:      "Malagásí Lángwej",
			language.MGH:     "Makúwá-Mító",
			language.MGO:     "Mẹta’ Lángwej",
			language.MI:      "Maórí Lángwej",
			language.MK:      "Masẹdóniá Lángwej",
			language.ML:      "Maléyálám Lángwej",
			language.MN:      "Mọngóliá Lángwej",
			language.MR:      "Marátí Lángwej",
			language.MS:      "Malé Lángwej",
			language.MT:      "Mọ́ltá Lángwej",
			language.MUA:     "Mundáng Lángwej",
			language.MUL:     "Plẹ́ntí Lángwej-dẹm",
			language.MY:      "Bọ́ma Lángwej",
			language.MZN:     "Mazandẹrání Lángwej",
			language.NAQ:     "Naámá Lángwej",
			language.NB:      "Nọwẹ́jiá Bokmál Lángwej",
			language.ND:      "Nọ́tán Ndẹbẹlẹ Lángwej",
			language.NDS:     "Ló Jámán Lángwej",
			language.NE:      "Nẹpálí Lángwej",
			language.NL:      "Dọch Lángwej",
			language.NL_BE:   "Flẹ́mish Lángwej",
			language.NMG:     "Kwasió Lángwej",
			language.NN:      "Nọwẹ́jiá Niúnọsk",
			language.NNH:     "Ngiẹ́mbọn Lángwej",
			language.NUS:     "Núa",
			language.NY:      "Nyánja",
			language.NYN:     "Nyankólẹ",
			language.OM:      "Orómó",
			language.OR:      "Ódiá",
			language.OS:      "Osẹ́tik",
			language.PA:      "Punjábi",
			language.PCM:     "Naijíriá Píjin",
			language.PL:      "Pólánd Lángwej",
			language.PRG:     "Prúshia",
			language.PS:      "Páshto",
			language.PT:      "Pọtiugiz",
			language.PT_BR:   "Brazíl Pọtiugíz",
			language.PT_PT:   "Yúróp Pọtiugíz",
			language.QU:      "Kẹchuá",
			language.RM:      "Románsh",
			language.RN:      "Rúndi",
			language.RO:      "Romániá Lángwej",
			language.ROF:     "Rómbo",
			language.RU:      "Rọshiá Lángwej",
			language.RW:      "Kinyarwánda Lángwej",
			language.RWK:     "Rwa",
			language.SA:      "Sánskrit",
			language.SAH:     "Sakhá",
			language.SAQ:     "Sambúru",
			language.SBP:     "Sangu",
			language.SD:      "Síndí",
			language.SE:      "Nọ́tán Sámí Lángwej",
			language.SEH:     "Sẹ́ná",
			language.SES:     "Kóiraboró Sẹ́nní Lángwej",
			language.SG:      "sàngo",
			language.SHI:     "Táchẹ́lit",
			language.SI:      "Sínhala",
			language.SK:      "Slóvak",
			language.SL:      "Slovẹ́niá Lángwej",
			language.SM:      "Samóá Lángwej",
			language.SMN:     "Ínárí Sámí Lángwej",
			language.SN:      "Shóna",
			language.SO:      "Sọmáli",
			language.SQ:      "Albéniá Lángwej",
			language.SR:      "Sẹrbiá Lángwej",
			language.ST:      "Saútán Sóto",
			language.SU:      "Sọ́ndaniz",
			language.SV:      "Suwídẹ́n Lángwej",
			language.SW:      "Swahíli",
			language.TA:      "tàmil",
			language.TE:      "Tẹlugu",
			language.TEO:     "Tẹ́so",
			language.TG:      "Tájik",
			language.TH:      "Taí",
			language.TI:      "Tigrínyá",
			language.TK:      "Tọ́kmẹn",
			language.TO:      "Tóngan",
			language.TR:      "Tọ́ki",
			language.TT:      "Tatá",
			language.TWQ:     "Tasawak",
			language.TZM:     "Mídúl Atlás Támazígt Lángwej",
			language.UG:      "Wiúgọ",
			language.UK:      "Yukrénia",
			language.UND:     "Lángwej wé nóbọ́di sabi",
			language.UR:      "Úrdú",
			language.UZ:      "Úzbẹk",
			language.VAI:     "Vaí",
			language.VI:      "Viẹ́tnám Lángwej",
			language.VO:      "Vólapiuk",
			language.VUN:     "Vúnjo",
			language.WAE:     "Wọ́lsa",
			language.WO:      "Wólof",
			language.XH:      "Kọ́sa",
			language.XOG:     "sóga",
			language.YAV:     "Yangbẹn",
			language.YI:      "Yídish",
			language.YO:      "Yorubá",
			language.YUE:     "Chainiz Kántọniz",
			language.ZGH:     "Gẹ́nárál Morókó Támazígt Lángwej",
			language.ZH:      "Chainiz, Mandarin",
			language.ZH_HANS: "Ízí Mandarín Chainíz Lángwej",
			language.ZH_HANT: "Tradíshọ́nál Mandarín Chainíz Lángwej",
			language.ZU:      "Zúlu",
			language.ZXX:     "Nó Lángwéj Kọ́ntẹnt",
		},
		Territories: cldr.Territories{
			territory.V_001: "Wọld",
			territory.V_002: "Áfríka",
			territory.V_003: "Nọ́t Amẹ́ríka",
			territory.V_005: "Saút Amẹ́ríka",
			territory.V_009: "Oshẹnia",
			territory.V_011: "Wẹ́stán Áfríka",
			territory.V_013: "Mídúl Amẹ́ríka",
			territory.V_014: "Ístán Áfríká",
			territory.V_015: "Nọ́tán Áfríka",
			territory.V_017: "Mídúl Áfríka",
			territory.V_018: "Saútán Áfríka",
			territory.V_019: "Amẹ́ríkas",
			territory.V_021: "Nọ́tán Amẹ́ríka",
			territory.V_029: "Karíbián",
			territory.V_030: "Ístán Éshia",
			territory.V_034: "Saútán Éshia",
			territory.V_035: "Saútíst Éshiá",
			territory.V_039: "Saútán Yúrop",
			territory.V_053: "Ọstraléshia",
			territory.V_054: "Mẹlanẹíshia",
			territory.V_057: "Maikroníshia Ríjọn",
			territory.V_061: "Poliníshiá",
			territory.V_142: "Éshia",
			territory.V_143: "Mídúl Éshia",
			territory.V_145: "Wẹ́stán Éshia",
			territory.V_150: "Yúrop",
			territory.V_151: "Ístán Yúrop",
			territory.V_154: "Nọ́tán Yúrop",
			territory.V_155: "Wẹ́stán Yúrop",
			territory.V_202: "Áfríka Éria Biló Sahára",
			territory.V_419: "Látín Amẹ́ríka",
			territory.AC:    "Asẹ́nshọ́n Aíland",
			territory.AD:    "Andọ́ra",
			territory.AE:    "Yunaítẹ́d Áráb Ẹ́mírets",
			territory.AF:    "Afgánístan",
			territory.AG:    "Antígwua & Barbúda",
			territory.AI:    "Angwíla",
			territory.AL:    "Albénia",
			territory.AM:    "Armẹ́niá",
			territory.AO:    "Angóla",
			territory.AQ:    "Antáktíka",
			territory.AR:    "Ajẹntína",
			territory.AS:    "Amẹ́ríká Samoa",
			territory.AT:    "Ọ́stria",
			territory.AU:    "Ọstrélia",
			territory.AW:    "Arúba",
			territory.AX:    "Ọ́lánd Aílands",
			territory.AZ:    "Azẹrbaijan",
			territory.BA:    "Bọ́zniá & Hẹzẹgovína",
			territory.BB:    "Barbédọs",
			territory.BD:    "Bangladẹsh",
			territory.BE:    "Bẹ́ljọm",
			territory.BF:    "Burkína Fáso",
			territory.BG:    "Bọlgéria",
			territory.BH:    "Barein",
			territory.BI:    "Burúndi",
			territory.BJ:    "Binin",
			territory.BL:    "Sént Batẹlẹ́mi",
			territory.BM:    "Bẹmiúda",
			territory.BN:    "Brunẹi",
			territory.BO:    "Bolívia",
			territory.BQ:    "Karíbián Nẹ́dalands",
			territory.BR:    "Brázil",
			territory.BS:    "Bahámas",
			territory.BT:    "Butan",
			territory.BV:    "Buvẹ́ Aíland",
			territory.BW:    "Botswána",
			territory.BY:    "Bẹ́larus",
			territory.BZ:    "Bẹliz",
			territory.CA:    "Kánáda",
			territory.CC:    "Kókós Aílands",
			territory.CD:    "Kóngo (DRC)",
			territory.CF:    "Sẹ́ntrál Áfríkán Ripọ́blik",
			territory.CG:    "Kóngó (Ripọ́blik)",
			territory.CH:    "Swítsaland",
			territory.CI:    "Kót Divua",
			territory.CK:    "Kúk Aílands",
			territory.CL:    "Chílẹ",
			territory.CM:    "Kamẹrun",
			territory.CN:    "Chaína",
			territory.CO:    "Kolómbia",
			territory.CP:    "Klipatọ́n Aíland",
			territory.CR:    "Kósta Ríka",
			territory.CU:    "Kiúbá",
			territory.CV:    "Kép Vẹ́d",
			territory.CW:    "Kiurásao",
			territory.CX:    "Krísmás Aíland",
			territory.CY:    "Saíprọs",
			territory.CZ:    "Chẹ́k Ripọ́blik",
			territory.DE:    "Jámáni",
			territory.DG:    "Diẹ́gó Garsia",
			territory.DJ:    "Jibúti",
			territory.DK:    "Dẹ́nmak",
			territory.DM:    "Dọmíníka",
			territory.DO:    "Dọmíníka Ripọ́blik",
			territory.DZ:    "Aljíria",
			territory.EA:    "Sẹúta & Mẹ́líla",
			territory.EC:    "Ẹ́kwuádọ",
			territory.EE:    "Ẹstónia",
			territory.EG:    "Íjipt",
			territory.EH:    "Wẹ́stán Sahára",
			territory.ER:    "Ẹritrẹ́a",
			territory.ES:    "Spen",
			territory.ET:    "Ẹtiópia",
			territory.EU:    "Yurópián Yúniọ́n",
			territory.EZ:    "Yúróéria",
			territory.FI:    "Fínland",
			territory.FJ:    "Fíji",
			territory.FK:    "Fọ́klánd Aílands (Íslás Malvínas)",
			territory.FM:    "Maikroníshia",
			territory.FO:    "Fáro Aílands",
			territory.FR:    "Frans",
			territory.GA:    "Gabọn",
			territory.GB:    "UK",
			territory.GD:    "Grẹnéda",
			territory.GE:    "Jọ́jia",
			territory.GF:    "Frẹ́nch Giána",
			territory.GG:    "Guẹnzi",
			territory.GH:    "Gána",
			territory.GI:    "Jibrọ́lta",
			territory.GL:    "Grínland",
			territory.GM:    "Gámbia",
			territory.GN:    "Gíni",
			territory.GP:    "Guadalúpẹ",
			territory.GQ:    "Ikwétóriál Gíni",
			territory.GR:    "Gris",
			territory.GS:    "Saút Jọ́jia an Saút Sándwích Aílands",
			territory.GT:    "Guátẹmála",
			territory.GU:    "Guam",
			territory.GW:    "Gíní-Bisáu",
			territory.GY:    "Gayána",
			territory.HK:    "Họng Kọng",
			territory.HM:    "Hiád & MakDónáld Aílands",
			territory.HN:    "Họndúras",
			territory.HR:    "Kroéshia",
			territory.HT:    "Haíti",
			territory.HU:    "Họ́ngári",
			territory.IC:    "Kenerí Aílands",
			territory.ID:    "Indoníshia",
			territory.IE:    "Ayaland",
			territory.IL:    "Ízrẹl",
			territory.IM:    "Aíl ọf Man",
			territory.IN:    "Índia",
			territory.IO:    "Brítísh Índián Óshen Tẹ́rẹ́tri",
			territory.IQ:    "Irak",
			territory.IR:    "Iran",
			territory.IS:    "Aísland",
			territory.IT:    "Ítáli",
			territory.JE:    "Jẹ́si",
			territory.JM:    "Jamaíka",
			territory.JO:    "Jọ́dan",
			territory.JP:    "Japan",
			territory.KE:    "Kẹ́nya",
			territory.KG:    "Kẹjístan",
			territory.KH:    "Kambódia",
			territory.KI:    "Kiribáti",
			territory.KM:    "Kọ́mọ́ros",
			territory.KN:    "Sent Kits & Nẹ́vis",
			territory.KP:    "Nọ́t Koria",
			territory.KR:    "Saút Koria",
			territory.KW:    "Kuwét",
			territory.KY:    "Kéman Aílands",
			territory.KZ:    "Kazakstan",
			territory.LA:    "Laos",
			territory.LB:    "Lẹ́bánọn",
			territory.LC:    "Sent Lúshia",
			territory.LI:    "Líktẹ́nstain",
			territory.LK:    "Sri Lánka",
			territory.LR:    "Laibẹ́ria",
			territory.LS:    "Lẹsóto",
			territory.LT:    "Lituénia",
			territory.LU:    "Lọ́ksẹ́mbọg",
			territory.LV:    "Látvia",
			territory.LY:    "Líbia",
			territory.MA:    "Morọko",
			territory.MC:    "Mọ́náko",
			territory.MD:    "Mọldóva",
			territory.ME:    "Mọntinígro",
			territory.MF:    "Sent Mátin",
			territory.MG:    "Madagáska",
			territory.MH:    "Máshál Aílands",
			territory.MK:    "Nọ́t Masidónia",
			territory.ML:    "Máli",
			territory.MM:    "Miánma (Bọ́ma)",
			territory.MN:    "Mọngólia",
			territory.MO:    "Makáo",
			territory.MP:    "Nọ́tán Mariána Aílands",
			territory.MQ:    "Matínik",
			territory.MR:    "Mọriténia",
			territory.MS:    "Mọntsẹrat",
			territory.MT:    "Mọ́lta",
			territory.MU:    "Mọríshọs",
			territory.MV:    "Mọ́ldivs",
			territory.MW:    "Maláwi",
			territory.MX:    "Mẹ́ksíko",
			territory.MY:    "Maléshia",
			territory.MZ:    "Mozámbik",
			territory.NA:    "Namíbia",
			territory.NC:    "Niú Kalẹdónia",
			territory.NE:    "Nizhẹr",
			territory.NF:    "Nọ́fọlk Aíland",
			territory.NG:    "Naijíria",
			territory.NI:    "Nikarágwua",
			territory.NL:    "Nẹ́dalands",
			territory.NO:    "Nọ́we",
			territory.NP:    "Nẹ́pal",
			territory.NR:    "Náuru",
			territory.NU:    "Niúẹ",
			territory.NZ:    "Niú Zíland",
			territory.OM:    "Oman",
			territory.PA:    "Pánáma",
			territory.PE:    "Pẹ́ru",
			territory.PF:    "Frẹ́nch Poliníshia",
			territory.PG:    "Pápuá Niú Gíni",
			territory.PH:    "Fílípins",
			territory.PK:    "Pakístan",
			territory.PL:    "Póland",
			territory.PM:    "Sent Piẹr & Míkẹlọn",
			territory.PN:    "Pítkén Aílands",
			territory.PR:    "Puẹ́rto Ríkọ",
			territory.PS:    "Pálẹ́stain",
			territory.PT:    "Pọ́túgal",
			territory.PW:    "Palau",
			territory.PY:    "Párágwue",
			territory.QA:    "Kata",
			territory.QO:    "Rimót Pát ọf Oshẹ́nia",
			territory.RE:    "Réyúniọn",
			territory.RO:    "Ruménia",
			territory.RS:    "Sẹ́bia",
			territory.RU:    "Rọ́shia",
			territory.RW:    "Ruwánda",
			territory.SA:    "Saúdí Arébia",
			territory.SB:    "Sólómọ́n Aílands",
			territory.SC:    "Sẹ́chẹls",
			territory.SD:    "Sudan",
			territory.SE:    "Swídẹn",
			territory.SG:    "Singapọ",
			territory.SH:    "Sent Hẹlẹ́na",
			territory.SI:    "Slovẹ́nia",
			territory.SJ:    "Sválbad & Jén Meyẹn",
			territory.SK:    "Slovékia",
			territory.SL:    "Siẹ́ra Líon",
			territory.SM:    "San Maríno",
			territory.SN:    "Sẹ́nẹ́gal",
			territory.SO:    "Sọmália",
			territory.SR:    "Súrínam",
			territory.SS:    "Saút Sudan",
			territory.ST:    "Sao Tómé & Prínsípẹ",
			territory.SV:    "El Sálvádọ",
			territory.SX:    "Sint Mátin",
			territory.SY:    "Síria",
			territory.SZ:    "Swáziland",
			territory.TA:    "Trístán da Kúna",
			territory.TC:    "Tọks an Kaíkọ́s Aílands",
			territory.TD:    "Chad",
			territory.TF:    "Frẹ́nch Saútán Tẹ́rẹ́tris",
			territory.TG:    "Tógo",
			territory.TH:    "Taíland",
			territory.TJ:    "Tajíkstan",
			territory.TK:    "Tókẹ́lau",
			territory.TL:    "Íst Tímọ",
			territory.TM:    "Tọkmẹ́nístan",
			territory.TN:    "Tuníshia",
			territory.TO:    "Tónga",
			territory.TR:    "Tọ́ki",
			territory.TT:    "Trínídad & Tobágo",
			territory.TV:    "Tuválu",
			territory.TW:    "Taiwan",
			territory.TZ:    "Tanzánia",
			territory.UA:    "Yukrein",
			territory.UG:    "Yugánda",
			territory.UM:    "U.S. Faá Faá Aílands",
			territory.UN:    "Yunaítẹd Néshọns",
			territory.US:    "US",
			territory.UY:    "Yúrugwue",
			territory.UZ:    "Uzbẹ́kistan",
			territory.VA:    "Vátíkán Síti",
			territory.VC:    "Sent Vínsẹnt & Grẹ́nádians",
			territory.VE:    "Vẹnẹzuẹ́la",
			territory.VG:    "Brítísh Vájín Aílands",
			territory.VI:    "U.S. Vájín Aílands",
			territory.VN:    "Viẹ́tnam",
			territory.VU:    "Vanuátu",
			territory.WF:    "Wọ́lis & Fiutúna",
			territory.WS:    "Samóa",
			territory.XA:    "To yúz atifíshál vọis wẹ́n yu de tọk",
			territory.XB:    "Atífíshál Tú-Wé Dairẹ́kshọn",
			territory.XK:    "Kósóvo",
			territory.YE:    "Yẹ́mẹn",
			territory.YT:    "Meyọt",
			territory.ZA:    "Saút Áfríka",
			territory.ZM:    "Zámbia",
			territory.ZW:    "Zimbábwẹ",
			territory.ZZ:    "Ríjọn Wé Nóbọ́di Sabí",
		},
	}
}
