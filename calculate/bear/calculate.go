package calculate

import (
	"fmt"
	info "info/bear"
	"math"
	"math/rand"
	rngtool "rngtools/bear"
	scoretool "scoretool/bear"
	table "table/bear"
	"time"
)

type TotalResult struct {
	RTP                 float32
	Totalpay            int
	NGtotalpay          int
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
	Totalcombo          [][]int
	NGFGcombo           EachRoundResult
	RetriggerTimes      int
	TotalVariance       float64
	TotalStd            float64
	Costtime            interface{}
}

type EachRoundResult struct {
	NGresult        [][]int
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
}

type EachRoundFGResult struct {
	FGEachresult [][]int
	FGEachpay    int
	FGVariance   float64
}

const gamename string = info.GameName

///Reel 轉輪數///
const reelamount int = info.Reelamount

///Column 橫排數///
const col int = info.Col

///基本投注額Bet///
const bet int = info.Bet

///獎圖總數///
const symbolamount = info.Symbolamount

///combo數目0~5combo///
const comboresultnum = info.Comboresultnum

///線數///
const linenum = info.Linenum

///派彩表///
var paytable [][]int = table.PublicTablefunc().PayTable

///理論值ＲＴＰ///
var thmRTP = info.ThmRTP

///倍數統計///
var multiplestatisics = table.PublicTablefunc().MultipleTalbe

func Run(session int, runtime int, rtp int) {

	for i := 0; i < runtime; i++ {
		a := Totalresult(session, rtp)
		if rtp != 0 {
			fmt.Println("THM RTP :", rtp)
		}

		fmt.Println("遊戲名稱 :", gamename)
		fmt.Println("試驗次數 ：", session)
		fmt.Println("總派彩 :", a.Totalpay)

		fmt.Println("ＲＴＰ :", a.RTP)

		fmt.Println()
		fmt.Println("XXXXXXXXXXXXXXXXX細項XXXXXXXXXXXXXXXXXXXXX")
		fmt.Println()

		fmt.Println("NG派彩 :", a.NGtotalpay)
		fmt.Println("NGRTP :", a.NGRTP)
		fmt.Println("NGsctterpay :", a.NGScattertotalpay)
		fmt.Println("NGscatter RTP :", a.NGScatterRTP)
		fmt.Println("FG派彩 :", a.FGtotalpay)
		fmt.Println("FGRTP :", a.FGRTP)
		fmt.Println("倍數區間 ：  ", "0,~10,~20,~30,~40,~50,50~")
		fmt.Println("NG倍數統計 ：", a.NGMultiplestatisics)
		fmt.Println("NG倍數頻率 ：", a.NGMultiplefrequency)
		fmt.Println("FG倍數統計 ：", a.FGMultiplestatisics[1:])
		fmt.Println("FG倍數頻率 ：", a.FGMultiplefrequency[1:])
		fmt.Println("變異數 :", a.TotalVariance)
		fmt.Println("標準差 ：", a.TotalStd)
		fmt.Println("花費時間 ：", a.Costtime)

		fmt.Println()
		fmt.Println("==================================================================================")

	}

}

func Totalresult(session int, rtp int) TotalResult {
	s := time.Now()
	rand.Seed(int64(time.Now().UnixNano()))
	var Total = TotalResult{}
	// Total.NGFGcombo.NGresult = table.NGTablefunc().NGComboStatistic
	// Total.NGFGcombo.FGresult = table.FGTablefunc().FGComboStatistic
	//Total.Totalcombo = table.PublicTablefunc().TotalTable
	var stritableweight [reelamount][]int
	switch rtp {
	case 95:
		stritableweight = rngtool.Stritableweightf(table.NGTablefunc().NGWeightTable95)
	case 98:
		stritableweight = rngtool.Stritableweightf(table.NGTablefunc().NGWeightTable98)
	}

	for k := 0; k < session; k++ {
		//Total.Totalpay += eachroundresult().FGpay + eachroundresult().NGpay + eachroundresult().Scatterpay

		a := eachroundresult(rtp, stritableweight)
		///計算總派彩 ＮＧ總派採 ＦＧ總派彩 Ｓcatter總派彩 總變異數 總觸發ＦＧ次數 ＮＧＦＧ區間統計///
		Total.Totalpay += a.NGpay + a.FGpay + a.Scatterpay
		Total.NGtotalpay += a.NGpay + a.Scatterpay
		Total.FGtotalpay += a.FGpay
		Total.NGScattertotalpay += a.Scatterpay
		Total.TotalVariance += float64(a.NGVariance+a.ScatterVariance+a.FGVariance) / float64(session)
		Total.RetriggerTimes += a.EaceRetrigger
		Total.NGMultiplestatisics[a.NGMultiple]++
		Total.FGMultiplestatisics[a.FGMultiple]++

	}

	///計算總ＲＴＰ，ＮＧＲＴＰ，ＦＧＲＴＰ///
	Total.RTP = float32(Total.Totalpay) / (float32(session) * float32(bet))
	Total.NGRTP = float32(Total.NGtotalpay) / (float32(session) * float32(bet))
	Total.FGRTP = float32(Total.FGtotalpay) / (float32(session) * float32(bet))
	Total.NGScatterRTP = float32(Total.NGScattertotalpay) / (float32(session) * float32(bet))

	Total.Totalsession = session
	///計算ＮＧＦＧ標準差／／／
	Total.TotalStd = math.Sqrt(Total.TotalVariance)
	for i := 0; i < len(Total.NGMultiplestatisics); i++ {
		Total.NGMultiplefrequency = append(Total.NGMultiplefrequency, float32(Total.NGMultiplestatisics[i])/float32(session))
		Total.FGMultiplefrequency = append(Total.FGMultiplefrequency, float32(Total.FGMultiplestatisics[i])/float32(Total.RetriggerTimes))
	}

	Total.Costtime = time.Since(s)

	return Total

}

///計算各輪結果///
func eachroundresult(rtp int, stritable [reelamount][]int) EachRoundResult {
	var Eachresult = EachRoundResult{}

	///計算ＮＧcombo///
	Eachresult.NGresult = table.NGTablefunc().NGComboStatistic

	result := rngtool.BaseGameRng()
	if rtp != 0 {
		result = rngtool.BaseGameWeightRng(stritable)
	}

	///	工程盤面///
	//result = [4][5]int{{2, 2, 2, 2, 2}, {0, 2, 4, 2, 14}, {3, 0, 4, 2, 8}, {0, 2, 2, 12, 11}}

	resultmap := rngtool.Resultmapf(result)
	///判斷進FG,並計分///
	if resultmap[0] > 2 {
		a := fgresult()

		Eachresult.EaceRetrigger++

		Eachresult.FGresult = a.FGEachresult
		Eachresult.FGpay += a.FGEachpay
		Eachresult.FGMultiple = scoretool.Multiplejudge(Eachresult.FGpay)

		Eachresult.FGVariance += a.FGVariance
		scatterpay := scoretool.Scatterpayf(resultmap)
		Eachresult.Scatterpay += scatterpay
		Eachresult.ScatterVariance += math.Pow((float64(Eachresult.Scatterpay) - float64(thmRTP)), 2)
	}

	resultline := rngtool.Linetablef(result)

	for i := 0; i < len(resultline); i++ {

		realsymbo, realcombo := scoretool.Combojudge(resultline[i])
		///陣列combo累加///
		Eachresult.NGresult[realsymbo][realcombo]++
		///計分///
		Eachresult.NGpay += paytable[realsymbo][realcombo]
		Eachresult.NGVariance += math.Pow((float64(Eachresult.NGpay)/float64(linenum) - float64(thmRTP)), 2)
		Eachresult.NGMultiple = scoretool.Multiplejudge(Eachresult.NGpay)
	}

	return Eachresult

}

///FG count payoff
func fgresult() EachRoundFGResult {
	var result = EachRoundFGResult{}
	result.FGEachresult = table.FGTablefunc().FGComboStatistic

	var wildposition []int

	totalsession := 10

	for i := 0; i < totalsession; i++ {
		freeresult := rngtool.FreeGameRng()
		// freeresult = [4][5]int{
		// 	{5, 5, 5, 4, 4}, {11, 2, 4, 2, 14}, {3, 0, 4, 2, 8}, {7, 2, 2, 12, 11},
		// }

		resultmap := rngtool.Resultmapf(freeresult)

		if resultmap[0] > 2 {
			result.FGEachpay += scoretool.Scatterpayf(resultmap)

			if totalsession < 20 {

				totalsession += 10

			}
		}

		resultinti := rngtool.Linetablef(freeresult)
		var lockwild []int
		lockwild = lockwildf(freeresult)

		wildposition = append(wildposition, lockwild...) ///append wild position///

		for i := 0; i < len(resultinti); i++ {
			//for i := 0; i < 2; i++ {
			resultline := resultinti[i]

			for j := 0; j < len(wildposition); j++ {
				resultline[wildposition[j]] = 1
			}

			realsymbo, realcombo := scoretool.Combojudge(resultline)

			result.FGEachresult[realsymbo][realcombo]++
			result.FGEachpay += paytable[realsymbo][realcombo]

		}

	}
	result.FGVariance = math.Pow((float64(result.FGEachpay)/float64(linenum) - float64(thmRTP)), 2)

	return result
}

/// 丟進freegame rng return lockwild position
func lockwildf(freeresult [col][reelamount]int) []int {

	var wildposition []int

	for i := 0; i < reelamount; i++ {
		if freeresult[0][i] == 1 && freeresult[1][i] == 1 && freeresult[2][i] == 1 && freeresult[3][i] == 1 {
			wildposition = append(wildposition, i)
		}
	}
	return wildposition

}
