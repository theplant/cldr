// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_yo_BJ() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "yo_BJ",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMM y", Long: "d MMM y", Medium: "d MM y", Short: "d/M/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:m:s", Short: "H:m"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Shɛ́", Feb: "Èr", Mar: "Ɛr", Apr: "Ìg", May: "Ɛ̀b", Jun: "Òk", Jul: "Ag", Aug: "Òg", Sep: "Ow", Oct: "Ɔ̀w", Nov: "Bé", Dec: "Ɔ̀p"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "S", Feb: "È", Mar: "Ɛ", Apr: "Ì", May: "Ɛ̀", Jun: "Ò", Jul: "A", Aug: "Ò", Sep: "O", Oct: "Ɔ̀", Nov: "B", Dec: "Ɔ̀"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Shɛ́rɛ́", Feb: "Èrèlè", Mar: "Ɛrɛ̀nà", Apr: "Ìgbé", May: "Ɛ̀bibi", Jun: "Òkúdu", Jul: "Agɛmɔ", Aug: "Ògún", Sep: "Owewe", Oct: "Ɔ̀wàrà", Nov: "Bélú", Dec: "Ɔ̀pɛ̀"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Àìk", Mon: "Aj", Tue: "Ìsɛ́g", Wed: "Ɔjɔ́r", Thu: "Ɔjɔ́b", Fri: "Ɛt", Sat: "Àbám"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "À", Mon: "A", Tue: "Ì", Wed: "Ɔ", Thu: "Ɔ", Fri: "Ɛ", Sat: "À"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Àìkú", Mon: "Ajé", Tue: "Ìsɛ́gun", Wed: "Ɔjɔ́rú", Thu: "Ɔjɔ́bɔ", Fri: "Ɛtì", Sat: "Àbámɛ́ta"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Àìkú", Mon: "Ajé", Tue: "Ìsɛ́gun", Wed: "Ɔjɔ́rú", Thu: "Ɔjɔ́bɔ", Fri: "Ɛtì", Sat: "Àbámɛ́ta"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "Àárɔ̀", PM: "Ɔ̀sán"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "Àárɔ̀", PM: "Ɔ̀sán"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "Àárɔ̀", PM: "Ɔ̀sán"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "0 ɛgbɛ̀rún", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Diami ti Awon Orílɛ́ède Arabu", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Afugánì Afuganísítàànì", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "Lɛ́kɛ̀ Àlìbéníà", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "Dírààmù Àmɛ́níà", Symbol: ""},
				currency.ANG: cldr.Currency{DisplayName: "Gílídà Netherlands Antillean", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Wansa ti Orílɛ́ède Àngólà", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "Pɛ́sò Agɛntínà", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Dɔla ti Orílɛ́ède Ástràlìá", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Fuloríìnì Àrúbà", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Mánààtì Àsàbáíjáì", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Àmi Yíyípadà Bosnia-Herzegovina", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "Dɔ́là Bábádɔ̀ɔ̀sì", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "Tákà Báńgíládɛ̀ɛ̀shì", Symbol: "৳"},
				currency.BGN: cldr.Currency{DisplayName: "Owó Lɛ́fì Bɔ̀lìgéríà", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Dina ti Orílɛ́ède Báránì", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Faransi ti Orílɛ́ède Bùùrúndì", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "Dɔ́là Bɛ̀múdà", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "Dɔ́là Bùrùnéì", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "Bɔlifiánò Bɔ̀lífíà", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "Owó ti Orílɛ̀-èdè Brazil", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Dɔ́là Bàhámà", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "Ìngɔ́tírɔ̀mù Bútàànì", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Pula ti Orílɛ́ède Bɔ̀tìsúwánà", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "Rɔ́bù Bɛ̀lárùùsì", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "Dɔ́là Bɛ̀lísè", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Dɔla ti Orílɛ́ède Kánádà", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Faransi ti Orílɛ́ède Kóngò", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Faransi ti Orílɛ́ède Siwisi", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Pɛ́sò Shílè", Symbol: "$"},
				currency.CNH: cldr.Currency{DisplayName: "Yúànì Sháínà", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "Reminibi ti Orílɛ́ède sháínà", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Pɛ́sò Kòlóḿbíà", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "Kólɔ́ɔ̀nì Kosita Ríkà", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "Pɛ́sò Yíyípadà Kúbà", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "Pɛ́sò Kúbà", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Kabofediano ti Orílɛ́ède Esuodo", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Koruna Shɛ́ɛ̀kì", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Faransi ti Orílɛ́ède Dibouti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "Kírónì Dáníshì", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "Pɛ́sò Dòníníkà", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dina ti Orílɛ́ède Àlùgèríánì", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "pɔɔn ti Orílɛ́ède Egipiti", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Nakifa ti Orílɛ́ède Eriteriani", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Biri ti Orílɛ́ède Eutopia", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "owó Yúrò", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Dɔ́là Fíjì", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "Pɔ́n-ùn Erékùsù Falkland", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Pɔ́n-ùn ti Orilɛ̀-èdè Gɛ̀ɛ́sì", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Lárì Jɔ́jíà", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "shidi ti Orílɛ́ède Gana", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Sídì Ghanaian", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "Pɔ́n-ùn Gibraltar", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi ti Orílɛ́ède Gamibia", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "Fírànkì Gíníànì", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Faransi ti Orílɛ́ède Gini", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "Kúɛ́tísààlì Guatimílà", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "Dɔ́là Gùyánà", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "Dɔ́là Hong Kong", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Lɛmipírà Ɔ́ńdúrà", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "Kúnà Croatian", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "Gɔ́dì Àítì", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "Fɔ́ríǹtì Hɔ̀ngérí", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "Rúpìyá Indonésíà", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "Shékélì Tuntun Ísírɛ̀ɛ̀lì", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupi ti Orílɛ́ède Indina", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Dínárì Ìráákì", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "Rial Iranian", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Kòrónà Icelandic", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "Dɔ́là Jàmáíkà", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "Dínárì Jɔ́dàànì", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "Yeni ti Orílɛ́ède Japani", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "shiili ti Orílɛ́ède Kenya", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "Sómú Kirijísítàànì", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "Ráyò Kàm̀bɔ́díà", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Faransi ti Orílɛ́ède shomoriani", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "Wɔ́ɔ̀nù Àríwá Kòríà", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "Wɔ́ɔ̀nù Gúúsù Kòríà", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Dínárì Kuwaiti", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "Dɔ́là Erékùsù Cayman", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "Tɛngé Kasakísítàànì", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "Kíììpù Làótì", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "Pɔ́n-ùn Lebanese", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "Rúpìì Siri Láńkà", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Dɔla ti Orílɛ́ède Liberia", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Loti ti Orílɛ́ède Lesoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Dina ti Orílɛ́ède Libiya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Dirami ti Orílɛ́ède Moroko", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Owó Léhù Moldovan", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Faransi ti Orílɛ́ède Malagasi", Symbol: "Ar"},
				currency.MKD: cldr.Currency{DisplayName: "Dɛ́nà Masidóníà", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Kíyàtì Myanmar", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "Túgúrììkì Mòǹgólíà", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "Pàtákà Màkáò", Symbol: ""},
				currency.MRO: cldr.Currency{DisplayName: "Ouguiya ti Orílɛ́ède Maritania (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Ouguiya ti Orílɛ́ède Maritania", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupi ti Orílɛ́ède Maritiusi", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "Rúfìyá Mɔ̀lìdífà", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "Kasha ti Orílɛ́ède Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "Pɛ́sò Mɛ́síkò", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Ríngìtì Màléshíà", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Metika ti Orílɛ́ède Mosamibiki", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Mɛ́tíkààlì Mòsáḿbíìkì", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Dɔla ti Orílɛ́ède Namibia", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Náìrà ti Orílɛ̀-èdè Nàìjíríà", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "Kɔ̀dóbà Naikarágúà", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "Kírónì Nɔ́ɔ́wè", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "Rúpìì Nɛ̵́pààlì", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "Dɔ́là New Zealand", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Ráyò Omani", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "Bálíbóà Pànámà", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Sólì Pèrúù", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Kínà Papua Guinea Tuntun", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "Písò Fílípìnì", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "Rúpìì Pakisitánì", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "Sílɔ̀tì Pɔ́líshì", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "Gúáránì Párágúwè", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "Ráyò Kàtárì", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Léhù Ròméníà", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "Dínárì Sàbíà", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "Owó ruble ti ilɛ̀ Rɔ́shíà", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Faransi ti Orílɛ́ède Ruwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Riya ti Orílɛ́ède Saudi", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "Dɔ́là Erékùsù Sɔ́lómɔ́nì", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Rupi ti Orílɛ́ède Sayiselesi", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Dina ti Orílɛ́ède Sudani", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Pɔɔun ti Orílɛ́ède Sudani", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Kòrónà Súwídìn", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "Dɔ́là Síngápɔ̀", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Pɔɔun ti Orílɛ́ède ̣Elena", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Lioni", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Sile ti Orílɛ́ède Somali", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "Dɔ́là Súrínámì", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "Pɔ́n-ùn Gúúsù Sùdáànì", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Dobira ti Orílɛ́ède Sao tome Ati Pirisipe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Dobira ti Orílɛ́ède Sao tome Ati Pirisipe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "Pɔ́n-ùn Sírìà", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "Báàtì Tháì", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "Sómónì Tajikístàànì", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Mánààtì Tɔkimɛnístàànì", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "Dina ti Orílɛ́ède Tunisia", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "Pàángà Tóńgà", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "Lírà Tɔ́kì", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "Dɔ́là Trinidad & Tobago", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "Dɔ́là Tàìwánì Tuntun", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Sile ti Orílɛ́ède Tansania", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "Ɔrifiníyà Yukiréníà", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Siile ti Orílɛ́ède Uganda", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Dɔ́là", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "Pɛ́sò Úrúgúwè", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "Sómú Usibɛkísítàànì", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "Bɔ̀lífà Fɛnɛsuɛ́là", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "Dáhùn Vietnamese", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Fátù Vanuatu", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "Tálà Sàmóà", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "Faransi ti Orílɛ́ède BEKA", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Dɔ́là Ilà Oòrùn Karíbíà", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Faransi ti Orílɛ́ède BIKEAO", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "Fírànkì CFP", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "owóníná àìmɔ̀", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Ráyò Yɛ́mɛ̀nì", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Randi ti Orílɛ́ède Ariwa Afirika", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Kawasha ti Orílɛ́ède Saabia (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kawasha ti Orílɛ́ède Saabia", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Dɔla ti Orílɛ́ède Siibabuwe", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AF:      "Èdè Afrikani",
			language.AGQ:     "Ágɛ̀ɛ̀mù",
			language.AK:      "Èdè Akani",
			language.AM:      "Èdè Amariki",
			language.AR:      "Èdè Árábìkì",
			language.AS:      "Ti Assam",
			language.ASA:     "Asu",
			language.AST:     "Asturian",
			language.AZ:      "Èdè Azerbaijani",
			language.BAS:     "Basaa",
			language.BE:      "Èdè Belarusi",
			language.BEM:     "Béḿbà",
			language.BEZ:     "Bɛ́nà",
			language.BG:      "Èdè Bugaria",
			language.BM:      "Báḿbàrà",
			language.BN:      "Èdè Bengali",
			language.BO:      "Tibetán",
			language.BR:      "Èdè Bretoni",
			language.BRX:     "Bódò",
			language.BS:      "Èdè Bosnia",
			language.CA:      "Èdè Catala",
			language.CCP:     "Chakma",
			language.CE:      "Chechen",
			language.CEB:     "Cebuano",
			language.CGG:     "Chiga",
			language.CHR:     "Shɛ́rókiì",
			language.CKB:     "Ààrin Gbùngbùn Kurdish",
			language.CO:      "Corsican",
			language.CS:      "Èdè seeki",
			language.CU:      "Síláfííkì Ilé Ìjɔ́sìn",
			language.CY:      "Èdè Welshi",
			language.DA:      "Èdè Ilɛ̀ Denmark",
			language.DAV:     "Táítà",
			language.DE:      "Èdè Jámánì",
			language.DE_AT:   "Èdè Jámánì (Ɔ́síríà )",
			language.DE_CH:   "Èdè Ilɛ̀ Jámánì (Orílɛ́ède swítsàlandì)",
			language.DJE:     "Shárúmà",
			language.DSB:     "Shobíànù Ìpìlɛ̀",
			language.DUA:     "Duala",
			language.DYO:     "Jola-Fonyi",
			language.DZ:      "Dzongkha",
			language.EBU:     "Ɛmbù",
			language.EE:      "Ewè",
			language.EL:      "Èdè Giriki",
			language.EN:      "Èdè Gɛ̀ɛ́sì",
			language.EN_AU:   "Èdè Gɛ̀ɛ́sì (órílɛ̀-èdè Ɔsirélíà)",
			language.EN_CA:   "Èdè Gɛ̀ɛ́sì (Orílɛ̀-èdè Kánádà)",
			language.EN_GB:   "Èdè Gɛ̀ɛ́sì (GB)",
			language.EN_US:   "Èdè Gɛ̀ɛ́sì (US)",
			language.EO:      "Èdè Esperanto",
			language.ES:      "Èdè Sípáníìshì",
			language.ES_419:  "Èdè Sípáníìshì (orílɛ̀-èdè Látìn-Amɛ́ríkà) ( Èdè Sípáníìshì (Látìn-Amɛ́ríkà)",
			language.ES_ES:   "Èdè Sípáníìshì (orílɛ̀-èdè Yúróòpù)",
			language.ES_MX:   "Èdè Sípáníìshì (orílɛ̀-èdè Mɛ́síkò)",
			language.ET:      "Èdè Estonia",
			language.EU:      "Èdè Baski",
			language.EWO:     "Èwóǹdò",
			language.FA:      "Èdè Pasia",
			language.FF:      "Èdè Fúlàní",
			language.FI:      "Èdè Finisi",
			language.FIL:     "Èdè Filipino",
			language.FO:      "Èdè Faroesi",
			language.FR:      "Èdè Faransé",
			language.FR_CA:   "Èdè Faransé (orílɛ̀-èdè Kánádà)",
			language.FR_CH:   "Èdè Faranshé (Súwísàlaǹdì)",
			language.FUR:     "Firiúlíànì",
			language.FY:      "Èdè Frisia",
			language.GA:      "Èdè Ireland",
			language.GD:      "Èdè Gaelik ti Ilu Scotland",
			language.GL:      "Èdè Galicia",
			language.GN:      "Èdè Guarani",
			language.GSW:     "Súwísì ti Jámánì",
			language.GU:      "Èdè Gujarati",
			language.GUZ:     "Gusii",
			language.GV:      "Máǹkì",
			language.HA:      "Èdè Hausa",
			language.HAW:     "Hawaiian",
			language.HE:      "Èdè Heberu",
			language.HI:      "Èdè Híńdì",
			language.HMN:     "Hmong",
			language.HR:      "Èdè Kroatia",
			language.HSB:     "Sorbian Apá Òkè",
			language.HT:      "Haitian Creole",
			language.HU:      "Èdè Hungaria",
			language.HY:      "Èdè Ile Armenia",
			language.IA:      "Èdè pipo",
			language.ID:      "Èdè Indonéshíà",
			language.IE:      "Iru Èdè",
			language.IG:      "Èdè Yíbò",
			language.II:      "Shíkuán Yì",
			language.IS:      "Èdè Icelandic",
			language.IT:      "Èdè Ítálì",
			language.JA:      "Èdè Jàpáànù",
			language.JGO:     "Ńgòmbà",
			language.JMC:     "Máshámè",
			language.JV:      "Èdè Javanasi",
			language.KA:      "Èdè Georgia",
			language.KAB:     "Kabilè",
			language.KAM:     "Káńbà",
			language.KDE:     "Mákondé",
			language.KEA:     "Kabufadíánù",
			language.KHQ:     "Koira Shíínì",
			language.KI:      "Kíkúyù",
			language.KK:      "Kashakì",
			language.KKJ:     "Kàkó",
			language.KL:      "Kalaalísùtì",
			language.KLN:     "Kálɛnjín",
			language.KM:      "Èdè kameri",
			language.KN:      "Èdè Kannada",
			language.KO:      "Èdè Kòríà",
			language.KOK:     "Kónkánì",
			language.KS:      "Kashímirì",
			language.KSB:     "Sháńbálà",
			language.KSF:     "Báfíà",
			language.KSH:     "Colognian",
			language.KU:      "Kɔdishì",
			language.KW:      "Kɔ́nììshì",
			language.KY:      "Kírígíìsì",
			language.LA:      "Èdè Latini",
			language.LAG:     "Láńgì",
			language.LB:      "Lùshɛ́mbɔ́ɔ̀gì",
			language.LG:      "Ganda",
			language.LKT:     "Lákota",
			language.LN:      "Lìǹgálà",
			language.LO:      "Láò",
			language.LRC:     "Apáàríwá Lúrì",
			language.LT:      "Èdè Lithuania",
			language.LU:      "Lúbà-Katanga",
			language.LUY:     "Luyíà",
			language.LV:      "Èdè Latvianu",
			language.MAS:     "Másáì",
			language.MER:     "Mérù",
			language.MFE:     "Morisiyen",
			language.MG:      "Malagasì",
			language.MGH:     "Makhuwa-Meeto",
			language.MGO:     "Métà",
			language.MI:      "Màórì",
			language.MK:      "Èdè Macedonia",
			language.ML:      "Málàyálámù",
			language.MN:      "Mòngólíà",
			language.MR:      "Èdè marathi",
			language.MS:      "Èdè Malaya",
			language.MT:      "Èdè Malta",
			language.MUA:     "Múndàngì",
			language.MUL:     "Ɔlɔ́pɔ̀ èdè",
			language.MY:      "Èdè Bumiisi",
			language.MZN:     "Masanderani",
			language.NAQ:     "Námà",
			language.NB:      "Nɔ́ɔ́wè Bokímàl",
			language.ND:      "Àríwá Ndebele",
			language.NDS:     "Jámánì ìpìlɛ̀",
			language.NE:      "Èdè Nepali",
			language.NL:      "Èdè Dɔ́ɔ̀shì",
			language.NMG:     "Kíwáshíò",
			language.NN:      "Nɔ́ɔ́wè Nínɔ̀sìkì",
			language.NNH:     "Ngiembùnù",
			language.NO:      "Èdè Norway",
			language.NUS:     "Núɛ̀",
			language.NY:      "Ńyájà",
			language.NYN:     "Ńyákɔ́lè",
			language.OC:      "Èdè Occitani",
			language.OM:      "Òròmɔ́",
			language.OR:      "Òdíà",
			language.OS:      "Ɔshɛ́tíìkì",
			language.PA:      "Èdè Punjabi",
			language.PL:      "Èdè Póláǹdì",
			language.PRG:     "Púrúshíànù",
			language.PS:      "Páshítò",
			language.PT:      "Èdè Pɔtogí",
			language.PT_BR:   "Èdè Pɔtogí (Orilɛ̀-èdè Bràsíl)",
			language.PT_PT:   "Èdè Pɔtogí (orílɛ̀-èdè Yúróòpù)",
			language.QU:      "Kúɛ́ńjùà",
			language.RM:      "Rómáǹshì",
			language.RN:      "Rúńdì",
			language.RO:      "Èdè Romania",
			language.ROF:     "Róńbò",
			language.RU:      "Èdè Rɔ́shíà",
			language.RW:      "Èdè Ruwanda",
			language.RWK:     "Riwa",
			language.SA:      "Èdè awon ara Indo",
			language.SAH:     "Sàkíhà",
			language.SAQ:     "Samburu",
			language.SBP:     "Sangu",
			language.SD:      "Èdè Sindhi",
			language.SE:      "Apáàríwá Sami",
			language.SEH:     "Shɛnà",
			language.SES:     "Koiraboro Seni",
			language.SG:      "Sango",
			language.SH:      "Èdè Serbo-Croatiani",
			language.SHI:     "Tashelíìtì",
			language.SI:      "Èdè Sinhalese",
			language.SK:      "Èdè Slovaki",
			language.SL:      "Èdè Slovenia",
			language.SM:      "Sámóánù",
			language.SMN:     "Inari Sami",
			language.SN:      "Shɔnà",
			language.SO:      "Èdè ara Somalia",
			language.SQ:      "Èdè Albania",
			language.SR:      "Èdè Serbia",
			language.ST:      "Èdè Sesoto",
			language.SU:      "Èdè Sudani",
			language.SV:      "Èdè Suwidiisi",
			language.SW:      "Èdè Swahili",
			language.TA:      "Èdè Tamili",
			language.TE:      "Èdè Telugu",
			language.TEO:     "Tɛ́sò",
			language.TG:      "Tàjíìkì",
			language.TH:      "Èdè Tai",
			language.TI:      "Èdè Tigrinya",
			language.TK:      "Èdè Turkmen",
			language.TLH:     "Èdè Klingoni",
			language.TO:      "Tóńgàn",
			language.TR:      "Èdè Tɔɔkisi",
			language.TT:      "Tatarí",
			language.TWQ:     "Tasawak",
			language.TZM:     "Ààrin Gbùngbùn Atlas Tamazight",
			language.UG:      "Yúgɔ̀",
			language.UK:      "Èdè Ukania",
			language.UND:     "Èdè àìmɔ̀",
			language.UR:      "Èdè Udu",
			language.UZ:      "Èdè Uzbek",
			language.VI:      "Èdè Jetinamu",
			language.VO:      "Fɔ́lápùùkù",
			language.VUN:     "Funjo",
			language.WAE:     "Wɔsà",
			language.WO:      "Wɔ́lɔ́ɔ̀fù",
			language.XH:      "Èdè Xhosa",
			language.XOG:     "Shógà",
			language.YAV:     "Yangbɛn",
			language.YI:      "Èdè Yiddishi",
			language.YO:      "Èdè Yorùbá",
			language.YUE:     "Cantonese",
			language.ZGH:     "Àfɛnùkò Támásáìtì ti Mòrókò",
			language.ZH:      "Sháídà, Mandrínì",
			language.ZH_HANT: "Èdè Ìbílɛ̀ Sháínà",
			language.ZU:      "Èdè Shulu",
			language.ZXX:     "Kò sí àkóònú elédè",
		},
		Territories: cldr.Territories{
			territory.V_001: "Agbáyé",
			territory.V_002: "Áfíríkà",
			territory.V_003: "Àríwá Amɛ́ríkà",
			territory.V_005: "Gúúshù Amɛ́ríkà",
			territory.V_009: "Oceania",
			territory.V_011: "Ìwɔ̀ oorùn Afíríkà",
			territory.V_013: "Ààrin Gbùgbùn Àmɛ́ríkà",
			territory.V_014: "Ìlà Oorùn Áfíríkà",
			territory.V_015: "Northern Africa",
			territory.V_017: "Ààrín gbùngbùn Afíríkà",
			territory.V_018: "Apágúúsù Áfíríkà",
			territory.V_019: "Amɛ́ríkà",
			territory.V_021: "Apáàríwá Amɛ́ríkà",
			territory.V_029: "Káríbíànù",
			territory.V_030: "Ìlà Òòrùn Eshíà",
			territory.V_034: "Gúúshù Eshíà",
			territory.V_035: "Gúúshù ìlà òòrùn Éshíà",
			territory.V_039: "Gúúshù Yúróòpù",
			territory.V_053: "Ɔshirélashíà",
			territory.V_054: "Mɛlanéshíà",
			territory.V_057: "Agbègbè Maikironéshíà",
			territory.V_061: "Polineshíà",
			territory.V_142: "Asia",
			territory.V_143: "Ààrin Gbùngbùn Éshíà",
			territory.V_145: "Ìwɔ̀ Òòrùn Eshíà",
			territory.V_150: "Europe",
			territory.V_151: "Ìlà Òrùn Yúrópù",
			territory.V_154: "Northern Europe",
			territory.V_155: "Ìwɔ̀ Òòrùn Yúrópù",
			territory.V_202: "Sàhárà Áfíríkà",
			territory.V_419: "Látín Amɛ́ríkà",
			territory.AC:    "Erékùsù Ascension",
			territory.AD:    "Orílɛ́ède Ààndórà",
			territory.AE:    "Orílɛ́ède Ɛmirate ti Awɔn Arabu",
			territory.AF:    "Orílɛ́ède Àfùgànístánì",
			territory.AG:    "Orílɛ́ède Ààntígúà àti Báríbúdà",
			territory.AI:    "Orílɛ́ède Ààngúlílà",
			territory.AL:    "Orílɛ́ède Àlùbàníánì",
			territory.AM:    "Orílɛ́ède Améníà",
			territory.AO:    "Orílɛ́ède Ààngólà",
			territory.AQ:    "Antakítíkà",
			territory.AR:    "Orílɛ́ède Agentínà",
			territory.AS:    "Sámóánì ti Orílɛ́ède Àméríkà",
			territory.AT:    "Orílɛ́ède Asítíríà",
			territory.AU:    "Orílɛ́ède Ástràlìá",
			territory.AW:    "Orílɛ́ède Árúbà",
			territory.AX:    "Àwɔn Erékùsù ti Åland",
			territory.AZ:    "Orílɛ́ède Asɛ́bájánì",
			territory.BA:    "Orílɛ́ède Bɔ̀síníà àti Ɛtisɛgófínà",
			territory.BB:    "Orílɛ́ède Bábádósì",
			territory.BD:    "Orílɛ́ède Bángáládésì",
			territory.BE:    "Orílɛ́ède Bégíɔ́mù",
			territory.BF:    "Orílɛ́ède Bùùkíná Fasò",
			territory.BG:    "Orílɛ́ède Bùùgáríà",
			territory.BH:    "Orílɛ́ède Báránì",
			territory.BI:    "Orílɛ́ède Bùùrúndì",
			territory.BJ:    "Orílɛ́ède Bɛ̀nɛ̀",
			territory.BL:    "St. Barthélemy",
			territory.BM:    "Orílɛ́ède Bémúdà",
			territory.BN:    "Orílɛ́ède Búrúnɛ́lì",
			territory.BO:    "Orílɛ́ède Bɔ̀lífíyà",
			territory.BQ:    "Caribbean Netherlands",
			territory.BR:    "Orilɛ̀-èdè Bàràsílì",
			territory.BS:    "Orílɛ́ède Bàhámásì",
			territory.BT:    "Orílɛ́ède Bútánì",
			territory.BV:    "Erékùsù Bouvet",
			territory.BW:    "Orílɛ́ède Bɔ̀tìsúwánà",
			territory.BY:    "Orílɛ́ède Bélárúsì",
			territory.BZ:    "Orílɛ́ède Bèlísɛ̀",
			territory.CA:    "Orílɛ́ède Kánádà",
			territory.CC:    "Erékùsù Cocos (Keeling)",
			territory.CD:    "Orilɛ́ède Kóngò",
			territory.CF:    "Orílɛ́ède Àrin gùngun Áfíríkà",
			territory.CG:    "Orílɛ́ède Kóngò",
			territory.CH:    "Orílɛ́ède switishilandi",
			territory.CI:    "Orílɛ́ède Kóútè forà",
			territory.CK:    "Orílɛ́ède Etíokun Kùúkù",
			territory.CL:    "Orílɛ́ède shílè",
			territory.CM:    "Orílɛ́ède Kamerúúnì",
			territory.CN:    "Orilɛ̀-èdè Sháínà",
			territory.CO:    "Orílɛ́ède Kòlómíbìa",
			territory.CP:    "Erékùsù Clipperston",
			territory.CR:    "Orílɛ́ède Kuusita Ríkà",
			territory.CU:    "Orílɛ́ède Kúbà",
			territory.CV:    "Orílɛ́ède Etíokun Kápé féndè",
			territory.CW:    "Curaçao",
			territory.CX:    "Erékùsù Christmas",
			territory.CY:    "Orílɛ́ède Kúrúsì",
			territory.CZ:    "Orílɛ́ède shɛ́ɛ́kì",
			territory.DE:    "Orílɛèdè Jámánì",
			territory.DG:    "Diego Gashia",
			territory.DJ:    "Orílɛ́ède Díbɔ́ótì",
			territory.DK:    "Orílɛ́ède Dɛ́mákì",
			territory.DM:    "Orílɛ́ède Dòmíníkà",
			territory.DO:    "Orilɛ́ède Dòmíníkánì",
			territory.DZ:    "Orílɛ́ède Àlùgèríánì",
			territory.EA:    "Seuta àti Melilla",
			territory.EC:    "Orílɛ́ède Ekuádò",
			territory.EE:    "Orílɛ́ède Esitonia",
			territory.EG:    "Orílɛ́ède Égípítì",
			territory.EH:    "Ìwɔ̀òòrùn Sàhárà",
			territory.ER:    "Orílɛ́ède Eritira",
			territory.ES:    "Orílɛ́ède Sipani",
			territory.ET:    "Orílɛ́ède Etopia",
			territory.EU:    "Ìshɔ̀kan Yúròpù",
			territory.EZ:    "Agbègbè Euro",
			territory.FI:    "Orílɛ́ède Filandi",
			territory.FJ:    "Orílɛ́ède Fiji",
			territory.FK:    "Orílɛ́ède Etikun Fakalandi",
			territory.FM:    "Orílɛ́ède Makoronesia",
			territory.FO:    "Àwɔn Erékùsù ti Faroe",
			territory.FR:    "Orílɛ́ède Faranse",
			territory.GA:    "Orílɛ́ède Gabon",
			territory.GB:    "Orílɛ́èdè Gɛ̀ɛ́sì",
			territory.GD:    "Orílɛ́ède Genada",
			territory.GE:    "Orílɛ́ède Gɔgia",
			territory.GF:    "Orílɛ́ède Firenshi Guana",
			territory.GG:    "Guernsey",
			territory.GH:    "Orílɛ́ède Gana",
			territory.GI:    "Orílɛ́ède Gibaratara",
			territory.GL:    "Orílɛ́ède Gerelandi",
			territory.GM:    "Orílɛ́ède Gambia",
			territory.GN:    "Orílɛ́ède Gene",
			territory.GP:    "Orílɛ́ède Gadelope",
			territory.GQ:    "Orílɛ́ède Ekutoria Gini",
			territory.GR:    "Orílɛ́ède Geriisi",
			territory.GS:    "Gúúsù Georgia àti Gúúsù Àwɔn Erékùsù Sandwich",
			territory.GT:    "Orílɛ́ède Guatemala",
			territory.GU:    "Orílɛ́ède Guamu",
			territory.GW:    "Orílɛ́ède Gene-Busau",
			territory.GY:    "Orílɛ́ède Guyana",
			territory.HK:    "Hong Kong SAR ti Sháìnà",
			territory.HM:    "Erékùsù Heard àti Erékùsù McDonald",
			territory.HN:    "Orílɛ́ède Hondurasi",
			territory.HR:    "Orílɛ́ède Kòróátíà",
			territory.HT:    "Orílɛ́ède Haati",
			territory.HU:    "Orílɛ́ède Hungari",
			territory.IC:    "Ɛrékùsù Kánárì",
			territory.ID:    "Orílɛ́ède Indonesia",
			territory.IE:    "Orílɛ́ède Ailandi",
			territory.IL:    "Orílɛ́ède Iserɛli",
			territory.IM:    "Isle of Man",
			territory.IN:    "Orílɛ́ède India",
			territory.IO:    "Orílɛ́ède Etíkun Índíánì ti Ìlú Bírítísì",
			territory.IQ:    "Orílɛ́ède Iraki",
			territory.IR:    "Orílɛ́ède Irani",
			territory.IS:    "Orílɛ́ède Ashilandi",
			territory.IT:    "Orílɛ́ède Itáli",
			territory.JE:    "Jersey",
			territory.JM:    "Orílɛ́ède Jamaika",
			territory.JO:    "Orílɛ́ède Jɔdani",
			territory.JP:    "Orílɛ́ède Japani",
			territory.KE:    "Orílɛ́ède Kenya",
			territory.KG:    "Orílɛ́ède Kurishisitani",
			territory.KH:    "Orílɛ́ède Kàmùbódíà",
			territory.KI:    "Orílɛ́ède Kiribati",
			territory.KM:    "Orílɛ́ède Kòmòrósì",
			territory.KN:    "Orílɛ́ède Kiiti ati Neefi",
			territory.KP:    "Orílɛ́ède Guusu Kɔria",
			territory.KR:    "Orílɛ́ède Ariwa Kɔria",
			territory.KW:    "Orílɛ́ède Kuweti",
			territory.KY:    "Orílɛ́ède Etíokun Kámánì",
			territory.KZ:    "Orílɛ́ède Kashashatani",
			territory.LA:    "Orílɛ́ède Laosi",
			territory.LB:    "Orílɛ́ède Lebanoni",
			territory.LC:    "Orílɛ́ède Lushia",
			territory.LI:    "Orílɛ́ède Lɛshitɛnisiteni",
			territory.LK:    "Orílɛ́ède Siri Lanka",
			territory.LR:    "Orílɛ́ède Laberia",
			territory.LS:    "Orílɛ́ède Lesoto",
			territory.LT:    "Orílɛ́ède Lituania",
			territory.LU:    "Orílɛ́ède Lusemogi",
			territory.LV:    "Orílɛ́ède Latifia",
			territory.LY:    "Orílɛ́ède Libiya",
			territory.MA:    "Orílɛ́ède Moroko",
			territory.MC:    "Orílɛ́ède Monako",
			territory.MD:    "Orílɛ́ède Modofia",
			territory.ME:    "Montenegro",
			territory.MF:    "St. Martin",
			territory.MG:    "Orílɛ́ède Madasika",
			territory.MH:    "Orílɛ́ède Etikun Máshali",
			territory.MK:    "Àríwá Macedonia",
			territory.ML:    "Orílɛ́ède Mali",
			territory.MM:    "Orílɛ́ède Manamari",
			territory.MN:    "Orílɛ́ède Mogolia",
			territory.MO:    "Macao SAR ti Sháìnà",
			territory.MP:    "Orílɛ́ède Etikun Guusu Mariana",
			territory.MQ:    "Orílɛ́ède Matinikuwi",
			territory.MR:    "Orílɛ́ède Maritania",
			territory.MS:    "Orílɛ́ède Motserati",
			territory.MT:    "Orílɛ́ède Malata",
			territory.MU:    "Orílɛ́ède Maritiusi",
			territory.MV:    "Orílɛ́ède Maladifi",
			territory.MW:    "Orílɛ́ède Malawi",
			territory.MX:    "Orílɛ́ède Mesiko",
			territory.MY:    "Orílɛ́ède Malasia",
			territory.MZ:    "Orílɛ́ède Moshamibiku",
			territory.NA:    "Orílɛ́ède Namibia",
			territory.NC:    "Orílɛ́ède Kaledonia Titun",
			territory.NE:    "Orílɛ́ède Nàìjá",
			territory.NF:    "Orílɛ́ède Etikun Nɔ́úfókì",
			territory.NG:    "Orilɛ̀-èdè Nàìjíríà",
			territory.NI:    "Orílɛ́ède NIkaragua",
			territory.NL:    "Orílɛ́ède Nedalandi",
			territory.NO:    "Orílɛ́ède Nɔɔwii",
			territory.NP:    "Orílɛ́ède Nepa",
			territory.NR:    "Orílɛ́ède Nauru",
			territory.NU:    "Orílɛ́ède Niue",
			territory.NZ:    "Orílɛ́ède shilandi Titun",
			territory.OM:    "Orílɛ́ède Ɔɔma",
			territory.PA:    "Orílɛ́ède Panama",
			territory.PE:    "Orílɛ́ède Peru",
			territory.PF:    "Orílɛ́ède Firenshi Polinesia",
			territory.PG:    "Orílɛ́ède Paapu ti Giini",
			territory.PH:    "Orílɛ́ède filipini",
			territory.PK:    "Orílɛ́ède Pakisitan",
			territory.PL:    "Orílɛ́ède Polandi",
			territory.PM:    "Orílɛ́ède Pɛɛri ati mikuloni",
			territory.PN:    "Orílɛ́ède Pikarini",
			territory.PR:    "Orílɛ́ède Pɔto Riko",
			territory.PS:    "Palɛsitín",
			territory.PT:    "Orílɛ́ède Pɔ́túgà",
			territory.PW:    "Orílɛ́ède Paalu",
			territory.PY:    "Orílɛ́ède Paraguye",
			territory.QA:    "Orílɛ́ède Kota",
			territory.QO:    "Agbègbè Oceania",
			territory.RE:    "Orílɛ́ède Riuniyan",
			territory.RO:    "Orílɛ́ède Romaniya",
			territory.RS:    "Serbia",
			territory.RU:    "Orílɛ́ède Rɔshia",
			territory.RW:    "Orílɛ́ède Ruwanda",
			territory.SA:    "Orílɛ́ède Saudi Arabia",
			territory.SB:    "Orílɛ́ède Etikun Solomoni",
			territory.SC:    "Orílɛ́ède seshɛlɛsi",
			territory.SD:    "Orílɛ́ède Sudani",
			territory.SE:    "Orílɛ́ède Swidini",
			territory.SG:    "Orílɛ́ède Singapo",
			territory.SH:    "Orílɛ́ède Hɛlena",
			territory.SI:    "Orílɛ́ède Silofania",
			territory.SJ:    "Svalbard & Jan Mayen",
			territory.SK:    "Orílɛ́ède Silofakia",
			territory.SL:    "Orílɛ́ède Siria looni",
			territory.SM:    "Orílɛ́ède Sani Marino",
			territory.SN:    "Orílɛ́ède Sɛnɛga",
			territory.SO:    "Orílɛ́ède Somalia",
			territory.SR:    "Orílɛ́ède Surinami",
			territory.SS:    "Gúúsù Sudan",
			territory.ST:    "Orílɛ́ède Sao tomi ati piriishipi",
			territory.SV:    "Orílɛ́ède Ɛɛsáfádò",
			territory.SX:    "Sint Maarten",
			territory.SY:    "Orílɛ́ède Siria",
			territory.SZ:    "Orílɛ́ède Sashiland",
			territory.TA:    "Tristan da Kunha",
			territory.TC:    "Orílɛ́ède Tɔɔki ati Etikun Kakɔsi",
			territory.TD:    "Orílɛ́ède shààdì",
			territory.TF:    "Agbègbè Gúúsù Faranshé",
			territory.TG:    "Orílɛ́ède Togo",
			territory.TH:    "Orílɛ́ède Tailandi",
			territory.TJ:    "Orílɛ́ède Takisitani",
			territory.TK:    "Orílɛ́ède Tokelau",
			territory.TL:    "Ìlà Òòrùn Tímɔ̀",
			territory.TM:    "Orílɛ́ède Tɔɔkimenisita",
			territory.TN:    "Orílɛ́ède Tunishia",
			territory.TO:    "Orílɛ́ède Tonga",
			territory.TR:    "Orílɛ́ède Tɔɔki",
			territory.TT:    "Orílɛ́ède Tirinida ati Tobaga",
			territory.TV:    "Orílɛ́ède Tufalu",
			territory.TW:    "Orílɛ́ède Taiwani",
			territory.TZ:    "Orílɛ́ède Tàǹsáníà",
			territory.UA:    "Orílɛ́ède Ukarini",
			territory.UG:    "Orílɛ́ède Uganda",
			territory.UM:    "Àwɔn Erékùsù Kékèké Agbègbè US",
			territory.UN:    "Ìshɔ̀kan àgbáyé",
			territory.US:    "Orílɛ̀-èdè Amɛrikà",
			territory.UY:    "Orílɛ́ède Nruguayi",
			territory.UZ:    "Orílɛ́ède Nshibɛkisitani",
			territory.VA:    "Ìlú Vatican",
			territory.VC:    "Orílɛ́ède Fisɛnnti ati Genadina",
			territory.VE:    "Orílɛ́ède Fɛnɛshuɛla",
			territory.VG:    "Orílɛ́ède Etíkun Fágínì ti ìlú Bírítísì",
			territory.VI:    "Orílɛ́ède Etikun Fagini ti Amɛrika",
			territory.VN:    "Orílɛ́ède Fɛtinami",
			territory.VU:    "Orílɛ́ède Faniatu",
			territory.WF:    "Orílɛ́ède Wali ati futuna",
			territory.WS:    "Orílɛ́ède Samɔ",
			territory.XA:    "Pseudo-Accents",
			territory.XB:    "Pseudo-Bidi",
			territory.XK:    "Kòsófò",
			territory.YE:    "Orílɛ́ède yemeni",
			territory.YT:    "Orílɛ́ède Mayote",
			territory.ZA:    "Gúúshù Áfíríkà",
			territory.ZM:    "Orílɛ́ède shamibia",
			territory.ZW:    "Orílɛ́ède shimibabe",
			territory.ZZ:    "Àgbègbè àìmɔ̀",
		},
	}
}