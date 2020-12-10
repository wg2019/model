package main

import (
	"fmt"
	"log"
	"strings"
)

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
	// Desc 描述
	Desc = "desc"
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
	Desc      string
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
	case "boolean":
		return "bool"
	case "tinyint":
		return "int8"
	case "smallint", "year":
		return "int16"
	case "integer", "mediumint", "int":
		return "int32"
	case "bigint":
		return "int64"
	case "date", "timestamp without time zone", "timestamp with time zone", "time with time zone", "time without time zone",
		"timestamp", "datetime", "time":
		return "string"
	case "bytea",
		"binary", "varbinary", "tinyblob", "blob", "mediumblob", "longblob":
		return "[]byte"
	case "text", "character", "character varying", "tsvector", "bit", "bit varying", "money", "json", "jsonb", "xml", "point", "interval", "line", "ARRAY",
		"char", "varchar", "tinytext", "mediumtext", "longtext":
		return "string"
	case "real":
		return "float32"
	case "numeric", "decimal", "double precision", "float", "double":
		return "float64"
	default:
		return "string"
	}
}

func (t *TableField) GetDesc() string {
	//获取字段gorm描述
	desc, ok := FieldDescMap[Input.Desc]
	if !ok {
		desc = FieldDescMap["json"]
	}
	log.Printf("desc: %s", Input.Desc)
	return fmt.Sprintf(desc, t.Field)
}

var FieldDescMap = map[string]string{
	"gorm": "`gorm:\"column:%s\"`",
	"json": "`json:\"%s\"`",
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
	{{CamelizeStr .Field true}} {{.GoType}} {{.GetDesc}}	` + "// {{.Comment}}" + `
{{- end}}
}
{{$firstChar := FirstCharacter .Name}}
func ({{$firstChar}} *{{$tableName }}) TableName() string {
	return "{{.Name}}"
}
`
