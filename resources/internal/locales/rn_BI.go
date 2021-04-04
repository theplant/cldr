// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_rn_BI() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "rn_BI",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Mut.", Feb: "Gas.", Mar: "Wer.", Apr: "Mat.", May: "Gic.", Jun: "Kam.", Jul: "Nya.", Aug: "Kan.", Sep: "Nze.", Oct: "Ukw.", Nov: "Ugu.", Dec: "Uku."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Nzero", Feb: "Ruhuhuma", Mar: "Ntwarante", Apr: "Ndamukiza", May: "Rusama", Jun: "Ruheshi", Jul: "Mukakaro", Aug: "Nyandagaro", Sep: "Nyakanga", Oct: "Gitugutu", Nov: "Munyonyo", Dec: "Kigarama"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "cu.", Mon: "mbe.", Tue: "kab.", Wed: "gtu.", Thu: "kan.", Fri: "gnu.", Sat: "gnd."}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "W", Thu: "T", Fri: "F", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Ku w’indwi", Mon: "Ku wa mbere", Tue: "Ku wa kabiri", Wed: "Ku wa gatatu", Thu: "Ku wa kane", Fri: "Ku wa gatanu", Sat: "Ku wa gatandatu"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "Z.MU.", PM: "Z.MW."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "Z.MU.", PM: "Z.MW."}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00¤", CurrencyAccounting: "", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Idiramu ryo muri Leta Zunze Ubumwe z’Abarabu", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Ikwanza ryo muri Angola", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Idolari ryo muri Ositaraliya", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "Idinari ry’iribahireyini", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Ifaranga ry’Uburundi", Symbol: "FBu"},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Ipula ryo muri Botswana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Idolari rya Kanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Ifaranga rya Kongo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Ifaranga ry’Ubusuwisi", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Iyuwani ryo mu Bushinwa", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Irikaboveridiyano ryo muri Esikudo", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Ifaranga ryo muri Jibuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Idinari ryo muri Alijeriya", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Ipawundi rya Misiri", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Irinakufa ryo muri Eritereya", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Ibiri ryo muri Etiyopiya", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Iyero", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Ipawundi ryo mu Bwongereza", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Icedi ryo muri Gana", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Idalasi ryo muri Gambiya", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Ifaranga ryo muri Gineya", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Irupiya ryo mu Buhindi", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Iyeni ry’Ubuyapani", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Ishilingi rya Kenya", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Ifaranga rya Komore", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Idolari rya Liberiya", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Iloti ryo muro Lesoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Idinari rya Libiya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Idiramu ryo muri Maroke", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Iriyari ryo muri Madagasikari", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Ugwiya ryo muri Moritaniya (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Ugwiya ryo muri Moritaniya", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Irupiya ryo mu birwa bya Morise", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Ikwaca ryo muri Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Irimetikali ryo muri Mozambike", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Idolari rya Namibiya", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Inayira ryo muri Nijeriya", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "Ifaranga ry’u Rwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Iriyari ryo muri Arabiya Sawudite", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Irupiya ryo mu birwa bya Sayisheli", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Ipawundi rya Sudani", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Ipawundi rya Sente Helena", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Ilewone", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Ishilingi ryo muri Somaliya", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Idobura ryo muri Sawotome na Perensipe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Idobura ryo muri Sawotome na Perensipe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Ililangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Idinari ryo muri Tuniziya", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Ishilingi rya Tanzaniya", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Ishilingi ry’Ubugande", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Idolari ry’abanyamerika", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Irandi ryo muri Afurika y’Epfo", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Ikwaca ryo muri Zambiya (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Ikwaca ryo muri Zambiya", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Idolari ryo muri Zimbabwe", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK: "Igikani",
			language.AM: "Ikimuhariki",
			language.AR: "Icarabu",
			language.BE: "Ikibelarusiya",
			language.BG: "Ikinyabuligariya",
			language.BN: "Ikibengali",
			language.CS: "Igiceke",
			language.DE: "Ikidage",
			language.EL: "Ikigereki",
			language.EN: "Icongereza",
			language.ES: "Icesipanyolo",
			language.FA: "Igiperisi",
			language.FR: "Igifaransa",
			language.HA: "Igihawusa",
			language.HI: "Igihindi",
			language.HU: "Ikinyahongiriya",
			language.ID: "Ikinyendoziya",
			language.IG: "Ikigubo",
			language.IT: "Igitaliyani",
			language.JA: "Ikiyapani",
			language.JV: "Ikinyejava",
			language.KM: "Igikambodiya",
			language.KO: "Ikinyakoreya",
			language.MS: "Ikinyamaleziya",
			language.MY: "Ikinyabirimaniya",
			language.NE: "Ikinepali",
			language.NL: "Igiholandi",
			language.PA: "Igipunjabi",
			language.PL: "Ikinyapolonye",
			language.PT: "Igiporutugari",
			language.RN: "Ikirundi",
			language.RO: "Ikinyarumaniya",
			language.RU: "Ikirusiya",
			language.RW: "Ikinyarwanda",
			language.SO: "Igisomali",
			language.SV: "Igisuweduwa",
			language.TA: "Igitamili",
			language.TH: "Ikinyatayilandi",
			language.TR: "Igiturukiya",
			language.UK: "Ikinyayukereni",
			language.UR: "Inyeyurudu",
			language.VI: "Ikinyaviyetinamu",
			language.YO: "Ikiyoruba",
			language.ZH: "Igishinwa",
			language.ZU: "Ikizulu",
		},
		Territories: cldr.Territories{
			territory.AD: "Andora",
			territory.AE: "Leta Zunze Ubumwe z’Abarabu",
			territory.AF: "Afuganisitani",
			territory.AG: "Antigwa na Baribuda",
			territory.AI: "Angwila",
			territory.AL: "Alubaniya",
			territory.AM: "Arumeniya",
			territory.AO: "Angola",
			territory.AR: "Arijantine",
			territory.AS: "Samowa nyamerika",
			territory.AT: "Otirishe",
			territory.AU: "Ositaraliya",
			territory.AW: "Aruba",
			territory.AZ: "Azerubayijani",
			territory.BA: "Bosiniya na Herigozevine",
			territory.BB: "Barubadosi",
			territory.BD: "Bangaladeshi",
			territory.BE: "Ububiligi",
			territory.BF: "Burukina Faso",
			territory.BG: "Buligariya",
			territory.BH: "Bahareyini",
			territory.BI: "Uburundi",
			territory.BJ: "Bene",
			territory.BM: "Berimuda",
			territory.BN: "Buruneyi",
			territory.BO: "Boliviya",
			territory.BR: "Burezili",
			territory.BS: "Bahamasi",
			territory.BT: "Butani",
			territory.BW: "Botswana",
			territory.BY: "Belausi",
			territory.BZ: "Belize",
			territory.CA: "Kanada",
			territory.CD: "Repubulika Iharanira Demokarasi ya Kongo",
			territory.CF: "Repubulika ya Santarafurika",
			territory.CG: "Kongo",
			territory.CH: "Ubusuwisi",
			territory.CI: "Kotedivuware",
			territory.CK: "Izinga rya Kuku",
			territory.CL: "Shili",
			territory.CM: "Kameruni",
			territory.CN: "Ubushinwa",
			territory.CO: "Kolombiya",
			territory.CR: "Kositarika",
			territory.CU: "Kiba",
			territory.CV: "Ibirwa bya Kapuveri",
			territory.CY: "Izinga rya Shipure",
			territory.CZ: "Repubulika ya Ceke",
			territory.DE: "Ubudage",
			territory.DJ: "Jibuti",
			territory.DK: "Danimariki",
			territory.DM: "Dominika",
			territory.DO: "Repubulika ya Dominika",
			territory.DZ: "Alijeriya",
			territory.EC: "Ekwateri",
			territory.EE: "Esitoniya",
			territory.EG: "Misiri",
			territory.ER: "Elitereya",
			territory.ES: "Hisipaniya",
			territory.ET: "Etiyopiya",
			territory.FI: "Finilandi",
			territory.FJ: "Fiji",
			territory.FK: "Izinga rya Filikilandi",
			territory.FM: "Mikoroniziya",
			territory.FR: "Ubufaransa",
			territory.GA: "Gabo",
			territory.GB: "Ubwongereza",
			territory.GD: "Gerenada",
			territory.GE: "Jeworujiya",
			territory.GF: "Gwayana y’Abafaransa",
			territory.GH: "Gana",
			territory.GI: "Juburalitari",
			territory.GL: "Gurunilandi",
			territory.GM: "Gambiya",
			territory.GN: "Guneya",
			territory.GP: "Gwadelupe",
			territory.GQ: "Gineya Ekwatoriyali",
			territory.GR: "Ubugereki",
			territory.GT: "Gwatemala",
			territory.GU: "Gwamu",
			territory.GW: "Gineya Bisawu",
			territory.GY: "Guyane",
			territory.HN: "Hondurasi",
			territory.HR: "Korowasiya",
			territory.HT: "Hayiti",
			territory.HU: "Hungariya",
			territory.ID: "Indoneziya",
			territory.IE: "Irilandi",
			territory.IL: "Isiraheli",
			territory.IN: "Ubuhindi",
			territory.IO: "Intara y’Ubwongereza yo mu birwa by’Abahindi",
			territory.IQ: "Iraki",
			territory.IR: "Irani",
			territory.IS: "Ayisilandi",
			territory.IT: "Ubutaliyani",
			territory.JM: "Jamayika",
			territory.JO: "Yorudaniya",
			territory.JP: "Ubuyapani",
			territory.KE: "Kenya",
			territory.KG: "Kirigisitani",
			territory.KH: "Kamboje",
			territory.KI: "Kiribati",
			territory.KM: "Izinga rya Komore",
			territory.KN: "Sekitsi na Nevisi",
			territory.KP: "Koreya y’amajaruguru",
			territory.KR: "Koreya y’amajepfo",
			territory.KW: "Koweti",
			territory.KY: "Ibirwa bya Keyimani",
			territory.KZ: "Kazakisitani",
			territory.LA: "Layosi",
			territory.LB: "Libani",
			territory.LC: "Selusiya",
			territory.LI: "Lishyitenshitayini",
			territory.LK: "Sirilanka",
			territory.LR: "Liberiya",
			territory.LS: "Lesoto",
			territory.LT: "Lituwaniya",
			territory.LU: "Lukusamburu",
			territory.LV: "Lativa",
			territory.LY: "Libiya",
			territory.MA: "Maroke",
			territory.MC: "Monako",
			territory.MD: "Moludavi",
			territory.MG: "Madagasikari",
			territory.MH: "Izinga rya Marishari",
			territory.ML: "Mali",
			territory.MM: "Birimaniya",
			territory.MN: "Mongoliya",
			territory.MP: "Amazinga ya Mariyana ryo mu majaruguru",
			territory.MQ: "Maritiniki",
			territory.MR: "Moritaniya",
			territory.MS: "Monteserati",
			territory.MT: "Malita",
			territory.MU: "Izinga rya Morise",
			territory.MV: "Moludave",
			territory.MW: "Malawi",
			territory.MX: "Migizike",
			territory.MY: "Maleziya",
			territory.MZ: "Mozambiki",
			territory.NA: "Namibiya",
			territory.NC: "Niyukaledoniya",
			territory.NE: "Nijeri",
			territory.NF: "izinga rya Norufoluke",
			territory.NG: "Nijeriya",
			territory.NI: "Nikaragwa",
			territory.NL: "Ubuholandi",
			territory.NO: "Noruveji",
			territory.NP: "Nepali",
			territory.NR: "Nawuru",
			territory.NU: "Niyuwe",
			territory.NZ: "Nuvelizelandi",
			territory.OM: "Omani",
			territory.PA: "Panama",
			territory.PE: "Peru",
			territory.PF: "Polineziya y’Abafaransa",
			territory.PG: "Papuwa Niyugineya",
			territory.PH: "Amazinga ya Filipine",
			territory.PK: "Pakisitani",
			territory.PL: "Polonye",
			territory.PM: "Sempiyeri na Mikeloni",
			territory.PN: "Pitikeyirini",
			territory.PR: "Puwetoriko",
			territory.PS: "Palesitina Wesitibanka na Gaza",
			territory.PT: "Porutugali",
			territory.PW: "Palawu",
			territory.PY: "Paragwe",
			territory.QA: "Katari",
			territory.RE: "Amazinga ya Reyiniyo",
			territory.RO: "Rumaniya",
			territory.RU: "Uburusiya",
			territory.RW: "u Rwanda",
			territory.SA: "Arabiya Sawudite",
			territory.SB: "Amazinga ya Salumoni",
			territory.SC: "Amazinga ya Seyisheli",
			territory.SD: "Sudani",
			territory.SE: "Suwedi",
			territory.SG: "Singapuru",
			territory.SH: "Sehelene",
			territory.SI: "Siloveniya",
			territory.SK: "Silovakiya",
			territory.SL: "Siyeralewone",
			territory.SM: "Sanimarino",
			territory.SN: "Senegali",
			territory.SO: "Somaliya",
			territory.SR: "Suriname",
			territory.ST: "Sawotome na Perensipe",
			territory.SV: "Eli Saluvatori",
			territory.SY: "Siriya",
			territory.SZ: "Suwazilandi",
			territory.TC: "Amazinga ya Turkisi na Cayikosi",
			territory.TD: "Cadi",
			territory.TG: "Togo",
			territory.TH: "Tayilandi",
			territory.TJ: "Tajikisitani",
			territory.TK: "Tokelawu",
			territory.TL: "Timoru y’iburasirazuba",
			territory.TM: "Turukumenisitani",
			territory.TN: "Tuniziya",
			territory.TO: "Tonga",
			territory.TR: "Turukiya",
			territory.TT: "Tirinidadi na Tobago",
			territory.TV: "Tuvalu",
			territory.TW: "Tayiwani",
			territory.TZ: "Tanzaniya",
			territory.UA: "Ikerene",
			territory.UG: "Ubugande",
			territory.US: "Leta Zunze Ubumwe za Amerika",
			territory.UY: "Irigwe",
			territory.UZ: "Uzubekisitani",
			territory.VA: "Umurwa wa Vatikani",
			territory.VC: "Sevensa na Gerenadine",
			territory.VE: "Venezuwela",
			territory.VG: "Ibirwa by’isugi by’Abongereza",
			territory.VI: "Amazinga y’Isugi y’Abanyamerika",
			territory.VN: "Viyetinamu",
			territory.VU: "Vanuwatu",
			territory.WF: "Walisi na Futuna",
			territory.WS: "Samowa",
			territory.YE: "Yemeni",
			territory.YT: "Mayote",
			territory.ZA: "Afurika y’Epfo",
			territory.ZM: "Zambiya",
			territory.ZW: "Zimbabwe",
		},
	}
}
