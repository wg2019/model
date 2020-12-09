package main

import "strings"

const (
	// AppName 程序名称
	AppName = "model"
	// Host 主机
	Host = "host"
	// Port 端口号
	Port = "port"
	// Output 导出路径
	Output = "output"
	// Package 包名称
	Package = "package"
	// Database 数据库名称
	Database = "database"
	// TableName 表名，多表名用英文","分隔
	TableName = "table"
	// User 用户名称
	User = "user"
	// Password 用户密码
	Password = "password"
)

// InputStruct 输入结构
type InputStruct struct {
	Host      string
	Port      int64
	Output    string
	Package   string
	Database  string
	TableName string
	User      string
	Password  string
}

// Input 输入数据
var Input = new(InputStruct)

// TableStruct 表结构
type TableStruct struct {
	Name    string `gorm:"column:Name"`
	Comment string `gorm:"column:Comment"`
}

// TableField 表字段信息
type TableField struct {
	Field      string `gorm:"column:Field"`
	Type       string `gorm:"column:Type"`
	Null       string `gorm:"column:Null"`
	Key        string `gorm:"column:Key"`
	Default    string `gorm:"column:Default"`
	Extra      string `gorm:"column:Extra"`
	Privileges string `gorm:"column:Privileges"`
	Comment    string `gorm:"column:Comment"`
}

func (t *TableField) GoType() string {
	typeArr := strings.Split(t.Type, "(")

	switch typeArr[0] {
	case "int":
		return "int"
	case "integer":
		return "int"
	case "mediumint":
		return "int"
	case "bit":
		return "int"
	case "year":
		return "int"
	case "smallint":
		return "int"
	case "tinyint":
		return "int"
	case "bigint":
		return "int64"
	case "decimal":
		return "float32"
	case "double":
		return "float32"
	case "float":
		return "float32"
	case "real":
		return "float32"
	case "numeric":
		return "float32"
	case "timestamp":
		return "time.Time"
	case "datetime":
		return "time.Time"
	case "time":
		return "time.Time"
	default:
		return "string"
	}
}

// TableMetaData 表元数据
type TableMetaData struct {
	Package string
	*TableStruct
	Fields []*TableField
}

var tableModelTemplate = `package {{.Package}}
{{$tableName := CamelizeStr .Name true}}
type {{$tableName}} struct {
{{- range .Fields}}
	{{CamelizeStr .Field true}} {{.GoType}} ` + "// {{.Comment}}" + `
{{- end}}
}
{{$firstChar := FirstCharacter .Name}}
func ({{$firstChar}} *{{$tableName }}) TableName() string {
	return "{{.Name}}"
}
`
