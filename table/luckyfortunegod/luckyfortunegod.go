package table

import (
	"fmt"
	info "info/luckyfortunegod"
	"os"
	"runtime"
	"strconv"

	"github.com/Luxurioust/excelize"
)

var (
	NGStriTablertp95  [reelamount][]int
	NGStriTablertp965 [reelamount][]int
	NGStriTablertp99  [reelamount][]int

	FGStriTablertp95  [reelamount][]int
	FGStriTablertp965 [reelamount][]int
	FGStriTablertp99  [reelamount][]int
)

var (
	///NG combo統計表初始///
	NGComboStatistic = [symbolamount][comboresultnum]int{}
	///FG combo統計表初始///
	FGComboStatistic = [symbolamount][comboresultnum]int{}
)

type NGTable struct {
	NGStriTablertp95  [reelamount][]int
	NGStriTablertp965 [reelamount][]int
	NGStriTablertp98  [reelamount][]int
	NGStriTablertp99  [reelamount][]int

	NGComboStatistic [symbolamount][comboresultnum]int
}

type FGTable struct {
	FGStriTablertp95  [reelamount][]int
	FGStriTablertp965 [reelamount][]int
	FGStriTablertp98  [reelamount][]int
	FGStriTablertp99  [reelamount][]int

	FGWeightTable    [][]int
	FGComboStatistic [symbolamount][comboresultnum]int
}

type BGTable struct {
	BGStriTable [][]int

	BGWeightTable    [][]int
	BGComboStatistic [symbolamount][comboresultnum]int
}

type PublicTable struct {
	PayTable      [][]int
	LineTable     [][]int
	TotalTable    [symbolamount][comboresultnum]int
	MultipleTalbe []int
}

///每輪金幣出現機率
var EachReelMoneyRate = []float32{
	float32(2) / float32(20), //R1
	float32(2) / float32(20), //R2
	float32(4) / float32(20), //R3
	float32(5) / float32(20), //R4
	float32(5) / float32(20), //R5
}

var (
	Moneyweight = [2][]int{
		{0, 2, 3, 4, 5, 6, 8, 10},
		{0, 1, 2, 3, 4, 5, 6, 7},
	}

	Multipleweight = [2][]int{
		{0, 2, 3, 4, 5, 6, 8, 10},
		{0, 1, 2, 3, 4, 5, 6, 7},
	}
)

///獎圖總數+1///
const symbolamount = info.Symbolamount

///combo數目0~5combo///
const comboresultnum = info.Comboresultnum

const reelamount = info.Reelamount

func PublicTablefunc() PublicTable {

	// fmt.Println("use public table")
	// fmt.Println("誰在調用", PrintCallfuncName())
	// fmt.Println()
	var Public = PublicTable{}

	/// 派彩表 對應0~5combo///
	Public.PayTable = [][]int{

		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 100, 200, 1000}, //M1
		{0, 0, 0, 50, 100, 500},   //M2
		{0, 0, 0, 50, 100, 500},   //M3
		{0, 0, 0, 30, 60, 300},    //M4
		{0, 0, 0, 30, 60, 300},    //M5
		{0, 0, 0, 20, 40, 200},    //M6
		{0, 0, 0, 10, 20, 100},    //A
		{0, 0, 0, 10, 20, 100},    //K
		{0, 0, 0, 5, 10, 50},      //Q
		{0, 0, 0, 5, 10, 50},      //J
		{0, 0, 0, 5, 10, 50},      //TE
		{0, 0, 0, 0, 0, 0},        //Sc
		{0, 0, 0, 0, 0, 0},        //Wild
	}

	Public.LineTable = [][]int{}
	Public.TotalTable = [symbolamount][comboresultnum]int{}

	Public.MultipleTalbe = make([]int, len(info.Multiple)+1)

	return Public
}

func BGTablefunc() BGTable {
	var BGresult = BGTable{}

	///BG轉輪表///
	BGresult.BGStriTable = [][]int{}

	///BG 權重表///
	BGresult.BGWeightTable = [][]int{}

	///BG combo統計表初始///
	BGresult.BGComboStatistic = [symbolamount][comboresultnum]int{}

	return BGresult

}

func Getexcelparsheet() {
	xlsxfg, err := excelize.OpenFile(info.Excelroutiefg)
	// fmt.Println("free useexcel")
	// fmt.Println("誰在調用", PrintCallfuncName(), PrintCallfuncNameUpper())
	// fmt.Println()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	xlsxng, err := excelize.OpenFile(info.Excelroutieng)
	// fmt.Println("free useexcel")
	// fmt.Println("誰在調用", PrintCallfuncName(), PrintCallfuncNameUpper())
	// fmt.Println()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var rtproutie = []string{"95", "965", "99"}

	for i := 0; i < len(rtproutie); i++ {
		rowng, err := xlsxng.GetRows("rtp" + rtproutie[i])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		stritable := [reelamount][]int{}

		for i := 0; i < len(rowng); i++ {
			for k := 0; k < len(rowng[i]); k++ {
				if rowng[i][k] == "" {
					continue
				} else {
					element, err := strconv.Atoi(rowng[i][k])
					if err != nil {
						panic(err)
					}
					stritable[k] = append(stritable[k], element)
				}

			}
		}
		switch rtproutie[i] {
		case "95":
			temp := &NGStriTablertp95
			*temp = stritable
		case "965":
			temp := &NGStriTablertp965
			*temp = stritable
		case "99":
			temp := &NGStriTablertp99
			*temp = stritable
		}

	}

	for i := 0; i < len(rtproutie); i++ {
		rowfg, err := xlsxfg.GetRows("rtp" + rtproutie[i])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		stritable := [reelamount][]int{}

		for i := 0; i < len(rowfg); i++ {
			for k := 0; k < len(rowfg[i]); k++ {
				if rowfg[i][k] == "" {
					continue
				} else {
					element, err := strconv.Atoi(rowfg[i][k])
					if err != nil {
						panic(err)
					}
					stritable[k] = append(stritable[k], element)
				}

			}
		}
		switch rtproutie[i] {
		case "95":
			temp := &FGStriTablertp95
			*temp = stritable
		case "965":
			temp := &FGStriTablertp965
			*temp = stritable
		case "99":
			temp := &FGStriTablertp99
			*temp = stritable
		}

	}

}

func PrintCallfuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func PrintCallfuncNameUpper() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}
