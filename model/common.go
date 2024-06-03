package model

import (
	"bytes"
	"strconv"
)

// Uint64 support string quoted number in json
type Uint64 uint64

// UnmarshalJSON implement json Unmarshal interface
func (u64 *Uint64) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	b = bytes.ReplaceAll(b, []byte(","), []byte(""))
	i, _ := strconv.ParseUint(string(b), 10, 64)
	*u64 = Uint64(i)
	return
}

func (u64 Uint64) Value() uint64 {
	return uint64(u64)
}

func (u64 Uint64) String() string {
	return strconv.FormatUint(uint64(u64), 10)
}

// JSONUint64 support string quoted number in json and marshal to string
type JSONUint64 uint64

func JSONUint64FromUint64(v uint64) JSONUint64 {
	return JSONUint64(v)
}

// UnmarshalJSON implement json Unmarshal interface
func (u64 *JSONUint64) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	i, _ := strconv.ParseUint(string(b), 10, 64)
	*u64 = JSONUint64(i)
	return
}

func (u64 JSONUint64) MarshalJSON() ([]byte, error) {
	return []byte(`"` + strconv.FormatUint(uint64(u64), 10) + `"`), nil
}

func (u64 JSONUint64) Value() uint64 {
	return uint64(u64)
}

func (u64 JSONUint64) String() string {
	return strconv.FormatUint(u64.Value(), 10)
}

// Int64 support string quoted number in json
type Int64 int64

// UnmarshalJSON implement json Unmarshal interface
func (i64 *Int64) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	i, _ := strconv.ParseInt(string(b), 10, 64)
	*i64 = Int64(i)
	return
}

func (i64 Int64) Value() int64 {
	return int64(i64)
}

func (i64 Int64) String() string {
	return strconv.FormatInt(int64(i64), 10)
}

// JSONInt64 support string quoted number in json and marshal to string
type JSONInt64 int64

func JSONInt64FromInt64(v int64) JSONInt64 {
	return JSONInt64(v)
}

// UnmarshalJSON implement json Unmarshal interface
func (i64 *JSONInt64) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	b = bytes.ReplaceAll(b, []byte(","), []byte(""))
	i, _ := strconv.ParseInt(string(b), 10, 64)
	*i64 = JSONInt64(i)
	return
}

func (i64 JSONInt64) MarshalJSON() ([]byte, error) {
	return []byte(`"` + strconv.FormatInt(int64(i64), 10) + `"`), nil
}

func (i64 JSONInt64) Value() int64 {
	return int64(i64)
}

// Int support string quoted number in json
type Int int

// UnmarshalJSON implement json Unmarshal interface
func (i *Int) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	b = bytes.ReplaceAll(b, []byte(","), []byte(""))
	v, _ := strconv.Atoi(string(b))
	*i = Int(v)
	return
}

func (i Int) Value() int {
	return int(i)
}

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

// JSONInt support string quoted number in json and marshal to string
type JSONInt int

func JSONIntFromInt(v int) JSONInt {
	return JSONInt(v)
}

// UnmarshalJSON implement json Unmarshal interface
func (i *JSONInt) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	b = bytes.ReplaceAll(b, []byte(","), []byte(""))
	v, _ := strconv.Atoi(string(b))
	*i = JSONInt(v)
	return
}

func (i JSONInt) MarshalJSON() ([]byte, error) {
	return []byte(`"` + strconv.FormatInt(int64(i), 10) + `"`), nil
}

func (i JSONInt) Value() int {
	return int(i)
}

// Float64 support string quoted number in json
type Float64 float64

// UnmarshalJSON implement json Unmarshal interface
func (f64 *Float64) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	var isPercent bool
	if b[len(b)-1] == '%' {
		b = b[0 : len(b)-2]
		isPercent = true
	}
	b = bytes.ReplaceAll(b, []byte(","), []byte(""))
	i, _ := strconv.ParseFloat(string(b), 64)
	if isPercent {
		i /= 100
	}
	*f64 = Float64(i)
	return
}

func (f64 Float64) Value() float64 {
	return float64(f64)
}

func (f64 Float64) String(prec int) string {
	return strconv.FormatFloat(float64(64), 'f', prec, 64)
}

// Bool support number/string in json
type Bool bool

// UnmarshalJSON implement json Unmarshal interface
func (bl *Bool) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	var ret bool
	str := string(b)
	if str == "true" {
		ret = true
	} else if str == "false" {
		ret = false
	} else if i, err := strconv.ParseInt(str, 10, 64); err != nil {
		return err
	} else if i == 0 {
		ret = false
	} else {
		ret = true
	}
	*bl = Bool(ret)
	return
}

func (b Bool) Value() bool {
	return bool(b)
}

func (b Bool) String() string {
	if b {
		return "true"
	}
	return "false"
}

type CodeNamePair struct {
	// Code code
	Code string `json:"code,omitempty"`
	// Name 名称
	Name string `json:"name,omitempty"`
	// Children 子节点
	Children []CodeNamePair `json:"children,omitempty"`
}

// PageDTO 分页信息
type PageDTO struct {
	// PageIndex 页码
	PageIndex int64 `json:"page_index,omitempty"`
	// PageSize 每页大小
	PageSize int64 `json:"page_size,omitempty"`
}

// PageRespDTO 分页信息
type PageRespDTO struct {
	// PageIndex 页码
	PageIndex int64 `json:"page_index,omitempty"`
	// TotalCount 总数量
	TotalCount int64 `json:"total_count,omitempty"`
}
