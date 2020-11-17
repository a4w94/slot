package tools

///盤面工具///
import (
	info "info/luckyfortune"
	"math/rand"
	table "table/luckyfortune"
)

const (
	///Reel 轉輪數///
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

///NG轉輪表///
var (
	ngstritable = table.NGTablefunc().NGStriTablertp965

	///FG轉輪表///
	fgstritable = table.FGTablefunc().FGStriTablertp965

	///BG轉輪表///
	bgstritable = table.BGTablefunc().BGStriTable

	///賠付表///
	paytable = table.PublicTablefunc().PayTable

	///線表///
	linetable = table.PublicTablefunc().LineTable
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

func OutputJudge(input Input) Input {
	result := Input{}
	result.Accumulatebet = input.Accumulatebet ///0
	result.Bet = input.Bet///250000
	result.Accumulatepayoff = input.Accumulatepayoff///241250
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

	InputInfo(Choosertp)
	return result

}

func InputInfo(rtp int) {
	p := &ngstritable
	fgpoint := &fgstritable
	switch rtp {
	case 95:

		*p = table.NGTablefunc().NGStriTablertp95
		*fgpoint = table.FGTablefunc().FGStriTablertp95
	case 965:
		*p = table.NGTablefunc().NGStriTablertp965
		*fgpoint = table.FGTablefunc().FGStriTablertp965
	case 98:
		*p = table.NGTablefunc().NGStriTablertp98
		*fgpoint = table.FGTablefunc().FGStriTablertp98
	case 99:
		*p = table.NGTablefunc().NGStriTablertp99
		*fgpoint = table.FGTablefunc().FGStriTablertp99
	}
	//fmt.Println("修改轉輪表ＲＴＰ", rtp)
}

///NG product rng
func BaseGameRng() [col][reelamount]int {

	var row1 []int
	var rngresult [col][reelamount]int

	for i := 0; i < reelamount; i++ {

		randnumber := rand.Intn(len(ngstritable[i]))
		row1 = append(row1, randnumber)

	}

	for i := 0; i < reelamount; i++ {
		for j := 0; j < col; j++ {

			rngresult[j][i] = ngstritable[i][(row1[i]+j)%(len(ngstritable[i]))]
		}
	}

	return rngresult

}

func FreeGameallcombo(rtp int, row1 [5]int) [col][reelamount]int {
	var stritable [][]int
	stritable = table.FGTablefunc().FGStriTablertp965
	var rngresult [col][reelamount]int

	for i := 0; i < reelamount; i++ {
		for j := 0; j < col; j++ {

			rngresult[j][i] = stritable[i][(row1[i]+j)%(len(stritable[i]))]
		}
	}

	return rngresult

}

// ///NG product rng
// func BaseGameWeightRng(stritableweight [reelamount][]int) [col][reelamount]int {
// 	var stritable = ngstritable.NGStriTable8rtp95
// 	var row1 []int
// 	var result [col][reelamount]int

// 	for i := 0; i < reelamount; i++ {

// 		randpostiton := rand.Intn(len(stritableweight[i]))
// 		randnumber := stritableweight[i][randpostiton]
// 		row1 = append(row1, randnumber)

// 	}

// 	for i := 0; i < reelamount; i++ {
// 		for j := 0; j < col; j++ {

// 			result[j][i] = stritable[i][(row1[i]+j)%(len(stritable[i]))]
// 		}
// 	}

// 	return result

// }

///FG product rng
func FreeGameRng() [col][reelamount]int {

	var row1 []int
	var rngresult [col][reelamount]int

	//switch rtp {
	// case 95:
	// 	stritable = fgstritable.FGStriTablertp95
	// case 965:
	// 	stritable = fgstritable.FGStriTablertp965
	// case 98:
	// 	stritable = fgstritable.FGStriTablertp98
	// }
	// stritable = fgstritable.FGStriTablertp965

	for i := 0; i < reelamount; i++ {

		randnumber := rand.Intn(len(fgstritable[i]))
		row1 = append(row1, randnumber)

	}

	for i := 0; i < reelamount; i++ {
		for j := 0; j < col; j++ {

			rngresult[j][i] = fgstritable[i][(row1[i]+j)%(len(fgstritable[i]))]
		}
	}

	return rngresult

}

// ///BG product rng
// func BounsGameRng() [col][reelamount]int {

// 	var row1 []int
// 	var rngresult [col][reelamount]int

// 	for i := 0; i < reelamount; i++ {

// 		randnumber := rand.Intn(len(bgstritable[i]))
// 		row1 = append(row1, randnumber)

// 	}

// 	for i := 0; i < reelamount; i++ {
// 		for j := 0; j < col; j++ {

// 			rngresult[j][i] = bgstritable[i][(row1[i]+j)%(len(fgstritable[i]))]
// 		}
// 	}

// 	return rngresult

// }

///統計盤面各獎圖個數///
func Resultmapf(result [col][reelamount]int) int {

	var scatteramount int
	for i := 0; i < reelamount; i++ {
		var tmp int
		for k := 0; k < col; k++ {
			if result[k][i] == scatter {
				scatteramount++
				tmp++
			}
		}
		if tmp == 0 {
			break
		}

	}

	return scatteramount

}

// func Eachreelmapf(reel int, result [col][reelamount]int) map[int]int {
// 	resultmap := make(map[int]int)

// 	for k := 0; k < col; k++ {
// 		num := result[k][reel]
// 		resultmap[num] = resultmap[num] + 1

// 	}

// 	return resultmap
// }

// func eachreellenf(stritalbe [][]int) []int {
// 	eachreellen := []int{}

// 	for i := 0; i < reelamount; i++ {
// 		eachreellen = append(eachreellen, len(stritalbe[i]))
// 	}
// 	return eachreellen

// }

// func Stritableweightf(weight [][]int) [reelamount][]int {
// 	eachreellen := eachreellenf(ngstritable)
// 	var striweighttable [reelamount][]int

// 	var weightsum [reelamount][]int

// 	for i := 0; i < len(eachreellen); i++ {
// 		arr := make([]int, eachreellen[i])
// 		weightsum[i] = append(weightsum[i], arr...)
// 	}

// 	for k := 0; k < reelamount; k++ {
// 		weightsum[k][0] = weight[k][0]
// 		for i := 1; i < len(weightsum[k]); i++ {
// 			weightsum[k][i] = weight[k][i] + weightsum[k][i-1]
// 		}

// 		for j := 0; j < weightsum[k][0]; j++ {
// 			striweighttable[k] = append(striweighttable[k], 0)
// 		}
// 		for i := 0; i < len(weightsum[k])-1; i++ {

// 			for j := weightsum[k][i]; j < weightsum[k][i+1]; j++ {
// 				striweighttable[k] = append(striweighttable[k], i+1)
// 			}

// 		}
// 	}
// 	//fmt.Println(striweighttable)

// 	return striweighttable

// }
