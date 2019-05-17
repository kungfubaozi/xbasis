package types

type StatisticsType int64

const (
	//时间
	SATTime StatisticsType = iota

	//次数
	SATCounter

	//累加
	SATAccumulation
)
