// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_ar_OM() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE، d MMMM y", Long: "d MMMM y", Medium: "dd\u200f/MM\u200f/y", Short: "d\u200f/M\u200f/y"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "غرينتش{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "يناير", Feb: "فبراير", Mar: "مارس", Apr: "أبريل", May: "مايو", Jun: "يونيو", Jul: "يوليو", Aug: "أغسطس", Sep: "سبتمبر", Oct: "أكتوبر", Nov: "نوفمبر", Dec: "ديسمبر"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "ي", Feb: "ف", Mar: "م", Apr: "أ", May: "و", Jun: "ن", Jul: "ل", Aug: "غ", Sep: "س", Oct: "ك", Nov: "ب", Dec: "د"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "يناير", Feb: "فبراير", Mar: "مارس", Apr: "أبريل", May: "مايو", Jun: "يونيو", Jul: "يوليو", Aug: "أغسطس", Sep: "سبتمبر", Oct: "أكتوبر", Nov: "نوفمبر", Dec: "ديسمبر"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "الأحد", Mon: "الاثنين", Tue: "الثلاثاء", Wed: "الأربعاء", Thu: "الخميس", Fri: "الجمعة", Sat: "السبت"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "ح", Mon: "ن", Tue: "ث", Wed: "ر", Thu: "خ", Fri: "ج", Sat: "س"}, Short: cldr.CalendarDayFormatNameValue{Sun: "أحد", Mon: "إثنين", Tue: "ثلاثاء", Wed: "أربعاء", Thu: "خميس", Fri: "جمعة", Sat: "سبت"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "الأحد", Mon: "الاثنين", Tue: "الثلاثاء", Wed: "الأربعاء", Thu: "الخميس", Fri: "الجمعة", Sat: "السبت"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "ص", PM: "م"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "ص", PM: "م"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "صباحًا", PM: "مساءً"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_ar]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "٫", Group: "٬", Negative: "\u061c-", Percent: "٪\u061c", PerMille: "؉"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "بيستا أندوري", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "درهم إماراتي", Symbol: "د.إ.\u200f"},
				currency.AFA: cldr.Currency{DisplayName: "أفغاني - 1927-2002", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "أفغاني", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "ليك ألباني", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "درام أرميني", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "غيلدر أنتيلي هولندي", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "كوانزا أنغولي", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "كوانزا أنجولي - 1977-1990", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "كوانزا أنجولي جديدة - 1990-2000", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "كوانزا أنجولي معدلة - 1995 - 1999", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "استرال أرجنتيني", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "بيزو أرجنتيني - 1983-1985", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "بيزو أرجنتيني", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "شلن نمساوي", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "دولار أسترالي", Symbol: "AU$"},
				currency.AWG: cldr.Currency{DisplayName: "فلورن أروبي", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "مانات أذريبجاني", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "مانات أذربيجان", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "دينار البوسنة والهرسك", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "مارك البوسنة والهرسك قابل للتحويل", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "دولار بربادوسي", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "تاكا بنغلاديشي", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "فرنك بلجيكي قابل للتحويل", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "فرنك بلجيكي", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "فرنك بلجيكي مالي", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "ليف بلغاري", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "دينار بحريني", Symbol: "د.ب.\u200f"},
				currency.BIF: cldr.Currency{DisplayName: "فرنك بروندي", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "دولار برمودي", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "دولار بروناي", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "بوليفيانو بوليفي", Symbol: "BOB"},
				currency.BOP: cldr.Currency{DisplayName: "بيزو بوليفي", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "مفدول بوليفي", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "نوفو كروزايرو برازيلي - 1967-1986", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "كروزادو برازيلي", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "كروزايرو برازيلي - 1990-1993", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "ريال برازيلي", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "دولار باهامي", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "نولتوم بوتاني", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "كيات بورمي", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "بولا بتسواني", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "روبل بيلاروسي جديد - 1994-1999", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "روبل بيلاروسي", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "روبل بيلاروسي (٢٠٠٠–٢٠١٦)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "دولار بليزي", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "دولار كندي", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "فرنك كونغولي", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "فرنك سويسري", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "بيزو تشيلي", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "يوان صيني (في الخارج)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "يوان صيني", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "بيزو كولومبي", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "كولن كوستاريكي", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "دينار صربي قديم", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "كرونة تشيكوسلوفاكيا", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "بيزو كوبي قابل للتحويل", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "بيزو كوبي", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "اسكودو الرأس الأخضر", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "جنيه قبرصي", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "كرونة تشيكية", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "أوستمارك ألماني شرقي", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "مارك ألماني", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "فرنك جيبوتي", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "كرونة دنماركية", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "بيزو الدومنيكان", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "دينار جزائري", Symbol: "د.ج.\u200f"},
				currency.EEK: cldr.Currency{DisplayName: "كرونة استونية", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "جنيه مصري", Symbol: "ج.م.\u200f"},
				currency.ERN: cldr.Currency{DisplayName: "ناكفا أريتري", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "بيزيتا إسباني", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "بير أثيوبي", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "يورو", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "ماركا فنلندي", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "دولار فيجي", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "جنيه جزر فوكلاند", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "فرنك فرنسي", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "جنيه إسترليني", Symbol: "UK£"},
				currency.GEL: cldr.Currency{DisplayName: "لارى جورجي", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "سيدي غاني", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "سيدي غانا", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "جنيه جبل طارق", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "دلاسي غامبي", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "فرنك غينيا", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "سيلي غينيا", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "اكويل جونينا غينيا الاستوائيّة", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "دراخما يوناني", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "كوتزال غواتيمالا", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "اسكود برتغالي غينيا", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "بيزو غينيا بيساو", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "دولار غيانا", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "دولار هونغ كونغ", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "ليمبيرا هنداروس", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "دينار كرواتي", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "كونا كرواتي", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "جوردى هايتي", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "فورينت هنغاري", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "روبية إندونيسية", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "جنيه إيرلندي", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "جنيه إسرائيلي", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "شيكل إسرائيلي جديد", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "روبية هندي", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "دينار عراقي", Symbol: "د.ع.\u200f"},
				currency.IRR: cldr.Currency{DisplayName: "ريال إيراني", Symbol: "ر.إ."},
				currency.ISK: cldr.Currency{DisplayName: "كرونة أيسلندية", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "ليرة إيطالية", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "دولار جامايكي", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "دينار أردني", Symbol: "د.أ.\u200f"},
				currency.JPY: cldr.Currency{DisplayName: "ين ياباني", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "شلن كينيي", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "سوم قيرغستاني", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "رييال كمبودي", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "فرنك جزر القمر", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "وون كوريا الشمالية", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "وون كوريا الجنوبية", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "دينار كويتي", Symbol: "د.ك.\u200f"},
				currency.KYD: cldr.Currency{DisplayName: "دولار جزر كيمن", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "تينغ كازاخستاني", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "كيب لاوسي", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "جنيه لبناني", Symbol: "ل.ل.\u200f"},
				currency.LKR: cldr.Currency{DisplayName: "روبية سريلانكية", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "دولار ليبيري", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "لوتي ليسوتو", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "ليتا ليتوانية", Symbol: "Lt"},
				currency.LTT: cldr.Currency{DisplayName: "تالوناس ليتواني", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "فرنك لوكسمبرج قابل للتحويل", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "فرنك لوكسمبرج", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "فرنك لوكسمبرج المالي", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "لاتس لاتفيا", Symbol: "Ls"},
				currency.LVR: cldr.Currency{DisplayName: "روبل لاتفيا", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "دينار ليبي", Symbol: "د.ل.\u200f"},
				currency.MAD: cldr.Currency{DisplayName: "درهم مغربي", Symbol: "د.م.\u200f"},
				currency.MAF: cldr.Currency{DisplayName: "فرنك مغربي", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "ليو مولدوفي", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "أرياري مدغشقر", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "فرنك مدغشقر", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "دينار مقدوني", Symbol: "MKD"},
				currency.MLF: cldr.Currency{DisplayName: "فرنك مالي", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "كيات ميانمار", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "توغروغ منغولي", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "باتاكا ماكاوي", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "أوقية موريتانية - 1973-2017", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "أوقية موريتانية", Symbol: "أ.م."},
				currency.MTL: cldr.Currency{DisplayName: "ليرة مالطية", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "جنيه مالطي", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "روبية موريشيوسية", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "روفيه جزر المالديف", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "كواشا مالاوي", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "بيزو مكسيكي", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "بيزو فضي مكسيكي - 1861-1992", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "رينغيت ماليزي", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "اسكود موزمبيقي", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "متكال موزمبيقي", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "دولار ناميبي", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "نايرا نيجيري", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "كوردوبة نيكاراجوا", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "قرطبة نيكاراغوا", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "جلدر هولندي", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "كرونة نرويجية", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "روبية نيبالي", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "دولار نيوزيلندي", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "ريال عماني", Symbol: "ر.ع.\u200f"},
				currency.PAB: cldr.Currency{DisplayName: "بالبوا بنمي", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "سول بيروفي", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "كينا بابوا غينيا الجديدة", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "بيزو فلبيني", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "روبية باكستاني", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "زلوتي بولندي", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "زلوتي بولندي - 1950-1995", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "اسكود برتغالي", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "غواراني باراغواي", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "ريال قطري", Symbol: "ر.ق.\u200f"},
				currency.RHD: cldr.Currency{DisplayName: "دولار روديسي", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "ليو روماني قديم", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "ليو روماني", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "دينار صربي", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "روبل روسي", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "روبل روسي - 1991-1998", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "فرنك رواندي", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "ريال سعودي", Symbol: "ر.س.\u200f"},
				currency.SBD: cldr.Currency{DisplayName: "دولار جزر سليمان", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "روبية سيشيلية", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "دينار سوداني", Symbol: "د.س.\u200f"},
				currency.SDG: cldr.Currency{DisplayName: "جنيه سوداني", Symbol: "ج.س."},
				currency.SDP: cldr.Currency{DisplayName: "جنيه سوداني قديم", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "كرونة سويدية", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "دولار سنغافوري", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "جنيه سانت هيلين", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "تولار سلوفيني", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "كرونة سلوفاكية", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "ليون سيراليوني", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "شلن صومالي", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "دولار سورينامي", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "جلدر سورينامي", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "جنيه جنوب السودان", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "دوبرا ساو تومي وبرينسيبي - 1977-2017", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "دوبرا ساو تومي وبرينسيبي", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "روبل سوفيتي", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "كولون سلفادوري", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "ليرة سورية", Symbol: "ل.س.\u200f"},
				currency.SZL: cldr.Currency{DisplayName: "ليلانجيني سوازيلندي", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "باخت تايلاندي", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "روبل طاجيكستاني", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "سوموني طاجيكستاني", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "مانات تركمنستاني", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "مانات تركمانستان", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "دينار تونسي", Symbol: "د.ت.\u200f"},
				currency.TOP: cldr.Currency{DisplayName: "بانغا تونغا", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "اسكود تيموري", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "ليرة تركي", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "ليرة تركية", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "دولار ترينداد وتوباغو", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "دولار تايواني", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "شلن تنزاني", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "هريفنيا أوكراني", Symbol: "UAH"},
				currency.UGS: cldr.Currency{DisplayName: "شلن أوغندي - 1966-1987", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "شلن أوغندي", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "دولار أمريكي", Symbol: "US$"},
				currency.USN: cldr.Currency{DisplayName: "دولار أمريكي (اليوم التالي)\u200f", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "دولار أمريكي (نفس اليوم)\u200f", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "بيزو أوروجواي - 1975-1993", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "بيزو اوروغواي", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "سوم أوزبكستاني", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "بوليفار فنزويلي - 1871-2008", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "بوليفار فنزويلي - 2008–2018", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "بوليفار فنزويلي", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "دونج فيتنامي", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "فاتو فانواتو", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "تالا ساموا", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "فرنك وسط أفريقي", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "فضة", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "ذهب", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "الوحدة الأوروبية المركبة", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "الوحدة المالية الأوروبية", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "الوحدة الحسابية الأوروبية", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "(XBD)وحدة الحساب الأوروبية", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "دولار شرق الكاريبي", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "حقوق السحب الخاصة", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "وحدة النقد الأوروبية", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "فرنك فرنسي ذهبي", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "(UIC)فرنك فرنسي", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "فرنك غرب أفريقي", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "بالاديوم", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "فرنك سي إف بي", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "البلاتين", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "كود اختبار العملة", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "عملة غير معروفة", Symbol: "***"},
				currency.YDD: cldr.Currency{DisplayName: "دينار يمني", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "ريال يمني", Symbol: "ر.ي.\u200f"},
				currency.YUD: cldr.Currency{DisplayName: "دينار يوغسلافي", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "دينار يوغسلافي قابل للتحويل", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "راند جنوب أفريقيا -مالي", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "راند جنوب أفريقيا", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "كواشا زامبي - 1968-2012", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "كواشا زامبي", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "زائير زائيري جديد", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "زائير زائيري", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "دولار زمبابوي", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "دولار زمبابوي 2009", Symbol: ""},
			},
		},
	}
}
