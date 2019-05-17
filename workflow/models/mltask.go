package models

//统计服务
type StatisticsTask struct {
	*Info
	RefKey string `bson:"ref_key" json:"ref_key"`
	Type   int64  `bson:"type" json:"type"`
}

//时间序列预测
type TimeSeriesPredictionTask struct {
	*Info
	RefKey string `bson:"ref_key" json:"ref_key"`
}
