package main

import "strings"

const IcmpLength = 8


func main() {

	cmd := NewCmd()
	//fmt.Println(cmd)
	if len(strings.Split(cmd.Ip,".")) != 3 {
		cmd.Ip = "192.168.0"
	}
	test := NewQP(cmd.Ip)
	for i:=0;i<256;i++{
		go Ping(test.Prefix,i,1)
	}
	cnt := 0
	for i := range AllResuiltChannel{
		test.Print[i.num] = i.success
		cnt++
		if cnt == 256 {
			break
		}
	}
	Format(test,20,false)
}
