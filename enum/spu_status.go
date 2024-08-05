package enum

// SpuStatus 状态
type SpuStatus int

const (
	// SpuStatus_CAN_BIND 可绑定
	SpuStatus_CAN_BIND SpuStatus = 1
	// SpuStatus_AUDITING 审核中
	SpuStatus_AUDITING SpuStatus = 2
	// SpuStatus_REJECTED 审核不通过
	SpuStatus_REJECTED SpuStatus = 3
)
