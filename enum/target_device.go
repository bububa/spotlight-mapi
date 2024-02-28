package enum

// TargetDevice 定向设备
type TargetDevice string

const (
	// TargetDevice_IOS ios
	TargetDevice_IOS TargetDevice = "ios"
	// TargetDevice_ANDROID android
	TargetDevice_ANDROID TargetDevice = "android"
	// TargetDevice_ALL all
	TargetDevice_ALL TargetDevice = "all"
)
