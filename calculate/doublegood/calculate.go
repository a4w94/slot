package calculate

import (
	"fmt"
	info "info/doublegood"
	"math/rand"
	rngtool "rngtools/doublegood"
	scoretool "scoretool/doublegood"
	table "table/doublegood"
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
	FreeGameStatus  bool

	Risk
}
type Ngjudgepass struct {
	NGpay      int
	Scatterpay int
	Trigger    bool
	Fgsession  int
	Risk
}

type Fgjudgepass struct {
	FGpay      int
	Scatterpay int
	Trigger    bool
	Fgsession  int
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

	///理論值ＲＴＰ///
	thmRTP = info.ThmRTP

	///Reel 轉輪數///
	reelamount int = info.Reelamount

	///Column 橫排數///
	col int = info.Col

	linenum int = info.Linenum

	///基本投注額Bet///
	bet int = info.Bet

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

	///scatter 贈送場數／／／
	scatterinfo = info.Scatterinfo

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
var web = TotalResult{}

func Run(session int, runtime int, costlevelinput int, rtp int) TotalResult {
	k := &costlevel
	*k = costlevelinput
	m := &Betmultiple
	*m = costlevelinput / bet

	//背景程式///

	for i := 0; i < runtime; i++ {
		// rngtool.Background()

		// fmt.Println(table.NGStriTablertp95)
		// fmt.Println(rngtool.Ngstritable)
		a := Totalresult(session)

		fmt.Println("遊戲名稱 :", gamename)
		fmt.Println("下注額 :", costlevel)
		fmt.Println("下注倍數 :", Betmultiple)

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

		///倍數區間
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
			if a.FGMultiplestatisics[time] == 0 {
				a.FGMultiplefrequency[time] = 0
			}
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

		///標準差
		fmt.Println("變異數 :", a.TotalVariance)
		fmt.Println("標準差 ：", a.TotalStd)

		fmt.Println("花費時間 ：", a.Costtime)
		fmt.Println("freegame機率", float32(a.RetriggerTimes)/float32(session))
		// var ngcombofrequency [symbolamount][comboresultnum]float32
		// var fgcombofrequency [symbolamount][comboresultnum]float32
		fmt.Println("95,965,99", Rtp95, Rtp965, Rtp99)
		fmt.Println(a.Risk)
		// for i := 0; i < symbolamount; i++ {
		// 	for k := 0; k < comboresultnum; k++ {
		// 		ngcombofrequency[i][k] = float32(a.NGFGcombo.NGresult[i][k]) / float32(session)
		// 		fgcombofrequency[i][k] = float32(a.NGFGcombo.FGresult[i][k]) / float32(a.RetriggerTimes)

		// 	}
		// }
		// // fmt.Println(a.NGFGcombo.NGresult)
		// // fmt.Println(ngcombofrequency)
		// fmt.Println(fgcombofrequency)
		fmt.Println(levelupper)
		fmt.Println()
		fmt.Println("==============================================================================================================================================================================================================================")
		web = a

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
	RiskControl.Accumulatebet = 0
	RiskControl.Accumulatepayoff = 241250
	RiskControl.Bet = 250000
	RiskControl.Currentrtp = 0.965
	for k := 0; k < session; k++ {

		//Total.Totalpay += eachroundresult().FGpay + eachroundresult().NGpay + eachroundresult().Scatterpay

		///風控流程///

		a := eachroundresult(RiskControl, k)

		///風控
		RiskControl = a.Risk
		//fmt.Println(RiskControl)

		///計算總派彩 ＮＧ總派採 ＦＧ總派彩 Ｓcatter總派彩 總變異數 總觸發ＦＧ次數 ＮＧＦＧ區間統計///

		Total.Totalpay += a.NGpay + a.FGpay + a.Scatterpay
		Total.NGtotalpay += a.NGpay + a.Scatterpay
		Total.FGtotalpay += a.FGpay
		Total.NGScattertotalpay += a.Scatterpay
		//Total.TotalVariance += float64(a.NGVariance+a.ScatterVariance+a.FGVariance) / float64(session)
		Total.RetriggerTimes += a.EaceRetrigger

		///倍數區間累加///
		Total.NGMultiplestatisics[a.NGMultiple]++
		if a.FreeGameStatus == true {
			Total.FGMultiplestatisics[a.FGMultiple]++
		}

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
	//Total.TotalStd = math.Sqrt(Total.TotalVariance)

	///計算倍數區間頻率
	for i := 0; i < len(Total.NGMultiplestatisics); i++ {
		Total.NGMultiplefrequency = append(Total.NGMultiplefrequency, float32(Total.NGMultiplestatisics[i])/float32(session))
		Total.FGMultiplefrequency = append(Total.FGMultiplefrequency, float32(Total.FGMultiplestatisics[i])/float32(Total.RetriggerTimes))
	}

	///
	// for i := 0; i < symbolamount; i++ {
	// 	for k := 0; k < comboresultnum; k++ {
	// 		Total.NGFGcombo.NGresultfrequency[i][k] = float32(Total.NGFGcombo.NGresult[i][k])
	// 		Total.NGFGcombo.NGresultfrequency[i][k] = Total.NGFGcombo.NGresultfrequency[i][k] / float32(session)
	// 	}
	// }

	Total.Costtime = time.Since(s)

	Total.Risk = RiskControl

	// if Total.Risk.Choosertp == 95 {
	// 	fmt.Println(Total.Risk)
	// 	fmt.Println("PAY", Total.NGtotalpay)
	// }

	return Total

}

///計算各輪結果///
func eachroundresult(risk Risk, sessiont int) EachRoundResult {
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

		Eachresult.FreeGameStatus = true
		a := fgresult(Eachresult.Risk, ngresult.Fgsession)

		Eachresult.EaceRetrigger++

		Eachresult.FGresult = a.FGEachresult
		Eachresult.FGpay += a.FGEachpay

		///計算ＦＧ得分倍數區間
		Eachresult.FGMultiple = scoretool.Multiplejudge(Eachresult.FGpay)
		//Eachresult.FGVariance += a.FGVariance

		Eachresult.Risk.Input = a.Risk.Input
	}
	if usertp == 99 {
		fmt.Println("場次", sessiont)
		fmt.Println("ngpay:", ngresult.NGpay)
		fmt.Println("BET", ngresult.Bet)
		fmt.Println("Payoff:", ngresult.Accumulatepayoff)
		fmt.Println("累積BET", ngresult.Accumulatebet)
		fmt.Println("轉輪表：", usertp)
		if Eachresult.FGpay > 0 {
			fmt.Println("FGpay:", Eachresult.FGpay, "====================")
		}

		fmt.Println()

	}

	//計算標準差//
	//Eachresult.NGVariance += math.Pow((float64(Eachresult.NGpay)/float64(costlevel) - float64(thmRTP)), 2)

	//計算ＮＧ得分倍數區間／／／
	Eachresult.NGMultiple = scoretool.Multiplejudge(Eachresult.NGpay)

	return Eachresult

}

type Scatter struct {
	scoretool.Scatterinfo
}

func Ngjudgepassf(risk Risk) Ngjudgepass {

	var result = Ngjudgepass{}

	reproducttime := 1
	for i := 0; i < reproducttime; i++ {
		result.Fgsession = 0

		var normalpay int

		var trigger bool
		var payoff int
		ngresult := rngtool.BaseGameRng()

		///工程盤面///
		//fmt.Println(ngresult)
		//ngresult = [3][5]int{{13, 13, 23, 21, 2}, {5, 7, 7, 17, 6}, {6, 4, 13, 23, 23}}

		resultline := rngtool.Linetablef(ngresult)

		for i := 0; i < linenum; i++ {

			realsymbo, realcombo := scoretool.CombojudgeLineGame(resultline[i])

			// ///陣列combo累加///
			// Eachresult.NGresult[realsymbo][realcombo]++

			///計分///
			normalpay += paytable[realsymbo][realcombo]

		}

		///scatter pay///
		var scatterinfo = Scatter{}.Scatterinfo
		resultmap := rngtool.Resultmapf(ngresult)
		if resultmap > 2 {
			trigger = true
			scatterinfo = scoretool.Scatterpayf(costlevel, resultmap)
			result.Fgsession = scatterinfo.Fgsession

			//result.ScatterVariance += math.Pow((float64(Eachresult.Scatterpay) - float64(thmRTP)), 2)
		}

		payoff = normalpay + scatterinfo.Scatterpay

		///風控計算目前派彩，累積投注額，ＲＴＰ///
		type record struct {
			Risk
		}
		var tmp = record{}
		tmp.Bet = risk.Input.Bet + float32(costlevel)

		tmp.Accumulatepayoff = risk.Input.Accumulatepayoff + float32(normalpay+scatterinfo.Scatterpay)
		tmp.Accumulatebet = risk.Input.Accumulatebet + float32(costlevel)

		tmp.Currentrtp = tmp.Input.Accumulatepayoff / tmp.Input.Bet

		if payoff < payofflimit {
			result.NGpay = normalpay
			result.Scatterpay = scatterinfo.Scatterpay
			result.Trigger = trigger
			result.Risk = tmp.Risk
			//fmt.Println("pay<=50000", result, "重抽次數", reproducttime, result.Risk)
			break
		} else {
			if tmp.Currentrtp <= levelupper || reproducttime > 10 {
				result.NGpay = normalpay
				result.Scatterpay = scatterinfo.Scatterpay
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
func fgresult(risk Risk, fgsession int) EachRoundFGResult {

	var result = EachRoundFGResult{}
	result.FGEachresult = table.FGTablefunc().FGComboStatistic
	result.Risk = risk

	totalsession := fgsession
	//retrigger := 1

	for i := 0; i < totalsession; i++ {

		freeresult := Fgjudgepassf(result.Risk)

		if freeresult.Trigger == true {

			totalsession += freeresult.Fgsession
			// retrigger++

		}

		result.FGEachpay += freeresult.FGpay + freeresult.Scatterpay
		///風控計算目前派彩，累積投注額，ＲＴＰ///

		result.Risk.Input = rngtool.OutputJudge(freeresult.Input)

	}
	//result.FGVariance = math.Pow((float64(result.FGEachpay)/float64(costlevel) - float64(thmRTP)), 2)

	// fmt.Println("FGpay", freeresult.FGpay)
	// fmt.Println("BET", result.Risk.Bet)
	// fmt.Println("Payoff", result.Risk.Accumulatepayoff)
	// fmt.Println()

	return result
}

func Fgjudgepassf(risk Risk) Fgjudgepass {
	var result = Fgjudgepass{}

	reproducttime := 1
	for i := 0; i < reproducttime; i++ {
		result.Fgsession = 0

		var normalpay int

		var trigger bool
		var payoff int
		fgresult := rngtool.FreeGameRng()

		///工程盤面///
		//fgresult = [3][5]int{{13, 13, 13, 2, 2}, {5, 7, 7, 17, 6}, {6, 4, 1, 3, 4}}

		resultline := rngtool.Linetablef(fgresult)

		for i := 0; i < linenum; i++ {

			realsymbo, realcombo := scoretool.CombojudgeLineGame(resultline[i])
			// ///陣列combo累加///
			// Eachresult.NGresult[realsymbo][realcombo]++
			///計分///
			normalpay += paytable[realsymbo][realcombo]

			// Eachresult.NGVariance += math.Pow((float64(Eachresult.NGpay)/float64(linenum) - float64(thmRTP)), 2)
			// Eachresult.NGMultiple = scoretool.Multiplejudge(Eachresult.NGpay)
		}

		///scatter pay///
		var scatterinfo = Scatter{}.Scatterinfo
		resultmap := rngtool.Resultmapf(fgresult)
		if resultmap > 2 {
			trigger = true
			scatterinfo = scoretool.Scatterpayf(costlevel, resultmap)
			result.Fgsession = scatterinfo.Fgsession

			//Eachresult.ScatterVariance += math.Pow((float64(Eachresult.Scatterpay) - float64(thmRTP)), 2)
		}

		payoff = normalpay + scatterinfo.Fgsession

		///風控計算目前派彩，累積投注額，ＲＴＰ///
		type record struct {
			Risk
		}
		var tmp = record{}
		tmp.Bet = risk.Input.Bet

		tmp.Accumulatepayoff = risk.Input.Accumulatepayoff + float32(normalpay+scatterinfo.Fgsession)
		tmp.Accumulatebet = risk.Input.Accumulatebet

		tmp.Currentrtp = tmp.Input.Accumulatepayoff / tmp.Input.Bet

		if payoff < payofflimit {
			result.FGpay = normalpay
			result.Scatterpay = scatterinfo.Fgsession
			result.Trigger = trigger
			result.Risk = tmp.Risk

			//fmt.Println("pay<=50000", result, "重抽次數", reproducttime, result.Risk)
			break
		} else {
			if tmp.Currentrtp <= levelupper || reproducttime > 10 {
				result.FGpay = normalpay
				result.Scatterpay = scatterinfo.Fgsession
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
