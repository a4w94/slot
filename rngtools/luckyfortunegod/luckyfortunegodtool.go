package tools

///盤面工具///
import (
	"fmt"
	info "info/luckyfortunegod"
	"math/rand"
	table "table/luckyfortunegod"
)

///Reel 轉輪數///
const (
	reelamount int = info.Reelamount

	///Column 橫排數///
	col int = info.Col

	///基本投注額Bet///
	bet int = info.Bet

	///獎圖總數///
	symbolamount = info.Symbolamount

	///combo數目0~5combo///
	comboresultnum = info.Comboresultnum

	///Scatter 代號///
	scatter = info.Scatter

	///錢幣 代號///
	money = info.Money

	///風控／／／

	payofflimit              = info.Payofflimit
	levelupper               = info.LevelUpper
	levellower               = info.LevelLower
	initialaccumulatebet     = info.InitialAccumulatebet
	initialaccumulatepayoff  = info.InitialAccumulatepayoff
	initialbet               = info.InitialBet
	firstaccumulatebetcheck  = info.FirstAccumulatebetCheck
	secondaccumulatebetcheck = info.SecondAccumulatebetCheck
)

var (
	///NG轉輪表///
	Ngstritable = table.NGStriTablertp965

	///FG轉輪表///
	Fgstritable = table.FGStriTablertp965

	///BG轉輪表///
	bgstritable = table.BGTablefunc().BGStriTable

	///線表///
	linetable = table.PublicTablefunc().LineTable
)

var (
	///NG轉輪表///
	ngstritable965 = table.NGStriTablertp965

	ngstritable95 = table.NGStriTablertp95

	ngstritable99 = table.NGStriTablertp99

	///FG轉輪表///
	fgstritable965 = table.FGStriTablertp965
	fgstritable95  = table.FGStriTablertp95
	fgstritable99  = table.FGStriTablertp99

	///FG金幣機率表
	eachreelmoneyrate = table.EachReelMoneyRate
	moneyweight       = table.Moneyweight
	multipleweight    = table.Multipleweight
)

type Judgelevel struct {
	Choosertp int
	Input
}

type Input struct {
	Bet              float32
	Accumulatebet    float32
	Accumulatepayoff float32
	Currentrtp       float32
	Choosertp        int
}

func PointStritable() {
	ngrtp95, ngrtp965, ngrtp99 := &ngstritable95, &ngstritable965, &ngstritable99
	*ngrtp95, *ngrtp965, *ngrtp99 = table.NGStriTablertp95, table.NGStriTablertp965, table.NGStriTablertp99

	fgrtp95, fgrtp965, fgrtp99 := &fgstritable95, &fgstritable965, &fgstritable99
	*fgrtp95, *fgrtp965, *fgrtp99 = table.FGStriTablertp95, table.FGStriTablertp965, table.FGStriTablertp99

	Ngstri, Fgstri := &Ngstritable, &Fgstritable
	*Ngstri, *Fgstri = table.NGStriTablertp965, table.FGStriTablertp965

}

func OutputJudge(input Input) Input {
	result := Input{}
	result.Bet = input.Bet
	result.Accumulatebet = input.Accumulatebet
	result.Accumulatepayoff = input.Accumulatepayoff
	result.Currentrtp = input.Currentrtp
	Choosertp := 965

	//fmt.Println("傳入值 ", result)

	if input.Accumulatebet > firstaccumulatebetcheck {
		if input.Currentrtp > levellower && input.Currentrtp < levelupper {
			result.Accumulatebet = initialaccumulatebet
			result.Accumulatepayoff = initialaccumulatepayoff
			result.Bet = initialbet
		} else {
			if input.Accumulatebet > secondaccumulatebetcheck {
				result.Accumulatebet = initialaccumulatebet
				result.Accumulatepayoff = initialaccumulatepayoff
				result.Bet = initialbet
			} else {
				switch {
				case input.Currentrtp <= levellower:

					Choosertp = 99
				case input.Currentrtp >= levelupper:

					Choosertp = 95
				case input.Currentrtp > levellower && input.Currentrtp < levelupper:

					Choosertp = 965
				}

			}
		}

	} else {
		switch {
		case input.Currentrtp <= levellower:

			Choosertp = 99

		case input.Currentrtp >= levelupper:

			Choosertp = 95
		case input.Currentrtp > levellower && input.Currentrtp < levelupper:

			Choosertp = 965
		}
	}
	//fmt.Println("回傳值", result)
	Choosertp = 965
	result.Choosertp = Choosertp

	InputInfo(Choosertp)
	return result

}

func InputInfo(rtp int) {

	p := &Ngstritable
	fgpoint := &Fgstritable
	switch rtp {
	case 95:

		*p = ngstritable95
		*fgpoint = fgstritable95
	case 965:
		*p = ngstritable965
		*fgpoint = fgstritable965
	// case 98:
	// 	*p = table.NGTablefunc().NGStriTablertp98
	// 	*fgpoint = table.FGTablefunc().FGStriTablertp98
	case 99:
		*p = ngstritable99
		*fgpoint = fgstritable99
	}
}

///NG product rng
func BaseGameRng() [col][reelamount]int {

	var row1 []int
	var rngresult [col][reelamount]int

	for i := 0; i < reelamount; i++ {

		randnumber := rand.Intn(len(Ngstritable[i]))
		row1 = append(row1, randnumber)

	}

	for i := 0; i < reelamount; i++ {
		for j := 0; j < col; j++ {

			rngresult[j][i] = Ngstritable[i][(row1[i]+j)%(len(Ngstritable[i]))]
		}
	}

	return rngresult

}

// ///傳入盤面 回傳線數總結果///
// func Linetablef(result [col][reelamount]int) [linenum][reelamount]int {

// 	var resultlinearr [linenum][reelamount]int
// 	for i := 0; i < linenum; i++ {
// 		for j := 0; j < reelamount; j++ {

// 			resultlinearr[i][j] = result[(linetable[i][j])][j]

// 		}

// 	}
// 	return resultlinearr
// }

///FG product rng
func FreeGameRng() [col][reelamount]int {

	var row1 []int
	var rngresult [col][reelamount]int

	for i := 0; i < reelamount; i++ {

		randnumber := rand.Intn(len(Fgstritable[i]))
		row1 = append(row1, randnumber)

	}

	for i := 0; i < reelamount; i++ {
		for j := 0; j < col; j++ {

			rngresult[j][i] = Fgstritable[i][(row1[i]+j)%(len(Fgstritable[i]))]
		}
	}

	return rngresult

}

///金幣產盤
func FreeGameMoneyRng(inputbet int, moneyamount int) ([col][reelamount]int, int) {
	multiprate := float32(1) / float32(4)
	var multiple int

	var rngresult [col][reelamount]int

	for i := 0; i < col; i++ {
		for k := 0; k < reelamount; k++ {
			seed := rand.Float32()
			if seed < eachreelmoneyrate[k] {
				rngresult[i][k] = ProductMoney(inputbet)
			}
		}
	}
	fmt.Println("金幣盤面", rngresult)

	if moneyamount != 0 && moneyamount != 15 {
		seed := rand.Float32()
		if seed < multiprate {
			multiple = ProductMultiple()
			fmt.Println("get multiple", multiple)

		}
	}

	fmt.Println(rngresult)

	return rngresult, multiple
}

func InsertMultilple(rngresult [col][reelamount]int, multiple int) [col][reelamount]int {
	reelseed := rand.Intn(reelamount)

	for t := 0; t < col; t++ {
		if rngresult[t][reelseed] == 0 {
			rngresult[t][reelseed] = multiple
			break
		}
	}

	fmt.Println("乘倍放置盤面", rngresult)
	return rngresult
}

func ProductMoney(inputbet int) int {
	var money int

	seed := rand.Intn(moneyweight[1][len(moneyweight[1])-1])

	for i := 1; i < len(moneyweight[1]); i++ {
		if seed >= moneyweight[1][i-1] && seed < moneyweight[1][i] {
			money = inputbet * moneyweight[0][i]
		}
	}

	return money
}

func ProductMultiple() int {
	var multipe int
	seed := rand.Intn(multipleweight[1][len(multipleweight[1])-1])

	for i := 1; i < len(multipleweight[1]); i++ {
		if seed >= multipleweight[1][i-1] && seed < multipleweight[1][i] {
			multipe = multipleweight[0][i]
		}
	}

	return multipe
}

///統計盤面各獎圖個數///
func Resultmapf(result [col][reelamount]int) int {

	var scatteramount int

	for i := 0; i < reelamount; i++ {

		for k := 0; k < col; k++ {
			if result[k][i] == scatter {
				scatteramount++

			}
		}

	}

	return scatteramount

}

/// 丟進freegame rng return lockwild position

type Moneymap struct {
	MoneyPosition [col][]int
	MoneyMap      [col][]int
}

func LockMoneySymbol(freeresult [col][reelamount]int) Moneymap {
	var result = Moneymap{}

	for k := 0; k < col; k++ {

		for i := 0; i < reelamount; i++ {
			if freeresult[k][i] != 0 {
				result.MoneyPosition[k] = append(result.MoneyPosition[k], i)
				result.MoneyMap[k] = append(result.MoneyMap[k], freeresult[k][i])
			}
		}
	}
	return result

}

func ResultMoneymap(freeresult [col][reelamount]int) int {
	var moneysymbolamount int
	for k := 0; k < col; k++ {

		for i := 1; i < reelamount; i++ {
			if freeresult[k][i] != 0 {
				moneysymbolamount++
			}
		}
	}
	return moneysymbolamount
}
