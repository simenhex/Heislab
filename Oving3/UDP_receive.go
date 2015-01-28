package main

import (
	"net"
	"fmt"
	"time"
)




func main(){
	var buff[1024]byte

	defer sock.close()	


	addr,err :=net.ResolveUDPAddr( "udp",":30000")
	if err != nil{
		fmt.Println("ResolveUDPAddr failed badly")
		return;
	}
	fmt.Println("addr.IP:",addr.IP)
	


	sock,err := net.ListenUDP("udp", addr)
	if err != nil{
		fmt.Println("ListenUDP failed badly")
	}
	fmt.Println(sock)
	 
	for {
		rlen,remote,err := sock.ReadFromUDP(buff[:])
		if err != nil{
			fmt.Println("Not able to receive from",remote)
		}
		fmt.Println("Received ", rlen, " bytes from ", remote,"\n")
		fmt.Println(string(buff[:]))
		time.Sleep(2*time.Second)	
	
	
	}
	


	
	
	
}

