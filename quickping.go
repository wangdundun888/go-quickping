package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

var AllResuiltChannel = make(chan Result,0)
type QuickPing struct {
	Prefix string
	Print  []bool
}

type Result struct {
	num     int
	success bool
}

func NewQP(prefix string)QuickPing {
	qp := QuickPing{
		Prefix: prefix,
		Print:  make([]bool, 256),
	}
	return qp
}

func Ping(ipPrefix string, num int, interval int) {
	var r Result
	msg := getHead()
	r.num = num
	remoteAddr, err := net.ResolveIPAddr("ip", ipPrefix+"."+strconv.Itoa(num))
	if err != nil {
		fmt.Println(ipPrefix+"."+strconv.Itoa(num), err)
		AllResuiltChannel <- r
		return
	}
	conn, err := net.DialIP("ip:icmp", nil, remoteAddr)
	if err != nil {
		fmt.Println("error: ", err)
		AllResuiltChannel <- r
		return
	}
	//ping
	if _, err := conn.Write(msg); err != nil {
		fmt.Println("send data error: ", err)
		AllResuiltChannel <- r
		return
	}
	timeout := func(tt chan<- int) {
		msg := make([]byte, 512)
		conn.Read(msg)
		tt <- 1
		return
	}
	tt := make(chan int,0)
	var t int
	go timeout(tt)
	select {
	case t = <- tt:
		_ = t
		r.success = true
		AllResuiltChannel <- r
		return
	case <-time.After(time.Duration(interval) * time.Second):
		AllResuiltChannel <- r
		return
	}
	AllResuiltChannel <- r
	return
}

func getHead()[]byte{
	msg := make([]byte,8)
	msg[0] = 8 // type
	msg[1] = 0 // code
	msg[2] = 0 // checkSum -> 2 byte
	msg[3] = 0
	msg[4] = 0  // identifier[0]
	msg[5] = 1 // identifier[1]
	msg[6] = 0  // sequence[0]
	msg[7] = 1 // sequence[1]
	// 检验和
	check := CheckSum(msg[:IcmpLength])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)
	return msg
}

func CheckSum(msg []byte) uint16 {
	sum := 0
	for n := 0; n < len(msg); n += 2 {
		sum += int(msg[n])<<8 + int(msg[n+1])
	}
	sum = (sum >> 16) + sum&0xffff
	sum += sum >> 16
	return uint16(^sum)
}
