// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_yue_Hant() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y年M月d日 EEEE", Long: "y年M月d日", Medium: "y年M月d日", Short: "y/M/d"}, Time: cldr.CalendarDateFormat{Full: "ah:mm:ss [zzzz]", Long: "ah:mm:ss [z]", Medium: "ah:mm:ss", Short: "ah:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "1月", Feb: "2月", Mar: "3月", Apr: "4月", May: "5月", Jun: "6月", Jul: "7月", Aug: "8月", Sep: "9月", Oct: "10月", Nov: "11月", Dec: "12月"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "1月", Feb: "2月", Mar: "3月", Apr: "4月", May: "5月", Jun: "6月", Jul: "7月", Aug: "8月", Sep: "9月", Oct: "10月", Nov: "11月", Dec: "12月"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "星期日", Mon: "星期一", Tue: "星期二", Wed: "星期三", Thu: "星期四", Fri: "星期五", Sat: "星期六"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "日", Mon: "一", Tue: "二", Wed: "三", Thu: "四", Fri: "五", Sat: "六"}, Short: cldr.CalendarDayFormatNameValue{Sun: "日", Mon: "一", Tue: "二", Wed: "三", Thu: "四", Fri: "五", Sat: "六"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "星期日", Mon: "星期一", Tue: "星期二", Wed: "星期三", Thu: "星期四", Fri: "星期五", Sat: "星期六"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "上午", PM: "下午"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "上午", PM: "下午"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "上午", PM: "下午"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_yue]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "安道爾陪士特", Symbol: "ADP"},
				currency.AED: cldr.Currency{DisplayName: "阿拉伯聯合大公國迪爾汗", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "阿富汗尼 (1927–2002)", Symbol: "AFA"},
				currency.AFN: cldr.Currency{DisplayName: "阿富汗尼", Symbol: "AFN"},
				currency.ALK: cldr.Currency{DisplayName: "阿爾巴尼亞列克 (1946–1965)", Symbol: "ALK"},
				currency.ALL: cldr.Currency{DisplayName: "阿爾巴尼亞列克", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "亞美尼亞德拉姆", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "荷屬安地列斯盾", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "安哥拉寬扎", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "安哥拉寬扎 (1977–1990)", Symbol: "AOK"},
				currency.AON: cldr.Currency{DisplayName: "安哥拉新寬扎 (1990–2000)", Symbol: "AON"},
				currency.AOR: cldr.Currency{DisplayName: "安哥拉新調寬扎 (1995–1999)", Symbol: "AOR"},
				currency.ARA: cldr.Currency{DisplayName: "阿根廷奧斯特納爾", Symbol: "ARA"},
				currency.ARL: cldr.Currency{DisplayName: "阿根廷披索 (1970–1983)", Symbol: "ARL"},
				currency.ARM: cldr.Currency{DisplayName: "阿根廷披索 (1881–1970)", Symbol: "ARM"},
				currency.ARP: cldr.Currency{DisplayName: "阿根廷披索 (1983–1985)", Symbol: "ARP"},
				currency.ARS: cldr.Currency{DisplayName: "阿根廷披索", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "奧地利先令", Symbol: "ATS"},
				currency.AUD: cldr.Currency{DisplayName: "澳幣", Symbol: "AU$"},
				currency.AWG: cldr.Currency{DisplayName: "阿路巴盾", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "亞塞拜然馬納特 (1993–2006)", Symbol: "AZM"},
				currency.AZN: cldr.Currency{DisplayName: "亞塞拜然馬納特", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "波士尼亞-赫塞哥維納第納爾", Symbol: "BAD"},
				currency.BAM: cldr.Currency{DisplayName: "波士尼亞-赫塞哥維納可轉換馬克", Symbol: "BAM"},
				currency.BAN: cldr.Currency{DisplayName: "波士尼亞-赫塞哥維納新第納爾", Symbol: "BAN"},
				currency.BBD: cldr.Currency{DisplayName: "巴貝多元", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "孟加拉塔卡", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "比利時法郎（可轉換）", Symbol: "BEC"},
				currency.BEF: cldr.Currency{DisplayName: "比利時法郎", Symbol: "BEF"},
				currency.BEL: cldr.Currency{DisplayName: "比利時法郎（金融）", Symbol: "BEL"},
				currency.BGL: cldr.Currency{DisplayName: "保加利亞硬列弗", Symbol: "BGL"},
				currency.BGM: cldr.Currency{DisplayName: "保加利亞社會黨列弗", Symbol: "BGM"},
				currency.BGN: cldr.Currency{DisplayName: "保加利亞新列弗", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "保加利亞列弗 (1879–1952)", Symbol: "BGO"},
				currency.BHD: cldr.Currency{DisplayName: "巴林第納爾", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "蒲隆地法郎", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "百慕達幣", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "汶萊元", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "玻利維亞諾", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "玻利維亞玻利維亞諾 (1863–1963)", Symbol: "BOL"},
				currency.BOP: cldr.Currency{DisplayName: "玻利維亞披索", Symbol: "BOP"},
				currency.BOV: cldr.Currency{DisplayName: "玻利維亞幕多", Symbol: "BOV"},
				currency.BRB: cldr.Currency{DisplayName: "巴西克魯薩多農瓦 (1967–1986)", Symbol: "BRB"},
				currency.BRC: cldr.Currency{DisplayName: "巴西克魯賽羅 (1986–1989)", Symbol: "BRC"},
				currency.BRE: cldr.Currency{DisplayName: "巴西克魯賽羅 (1990–1993)", Symbol: "BRE"},
				currency.BRL: cldr.Currency{DisplayName: "巴西里拉", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "巴西克如爾達農瓦", Symbol: "BRN"},
				currency.BRR: cldr.Currency{DisplayName: "巴西克魯賽羅 (1993–1994)", Symbol: "BRR"},
				currency.BRZ: cldr.Currency{DisplayName: "巴西克魯賽羅 (1942 –1967)", Symbol: "BRZ"},
				currency.BSD: cldr.Currency{DisplayName: "巴哈馬元", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "不丹那特倫", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "緬甸基雅特", Symbol: "BUK"},
				currency.BWP: cldr.Currency{DisplayName: "波札那普拉", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "白俄羅斯新盧布 (1994–1999)", Symbol: "BYB"},
				currency.BYN: cldr.Currency{DisplayName: "白俄羅斯盧布", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "白俄羅斯盧布 (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "貝里斯元", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "加幣", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "剛果法郎", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "歐元 (WIR)", Symbol: "CHE"},
				currency.CHF: cldr.Currency{DisplayName: "瑞士法郎", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "法郎 (WIR)", Symbol: "CHW"},
				currency.CLE: cldr.Currency{DisplayName: "智利埃斯庫多", Symbol: "CLE"},
				currency.CLF: cldr.Currency{DisplayName: "卡林油達佛曼跎", Symbol: "CLF"},
				currency.CLP: cldr.Currency{DisplayName: "智利披索", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "人民幣 (離岸)", Symbol: "CNH"},
				currency.CNX: cldr.Currency{DisplayName: "", Symbol: "CNX"},
				currency.CNY: cldr.Currency{DisplayName: "人民幣", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "哥倫比亞披索", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "哥倫比亞幣 (COU)", Symbol: "COU"},
				currency.CRC: cldr.Currency{DisplayName: "哥斯大黎加科朗", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "舊塞爾維亞第納爾", Symbol: "CSD"},
				currency.CSK: cldr.Currency{DisplayName: "捷克斯洛伐克硬克朗", Symbol: "CSK"},
				currency.CUC: cldr.Currency{DisplayName: "古巴可轉換披索", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "古巴披索", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "維德角埃斯庫多", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "賽普勒斯鎊", Symbol: "CYP"},
				currency.CZK: cldr.Currency{DisplayName: "捷克克朗", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "東德奧斯特馬克", Symbol: "DDM"},
				currency.DEM: cldr.Currency{DisplayName: "德國馬克", Symbol: "DEM"},
				currency.DJF: cldr.Currency{DisplayName: "吉布地法郎", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "丹麥克朗", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "多明尼加披索", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "阿爾及利亞第納爾", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "厄瓜多蘇克雷", Symbol: "ECS"},
				currency.ECV: cldr.Currency{DisplayName: "厄瓜多爾由里達瓦康斯坦 (UVC)", Symbol: "ECV"},
				currency.EEK: cldr.Currency{DisplayName: "愛沙尼亞克朗", Symbol: "EEK"},
				currency.EGP: cldr.Currency{DisplayName: "埃及鎊", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "厄立特里亞納克法", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "西班牙比塞塔（會計單位）", Symbol: "ESA"},
				currency.ESB: cldr.Currency{DisplayName: "西班牙比塞塔（可轉換會計單位）", Symbol: "ESB"},
				currency.ESP: cldr.Currency{DisplayName: "西班牙陪士特", Symbol: "ESP"},
				currency.ETB: cldr.Currency{DisplayName: "衣索比亞比爾", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "歐元", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "芬蘭馬克", Symbol: "FIM"},
				currency.FJD: cldr.Currency{DisplayName: "斐濟元", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "福克蘭群島鎊", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "法國法郎", Symbol: "FRF"},
				currency.GBP: cldr.Currency{DisplayName: "英鎊", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "喬治亞庫旁拉里", Symbol: "GEK"},
				currency.GEL: cldr.Currency{DisplayName: "喬治亞拉里", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "迦納賽地 (1979–2007)", Symbol: "GHC"},
				currency.GHS: cldr.Currency{DisplayName: "迦納塞地", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "直布羅陀鎊", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "甘比亞達拉西", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "幾內亞法郎", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "幾內亞西里", Symbol: "GNS"},
				currency.GQE: cldr.Currency{DisplayName: "赤道幾內亞埃奎勒", Symbol: "GQE"},
				currency.GRD: cldr.Currency{DisplayName: "希臘德拉克馬", Symbol: "GRD"},
				currency.GTQ: cldr.Currency{DisplayName: "瓜地馬拉格查爾", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "葡屬幾內亞埃斯庫多", Symbol: "GWE"},
				currency.GWP: cldr.Currency{DisplayName: "幾內亞比索披索", Symbol: "GWP"},
				currency.GYD: cldr.Currency{DisplayName: "圭亞那元", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "港幣", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "洪都拉斯倫皮拉", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "克羅埃西亞第納爾", Symbol: "HRD"},
				currency.HRK: cldr.Currency{DisplayName: "克羅埃西亞庫納", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "海地古德", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "匈牙利福林", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "印尼盾", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "愛爾蘭鎊", Symbol: "IEP"},
				currency.ILP: cldr.Currency{DisplayName: "以色列鎊", Symbol: "ILP"},
				currency.ILR: cldr.Currency{DisplayName: "以色列謝克爾 (1980–1985)", Symbol: "ILR"},
				currency.ILS: cldr.Currency{DisplayName: "以色列新謝克爾", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "印度盧比", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "伊拉克第納爾", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "伊朗里亞爾", Symbol: "IRR"},
				currency.ISJ: cldr.Currency{DisplayName: "冰島克朗 (1918–1981)", Symbol: "ISJ"},
				currency.ISK: cldr.Currency{DisplayName: "冰島克朗", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "義大利里拉", Symbol: "ITL"},
				currency.JMD: cldr.Currency{DisplayName: "牙買加元", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "約旦第納爾", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "日圓", Symbol: "¥"},
				currency.KES: cldr.Currency{DisplayName: "肯尼亞先令", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "吉爾吉斯索姆", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "柬埔寨瑞爾", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "科摩羅法郎", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "北韓圓", Symbol: "KPW"},
				currency.KRH: cldr.Currency{DisplayName: "南韓圜", Symbol: "KRH"},
				currency.KRO: cldr.Currency{DisplayName: "南韓圓", Symbol: "KRO"},
				currency.KRW: cldr.Currency{DisplayName: "韓圓", Symbol: "￦"},
				currency.KWD: cldr.Currency{DisplayName: "科威特第納爾", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "開曼群島元", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "卡扎克斯坦坦吉", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "寮國基普", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "黎巴嫩鎊", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "斯里蘭卡盧比", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "賴比瑞亞元", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "賴索托洛蒂", Symbol: "LSL"},
				currency.LTL: cldr.Currency{DisplayName: "立陶宛立特", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "立陶宛特羅", Symbol: "LTT"},
				currency.LUC: cldr.Currency{DisplayName: "盧森堡可兌換法郎", Symbol: "LUC"},
				currency.LUF: cldr.Currency{DisplayName: "盧森堡法郎", Symbol: "LUF"},
				currency.LUL: cldr.Currency{DisplayName: "盧森堡金融法郎", Symbol: "LUL"},
				currency.LVL: cldr.Currency{DisplayName: "拉脫維亞拉特銀幣", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "拉脫維亞盧布", Symbol: "LVR"},
				currency.LYD: cldr.Currency{DisplayName: "利比亞第納爾", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "摩洛哥迪拉姆", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "摩洛哥法郎", Symbol: "MAF"},
				currency.MCF: cldr.Currency{DisplayName: "摩納哥法郎", Symbol: "MCF"},
				currency.MDC: cldr.Currency{DisplayName: "摩爾多瓦券", Symbol: "MDC"},
				currency.MDL: cldr.Currency{DisplayName: "摩杜雲列伊", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "馬達加斯加阿里亞里", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "馬達加斯加法郎", Symbol: "MGF"},
				currency.MKD: cldr.Currency{DisplayName: "馬其頓第納爾", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "馬其頓第納爾 (1992–1993)", Symbol: "MKN"},
				currency.MLF: cldr.Currency{DisplayName: "馬里法郎", Symbol: "MLF"},
				currency.MMK: cldr.Currency{DisplayName: "緬甸元", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "蒙古圖格里克", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "澳門元", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "茅利塔尼亞烏吉亞 (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "茅利塔尼亞烏吉亞", Symbol: ""},
				currency.MTL: cldr.Currency{DisplayName: "馬爾他里拉", Symbol: "MTL"},
				currency.MTP: cldr.Currency{DisplayName: "馬爾他鎊", Symbol: "MTP"},
				currency.MUR: cldr.Currency{DisplayName: "模里西斯盧比", Symbol: "MUR"},
				currency.MVP: cldr.Currency{DisplayName: "馬爾地夫盧比", Symbol: "MVP"},
				currency.MVR: cldr.Currency{DisplayName: "馬爾地夫盧非亞", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "馬拉維克瓦查", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "墨西哥披索", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "墨西哥銀披索 (1861–1992)", Symbol: "MXP"},
				currency.MXV: cldr.Currency{DisplayName: "墨西哥轉換單位 (UDI)", Symbol: "MXV"},
				currency.MYR: cldr.Currency{DisplayName: "馬來西亞令吉", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "莫三比克埃斯庫多", Symbol: "MZE"},
				currency.MZM: cldr.Currency{DisplayName: "莫三比克梅蒂卡爾 (1980–2006)", Symbol: "MZM"},
				currency.MZN: cldr.Currency{DisplayName: "莫三比克梅蒂卡爾", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "納米比亞元", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "奈及利亞奈拉", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "尼加拉瓜科多巴", Symbol: "NIC"},
				currency.NIO: cldr.Currency{DisplayName: "尼加拉瓜金科多巴", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "荷蘭盾", Symbol: "NLG"},
				currency.NOK: cldr.Currency{DisplayName: "挪威克朗", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "尼泊爾盧比", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "紐西蘭幣", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "阿曼里亞爾", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "巴拿馬巴波亞", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "祕魯因蒂", Symbol: "PEI"},
				currency.PEN: cldr.Currency{DisplayName: "秘魯太陽幣", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "秘魯索爾 (1863–1965)", Symbol: "PES"},
				currency.PGK: cldr.Currency{DisplayName: "巴布亞紐幾內亞基那", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "菲律賓披索", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "巴基斯坦盧比", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "波蘭茲羅提", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "波蘭茲羅提 (1950–1995)", Symbol: "PLZ"},
				currency.PTE: cldr.Currency{DisplayName: "葡萄牙埃斯庫多", Symbol: "PTE"},
				currency.PYG: cldr.Currency{DisplayName: "巴拉圭瓜拉尼", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "卡達里亞爾", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "羅德西亞元", Symbol: "RHD"},
				currency.ROL: cldr.Currency{DisplayName: "舊羅馬尼亞列伊", Symbol: "ROL"},
				currency.RON: cldr.Currency{DisplayName: "羅馬尼亞列伊", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "塞爾維亞戴納", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "俄羅斯盧布", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "俄羅斯盧布 (1991–1998)", Symbol: "RUR"},
				currency.RWF: cldr.Currency{DisplayName: "盧安達法郎", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "沙烏地里亞爾", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "索羅門群島元", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "塞席爾盧比", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "蘇丹第納爾", Symbol: "SDD"},
				currency.SDG: cldr.Currency{DisplayName: "蘇丹鎊", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "舊蘇丹鎊", Symbol: "SDP"},
				currency.SEK: cldr.Currency{DisplayName: "瑞典克朗", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "新加坡幣", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "聖赫勒拿鎊", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "斯洛維尼亞托勒", Symbol: "SIT"},
				currency.SKK: cldr.Currency{DisplayName: "斯洛伐克克朗", Symbol: "SKK"},
				currency.SLL: cldr.Currency{DisplayName: "獅子山利昂", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "索馬利亞先令", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "蘇利南元", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "蘇利南基爾", Symbol: "SRG"},
				currency.SSP: cldr.Currency{DisplayName: "南蘇丹鎊", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "聖多美島和普林西比島多布拉 (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "聖多美島和普林西比島多布拉", Symbol: "Db"},
				currency.SUR: cldr.Currency{DisplayName: "蘇聯盧布", Symbol: "SUR"},
				currency.SVC: cldr.Currency{DisplayName: "薩爾瓦多科郎", Symbol: "SVC"},
				currency.SYP: cldr.Currency{DisplayName: "敘利亞鎊", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "史瓦濟蘭里朗吉尼", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "泰銖", Symbol: "THB"},
				currency.TJR: cldr.Currency{DisplayName: "塔吉克盧布", Symbol: "TJR"},
				currency.TJS: cldr.Currency{DisplayName: "塔吉克索莫尼", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "土庫曼馬納特 (1993–2009)", Symbol: "TMM"},
				currency.TMT: cldr.Currency{DisplayName: "土庫曼馬納特", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "突尼西亞第納爾", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "東加潘加", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "帝汶埃斯庫多", Symbol: "TPE"},
				currency.TRL: cldr.Currency{DisplayName: "土耳其里拉", Symbol: "TRL"},
				currency.TRY: cldr.Currency{DisplayName: "新土耳其里拉", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "千里達及托巴哥元", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "新台幣", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "坦尚尼亞先令", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "烏克蘭格里夫納", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "烏克蘭卡本瓦那茲", Symbol: "UAK"},
				currency.UGS: cldr.Currency{DisplayName: "烏干達先令 (1966–1987)", Symbol: "UGS"},
				currency.UGX: cldr.Currency{DisplayName: "烏干達先令", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "美元", Symbol: "US$"},
				currency.USN: cldr.Currency{DisplayName: "美元（次日）", Symbol: "USN"},
				currency.USS: cldr.Currency{DisplayName: "美元（當日）", Symbol: "USS"},
				currency.UYI: cldr.Currency{DisplayName: "烏拉圭披索（指數單位）", Symbol: "UYI"},
				currency.UYP: cldr.Currency{DisplayName: "烏拉圭披索 (1975–1993)", Symbol: "UYP"},
				currency.UYU: cldr.Currency{DisplayName: "烏拉圭披索", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "烏茲別克索姆", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "委內瑞拉玻利瓦 (1871–2008)", Symbol: "VEB"},
				currency.VEF: cldr.Currency{DisplayName: "委內瑞拉玻利瓦 (VEF)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "委內瑞拉玻利瓦", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "越南盾", Symbol: "₫"},
				currency.VNN: cldr.Currency{DisplayName: "越南盾 (1978–1985)", Symbol: "VNN"},
				currency.VUV: cldr.Currency{DisplayName: "萬那杜瓦圖", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "西薩摩亞塔拉", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "法郎 (CFA–BEAC)", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "白銀", Symbol: "XAG"},
				currency.XAU: cldr.Currency{DisplayName: "黃金", Symbol: "XAU"},
				currency.XBA: cldr.Currency{DisplayName: "歐洲綜合單位", Symbol: "XBA"},
				currency.XBB: cldr.Currency{DisplayName: "歐洲貨幣單位 (XBB)", Symbol: "XBB"},
				currency.XBC: cldr.Currency{DisplayName: "歐洲會計單位 (XBC)", Symbol: "XBC"},
				currency.XBD: cldr.Currency{DisplayName: "歐洲會計單位 (XBD)", Symbol: "XBD"},
				currency.XCD: cldr.Currency{DisplayName: "格瑞那達元", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "特殊提款權", Symbol: "XDR"},
				currency.XEU: cldr.Currency{DisplayName: "歐洲貨幣單位 (XEU)", Symbol: "XEU"},
				currency.XFO: cldr.Currency{DisplayName: "法國金法郎", Symbol: "XFO"},
				currency.XFU: cldr.Currency{DisplayName: "法國法郎 (UIC)", Symbol: "XFU"},
				currency.XOF: cldr.Currency{DisplayName: "法郎 (CFA–BCEAO)", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "帕拉狄昂", Symbol: "XPD"},
				currency.XPF: cldr.Currency{DisplayName: "法郎 (CFP)", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "白金", Symbol: "XPT"},
				currency.XRE: cldr.Currency{DisplayName: "RINET 基金", Symbol: "XRE"},
				currency.XSU: cldr.Currency{DisplayName: "蘇克雷貨幣", Symbol: "XSU"},
				currency.XTS: cldr.Currency{DisplayName: "測試用貨幣代碼", Symbol: "XTS"},
				currency.XUA: cldr.Currency{DisplayName: "亞洲開發銀行計價單位", Symbol: "XUA"},
				currency.XXX: cldr.Currency{DisplayName: "未知貨幣", Symbol: "XXX"},
				currency.YDD: cldr.Currency{DisplayName: "葉門第納爾", Symbol: "YDD"},
				currency.YER: cldr.Currency{DisplayName: "葉門里亞爾", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "南斯拉夫第納爾硬幣", Symbol: "YUD"},
				currency.YUM: cldr.Currency{DisplayName: "南斯拉夫挪威亞第納爾", Symbol: "YUM"},
				currency.YUN: cldr.Currency{DisplayName: "南斯拉夫可轉換第納爾", Symbol: "YUN"},
				currency.YUR: cldr.Currency{DisplayName: "南斯拉夫改革第納爾 (1992–1993)", Symbol: "YUR"},
				currency.ZAL: cldr.Currency{DisplayName: "南非蘭特（金融）", Symbol: "ZAL"},
				currency.ZAR: cldr.Currency{DisplayName: "南非蘭特", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "尚比亞克瓦查 (1968–2012)", Symbol: "ZMK"},
				currency.ZMW: cldr.Currency{DisplayName: "尚比亞克瓦查", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "薩伊新扎伊爾", Symbol: "ZRN"},
				currency.ZRZ: cldr.Currency{DisplayName: "薩伊扎伊爾", Symbol: "ZRZ"},
				currency.ZWD: cldr.Currency{DisplayName: "辛巴威元 (1980–2008)", Symbol: "ZWD"},
				currency.ZWL: cldr.Currency{DisplayName: "辛巴威元 (2009)", Symbol: "ZWL"},
				currency.ZWR: cldr.Currency{DisplayName: "辛巴威元 (2008)", Symbol: "ZWR"},
			},
		},
	}
}
