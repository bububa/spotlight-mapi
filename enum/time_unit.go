package enum

// TimeUnit 时间维度
type TimeUnit string

const (
	// TimeUnit_DAY 分天
	TimeUnit_DAY TimeUnit = "DAY"
	// TimeUnit_HOUR 分时
	TimeUnit_HOUR TimeUnit = "HOUR"
	// TimeUnit_SUMMARY 汇总,默认分天
	TimeUnit_SUMMARY TimeUnit = "SUMMARY"
)
