package tools

import (
	"errors"
	"reflect"
)

// TODO 同时也可以获取默认值
// Default 通过反射将传入数据的字段设置为默认值。
// 参数 data 是一个指向结构体的指针。
// 返回错误如果传入的数据不是指针。
func Default(data any) error {
	// 获取传入数据的类型和值。
	typeOf := reflect.TypeOf(data)
	valueOf := reflect.ValueOf(data)
	// 检查传入的数据是否为指针。
	if typeOf.Kind() != reflect.Pointer {
		return errors.New("must be pointer")
	}
	// 处理结构体字段。
	ele := typeOf.Elem()
	valueEle := valueOf.Elem()
	for i := 0; i < ele.NumField(); i++ {
		field := ele.Field(i)
		value := valueEle.Field(i)
		// 根据字段类型设置默认值。
		kind := field.Type.Kind()
		if kind == reflect.Int {
			value.Set(defaultInt())
		}
		if kind == reflect.Int32 {
			value.Set(defaultInt32())
		}
		if kind == reflect.Int64 {
			value.Set(defaultInt64())
		}
		if kind == reflect.String {
			value.Set(defaultString())
		}
		if kind == reflect.Float64 {
			value.Set(defaultFloat64())
		}
		if kind == reflect.Float32 {
			value.Set(defaultFloat32())
		}
	}
	return nil
}

// defaultString 返回一个默认值为空字符串的reflect.Value。
func defaultString() reflect.Value {
	var i = ""
	return reflect.ValueOf(i)
}

// defaultInt 返回一个默认值为0的int类型的reflect.Value。
func defaultInt() reflect.Value {
	var i int = 0
	return reflect.ValueOf(i)
}

// defaultInt32 返回一个默认值为0的int32类型的reflect.Value。
func defaultInt32() reflect.Value {
	var i int32 = 0
	return reflect.ValueOf(i)
}

// defaultInt64 返回一个默认值为0的int64类型的reflect.Value。
func defaultInt64() reflect.Value {
	var i int64 = 0
	return reflect.ValueOf(i)
}

// defaultFloat64 返回一个默认值为0的float64类型的reflect.Value。
func defaultFloat64() reflect.Value {
	var i float64 = 0
	return reflect.ValueOf(i)
}

// defaultFloat32 返回一个默认值为0的float32类型的reflect.Value。
func defaultFloat32() reflect.Value {
	var i float32 = 0
	return reflect.ValueOf(i)
}
