package workflow

//请求任务
type apiTask struct {
	*basicModel
	Url             string   `bson:"url" json:"url"`
	RequestMethod   string   `bson:"request_method" json:"request_method"`
	ResponseHeaders []string `bson:"response_headers" json:"response_headers"`
	RequestHeaders  []string `bson:"request_headers" json:"request_headers"`
	RequestModel    string   `bson:"request_model" json:"request_model"`
}

//统计任务
type statisticsTask struct {
}

//数据访问任务
type databaseTask struct {
	*basicModel
	Url      string `bson:"url" json:"url"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	Port     int64  `bson:"port" json:"port"`
}

//时间序列预测
type timeseqTask struct {
	*basicModel
	Source string `bson:"source" json:"source"`
}
