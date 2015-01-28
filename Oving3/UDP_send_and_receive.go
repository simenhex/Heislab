package main

import (
	"net"
	"fmt"
	"time"
)



func main(){
	
	send()
	receive()

}


func send(){
	//broadcastIP := "129.241.187.255"
	//port:= "20012"
	
	server_addr,err:=net.ResolveUDPAddr("udp","129.241.187.255:20012")
	if err != nil{
		fmt.Println("ResolveUDPAddr failed badly")
	}

	sendSock,err:=net.DialUDP("udp",nil,server_addr)
	if err != nil{
		fmt.Println("DialUDP failed badly")
	}

	message:="Å hei hvor det går"

	fmt.Println(server_addr,"\n")
	
	n,err:= sendSock.Write([]byte(message))
	if err != nil{
		fmt.Println("WriteToUDP failed badly")
	}


	fmt.Println("ukjente n er:",n,"\n")

}




func receive(){

	var buff[1024]byte
	//defer sock.close()	


	addr,err :=net.ResolveUDPAddr( "udp",":20012")
	if err != nil{
		fmt.Println("ResolveUDPAddr failed badly")
		return
	}
	//fmt.Println("addr.IP:",addr.IP)
	


	sock,err := net.ListenUDP("udp", addr)
	if err != nil{
		fmt.Println("ListenUDP failed badly")
		return
	}
	//fmt.Println(sock)
	 
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
