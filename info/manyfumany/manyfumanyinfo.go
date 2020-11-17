package info

import "fmt"

///遊戲名稱///
const GameName string = "多福多彩"

///Reel 轉輪數///
const Reelamount int = 5

///Column 橫排數///
const Col int = 3

///free game 場數///
const FGsession int = 10

///基本投注額Bet///
const Bet int = 88

///獎圖總數///
const Symbolamount int = 18

///combo數目0~5combo///
const Comboresultnum int = 6

///WILD 代號///
const Wild int = 17

///Scatter 代號///
const Scatter int = 16

///理論值ＲＴＰ///
const ThmRTP float32 = 0.965

///ＲＴＰ95///
const RTP95 float32 = 0.95

///ＲＴＰ98///
const RTP98 float32 = 0.98

///倍數表///
var Multiple = []int{0, 0, 1, 5, 10, 30, 50, 100}

func Infomation() {
	fmt.Println("遊戲名稱 : ", GameName)
	fmt.Println("轉輪數 : ", Reelamount)
	fmt.Println("橫排數 : ", Col)
	fmt.Println("基本投注額 : ", Bet)
	fmt.Println("獎圖總數 : ", Symbolamount)
	fmt.Println("combo數目0~5combo : ", Comboresultnum)

	fmt.Println("理論值ＲＴＰ ：", ThmRTP)

}
