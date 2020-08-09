// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_fy() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "fy",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd-MM-yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'om' {0}", Long: "{1} 'om' {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mrt", Apr: "Apr", May: "Mai", Jun: "Jun", Jul: "Jul", Aug: "Aug", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Des"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Jannewaris", Feb: "Febrewaris", Mar: "Maart", Apr: "April", May: "Maaie", Jun: "Juny", Jul: "July", Aug: "Augustus", Sep: "Septimber", Oct: "Oktober", Nov: "Novimber", Dec: "Desimber"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "si", Mon: "mo", Tue: "ti", Wed: "wo", Thu: "to", Fri: "fr", Sat: "so"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "W", Thu: "T", Fri: "F", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "si", Mon: "mo", Tue: "ti", Wed: "wo", Thu: "to", Fri: "fr", Sat: "so"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "snein", Mon: "moandei", Tue: "tiisdei", Wed: "woansdei", Thu: "tongersdei", Fri: "freed", Sat: "sneon"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_fy]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a0#,##0.00;¤\u00a0#,##0.00-", CurrencyAccounting: "¤\u00a0#,##0.00;(¤\u00a0#,##0.00)", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "Andorrese peseta", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "Verenigde Arabyske Emiraten-dirham", Symbol: ""},
				currency.AFA: cldr.Currency{DisplayName: "Afghani (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Afghaanske afghani", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "Albanese lek", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "Armeense dram", Symbol: ""},
				currency.ANG: cldr.Currency{DisplayName: "Nederlânsk-Antilliaanske gûne", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Angolese kwanza", Symbol: "Kz"},
				currency.AOK: cldr.Currency{DisplayName: "Angolese kwanza (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "Angolese nieuwe kwanza (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "Angolese kwanza reajustado (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "Argentynske austral", Symbol: ""},
				currency.ARL: cldr.Currency{DisplayName: "Argentynske peso ley (1970–1983)", Symbol: ""},
				currency.ARM: cldr.Currency{DisplayName: "Argentynske peso (1881–1970)", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "Argentynske peso (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "Argentynske peso", Symbol: "$"},
				currency.ATS: cldr.Currency{DisplayName: "Eastenrykse schilling", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "Australyske dollar", Symbol: "AU$"},
				currency.AWG: cldr.Currency{DisplayName: "Arubaanske gulden", Symbol: ""},
				currency.AZM: cldr.Currency{DisplayName: "Azerbeidzjaanske manat (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Azerbeidzjaanske manat", Symbol: ""},
				currency.BAD: cldr.Currency{DisplayName: "Bosnyske dinar", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Bosnyske convertibele mark", Symbol: "KM"},
				currency.BAN: cldr.Currency{DisplayName: "Nije Bosnyske dinar (1994–1997)", Symbol: ""},
				currency.BBD: cldr.Currency{DisplayName: "Barbadaanske dollar", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "Bengalese taka", Symbol: "৳"},
				currency.BEC: cldr.Currency{DisplayName: "Belgyske frank (convertibel)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "Belgyske frank", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "Belgyske frank (finansjeel)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "Bulgaarse harde lev", Symbol: ""},
				currency.BGM: cldr.Currency{DisplayName: "Bulgaarse socialistyske lev", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "Bulgaarse lev", Symbol: ""},
				currency.BGO: cldr.Currency{DisplayName: "Bulgaarse lev (1879–1952)", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Bahreinse dinar", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Burundese frank", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "Bermuda-dollar", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "Bruneise dollar", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "Boliviaanske boliviano", Symbol: "Bs"},
				currency.BOL: cldr.Currency{DisplayName: "Boliviaanske boliviano (1863–1963)", Symbol: ""},
				currency.BOP: cldr.Currency{DisplayName: "Boliviaanske peso", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "Boliviaanske mvdol", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "Braziliaanske cruzeiro novo (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "Braziliaanske cruzado", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "Braziliaanske cruzeiro (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "Braziliaanske real", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "Braziliaanske cruzado novo", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "Braziliaanske cruzeiro", Symbol: ""},
				currency.BRZ: cldr.Currency{DisplayName: "Braziliaanske cruzeiro (1942–1967)", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "Bahamaanske dollar", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "Bhutaanske ngultrum", Symbol: ""},
				currency.BUK: cldr.Currency{DisplayName: "Birmese kyat", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Botswaanske pula", Symbol: "P"},
				currency.BYB: cldr.Currency{DisplayName: "Wit-Russyske nieuwe roebel (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "Wit-Russyske roebel", Symbol: "р."},
				currency.BYR: cldr.Currency{DisplayName: "Wit-Russyske roebel (2000–2016)", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "Belizaanske dollar", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Canadese dollar", Symbol: "C$"},
				currency.CDF: cldr.Currency{DisplayName: "Congolese frank", Symbol: ""},
				currency.CHE: cldr.Currency{DisplayName: "WIR euro", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Zwitserse frank", Symbol: ""},
				currency.CHW: cldr.Currency{DisplayName: "WIR franc", Symbol: ""},
				currency.CLE: cldr.Currency{DisplayName: "Sileenske escudo", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "Sileenske unidades de fomento", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Sileenske peso", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Sineeske yuan renminbi", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Kolombiaanske peso", Symbol: "$"},
				currency.COU: cldr.Currency{DisplayName: "Unidad de Valor Real", Symbol: ""},
				currency.CRC: cldr.Currency{DisplayName: "Costaricaanske colón", Symbol: "₡"},
				currency.CSD: cldr.Currency{DisplayName: "Alde Servyske dinar", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "Tsjechoslowaakse harde koruna", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "Kubaanske convertibele peso", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "Kubaanske peso", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Kaapverdyske escudo", Symbol: ""},
				currency.CYP: cldr.Currency{DisplayName: "Cyprysk pûn", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Tsjechyske kroon", Symbol: "Kč"},
				currency.DDM: cldr.Currency{DisplayName: "East-Dútske ostmark", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "Dútske mark", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "Djiboutiaanske frank", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "Deenske kroon", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "Dominikaanske peso", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Algerynske dinar", Symbol: ""},
				currency.ECS: cldr.Currency{DisplayName: "Ecuadoraanske sucre", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "Ecuadoraanske unidad de valor constante (UVC)", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "Estlânske kroon", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Egyptysk pûn", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Eritrese nakfa", Symbol: ""},
				currency.ESA: cldr.Currency{DisplayName: "Spaanske peseta (account A)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "Spaanske peseta (convertibele account)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "Spaanske peseta", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Ethiopyske birr", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Finse markka", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "Fiji-dollar", Symbol: "FJ$"},
				currency.FKP: cldr.Currency{DisplayName: "Falklâneilânske pûn", Symbol: "£"},
				currency.FRF: cldr.Currency{DisplayName: "Franske franc", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "Brits pûn", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "Georgyske kupon larit", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "Georgyske lari", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Ghanese cedi (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Ghanese cedi", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "Gibraltarees pûn", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Gambiaanske dalasi", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "Guinese franc", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Guinese syli", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "Equatoriaal-Guinese ekwele guineana", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "Grykse drachme", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "Guatemalteekse quetzal", Symbol: "Q"},
				currency.GWE: cldr.Currency{DisplayName: "Portugees-Guinese escudo", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "Guinee-Bissause peso", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "Guyaanske dollar", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "Hongkongske dollar", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Hondurese lempira", Symbol: "L"},
				currency.HRD: cldr.Currency{DisplayName: "Kroatyske dinar", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "Kroatyske kuna", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "Haïtiaanske gourde", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "Hongaarse forint", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "Indonesyske roepia", Symbol: "Rp"},
				currency.IEP: cldr.Currency{DisplayName: "Ierske pûn", Symbol: "IEP"},
				currency.ILP: cldr.Currency{DisplayName: "Israëlysk pûn", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "Israëlyske nieuwe shekel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Indiase roepie", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Iraakse dinar", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "Iraanske rial", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Yslânske kroon", Symbol: "kr"},
				currency.ITL: cldr.Currency{DisplayName: "Italiaanske lire", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "Jamaikaanske dollar", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "Jordaanske dinar", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "Japanse yen", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Keniaanske shilling", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "Kirgizyske som", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "Kambodjaanske riel", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Komorese frank", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "Noard-Koreaanske won", Symbol: "₩"},
				currency.KRH: cldr.Currency{DisplayName: "Sûd-Koreaanske hwan (1953–1962)", Symbol: ""},
				currency.KRO: cldr.Currency{DisplayName: "Alde Sûd-Koreaanske won (1945–1953)", Symbol: ""},
				currency.KRW: cldr.Currency{DisplayName: "Sûd-Koreaanske won", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Koeweitse dinar", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "Caymaneilânske dollar", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "Kazachstaanske tenge", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "Laotiaanske kip", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "Libaneeske pûn", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "Sri Lankaanske roepie", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Liberiaanske dollar", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Lesothaanske loti", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Litouwse litas", Symbol: "Lt"},
				currency.LTT: cldr.Currency{DisplayName: "Litouwse talonas", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "Lúksemboargske convertibele franc", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "Lúksemboargske frank", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "Lúksemboargske finansjele franc", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "Letse lats", Symbol: "Ls"},
				currency.LVR: cldr.Currency{DisplayName: "Letse roebel", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "Libyske dinar", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Marokkaanske dirham", Symbol: ""},
				currency.MAF: cldr.Currency{DisplayName: "Marokkaanske franc", Symbol: ""},
				currency.MCF: cldr.Currency{DisplayName: "Monegaskyske frank", Symbol: ""},
				currency.MDC: cldr.Currency{DisplayName: "Moldavyske cupon", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Moldavyske leu", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Malagassyske ariary", Symbol: "Ar"},
				currency.MGF: cldr.Currency{DisplayName: "Malagassyske franc", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "Macedonyske denar", Symbol: ""},
				currency.MKN: cldr.Currency{DisplayName: "Macedonyske denar (1992–1993)", Symbol: ""},
				currency.MLF: cldr.Currency{DisplayName: "Malinese franc", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Myanmarese kyat", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "Mongoalske tugrik", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "Macause pataca", Symbol: ""},
				currency.MRO: cldr.Currency{DisplayName: "Mauritaanske ouguiya (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Mauritaanske ouguiya", Symbol: ""},
				currency.MTL: cldr.Currency{DisplayName: "Maltese lire", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "Maltees pûn", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Mauritiaanske roepie", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "Maldivyske rufiyaa", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "Malawyske kwacha", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "Meksikaanske peso", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "Meksikaanske sulveren peso (1861–1992)", Symbol: ""},
				currency.MXV: cldr.Currency{DisplayName: "Meksikaanske unidad de inversion (UDI)", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "Maleisyske ringgit", Symbol: "RM"},
				currency.MZE: cldr.Currency{DisplayName: "Mozambikaanske escudo", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "Alde Mozambikaanske metical", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Mozambikaanske metical", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Namibyske dollar", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Nigeriaanske naira", Symbol: "₦"},
				currency.NIC: cldr.Currency{DisplayName: "Nicaraguaanske córdoba (1988–1991)", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "Nicaraguaanske córdoba", Symbol: "C$"},
				currency.NLG: cldr.Currency{DisplayName: "Nederlânske gûne", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "Noarske kroon", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "Nepalese roepie", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "Nij-Seelânske dollar", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Omaanske rial", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "Panamese balboa", Symbol: ""},
				currency.PEI: cldr.Currency{DisplayName: "Peruaanske inti", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Peruaanske sol", Symbol: ""},
				currency.PES: cldr.Currency{DisplayName: "Peruaanske sol (1863–1985)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Papuaanske kina", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "Filipynske peso", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "Pakistaanske roepie", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "Poalske zloty", Symbol: "zł"},
				currency.PLZ: cldr.Currency{DisplayName: "Poalske zloty (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "Portugeeske escudo", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "Paraguayaanske guarani", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "Katarese rial", Symbol: ""},
				currency.RHD: cldr.Currency{DisplayName: "Rhodesyske dollar", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "Alde Roemeenske leu", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Roemeenske leu", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "Servyske dinar", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "Russyske roebel", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "Russyske roebel (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Rwandese frank", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Saoedi-Arabyske riyal", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "Salomon-dollar", Symbol: "SI$"},
				currency.SCR: cldr.Currency{DisplayName: "Seychelse roepie", Symbol: ""},
				currency.SDD: cldr.Currency{DisplayName: "Soedaneeske dinar", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Soedaneeske pûn", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Soedaneeske pûn (1957–1998)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Sweedske kroon", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "Singaporese dollar", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Sint-Heleenske pûn", Symbol: "£"},
				currency.SIT: cldr.Currency{DisplayName: "Sloveenske tolar", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "Slowaakse koruna", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Sierraleoonse leone", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Somalyske shilling", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "Surinaamske dollar", Symbol: "$"},
				currency.SRG: cldr.Currency{DisplayName: "Surinaamske gulden", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "Sûd-Soedaneeske pûn", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Santomese dobra (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Santomese dobra", Symbol: "Db"},
				currency.SUR: cldr.Currency{DisplayName: "Sovjet-roebel", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "Salvadoraanske colón", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "Syrysk pûn", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Swazyske lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "Thaise baht", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "Tadzjikistaanske roebel", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "Tadzjikistaanske somoni", Symbol: ""},
				currency.TMM: cldr.Currency{DisplayName: "Turkmeense manat (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Turkmeense manat", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "Tunesyske dinar", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "Tongaanske paʻanga", Symbol: "T$"},
				currency.TPE: cldr.Currency{DisplayName: "Timorese escudo", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "Turkse lire", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "Turkse lira", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "Trinidad en Tobago-dollar", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "Nije Taiwanese dollar", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Tanzaniaanske shilling", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "Oekraïense hryvnia", Symbol: "₴"},
				currency.UAK: cldr.Currency{DisplayName: "Oekraïense karbovanetz", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "Oegandese shilling (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "Oegandese shilling", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Amerikaanske dollar", Symbol: "US$"},
				currency.USN: cldr.Currency{DisplayName: "Amerikaanske dollar (folgjende dei)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "Amerikaanske dollar (zelfde dei)", Symbol: ""},
				currency.UYI: cldr.Currency{DisplayName: "Uruguayaanske peso en geïndexeerde eenheden", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "Uruguayaanske peso (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "Uruguayaanske peso", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "Oezbekistaanske sum", Symbol: ""},
				currency.VEB: cldr.Currency{DisplayName: "Fenezolaanske bolivar (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "Fenezolaanske bolivar (2008–2018)", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "Fenezolaanske bolivar", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "Fietnameeske dong", Symbol: "₫"},
				currency.VNN: cldr.Currency{DisplayName: "Alde Fietnameeske dong (1978–1985)", Symbol: ""},
				currency.VUV: cldr.Currency{DisplayName: "Vanuatuaanske vatu", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "Samoaanske tala", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "CFA-frank", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "Sulver", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "Goud", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "Europeeske gearfoege ienheid", Symbol: "XBA"},
				currency.XBB: cldr.Currency{DisplayName: "Europeeske monetaire ienheid", Symbol: "XBB"},
				currency.XBC: cldr.Currency{DisplayName: "Europeeske rekkenienheid (XBC)", Symbol: "XBC"},
				currency.XBD: cldr.Currency{DisplayName: "Europeeske rekkenienheid (XBD)", Symbol: "XBD"},
				currency.XCD: cldr.Currency{DisplayName: "East-Karibyske dollar", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "Special Drawing Rights", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "European Currency Unit", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "Franse gouden franc", Symbol: "XFO"},
				currency.XFU: cldr.Currency{DisplayName: "Franse UIC-franc", Symbol: "XFU"},
				currency.XOF: cldr.Currency{DisplayName: "CFA-franc BCEAO", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "Palladium", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "CFP-franc", Symbol: "XPF"},
				currency.XPT: cldr.Currency{DisplayName: "Platina", Symbol: ""},
				currency.XRE: cldr.Currency{DisplayName: "RINET-fondsen", Symbol: "XRE"},
				currency.XSU: cldr.Currency{DisplayName: "Sucre", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "Valutacode voor testdoeleinden", Symbol: "XTS"},
				currency.XUA: cldr.Currency{DisplayName: "ADB-rekkenienheid", Symbol: "XUA"},
				currency.XXX: cldr.Currency{DisplayName: "Unbekende muntienheid", Symbol: "XXX"},
				currency.YDD: cldr.Currency{DisplayName: "Jemenityske dinar", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "Jemenityske rial", Symbol: ""},
				currency.YUD: cldr.Currency{DisplayName: "Joegoslavyske harde dinar", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "Joegoslavyske noviy-dinar", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "Joegoslavyske convertibele dinar", Symbol: ""},
				currency.YUR: cldr.Currency{DisplayName: "Joegoslavyske herfoarme dinar (1992–1993)", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "Sûd-Afrikaanske rand (finansjeel)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Sûd-Afrikaanske rand", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Sambiaanske kwacha (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Sambiaanske kwacha", Symbol: "ZK"},
				currency.ZRN: cldr.Currency{DisplayName: "Saïreeske nije Saïre", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "Saïreeske Saïre", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "Simbabwaanske dollar", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "Simbabwaanske dollar (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "Simbabwaanske dollar (2008)", Symbol: ""},
			},
		},
	}
}
