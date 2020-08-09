// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_ko_KP() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y년 M월 d일 EEEE", Long: "y년 M월 d일", Medium: "y. M. d.", Short: "yy. M. d."}, Time: cldr.CalendarDateFormat{Full: "a h시 m분 s초 zzzz", Long: "a h시 m분 s초 z", Medium: "a h:mm:ss", Short: "a h:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "1월", Feb: "2월", Mar: "3월", Apr: "4월", May: "5월", Jun: "6월", Jul: "7월", Aug: "8월", Sep: "9월", Oct: "10월", Nov: "11월", Dec: "12월"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1월", Feb: "2월", Mar: "3월", Apr: "4월", May: "5월", Jun: "6월", Jul: "7월", Aug: "8월", Sep: "9월", Oct: "10월", Nov: "11월", Dec: "12월"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "1월", Feb: "2월", Mar: "3월", Apr: "4월", May: "5월", Jun: "6월", Jul: "7월", Aug: "8월", Sep: "9월", Oct: "10월", Nov: "11월", Dec: "12월"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "일", Mon: "월", Tue: "화", Wed: "수", Thu: "목", Fri: "금", Sat: "토"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "일", Mon: "월", Tue: "화", Wed: "수", Thu: "목", Fri: "금", Sat: "토"}, Short: cldr.CalendarDayFormatNameValue{Sun: "일", Mon: "월", Tue: "화", Wed: "수", Thu: "목", Fri: "금", Sat: "토"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "일요일", Mon: "월요일", Tue: "화요일", Wed: "수요일", Thu: "목요일", Fri: "금요일", Sat: "토요일"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "오전", PM: "오후"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_ko]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "", Group: "", Negative: "\u200f-", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "안도라 페세타", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "아랍에미리트 디르함", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "아프가니 (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "아프가니스탄 아프가니", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "알바니아 레크", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "아르메니아 드람", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "네덜란드령 안틸레스 길더", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "앙골라 콴자", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "앙골라 콴자 (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "앙골라 신콴자 (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "앙골라 재조정 콴자 (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "아르헨티나 오스트랄", Symbol: ""},
				currency.ARL: cldr.Currency{DisplayName: "아르헨티나 페소 레이 (1970–1983)", Symbol: ""},
				currency.ARM: cldr.Currency{DisplayName: "아르헨티나 페소 (18810–1970)", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "아르헨티나 페소 (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "아르헨티나 페소", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "호주 실링", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "호주 달러", Symbol: "AU$"},
				currency.AWG: cldr.Currency{DisplayName: "아루바 플로린", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "아제르바이젠 마나트(1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "아제르바이잔 마나트", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "보스니아-헤르체고비나 디나르", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "보스니아-헤르체고비나 태환 마르크", Symbol: "BAM"},
				currency.BAN: cldr.Currency{DisplayName: "보스니아-헤르체고비나 신디나르 (1994–1997)", Symbol: ""},
				currency.BBD: cldr.Currency{DisplayName: "바베이도스 달러", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "방글라데시 타카", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "벨기에 프랑 (태환)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "벨기에 프랑", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "벨기에 프랑 (금융)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "불가리아 동전 렛", Symbol: ""},
				currency.BGM: cldr.Currency{DisplayName: "불가리아 사회주의자 렛", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "불가리아 레프", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "불가리아 렛 (1879–1952)", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "바레인 디나르", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "부룬디 프랑", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "버뮤다 달러", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "부루나이 달러", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "볼리비아노", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "볼리비아 볼리비아노 (1863–1963)", Symbol: ""},
				currency.BOP: cldr.Currency{DisplayName: "볼리비아노 페소", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "볼리비아노 Mvdol(기금)", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "볼리비아노 크루제이루 노보 (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "브라질 크루자두", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "브라질 크루제이루 (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "브라질 레알", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "브라질 크루자두 노보", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "브라질 크루제이루", Symbol: ""},
				currency.BRZ: cldr.Currency{DisplayName: "브라질 크루제이루 (1942–1967)", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "바하마 달러", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "부탄 눌투눔", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "버마 차트", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "보츠와나 폴라", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "벨라루스 신권 루블 (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "벨라루스 루블", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "벨라루스 루블 (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "벨리즈 달러", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "캐나다 달러", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "콩고 프랑 콩골라스", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "유로 (WIR)", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "스위스 프랑", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "프랑 (WIR)", Symbol: ""},
				currency.CLE: cldr.Currency{DisplayName: "칠레 에스쿠도", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "칠레 (UF)", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "칠레 페소", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "중국 위안화(역외)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "중국 위안화", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "콜롬비아 페소", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "콜롬비아 실가 단위", Symbol: ""},
				currency.CRC: cldr.Currency{DisplayName: "코스타리카 콜론", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "고 세르비아 디나르", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "체코슬로바키아 동전 코루나", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "쿠바 태환 페소", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "쿠바 페소", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "카보베르데 에스쿠도", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "싸이프러스 파운드", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "체코 공화국 코루나", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "동독 오스트마르크", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "독일 마르크", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "지부티 프랑", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "덴마크 크로네", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "도미니카 페소", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "알제리 디나르", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "에쿠아도르 수크레", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "에콰도르 (UVC)", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "에스토니아 크룬", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "이집트 파운드", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "에리트리아 나크파", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "스페인 페세타(예금)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "스페인 페세타(변환 예금)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "스페인 페세타", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "에티오피아 비르", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "유로", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "핀란드 마르카", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "피지 달러", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "포클랜드제도 파운드", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "프랑스 프랑", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "파운드", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "그루지야 지폐 라리트", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "조지아 라리", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "가나 시디 (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "가나 시디", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "지브롤터 파운드", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "감비아 달라시", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "기니 프랑", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "기니 시리", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "적도 기니 에쿨 (Ekwele)", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "그리스 드라크마", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "과테말라 케트살", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "포르투갈령 기니 에스쿠도", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "기네비쏘 페소", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "가이아나 달러", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "홍콩 달러", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "온두라스 렘피라", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "크로아티아 디나르", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "크로아티아 쿠나", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "하이티 구르드", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "헝가리 포린트", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "인도네시아 루피아", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "아일랜드 파운드", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "이스라엘 파운드", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "이스라엘 신권 세켈", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "인도 루피", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "이라크 디나르", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "이란 리얄", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "아이슬란드 크로나", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "이탈리아 리라", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "자메이카 달러", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "요르단 디나르", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "일본 엔화", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "케냐 실링", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "키르기스스탄 솜", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "캄보디아 리얄", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "코모르 프랑", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "조선 민주주의 인민 공화국 원", Symbol: "KPW"},
				currency.KRH: cldr.Currency{DisplayName: "대한민국 환 (1953–1962)", Symbol: ""},
				currency.KRW: cldr.Currency{DisplayName: "대한민국 원", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "쿠웨이트 디나르", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "케이맨 제도 달러", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "카자흐스탄 텐게", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "라오스 키프", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "레바논 파운드", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "스리랑카 루피", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "라이베리아 달러", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "레소토 로티", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "리투아니아 리타", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "룩셈부르크 타로나", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "룩셈부르크 변환 프랑", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "룩셈부르크 프랑", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "룩셈부르크 재정 프랑", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "라트비아 라트", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "라트비아 루블", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "리비아 디나르", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "모로코 디렘", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "모로코 프랑", Symbol: ""},
				currency.MCF: cldr.Currency{DisplayName: "모나코 프랑", Symbol: ""},
				currency.MDC: cldr.Currency{DisplayName: "몰도바 쿠폰", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "몰도바 레이", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "마다가스카르 아리아리", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "마다가스카르 프랑", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "마케도니아 디나르", Symbol: "MKD"},
				currency.MLF: cldr.Currency{DisplayName: "말리 프랑", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "미얀마 키얏", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "몽골 투그릭", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "마카오 파타카", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "모리타니 우기야 (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "모리타니 우기야", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "몰타 리라", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "몰타 파운드", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "모리셔스 루피", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "몰디브 제도 루피아", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "말라위 콰쳐", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "멕시코 페소", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "멕시코 실버 페소 (1861–1992)", Symbol: ""},
				currency.MXV: cldr.Currency{DisplayName: "멕시코 (UDI)", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "말레이시아 링깃", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "모잠비크 에스쿠도", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "고 모잠비크 메티칼", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "모잠비크 메티칼", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "나미비아 달러", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "니제르 나이라", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "니카라과 코르도바", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "니카라과 코르도바 오로", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "네델란드 길더", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "노르웨이 크로네", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "네팔 루피", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "뉴질랜드 달러", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "오만 리얄", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "파나마 발보아", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "페루 인티", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "페루 솔", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "페루 솔 (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "파푸아뉴기니 키나", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "필리핀 페소", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "파키스탄 루피", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "폴란드 즐로티", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "폴란드 즐로티 (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "포르투갈 에스쿠도", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "파라과이 과라니", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "카타르 리얄", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "로디지아 달러", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "루마니아 레이", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "루마니아 레우", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "세르비아 디나르", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "러시아 루블", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "러시아 루블 (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "르완다 프랑", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "사우디아라비아 리얄", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "솔로몬 제도 달러", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "세이셸 루피", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "수단 디나르", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "수단 파운드", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "고 수단 파운드", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "스웨덴 크로나", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "싱가폴 달러", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "세인트헬레나 파운드", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "슬로베니아 톨라르", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "슬로바키아 코루나", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "시에라리온 리온", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "소말리아 실링", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "수리남 달러", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "수리남 길더", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "남수단 파운드", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "상투메 프린시페 도브라 (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "상투메 프린시페 도브라", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "소련 루블", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "엘살바도르 콜론", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "시리아 파운드", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "스와질란드 릴랑게니", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "태국 바트", Symbol: "THB"},
				currency.TJR: cldr.Currency{DisplayName: "타지키스탄 루블", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "타지키스탄 소모니", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "투르크메니스탄 마나트 (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "투르크메니스탄 마나트", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "튀니지 디나르", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "통가 파앙가", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "티모르 에스쿠도", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "터키 리라", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "신 터키 리라", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "트리니다드 토바고 달러", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "신 타이완 달러", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "탄자니아 실링", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "우크라이나 그리브나", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "우크라이나 카보바네츠", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "우간다 실링 (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "우간다 실링", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "미국 달러", Symbol: "US$"},
				currency.USN: cldr.Currency{DisplayName: "미국 달러(다음날)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "미국 달러(당일)", Symbol: ""},
				currency.UYI: cldr.Currency{DisplayName: "우루과이 페소 (UI)", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "우루과이 페소 (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "우루과이 페소 우루과요", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "우즈베키스탄 숨", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "베네주엘라 볼리바르 (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "베네수엘라 볼리바르 (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "베네수엘라 볼리바르", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "베트남 동", Symbol: "₫"},
				currency.VNN: cldr.Currency{DisplayName: "베트남 동 (1978–1985)", Symbol: ""},
				currency.VUV: cldr.Currency{DisplayName: "바누아투 바투", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "서 사모아 탈라", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "중앙아프리카 CFA 프랑", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "은화", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "금", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "유르코 (유럽 회계 단위)", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "유럽 통화 동맹", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "유럽 계산 단위 (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "유럽 계산 단위 (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "동카리브 달러", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "특별인출권", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "유럽 환율 단위", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "프랑스 프랑 (Gold)", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "프랑스 프랑 (UIC)", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "서아프리카 CFA 프랑", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "팔라듐", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "CFP 프랑", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "백금", Symbol: ""},
				currency.XRE: cldr.Currency{DisplayName: "RINET 기금", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "테스트 통화 코드", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "알 수 없는 통화 단위", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "예멘 디나르", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "예멘 리알", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "유고슬라비아 동전 디나르", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "유고슬라비아 노비 디나르", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "유고슬라비아 전환 디나르", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "남아프리카 랜드 (금융)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "남아프리카 랜드", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "쟘비아 콰쳐 (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "잠비아 콰쳐", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "자이르 신권 자이르", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "자이르 자이르", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "짐바브웨 달러", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "짐바브웨 달러 (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "짐바브웨 달러 (2008)", Symbol: ""},
			},
		},
	}
}
