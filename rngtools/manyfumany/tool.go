package tools

///盤面工具///
import (
	info "info/manyfumany"
	"math/rand"
	table "table/manyfumany"
)

///Reel 轉輪數///
const reelamount int = info.Reelamount

///Column 橫排數///
const col int = info.Col

///基本投注額Bet///
const bet int = info.Bet

///獎圖總數///
const symbolamount = info.Symbolamount

///combo數目0~5combo///
const comboresultnum = info.Comboresultnum

///Scatter 代號///
const scatter = info.Scatter

///NG轉輪表///

var ngstritable = table.NGTablefunc()

///FG轉輪表///
var fgstritable = table.FGTablefunc()

///BG轉輪表///
var bgstritable = table.BGTablefunc().BGStriTable

///線表///
var linetable = table.PublicTablefunc().LineTable

///NG product rng
func BaseGameRng(costlevel int, rtp int) [col][reelamount]int {

	var row1 []int
	var rngresult [col][reelamount]int
	var stritable [][]int

	switch costlevel {
	case 8:
		switch rtp {
		case 95:
			stritable = ngstritable.NGStriTable8rtp95
		case 965:
			stritable = ngstritable.NGStriTable8rtp965
		case 98:
			stritable = ngstritable.NGStriTable8rtp98
		}
	case 18:
		switch rtp {
		case 95:
			stritable = ngstritable.NGStriTable18rtp95
		case 965:
			stritable = ngstritable.NGStriTable18rtp965
		case 98:
			stritable = ngstritable.NGStriTable18rtp98
		}
	case 38:
		switch rtp {
		case 95:
			stritable = ngstritable.NGStriTable38rtp95
		case 965:
			stritable = ngstritable.NGStriTable38rtp965
		case 98:
			stritable = ngstritable.NGStriTable38rtp98
		}
	case 68:
		switch rtp {
		case 95:
			stritable = ngstritable.NGStriTable68rtp95
		case 965:
			stritable = ngstritable.NGStriTable68rtp965
		case 98:
			stritable = ngstritable.NGStriTable68rtp98
		}
	case 88:
		switch rtp {
		case 95:
			stritable = ngstritable.NGStriTable88rtp95
		case 965:
			stritable = ngstritable.NGStriTable88rtp965
		case 98:
			stritable = ngstritable.NGStriTable88rtp98
		}

	}

	for i := 0; i < reelamount; i++ {

		randnumber := rand.Intn(len(stritable[i]))
		row1 = append(row1, randnumber)

	}

	for i := 0; i < reelamount; i++ {
		for j := 0; j < col; j++ {

			rngresult[j][i] = stritable[i][(row1[i]+j)%(len(stritable[i]))]
		}
	}

	return rngresult

}

// func FreeGameallcombo(rtp int, row1 [5]int) [col][reelamount]int {

// 	var rngresult [col][reelamount]int

// 	for i := 0; i < reelamount; i++ {
// 		for j := 0; j < col; j++ {

// 			rngresult[j][i] = fgstritable[i][(row1[i]+j)%(len(fgstritable[i]))]
// 		}
// 	}

// 	return rngresult

// }

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
func FreeGameRng(costlevel int, rtp int) [col][reelamount]int {

	var row1 []int
	var rngresult [col][reelamount]int
	var stritable [][]int

	switch costlevel {
	case 8:
		switch rtp {
		case 95:
			stritable = fgstritable.FGStriTable8rtp95
		case 965:
			stritable = fgstritable.FGStriTable8rtp965
		case 98:
			stritable = fgstritable.FGStriTable8rtp98
		}
	case 18:
		switch rtp {
		case 95:
			stritable = fgstritable.FGStriTable18rtp95
		case 965:
			stritable = fgstritable.FGStriTable18rtp965
		case 98:
			stritable = fgstritable.FGStriTable18rtp98
		}
	case 38:
		switch rtp {
		case 95:
			stritable = fgstritable.FGStriTable38rtp95
		case 965:
			stritable = fgstritable.FGStriTable38rtp965
		case 98:
			stritable = fgstritable.FGStriTable38rtp98
		}
	case 68:
		switch rtp {
		case 95:
			stritable = fgstritable.FGStriTable68rtp95
		case 965:
			stritable = fgstritable.FGStriTable68rtp965
		case 98:
			stritable = fgstritable.FGStriTable68rtp98
		}
	case 88:
		switch rtp {
		case 95:
			stritable = fgstritable.FGStriTable88rtp95
		case 965:
			stritable = fgstritable.FGStriTable88rtp965

		case 98:
			stritable = fgstritable.FGStriTable88rtp98
		}

	}

	for i := 0; i < reelamount; i++ {

		randnumber := rand.Intn(len(stritable[i]))
		row1 = append(row1, randnumber)

	}

	for i := 0; i < reelamount; i++ {
		for j := 0; j < col; j++ {

			rngresult[j][i] = stritable[i][(row1[i]+j)%(len(stritable[i]))]
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
