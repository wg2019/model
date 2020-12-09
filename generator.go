package main

import (
	"bytes"
	"text/template"
)

// GetTables 获取表列表
func GetTables() (tables []*TableStruct, err error) {
	tables = make([]*TableStruct, 0)
	if Input.TableName == "" {
		Db.Raw("SELECT TABLE_NAME as Name,TABLE_COMMENT as Comment FROM information_schema.TABLES WHERE table_schema='" + Input.Database + "';").Find(&tables)
	} else {
		Db.Raw("SELECT TABLE_NAME as Name,TABLE_COMMENT as Comment FROM information_schema.TABLES WHERE TABLE_NAME IN (" + Input.TableName + ") AND table_schema='" + Input.Database + "';").Find(&tables)
	}
	return tables, nil
}

//获取所有字段信息
func GetTableFields(tableName string) []*TableField {
	fields := make([]*TableField, 0)
	Db.Raw("show FULL COLUMNS from " + tableName + ";").Find(&fields)
	return fields
}

func GenerateTemplate(templateData *TableMetaData) ([]byte, error) {
	t, err := template.New("tableTemplate").Funcs(template.FuncMap{
		"CamelizeStr":    CamelizeStr,
		"FirstCharacter": FirstCharacter,
	}).Parse(tableModelTemplate)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, templateData); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
