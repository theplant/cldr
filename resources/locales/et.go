// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_et() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM y", Short: "dd.MM.yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT {0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jaan", Feb: "veebr", Mar: "märts", Apr: "apr", May: "mai", Jun: "juuni", Jul: "juuli", Aug: "aug", Sep: "sept", Oct: "okt", Nov: "nov", Dec: "dets"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "V", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "jaanuar", Feb: "veebruar", Mar: "märts", Apr: "aprill", May: "mai", Jun: "juuni", Jul: "juuli", Aug: "august", Sep: "september", Oct: "oktoober", Nov: "november", Dec: "detsember"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "P", Mon: "E", Tue: "T", Wed: "K", Thu: "N", Fri: "R", Sat: "L"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "P", Mon: "E", Tue: "T", Wed: "K", Thu: "N", Fri: "R", Sat: "L"}, Short: cldr.CalendarDayFormatNameValue{Sun: "P", Mon: "E", Tue: "T", Wed: "K", Thu: "N", Fri: "R", Sat: "L"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "pühapäev", Mon: "esmaspäev", Tue: "teisipäev", Wed: "kolmapäev", Thu: "neljapäev", Fri: "reede", Sat: "laupäev"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_et]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "−", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤;(#,##0.00\u00a0¤)", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "Andorra peseeta", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "Araabia Ühendemiraatide dirhem", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "Afganistani afgaani (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Afganistani afgaani", Symbol: "AFN"},
				currency.ALK: cldr.Currency{DisplayName: "Albaania lekk (1946–1965)", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "Albaania lekk", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Armeenia dramm", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Hollandi Antillide kulden", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Angola kvanza", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "Angola kvanza (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "Angola kvanza (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "Angola reformitud kvanza, 1995–1999", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "Argentina austral", Symbol: ""},
				currency.ARM: cldr.Currency{DisplayName: "Argentina peeso (1881–1970)", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "Argentina peeso (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "Argentina peeso", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "Austria šilling", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "Austraalia dollar", Symbol: "AU$"},
				currency.AWG: cldr.Currency{DisplayName: "Aruba kulden", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "Aserbaidžaani manat (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Aserbaidžaani manat", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "Bosnia ja Hertsegoviina dinaar (1992–1994)", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Bosnia ja Hertsegoviina konverteeritav mark", Symbol: "BAM"},
				currency.BAN: cldr.Currency{DisplayName: "Bosnia ja Hertsegoviina uus dinaar (1994–1997)", Symbol: ""},
				currency.BBD: cldr.Currency{DisplayName: "Barbadose dollar", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Bangladeshi taka", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "Belgia konverteeritav frank", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "Belgia frank", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "Belgia arveldusfrank", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "Bulgaaria püsiv leev", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "Bulgaaria leev", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "Bulgaaria leev (1879–1952)", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Bahreini dinaar", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Burundi frank", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Bermuda dollar", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Brunei dollar", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Boliivia boliviaano", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "Boliivia boliviaano (1863–1963)", Symbol: ""},
				currency.BOP: cldr.Currency{DisplayName: "Boliivia peeso", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "Brasiilia uus kruseiro (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "Brasiilia krusado", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "Brasiilia kruseiro (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "Brasiilia reaal", Symbol: "R$"},
				currency.BRR: cldr.Currency{DisplayName: "Brasiilia kruseiro (1993–1994)", Symbol: ""},
				currency.BRZ: cldr.Currency{DisplayName: "Brasiilia kruseiro (1942–1967)", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "Bahama dollar", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Bhutani ngultrum", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "Birma kjatt", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Botswana pula", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "Valgevene uus rubla (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "Valgevene rubla", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Valgevene rubla (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "Belize’i dollar", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Kanada dollar", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Kongo frank", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "Šveitsi frank", Symbol: "CHF"},
				currency.CLE: cldr.Currency{DisplayName: "Tšiili eskuudo", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Tšiili peeso", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Hiina jüaan (välismaine turg)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "Hiina jüaan", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Colombia peeso", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "Costa Rica koloon", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "Serbia dinaar (2002–2006)", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "Kuuba konverteeritav peeso", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Kuuba peeso", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Cabo Verde eskuudo", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "Küprose nael", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Tšehhi kroon", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "Ida-Saksa mark", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "Saksa mark", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "Djibouti frank", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Taani kroon", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Dominikaani peeso", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Alžeeria dinaar", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "Ecuadori sukre", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "Eesti kroon", Symbol: "kr"},
				currency.EGP: cldr.Currency{DisplayName: "Egiptuse nael", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Eritrea nakfa", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "Hispaania peseeta", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Etioopia birr", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Soome mark", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "Fidži dollar", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Falklandi saarte nael", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "Prantsuse frank", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "Suurbritannia naelsterling", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Gruusia lari", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "Ghana sedi (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Ghana sedi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Gibraltari nael", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Gambia dalasi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Guinea frank", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "Guinea syli", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "Kreeka drahm", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "Guatemala ketsaal", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "Portugali Guinea eskuudo", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "Guinea-Bissau peeso", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "Guyana dollar", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Hongkongi dollar", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Hondurase lempiira", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "Horvaatia dinaar", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "Horvaatia kuna", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Haiti gurd", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Ungari forint", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Indoneesia ruupia", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "Iiri nael", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "Iisraeli nael", Symbol: ""},
				currency.ILR: cldr.Currency{DisplayName: "Iisraeli seekel (1980–1985)", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "Iisraeli uus seekel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "India ruupia", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Iraagi dinaar", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Iraani riaal", Symbol: "IRR"},
				currency.ISJ: cldr.Currency{DisplayName: "Islandi kroon (1918–1981)", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Islandi kroon", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "Itaalia liir", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "Jamaica dollar", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Jordaania dinaar", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Jaapani jeen", Symbol: "¥"},
				currency.KES: cldr.Currency{DisplayName: "Keenia šilling", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Kõrgõzstani somm", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Kambodža riaal", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Komoori frank", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Põhja-Korea vonn", Symbol: "KPW"},
				currency.KRO: cldr.Currency{DisplayName: "Lõuna-Korea vonn (1945–1953)", Symbol: ""},
				currency.KRW: cldr.Currency{DisplayName: "Lõuna-Korea vonn", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Kuveidi dinaar", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Kaimanisaarte dollar", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Kasahstani tenge", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Laose kiip", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Liibanoni nael", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Sri Lanka ruupia", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Libeeria dollar", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Lesotho loti", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Leedu litt", Symbol: "Lt"},
				currency.LUC: cldr.Currency{DisplayName: "Luksemburgi konverteeritav frank", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "Luksemburgi frank", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "Läti latt", Symbol: "Ls"},
				currency.LVR: cldr.Currency{DisplayName: "Läti rubla", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "Liibüa dinaar", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Maroko dirhem", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "Maroko frank", Symbol: ""},
				currency.MCF: cldr.Currency{DisplayName: "Monaco frank", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Moldova leu", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Madagaskari ariari", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "Madagaskar frank", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "Makedoonia dinaar", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "Makedoonia dinaar (1992–1993)", Symbol: ""},
				currency.MLF: cldr.Currency{DisplayName: "Mali frank", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Myanmari kjatt", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Mongoolia tugrik", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Macau pataaka", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Mauritaania ugia (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Mauritaania ugia", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "Malta liir", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "Malta nael", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Mauritiuse ruupia", Symbol: "MUR"},
				currency.MVP: cldr.Currency{DisplayName: "Maldiivi ruupia (1947–1981)", Symbol: ""},
				currency.MVR: cldr.Currency{DisplayName: "Maldiivi ruupia", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Malawi kvatša", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Mehhiko peeso", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "Mehhiko peeso (1861–1992)", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "Malaisia ringgit", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "Mosambiigi eskuudo", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "Mosambiigi metikal (1980–2006)", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Mosambiigi metikal", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Namiibia dollar", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Nigeeria naira", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "Nicaragua kordoba (1988–1991)", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "Nicaragua kordoba", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "Hollandi kulden", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "Norra kroon", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Nepali ruupia", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Uus-Meremaa dollar", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Omaani riaal", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Panama balboa", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "Peruu inti", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Peruu soll", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "Peruu soll (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Paapua Uus-Guinea kina", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Filipiini peeso", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Pakistani ruupia", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Poola zlott", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "Poola zlott (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "Portugali eskuudo", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "Paraguay guaranii", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Katari riaal", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "Rodeesia dollar", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "Rumeenia leu (1952–2006)", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Rumeenia leu", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Serbia dinaar", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Venemaa rubla", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "Venemaa rubla (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Rwanda frank", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Saudi Araabia riaal", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Saalomoni Saarte dollar", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Seišelli ruupia", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "Sudaani dinaar (1992–2007)", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Sudaani nael", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "Sudaani nael (1957–1998)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Rootsi kroon", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Singapuri dollar", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Saint Helena nael", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "Sloveenia tolar", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "Slovaki kroon", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Sierra Leone leoone", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Somaalia šilling", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Suriname dollar", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "Suriname kulden", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "Lõuna-Sudaani nael", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "São Tomé ja Príncipe dobra (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "São Tomé ja Príncipe dobra", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "NSVL-i rubla", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "El Salvadori koloon", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "Süüria nael", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Svaasimaa lilangeni", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Tai baat", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "Tadžikistani rubla", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "Tadžikistani somoni", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "Türkmenistani manat (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Türkmenistani manat", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Tuneesia dinaar", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Tonga pa’anga", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "Timori eskuudo", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "Türgi liir (1922–2005)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "Türgi liir", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Trinidadi ja Tobago dollar", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "uus Taiwani dollar", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Tansaania šilling", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Ukraina grivna", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "Ukraina karbovanets", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "Uganda šilling (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "Uganda šilling", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "USA dollar", Symbol: "$"},
				currency.USN: cldr.Currency{DisplayName: "USA järgmise päeva dollar", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "USA sama päeva dollar", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "Uruguay peeso (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "Uruguay peeso", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Usbekistani somm", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "Venezuela boliivar (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "Venezuela boliivar (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Venezuela boliivar", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Vietnami dong", Symbol: "₫"},
				currency.VNN: cldr.Currency{DisplayName: "Vietnami dong (1978–1985)", Symbol: ""},
				currency.VUV: cldr.Currency{DisplayName: "Vanuatu vatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Samoa taala", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Kesk-Aafrika CFA frank", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "hõbe", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "kuld", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "EURCO", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "Euroopa rahaühik", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "Euroopa rahaline arvestusühik (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "Euroopa rahaline arvestusühik (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "Ida-Kariibi dollar", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "Rahvusvahelise Valuutafondi arvestusühik", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "eküü", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "Prantsuse kuldfrank", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "Prantsuse UIC-frank", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "Lääne-Aafrika CFA frank", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "pallaadium", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "CFP frank", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "plaatina", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "vääringute testkood", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "määramata rahaühik", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "Jeemeni dinaar", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "Jeemeni riaal", Symbol: "YER"},
				currency.YUM: cldr.Currency{DisplayName: "Jugoslaavia uus dinaar (1994–2002)", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "Jugoslaavia konverteeritav dinaar (1990–1992)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Lõuna-Aafrika rand", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Sambia kvatša (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Sambia kvatša", Symbol: "ZMW"},
				currency.ZRZ: cldr.Currency{DisplayName: "Sairi zaire", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "Zimbabwe dollar (1980–2008)", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "Zimbabwe dollar (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "Zimbabwe dollar (2008)", Symbol: ""},
			},
		},
	}
}
