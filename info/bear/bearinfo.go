package info

import "fmt"

///遊戲名稱///
const GameName string = "台灣黑熊"

///Reel 轉輪數///
const Reelamount int = 5

///Column 橫排數///
const Col int = 4

///基本投注額Bet///
const Bet int = 40

///獎圖總數///
const Symbolamount int = 12

///combo數目0~5combo///
const Comboresultnum int = 6

///線數///
const Linenum int = 40

///理論值ＲＴＰ///
const ThmRTP float32 = 0.9566

///ＲＴＰ95///
const RTP95 float32 = 0.95

///ＲＴＰ98///
const RTP98 float32 = 0.98

///倍數表///
var Multiple = []int{0, 10, 20, 30, 40, 50}

func Infomation() {

	fmt.Println("轉輪數 : ", Reelamount)
	fmt.Println("橫排數 : ", Col)
	fmt.Println("基本投注額 : ", Bet)
	fmt.Println("獎圖總數 : ", Symbolamount)
	fmt.Println("combo數目0~5combo : ", Comboresultnum)
	fmt.Println("線數 : ", Linenum)
	fmt.Println("理論值ＲＴＰ ：", ThmRTP)

}
