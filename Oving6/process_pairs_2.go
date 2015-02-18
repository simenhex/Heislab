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
}

func backUp(){

	var buff[1024]byte
	var local_IP string = "129.241.187.144"

 
	addr,err :=net.ResolveUDPAddr( "udp",":20010")	
	if err != nil{
		fmt.Println("ResolveUDPAddr failed")
		return
	}

	sock,err := net.ListenUDP("udp", addr)
	if err != nil{
		fmt.Println("ListenUDP failed")
		return
	}


	sock.SetReadDeadline(time.Now().Add(1500*time.Millisecond))

	//ny
	stateTemp:=State{
		Counter:	0,
		Master: 	1,
	}

	channel:=make(chan State )
	 
	for {
		rlen,_,err := sock.ReadFromUDP(buff[:])
		sock.SetReadDeadline(time.Now().Add(1500*time.Millisecond))
		//Timeout?
		if err != nil{
			fmt.Println(err,"\n")
			stateTemp.Master:=true
			channel <- stateTemp

			//send state struct til main med Counter=currentState og Master=true gjennom channel
			return
		}

		if err==nil {	
			fmt.Println("rlen og local IP:",rlen,local_IP)
			var k int
			json.Unmarshal(buff[:rlen],&k)
			fmt.Println(k,"\n")
			stateTemp.Counter:=k

			/*currentState=true
			if(currentState!=lastState)
				resetTimer ??	
			}
			lastState=currentState*/


		}
	}
}


func main(){
	fmt.Println("\n\n\n") //bedre oversikt
	runtime.GOMAXPROCS(runtime.NumCPU())

	state := State{
		Counter:	0,
		Master:		false,	//ny
	}

	go backUp()
	time.Sleep(2*time.Second)

	for{
		if(state.Master==true){
			//close backUp thread ?? blir vel gjort i backup ved bruk av return?
			server_addr,err:=net.ResolveUDPAddr("udp","129.241.187.255:20010")
			if err != nil{
				fmt.Println("ResolveUDPAddr failed")
				return
			}

			sendSock,err:=net.DialUDP("udp",nil,server_addr)
			if err != nil{
				fmt.Println("DialUDP failed")
				return
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
		if state.Master == false{	//ny
			//oppdaterer structen
			state <- channel	//er channel tilgengelig her  når den blir deklarert i en annen terminal?


		}

		}
	}
}

