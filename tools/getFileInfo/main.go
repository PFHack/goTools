/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-08-31 17:03:50
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */

package main

import (
	"fmt"
	"log"
	"os"
)

var (
	fileInfo os.FileInfo
	err      error
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stdout, "Using: %s file\n", os.Args[0])
		os.Exit(1)
	}
	fileInfo, err = os.Stat(os.Args[1])
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.",err)
		}
	}
	fmt.Println("|=======================================\n")
	fmt.Println("| File name:", fileInfo.Name())
	fmt.Println("|---------------------------------------\n")
	fmt.Println("| Size in bytes:", fileInfo.Size())
	fmt.Println("|---------------------------------------\n")
	fmt.Println("| Permissions:", fileInfo.Mode())
	fmt.Println("|---------------------------------------\n")
	fmt.Println("| Last modified:", fileInfo.ModTime())
	fmt.Println("|---------------------------------------\n")
	fmt.Println("| Is Directory: ", fileInfo.IsDir())
	fmt.Println("|---------------------------------------\n")
	fmt.Printf("| System interface type: %T\n", fileInfo.Sys())
	fmt.Println("|---------------------------------------\n")
	fmt.Printf("| System info: %+v\n\n", fileInfo.Sys())
	fmt.Println("|=======================================\n")
}
