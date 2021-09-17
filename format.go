package main

import (
	"fmt"
	"strconv"
	"time"
)

func formatValue(n int)(str string,err error){
	if n < 0 || n > 255{
		return  "",fmt.Errorf("%d out of range,should be 0 <= n <= 255",n)
	}
	strN := strconv.Itoa(n)
	l := len(strN)
	if l == 1 {
		str = "  " + strN + "  "
	}else if l == 2 {
		str = " " + strN + "  "
	}else{
		str = " " + strN + " "
	}
	return
}



func Format(qp QuickPing,row int,fileModel bool){
	s := "+-----"
	t := row
	lastCnt := 0
	used := 0
	fmt.Printf("########## %s.0---%s.255 ##########\n",qp.Prefix,qp.Prefix)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//(256+t-1)/t 向上取整,保证在一行有r列的情况下能装下256个ip
	for i:=0;i<(256+t-1)/t;i++ {
		for j:=0;j<20;j++{
			fmt.Print(s)
		}
		fmt.Println("+")
		m := t * i
		lc := 0
		for j:=0;j<20;j++{
			num := m + j
			if num > 255 {
				lastCnt = lc
				break
			}
			str := "     "
			var err error

			if qp.Print[num] {
				str,err = formatValue(num)
				used++
			}
			if err != nil {
				lastCnt = lc
				break
			}
			fmt.Print("|")
			fmt.Print(str)
			lc++
		}
		fmt.Println("|")
	}
	for lastCnt != 0 {
		fmt.Print(s)
		lastCnt--
	}
	fmt.Println("+")
	fmt.Printf("used %d  !!!\n",used)
}
