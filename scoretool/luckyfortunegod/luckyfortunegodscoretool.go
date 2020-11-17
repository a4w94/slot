package scoretools

import (
	info "info/luckyfortunegod"
	table "table/luckyfortunegod"
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
				tmp = i - 1

			} else if div >= multiple[len(multiple)-1] {
				tmp = len(multiple) - 1
			}
		}
	}
	//fmt.Println("分數", payoff, "倍數", div, "傳出tmp", tmp)

	return tmp
}

///計算scatter分數／／／

type Scatterinfo struct {
	Scatterpay int
	Fgsession  int
}

func Scatterpayf(costlevel int, scattermap int) Scatterinfo {
	var scatterinfo = Scatterinfo{}

	for i := 0; i < 3; i++ {
		if scattermap == info.Scatterinfo[0][i] {
			scatterinfo.Scatterpay = info.Scatterinfo[1][i] * costlevel
			scatterinfo.Fgsession = info.Scatterinfo[2][i]

		}
	}
	return scatterinfo
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
