/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-02 17:03:44
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */

package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	// "sync"
)

type Scan struct {
	IP      string
	Port    string
	Process int
	Debug bool
}

func (s *Scan) Parse() error {
	var pageCount int
	s.Debug = true
	if s.Port == "" {
		s.Port = "21,22,80,8080,3306,9000"
	}
	if s.Process == 0 {
		s.Process = 1
	}
	ports, _ := s.getAllPort(s.Port)
	total := len(ports)

	if total < s.Process {
		pageCount = total
	} else {
		pageCount = s.Process
	}
	
	num := int(math.Ceil(float64(total) / float64(pageCount)))
	
	s.sendLog(fmt.Sprintf("%v 【%v】需要扫描端口总数:%v 个，总协程:%v 个，每个协程处理:%v 个", time.Now().Format("2006-01-02 15:04:05"), s.IP, total, pageCount, num))
	
	//start := time.Now()
	all := map[int][]int{}
	for i := 1; i <= pageCount; i++ {
		for j := 0; j < num; j++ {
			tmp := (i-1)*num + j
			if tmp < total {
				all[i] = append(all[i], ports[tmp])
			}
		}
	}




	return nil
}

func (s *Scan) getAllPort(port string) ([]int, error) {
	var ports []int
	//处理 ","号 如 80,81,88 或 80,88-100
	portArr := strings.Split(strings.Trim(port, ","), ",")
	for _, v := range portArr {
		portArr2 := strings.Split(strings.Trim(v, "-"), "-")
		startPort, err := s.filterPort(portArr2[0])
		if err != nil {
			continue
		}
		//第一个端口先添加
		ports = append(ports, startPort)
		if len(portArr2) > 1 {
			//添加第一个后面的所有端口
			endPort, _ := s.filterPort(portArr2[1])
			if endPort > startPort {
				for i := 1; i <= endPort-startPort; i++ {
					ports = append(ports, startPort+i)
				}
			}
		}
	}
	//去重复
	ports = s.arrayUnique(ports)

	return ports, nil
}

func (s *Scan) sendLog(str string) {
	if s.Debug == true {
		fmt.Println(str)
	}
}

func (s *Scan) filterPort(str string) (int, error) {
	port, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	if port < 1 || port > 65535 {
		return 0, errors.New("端口号范围超出")
	}
	return port, nil
}

func (s *Scan) arrayUnique(arr []int) []int {
	var newArr []int
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}
