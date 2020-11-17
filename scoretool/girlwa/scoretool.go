package scoretools

import (
	info "info/girlwa"
	table "table/girlwa"
)

///基本投注額Bet///
const bet int = info.Bet

///Reel 轉輪數///
const reelamount int = info.Reelamount

///派彩表///
var paytable [][]int = table.PublicTablefunc().PayTable

///倍數表///
var multiple = info.Multiple

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

func Scatterpayf(scattermap map[int]int) int {
	var scatterpay int

	scatterpaytable := [][]int{{0, 3, 10}}

	for i := 0; i < len(scatterpaytable); i++ {
		if scattermap[0] == scatterpaytable[i][1] {
			scatterpay = scatterpaytable[i][2] * bet

		}
	}
	return scatterpay
}

///  計算combo數與回傳SYMBOL///
func Combojudge(resultline [reelamount]int) (int, int) {

	comboquantity := 1

	var symbokind int

	if resultline[0] == 1 {
		comboquantity = 0

		for i := 0; i < reelamount; i++ {
			if resultline[i] == 1 {
				symbokind = 1
				continue
			} else if resultline[i] != 1 {
				symbokind = resultline[i]

				break
			}

		}
		for i := 0; i < reelamount; i++ {
			if symbokind == resultline[i] || resultline[i] == 1 {
				comboquantity = comboquantity + 1
			} else {
				break
			}
		}
	} else {
		symbokind = resultline[0]

		for i := 1; i < reelamount; i++ {
			if resultline[0] == resultline[i] || resultline[i] == 1 {
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
		wildpayresult = paytable[1][wildquantity]

	}
	var realcombo int
	var realsymbo int

	if wildpayresult > payresult {
		realsymbo = 1
		realcombo = wildquantity
	} else {

		realsymbo = symbokind
		realcombo = comboquantity

	}
	return realsymbo, realcombo

}
