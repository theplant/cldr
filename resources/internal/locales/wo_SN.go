// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_wo_SN() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "wo_SN",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "dd-MM-y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'ci' {0}", Long: "{1} 'ci' {0}", Medium: "{1} - {0}", Short: "{1} - {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Sam", Feb: "Few", Mar: "Mar", Apr: "Awr", May: "Mee", Jun: "Suw", Jul: "Sul", Aug: "Ut", Sep: "Sàt", Oct: "Okt", Nov: "Now", Dec: "Des"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Samwiyee", Feb: "Fewriyee", Mar: "Mars", Apr: "Awril", May: "Mee", Jun: "Suwe", Jul: "Sulet", Aug: "Ut", Sep: "Sàttumbar", Oct: "Oktoobar", Nov: "Nowàmbar", Dec: "Desàmbar"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Dib", Mon: "Alt", Tue: "Tal", Wed: "Àla", Thu: "Alx", Fri: "Àjj", Sat: "Ase"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Dib", Mon: "Alt", Tue: "Tal", Wed: "Àla", Thu: "Alx", Fri: "Àjj", Sat: "Ase"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Dib", Mon: "Alt", Tue: "Tal", Wed: "Àla", Thu: "Alx", Fri: "Àjj", Sat: "Ase"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Dibéer", Mon: "Altine", Tue: "Talaata", Wed: "Àlarba", Thu: "Alxamis", Fri: "Àjjuma", Sat: "Aseer"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "Sub", PM: "Ngo"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "Sub", PM: "Ngo"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "Sub", PM: "Ngo"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a0#,##0.00", CurrencyAccounting: "¤\u00a0#,##0.00", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AOA: cldr.Currency{DisplayName: "", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "Real bu Bresil", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "", Symbol: "CA$"},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Yuan bu Siin", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.EGP: cldr.Currency{DisplayName: "", Symbol: "E£"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Pound bu Grànd Brëtaañ", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupee bu End", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Yen bu Sapoŋ", Symbol: "JP¥"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.MGA: cldr.Currency{DisplayName: "", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MUR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.NAD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "", Symbol: "NZ$"},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "", Symbol: "₲"},
				currency.RON: cldr.Currency{DisplayName: "", Symbol: "lei"},
				currency.RUB: cldr.Currency{DisplayName: "Ruble bi Rsis", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "", Symbol: "RF"},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.USD: cldr.Currency{DisplayName: "Dolaaru US", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Franc CFA bu Afrik Sowwu-jant", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Xaalis buñ Xamul", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AF:      "Afrikaans",
			language.AM:      "Amharik",
			language.AR:      "Araab",
			language.AS:      "Asame",
			language.AZ:      "Aserbayjane",
			language.BA:      "Baskir",
			language.BAN:     "Bali",
			language.BE:      "Belaris",
			language.BEM:     "Bemba",
			language.BG:      "Bilgaar",
			language.BN:      "Baŋla",
			language.BO:      "Tibetan",
			language.BR:      "Breton",
			language.BS:      "Bosñak",
			language.CA:      "Katalan",
			language.CEB:     "Sibiyanoo",
			language.CHM:     "Mari",
			language.CHR:     "Ceroki",
			language.CKB:     "Kurdi gu Diggu",
			language.CO:      "Kors",
			language.CS:      "Cek",
			language.CY:      "Wels",
			language.DA:      "Danuwa",
			language.DE:      "Almaa",
			language.DSB:     "Sorab-Suuf",
			language.DV:      "Diweyi",
			language.DZ:      "Dsongkaa",
			language.EL:      "Gereg",
			language.EN:      "Àngale",
			language.EN_GB:   "Àngale (RI)",
			language.EN_US:   "Àngale (ES)",
			language.EO:      "Esperantoo",
			language.ES:      "Español",
			language.ES_419:  "Español (Amerik Latin)",
			language.ET:      "Estoñiye",
			language.EU:      "Bask",
			language.FA:      "Pers",
			language.FF:      "Pël",
			language.FI:      "Feylànde",
			language.FIL:     "Filipiye",
			language.FO:      "Feroos",
			language.FR:      "Farañse",
			language.GA:      "Irlànde",
			language.GD:      "Galuwaa bu Ekos",
			language.GL:      "Galisiye",
			language.GN:      "Garani",
			language.GU:      "Gujarati",
			language.HA:      "Hawsa",
			language.HAW:     "Hawaye",
			language.HE:      "Ebrë",
			language.HI:      "Endo",
			language.HIL:     "Hiligaynon",
			language.HR:      "Krowat",
			language.HSB:     "Sorab-Kaw",
			language.HT:      "Kereyolu Ayti",
			language.HU:      "Ongruwaa",
			language.HY:      "Armaniye",
			language.HZ:      "Herero",
			language.IBB:     "Ibibiyo",
			language.ID:      "Endonesiye",
			language.IG:      "Igbo",
			language.IS:      "Islànde",
			language.IT:      "Italiye",
			language.IU:      "Inuktitit",
			language.JA:      "Sapone",
			language.KA:      "Sorsiye",
			language.KK:      "Kasax",
			language.KM:      "Xmer",
			language.KN:      "Kannadaa",
			language.KO:      "Koreye",
			language.KOK:     "Konkani",
			language.KR:      "Kanuri",
			language.KRU:     "Kuruks",
			language.KS:      "Kashmiri",
			language.KU:      "Kurdi",
			language.KY:      "Kirgiis",
			language.LA:      "Latin",
			language.LB:      "Liksàmbursuwaa",
			language.LO:      "Laaw",
			language.LT:      "Lituyaniye",
			language.LV:      "Letoniye",
			language.MEN:     "Mende",
			language.MG:      "Malagasi",
			language.MI:      "Mawri",
			language.MK:      "Maseduwaane",
			language.ML:      "Malayalam",
			language.MN:      "Mongoliye",
			language.MNI:     "Manipuri",
			language.MOH:     "Mowak",
			language.MR:      "Marati",
			language.MS:      "Malay",
			language.MT:      "Malt",
			language.MY:      "Birmes",
			language.NE:      "Nepale",
			language.NIU:     "Niweyan",
			language.NL:      "Neyerlànde",
			language.NO:      "Nerwesiye",
			language.NY:      "Sewa",
			language.OC:      "Ositan",
			language.OM:      "Oromo",
			language.OR:      "Oja",
			language.PA:      "Punjabi",
			language.PAP:     "Papiyamento",
			language.PL:      "Polone",
			language.PS:      "Pasto",
			language.PT:      "Purtugees",
			language.QU:      "Kesuwa",
			language.QUC:     "Kishe",
			language.RM:      "Romaas",
			language.RO:      "Rumaniyee",
			language.RU:      "Rus",
			language.RW:      "Kinyarwànda",
			language.SA:      "Sanskrit",
			language.SAH:     "Saxa",
			language.SAT:     "Santali",
			language.SD:      "Sindi",
			language.SE:      "Penku Sami",
			language.SI:      "Sinala",
			language.SK:      "Eslowaki (Eslowak)",
			language.SL:      "Esloweniye",
			language.SMA:     "Sami gu Saalum",
			language.SMJ:     "Lule Sami",
			language.SMN:     "Inari Sami",
			language.SMS:     "Eskolt Sami",
			language.SO:      "Somali (làkk)",
			language.SQ:      "Albane",
			language.SR:      "Serb",
			language.SV:      "Suweduwaa",
			language.SYR:     "Siryak",
			language.TA:      "Tamil",
			language.TE:      "Telugu",
			language.TG:      "Tajis",
			language.TH:      "Tay",
			language.TI:      "Tigriña",
			language.TK:      "Tirkmen",
			language.TO:      "Tongan",
			language.TR:      "Tirk",
			language.TT:      "Tatar",
			language.TZM:     "Tamasis gu Digg Atlaas",
			language.UG:      "Uygur",
			language.UK:      "Ikreniye",
			language.UND:     "Làkk wuñ xamul",
			language.UR:      "Urdu",
			language.UZ:      "Usbek",
			language.VE:      "Wenda",
			language.VI:      "Wiyetnaamiye",
			language.WO:      "Wolof",
			language.YI:      "Yidis",
			language.YO:      "Yoruba",
			language.ZH:      "Sinuwaa",
			language.ZH_HANS: "Sinuwaa buñ woyofal",
			language.ZH_HANT: "Sinuwaa bu cosaan",
		},
		Territories: cldr.Territories{
			territory.AD: "Andoor",
			territory.AE: "Emira Arab Ini",
			territory.AF: "Afganistaŋ",
			territory.AG: "Antiguwa ak Barbuda",
			territory.AI: "Angiiy",
			territory.AL: "Albani",
			territory.AM: "Armeni",
			territory.AO: "Àngolaa",
			territory.AQ: "Antarktik",
			territory.AR: "Arsàntin",
			territory.AS: "Samowa bu Amerig",
			territory.AT: "Ótiriis",
			territory.AU: "Ostarali",
			territory.AW: "Aruba",
			territory.AX: "Duni Aalànd",
			territory.AZ: "Aserbayjaŋ",
			territory.BA: "Bosni Ersegowin",
			territory.BB: "Barbad",
			territory.BD: "Bengalades",
			territory.BE: "Belsig",
			territory.BF: "Burkina Faaso",
			territory.BG: "Bilgari",
			territory.BH: "Bahreyin",
			territory.BI: "Burundi",
			territory.BJ: "Benee",
			territory.BL: "Saŋ Bartalemi",
			territory.BM: "Bermid",
			territory.BN: "Burney",
			territory.BO: "Boliwi",
			territory.BR: "Beresil",
			territory.BS: "Bahamas",
			territory.BT: "Butaŋ",
			territory.BV: "Dunu Buwet",
			territory.BW: "Botswana",
			territory.BY: "Belaris",
			territory.BZ: "Belis",
			territory.CA: "Kanadaa",
			territory.CC: "Duni Koko (Kilin)",
			territory.CD: "Kongo (R K D)",
			territory.CF: "Repiblik Sàntar Afrik",
			territory.CG: "Réewum Kongo",
			territory.CH: "Siwis",
			territory.CI: "Kodiwaar",
			territory.CK: "Duni Kuuk",
			territory.CL: "Sili",
			territory.CM: "Kamerun",
			territory.CN: "Siin",
			territory.CO: "Kolombi",
			territory.CR: "Kosta Rika",
			territory.CU: "Kuba",
			territory.CV: "Kabo Werde",
			territory.CW: "Kursawo",
			territory.CX: "Dunu Kirismas",
			territory.CY: "Siipar",
			territory.CZ: "Réewum Cek",
			territory.DE: "Almaañ",
			territory.DJ: "Jibuti",
			territory.DK: "Danmàrk",
			territory.DM: "Dominik",
			territory.DO: "Repiblik Dominiken",
			territory.DZ: "Alseri",
			territory.EC: "Ekwaatër",
			territory.EE: "Estoni",
			territory.EG: "Esipt",
			territory.ER: "Eritere",
			territory.ES: "Españ",
			territory.ET: "Ecopi",
			territory.FI: "Finlànd",
			territory.FJ: "Fijji",
			territory.FK: "Duni Falkland",
			territory.FM: "Mikoronesi",
			territory.FO: "Duni Faro",
			territory.FR: "Faraans",
			territory.GA: "Gaboŋ",
			territory.GB: "Ruwaayom Ini",
			territory.GD: "Garanad",
			territory.GE: "Seworsi",
			territory.GF: "Guyaan Farañse",
			territory.GG: "Gernase",
			territory.GH: "Gana",
			territory.GI: "Sibraltaar",
			territory.GL: "Girinlànd",
			territory.GM: "Gàmbi",
			territory.GN: "Gine",
			territory.GP: "Guwaadelup",
			territory.GQ: "Gine Ekuwatoriyal",
			territory.GR: "Gerees",
			territory.GS: "Seworsi di Sid ak Duni Sàndwiis di Sid",
			territory.GT: "Guwatemala",
			territory.GU: "Guwam",
			territory.GW: "Gine-Bisaawóo",
			territory.GY: "Giyaan",
			territory.HK: "Ooŋ Koŋ",
			territory.HM: "Duni Hërd ak Duni MakDonald",
			territory.HN: "Onduraas",
			territory.HR: "Korowasi",
			territory.HT: "Ayti",
			territory.HU: "Ongari",
			territory.ID: "Indonesi",
			territory.IE: "Irlànd",
			territory.IL: "Israyel",
			territory.IM: "Dunu Maan",
			territory.IN: "End",
			territory.IO: "Terituwaaru Brëtaañ ci Oseyaa Enjeŋ",
			territory.IQ: "Irag",
			territory.IR: "Iraŋ",
			territory.IS: "Islànd",
			territory.IT: "Itali",
			territory.JE: "Serse",
			territory.JM: "Samayig",
			territory.JO: "Sordani",
			territory.JP: "Sàppoŋ",
			territory.KE: "Keeña",
			territory.KG: "Kirgistaŋ",
			territory.KH: "Kàmboj",
			territory.KI: "Kiribati",
			territory.KM: "Komoor",
			territory.KN: "Saŋ Kits ak Newis",
			territory.KP: "Kore Noor",
			territory.KW: "Kowet",
			territory.KY: "Duni Kaymaŋ",
			territory.KZ: "Kasaxstaŋ",
			territory.LA: "Lawos",
			territory.LB: "Libaa",
			territory.LC: "Saŋ Lusi",
			territory.LI: "Liktensteyin",
			territory.LK: "Siri Lànka",
			territory.LR: "Liberiya",
			territory.LS: "Lesoto",
			territory.LT: "Litiyani",
			territory.LU: "Liksàmbur",
			territory.LV: "Letoni",
			territory.LY: "Libi",
			territory.MA: "Marog",
			territory.MC: "Monako",
			territory.MD: "Moldawi",
			territory.ME: "Montenegoro",
			territory.MF: "Saŋ Marteŋ",
			territory.MG: "Madagaskaar",
			territory.MH: "Duni Marsaal",
			territory.MK: "Maseduwaan bëj Gànnaar",
			territory.ML: "Mali",
			territory.MM: "Miyanmaar",
			territory.MN: "Mongoli",
			territory.MO: "Makaawo",
			territory.MP: "Duni Mariyaan Noor",
			territory.MQ: "Martinik",
			territory.MR: "Mooritani",
			territory.MS: "Mooseraa",
			territory.MT: "Malt",
			territory.MU: "Moriis",
			territory.MV: "Maldiiw",
			territory.MW: "Malawi",
			territory.MX: "Meksiko",
			territory.MY: "Malesi",
			territory.MZ: "Mosàmbig",
			territory.NA: "Namibi",
			territory.NC: "Nuwel Kaledoni",
			territory.NE: "Niiseer",
			territory.NF: "Dunu Norfolk",
			territory.NG: "Niseriya",
			territory.NI: "Nikaraguwa",
			territory.NL: "Peyi Baa",
			territory.NO: "Norwees",
			territory.NP: "Nepaal",
			territory.NR: "Nawru",
			territory.NU: "Niw",
			territory.NZ: "Nuwel Selànd",
			territory.OM: "Omaan",
			territory.PA: "Panama",
			territory.PE: "Peru",
			territory.PF: "Polinesi Farañse",
			territory.PG: "Papuwasi Gine Gu Bees",
			territory.PH: "Filipin",
			territory.PK: "Pakistaŋ",
			territory.PL: "Poloñ",
			territory.PM: "Saŋ Peer ak Mikeloŋ",
			territory.PN: "Duni Pitkayirn",
			territory.PR: "Porto Riko",
			territory.PT: "Portigaal",
			territory.PW: "Palaw",
			territory.PY: "Paraguwe",
			territory.QA: "Kataar",
			territory.RE: "Reeñoo",
			territory.RO: "Rumani",
			territory.RS: "Serbi",
			territory.RU: "Risi",
			territory.RW: "Ruwànda",
			territory.SA: "Arabi Sawudi",
			territory.SB: "Duni Salmoon",
			territory.SC: "Seysel",
			territory.SD: "Sudaŋ",
			territory.SE: "Suwed",
			territory.SG: "Singapuur",
			territory.SH: "Saŋ Eleen",
			territory.SI: "Esloweni",
			territory.SJ: "Swalbaar ak Jan Mayen",
			territory.SK: "Eslowaki",
			territory.SL: "Siyera Lewon",
			territory.SM: "San Marino",
			territory.SN: "Senegaal",
			territory.SO: "Somali",
			territory.SR: "Sirinam",
			territory.SS: "Sudaŋ di Sid",
			territory.ST: "Sawo Tome ak Pirinsipe",
			territory.SV: "El Salwadoor",
			territory.SX: "Sin Marten",
			territory.SY: "Siri",
			territory.SZ: "Suwasilànd",
			territory.TC: "Duni Tirk ak Kaykos",
			territory.TD: "Càdd",
			territory.TF: "Teer Ostraal gu Fraas",
			territory.TG: "Togo",
			territory.TH: "Taylànd",
			territory.TJ: "Tajikistaŋ",
			territory.TK: "Tokoloo",
			territory.TL: "Timor Leste",
			territory.TM: "Tirkmenistaŋ",
			territory.TN: "Tinisi",
			territory.TO: "Tonga",
			territory.TR: "Tirki",
			territory.TT: "Tirinite ak Tobago",
			territory.TV: "Tuwalo",
			territory.TW: "Taywan",
			territory.TZ: "Taŋsani",
			territory.UA: "Ikeren",
			territory.UG: "Ugànda",
			territory.UM: "Duni Amerig Utar meer",
			territory.US: "Etaa Sini",
			territory.UY: "Uruge",
			territory.UZ: "Usbekistaŋ",
			territory.VA: "Site bu Watikaa",
			territory.VC: "Saŋ Weesaa ak Garanadin",
			territory.VE: "Wenesiyela",
			territory.VG: "Duni Wirsin yu Brëtaañ",
			territory.VI: "Duni Wirsin yu Etaa-sini",
			territory.VN: "Wiyetnam",
			territory.VU: "Wanuatu",
			territory.WF: "Walis ak Futuna",
			territory.WS: "Samowa",
			territory.XK: "Kosowo",
			territory.YE: "Yaman",
			territory.YT: "Mayot",
			territory.ZA: "Afrik di Sid",
			territory.ZM: "Sàmbi",
			territory.ZW: "Simbabwe",
			territory.ZZ: "Gox buñ xamul",
		},
	}
}
