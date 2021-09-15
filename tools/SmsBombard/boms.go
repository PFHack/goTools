/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-13 18:22:02
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	js "github.com/bitly/go-simplejson"
)

type Bom struct {
	Phone string
	Num   int
	Debug bool
}

func (b *Bom) GetApiUrls() (map[string]interface{}, error) {
	var result map[string]interface{}
	jsonFile, err := os.Open("api.json")
	if err != nil {
		return result, err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	result, err := js.NewJson([]byte(byteValue))
	if err != nil {
		return result, err
	}
	return result, nil
}

func (b *Bom) Parse() error {
	b.GetApiUrls()
	fmt.Println()
	return nil
}
