package info

import "fmt"

const (
	///遊戲名稱///
	GameName string = "招財進寶"

	///Reel 轉輪數///
	Reelamount int = 5

	///Column 橫排數///
	Col int = 3

	///free game 場數///
	FGsession int = 10

	///基本投注額Bet///
	Bet int = 50

	///獎圖總數+1 沒有零號///
	Symbolamount int = 14

	///combo數目0~5combo///
	Comboresultnum int = 6

	///WILD 代號///
	Wild int = 12

	///Scatter 代號///
	Scatter int = 13

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
	InitialAccumulatepayoff = 241250
	InitialBet              = 250000

	FirstAccumulatebetCheck  = 800000
	SecondAccumulatebetCheck = 1200000
)

///倍數表///
var Multiple = []int{0, 10, 20, 30, 40, 50}

func Infomation() {
	fmt.Println("遊戲名稱 : ", GameName)
	fmt.Println("轉輪數 : ", Reelamount)
	fmt.Println("橫排數 : ", Col)
	fmt.Println("基本投注額 : ", Bet)
	fmt.Println("獎圖總數 : ", Symbolamount)
	fmt.Println("combo數目0~5combo : ", Comboresultnum)

	fmt.Println("理論值ＲＴＰ ：", ThmRTP)

}
