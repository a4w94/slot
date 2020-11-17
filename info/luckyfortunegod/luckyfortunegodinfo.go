package info

import "fmt"

const (
	///遊戲名稱///
	GameName string = "多福財神"

	///遊戲模式///
	GamePayLines string = "243ways"

	///Reel 轉輪數///
	Reelamount int = 5

	///Column 橫排數///
	Col int = 3

	///基本投注額Bet///
	Bet int = 50

	///線數///
	//Linenum int = 1

	///獎圖總數+1 沒有零號///
	Symbolamount int = 14

	///combo數目0~5combo///
	Comboresultnum int = 6

	///WILD 代號///
	Wild int = 12

	///Scatter 代號///
	Scatter int = 13

	///錢幣///
	Money int = 1

	///理論值ＲＴＰ///
	ThmRTP float32 = 0.965

	///ＲＴＰ95///
	RTP95 float32 = 0.95

	///ＲＴＰ99///
	RTP99 float32 = 0.99

	///     風控     ／／／

	Payofflimit = 50000
	///水位上下限．．///
	LevelUpper = 0.967
	LevelLower = 0.963
	///初始化///
	InitialAccumulatebet    = 0
	InitialAccumulatepayoff = 2412500
	InitialBet              = 2500000

	FirstAccumulatebetCheck  = 800000
	SecondAccumulatebetCheck = 1200000
)

///倍數表///
var Multiple = []int{0, 0, 1, 5, 10, 30, 50, 100}

var Excelroutieng = "/Users/terry_hsiesh/go/game/luckyfortunegod/parsheet/ngparsheet.xlsx"
var Excelroutiefg = "/Users/terry_hsiesh/go/game/luckyfortunegod/parsheet/fgparsheet.xlsx"

var (
	Scatterinfo = [3][]int{
		{3, 4, 5},
		{5, 10, 50},
		{3, 3, 3},
	}

	Scattercombo = [Comboresultnum]int{}
)

func Infomation() {
	fmt.Println("遊戲名稱 : ", GameName)
	fmt.Println("PayLines : ", GamePayLines)
	fmt.Println("轉輪數 : ", Reelamount)
	fmt.Println("橫排數 : ", Col)
	fmt.Println("基本投注額 : ", Bet)
	fmt.Println("獎圖總數 : ", Symbolamount)
	fmt.Println("combo數目0~5combo : ", Comboresultnum)

	fmt.Println("理論值ＲＴＰ ：", ThmRTP)

}
