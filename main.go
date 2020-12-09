package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app.HideHelp = true
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

var app = &cli.App{
	Name:        AppName,
	Usage:       "实体生成器",
	Description: "啊吧啊吧啊吧",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    Host,
			Aliases: []string{"h"},
			Usage:   "主机",
			Value:   "127.0.0.1",
		},
		&cli.Int64Flag{
			Name:    Port,
			Aliases: []string{"P"},
			Usage:   "端口",
			Value:   3306,
		},
		&cli.StringFlag{
			Name:    Output,
			Aliases: []string{"o"},
			Usage:   "生成文件夹",
			Value:   "./",
		},
		&cli.StringFlag{
			Name:    Package,
			Aliases: []string{"pkg"},
			Usage:   "包名称",
			Value:   "db",
		},
		&cli.StringFlag{
			Name:     Database,
			Aliases:  []string{"db"},
			Usage:    "数据库名称",
			Required: true,
		},
		&cli.StringFlag{
			Name:    TableName,
			Aliases: []string{"t"},
			Usage:   "表名称",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    User,
			Aliases: []string{"u"},
			Usage:   "用户名称",
			Value:   "root",
		},
		&cli.StringFlag{
			Name:    Password,
			Aliases: []string{"p"},
			Usage:   "密码",
			Value:   "",
		},
	},
	Action: ModelByInputArgs,
}
