// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_az_Latn_AZ() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "d MMMM y, EEEE", Long: "d MMMM y", Medium: "d MMM y", Short: "dd.MM.yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "yan", Feb: "fev", Mar: "mar", Apr: "apr", May: "may", Jun: "iyn", Jul: "iyl", Aug: "avq", Sep: "sen", Oct: "okt", Nov: "noy", Dec: "dek"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "yanvar", Feb: "Fevral", Mar: "mart", Apr: "Aprel", May: "May", Jun: "İyun", Jul: "İyul", Aug: "Avqust", Sep: "Sentyabr", Oct: "Oktyabr", Nov: "Noyabr", Dec: "dekabr"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "B.", Mon: "B.E.", Tue: "Ç.A.", Wed: "Ç.", Thu: "C.A.", Fri: "C.", Sat: "Ş."}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "7", Mon: "1", Tue: "2", Wed: "3", Thu: "4", Fri: "5", Sat: "6"}, Short: cldr.CalendarDayFormatNameValue{Sun: "B.", Mon: "B.E.", Tue: "Ç.A.", Wed: "Ç.", Thu: "C.A.", Fri: "C.", Sat: "Ş."}, Wide: cldr.CalendarDayFormatNameValue{Sun: "bazar", Mon: "bazar ertəsi", Tue: "çərşənbə axşamı", Wed: "çərşənbə", Thu: "cümə axşamı", Fri: "cümə", Sat: "şənbə"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_az]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "Andora Pesetası", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "Birləşmiş Ərəb Əmirlikləri Dirhəmi", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "Əfqanıstan Əfqanisi (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Əfqanıstan Əfqanisi", Symbol: "AFN"},
				currency.ALK: cldr.Currency{DisplayName: "Albaniya Leki (1946–1965)", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "Albaniya Leki", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Ermənistan Dramı", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Niderland Antilyası Gilderi", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Anqola Kvanzası", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "Anqola Kvanzasi (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "Anqola Yeni Kvanzası (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "Anqola Kvanzası (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "Argentina avstralı", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "Argentina pesosu (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "Argentina Pesosu", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "Avstriya Şillinqi", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "Avstraliya Dolları", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Aruba Florini", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "Azərbaycan Manatı (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Azərbaycan Manatı", Symbol: "₼"},
				currency.BAD: cldr.Currency{DisplayName: "Bosniya-Herseqovina Dinarı", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Bosniya-Herseqovina Markası", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "Barbados Dolları", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Banqladeş Takası", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "Belçika Frankı (deyşirik)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "Belçika Frankı", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "Belçika Frankı (finans)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "Bolqarıstan Levası", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "Bolqarıstan Levi", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "Bəhreyn Dinarı", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Burundi Frankı", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Bermuda Dolları", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Bruney Dolları", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Boliviya Bolivianosu", Symbol: "BOB"},
				currency.BOP: cldr.Currency{DisplayName: "Boliviya pesosu", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "Boliviya mvdolı", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "Braziliya kruzeyro novası", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "Braziliya kruzadosu", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "Braziliya kruzeyrosu (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "Braziliya Realı", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "Braziliya kruzado novası", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "Braziliya kruzeyrosu", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "Baham Dolları", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Butan Nqultrumu", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "Burmis Kyatı", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Botsvana Pulası", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "Belarus Yeni Rublu (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "Belarus Rublu", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Belarus Rublu (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "Beliz Dolları", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Kanada Dolları", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Konqo Frankı", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "WIR Avro", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "İsveçrə Frankı", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "WIR Frankası", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Çili Pesosu", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Çin Yuanı (ofşor)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "Çin Yuanı", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Kolumbiya Pesosu", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "Kosta Rika Kolonu", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "Serbiya Dinarı (2002–2006)", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "Çexoslavakiya Korunası", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "Kuba Çevrilən Pesosu", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Kuba Pesosu", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Kape Verde Eskudosu", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "Kipr Paundu", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Çexiya Korunası", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "Şərq Almaniya Ostmarkı", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "Alman Markası", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "Cibuti Frankı", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Danimarka Kronu", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Dominika Pesosu", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Əlcəzair Dinarı", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "Ekvador Sukresi", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "Estoniya Krunu", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Misir Funtu", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Eritreya Nakfası", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "İspan Pesetası (A account)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "İspan Pesetası (dəyşirik)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "İspan Pesetası", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Efiopiya Bırrı", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Avro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Fin Markası", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "Fici Dolları", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Folklend Adaları Funtu", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "Fransız Markası", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "Britaniya Funtu", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "Gürcüstan Kupon Lariti", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "Gürcüstan Larisi", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "Qana Sedisi (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Qana Sedisi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Cəbəli-Tariq Funtu", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Qambiya Dalasisi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Qvineya Frankı", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "Qvineya Sulisi", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "Ekvatoriya Gvineya Ekvele Quneanası", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "Yunan Draçması", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "Qvatemala Küetzalı", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "Portugal Qvineya Eskudosu", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "Qvineya-Bisau Pesosu", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "Qayana Dolları", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Honq Konq Dolları", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Honduras Lempirası", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "Xorvatiya Dinarı", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "Xorvatiya Kunası", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Haiti Qourdu", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Macarıstan Forinti", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "İndoneziya Rupisi", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "İrlandiya Paundu", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "İzrail Paundu", Symbol: ""},
				currency.ILR: cldr.Currency{DisplayName: "İsrail Şekeli (1980–1985)", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "İsrail Yeni Şekeli", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Hindistan Rupisi", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "İraq Dinarı", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "İran Rialı", Symbol: "IRR"},
				currency.ISJ: cldr.Currency{DisplayName: "İslandiya Kronu (1918–1981)", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "İslandiya Kronu", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "İtaliya Lirası", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "Yamayka Dolları", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "İordaniya Dinarı", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Yaponiya Yeni", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Keniya Şillinqi", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Qırğızıstan Somu", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Kamboca Rieli", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Komor Frankı", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Şimali Koreya Vonu", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "Cənubi Koreya Vonu", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Küveyt Dinarı", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Kayman Adaları Dolları", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Qazaxıstan Tengesi", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Laos Kipi", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Livan Funtu", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Şri Lanka Rupisi", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Liberiya Dolları", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Lesoto Lotisi", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Litva Liti", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "Litva Talonası", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "Luksemburq Frankası (dəyişik)", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "Luksemburq Frankası", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "Luksemburq Frankası (finans)", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "Latviya Latı", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "Latviya Rublu", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "Liviya Dinarı", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Mərakeş Dirhəmi", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "Mərakeş Frankası", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Moldova Leyi", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Madaqaskar Ariarisi", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "Madaqaskar Frankası", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "Makedoniya Dinarı", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "Makedoniya Dinarı (1992–1993)", Symbol: ""},
				currency.MLF: cldr.Currency{DisplayName: "Mali Frankı", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Myanma Kiyatı", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Monqoliya Tuqriki", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Makao Patakası", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Mavritaniya Ugiyası (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Mavritaniya Ugiyası", Symbol: "MRU"},
				currency.MTP: cldr.Currency{DisplayName: "Maltiz Paundu", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Mavriki Rupisi", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Maldiv Rufiyası", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Malavi Kvaçası", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Meksika Pesosu", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "Meksika gümüş pesosu", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "Malayziya Ringiti", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "Mozambik Eskudosu", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "Mozambik Metikalı (1980–2006)", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Mozambik Metikalı", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Namibiya Dolları", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Nigeriya Nairası", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "Nikaraqua kordobu", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "Nikaraqua Kordobası", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "Hollandiya Gilderi", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "Norveç Kronu", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Nepal Rupisi", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Yeni Zelandiya Dolları", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Oman Rialı", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Panama Balboası", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "Peru Inti", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Peru Solu", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "Peru Solu (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Papua Yeni Qvineya Kinası", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Filippin Pesosu", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Pakistan Rupisi", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Polşa Zlotısı", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "Polşa Zlotısı (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "Portuqal Eskudosu", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "Paraqvay Quaranisi", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Qatar Rialı", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "Rodezian Dolları", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "Rumıniya Leyi (1952–2006)", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Rumıniya Leyi", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Serbiya Dinarı", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Rusiya Rublu", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "Rusiya Rublu (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Ruanda Frankı", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Səudiyyə Riyalı", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Solomon Adaları Dolları", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Seyşel Rupisi", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "Sudan Funtu", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "İsveç Kronu", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Sinqapur Dolları", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Müqəddəs Yelena Funtu", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "Sloveniya Toları", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "Slovak Korunası", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Sierra Leon Leonu", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Somali Şillinqi", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Surinam Dolları", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "Cənubi Sudan Funtu", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "San Tom və Prinsip Dobrası (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "San Tom və Prinsip Dobrası", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "Sovet Rublu", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "El Salvador kolonu", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "Suriya Funtu", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Svazilend Lilangenini", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Tayland Batı", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "Tacikistan Rublu", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "Tacikistan Somonisi", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "Türkmənistan Manatı (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Türkmənistan Manatı", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Tunis Dinarı", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Tonqa Panqası", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "Timor Eskudu", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "Türkiyə Lirəsi (1922–2005)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "Türkiyə Lirəsi", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Trinidad və Tobaqo Dolları", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Tayvan Yeni Dolları", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Tanzaniya Şillinqi", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Ukrayna Qrivnası", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "Ukrayna Karbovenesası", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "Uqanda Şillinqi (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "Uqanda Şillinqi", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "ABŞ Dolları", Symbol: "US$"},
				currency.USN: cldr.Currency{DisplayName: "ABŞ dolları (yeni gün)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "ABŞ dolları (həmin gün)", Symbol: ""},
				currency.UYI: cldr.Currency{DisplayName: "Uruqvay pesosu Unidades Indexadas", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "Uruqvay Pesosu (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "Uruqvay Pesosu", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Özbəkistan Somu", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "Venesuela Bolivarı (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "Venesuela Bolivarı (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Venesuela Bolivarı", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Vyetnam Donqu", Symbol: "₫"},
				currency.VNN: cldr.Currency{DisplayName: "Vyetnam Donqu (1978–1985)", Symbol: ""},
				currency.VUV: cldr.Currency{DisplayName: "Vanuatu Vatusu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Samoa Talası", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Kamerun Frankı", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "gümüş", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "qızıl", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "Şərqi Karib Dolları", Symbol: "EC$"},
				currency.XFO: cldr.Currency{DisplayName: "Fransız Gızıl Frankı", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "Fransız UİC Frankı", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "Fil Dişi Sahili Frankı", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "Palladium", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "Fransız Polineziyası Frankı", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "Platinum", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "Naməlum Valyuta", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "Yəmən Dinarı", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "Yəmən Rialı", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "Yuqoslaviya Dinarı (1966–1990)", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "Yuqoslaviya Yeni Dinarı (1994–2002)", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "Yuqoslaviya Dinarı (1990–1992)", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "Cənubi Afrika Randı (finans)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Cənubi Afrika Randı", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Zambiya Kvaçası (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Zambiya Kvaçası", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "Zair Yeni Zairi (1993–1998)", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "Zair Zairi (1971–1993)", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "Zimbabve Dolları (1980–2008)", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "Zimbabve Dolları (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "Zimbabve Dolları (2008)", Symbol: ""},
			},
		},
	}
}
