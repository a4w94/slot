package calculate

import (
	"fmt"
	info "info/doublegood"
	"math"
	"math/rand"
	per "permutation/doublegood"
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
}

type EachRoundFGResult struct {
	FGEachresult [symbolamount][comboresultnum]int
	FGEachpay    int

	FGVariance float64
}

type EachRoundBonusResult struct {
	BGEachresult [symbolamount][comboresultnum]int
	BGEachpay    int
}

const gamename string = info.GameName

///Reel 轉輪數///
const reelamount int = info.Reelamount

///Column 橫排數///
const col int = info.Col

///基本投注額Bet///
var costlevel int

///獎圖總數///
const symbolamount = info.Symbolamount

///combo數目0~5combo///
const comboresultnum = info.Comboresultnum

///線數///
const linenum int = info.Linenum

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

///free game 贈送場數／／／
var scatterinfo = info.Scatterinfo

var ngstritable = [][]int{}
var fgstritable = [][]int{}

var ngcombototal = rngtool.Eachreellen("NG", 965)
var fgcombototal = rngtool.Eachreellen("FG", 965)

func Run(session int, runtime int, costlevelinput int, rtp int) {
	fmt.Println("ng總combo", ngcombototal)
	fmt.Println("fg總combo", fgcombototal)
	k := &costlevel
	*k = costlevelinput

	for i := 0; i < runtime; i++ {
		a := Totalresult(session)

		fmt.Println("遊戲名稱 :", gamename)
		fmt.Println("等級 :", costlevel)
		fmt.Println("理論RTP :", rtp)
		fmt.Println("試驗次數 ：", session)
		fmt.Println("總派彩 :", a.Totalpay)

		//fmt.Println("ＲＴＰ :", a.RTP)

		fmt.Println()
		fmt.Println("XXXXXXXXXXXXXXXXX細項XXXXXXXXXXXXXXXXXXXXX")
		fmt.Println()

		fmt.Println("NG派彩 :", a.NGtotalpay)
		fmt.Println("NGRTP :", a.NGRTP)
		//fmt.Println("NGcombo頻率", a.NGFGcombo.NGresultfrequency)
		fmt.Println("NGsctterpay :", a.NGScattertotalpay)
		fmt.Println("NGscatter RTP :", a.NGScatterRTP)
		fmt.Println("freegame次數", a.RetriggerTimes)
		fmt.Println("freegame機率", float32(a.RetriggerTimes)/float32(ngcombototal))
		fmt.Println()
		fmt.Println("FG派彩 :", a.FGtotalpay)
		//fmt.Println("FG獨立ＲＴＰ :", float32(a.FGtotalpay)/float32(a.RetriggerTimes))
		fmt.Println("FGRTP :", a.FGRTP)
		//fmt.Println("倍數區間 ：  ", "0,~10,~20,~30,~40,~50,50~")
		//fmt.Println("NG倍數統計 ：", a.NGMultiplestatisics)
		//fmt.Println("NG倍數頻率 ：", a.NGMultiplefrequency)
		//fmt.Println("FG倍數統計 ：", a.FGMultiplestatisics)
		//fmt.Println("FG倍數頻率 ：", a.FGMultiplefrequency)
		//fmt.Println("變異數 :", a.TotalVariance)
		//fmt.Println("標準差 ：", a.TotalStd)
		fmt.Println("花費時間 ：", a.Costtime)

		// var ngcombofrequency [symbolamount][comboresultnum]float32
		// var fgcombofrequency [symbolamount][comboresultnum]float32

		// for i := 0; i < symbolamount; i++ {
		// 	for k := 0; k < comboresultnum; k++ {
		// 		ngcombofrequency[i][k] = float32(a.NGFGcombo.NGresult[i][k]) / float32(ngcombototal)
		// 		fgcombofrequency[i][k] = float32(a.NGFGcombo.FGresult[i][k]) / float32(a.RetriggerTimes)

		// 	}
		// }
		fmt.Println("NGpaycombo")
		for t := 0; t < symbolamount; t++ {

			fmt.Println(t, "", a.NGFGcombo.NGresult[t])
		}

		fmt.Println("FGpaycombo")
		for t := 0; t < symbolamount; t++ {

			fmt.Println(t, "", a.NGFGcombo.FGresult[t])
		}
		// // fmt.Println(ngcombofrequency)
		// fmt.Println(fgcombofrequency)

		fmt.Println()
		fmt.Println("==============================================================================================================================================================================================================================")

	}

}

func Totalresult(session int) TotalResult {
	s := time.Now()
	rand.Seed(int64(time.Now().UnixNano()))
	var Total = TotalResult{}
	Total.NGFGcombo.NGresult = table.NGTablefunc().NGComboStatistic
	Total.NGFGcombo.FGresult = table.FGTablefunc().FGComboStatistic
	//Total.Totalcombo = table.PublicTablefunc().TotalTable

	for k := 0; k < session; k++ {
		//Total.Totalpay += eachroundresult().FGpay + eachroundresult().NGpay + eachroundresult().Scatterpay

		a := eachroundresult()
		///計算總派彩 ＮＧ總派採 ＦＧ總派彩 Ｓcatter總派彩 總變異數 總觸發ＦＧ次數 ＮＧＦＧ區間統計///
		Total.Totalpay += a.NGpay + a.FGpay + a.Scatterpay
		Total.NGtotalpay += a.NGpay + a.Scatterpay
		Total.FGtotalpay += a.FGpay
		Total.NGScattertotalpay += a.Scatterpay
		//Total.TotalVariance += float64(a.NGVariance+a.ScatterVariance+a.FGVariance) / float64(ngcombototal)
		Total.RetriggerTimes += a.EaceRetrigger
		//Total.NGMultiplestatisics[a.NGMultiple]++
		//Total.FGMultiplestatisics[a.FGMultiple]++

		for i := 0; i < symbolamount; i++ {
			for k := 0; k < comboresultnum; k++ {
				Total.NGFGcombo.NGresult[i][k] += a.NGresult[i][k]
				Total.NGFGcombo.FGresult[i][k] += a.FGresult[i][k]
			}
		}

	}

	///計算總ＲＴＰ，ＮＧＲＴＰ，ＦＧＲＴＰ///
	Total.RTP = float32(Total.Totalpay) / (float32(ngcombototal) * float32(costlevel))
	Total.NGRTP = float32(Total.NGtotalpay) / (float32(ngcombototal) * float32(costlevel))
	Total.FGRTP = float32(Total.FGtotalpay) / (float32(fgcombototal) * float32(costlevel))
	Total.NGScatterRTP = float32(Total.NGScattertotalpay) / (float32(ngcombototal) * float32(costlevel))

	Total.Totalsession = session
	///計算ＮＧＦＧ標準差／／／
	// Total.TotalStd = math.Sqrt(Total.TotalVariance)
	// for i := 0; i < len(Total.NGMultiplestatisics); i++ {
	// 	Total.NGMultiplefrequency = append(Total.NGMultiplefrequency, float32(Total.NGMultiplestatisics[i])/float32(ngcombototal))
	// 	Total.FGMultiplefrequency = append(Total.FGMultiplefrequency, float32(Total.FGMultiplestatisics[i])/float32(Total.RetriggerTimes))
	// }
	// for i := 0; i < symbolamount; i++ {
	// 	for k := 0; k < comboresultnum; k++ {
	// 		Total.NGFGcombo.NGresultfrequency[i][k] = float32(Total.NGFGcombo.NGresult[i][k])
	// 		Total.NGFGcombo.NGresultfrequency[i][k] = Total.NGFGcombo.NGresultfrequency[i][k] / float32(ngcombototal)
	// 	}
	// }

	Total.Costtime = time.Since(s)

	return Total

}

///計算各輪結果///
func eachroundresult() EachRoundResult {
	var Eachresult = EachRoundResult{}

	///計算ＮＧcombo///
	Eachresult.NGresult = table.NGTablefunc().NGComboStatistic
	Eachresult.FGresult = table.FGTablefunc().FGComboStatistic

	var ngstri = rngtool.Ngstritable
	var fgstri = rngtool.Fgstritable
	fmt.Println(ngstri)
	fmt.Println(fgstri)

	for a := 0; a < len(ngstri[0]); a++ {
		for b := 0; b < len(ngstri[1]); b++ {
			for c := 0; c < len(ngstri[2]); c++ {
				for d := 0; d < len(ngstri[3]); d++ {
					for e := 0; e < len(ngstri[4]); e++ {
						row1 := [5]int{a, b, c, d, e}
						ngresult := per.BaseGameallcombo(row1)
						resultmap := rngtool.Resultmapf(ngresult)

						///判斷進FG,並計分///
						if resultmap > 2 {
							ngscatterinfo := scoretool.Scatterpayf(costlevel, resultmap)
							//a := fgresult(ngscatterinfo.Fgsession)

							Eachresult.EaceRetrigger++

							//Eachresult.FGresult = a.FGEachresult
							//Eachresult.FGpay += a.FGEachpay

							//Eachresult.FGMultiple = scoretool.Multiplejudge(Eachresult.FGpay)
							//Eachresult.FGVariance += a.FGVariance

							Eachresult.Scatterpay += ngscatterinfo.Scatterpay
							//Eachresult.ScatterVariance += math.Pow((float64(Eachresult.Scatterpay) - float64(thmRTP)), 2)
						}

						resultline := rngtool.Linetablef(ngresult)

						for i := 0; i < linenum; i++ {

							realsymbo, realcombo := scoretool.CombojudgeLineGame(resultline[i])
							///陣列combo累加///
							Eachresult.NGresult[realsymbo][realcombo]++
							///計分///
							Eachresult.NGpay += paytable[realsymbo][realcombo]
							Eachresult.NGVariance += math.Pow((float64(Eachresult.NGpay)/float64(linenum) - float64(thmRTP)), 2)
							Eachresult.NGMultiple = scoretool.Multiplejudge(Eachresult.NGpay)
						}

					}
				}
			}
		}
	}

	for a := 0; a < len(fgstri[0]); a++ {
		for b := 0; b < len(fgstri[1]); b++ {
			for c := 0; c < len(fgstri[2]); c++ {
				for d := 0; d < len(fgstri[3]); d++ {
					for e := 0; e < len(fgstri[4]); e++ {
						row1 := [5]int{a, b, c, d, e}
						fgresult := per.BaseGameallcombo(row1)
						resultmap := rngtool.Resultmapf(fgresult)

						///判斷進FG,並計分///
						if resultmap > 2 {
							fgscatterinfo := scoretool.Scatterpayf(costlevel, resultmap)
							//a := fgresult(ngscatterinfo.Fgsession)

							//Eachresult.EaceRetrigger++

							//Eachresult.FGresult = a.FGEachresult
							//Eachresult.FGpay += a.FGEachpay

							//Eachresult.FGMultiple = scoretool.Multiplejudge(Eachresult.FGpay)
							//Eachresult.FGVariance += a.FGVariance

							Eachresult.FGpay += fgscatterinfo.Scatterpay
							//Eachresult.ScatterVariance += math.Pow((float64(Eachresult.Scatterpay) - float64(thmRTP)), 2)
						}

						resultline := rngtool.Linetablef(fgresult)

						for i := 0; i < linenum; i++ {

							realsymbo, realcombo := scoretool.CombojudgeLineGame(resultline[i])
							///陣列combo累加///
							Eachresult.FGresult[realsymbo][realcombo]++
							///計分///
							Eachresult.FGpay += paytable[realsymbo][realcombo]
							Eachresult.FGVariance += math.Pow((float64(Eachresult.NGpay)/float64(linenum) - float64(thmRTP)), 2)
							Eachresult.FGMultiple = scoretool.Multiplejudge(Eachresult.NGpay)
						}

					}
				}
			}
		}
	}

	return Eachresult

}

///FG count payoff
func fgresult(fgsession int) EachRoundFGResult {
	var result = EachRoundFGResult{}
	result.FGEachresult = table.FGTablefunc().FGComboStatistic
	var fgstri = rngtool.Fgstritable

	//var wildposition []int

	totalsession := fgsession
	// retrigger := 1

	for i := 0; i < totalsession; i++ {

		for a := 0; a < len(fgstri[0]); a++ {
			for b := 0; b < len(fgstri[1]); b++ {
				for c := 0; c < len(fgstri[2]); c++ {
					for d := 0; d < len(fgstri[3]); d++ {
						for e := 0; e < len(fgstri[4]); e++ {
							row1 := [5]int{a, b, c, d, e}
							fgresult := per.BaseGameallcombo(row1)
							resultmap := rngtool.Resultmapf(fgresult)

							///判斷FGretigger,並計分///

							if resultmap > 2 {
								fgscatterinfo := scoretool.Scatterpayf(costlevel, resultmap)
								totalsession += fgscatterinfo.Fgsession
								result.FGEachpay += fgscatterinfo.Scatterpay

							}

							resultinti := rngtool.Linetablef(fgresult)

							for i := 0; i < linenum; i++ {

								realsymbo, realcombo := scoretool.CombojudgeLineGame(resultinti[i])

								result.FGEachresult[realsymbo][realcombo]++
								result.FGEachpay += paytable[realsymbo][realcombo]

							}

						}
					}
				}
			}
		}

	}

	//result.FGVariance = math.Pow((float64(result.FGEachpay)/float64(costlevel) - float64(thmRTP)), 2)

	return result
}
