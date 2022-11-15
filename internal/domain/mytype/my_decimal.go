package mytype

import (
	"database/sql/driver"
	"errors"

	"github.com/shopspring/decimal"
)

// 废弃
//MyTime 自定义时间
type MyDecimal decimal.Decimal

func (t *MyDecimal) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的decimal字符串
	str := string(data)
	d, err := decimal.NewFromString(str)
	if err != nil {
		return nil
	}
	*t = MyDecimal(d)
	return err
}

func (t MyDecimal) MarshalJSON() ([]byte, error) {
	str := decimal.Decimal(t).String()
	return []byte(str), nil
}

func (t MyDecimal) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	d := decimal.Decimal(t)
	return d.String(), nil
}

func (t *MyDecimal) Scan(v interface{}) error {
	/* typeOfResp := reflect.TypeOf(v)
	fmt.Println(fmt.Sprintf(" resp type is %s, kind is %s",
		typeOfResp, typeOfResp.Kind())) */
	switch vt := v.(type) {
	case decimal.Decimal:
		// 字符串转成 time.Time 类型
		*t = MyDecimal(vt)
	case string: //没有进来，db是 decimal
		// 字符串转成 time.Time 类型
		d, err := decimal.NewFromString(vt)
		if err != nil {
			return err
		}
		*t = MyDecimal(d)
	case []uint8: //db是 decimal
		// 字符串转成 time.Time 类型
		str := string(vt)
		d, err := decimal.NewFromString(str)
		if err != nil {
			return err
		}
		*t = MyDecimal(d)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *MyDecimal) String() string {
	return decimal.Decimal(*t).String()
}
func (ts *MyDecimal) ToDecimal() decimal.Decimal {
	return decimal.Decimal(*ts)
}
