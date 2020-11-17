package info

import "fmt"

const (
	///遊戲名稱///
	GameName string = "好事成雙"

	///Reel 轉輪數///
	Reelamount int = 5

	///Column 橫排數///
	Col int = 3

	///基本投注額Bet///
	Bet int = 30

	///線數///
	Linenum int = 30

	///獎圖總數+1 沒有零號///
	Symbolamount int = 24

	///combo數目0~5combo///
	Comboresultnum int = 11

	///WILD 代號///
	Wild       int = 21
	Wilddouble int = 22

	///Scatter 代號///
	Scatter int = 23

	///理論值ＲＴＰ///
	ThmRTP float32 = 0.965

	///ＲＴＰ95///
	RTP95 float32 = 0.95

	///ＲＴＰ98///
	RTP98 float32 = 0.98

	///     風控     ／／／

	Payofflimit = 50000
	///水位上下限．．///
	LevelUpper = 0.967
	LevelLower = 0.963
	///初始化///
	InitialAccumulatebet    = 0
	InitialAccumulatepayoff = 2412500
	InitialBet              = 250000

	FirstAccumulatebetCheck  = 800000
	SecondAccumulatebetCheck = 1200000
)

var Excelroutieng = "/Users/terry_hsiesh/go/game/doublegood/parsheet/ngparsheet.xlsx"
var Excelroutiefg = "/Users/terry_hsiesh/go/game/doublegood/parsheet/fgparsheet.xlsx"

///倍數表///
var Multiple = []int{0, 0, 1, 5, 10, 30, 50, 100}

///free game 場數///
var Scatterinfo = [3][]int{
	///scatter數目
	{3, 4, 5},

	///賠率
	{2, 10, 200},

	///贈送場次///
	{10, 15, 20},
}

func Infomation() {
	fmt.Println("遊戲名稱 : ", GameName)
	fmt.Println("轉輪數 : ", Reelamount)
	fmt.Println("橫排數 : ", Col)
	fmt.Println("基本投注額 : ", Bet)
	fmt.Println("獎圖總數 : ", Symbolamount)
	fmt.Println("combo數目0~5combo : ", Comboresultnum)

	fmt.Println("理論值ＲＴＰ ：", ThmRTP)

}
