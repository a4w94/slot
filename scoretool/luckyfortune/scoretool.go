package scoretools

import (
	info "info/luckyfortune"
	table "table/luckyfortune"
)

///基本投注額Bet///
const bet int = info.Bet

///Reel 轉輪數///
const reelamount int = info.Reelamount

///Column 橫排數///
const col int = info.Col

///派彩表///
var paytable [][]int = table.PublicTablefunc().PayTable

///倍數表///
var multiple = info.Multiple

///WILD 代號///
const wild = info.Wild

///Scatter 代號///
const scatter = info.Scatter

///計算分數落在哪個區間///
func Multiplejudge(payoff int) int {
	div := payoff / bet
	var tmp int
	if payoff == 0 {
		tmp = 0
	} else {
		for i := 1; i < len(multiple); i++ {
			if div >= multiple[i-1] && div < multiple[i] {
				tmp = i

			} else if div >= multiple[len(multiple)-1] {
				tmp = len(multiple)
			}
		}
	}

	return tmp
}

///計算scatter分數／／／

func Scatterpayf(costlevel int, scattermap int) int {
	var scatterpay int

	scatterpaytable := [][]int{{3, 4, 5}, {5, 10, 50}}

	for i := 0; i < 3; i++ {
		if scattermap == scatterpaytable[0][i] {
			scatterpay = scatterpaytable[1][i] * costlevel

		}
	}
	return scatterpay
}

///  計算combo數與回傳SYMBOL///
func CombojudgeLineGame(resultline [reelamount]int) (int, int) {

	comboquantity := 1

	var symbokind int

	if resultline[0] == wild {
		comboquantity = 0

		for i := 0; i < reelamount; i++ {
			if resultline[i] == wild {
				symbokind = wild
				continue
			} else if resultline[i] != wild {
				symbokind = resultline[i]

				break
			}

		}
		for i := 0; i < reelamount; i++ {
			if symbokind == resultline[i] || resultline[i] == wild {
				comboquantity = comboquantity + 1
			} else {
				break
			}
		}
	} else {
		symbokind = resultline[0]

		for i := 1; i < reelamount; i++ {
			if resultline[0] == resultline[i] || resultline[i] == wild {
				comboquantity = comboquantity + 1

			} else {
				break
			}
		}

	}

	var payresult int
	payresult = paytable[symbokind][comboquantity]

	wildquantity := 1
	var wildpayresult int

	if resultline[0] == 1 {
		for i := 1; i < reelamount; i++ {
			if resultline[0] == resultline[i] {
				wildquantity = wildquantity + 1

			} else {
				break
			}
		}
		wildpayresult = paytable[wild][wildquantity]

	}
	var realcombo int
	var realsymbo int

	if wildpayresult > payresult {
		realsymbo = wild
		realcombo = wildquantity
	} else {

		realsymbo = symbokind
		realcombo = comboquantity

	}
	return realsymbo, realcombo

}

func CombojudgeWayGame(result [col][reelamount]int) [col][3]int {
	var symbol int
	var resulttable [col][3]int

	for i := 0; i < col; i++ {

		symbol = result[i][0]
		var combo int
		linequantity := 1
		for j := 1; j < reelamount; j++ {
			eachquantity := 0
			for k := 0; k < col; k++ {
				if result[k][j] == symbol || result[k][j] == wild {
					eachquantity = eachquantity + 1
				}
			}

			if eachquantity == 0 {
				linequantity = linequantity * 1
			} else {
				linequantity = linequantity * eachquantity
			}

			if eachquantity == 0 {
				combo = j
				break
			} else {
				combo = 5
			}
		}

		resulttable[i][0] = symbol
		resulttable[i][1] = combo
		resulttable[i][2] = linequantity

	}
	return resulttable
}
