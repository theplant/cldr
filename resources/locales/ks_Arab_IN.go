// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_ks_Arab_IN() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "ج", Feb: "ف", Mar: "م", Apr: "ا", May: "م", Jun: "ج", Jul: "ج", Aug: "ا", Sep: "س", Oct: "س", Nov: "ا", Dec: "ن"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "جنؤری", Feb: "فرؤری", Mar: "مارٕچ", Apr: "اپریل", May: "میٔ", Jun: "جوٗن", Jul: "جوٗلایی", Aug: "اگست", Sep: "ستمبر", Oct: "اکتوٗبر", Nov: "نومبر", Dec: "دسمبر"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "آتھوار", Mon: "ژٔندٕروار", Tue: "بۆموار", Wed: "بودوار", Thu: "برؠسوار", Fri: "جُمہ", Sat: "بٹوار"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "ا", Mon: "ژ", Tue: "ب", Wed: "ب", Thu: "ب", Fri: "ج", Sat: "ب"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "اَتھوار", Mon: "ژٔندرٕروار", Tue: "بۆموار", Wed: "بودوار", Thu: "برؠسوار", Fri: "جُمہ", Sat: "بٹوار"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_ks]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "", Group: "", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "#,##,##0.###", Currency: "¤\u00a0#,##,##0.00", Percent: "#,##,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "اؠڑورَن پیسِٹا", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "مُتحدہ عرب اِمارات دِرہم", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "افغان افغٲنی", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "اؠلبینِیَن لِک", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "اَرمانؠن ڈرؠم", Symbol: ""},
				currency.ANG: cldr.Currency{DisplayName: "نایدَرلینڑ اؠنٹٕلیٖیَن گِلڑَر", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "اؠنگولَن کوانزا", Symbol: "Kz"},
				currency.AOR: cldr.Currency{DisplayName: "اؠنگولَن کوانزا رؠجِسٹاڑو", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "أرجَنٹیٖن اَسٹرل", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "أرجَنٹیٖن پِسو", Symbol: "$"},
				currency.ATS: cldr.Currency{DisplayName: "آسٹریَن شِلِنگ", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "آسٹریلِیَن ڈالَر", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "اَروبَن فِلورِن", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "آزَرباجانی مَنَٹ", Symbol: ""},
				currency.BAD: cldr.Currency{DisplayName: "بوزنِیاہَرزِگووِنا دیٖنار", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "بوزنِیاہَرزِگووِنا کَنوٲٹیبٕل مارٕک", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "بابیڑِیَن ڈالَر", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "بَنگلادیٖشی ٹَکا", Symbol: "৳"},
				currency.BEF: cldr.Currency{DisplayName: "بَلجِیَن فرینک", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "بیلگیرِیَن ہاڑ لِو", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "بحریٖنی دیٖنار", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "بُرُنڑِین فرینک", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "بٔرمیوٗڑَن ڈالَر", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "برونی ڈالَر", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "بولِوِیَن بولوینو", Symbol: "Bs"},
				currency.BOP: cldr.Currency{DisplayName: "بولویَن پِسو", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "بولوِیَن مَوڈال", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "برازیٖلین کروزِرو نووو", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "برازیٖلین کروزیڑو", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "برازیٖلین کروزِرو", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "برازیٖلین رِیَل", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "برازیٖلین کروزیڑو نووو", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "برازیٖلین کروزیرو", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "بہامِیَن ڈالر", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "بوٗٹینیٖز نگُلٹرم", Symbol: ""},
				currency.BUK: cldr.Currency{DisplayName: "بٔرمیٖز کیٹ", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "بوٹٕسوانَن پُلا", Symbol: "P"},
				currency.BYB: cldr.Currency{DisplayName: "بِلیروشِیَن نِو رِبٕل", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "بِلیروشِیَن رِبٕل", Symbol: "р."},
				currency.BYR: cldr.Currency{DisplayName: "بِلیروشِیَن رِبٕل (۲۰۰۰–۲۰۱۶)", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "بِلِزی ڈالر", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "کینَڑِیَن ڈالر", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "کونگولیٖز فریک", Symbol: ""},
				currency.CHE: cldr.Currency{DisplayName: "وِر یوٗرو", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "سُوِز فریک", Symbol: ""},
				currency.CHW: cldr.Currency{DisplayName: "وِر فریک", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "چِلِن یوٗنِڑیدیٖز ڑِ فومیٹو", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "چِلِن پِسو", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "چینیٖز یَن رِنمِنبی", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "کولَمبِین پِسو", Symbol: "$"},
				currency.COU: cldr.Currency{DisplayName: "ِٖیوٗنِڑیڑ ڑِ ویلور رِیل", Symbol: ""},
				currency.CRC: cldr.Currency{DisplayName: "کوسٹا رِکَن کولَن", Symbol: "₡"},
				currency.CSD: cldr.Currency{DisplayName: "پرون سٔربِین ڈالر", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "چِکوسولوواک ہاڑ کوروٗنا", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "کیوٗبَن پِسو", Symbol: "$"},
				currency.CYP: cldr.Currency{DisplayName: "کیپروٹ پَوُڑ", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "چیک کوریٖنا", Symbol: "Kč"},
				currency.DDM: cldr.Currency{DisplayName: "مٔشرِقی جٔرمَن مارٕک", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "جٔرمَن مارٕک", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "ڈٔنِش کرون", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "ڈومِنِکَن پِسو", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "اؠلجیرِیَن ڈیٖنار", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "اؠسٹونِیَن کرون", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "اِجِپٹِیَن پَوُنڑ", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "رِٹریٖن نَفکا", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "سِپینِش پیسِٹا", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "اِتھوپِیَن بِر", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "یوٗرو", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "فِنِش مارکا", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "فِجین ڈالر", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "فیکلینڑِس آیلینڑ پونڑ", Symbol: "£"},
				currency.FRF: cldr.Currency{DisplayName: "فرانسِسی فریک", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "برطٲنوی پاونڑ سٹٔرلِنگ", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "جارجِیَن کیوٗپَن لَرِٹ", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "جارجِیَن لاری", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "گَنییَن سؠڑی(۱۹۷۹–۲٠٠۷)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "گَنییَن سؠڑی", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "گِبریلٹَر پَاونڑ", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "گیمبِیاہُک دلاسی", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "گِنِیَن فرینک", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "گِنِیَن سِلی", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "اِکویٹورِیَل گِنِیَن اؠکویٖل", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "گریٖسُک ڑرؠکما", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "گواٹَمالَن قیوٗٹزَل", Symbol: "Q"},
				currency.GWE: cldr.Currency{DisplayName: "پورتگیٖزُک گِنی اؠسکیوٗڑو", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "گِنی بِساوُک پؠسو", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "گَیَنیٖزُک ڑالَر", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "حانگ کانگُک ڑالَر", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "حونڑورنُک لؠمپیٖرا", Symbol: "L"},
				currency.HRD: cldr.Currency{DisplayName: "کروایشنُک دیٖنار", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "کروایشنُک کوٗنا", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "حیشَنُک گوڑ", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "حَنگیرِیَن فورِنٹ", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "اِنڑونیشیاہُک رُپِیاہ", Symbol: "Rp"},
				currency.IEP: cldr.Currency{DisplayName: "اَیرلینڑُک پاونڑ", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "اِزرٲیِلی پاونڑ", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "اِزرٲیِلی نٔوۍ شؠقٕل", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "ہِندُستٲنۍ رۄپَے", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "ایٖراقُک دیٖنار", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "ایٖرانُک رِیال", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "اَیسلینڑُک کرونا", Symbol: "kr"},
				currency.ITL: cldr.Currency{DisplayName: "اِٹلیٖ یُک لیٖرا", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "جَمَیکاہُک ڑالَر", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "جَرڑینیاہُک دیٖنار", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "جاپانُک یَن", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "کؠنیَن شِلِنگ", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "کِرگِستانُک سوم", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "کَمبوڑِیاہُک رِیال", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "کومورِیَن فرینک", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "جنوٗبی کورِیَن وَن", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "ساوتھ کورِیَن وَن", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "قُویتُک دیٖنار", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "کیمین ججیٖرُک ڑالَر", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "کزاکِستان ٹینج", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "لَوٹِیَن کِپ", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "لیبنیٖزُک پاونڑ", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "سری لَنکاہٕچ رۄپَے", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "لَیبیرِیَن ڑالَر", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "لِسوتھو لوٹی", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "لِتھوینِیَن لِٹاس", Symbol: "Lt"},
				currency.LTT: cldr.Currency{DisplayName: "لِتھوینِیَن ٹؠلوناس", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "لَکزٕمبورگِیَن کَنؤرٹِبٕل فرینک", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "لَکزٕمبورگِیَن فرینک", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "لَکزٕمبوگ فَینانشَل فرینک", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "لَتوِیَن لیٹس", Symbol: "Ls"},
				currency.LVR: cldr.Currency{DisplayName: "لَتوِیَن رَبٕل", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "لِبیَن دیٖنار", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "موروکَن دِرہَم", Symbol: ""},
				currency.MAF: cldr.Currency{DisplayName: "موروکَن فرینک", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "مولڑووین لیوٗ", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "مؠلؠگیسی اؠریَری", Symbol: "Ar"},
				currency.MGF: cldr.Currency{DisplayName: "مؠلؠگیسی فرینک", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "مؠکَڑونِیَن دیٖنار", Symbol: ""},
				currency.MLF: cldr.Currency{DisplayName: "میلِیَن فرینک", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "مِیانما کیاٹ", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "مۄنگولِیَن ٹُگرِک", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "مؠکانیٖز پَٹاکا", Symbol: ""},
				currency.MRO: cldr.Currency{DisplayName: "مورِٹینِیَن عوگیوٗیا (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "مورِٹینِیَن عوگیوٗیا", Symbol: ""},
				currency.MTL: cldr.Currency{DisplayName: "مالٹیٖزُک لیٖرا", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "مالٹیٖزُک پاونڑ", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "مورؠشِیاہٕچ رۄپَے", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "مالدِیٖوِیَن رُفِیا", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "مؠلیوِیَن کواچا", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "مؠکسِکَن پؠسو", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "مؠکسِکَن سِلوَر پؠسو (۱۸۶۱–۱۹۹۲)", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "مَلیشِیَن رِنگِٹ", Symbol: "RM"},
				currency.MZE: cldr.Currency{DisplayName: "موزیمبِکَن سکیوٗڑو", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "پرون موزیمبِکَن مؠٹِکَل", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "موزیمبِکَن مؠٹِکَل", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "نامِبِیَن ڑالَر", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "نَیجیرِیَن ڑالَر", Symbol: "₦"},
				currency.NIC: cldr.Currency{DisplayName: "نِکؠراگُوؠن کورڑوبا", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "نِکؠراگُوؠن کورڑوبا اورو", Symbol: "C$"},
				currency.NLG: cldr.Currency{DisplayName: "ڈَچ گِلڑَر", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "نوروییِنُک کرون", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "نیپالٕچ رۄپَے", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "نِوزیٖلینڑُک ڑالَر", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "اومِنی رِیال", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "پانامانِیَن بالبوز", Symbol: ""},
				currency.PEI: cldr.Currency{DisplayName: "پٔریوٗوِیَن اِنٹی", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "پٔریوٗوِیَن سولٕز", Symbol: ""},
				currency.PES: cldr.Currency{DisplayName: "پٔریوٗوِیَن سول (۱۸۶۳–۱۹۶۵)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "نیوٗ پیپُعا گِنِیَن کیٖنا", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "پھِلِپایِٔن پؠسو", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "پاکِستٲنۍ رۄپَے", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "پولِش زلوٹی", Symbol: "zł"},
				currency.PLZ: cldr.Currency{DisplayName: "پولِش زلوٹی(۱۹۵٠–۱۹۹۵)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "پورتُگیٖز اؠسکیوٗڑو", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "پؠرؠگیوٗوَیَن گُعارانی", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "قطاری رِیال", Symbol: ""},
				currency.RHD: cldr.Currency{DisplayName: "رھوڑیشِیَن ڑالَر", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "اولڑ رومانِیَن لؠیوٗ", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "رومانِیَن لؠیوٗ", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "سٔربِیَن دیٖنار", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "رٔشیَن رَبٕل", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "رٔشیَن رَبٕل(۱۹۹۱–۱۹۹۸)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "روانڑَن فرانک", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "سودیٖیُک رِیال", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "سولَمَن جٔزیٖرُک ڈالَر", Symbol: "$"},
				currency.SDD: cldr.Currency{DisplayName: "پرون سوٗڈانُک دیٖنار", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "سوٗڈانُک پونڈ", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "پرون سوٗڈانُک پونڈ", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "سویٖڈِش کَرونا", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "سِنگاپورُک ڈالَر", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "سینٹ ہیلِنا پونڈ", Symbol: "£"},
				currency.SIT: cldr.Currency{DisplayName: "سلووینُک ٹولَر", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "سلووَک کَرونا", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "سومالی شِلِنگ", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "سُریٖنامُک ڈالَر", Symbol: "$"},
				currency.SRG: cldr.Currency{DisplayName: "سُریٖنام گِلدَر", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "", Symbol: "Db"},
				currency.SUR: cldr.Currency{DisplayName: "سوویت روبٕل", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "سَلویدَرُک کولَن", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "سیٖریاہُک پونڈ", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "سوازی لِلَنگینی", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "تھایھک بات", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "تاجکِستانُک رَبٕل", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "تاجِکتانُک سَمونی", Symbol: ""},
				currency.TMM: cldr.Currency{DisplayName: "تُکَمَنِستانُک مَنَت", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "ٹُنیشیاہُک دیٖنار", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "ٹونگَن پانگا", Symbol: "T$"},
				currency.TPE: cldr.Currency{DisplayName: "ٹیموریسو ایکیٖڈو", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "پرون تُرکِش لیرا", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "تُرکیہُک لیرا", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "ٹرِنہِ ڈیڈ تہٕ ٹوبیگو ڈالَر", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "نۆو تیوانُک ڈالَر", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "تَنزانیاہُک شِلِنگ", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "یوٗکرینیاہُک ہرِوِنیا", Symbol: "₴"},
				currency.UAK: cldr.Currency{DisplayName: "یوٗکرینیاہُک کاربووَنیٹس", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "اُگاداہُک شِلِنگ(۱۹۶۶–۱۹۸۷)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "اُگاداہُک شِلِنگ", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "یوٗ ایس ڈالَر", Symbol: "US$"},
				currency.USN: cldr.Currency{DisplayName: "یوٗ ایس ڈالَر(پَگاہ)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "یوٗ ایس ڈالَر(تَمی دًۄہ)", Symbol: ""},
				currency.UYI: cldr.Currency{DisplayName: "اُرگایَن پیسو یوٗنِڈیڈَس اِنڈیکسَس", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "اُرگایَن پیسو(۱۹۷۵–۱۹۹۳)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "اُروٗگایَن پیسو", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "اُبیکِستان سوم", Symbol: ""},
				currency.VEB: cldr.Currency{DisplayName: "وینٕزوٗلیُک بولِوَر (۱۸۷۱–۲۰۰۸)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "وینٕزوٗلیُک بولِوَر (2008–2018)", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "وینٕزوٗلیُک بولِوَر", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "وِیَنَمُک ڈانگ", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "وَنوٗاَتوٗ وَتوٗ", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "سَمون تَلا", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "سی ایف اے فرینک بی ایٖ اے سی", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "رۄپھ", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "سۄن", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "یوٗرپی کَمپوسِٹ یوٗنِٹ", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "یوٗرپی مونِٹَری یوٗنِٹ", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "یوٗرپی یوٗنِٹ آف ایکاوُنٹ (ایکس بی سی)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "یوٗرپی یوٗنِٹ آف ایکاوُنٹ (ایکس بی ڈی)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "مَشرِقی کیرِبِیَن ڈالَر", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "خاص ڈرایِنگ رایٹس", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "یوٗرپی کَرَنسی یوٗنِٹ", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "فرینچ گولڈ فرینک", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "فرینچ یوٗ اے سی فرینک", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "سی ایف اے فرینک بی سی ایٖ اے او", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "پُلیڈیَم", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "سی ایف پی فرینک", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "پلیٹِنَم", Symbol: ""},
				currency.XRE: cldr.Currency{DisplayName: "آر آے این ایٖ ٹی فَنڈ", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "ٹیسٹِنگ کَرَنسی کوڈ", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "اَنزٲنۍ یا نالَگہٕ ہار سِکہٕ", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "یَمنُک دیٖنار", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "یَمنُک رِیال", Symbol: ""},
				currency.YUD: cldr.Currency{DisplayName: "یوگوسلاوِیَن ہاڑ دیٖنار", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "یوگوسلاوِیَن نووِے دیٖنار", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "یوگوسلاوِیَن کَنؤٹِبٕل دیٖنار", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "ساوُتھ افریٖکاہُک رینڈ", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "زِمبابیُک کواچا (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "زِمبابیُک کواچا", Symbol: "ZK"},
				currency.ZRN: cldr.Currency{DisplayName: "زایرِیَن نِو زایِر", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "زَیرُک ڈالَر", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "زِمبابِیُک ڈالَر", Symbol: ""},
			},
		},
	}
}
