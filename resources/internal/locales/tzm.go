// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_tzm() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "tzm",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Yen", Feb: "Yeb", Mar: "Mar", Apr: "Ibr", May: "May", Jun: "Yun", Jul: "Yul", Aug: "Ɣuc", Sep: "Cut", Oct: "Kṭu", Nov: "Nwa", Dec: "Duj"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "Y", Feb: "Y", Mar: "M", Apr: "I", May: "M", Jun: "Y", Jul: "Y", Aug: "Ɣ", Sep: "C", Oct: "K", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Yennayer", Feb: "Yebrayer", Mar: "Mars", Apr: "Ibrir", May: "Mayyu", Jun: "Yunyu", Jul: "Yulyuz", Aug: "Ɣuct", Sep: "Cutanbir", Oct: "Kṭuber", Nov: "Nwanbir", Dec: "Dujanbir"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Asa", Mon: "Ayn", Tue: "Asn", Wed: "Akr", Thu: "Akw", Fri: "Asm", Sat: "Asḍ"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "A", Mon: "A", Tue: "A", Wed: "A", Thu: "A", Fri: "A", Sat: "A"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Asamas", Mon: "Aynas", Tue: "Asinas", Wed: "Akras", Thu: "Akwas", Fri: "Asimwas", Sat: "Asiḍyas"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "Zdat azal", PM: "Ḍeffir aza"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "Zdat azal", PM: "Ḍeffir aza"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Derhem Uymarati", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Kwanza Unguli", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Ḍular Usṭrali", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "Ḍinar Ubaḥrayni", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Frank Uburundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Pula Ubutswani", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Ḍular Ukanadi", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Frank Ukunguli", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Frank Uswisri", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Ywan Renminbi Ucinwi", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Iskudu Ukabuvirdyani", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Frank Uğibuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Ḍinar Udzayri", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Junih Umiṣṛi", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfa Uyritri", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Birr Uyityuppi", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Uṛu", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Junih Ubriṭani", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Sidi Uɣani", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi Agambi", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Frank Uɣini", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupi Uḥindi", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Yann Ujappuni", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Cillin Ukini", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Frank Uqumuri", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Ḍular Ulibiri", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Luti Ulusuṭi", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Ḍinar Ulibi", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Derhem Umeṛṛuki", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Aryari Umalɣaci", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Uqiyya Umuritani (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Uqiyya Umuritani", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupi Umurisi", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Kwača Umalawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Mitikal Umuzambiqi", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Ḍular Unamibi", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Nayra Unijiri", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "Frank Urwandi", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Ryal Usaεudi", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Rupi Usicili", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Junih Usudani", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Junih Usudani (1956–2007)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Junih Usantehilini", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Lyun Usirralyuni", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Cilin Uṣumali", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Dubra Usawṭumi (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Dubra Usawṭumi", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilanjini Uswazi", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Ḍinar Utunsi", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Cilin Uṭanzani", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Cilin Uɣandi (1966–1987)", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Ḍular Umirikani", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "Frank CFA (BEAC)", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Frank CFA (BCEAO)", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Rand Ufriki Unzul", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwača Uzambi (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kwača Uzambi", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Ḍular Uzimbabwi", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK:  "Takanit",
			language.AM:  "Tamharit",
			language.AR:  "Taεrabt",
			language.BE:  "Tabilarusit",
			language.BG:  "Tabelɣarit",
			language.BN:  "Tabinɣalit",
			language.CS:  "Tačikt",
			language.DE:  "Talmanit",
			language.EL:  "Tayunanit",
			language.EN:  "Tanglizt",
			language.ES:  "tasbelyunit",
			language.FA:  "Tafarisit",
			language.FR:  "Tafṛansist",
			language.HA:  "Tahawsat",
			language.HI:  "Tahindit",
			language.HU:  "Tahenɣarit",
			language.ID:  "Tindunisit",
			language.IG:  "Tigbut",
			language.IT:  "Taṭalyant",
			language.JA:  "Tajappunit",
			language.JV:  "Tajavanit",
			language.KM:  "Taxmert ,Talammast",
			language.KO:  "Takurit",
			language.MS:  "Tamalizit",
			language.MY:  "Taburmanit",
			language.NE:  "Tanippalit",
			language.NL:  "Tahulanḍit",
			language.PA:  "Tabenjabit",
			language.PL:  "Tappulunit",
			language.PT:  "Taburtuɣalit",
			language.RO:  "Taṛumanit",
			language.RU:  "Tarusit",
			language.RW:  "Tarwandit",
			language.SO:  "Taṣumalit",
			language.SV:  "Taswidit",
			language.TA:  "Tatamilt",
			language.TH:  "Taṭayt",
			language.TR:  "Taturkit",
			language.TZM: "Tamaziɣt n laṭlaṣ",
			language.UK:  "Tukranit",
			language.UR:  "Turdut",
			language.VI:  "Taviṭnamit",
			language.YO:  "Tayurubat",
			language.ZH:  "Tacinwit,Mandarin",
			language.ZU:  "tazulut",
		},
		Territories: cldr.Territories{
			territory.AD: "Anḍurra",
			territory.AE: "Imarat Tiεrabin Tidduklin",
			territory.AF: "Afɣanistan",
			territory.AG: "Antigwa d Barbuda",
			territory.AI: "Angwilla",
			territory.AL: "Albanya",
			territory.AM: "Arminya",
			territory.AO: "Angula",
			territory.AR: "Arjuntin",
			territory.AS: "Samwa Imirikaniyyin",
			territory.AT: "Ustriyya",
			territory.AU: "Usṭralya",
			territory.AW: "Aruba",
			territory.AZ: "Azerbiǧan",
			territory.BA: "Busna-d-Hirsik",
			territory.BB: "Barbadus",
			territory.BD: "Bangladic",
			territory.BE: "Beljika",
			territory.BF: "Burkina Fasu",
			territory.BG: "Belɣarya",
			territory.BH: "Baḥrayn",
			territory.BI: "Burundi",
			territory.BJ: "Binin",
			territory.BM: "Birmuda",
			territory.BN: "Brunay",
			territory.BO: "Bulivya",
			territory.BR: "Bṛazil",
			territory.BS: "Bahamas",
			territory.BT: "Buṭan",
			territory.BW: "Butswana",
			territory.BY: "Bilarusya",
			territory.BZ: "Biliz",
			territory.CA: "Kanada",
			territory.CD: "Tagduda Tadimuqraṭit n Kungu",
			territory.CF: "Tagduda n Afrika Wammas",
			territory.CG: "Kungu",
			territory.CH: "Swisra",
			territory.CI: "Taɣazut n Uszer",
			territory.CK: "Tigzirin n Kuk",
			territory.CL: "Ccili",
			territory.CM: "Kamerun",
			territory.CN: "Ṣṣin",
			territory.CO: "Kulumbya",
			territory.CR: "Kusṭa Rika",
			territory.CU: "kuba",
			territory.CV: "Tigzirin n Iɣf Uzegzaw",
			territory.CY: "Qubrus",
			territory.CZ: "Tagduda n Čik",
			territory.DE: "Almanya",
			territory.DJ: "Ǧibuti",
			territory.DK: "Danmark",
			territory.DM: "Ḍuminika",
			territory.DO: "Tagduda n Ḍuminikan",
			territory.DZ: "Dzayer",
			territory.EC: "Ikwaḍur",
			territory.EE: "Isṭunya",
			territory.EG: "Miṣr",
			territory.ER: "Iritrya",
			territory.ES: "Sbanya",
			territory.ET: "Ityupya",
			territory.FI: "Finlanḍa",
			territory.FJ: "Fiji",
			territory.FK: "Tigzirin n Falkland",
			territory.FM: "Mikrunizya",
			territory.FR: "Fṛansa",
			territory.GA: "Gabun",
			territory.GB: "Tagelda Taddukelt",
			territory.GD: "Grinada",
			territory.GE: "Jyurjya",
			territory.GF: "Guyana Tafransist",
			territory.GH: "Ɣana",
			territory.GI: "Jibralṭar",
			territory.GL: "Grinlanḍa",
			territory.GM: "Gambya",
			territory.GN: "Ɣinya",
			territory.GP: "Gwadalup",
			territory.GQ: "Ɣinya Tikwaṭur it",
			territory.GR: "Yunan",
			territory.GT: "Gwatimala",
			territory.GU: "Gwam",
			territory.GW: "Ɣinya-Bissaw",
			territory.GY: "Guyana",
			territory.HN: "Hinduras",
			territory.HR: "Krwatya",
			territory.HT: "Hayti",
			territory.HU: "Henɣarya",
			territory.ID: "Indunizya",
			territory.IE: "Irlanḍa",
			territory.IL: "Israeil",
			territory.IN: "Hind",
			territory.IO: "Amur n Agaraw Uhindi Ubṛiṭani",
			territory.IQ: "Ɛiraq",
			territory.IR: "Iran",
			territory.IS: "Islanḍa",
			territory.IT: "Iṭalya",
			territory.JM: "Jamayka",
			territory.JO: "Urḍun",
			territory.JP: "Jjappun",
			territory.KE: "Kinya",
			territory.KG: "Kirɣistan",
			territory.KH: "Kambudj",
			territory.KI: "Kiribati",
			territory.KM: "Qumur",
			territory.KN: "Santekits d Nivis",
			territory.KP: "Kurya Tugafat",
			territory.KR: "Kurya Tunẓult",
			territory.KW: "Kuwwayt",
			territory.KY: "Tigzirin n Kayman",
			territory.KZ: "Kazaxistan",
			territory.LA: "Laws",
			territory.LB: "Lubnan",
			territory.LC: "Santelusya",
			territory.LI: "Lictencṭayn",
			territory.LK: "Srilanka",
			territory.LR: "Libirya",
			territory.LS: "Lisuṭu",
			territory.LT: "Litwanya",
			territory.LU: "Liksumburg",
			territory.LV: "Liṭṭunya",
			territory.LY: "Libya",
			territory.MA: "Meṛṛuk",
			territory.MC: "Munaku",
			territory.MD: "Mulḍavya",
			territory.MG: "Madaɣacqar",
			territory.MH: "Tigzirin n Marcal",
			territory.ML: "Mali",
			territory.MM: "Myanmar",
			territory.MN: "Manɣulya",
			territory.MP: "Tigzirin n Maryana Tugafat",
			territory.MQ: "Martinik",
			territory.MR: "Muritanya",
			territory.MS: "Muntsirra",
			territory.MT: "Malṭa",
			territory.MU: "Muris",
			territory.MV: "Maldiv",
			territory.MW: "Malawi",
			territory.MX: "Miksik",
			territory.MY: "Malizya",
			territory.MZ: "Muzambiq",
			territory.NA: "Namibya",
			territory.NC: "kalidunya Tamaynut",
			territory.NE: "Nnijer",
			territory.NF: "Tigzirt Nurfulk",
			territory.NG: "Nijiria",
			territory.NI: "Nikaragwa",
			territory.NL: "Hulanḍa",
			territory.NO: "Nnurwij",
			territory.NP: "Nippal",
			territory.NR: "Nawru",
			territory.NU: "Niwi",
			territory.NZ: "Zilanḍa Tamaynut",
			territory.OM: "Ɛumman",
			territory.PA: "Panama",
			territory.PE: "Piru",
			territory.PF: "Pulinizya Tafransist",
			territory.PG: "Papwa Ɣinya Tamaynut",
			territory.PH: "Filippin",
			territory.PK: "Pakistan",
			territory.PL: "Pulunya",
			territory.PM: "Santepyir d Mikelun",
			territory.PN: "Pitkirn",
			territory.PR: "Purturiku",
			territory.PS: "Agemmaḍ Ugut d Ɣazza Ifilisṭiniyen",
			territory.PT: "Purtuɣal",
			territory.PW: "Palu",
			territory.PY: "Paragway",
			territory.QA: "Qaṭar",
			territory.RE: "Riyyunyun",
			territory.RO: "Ṛumanya",
			territory.RU: "Rusya",
			territory.RW: "Ruwwanḍa",
			territory.SA: "Ssaεudiyya Taεrabt",
			territory.SB: "Tigzirin n Salumun",
			territory.SC: "Ssicil",
			territory.SD: "Ssudan",
			territory.SE: "Ssewwid",
			territory.SG: "Sanɣafura",
			territory.SH: "Santehilin",
			territory.SI: "Sluvinya",
			territory.SK: "Sluvakya",
			territory.SL: "Siralyun",
			territory.SM: "Sanmarinu",
			territory.SN: "Ssiniɣal",
			territory.SO: "Ṣṣumal",
			territory.SR: "Surinam",
			territory.ST: "Sawṭumi d Prinsip",
			territory.SV: "Salvaḍur",
			territory.SY: "Surya",
			territory.SZ: "Swazilanḍa",
			territory.TC: "Tigzirin Turkiyyin d Tikaykusin",
			territory.TD: "Tcad",
			territory.TG: "Ṭṭugu",
			territory.TH: "Ṭaylanḍa",
			territory.TJ: "Ṭaǧikistan",
			territory.TK: "Tuklu",
			territory.TL: "Timur Tagmuṭ",
			territory.TM: "Turkmanistan",
			territory.TN: "Tunes",
			territory.TO: "Ṭunga",
			territory.TR: "Turkya",
			territory.TT: "Trinidad d Ṭubagu",
			territory.TV: "Ṭuvalu",
			territory.TW: "Ṭaywan",
			territory.TZ: "Ṭanzanya",
			territory.UA: "Ukranya",
			territory.UG: "Uɣanda",
			territory.US: "Iwunak Idduklen n Amirika",
			territory.UY: "Urugway",
			territory.UZ: "Uzbakistan",
			territory.VA: "Awank iɣrem n Vatikan",
			territory.VC: "Santevinsent d Grinadin",
			territory.VE: "Vinzwilla",
			territory.VG: "Tigzirin (Virgin) Tibṛiṭaniyin",
			territory.VI: "Tigzirin n Virjin n Iwunak Yedduklen",
			territory.VN: "Viṭnam",
			territory.VU: "Vanwatu",
			territory.WF: "Walis d Futuna",
			territory.WS: "Samwa",
			territory.YE: "Yaman",
			territory.YT: "Mayuṭ",
			territory.ZA: "Tafrikt Tunẓul",
			territory.ZM: "Zambya",
			territory.ZW: "Zimbabwi",
		},
	}
}