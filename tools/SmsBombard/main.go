/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-13 18:17:00
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var bom Bom

func main() {
	app := &cli.App{
		Name:                   "SmsBombard",
		Version:                "Bt0.1",
		Usage:                  "",
		UsageText:              "Phone [参数选项] -p参数为必须",
		EnableBashCompletion:   true,
		Copyright:              "版权所有：PFinal南丞  email：lampxiezi@163.com",
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "Phone",
				Aliases:  []string{"p"},
				Usage:    "要轰炸的电话",
				Required: true,
			},
			&cli.BoolFlag{
				Name:    "Num",
				Aliases: []string{"n"},
				Usage:   "要轰炸的次数",
			},
		},
		Action: func(c *cli.Context) error {
			bom.Phone = c.String("Phone")
			if c.IsSet("Num") {
				bom.Num = c.Int("Num")
			}
			err := bom.Parse()
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
