package main

import(
	"fmt"
	"net"
	"log"
	"time"
	"encoding/json"
)
/*
func spawnBackup(count int){
	command := exec.Command("gnome-terminal", "-e", "go", "run", "processpairs.go", string(count))
	err := command.Start()
	if err != nil {
		log.Println(err)
	}
}
*/
func sendCount(count int, outSock *net.UDPConn, broadcastAdr *net.UDPAddr) {
	//Send current
	b, err := json.Marshal(count)
	_,err = outSock.Write([]byte(b))
	if err != nil {
		log.Println(err)
	}
}

func reciveMessage(count int, listenSocket *net.UDPConn) (int, error) {
	buffer := make([]byte, 1024)
	lenght,_,err := listenSocket.ReadFromUDP(buffer)
	if err != nil {
		log.Println(err)
	}else{
		var tempCount int = -1
		err2 := json.Unmarshal(buffer[:lenght], &tempCount)
		if err2 != nil {
			log.Println(err2)
		}else{
			fmt.Println("Recived count = ", tempCount)
		}
		return tempCount, err
	}
	return count, err
}

func main(){
	var(
		master bool = false
		count int = -1
	)
	
	switch master{
		case false:
		//Create UDP socket (listen)
		localAdr, err := net.ResolveUDPAddr("udp", ":6789")
		if err != nil {
			log.Println(err)
		}
		listenSocket, err := net.ListenUDP("udp", localAdr)
		if err != nil {
			log.Println(err)
		}
		listenSocket.SetDeadline(time.Now().Add(3*time.Second)) 
		
		//Listen on messages from master (while alive)
		for err == nil{
			count, err = reciveMessage(count, listenSocket)
			listenSocket.SetDeadline(time.Now().Add(1200*time.Millisecond)) 
		}

		//If this code run, the master is dead 
		
		//Close listenSocket
		listenSocket.Close()
		
		master = true
		fmt.Println("I am the master")
		fallthrough
		
		case true:
		//spawn backup
		/*spawnBackup(count)*/
		
		//create udp socket
		broadcastAdr, err:= net.ResolveUDPAddr("udp", "78.91.38.188:6789")
		if err != nil {
			log.Println(err)
		}
		outSocket, err := net.DialUDP("udp", nil, broadcastAdr)
		
		//send (while)
		for{
			count++
			fmt.Println(count)
			sendCount(count, outSocket, broadcastAdr)
			time.Sleep(500*time.Millisecond) 
		}
	}
}	