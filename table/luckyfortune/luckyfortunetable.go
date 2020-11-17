package table

import (
	info "info/luckyfortune"
)

type NGTable struct {
	NGStriTablertp95  [][]int
	NGStriTablertp965 [][]int
	NGStriTablertp98  [][]int
	NGStriTablertp99  [][]int

	NGComboStatistic [symbolamount][comboresultnum]int
}

type FGTable struct {
	FGStriTablertp95  [][]int
	FGStriTablertp965 [][]int
	FGStriTablertp98  [][]int
	FGStriTablertp99  [][]int

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

///獎圖總數+1///
const symbolamount = info.Symbolamount

///combo數目0~5combo///
const comboresultnum = info.Comboresultnum

//const reelamount = info.Reelamount

func PublicTablefunc() PublicTable {
	var Public = PublicTable{}

	/// 派彩表 對應0~5combo///
	Public.PayTable = [][]int{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 50, 200, 1000},
		{0, 0, 0, 25, 100, 300},
		{0, 0, 0, 25, 100, 300},
		{0, 0, 0, 10, 50, 100},
		{0, 0, 0, 10, 50, 100},
		{0, 0, 0, 5, 10, 50},
		{0, 0, 0, 5, 10, 50},
		{0, 0, 0, 5, 10, 50},
		{0, 0, 0, 5, 10, 50},
		{0, 0, 0, 5, 10, 50},
		{0, 0, 0, 5, 10, 50},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}

	Public.LineTable = [][]int{}
	Public.TotalTable = [symbolamount][comboresultnum]int{}

	Public.MultipleTalbe = make([]int, len(info.Multiple)+1)

	return Public
}
func NGTablefunc() NGTable {

	var NGresult = NGTable{}

	///NG轉輪表ＲＴＰ95///
	NGresult.NGStriTablertp95 = [][]int{
		{5, 5, 6, 4, 4, 7, 2, 2, 2, 10, 5, 6, 5, 9, 4, 4, 8, 8, 1, 11, 11, 13, 6, 6, 4, 11, 4, 8, 3, 3, 7, 2, 2, 2, 7, 3, 3, 3, 10, 11, 4, 4, 4, 11, 6, 13, 9, 7, 13, 8, 7, 5, 5, 5, 8, 9, 1, 1, 1},
		{11, 10, 2, 2, 2, 7, 5, 13, 6, 6, 13, 10, 10, 12, 8, 8, 4, 4, 4, 6, 9, 12, 1, 1, 2, 2, 12, 3, 3, 3, 12, 6, 10, 2, 2, 2, 8, 1, 1, 1, 9, 7, 3, 3, 2, 10, 3, 7, 2, 5, 7, 3, 11, 4, 5, 11, 2, 6, 13, 10, 10, 2, 8, 8, 4, 4, 4, 6, 9, 1, 1, 1, 8, 8, 12, 3, 3, 3, 12, 6, 10, 3},
		{4, 9, 6, 1, 2, 12, 10, 8, 12, 2, 2, 2, 5, 5, 5, 11, 10, 12, 8, 10, 2, 13, 1, 1, 1, 13, 6, 8, 4, 10, 2, 11, 7, 12, 6, 2, 11, 5, 6, 13, 3, 3, 3, 13, 6, 7, 12, 9, 10, 13, 8, 9, 13, 6, 8, 1, 1, 1, 9, 6},
		{1, 10, 2, 2, 2, 13, 7, 9, 3, 3, 3, 6, 10, 12, 8, 11, 13, 7, 1, 1, 12, 2, 10, 6, 5, 9, 4, 7, 3, 12, 7, 2, 10, 5, 3, 11, 7, 2, 9, 4, 6, 7, 2, 9, 11, 5, 6, 7, 3, 9, 2, 11, 9, 3, 11, 8},
		{11, 10, 2, 8, 6, 2, 8, 6, 4, 7, 6, 5, 11, 4, 4, 4, 6, 9, 3, 10, 8, 1, 11, 6, 1, 9, 6, 13, 3, 3, 3, 11, 6, 4, 7, 10, 5, 8, 7, 2, 2, 2, 6, 10, 1, 1, 1, 6, 7, 3, 9, 6, 4, 9, 11, 3, 9, 7, 2, 8, 9, 1, 11, 7, 5, 8, 7, 3},
	}

	///NG轉輪表ＲＴＰ96.5///
	NGresult.NGStriTablertp965 = [][]int{
		{5, 5, 6, 4, 4, 7, 2, 2, 2, 10, 5, 6, 5, 9, 4, 4, 8, 8, 13, 11, 11, 13, 6, 6, 4, 11, 4, 8, 3, 3, 7, 2, 2, 2, 7, 3, 3, 3, 10, 11, 4, 4, 4, 11, 6, 13, 9, 7, 13, 8, 7, 5, 5, 5, 8, 9, 1, 1, 1},
		{11, 10, 2, 2, 2, 8, 6, 13, 9, 9, 13, 10, 10, 12, 8, 8, 4, 4, 4, 6, 9, 12, 1, 1, 2, 2, 12, 3, 3, 3, 12, 6, 10, 2, 2, 2, 8, 1, 1, 1, 9, 7, 3, 3, 2, 10, 3, 8, 2, 5, 7, 3, 11, 4, 5, 11, 2, 6, 13, 10, 10, 2, 8, 8, 4, 4, 4, 6, 9, 1, 1, 1, 8, 8, 12, 3, 3, 3, 12, 6, 10, 3},
		{4, 9, 6, 1, 2, 12, 10, 8, 12, 2, 2, 2, 5, 5, 5, 11, 10, 12, 8, 10, 2, 13, 1, 1, 1, 13, 6, 8, 4, 10, 2, 11, 7, 12, 6, 2, 11, 5, 6, 13, 3, 3, 3, 13, 6, 7, 12, 9, 10, 13, 8, 9, 13, 6, 8, 1, 1, 1, 9, 6},
		{1, 10, 2, 2, 2, 13, 7, 9, 3, 3, 3, 6, 10, 12, 8, 11, 13, 7, 1, 1, 12, 2, 10, 6, 5, 9, 4, 7, 3, 12, 7, 2, 10, 5, 3, 11, 7, 2, 9, 4, 6, 7, 2, 9, 11, 5, 6, 7, 3, 9, 2, 11, 9, 3, 11, 8},
		{11, 10, 2, 8, 6, 2, 8, 6, 4, 10, 6, 5, 11, 10, 4, 4, 6, 9, 3, 10, 8, 1, 11, 6, 1, 9, 6, 13, 3, 3, 3, 11, 6, 4, 7, 10, 5, 8, 7, 2, 2, 2, 6, 10, 1, 1, 1, 6, 7, 3, 9, 6, 4, 9, 11, 3, 9, 7, 2, 8, 9, 1, 11, 7, 5, 8, 7, 3},
	}

	///NG轉輪表ＲＴＰ98///
	NGresult.NGStriTablertp98 = [][]int{
		{5, 5, 6, 4, 4, 7, 2, 2, 2, 10, 5, 6, 5, 9, 4, 4, 8, 8, 1, 11, 11, 13, 6, 6, 4, 11, 4, 8, 3, 3, 7, 2, 2, 2, 7, 3, 3, 3, 10, 11, 4, 4, 4, 11, 6, 13, 9, 7, 13, 8, 7, 5, 5, 5, 8, 9, 1, 1, 1},
		{11, 10, 12, 5, 5, 5, 12, 13, 6, 6, 13, 10, 10, 12, 8, 8, 4, 4, 4, 6, 9, 12, 1, 1, 2, 2, 12, 5, 5, 5, 12, 6, 10, 2, 2, 2, 8, 1, 1, 1, 9, 7, 3, 3, 2, 10, 5, 7, 5, 5, 7, 3, 11, 4, 5, 11, 2, 6, 13, 10, 10, 2, 8, 8, 4, 4, 4, 6, 9, 1, 1, 1, 8, 8, 12, 3, 3, 3, 12, 6, 10, 3},
		{4, 9, 6, 7, 5, 12, 10, 8, 12, 4, 2, 2, 5, 5, 5, 11, 10, 12, 8, 10, 2, 13, 1, 1, 1, 13, 6, 8, 4, 10, 2, 11, 7, 12, 6, 2, 11, 5, 6, 13, 3, 3, 3, 13, 6, 7, 12, 9, 10, 13, 8, 9, 13, 6, 8, 1, 1, 1, 9, 6},
		{1, 10, 2, 2, 2, 13, 7, 9, 3, 3, 3, 6, 10, 12, 8, 11, 13, 4, 1, 1, 12, 2, 10, 6, 5, 9, 4, 7, 3, 12, 7, 5, 10, 5, 3, 11, 7, 2, 9, 4, 6, 7, 12, 9, 11, 5, 6, 7, 3, 9, 2, 11, 9, 3, 11, 8},
		{11, 10, 2, 8, 6, 2, 6, 2, 4, 7, 6, 5, 6, 4, 4, 4, 6, 9, 3, 10, 8, 5, 11, 6, 1, 6, 6, 13, 3, 3, 3, 11, 6, 4, 7, 10, 5, 8, 7, 2, 2, 2, 6, 10, 1, 1, 1, 6, 7, 3, 9, 6, 4, 9, 11, 3, 9, 7, 2, 8, 9, 1, 11, 7, 5, 8, 7, 3},
	}

	///NG轉輪表ＲＴＰ99///
	NGresult.NGStriTablertp99 = [][]int{
		{5, 5, 6, 4, 4, 7, 2, 2, 2, 10, 5, 6, 5, 9, 4, 4, 8, 8, 1, 11, 11, 13, 6, 6, 4, 11, 4, 8, 3, 3, 7, 2, 2, 2, 7, 3, 3, 3, 10, 11, 4, 4, 4, 11, 6, 13, 9, 7, 13, 8, 7, 5, 5, 5, 8, 9, 1, 1, 1},
		{11, 10, 2, 2, 2, 7, 5, 13, 6, 6, 13, 10, 10, 12, 8, 8, 4, 4, 4, 6, 9, 12, 1, 1, 2, 2, 12, 3, 3, 3, 12, 6, 10, 2, 2, 2, 8, 1, 1, 1, 9, 7, 3, 3, 2, 10, 3, 7, 2, 5, 7, 3, 11, 4, 5, 11, 2, 6, 13, 10, 10, 2, 8, 8, 4, 4, 4, 6, 9, 1, 1, 1, 8, 8, 12, 3, 3, 3, 12, 6, 10, 3},
		{4, 9, 6, 1, 2, 12, 10, 8, 12, 4, 2, 2, 5, 5, 5, 11, 10, 12, 8, 10, 2, 13, 1, 1, 1, 13, 6, 8, 4, 10, 2, 11, 7, 12, 6, 2, 11, 5, 6, 13, 3, 3, 3, 13, 6, 7, 12, 9, 10, 13, 8, 9, 13, 6, 8, 1, 1, 1, 9, 6},
		{1, 10, 2, 2, 2, 13, 7, 9, 12, 3, 3, 5, 10, 12, 8, 11, 13, 7, 1, 1, 12, 2, 10, 6, 5, 9, 4, 7, 3, 12, 7, 2, 10, 5, 3, 11, 7, 2, 9, 4, 6, 7, 2, 9, 11, 5, 6, 7, 3, 9, 2, 11, 9, 3, 11, 8},
		{11, 10, 2, 8, 6, 2, 8, 4, 4, 7, 5, 5, 11, 4, 4, 4, 5, 9, 3, 10, 8, 1, 11, 6, 1, 9, 6, 13, 3, 3, 3, 11, 6, 4, 7, 10, 5, 8, 7, 2, 2, 2, 6, 10, 1, 1, 1, 6, 7, 3, 9, 6, 4, 9, 11, 3, 9, 7, 2, 8, 9, 1, 11, 7, 5, 8, 7, 3},
	}

	///NG combo統計表初始///
	NGresult.NGComboStatistic = [symbolamount][comboresultnum]int{}

	return NGresult

}

func FGTablefunc() FGTable {
	var FGresult = FGTable{}

	///FG轉輪表ＲＴＰ95///
	FGresult.FGStriTablertp95 = [][]int{
		{4, 4, 4, 2, 2, 3, 3, 3, 5, 5, 5, 2, 2, 2, 4, 4, 4, 3, 3, 3, 2, 2, 4, 4, 4, 1, 1, 1, 4, 4, 4, 1, 1, 1, 13, 3, 4, 4, 3, 3, 4, 4, 4, 3, 3, 3, 4, 4, 4, 5, 5, 3, 3, 4, 3, 3},
		{2, 2, 2, 3, 3, 3, 5, 5, 5, 2, 2, 4, 4, 4, 13, 2, 5, 5, 12, 2, 2, 2, 13, 5, 5, 5, 3, 12, 5, 1, 2, 3, 3, 2, 5, 13, 1, 5, 2, 1, 5, 2, 4, 3, 2, 5, 5, 12, 1, 1, 1, 2, 4, 4, 4, 12, 5, 5},
		{2, 5, 5, 13, 3, 3, 3, 2, 2, 12, 5, 5, 12, 4, 4, 2, 2, 5, 5, 3, 3, 1, 1, 5, 5, 13, 5, 5, 4, 3, 3, 1, 1, 1, 2, 2, 2, 13, 3, 2, 1, 1, 12, 5, 4, 3, 5, 12, 5, 3, 4, 2, 5, 3, 2, 5, 3},
		{1, 4, 4, 4, 2, 2, 5, 5, 5, 2, 2, 4, 4, 3, 1, 12, 3, 5, 5, 13, 3, 3, 5, 2, 2, 13, 4, 4, 4, 12, 2, 1, 1, 5, 5, 5, 12, 4, 4, 2, 5, 3, 2, 1, 3, 5, 5, 3, 1, 1, 1, 3, 4, 5, 5, 2},
		{1, 2, 5, 3, 2, 4, 3, 5, 4, 4, 2, 5, 3, 2, 5, 3, 4, 4, 3, 3, 2, 2, 2, 4, 4, 3, 3, 2, 4, 4, 3, 3, 3, 5, 2, 3, 4, 5, 5, 3, 2, 4, 4, 2, 5, 3, 2, 4, 4, 13, 1, 5, 5, 3, 4, 4, 3, 5, 1, 2, 5, 3, 2, 4, 3, 5, 5, 13},
	}

	///FG轉輪表ＲＴＰ965///
	FGresult.FGStriTablertp965 = [][]int{
		{4, 4, 4, 2, 2, 3, 3, 3, 5, 5, 5, 2, 2, 2, 4, 4, 4, 3, 3, 3, 2, 2, 4, 4, 4, 1, 1, 1, 4, 4, 4, 1, 1, 1, 13, 3, 4, 4, 3, 3, 4, 4, 4, 3, 3, 3, 4, 4, 4, 5, 5, 3, 3, 4, 3, 3},
		{2, 2, 2, 3, 3, 3, 5, 5, 5, 2, 2, 4, 4, 4, 13, 2, 5, 5, 12, 2, 2, 2, 13, 5, 5, 5, 3, 12, 5, 1, 2, 3, 3, 2, 5, 13, 1, 5, 2, 1, 5, 2, 4, 3, 2, 5, 5, 12, 1, 1, 1, 2, 4, 4, 4, 12, 5, 5},
		{2, 5, 5, 13, 3, 3, 3, 2, 2, 12, 5, 5, 12, 4, 4, 2, 2, 5, 5, 3, 3, 1, 1, 5, 5, 13, 5, 5, 4, 3, 3, 1, 1, 1, 2, 2, 2, 13, 3, 2, 1, 1, 12, 5, 4, 3, 5, 12, 5, 3, 4, 2, 5, 3, 2, 5, 3},
		{1, 4, 4, 4, 2, 2, 5, 5, 5, 2, 2, 4, 4, 3, 1, 12, 3, 5, 5, 13, 3, 3, 5, 2, 2, 13, 4, 4, 4, 12, 2, 1, 1, 5, 5, 5, 12, 4, 4, 2, 5, 3, 2, 1, 3, 5, 5, 3, 1, 1, 1, 3, 4, 5, 5, 2},
		{1, 2, 5, 3, 2, 4, 3, 5, 4, 4, 2, 5, 3, 2, 5, 3, 4, 4, 3, 3, 2, 2, 2, 4, 4, 3, 3, 2, 4, 4, 3, 3, 3, 5, 2, 3, 4, 5, 5, 3, 2, 4, 4, 2, 5, 3, 2, 4, 4, 13, 1, 5, 5, 3, 4, 4, 3, 5, 1, 2, 5, 3, 2, 4, 3, 5, 5, 13},
	}

	///FG轉輪表ＲＴＰ98///
	FGresult.FGStriTablertp98 = [][]int{
		{4, 4, 4, 2, 2, 3, 3, 3, 5, 5, 5, 2, 2, 2, 4, 4, 4, 3, 3, 3, 2, 2, 4, 4, 4, 1, 1, 1, 4, 4, 4, 1, 1, 1, 13, 3, 4, 4, 3, 3, 4, 4, 4, 3, 3, 3, 4, 4, 4, 5, 5, 3, 3, 4, 3, 3},
		{2, 2, 2, 3, 3, 3, 5, 5, 5, 2, 2, 4, 4, 4, 13, 2, 5, 5, 12, 2, 2, 2, 13, 5, 5, 5, 3, 12, 5, 1, 2, 3, 3, 2, 5, 13, 1, 5, 2, 1, 5, 2, 4, 3, 2, 5, 5, 12, 1, 1, 1, 2, 4, 4, 4, 12, 5, 5},
		{2, 5, 5, 13, 3, 3, 3, 2, 2, 12, 5, 5, 12, 4, 4, 2, 2, 5, 5, 3, 3, 1, 1, 5, 5, 13, 5, 5, 4, 3, 3, 1, 1, 1, 2, 2, 2, 13, 3, 2, 1, 1, 12, 5, 4, 3, 5, 12, 5, 3, 4, 2, 5, 3, 2, 5, 3},
		{1, 4, 4, 4, 2, 2, 5, 5, 5, 2, 2, 4, 4, 3, 1, 12, 3, 5, 5, 13, 3, 3, 5, 2, 2, 13, 4, 4, 4, 12, 2, 1, 1, 5, 5, 5, 12, 4, 4, 2, 5, 3, 2, 1, 3, 5, 5, 3, 1, 1, 1, 3, 4, 5, 5, 2},
		{1, 2, 5, 3, 2, 4, 3, 5, 4, 4, 2, 5, 3, 2, 5, 3, 4, 4, 3, 3, 2, 2, 2, 4, 4, 3, 3, 2, 4, 4, 3, 3, 3, 5, 2, 3, 4, 5, 5, 3, 2, 4, 4, 2, 5, 3, 2, 4, 4, 13, 1, 5, 5, 3, 4, 4, 3, 5, 1, 2, 5, 3, 2, 4, 3, 5, 5, 13},
	}

	///FG轉輪表ＲＴＰ99///
	FGresult.FGStriTablertp99 = [][]int{
		{4, 4, 4, 2, 2, 3, 3, 3, 5, 5, 5, 2, 2, 2, 4, 4, 4, 3, 3, 3, 2, 2, 4, 4, 4, 1, 1, 1, 4, 4, 4, 1, 1, 1, 13, 3, 4, 4, 3, 3, 4, 4, 4, 3, 3, 3, 4, 4, 4, 5, 5, 3, 3, 4, 3, 3},
		{2, 2, 2, 3, 3, 3, 5, 5, 5, 2, 2, 4, 4, 4, 13, 2, 5, 5, 12, 2, 2, 2, 13, 5, 5, 5, 3, 12, 5, 1, 2, 3, 3, 2, 5, 13, 1, 5, 2, 1, 5, 2, 4, 3, 2, 5, 5, 12, 1, 1, 1, 2, 4, 4, 4, 12, 5, 5},
		{2, 5, 5, 13, 3, 3, 3, 2, 2, 12, 5, 5, 12, 4, 4, 2, 2, 5, 5, 3, 3, 1, 1, 5, 5, 13, 5, 5, 4, 3, 3, 1, 1, 1, 2, 2, 2, 13, 3, 2, 1, 1, 12, 5, 4, 3, 5, 12, 5, 3, 4, 2, 5, 3, 2, 5, 3},
		{1, 4, 4, 4, 2, 2, 5, 5, 5, 2, 2, 4, 4, 3, 1, 12, 3, 5, 5, 13, 3, 3, 5, 2, 2, 13, 4, 4, 4, 12, 2, 1, 1, 5, 5, 5, 12, 4, 4, 2, 5, 3, 2, 1, 3, 5, 5, 3, 1, 1, 1, 3, 4, 5, 5, 2},
		{1, 2, 5, 3, 2, 4, 3, 5, 4, 4, 2, 5, 3, 2, 5, 3, 4, 4, 3, 3, 2, 2, 2, 4, 4, 3, 3, 2, 4, 4, 3, 3, 3, 5, 2, 3, 4, 5, 5, 3, 2, 4, 4, 2, 5, 3, 2, 4, 4, 13, 1, 5, 5, 3, 4, 4, 3, 5, 1, 2, 5, 3, 2, 4, 3, 5, 5, 13},
	}

	/// 權重表///
	FGresult.FGWeightTable = [][]int{}

	///FG combo統計表初始///
	FGresult.FGComboStatistic = [symbolamount][comboresultnum]int{}

	return FGresult

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
