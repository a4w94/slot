package calculate

import (
	"fmt"
	info "info/girlwa"
	"math"
	"math/rand"
	rngtool "rngtools/girlwa"
	scoretool "scoretool/girlwa"
	table "table/girlwa"
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
	Totalcombo          [][]int
	NGFGcombo           EachRoundResult
	RetriggerTimes      int
	TotalVariance       float64
	TotalStd            float64
	Costtime            interface{}
}

type EachRoundResult struct {
	NGresult          [][]int
	NGresultfrequency [linenum][comboresultnum]float32
	NGpay             int
	FGresult          [][]int
	FGpay             int
	Scatterpay        int
	EaceRetrigger     int
	NGVariance        float64
	FGVariance        float64
	ScatterVariance   float64
	NGMultiple        int
	FGMultiple        int
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

///線數///
const linenum = info.Linenum

///派彩表///
var paytable [][]int = table.PublicTablefunc().PayTable

///理論值ＲＴＰ///
var thmRTP = info.ThmRTP

///倍數統計///
var multiplestatisics = table.PublicTablefunc().MultipleTalbe

var FGenter int
var FGbouns int
var BG1 int
var BG2 int
var BG3 int
var BG4 int
var BG5 int
var BG6 int
var bgsit1 int
var bgsit1pay int
var bgsit2 int
var bgsit2pay int
var bgsit8 int
var bgsit8pay int

var FGbonusenter int

func Run(session int, runtime int, rtp int) {

	for i := 0; i < runtime; i++ {
		a := Totalresult(session, rtp)
		fmt.Println("unexpanding")
		fmt.Println("遊戲名稱 :", gamename)
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
		fmt.Println("FGRTP :", a.FGRTP)
		fmt.Println("倍數區間 ：  ", "0,~10,~20,~30,~40,~50,50~")
		fmt.Println("NG倍數統計 ：", a.NGMultiplestatisics)
		fmt.Println("NG倍數頻率 ：", a.NGMultiplefrequency)
		fmt.Println("FG倍數統計 ：", a.FGMultiplestatisics[1:])
		fmt.Println("FG倍數頻率 ：", a.FGMultiplefrequency[1:])
		fmt.Println("變異數 :", a.TotalVariance)
		fmt.Println("標準差 ：", a.TotalStd)
		fmt.Println("花費時間 ：", a.Costtime)

		fmt.Println("BONUS", float32(FGbouns)/float32(FGbonusenter))
		fmt.Println("BONUSfirst", float32(BG1)/float32(FGbonusenter))
		fmt.Println("BONUSsec", float32(BG2)/float32(FGbonusenter))
		fmt.Println("BONUS3", float32(BG3)/float32(FGbonusenter))

		fmt.Println("BONUS4", float32(BG4)/float32(FGbonusenter))

		fmt.Println("BONUS5", float32(BG5)/float32(FGbonusenter))
		fmt.Println("BONUS6", float32(BG6)/float32(FGbonusenter))

		fmt.Println("BONUSfreq", float32(FGbonusenter)/float32(FGenter*fgsession))

		fmt.Println("FGfreq", float32(FGenter)/float32(session))
		fmt.Println("BG第一種情況", float32(bgsit1pay)/float32(bgsit1))
		fmt.Println("BG第一種情況", float32(bgsit2pay)/float32(bgsit2))

		fmt.Println("BG第一種情況", float32(bgsit8pay)/float32(bgsit8))

		fmt.Println()
		fmt.Println("==============================================================================================================================================================================================================================")

	}

}

func Totalresult(session int, rtp int) TotalResult {
	s := time.Now()
	rand.Seed(int64(time.Now().UnixNano()))
	var Total = TotalResult{}
	Total.NGFGcombo.NGresult = table.NGTablefunc().NGComboStatistic
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

		for i := 0; i < symbolamount; i++ {
			for k := 0; k < comboresultnum; k++ {
				Total.NGFGcombo.NGresult[i][k] += a.NGresult[i][k]
			}
		}

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
	for i := 0; i < symbolamount; i++ {
		for k := 0; k < comboresultnum; k++ {
			Total.NGFGcombo.NGresultfrequency[i][k] = float32(Total.NGFGcombo.NGresult[i][k])
			Total.NGFGcombo.NGresultfrequency[i][k] = Total.NGFGcombo.NGresultfrequency[i][k] / float32(session)
		}
	}

	Total.Costtime = time.Since(s)

	return Total

}

///計算各輪結果///
func eachroundresult(rtp int, stritable [reelamount][]int) EachRoundResult {
	var Eachresult = EachRoundResult{}

	///計算ＮＧcombo///
	Eachresult.NGresult = table.NGTablefunc().NGComboStatistic

	result := rngtool.BaseGameRng(rtp)
	if rtp != 0 {
		result = rngtool.BaseGameWeightRng(stritable)
	}

	///工程盤面///
	//result = [4][5]int{{2, 2, 2, 2, 2}, {0, 2, 4, 2, 14}, {3, 0, 4, 2, 8}, {0, 2, 2, 12, 11}}

	resultmap := rngtool.Resultmapf(result)
	///判斷進FG,並計分///
	if resultmap[0] > 2 {
		FGenter++
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

	for i := 0; i < linenum; i++ {

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

	//var wildposition []int

	totalsession := fgsession

	for i := 0; i < totalsession; i++ {
		freeresult := rngtool.FreeGameRng()
		// freeresult = [4][5]int{
		// 	{5, 5, 5, 4, 4}, {11, 2, 4, 2, 14}, {3, 0, 4, 2, 8}, {7, 2, 2, 12, 11},
		// }

		resultmap := rngtool.Eachreelmapf(0, freeresult)

		if resultmap[1] > 2 {
			for k := 0; k < col; k++ {
				///R1 expanding wild///
				freeresult[k][0] = 1
			}
			//result.FGEachpay += bonusgamef(freeresult).BGEachpay
			FGbouns += bonusgamef(freeresult).BGEachpay
			FGbonusenter++
		}

		resultinti := rngtool.Linetablef(freeresult)

		for i := 0; i < 1; i++ {

			realsymbo, realcombo := scoretool.Combojudge(resultinti[i])
			// if resultinti[i][0] == 1 && resultinti[i][1] != 1 && resultinti[i][2] != 1 && resultinti[i][3] != 1 {
			// 	bgsit1++
			// 	bgsit1pay += paytable[realsymbo][realcombo]

			// }
			// if resultinti[i][0] == 1 && resultinti[i][1] == 1 && resultinti[i][2] != 1 && resultinti[i][3] != 1 {
			// 	bgsit2++
			// 	bgsit2pay += paytable[realsymbo][realcombo]

			// }
			// if resultinti[i][0] == 1 && resultinti[i][1] == 1 && resultinti[i][2] == 1 && resultinti[i][3] == 1 {
			// 	bgsit8++
			// 	bgsit8pay += paytable[realsymbo][realcombo]

			// }

			result.FGEachresult[realsymbo][realcombo]++
			result.FGEachpay += paytable[realsymbo][realcombo]

		}

	}
	result.FGVariance = math.Pow((float64(result.FGEachpay)/float64(linenum) - float64(thmRTP)), 2)

	return result
}

func bonusgamef(freeresult [col][reelamount]int) EachRoundBonusResult {
	var result = EachRoundBonusResult{}
	result.BGEachresult = table.BGTablefunc().BGComboStatistic

	/// 初始盤面wild數量///
	var lockwildamount = rngtool.Resultmapf(freeresult)[1]
	//fmt.Println("初始wild 數量", lockwildamount)

	var wildposition [col][]int

	var intilockwild [col][]int

	///回傳wild 位置///
	//fmt.Println("傳入freeresult", freeresult)
	intilockwild = lockwildf(freeresult)
	//fmt.Println("回傳intilockwild", intilockwild)
	///產出lockwild 位置／／／
	for t := 0; t < col; t++ {
		wildposition[t] = append(wildposition[t], intilockwild[t]...)
	}
	//fmt.Println("初始wild位置", wildposition)

	totalsession := 1
	for i := 0; i < totalsession; i++ {
		/// 產盤面///
		bonusresult := rngtool.BounsGameRng()
		//fmt.Println("bg每次rng盤面", bonusresult)
		for k := 0; k < col; k++ {
			///R1 expanding wild///
			bonusresult[k][0] = 1
		}

		var lockwild [col][]int

		///回傳wild 位置///
		lockwild = lockwildf(bonusresult)

		///產出lockwild 位置／／／
		for t := 0; t < col; t++ {
			wildposition[t] = append(wildposition[t], lockwild[t]...)
		}
		///依照lock位置改成wild///
		for k := 0; k < col; k++ {
			for j := 0; j < len(wildposition[k]); j++ {
				bonusresult[k][wildposition[k][j]] = 1
			}
		}
		//fmt.Println("lock後盤面", bonusresult)
		///比照數量 若wild數量比前一次多則增加一次respin///
		tmp := rngtool.Resultmapf(bonusresult)[1]
		//fmt.Println("盤面wild數量", tmp)
		if tmp > lockwildamount {
			totalsession++
			lockwildamount = tmp

		}
		///線表計算///
		resultline := rngtool.Linetablef(bonusresult)

		for j := 0; j < 1; j++ {

			realsymbo, realcombo := scoretool.Combojudge(resultline[i])
			if resultline[i][0] == 1 && resultline[i][1] != 1 && resultline[i][2] != 1 && resultline[i][3] != 1 {
				bgsit1++
				bgsit1pay += paytable[realsymbo][realcombo]

			}
			if resultline[i][0] == 1 && resultline[i][1] == 1 && resultline[i][2] != 1 && resultline[i][3] != 1 {
				bgsit2++
				bgsit2pay += paytable[realsymbo][realcombo]

			}
			if resultline[i][0] == 1 && resultline[i][1] == 1 && resultline[i][2] == 1 && resultline[i][3] == 1 {
				bgsit8++
				bgsit8pay += paytable[realsymbo][realcombo]

			}
			result.BGEachresult[realsymbo][realcombo]++
			result.BGEachpay += paytable[realsymbo][realcombo]
		}
		if totalsession == 1 {

			BG1 += result.BGEachpay
		}
	}
	switch totalsession {
	// case 1:
	// 	BG1 += result.BGEachpay
	case 2:
		BG2 += result.BGEachpay
	case 3:
		BG3 += result.BGEachpay
	case 4:
		BG4 += result.BGEachpay
	case 5:
		BG5 += result.BGEachpay
	case 6:
		BG6 += result.BGEachpay

	}

	return result
}

/// 丟進freegame rng return lockwild position
func lockwildf(freeresult [col][reelamount]int) [col][]int {

	var wildposition [col][]int
	for k := 0; k < col; k++ {

		for i := 1; i < reelamount; i++ {
			if freeresult[k][i] == 1 {
				wildposition[k] = append(wildposition[k], i)
			}
		}
	}
	return wildposition

}
