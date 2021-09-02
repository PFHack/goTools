/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-02 16:42:16
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var scantool Scan
var TimeOut = 5

func main() {
	app := &cli.App{
		Name:                   "PortScan",
		Version:                "Bt0.1",
		Usage:                  "",
		UsageText:              "PortS [参数选项] --ip参数为必须",
		EnableBashCompletion:   true,
		Copyright:              "版权所有：PFinal南丞  email：lampxiezi@163.com",
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "ip",
				Aliases:  []string{"i"},
				Usage:    "要扫描的目标IP",
				Required: true,
			},
			&cli.BoolFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Usage:   "要扫描的目标IP的端口号,默认 All",
			},
		},
		Action: func(c *cli.Context) error {
			scantool.IP = c.String("ip")
			if c.IsSet("port") {
				scantool.Port = c.String("port")
			}
			err := scantool.Parse()
			if err != nil {
				log.Fatal(err)
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
