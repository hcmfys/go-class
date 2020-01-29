package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	f, _ := os.Open("/Volumes/e/web/go-class/classReader/ldc/idc.txt")
	r := bufio.NewReader(f)
	for {
		b, _, e := r.ReadLine()
		if e != nil {
			break
		}
		line := string(b)
		arr:=strings.Split(line," ")
		//fmt.Println(arr)
		fmt.Printf("ldcMap[%s] = Ldc_Code{CodeId:%s,CodeTxt:\"%s\",CodeBytes:1,Desc:\"%s\"}\n",arr[0],arr[0],arr[1],arr[2])

	}
}
