package utils

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

// CustomTime 自定义时间类型
type CustomTime struct {
	time.Time
}

// Now 获取当前时间
func Now() CustomTime {
	return CustomTime{time.Now()}
}

// Format 返回 yyyy-mm-dd H:i:s 格式字符串
func (ct CustomTime) Format() string {
	if ct.IsZero() {
		return ""
	}
	return ct.Time.Format(time.DateTime)
}

// String 实现 Stringer 接口
func (ct CustomTime) String() string {
	return ct.Format()
}

// MarshalJSON 实现 JSON 序列化
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	if ct.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + ct.Format() + `"`), nil
}

// UnmarshalJSON 实现 JSON 反序列化
func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == "null" || str == `""` {
		ct.Time = time.Time{}
		return nil
	}

	// 去除引号
	str = strings.Trim(str, `"`)

	// 尝试解析多种格式
	formats := []string{
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05Z07:00",
		time.RFC3339,
		time.RFC3339Nano,
	}

	var err error
	for _, layout := range formats {
		ct.Time, err = time.Parse(layout, str)
		if err == nil {
			return nil
		}
	}

	return fmt.Errorf("invalid time format: %s", str)
}

// Value 实现 driver.Valuer 接口 (数据库写入)
func (ct CustomTime) Value() (driver.Value, error) {
	if ct.IsZero() {
		return nil, nil
	}
	return ct.Time, nil
}

// Scan 实现 sql.Scanner 接口 (数据库读取)
func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		ct.Time = time.Time{}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		ct.Time = v
		return nil
	case []byte:
		return ct.UnmarshalJSON(v)
	case string:
		return ct.UnmarshalJSON([]byte(v))
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
}

// GormDataType 定义 GORM 数据库类型
func (CustomTime) GormDataType() string {
	return "datetime"
}

// GormDBDataType 定义 GORM 数据库类型 (v2 版本)
func (CustomTime) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "mysql", "sqlite":
		return "datetime"
	case "postgres":
		return "timestamp"
	default:
		return "datetime"
	}
}
