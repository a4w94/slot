package tools

///盤面工具///
import (
	info "info/bear"
	"math/rand"
	table "table/bear"
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

///線數///
const linenum = info.Linenum

///NG轉輪表///
var ngstritable = table.NGTablefunc().NGStriTable

///NG轉輪表95///
var ngweight95 = table.NGTablefunc().NGWeightTable95

///NG轉輪表98///
var ngweight98 = table.NGTablefunc().NGWeightTable98

///FG轉輪表///
var fgstritable = table.FGTablefunc().FGStriTable

///線表///
var linetable = table.PublicTablefunc().LineTable

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

///NG product rng
func BaseGameWeightRng(stritableweight [reelamount][]int) [col][reelamount]int {
	var stritable = ngstritable
	var row1 []int
	var result [col][reelamount]int

	for i := 0; i < reelamount; i++ {

		randpostiton := rand.Intn(len(stritableweight[i]))
		randnumber := stritableweight[i][randpostiton]
		row1 = append(row1, randnumber)

	}

	for i := 0; i < reelamount; i++ {
		for j := 0; j < col; j++ {

			result[j][i] = stritable[i][(row1[i]+j)%(len(stritable[i]))]
		}
	}

	return result

}

///FG product rng
func FreeGameRng() [col][reelamount]int {

	var row1 []int
	var rngresult [col][reelamount]int

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

///統計盤面各獎圖個數///
func Resultmapf(result [col][reelamount]int) map[int]int {

	resultmap := make(map[int]int)

	for i := 0; i < col; i++ {
		for _, num := range result[i] {
			resultmap[num] = resultmap[num] + 1

		}
	}
	return resultmap

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

func eachreellenf(stritalbe [][]int) []int {
	eachreellen := []int{}

	for i := 0; i < reelamount; i++ {
		eachreellen = append(eachreellen, len(stritalbe[i]))
	}
	return eachreellen

}

func Stritableweightf(weight [][]int) [reelamount][]int {
	eachreellen := eachreellenf(ngstritable)
	var striweighttable [reelamount][]int

	var weightsum [reelamount][]int

	for i := 0; i < len(eachreellen); i++ {
		arr := make([]int, eachreellen[i])
		weightsum[i] = append(weightsum[i], arr...)
	}

	for k := 0; k < reelamount; k++ {
		weightsum[k][0] = weight[k][0]
		for i := 1; i < len(weightsum[k]); i++ {
			weightsum[k][i] = weight[k][i] + weightsum[k][i-1]
		}

		for j := 0; j < weightsum[k][0]; j++ {
			striweighttable[k] = append(striweighttable[k], 0)
		}
		for i := 0; i < len(weightsum[k])-1; i++ {

			for j := weightsum[k][i]; j < weightsum[k][i+1]; j++ {
				striweighttable[k] = append(striweighttable[k], i+1)
			}

		}
	}
	//fmt.Println(striweighttable)

	return striweighttable

}
