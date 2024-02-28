package util

func StringsJoin(strs ...string) string {
	var n int
	for i := 0; i < len(strs); i++ {
		n += len(strs[i])
	}
	if n <= 0 {
		return ""
	}
	builder := NewStringsBuilder()
	builder.Grow(n)
	for _, s := range strs {
		builder.WriteString(s)
	}
	ret := builder.String()
	ReleaseStringsBuilder(builder)
	return ret
}
