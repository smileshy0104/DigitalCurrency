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
// 返回值是一个*gorm.DB实例，用于后续的数据库操作。
func connectMysql() *gorm.DB {
	// 配置MySQL连接参数
	username := "root"
	password := "123456"
	host := "8.153.166.227"
	port := 3309
	dbName := "mscoin"

	// 构造MySQL的DSN（Data Source Name），用于数据库连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbName)

	// 使用gorm库打开与MySQL的连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("连接数据库失败, error=%v", err)
	}
	return db
}

// Result 结构体用于描述字段的多种标签和属性。
type Result struct {
	Field        string // 字段的原始名称
	MessageField string // 用于在错误消息中引用字段名称
	Type         string // 字段的数据类型
	Gorm         string // 字段在Gorm库中的标签名称
	Json         string // 字段在JSON序列化时的标签名称
	Form         string // 字段在处理表单数据时的标签名称
	JsonForm     string // 结合了JSON和表单处理的标签
}

// StructResult 结构体用于汇总结构体的验证结果。
type StructResult struct {
	StructName string    // 结构体的名称
	Result     []*Result // 字段的验证结果
}

// MessageResult 结构体用于汇总错误消息的验证结果。
type MessageResult struct {
	MessageName string    // 错误消息的名称
	Result      []*Result // 与该错误消息相关的字段验证结果
}

// GenModel 生成模型和protobuf消息定义。
func GenModel(table, name string) {
	GenStruct(table, name)
	//GenProtoMessage(table, name)
}

// GenStruct 根据数据库表结构生成Go结构体代码。
// table: 数据库表名。
// structName: 生成的Go结构体名。
func GenStruct(table, structName string) {
	// 连接MySQL数据库。
	db := connectMysql()

	// 定义一个Result类型的切片来存储查询结果。
	var results []*Result
	// 执行SQL查询以获取表结构信息，并将结果扫描到results变量中。
	if err := db.Raw(fmt.Sprintf("DESCRIBE %s", table)).Scan(&results).Error; err != nil {
		// 如果查询失败，则记录错误信息并终止程序执行。
		log.Fatalf("查询表结构失败: %v", err)
	}

	// 遍历查询结果，处理每个字段以生成相应的Go结构体字段。
	for _, v := range results {
		// 保存原始字段名。
		field := v.Field
		// 将字段名转换为大驼峰命名法。
		v.Field = Name(field)
		// 根据数据库字段类型获取对应的Go类型。
		v.Type = getType(v.Type)
		// 生成JSON标签。
		v.Json = fmt.Sprintf("`json:\"%s\"`", field)
		// 生成包含form和json标签的字符串。
		v.JsonForm = fmt.Sprintf("`json:\"%s\" form:\"%s\"`", field, field)
		// 生成Gorm标签，指定数据库列名。
		v.Gorm = fmt.Sprintf("`gorm:\"column:%s\"`", field)
	}

	// 调用函数生成并保存Go结构体文件。
	generateStructFile(structName, results)
	// 调用函数生成并保存Gorm结构体文件。
	generateGormStructFile(structName, results)
}

// GenProtoMessage 生成protobuf消息定义。
func GenProtoMessage(table, messageName string) {
	db := connectMysql()

	var results []*Result
	if err := db.Raw(fmt.Sprintf("DESCRIBE %s", table)).Scan(&results).Error; err != nil {
		log.Fatalf("查询表结构失败: %v", err)
	}

	for _, v := range results {
		v.MessageField = TFName(v.Field)
		v.Type = getMessageType(v.Type)
	}

	generateProtoFile(messageName, results)
}

// generateStructFile 生成结构体文件
// 该函数根据给定的结构体名称和结果集，生成一个Go语言结构体定义文件
// 参数:
//
//	structName - 结构体的名称
//	results - 一个Result对象的切片，包含要生成的结构体的字段信息
func generateStructFile(structName string, results []*Result) {
	// 创建StructResult实例，封装结构体名称和字段结果集
	sr := StructResult{StructName: structName, Result: results}

	// 解析模板文件，用于生成结构体代码
	tmpl, err := template.ParseFiles("./struct.tpl")
	if err != nil {
		log.Fatalf("解析模板失败: %v", err)
	}

	// 构造文件路径，生成的文件位于./gen目录下，文件名基于结构体名称
	filePath := fmt.Sprintf("./gen/%s.go", strings.ToLower(structName))

	// 确保./gen目录存在，如果不存在则创建
	if err := ensureDir("./gen"); err != nil {
		log.Fatalf("创建目录失败: %v", err)
	}

	// 创建结构体定义文件
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("创建文件失败: %v", err)
	}
	defer file.Close()

	// 使用模板和结构体结果集填充文件，生成结构体定义
	if err := tmpl.Execute(file, sr); err != nil {
		log.Fatalf("执行模板失败: %v", err)
	}
}

// generateGormStructFile 根据结构体名称和查询结果生成GORM结构体文件。
// structName 是结构体的名称，将用于生成文件名和结构体名。
// results 是查询结果的切片，包含要生成结构体的字段信息。
func generateGormStructFile(structName string, results []*Result) {
	// 创建 StructResult 实例，封装结构体名称和查询结果。
	sr := StructResult{StructName: structName, Result: results}

	// 解析模板文件，用于生成结构体代码。
	tmpl, err := template.ParseFiles("./struct_gorm.tpl")
	if err != nil {
		// 如果解析模板失败，记录错误并退出。
		log.Fatalf("解析模板失败: %v", err)
	}

	// 根据结构体名称生成文件路径。
	filePath := fmt.Sprintf("./gen/%s_gorm.go", strings.ToLower(structName))
	file, err := os.Create(filePath)
	if err != nil {
		// 如果创建文件失败，记录错误并退出。
		log.Fatalf("创建文件失败: %v", err)
	}
	// 确保文件在函数结束后关闭。
	defer file.Close()

	// 使用解析的模板和结构体结果生成代码并写入文件。
	if err := tmpl.Execute(file, sr); err != nil {
		// 如果执行模板失败，记录错误并退出。
		log.Fatalf("执行模板失败: %v", err)
	}
}

// generateProtoFile 生成protobuf消息定义文件。
func generateProtoFile(messageName string, results []*Result) {
	sr := MessageResult{MessageName: messageName, Result: results}
	tmpl, err := template.ParseFiles("./message.tpl")
	if err != nil {
		log.Fatalf("解析模板失败: %v", err)
	}

	filePath := fmt.Sprintf("./gen/%s.proto", strings.ToLower(messageName))
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("创建文件失败: %v", err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, sr); err != nil {
		log.Fatalf("执行模板失败: %v", err)
	}
}

// ensureDir 确保目录存在，如果不存在则创建。
func ensureDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.Mkdir(dir, 0755)
	}
	return nil
}

// getMessageType 根据数据库字段类型返回对应的消息类型。
func getMessageType(t string) string {
	switch {
	case strings.Contains(t, "bigint"):
		return "int64"
	case strings.Contains(t, "varchar"), strings.Contains(t, "text"):
		return "string"
	case strings.Contains(t, "tinyint"), strings.Contains(t, "int") && !strings.Contains(t, "bigint"):
		return "int32"
	case strings.Contains(t, "double"), strings.Contains(t, "decimal"):
		return "double"
	default:
		return "unknown"
	}
}

// getType 根据数据库字段类型返回对应的Go类型。
func getType(t string) string {
	switch {
	case strings.Contains(t, "bigint"):
		return "int64"
	case strings.Contains(t, "varchar"), strings.Contains(t, "text"):
		return "string"
	case strings.Contains(t, "tinyint"):
		return "int"
	case strings.Contains(t, "int") && !strings.Contains(t, "bigint"):
		return "int"
	case strings.Contains(t, "double"), strings.Contains(t, "decimal"):
		return "float64"
	default:
		return "unknown"
	}
}

// TFName 将下划线风格的字符串转换为小驼峰风格的字符串。
func TFName(name string) string {
	var sb strings.Builder
	names := strings.Split(name, "_")

	for i, part := range names {
		if i == 0 {
			sb.WriteString(strings.ToLower(part)) // 首个部分小写
		} else {
			sb.WriteString(strings.Title(part)) // 其他部分大写
		}
	}

	return sb.String()
}

// Name 将下划线分隔的名字转换为驼峰命名。
// 该函数接收一个字符串参数name，将其按'_下划线'进行分割，
// 并将每个分割后的部分首字母大写，最终合并为一个新的字符串返回。
// 这个函数主要用于将数据库列名或者其他以下划线分隔的字符串
// 转换为符合Go语言命名规范的驼峰命名。
func Name(name string) string {
	var sb strings.Builder
	names := strings.Split(name, "_")

	for _, part := range names {
		sb.WriteString(strings.Title(part)) // 每个部分大写
	}

	return sb.String()
}
