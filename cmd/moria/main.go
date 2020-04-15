package main

import (
	"fmt"
	"os"

	"github.com/astro-bug/gondor/webapi/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/k0kubun/pp"
	"github.com/urfave/cli"
)

var (
	app   *cli.App
	flags []cli.Flag
)

func init() {
	app = cli.NewApp()
	app.HideHelp = true
	app.Version = "0.1.0"
	app.Usage = "后台开发和管理辅助工具"
	flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose, v",
			Usage: "输出详细信息",
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "配置文件路径",
			Value: "settings.yml",
		},
	}
}

func main() {
	app.Commands = []cli.Command{
		{
			Name:   "reverse",
			Usage:  "从数据库导出对应的Model代码",
			Flags:  flags,
			Action: ReverseAction,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Errorf(err.Error())
	}
}

func ReverseAction(ctx *cli.Context) (err error) {
	var name string
	if name = ctx.Args().First(); name == "" {
		name = "default"
	}
	verbose := ctx.Bool("verbose")
	cfg, err := config.ReadSettings(ctx.String("config"))
	if err != nil {
		return err
	} else if verbose {
		pp.Println(cfg)
	}
	source, dbconf := cfg.GetSource(name)
	for _, target := range cfg.ReverseTargets {
		if target.Type == "codes" && target.Language == "" {
			target.Language = "golang"
		}
		if target.TablePrefix == "" {
			target.TablePrefix = dbconf.TablePrefix
		}
		if err = config.RunReverse(&source, &target); err != nil {
			return err
		}
	}
	return nil
}
