package calculate

import (
	"fmt"
	info "info/manyfumany"
	"math"
	"math/rand"
	rngtool "rngtools/manyfumany"
	scoretool "scoretool/manyfumany"
	table "table/manyfumany"
	"time"
)

type TotalResult struct {
	RTP        float32
	Totalpay   int
	NGtotalpay int

	NGScattertotalpay   int
	NGScatterRTP        float32
	NGRTP               float32
	NGMultiplestatisics [8]int
	NGMultiplefrequency []float32
	FGtotalpay          int
	FGRTP               float32
	FGMultiplestatisics [8]int
	FGMultiplefrequency []float32
	Totalsession        int
	Totalcombo          [][]int
	NGFGcombo           EachRoundResult
	RetriggerTimes      int
	TotalVariance       float64
	TotalStd            float64
	Costtime            interface{}
}

type EachRoundResult struct {
	NGresult [][]int
	//NGresultfrequency [linenum][comboresultnum]float32
	NGpay           int
	FGresult        [][]int
	FGpay           int
	Scatterpay      int
	EaceRetrigger   int
	NGVariance      float64
	FGVariance      float64
	ScatterVariance float64
	NGMultiple      int
	FGMultiple      int
	FreeGameStatus  bool
}

type EachRoundFGResult struct {
	FGEachresult [][]int
	FGEachpay    int

	FGVariance float64
}

type EachRoundBonusResult struct {
	BGEachresult [][]int
	BGEachpay    int
}

const gamename string = info.GameName

///Reel 轉輪數///
const reelamount int = info.Reelamount

///Column 橫排數///
const col int = info.Col

///基本投注額Bet///
const bet int = info.Bet

///free game 贈送場數／／／
const fgsession int = info.FGsession

///獎圖總數///
const symbolamount = info.Symbolamount

///combo數目0~5combo///
const comboresultnum = info.Comboresultnum

///WILD 代號///
const wild = info.Wild

///Scatter 代號///
const scatter = info.Scatter

///派彩表///
var paytable [][]int = table.PublicTablefunc().PayTable

///理論值ＲＴＰ///
var thmRTP = info.ThmRTP

///倍數統計///
var multiplestatisics = table.PublicTablefunc().MultipleTalbe

var ng = table.NGTablefunc()
var fg = table.FGTablefunc()

func costtable(costlevel int, rtp int) ([][]int, [][]int) {
	var ngstritable [][]int
	var fgstritable [][]int
	switch costlevel {
	case 8:
		switch rtp {
		case 95:
			ngstritable = ng.NGStriTable8rtp95
			fgstritable = fg.FGStriTable8rtp95
		case 98:
			ngstritable = ng.NGStriTable8rtp98
			fgstritable = fg.FGStriTable8rtp98
		}
	case 18:
		switch rtp {
		case 95:
			ngstritable = ng.NGStriTable18rtp95
			fgstritable = fg.FGStriTable18rtp95
		case 98:
			ngstritable = ng.NGStriTable18rtp98
			fgstritable = fg.FGStriTable18rtp98
		}
	case 38:
		switch rtp {
		case 95:
			ngstritable = ng.NGStriTable38rtp95
			fgstritable = fg.FGStriTable38rtp95
		case 98:
			ngstritable = ng.NGStriTable38rtp98
			fgstritable = fg.FGStriTable38rtp98
		}
	case 68:
		switch rtp {
		case 95:
			ngstritable = ng.NGStriTable68rtp95
			fgstritable = fg.FGStriTable68rtp95
		case 98:
			ngstritable = ng.NGStriTable68rtp98
			fgstritable = fg.FGStriTable68rtp98
		}
	case 88:
		switch rtp {
		case 95:
			ngstritable = ng.NGStriTable88rtp95
			fgstritable = fg.FGStriTable88rtp95
		case 98:
			ngstritable = ng.NGStriTable88rtp98
			fgstritable = fg.FGStriTable88rtp98
		}

	}
	return ngstritable, fgstritable

}

func Run(session int, runtime int, costlevel int, rtp int) {

	for i := 0; i < runtime; i++ {
		a := Totalresult(session, costlevel, rtp)

		fmt.Println("遊戲名稱 :", gamename)
		fmt.Println("等級 :", costlevel)
		fmt.Println("理論RTP :", rtp)
		fmt.Println("試驗次數 ：", session)
		fmt.Println("總派彩 :", a.Totalpay)

		fmt.Println("ＲＴＰ :", a.RTP)

		fmt.Println()
		fmt.Println("XXXXXXXXXXXXXXXXX細項XXXXXXXXXXXXXXXXXXXXX")
		fmt.Println()

		fmt.Println("NG派彩 :", a.NGtotalpay)
		fmt.Println("NGRTP :", a.NGRTP)
		//fmt.Println("NGcombo頻率", a.NGFGcombo.NGresultfrequency)
		fmt.Println("NGsctterpay :", a.NGScattertotalpay)
		fmt.Println("NGscatter RTP :", a.NGScatterRTP)
		fmt.Println("FG派彩 :", a.FGtotalpay)
		//fmt.Println("FG獨立ＲＴＰ :", float32(a.FGtotalpay)/float32(a.RetriggerTimes))
		fmt.Println("FGRTP :", a.FGRTP)

		fmt.Println(info.Multiple)
		fmt.Println(a.NGMultiplestatisics)
		fmt.Println("NG倍數統計")
		for time := 0; time < len(info.Multiple); time++ {
			if time == len(info.Multiple)-1 {
				fmt.Println("倍數區間 ：  ", info.Multiple[time], "以上")
				fmt.Println("NG倍數統計 ：", a.NGMultiplestatisics[time])
				fmt.Println("NG倍數頻率 ：", a.NGMultiplefrequency[time])
				fmt.Println()
			} else {

				fmt.Println("倍數區間 ：  ", info.Multiple[time], "~", info.Multiple[time+1])
				fmt.Println("NG倍數統計 ：", a.NGMultiplestatisics[time])
				fmt.Println("NG倍數頻率 ：", a.NGMultiplefrequency[time])
				fmt.Println()
			}
		}

		fmt.Println()
		fmt.Println("FG倍數統計")
		for time := 0; time < len(info.Multiple); time++ {
			if time == len(info.Multiple)-1 {
				fmt.Println("倍數區間 ：  ", info.Multiple[time], "以上")
				fmt.Println("FG倍數統計 ：", a.FGMultiplestatisics[time])
				fmt.Println("FG倍數頻率 ：", a.FGMultiplefrequency[time])
				fmt.Println()
			} else {
				fmt.Println("倍數區間 ：  ", info.Multiple[time], "~", info.Multiple[time+1])
				fmt.Println("FG倍數統計 ：", a.FGMultiplestatisics[time])
				fmt.Println("FG倍數頻率 ：", a.FGMultiplefrequency[time])
				fmt.Println()
			}
		}

		fmt.Println("變異數 :", a.TotalVariance)
		fmt.Println("標準差 ：", a.TotalStd)
		fmt.Println("花費時間 ：", a.Costtime)
		fmt.Println("freegame機率", float32(a.RetriggerTimes)/float32(session))
		// var ngcombofrequency [symbolamount][comboresultnum]float32
		// var fgcombofrequency [symbolamount][comboresultnum]float32

		// for i := 0; i < symbolamount; i++ {
		// 	for k := 0; k < comboresultnum; k++ {
		// 		ngcombofrequency[i][k] = float32(a.NGFGcombo.NGresult[i][k]) / float32(session)
		// 		fgcombofrequency[i][k] = float32(a.NGFGcombo.FGresult[i][k]) / float32(a.RetriggerTimes)

		// 	}
		// }
		// // fmt.Println(a.NGFGcombo.NGresult)
		// // fmt.Println(ngcombofrequency)
		// fmt.Println(fgcombofrequency)

		fmt.Println()
		fmt.Println("==============================================================================================================================================================================================================================")

	}

}

func Totalresult(session int, costlevel int, rtp int) TotalResult {
	s := time.Now()
	rand.Seed(int64(time.Now().UnixNano()))
	var Total = TotalResult{}
	Total.NGFGcombo.NGresult = table.NGTablefunc().NGComboStatistic
	Total.NGFGcombo.FGresult = table.FGTablefunc().FGComboStatistic
	//Total.Totalcombo = table.PublicTablefunc().TotalTable

	for k := 0; k < session; k++ {
		//Total.Totalpay += eachroundresult().FGpay + eachroundresult().NGpay + eachroundresult().Scatterpay

		a := eachroundresult(costlevel, rtp)
		///計算總派彩 ＮＧ總派採 ＦＧ總派彩 Ｓcatter總派彩 總變異數 總觸發ＦＧ次數 ＮＧＦＧ區間統計///
		Total.Totalpay += a.NGpay + a.FGpay + a.Scatterpay
		Total.NGtotalpay += a.NGpay + a.Scatterpay
		Total.FGtotalpay += a.FGpay
		Total.NGScattertotalpay += a.Scatterpay
		Total.TotalVariance += float64(a.NGVariance+a.ScatterVariance+a.FGVariance) / float64(session)
		Total.RetriggerTimes += a.EaceRetrigger
		Total.NGMultiplestatisics[a.NGMultiple]++
		if a.FreeGameStatus == true {
			Total.FGMultiplestatisics[a.FGMultiple]++
		}

		// for i := 0; i < symbolamount; i++ {
		// 	for k := 0; k < comboresultnum; k++ {
		// 		Total.NGFGcombo.NGresult[i][k] += a.NGresult[i][k]
		// 		Total.NGFGcombo.FGresult[i][k] += a.FGresult[i][k]
		// 	}
		// }

	}

	///計算總ＲＴＰ，ＮＧＲＴＰ，ＦＧＲＴＰ///
	Total.RTP = float32(Total.Totalpay) / (float32(session) * float32(costlevel))
	Total.NGRTP = float32(Total.NGtotalpay) / (float32(session) * float32(costlevel))
	Total.FGRTP = float32(Total.FGtotalpay) / (float32(session) * float32(costlevel))
	Total.NGScatterRTP = float32(Total.NGScattertotalpay) / (float32(session) * float32(costlevel))

	Total.Totalsession = session
	///計算ＮＧＦＧ標準差／／／
	Total.TotalStd = math.Sqrt(Total.TotalVariance)
	for i := 0; i < len(Total.NGMultiplestatisics); i++ {
		Total.NGMultiplefrequency = append(Total.NGMultiplefrequency, float32(Total.NGMultiplestatisics[i])/float32(session))
		Total.FGMultiplefrequency = append(Total.FGMultiplefrequency, float32(Total.FGMultiplestatisics[i])/float32(Total.RetriggerTimes))
	}
	// for i := 0; i < symbolamount; i++ {
	// 	for k := 0; k < comboresultnum; k++ {
	// 		Total.NGFGcombo.NGresultfrequency[i][k] = float32(Total.NGFGcombo.NGresult[i][k])
	// 		Total.NGFGcombo.NGresultfrequency[i][k] = Total.NGFGcombo.NGresultfrequency[i][k] / float32(session)
	// 	}
	// }

	Total.Costtime = time.Since(s)

	return Total

}

///計算各輪結果///
func eachroundresult(costlevel int, rtp int) EachRoundResult {
	var Eachresult = EachRoundResult{}

	///計算ＮＧcombo///
	Eachresult.NGresult = table.NGTablefunc().NGComboStatistic
	Eachresult.FGresult = table.FGTablefunc().FGComboStatistic

	ngresult := rngtool.BaseGameRng(costlevel, rtp)

	///工程盤面///
	//ngresult = [3][5]int{{16, 16, 16, 2, 2}, {5, 7, 7, 17, 6}, {6, 5, 5, 5, 5}}

	resultmap := rngtool.Resultmapf(ngresult)

	///判斷進FG,並計分///
	if resultmap > 2 {

		Eachresult.FreeGameStatus = true

		a := fgresult(costlevel, rtp)

		Eachresult.EaceRetrigger++

		Eachresult.FGresult = a.FGEachresult
		Eachresult.FGpay += a.FGEachpay
		Eachresult.FGMultiple = scoretool.Multiplejudge(Eachresult.FGpay)

		Eachresult.FGVariance += a.FGVariance
		scatterpay := scoretool.Scatterpayf(costlevel, resultmap)
		Eachresult.Scatterpay += scatterpay
		Eachresult.ScatterVariance += math.Pow((float64(Eachresult.Scatterpay) - float64(thmRTP)), 2)
	}

	symbocombototal := scoretool.CombojudgeWayGame(ngresult)
	///陣列combo累加///
	for k := 0; k < col; k++ {
		Eachresult.NGresult[symbocombototal[k][0]][symbocombototal[k][1]] += symbocombototal[k][2]
		///計分///
		Eachresult.NGpay += paytable[symbocombototal[k][0]][symbocombototal[k][1]] * symbocombototal[k][2]
	}

	Eachresult.NGVariance += math.Pow((float64(Eachresult.NGpay)/float64(costlevel) - float64(thmRTP)), 2)
	Eachresult.NGMultiple = scoretool.Multiplejudge(Eachresult.NGpay)

	return Eachresult

}

///FG count payoff
func fgresult(costlevel int, rtp int) EachRoundFGResult {
	var result = EachRoundFGResult{}
	result.FGEachresult = table.FGTablefunc().FGComboStatistic

	//var wildposition []int

	totalsession := fgsession
	retrigger := 1

	for i := 0; i < totalsession; i++ {
		freeresult := rngtool.FreeGameRng(costlevel, rtp)
		// freeresult = [3][5]int{
		// 	{16, 4, 16, 4, 4}, {5, 5, 5, 5, 5}, {3, 16, 4, 2, 8},
		// }

		resultmap := rngtool.Resultmapf(freeresult)

		if resultmap > 2 && retrigger < 10 {

			totalsession += fgsession
			retrigger++

		}
		if resultmap > 2 {
			freescatterpay := scoretool.Scatterpayf(costlevel, resultmap)
			result.FGEachpay += freescatterpay
		}
		symbocombototal := scoretool.CombojudgeWayGame(freeresult)
		for k := 0; k < col; k++ {
			//result.FGEachresult[symbocombototal[k][0]][symbocombototal[k][1]] += symbocombototal[k][2]
			result.FGEachpay += paytable[symbocombototal[k][0]][symbocombototal[k][1]] * symbocombototal[k][2]
		}
	}
	result.FGVariance = math.Pow((float64(result.FGEachpay)/float64(costlevel) - float64(thmRTP)), 2)

	return result
}
