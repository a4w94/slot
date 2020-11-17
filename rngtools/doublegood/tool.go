package tools

///盤面工具///
import (
	info "info/doublegood"
	"math/rand"
	table "table/doublegood"
)

///Reel 轉輪數///
const (
	reelamount int = info.Reelamount

	///Column 橫排數///
	col int = info.Col

	///基本投注額Bet///
	bet int = info.Bet

	///線數///
	linenum int = info.Linenum

	///獎圖總數///
	symbolamount = info.Symbolamount

	///combo數目0~5combo///
	comboresultnum = info.Comboresultnum

	///Scatter 代號///
	scatter = info.Scatter

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
	Ngstritable = table.NGTablefunc().NGStriTablertp965

	///FG轉輪表///
	Fgstritable = table.FGTablefunc().FGStriTablertp965

	///BG轉輪表///
	bgstritable = table.BGTablefunc().BGStriTable

	///賠付表///
	paytable = table.PublicTablefunc().PayTable

	///線表///
	linetable = table.PublicTablefunc().LineTable
)

var (
	///NG轉輪表///
	ngstritable965 = table.NGTablefunc().NGStriTablertp965

	ngstritable95 = table.NGTablefunc().NGStriTablertp95

	ngstritable99 = table.NGTablefunc().NGStriTablertp99

	///FG轉輪表///
	fgstritable965 = table.FGTablefunc().FGStriTablertp965
	fgstritable95  = table.FGTablefunc().FGStriTablertp95
	fgstritable99  = table.FGTablefunc().FGStriTablertp99
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

func Background() {
	table.Getexcelparsheet()
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

	result.Choosertp = Choosertp
	//fmt.Println("轉輪表", Choosertp)

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
	//fmt.Println(Ngstritable)

	//row1 := []int{}
	row1 := make([]int, reelamount, 5)
	var rngresult [col][reelamount]int

	for i := 0; i < reelamount; i++ {

		randnumber := rand.Intn(len(Ngstritable[i]))
		row1[i] = randnumber

	}

	for i := 0; i < reelamount; i++ {
		for j := 0; j < col; j++ {

			rngresult[j][i] = Ngstritable[i][(row1[i]+j)%(len(Ngstritable[i]))]
		}
	}

	return rngresult

}

///傳入盤面 回傳線數總結果///
func Linetablef(result [col][reelamount]int) [linenum][reelamount]int {

	var resultlinearr [linenum][reelamount]int
	for i := 0; i < linenum; i++ {
		for j := 0; j < reelamount; j++ {

			resultlinearr[i][j] = result[(linetable[i][j])][j]

		}

	}
	return resultlinearr
}

///FG product rng
func FreeGameRng() [col][reelamount]int {

	//row1 := []int{}
	row1 := make([]int, reelamount, 5)
	var rngresult [col][reelamount]int

	for i := 0; i < reelamount; i++ {

		randnumber := rand.Intn(len(Fgstritable[i]))
		row1[i] = randnumber
		//fmt.Println("種子", randnumber)

	}

	for i := 0; i < reelamount; i++ {
		for j := 0; j < col; j++ {

			rngresult[j][i] = Fgstritable[i][(row1[i]+j)%(len(Fgstritable[i]))]
		}
	}

	return rngresult

}

// }

///統計盤面各獎圖個數///
func Resultmapf(result [col][reelamount]int) int {

	var scatteramount int
	for i := 0; i < reelamount; i++ {
		//var tmp int
		for k := 0; k < col; k++ {
			if result[k][i] == scatter {
				scatteramount++
				//tmp++
			}
		}
		// if tmp == 0 {
		// 	break
		// }

	}

	return scatteramount

}

func Eachreellen(gamestatus string, rtp int) int {
	var stritable = [reelamount][]int{}
	if gamestatus == "NG" {
		switch rtp {
		case 95:
			stritable = ngstritable95
		case 965:
			stritable = ngstritable965
		case 99:
			stritable = ngstritable99

		}

	} else if gamestatus == "FG" {
		switch rtp {
		case 95:
			stritable = fgstritable95
		case 965:
			stritable = fgstritable965
		case 99:
			stritable = fgstritable99

		}
	}

	var reelallcombo int = 1
	for i := 0; i < reelamount; i++ {
		reelallcombo *= len(stritable[i])
	}
	return reelallcombo
}
