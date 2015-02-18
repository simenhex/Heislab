package main

import (
	"net"
	"fmt"
	"time"
	"encoding/json"
	"runtime"
)

type State struct{
	Counter 	int
	Master 		bool


func backUp(){

	var buff[1024]byte
	var local_IP string = "129.241.187.144"

 
	addr,err :=net.ResolveUDPAddr( "udp",":20011")	//returnerer (addr,err)
	if err != nil{
		fmt.Println("ResolveUDPAddr failed")
		return
	}
	//fmt.Println("\n addr.IP:",addr.IP)	


	sock,err := net.ListenUDP("udp", addr)
	if err != nil{
		fmt.Println("ListenUDP failed")
		return
	}


	sock.SetReadDeadline(time.Now().Add(1500*time.Millisecond))
	/*time_now=timenow()
	lastState int
	currentState int*/
	 
	for {
		rlen,_,err := sock.ReadFromUDP(buff[:])
		sock.SetReadDeadline(time.Now().Add(1500*time.Millisecond))
		if err != nil{
			fmt.Println(err,"\n")
			//opprett ny State struct
			//send state struct til main med Counter=currentState og Master=true gjennom channel
			return
		}
		if err==nil {
			fmt.Println("rlen og local IP:",rlen,local_IP)
			var k int
			json.Unmarshal(buff[:rlen],&k)
			fmt.Println(k,"\n")
			/*currentState=k
			if(currentState!=lastState)
				resetTimer
			}
			lastState=currentState*/

		}
	}
}


func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())



	}	
	state := State{
		Counter:	0,
		Master:		true,
	}

	go backUp()
	time.Sleep(2*time.Second)

	for{
		if(state.Master==true){
			//close backUp thread
			server_addr,err:=net.ResolveUDPAddr("udp","129.241.187.255:20011")
			if err != nil{
				fmt.Println("ResolveUDPAddr failed")
			}

			sendSock,err:=net.DialUDP("udp",nil,server_addr)
			if err != nil{
				fmt.Println("DialUDP failed")
			}

			for{
				fmt.Println("server_addr:",server_addr,"\n")
				b,err := json.Marshal(state.Counter)

				tull,err := sendSock.Write([]byte(b))
				if err != nil{
					fmt.Println("WriteToUDP failed")
				}	
				time.Sleep(1*time.Second)
				state.Counter+=1

				fmt.Println(tull,"is poop")
			}
		}
	}
}

