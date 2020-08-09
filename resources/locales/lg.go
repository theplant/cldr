// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_lg() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mar", Apr: "Apu", May: "Maa", Jun: "Juu", Jul: "Jul", Aug: "Agu", Sep: "Seb", Oct: "Oki", Nov: "Nov", Dec: "Des"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Janwaliyo", Feb: "Febwaliyo", Mar: "Marisi", Apr: "Apuli", May: "Maayi", Jun: "Juuni", Jul: "Julaayi", Aug: "Agusito", Sep: "Sebuttemba", Oct: "Okitobba", Nov: "Novemba", Dec: "Desemba"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Sab", Mon: "Bal", Tue: "Lw2", Wed: "Lw3", Thu: "Lw4", Fri: "Lw5", Sat: "Lw6"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "B", Tue: "L", Wed: "L", Thu: "L", Fri: "L", Sat: "L"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Sabbiiti", Mon: "Balaza", Tue: "Lwakubiri", Wed: "Lwakusatu", Thu: "Lwakuna", Fri: "Lwakutaano", Sat: "Lwamukaaga"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_lg]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "", Group: "", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "", Currency: "#,##0.00¤", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Diraamu eya Emireeti", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Kwanza ey’Angola", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Doola ey’Awusiturelya", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "Dinaali ey’eBaareeni", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Faranga ey’eburundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Pula ey’eBotiswana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Doola ey’eKanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Faranga ey’eKongo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Faranga ey’eSwitizirandi", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Yuwani Reniminibi ey’eCayina", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Esikudo ey’Keepu Veredi", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Faranga ey’eJjibuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dinaali ey’Aligerya", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Pawundi ey’eMisiri", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Nakifa ey’Eritureya", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Biiru ey’Esyopya", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Yuro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Pawundi ey’eBungereza", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Sedi ey’eGana", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi ey’eGambya", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Faranga ey’eGini", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupiya ey’eBuyindi", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Yeni ey’eJapani", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Silingi ey’eKenya", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Faranga ey’eKomoro", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Doola ey’eLiberya", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Loti ey’eLesoso", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Dinaali ey’eLibya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Diraamu ey’eMoroko", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Faranga ey’eMalagase", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Wugwiya ey’eMawritenya (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Wugwiya ey’eMawritenya", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupiya ey’eMawurisyasi", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Kwaca ey’eMalawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Metikaali ey’eMozambiiki", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Doola ey’eNamibiya", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Nayira ey’eNayijerya", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "", Symbol: "NZ$"},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "", Symbol: "₲"},
				currency.RON: cldr.Currency{DisplayName: "", Symbol: "lei"},
				currency.RUB: cldr.Currency{DisplayName: "", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Faranga ey’eRwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Riyaali ey’eBuwarabu", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Rupiya ey’eSesere", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Dinaali ey’eSudaani", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Pawundi ey’eSudaani", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Pawundi ey’eSenti Herena", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Lewone", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Silingi ey’eSomaliya", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Dobura ey’eSantome ne Purincipe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Dobura ey’eSantome ne Purincipe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Dinaali ey’eTunizya", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Silingi ey’eTanzaniya", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Silingi eya Yuganda", Symbol: "USh"},
				currency.USD: cldr.Currency{DisplayName: "Doola ey’Amerika", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "Faranga ey’omu Afirika eya wakati", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Faranga ey’omu Afirika ey’ebugwanjuba", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Randi ey’eSawusafirika", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwaca ey’eZambya (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kwaca ey’eZambya", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Doola ey’eZimbabwe", Symbol: ""},
			},
		},
	}
}
