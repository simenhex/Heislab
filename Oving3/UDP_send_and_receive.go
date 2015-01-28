package main

import (
	"net"
	"fmt"
	"runtime"
	"time"
)





func send(){
	
	server_addr,err:=net.ResolveUDPAddr("udp","129.241.187.255:20012")
	if err != nil{
		fmt.Println("ResolveUDPAddr failed")
	}

	sendSock,err:=net.DialUDP("udp",nil,server_addr)
	if err != nil{
		fmt.Println("DialUDP failed")
	}
	
	
	message:="Å hei hvor det går"

	fmt.Println("server_addr:",server_addr,"\n")
	
	for{
		n,err := sendSock.Write([]byte(message))
		if err != nil{
			fmt.Println("WriteToUDP failed")
		}
		time.Sleep(2*time.Second)
		if n == 2{
		}
	}

}




func receive(){
	var buff[1024]byte
	var local_IP string = "129.241.187.144"
	//defer sock.close()	


	addr,err :=net.ResolveUDPAddr( "udp",":20012")
	if err != nil{
		fmt.Println("ResolveUDPAddr failed")
		return
	}
	fmt.Println("\n addr.IP:",addr.IP)
	


	sock,err := net.ListenUDP("udp", addr)
	if err != nil{
		fmt.Println("ListenUDP failed")
		return
	}

	 
	for {
		rlen,remote,err := sock.ReadFromUDP(buff[:])
		if err != nil{
			fmt.Println("Not able to receive from",remote)
		}

		//for at sammenligningen skal fungere må local_IP være av samme type som remote.IP, dvs *AdrUDP eller noe  
		if string(remote.IP) != local_IP{
			fmt.Println("Received ", rlen, " bytes from ", remote)
			fmt.Println("Remote port:", remote.Port,"\n")
			fmt.Println(string(buff[:]),"\n")

		}
	}
}

func main(){
	

	fmt.Println("\n\n")	
 	runtime.GOMAXPROCS(runtime.NumCPU())
	deadChan :=make(chan int)

	go receive()
	go send()
	<-deadChan
}



