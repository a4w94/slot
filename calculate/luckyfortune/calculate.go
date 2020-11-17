package calculate

import (
	"fmt"
	info "info/luckyfortune"
	"math/rand"
	rngtool "rngtools/luckyfortune"
	scoretool "scoretool/luckyfortune"
	table "table/luckyfortune"
	"time"
)

type TotalResult struct {
	RTP        float32
	Totalpay   int
	NGtotalpay int

	NGScattertotalpay   int
	NGScatterRTP        float32
	NGRTP               float32
	NGMultiplestatisics [7]int
	NGMultiplefrequency []float32
	FGtotalpay          int
	FGRTP               float32
	FGMultiplestatisics [7]int
	FGMultiplefrequency []float32
	Totalsession        int
	Totalcombo          [symbolamount][comboresultnum]int
	NGFGcombo           EachRoundResult
	RetriggerTimes      int
	TotalVariance       float64
	TotalStd            float64
	Costtime            interface{}
	Risk
}

type Risk struct {
	rngtool.Input
}

type EachRoundResult struct {
	NGresult [symbolamount][comboresultnum]int
	//NGresultfrequency [linenum][comboresultnum]float32
	NGpay           int
	FGresult        [symbolamount][comboresultnum]int
	FGpay           int
	Scatterpay      int
	EaceRetrigger   int
	NGVariance      float64
	FGVariance      float64
	ScatterVariance float64
	NGMultiple      int
	FGMultiple      int
	Risk
}
type Ngjudgepass struct {
	NGpay      int
	Scatterpay int
	Trigger    bool
	Risk
}

type Fgjudgepass struct {
	FGpay      int
	Scatterpay int
	Trigger    bool
	Risk
}

type EachRoundFGResult struct {
	FGEachresult [symbolamount][comboresultnum]int
	FGEachpay    int
	FGVariance   float64
	Risk
}

type EachRoundBonusResult struct {
	BGEachresult [symbolamount][comboresultnum]int
	BGEachpay    int
}

const (
	///遊戲名稱///
	gamename string = info.GameName

	///Reel 轉輪數///
	reelamount int = info.Reelamount

	///Column 橫排數///
	col int = info.Col

	///基本投注額Bet///
	bet int = info.Bet

	///free game 贈送場數／／／
	fgsession int = info.FGsession

	///獎圖總數///
	symbolamount = info.Symbolamount

	///combo數目0~5combo///
	comboresultnum = info.Comboresultnum

	///WILD 代號///
	wild = info.Wild

	///Scatter 代號///
	scatter = info.Scatter

	///風控///
	payofflimit             = info.Payofflimit
	levelupper              = info.LevelUpper
	initialaccumulatebet    = info.InitialAccumulatebet
	initialaccumulatepayoff = info.InitialAccumulatepayoff
	initialbet              = info.InitialBet
)

var (
	costlevel int

	Betmultiple int

	///派彩表///
	paytable [][]int = table.PublicTablefunc().PayTable

	///理論值ＲＴＰ///
	thmRTP = info.ThmRTP

	///倍數統計///
	multiplestatisics = table.PublicTablefunc().MultipleTalbe

	///下注等級///
	betlevel int

	ngstritable = [][]int{}
	fgstritable = [][]int{}

	Rtp95  int
	Rtp965 int
	Rtp99  int
)

func Run(session int, runtime int, costlevelinput int, rtp int) TotalResult {
	k := &costlevel
	*k = costlevelinput
	m := &Betmultiple
	*m = costlevelinput / bet
	var web = TotalResult{}
	for i := 0; i < runtime; i++ {
		a := Totalresult(session)

		// fmt.Println("遊戲名稱 :", gamename)
		// fmt.Println("下注額 :", costlevel)
		// fmt.Println("下注倍數 :", Betmultiple)

		// fmt.Println("理論RTP :", rtp)
		// fmt.Println("試驗次數 ：", session)
		// fmt.Println("總派彩 :", a.Totalpay)

		// fmt.Println("ＲＴＰ :", a.RTP)

		// fmt.Println()
		// fmt.Println("XXXXXXXXXXXXXXXXX細項XXXXXXXXXXXXXXXXXXXXX")
		// fmt.Println()

		// fmt.Println("NG派彩 :", a.NGtotalpay)
		// fmt.Println("NGRTP :", a.NGRTP)
		// //fmt.Println("NGcombo頻率", a.NGFGcombo.NGresultfrequency)
		// fmt.Println("NGsctterpay :", a.NGScattertotalpay)
		// fmt.Println("NGscatter RTP :", a.NGScatterRTP)
		// fmt.Println("FG派彩 :", a.FGtotalpay)
		// //fmt.Println("FG獨立ＲＴＰ :", float32(a.FGtotalpay)/float32(a.RetriggerTimes))
		// fmt.Println("FGRTP :", a.FGRTP)
		// //fmt.Println("倍數區間 ：  ", "0,~10,~20,~30,~40,~50,50~")
		// //fmt.Println("NG倍數統計 ：", a.NGMultiplestatisics)
		// //fmt.Println("NG倍數頻率 ：", a.NGMultiplefrequency)
		// //fmt.Println("FG倍數統計 ：", a.FGMultiplestatisics)
		// //fmt.Println("FG倍數頻率 ：", a.FGMultiplefrequency)
		// //fmt.Println("變異數 :", a.TotalVariance)
		// //fmt.Println("標準差 ：", a.TotalStd)
		// fmt.Println("花費時間 ：", a.Costtime)
		// fmt.Println("freegame機率", float32(a.RetriggerTimes)/float32(session))
		// // var ngcombofrequency [symbolamount][comboresultnum]float32
		// // var fgcombofrequency [symbolamount][comboresultnum]float32
		// fmt.Println("95,965,99", rtp95, rtp965, rtp99)
		// fmt.Println(a.Risk)
		// // for i := 0; i < symbolamount; i++ {
		// // 	for k := 0; k < comboresultnum; k++ {
		// // 		ngcombofrequency[i][k] = float32(a.NGFGcombo.NGresult[i][k]) / float32(session)
		// // 		fgcombofrequency[i][k] = float32(a.NGFGcombo.FGresult[i][k]) / float32(a.RetriggerTimes)

		// // 	}
		// // }
		// // // fmt.Println(a.NGFGcombo.NGresult)
		// // // fmt.Println(ngcombofrequency)
		// // fmt.Println(fgcombofrequency)

		fmt.Println()
		web = a
		//fmt.Println("==============================================================================================================================================================================================================================")

	}
	return web
}

func Totalresult(session int) TotalResult {
	s := time.Now()
	rand.Seed(int64(time.Now().UnixNano()))
	var Total = TotalResult{}
	Total.NGFGcombo.NGresult = table.NGTablefunc().NGComboStatistic
	Total.NGFGcombo.FGresult = table.FGTablefunc().FGComboStatistic
	//Total.Totalcombo = table.PublicTablefunc().TotalTable
	var RiskControl = Risk{}
	RiskControl.Accumulatebet = initialaccumulatebet
	RiskControl.Accumulatepayoff = initialaccumulatepayoff
	RiskControl.Bet = initialbet
	RiskControl.Currentrtp = 0.965
	for k := 0; k < session; k++ {

		//Total.Totalpay += eachroundresult().FGpay + eachroundresult().NGpay + eachroundresult().Scatterpay

		///風控流程///

		a := eachroundresult(RiskControl)

		// RiskControl.Accumulatebet = a.Input.Accumulatebet
		// RiskControl.Accumulatepayoff = a.Input.Accumulatepayoff
		// RiskControl.Currentrtp = a.Input.Currentrtp
		RiskControl = a.Risk
		//fmt.Println(RiskControl)

		///計算總派彩 ＮＧ總派採 ＦＧ總派彩 Ｓcatter總派彩 總變異數 總觸發ＦＧ次數 ＮＧＦＧ區間統計///

		Total.Totalpay += a.NGpay + a.FGpay + a.Scatterpay
		Total.NGtotalpay += a.NGpay + a.Scatterpay
		Total.FGtotalpay += a.FGpay
		Total.NGScattertotalpay += a.Scatterpay
		//Total.TotalVariance += float64(a.NGVariance+a.ScatterVariance+a.FGVariance) / float64(session)
		Total.RetriggerTimes += a.EaceRetrigger
		//Total.NGMultiplestatisics[a.NGMultiple]++
		//Total.FGMultiplestatisics[a.FGMultiple]++
		//fmt.Println()
		//fmt.Println(Total.NGtotalpay, RiskControl)
	}

	///計算總ＲＴＰ，ＮＧＲＴＰ，ＦＧＲＴＰ///
	Total.RTP = float32(Total.Totalpay) / (float32(session) * float32(costlevel))
	Total.NGRTP = float32(Total.NGtotalpay) / (float32(session) * float32(costlevel))
	Total.FGRTP = float32(Total.FGtotalpay) / (float32(session) * float32(costlevel))
	Total.NGScatterRTP = float32(Total.NGScattertotalpay) / (float32(session) * float32(costlevel))

	Total.Totalsession = session
	///計算ＮＧＦＧ標準差／／／
	// Total.TotalStd = math.Sqrt(Total.TotalVariance)
	// for i := 0; i < len(Total.NGMultiplestatisics); i++ {
	// 	Total.NGMultiplefrequency = append(Total.NGMultiplefrequency, float32(Total.NGMultiplestatisics[i])/float32(session))
	// 	Total.FGMultiplefrequency = append(Total.FGMultiplefrequency, float32(Total.FGMultiplestatisics[i])/float32(Total.RetriggerTimes))
	// }
	// for i := 0; i < symbolamount; i++ {
	// 	for k := 0; k < comboresultnum; k++ {
	// 		Total.NGFGcombo.NGresultfrequency[i][k] = float32(Total.NGFGcombo.NGresult[i][k])
	// 		Total.NGFGcombo.NGresultfrequency[i][k] = Total.NGFGcombo.NGresultfrequency[i][k] / float32(session)
	// 	}
	// }

	Total.Costtime = time.Since(s)

	Total.Risk = RiskControl

	return Total

}

///計算各輪結果///
func eachroundresult(risk Risk) EachRoundResult {
	var Eachresult = EachRoundResult{}

	///計算ＮＧcombo///
	Eachresult.NGresult = table.NGTablefunc().NGComboStatistic
	Eachresult.FGresult = table.FGTablefunc().FGComboStatistic

	// ///風控計算目前派彩，累積投注額，ＲＴＰ///

	ngresult := Ngjudgepassf(risk)
	Eachresult.NGpay = ngresult.NGpay
	Eachresult.Scatterpay = ngresult.Scatterpay
	Eachresult.Risk.Input = rngtool.OutputJudge(ngresult.Input)

	usertp := Eachresult.Risk.Choosertp
	a := &Rtp95
	b := &Rtp965
	c := &Rtp99
	switch usertp {
	case 95:
		*a += 1
	case 965:
		*b += 1
	case 99:
		*c += 1
	}

	///判斷進FG,並計分///
	if ngresult.Trigger == true {

		a := fgresult(Eachresult.Risk)

		Eachresult.EaceRetrigger++

		Eachresult.FGresult = a.FGEachresult
		Eachresult.FGpay += a.FGEachpay
		//Eachresult.FGMultiple = scoretool.Multiplejudge(Eachresult.FGpay)
		//Eachresult.FGVariance += a.FGVariance

		Eachresult.Risk.Input = a.Risk.Input
	}

	//Eachresult.NGVariance += math.Pow((float64(Eachresult.NGpay)/float64(costlevel) - float64(thmRTP)), 2)
	//Eachresult.NGMultiple = scoretool.Multiplejudge(Eachresult.NGpay)

	return Eachresult

}

func Ngjudgepassf(risk Risk) Ngjudgepass {

	var result = Ngjudgepass{}

	reproducttime := 1
	for i := 0; i < reproducttime; i++ {

		var normalpay int
		var scatterpay int
		var trigger bool
		var payoff int
		ngresult := rngtool.BaseGameRng()

		///工程盤面///
		//ngresult = [3][5]int{{13, 13, 13, 2, 2}, {5, 7, 7, 17, 6}, {6, 4, 1, 3, 4}}

		symbocombototal := scoretool.CombojudgeWayGame(ngresult)
		///陣列combo累加///
		for k := 0; k < col; k++ {
			//Eachresult.NGresult[symbocombototal[k][0]][symbocombototal[k][1]] += symbocombototal[k][2]
			///計分///
			normalpay += paytable[symbocombototal[k][0]][symbocombototal[k][1]] * symbocombototal[k][2] * Betmultiple
		}

		///scatter pay///
		resultmap := rngtool.Resultmapf(ngresult)
		if resultmap > 2 {
			trigger = true
			scatterpay = scoretool.Scatterpayf(costlevel, resultmap)

			//Eachresult.ScatterVariance += math.Pow((float64(Eachresult.Scatterpay) - float64(thmRTP)), 2)
		}

		payoff = normalpay + scatterpay

		///風控計算目前派彩，累積投注額，ＲＴＰ///
		type record struct {
			Risk
		}
		var tmp = record{}
		tmp.Bet = risk.Input.Bet + float32(costlevel)

		tmp.Accumulatepayoff = risk.Input.Accumulatepayoff + float32(normalpay+scatterpay)
		tmp.Accumulatebet = risk.Input.Accumulatebet + float32(costlevel)

		tmp.Currentrtp = tmp.Input.Accumulatepayoff / tmp.Input.Bet

		if payoff < payofflimit {
			result.NGpay = normalpay
			result.Scatterpay = scatterpay
			result.Trigger = trigger
			result.Risk = tmp.Risk
			//fmt.Println("pay<=50000", result, "重抽次數", reproducttime, result.Risk)
			break
		} else {
			if tmp.Currentrtp <= levelupper || reproducttime > 10 {
				result.NGpay = normalpay
				result.Scatterpay = scatterpay
				result.Trigger = trigger
				result.Risk = tmp.Risk
				//fmt.Println("pay>50000", result, "重抽次數", reproducttime)

				break
			} else {
				reproducttime++
				//fmt.Println("重抽", tmp.Currentrtp, reproducttime, trigger)
				continue
			}
		}
	}

	return result
}

///FG count payoff
func fgresult(risk Risk) EachRoundFGResult {

	var result = EachRoundFGResult{}
	result.FGEachresult = table.FGTablefunc().FGComboStatistic
	result.Risk = risk

	totalsession := fgsession
	retrigger := 1

	for i := 0; i < totalsession; i++ {

		freeresult := Fgjudgepassf(result.Risk)

		if freeresult.Trigger == true && retrigger < 10 {

			totalsession += fgsession
			retrigger++

		}
		result.FGEachpay += freeresult.FGpay + freeresult.Scatterpay
		///風控計算目前派彩，累積投注額，ＲＴＰ///

		result.Risk.Input = rngtool.OutputJudge(freeresult.Input)

	}
	//result.FGVariance = math.Pow((float64(result.FGEachpay)/float64(costlevel) - float64(thmRTP)), 2)

	return result
}

func Fgjudgepassf(risk Risk) Fgjudgepass {
	var result = Fgjudgepass{}

	reproducttime := 1
	for i := 0; i < reproducttime; i++ {

		var normalpay int
		var scatterpay int
		var trigger bool
		var payoff int
		fgresult := rngtool.FreeGameRng()

		///工程盤面///
		//fgresult = [3][5]int{{13, 13, 13, 2, 2}, {5, 7, 7, 17, 6}, {6, 4, 1, 3, 4}}

		symbocombototal := scoretool.CombojudgeWayGame(fgresult)
		///陣列combo累加///
		for k := 0; k < col; k++ {
			//Eachresult.NGresult[symbocombototal[k][0]][symbocombototal[k][1]] += symbocombototal[k][2]
			///計分///
			normalpay += paytable[symbocombototal[k][0]][symbocombototal[k][1]] * symbocombototal[k][2] * Betmultiple
		}

		///scatter pay///
		resultmap := rngtool.Resultmapf(fgresult)
		if resultmap > 2 {
			trigger = true
			scatterpay = scoretool.Scatterpayf(costlevel, resultmap)

			//Eachresult.ScatterVariance += math.Pow((float64(Eachresult.Scatterpay) - float64(thmRTP)), 2)
		}

		payoff = normalpay + scatterpay

		///風控計算目前派彩，累積投注額，ＲＴＰ///
		type record struct {
			Risk
		}
		var tmp = record{}
		tmp.Bet = risk.Input.Bet

		tmp.Accumulatepayoff = risk.Input.Accumulatepayoff + float32(normalpay+scatterpay)
		tmp.Accumulatebet = risk.Input.Accumulatebet

		tmp.Currentrtp = tmp.Input.Accumulatepayoff / tmp.Input.Bet

		if payoff < payofflimit {
			result.FGpay = normalpay
			result.Scatterpay = scatterpay
			result.Trigger = trigger
			result.Risk = tmp.Risk

			//fmt.Println("pay<=50000", result, "重抽次數", reproducttime, result.Risk)
			break
		} else {
			if tmp.Currentrtp <= levelupper || reproducttime > 10 {
				result.FGpay = normalpay
				result.Scatterpay = scatterpay
				result.Trigger = trigger
				result.Risk = tmp.Risk
				//fmt.Println("pay>50000", result, "重抽次數", reproducttime)

				break
			} else {
				reproducttime++
				//fmt.Println("重抽", tmp.Currentrtp, reproducttime, trigger)
				continue
			}
		}
	}

	return result
}
