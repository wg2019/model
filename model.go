package main

import (
	"github.com/urfave/cli/v2"
	"log"
)

func ModelByInputArgs(c *cli.Context) error {
	input, err := GetInputParam(c)
	if err != nil {
		log.Fatalf("GetInputParam fail, err: %v", err)
	}
	Input = input
	if err := InitDbByInputArgs(); err != nil {
		log.Fatalf("InitDbByInputArgs fail, err: %v", err)
	}
	tables, err := GetTables()
	if err != nil {
		log.Fatalf("GetTables fail, err: %v", err)
	}
	for _, tableInfo := range tables {
		tableFields := GetTableFields(tableInfo.Name)
		tableMetaData := new(TableMetaData)
		tableMetaData.Package = Input.Package
		tableMetaData.TableStruct = tableInfo
		tableMetaData.Fields = tableFields
		content, err := GenerateTemplate(tableMetaData)
		if err != nil {
			log.Fatalf("GenerateTemplate fail, err: %v", err)
		}
		fileName := FileName(tableInfo.Name)
		err = SaveFile(Input.Output, fileName, content)
		if err != nil {
			log.Fatalf("SaveFile fail, err: %v", err)
		}
		log.Printf("%s \n%s\n==========", fileName, content)
	}
	return nil
}
