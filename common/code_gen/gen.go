package code_gen

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
	"text/template"
)

// connectMysql 用于建立与MySQL数据库的连接。
// 该函数无需参数。
// 返回值是一个*gorm.DB实例，用于后续的数据库操作。
// 此函数中包含了数据库连接字符串的构造以及与gorm库的交互。
func connectMysql() *gorm.DB {
	// 配置MySQL连接参数
	username := "root"      // 账号
	password := "123456"    // 密码
	host := "8.153.166.227" // 数据库地址，可以是Ip或者域名
	port := 3309            // 数据库端口
	Dbname := "mscoin"      // 数据库名
	// 构造MySQL的DSN（Data Source Name），用于数据库连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	var err error
	// 使用gorm库打开与MySQL的连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		// 如果连接失败，抛出panic
		panic("连接数据库失败, error=" + err.Error())
	}
	// 返回数据库连接实例
	return db
}

// Result 结构体用于描述字段的多种标签和属性。
// 它封装了一个字段在不同场景下的名称或标签，例如数据库、JSON序列化、表单等。
type Result struct {
	Field        string // Field 表示字段的原始名称。
	MessageField string // MessageField 用于在错误消息中引用字段名称。
	Type         string // Type 表示字段的数据类型。
	Gorm         string // Gorm 是字段在Gorm库中的标签名称。
	Json         string // Json 是字段在JSON序列化时的标签名称。
	Form         string // Form 是字段在处理表单数据时的标签名称。
	JsonForm     string // JsonForm 结合了JSON和表单处理的标签，用于特定场景下。
}

// StructResult 结构体用于汇总结构体的验证结果。
// 它包含了结构体名称和该结构体所有字段的验证结果。
type StructResult struct {
	StructName string    // StructName 表示结构体的名称。
	Result     []*Result // Result 是一个Result对象的切片，每个Result对象代表一个字段的验证结果。
}

// MessageResult 结构体用于汇总错误消息的验证结果。
// 它通过错误消息名称来组织验证结果，便于根据不同的错误消息进行结果的分类和处理。
type MessageResult struct {
	MessageName string    // MessageName 表示错误消息的名称。
	Result      []*Result // Result 是一个Result对象的切片，每个Result对象代表与该错误消息相关的字段验证结果。
}

func GenModel(table string, name string) {
	GenStruct(table, name)
	//GenProtoMessage(table, name)
}

// GenStruct 根据数据库表结构生成Go结构体代码。
// 参数:
//
//	table: 数据库表名。
//	structName: 生成的Go结构体名称。
func GenStruct(table string, structName string) {
	// 连接MySQL数据库。
	db := connectMysql()

	// 定义一个Result类型的切片来存储表结构信息。
	var results []*Result

	// 执行SQL查询，获取表结构信息。
	db.Raw(fmt.Sprintf("describe %s", table)).Scan(&results)

	// 遍历查询结果，处理每个字段的信息。
	for _, v := range results {
		// 将表字段名转换为符合Go命名规范的字段名。
		field := v.Field    // 数据表字段名称
		name := Name(field) // 表字段 aa_bb  字段名 AaBb

		// 使用驼峰命名法转换字段名。
		tfName := TFName(v.Field) //驼峰命名  aaBb

		// 更新字段名和类型，以及相关的标签信息。
		v.Field = name
		v.Type = getType(v.Type)
		v.Json = "`json:\"" + tfName + "\"`"
		v.JsonForm = "`json:\"" + tfName + "\" from:\"" + tfName + "\"`"
		v.Gorm = "`gorm:\"column:" + field + "\"`"
	}

	// 解析模板文件，用于生成结构体代码。
	tmpl, err := template.ParseFiles("./struct.tpl")
	log.Println(err)

	// 解析另一个模板文件，用于生成包含GORM标签的结构体代码。
	tmpl1, err := template.ParseFiles("./struct_gorm.tpl")
	log.Println(err)

	// 创建一个StructResult实例，包含结构体名称和字段信息。
	sr := StructResult{StructName: structName, Result: results}

	// 检查./gen目录是否存在，如果不存在则创建。
	_, err = os.Stat("./gen")
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir("./gen", 0666)
		}
	}

	// 创建Go结构体代码文件。
	file, err := os.Create("./gen/" + strings.ToLower(structName) + ".go")
	log.Println(err)

	// 执行模板，生成代码到文件。
	tmpl.Execute(file, sr)
	defer file.Close()

	// 创建包含GORM标签的Go结构体代码文件。
	file1, err := os.Create("./gen/" + strings.ToLower(structName) + "_gorm.go")
	defer file1.Close()
	log.Println(err)

	// 执行模板，生成代码到文件。
	tmpl1.Execute(file1, sr)
}

func GenProtoMessage(table string, messageName string) {
	db := connectMysql()
	var results []*Result
	db.Raw(fmt.Sprintf("describe %s", table)).Scan(&results)
	for _, v := range results {
		v.MessageField = TFName(v.Field)
		v.Type = getMessageType(v.Type)
	}
	var fm template.FuncMap = make(map[string]any)
	fm["Add"] = func(v int, add int) int {
		return v + add
	}
	t := template.New("message.tpl")
	t.Funcs(fm)
	tmpl, err := t.ParseFiles("./message.tpl")
	log.Println(err)
	sr := MessageResult{MessageName: messageName, Result: results}
	_, err = os.Stat("./gen")
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir("./gen", 0666)
		}
	}
	file, err := os.Create("./gen/" + strings.ToLower(messageName) + ".proto")
	defer file.Close()
	log.Println(err)
	err = tmpl.Execute(file, sr)
	log.Println(err)
}

// getMessageType 根据数据库字段类型返回对应的消息类型。
// 参数 t 是数据库字段类型的字符串表示。
// 返回值是字段在消息中的类型表示。
func getMessageType(t string) string {
	if strings.Contains(t, "bigint") {
		return "int64"
	}
	if strings.Contains(t, "varchar") {
		return "string"
	}
	if strings.Contains(t, "text") {
		return "string"
	}
	if strings.Contains(t, "tinyint") {
		return "int32"
	}
	if strings.Contains(t, "int") &&
		!strings.Contains(t, "tinyint") &&
		!strings.Contains(t, "bigint") {
		return "int32"
	}
	if strings.Contains(t, "double") {
		return "double"
	}
	if strings.Contains(t, "decimal") {
		return "double"
	}
	return ""
}

// getType 根据数据库字段类型返回对应的Go类型。
// 参数 t 是数据库字段类型的字符串表示。
// 返回值是字段在Go语言中的类型表示。
func getType(t string) string {
	if strings.Contains(t, "bigint") {
		return "int64"
	}
	if strings.Contains(t, "varchar") {
		return "string"
	}
	if strings.Contains(t, "text") {
		return "string"
	}
	if strings.Contains(t, "tinyint") {
		return "int"
	}
	if strings.Contains(t, "int") &&
		!strings.Contains(t, "tinyint") &&
		!strings.Contains(t, "bigint") {
		return "int"
	}
	if strings.Contains(t, "double") {
		return "float64"
	}
	if strings.Contains(t, "decimal") {
		return "float64"
	}
	return ""
}

// TFName 将下划线风格的字符串转换为小驼峰风格的字符串。
// 例如，将"user_name"转换为"userName"。
func TFName(name string) string {
	var names = name[:]    // 将输入字符串转换为切片进行操作
	isSkip := false        // 控制是否跳过下一个字符（用于处理下划线后的字符）
	var sb strings.Builder // 使用strings.Builder高效构建结果字符串

	for index, value := range names {
		if index == 0 {
			// 首字母转小写
			s := names[:index+1]
			s = strings.ToLower(s)
			sb.WriteString(s)
			continue
		}

		if isSkip {
			// 如果已经处理了下划线，则跳过当前循环
			isSkip = false
			continue
		}

		// 当遇到下划线时，将下划线后的字符转换为大写，并设置isSkip为true以跳过该字符
		if value == 95 { // ASCII码中95对应'_'
			s := names[index+1 : index+2] // 取下划线后的字符
			s = strings.ToUpper(s)        // 转换为大写
			sb.WriteString(s)
			isSkip = true // 设置标记，下次循环跳过该字符
			continue
		} else {
			// 普通字符直接追加到结果中
			s := names[index : index+1]
			sb.WriteString(s)
		}
	}

	return sb.String() // 返回最终的小驼峰格式字符串
}

// Name 将下划线风格的字符串转换为大驼峰风格的字符串。
// 例如，将"user_name"转换为"UserName"。
func Name(name string) string {
	var names = name[:]    // 将输入字符串转换为切片进行操作
	isSkip := false        // 控制是否跳过下一个字符（用于处理下划线后的字符）
	var sb strings.Builder // 使用strings.Builder高效构建结果字符串

	for index, value := range names {
		if index == 0 {
			// 首字母转大写
			s := names[:index+1]
			s = strings.ToUpper(s)
			sb.WriteString(s)
			continue
		}

		if isSkip {
			// 如果已经处理了下划线，则跳过当前循环
			isSkip = false
			continue
		}

		// 当遇到下划线时，将下划线后的字符转换为大写，并设置isSkip为true以跳过该字符
		if value == 95 { // ASCII码中95对应'_'
			s := names[index+1 : index+2] // 取下划线后的字符
			s = strings.ToUpper(s)        // 转换为大写
			sb.WriteString(s)
			isSkip = true // 设置标记，下次循环跳过该字符
			continue
		} else {
			// 普通字符直接追加到结果中
			s := names[index : index+1]
			sb.WriteString(s)
		}
	}

	return sb.String() // 返回最终的大驼峰格式字符串
}
