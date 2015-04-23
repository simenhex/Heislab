package networkModule

import(
	"fmt"
	"net"
	"encoding/json"
	"sleep"

)

runtime.GOMAXPROCS(runtime.NumCPU())

//Finn ut av: 
	//Hvordan oppdager en heis at den er koblet fra nettverket? 



type lastStateUpdate struct {
    last_floor  int 
    direction	int //-1,0 or 1
} 

type outsideOrder struct{	//opprett samme strukttype som i elev_brain
	test_var int
}

var last_state_update   lastStateUpdate
var outside_order 		outsideOrder


 var outsideOrderSocketSend 	*netUDPConn
 var outsideOrderSocketReceive 	*netUDPConn
 var stateSocketSend 			*netUDPConn
 var stateSocketReceive	 		*netUDPConn



func initSockets(ta inn de IP adressene som trengs){
	//Create sender socket for OutsideOrder
	broadcast_udp_addr,err:=net.ResolveUDPAddr("udp","129.241.187.255:7100")
	if err != nil{
		fmt.Println("ResolveUDPAddr failed in outisderOrder sending")
		return
	}


	outsideOrderSocketSend,err:=net.DialUDP("udp",nil,broadcast_udp_addr)
	if err != nil{
		fmt.Println("DialUDP failed in outsiderOrder sending")
		return
	}
	outsideOrderSocketSend = outsideOrderSocketSend

	//Create sender socket for PING
	broadcast_udp_addr,err:=net.ResolveUDPAddr("udp","129.241.187.255:8100")
	if err != nil{
		fmt.Println("ResolveUDPAddr failed in ping sending")
		return
	}

	stateSocketSend,err:=net.DialUDP("udp",nil,broadcast_udp_addr)
	if err != nil{
		fmt.Println("DialUDP failed in PING sending")
		return
	}
	stateSocketSend = stateSocketSend


	//Create receiver socket for outsideOrder
	listen_addr,err :=net.ResolveUDPAddr( "udp",":7100")
	if err != nil{
		fmt.Println("ResolveUDPAddr failed in otuside receiving")
		return
	}

	outsideOrderSocketReceive,err := net.ListenUDP("udp", listen_addr)
	if err != nil{
		fmt.Println("ListenUDP failed for ping")
		return
	}
	outsideOrderSocketReceive = outsideOrderSocketReceive

	//Create receiver socket for PING
	listen_addr,err :=net.ResolveUDPAddr( "udp",":8100")
	if err != nil{
		fmt.Println("ResolveUDPAddr failed in PING receiving")
		return
	}
	
	stateSocketReceive,err := net.ListenUDP("udp", listen_addr)
	if err != nil{
		fmt.Println("ListenUDP failed for ping")
		return
	}
	stateSocketReceive = stateSocketReceive
}




func NetworkModule(brainToNetworkOrder chan outsideOrder, networkToBrainOrder chan outsideOrder, brainToNetworkState chan lastStateUpdate, networkToBrainState lastStateUpdate){

	initSockets(Ta inn de IP adressene som trengs)

	go listenForOutsideOrder(networkToBrainOrder)
	go listenForNewStateUpdate(networkToBrainState)
	go broadcastLastStateUpdate() //ping


	for {
		select{
			case read_from_chan:= <- brainToNetworkOrder:	
				broadcastOutsideOrder(read_from_chan)

			case read_from_chan:= <- brainToNetworkState:	
				last_state_update = read_from_chan

			default:
				time.Sleep(10*time.Milliseconds)	//Viktig at kanalene leses raskere enn det sendes inn på dem
		}
	}
}


/////////////////Subfunctions////////////////////////////////////////////////////////


//TEST: Unmarshal fungerer
func listenForOutsideOrder(networkToBrainOrder chan outsideOrder){
	byte_array := make([] byte, 1024)
	var outside_order outsideOrder

	for {
		rec_len,received_ip,err := outsideOrderSocketReceive.ReadFromUDP(byte_array)
		if err != nil{
			fmt.Println("Not able to receive from ",received_ip)
		}
		err := json.Unmarshal(byte_array[0:rec_len], &outside_order)
		if(err != nil) {fmt.Println("Error with Unmarshal!!")}
		networkToBrainOrder <- outside_order 	
		byte_array = clearByteArray(byte_array,rec_len)	
		time.Sleep(10*time.Milliseconds)


	}
}

//TEST: Unmarshal fungerer
func listenForNewStateUpdate(networkToBrainState chan lastStateUpdate){
	byte_array := make([] byte, 1024)
	var last_state_update lastStateUpdate

	for {
		rec_len,received_ip,err := stateSocketReceive.ReadFromUDP(byte_array)
		if err != nil{
			fmt.Println("Not able to receive from ",received_ip)
		}
		err := json.Unmarshal(byte_array[0:rec_len], &last_state_update)
		if(err != nil) {fmt.Println("Error with Unmarshal i listenForNewStateUpdate()!!!")}
		networkToBrainState <- last_state_update 	

		byte_array = clearByteArray(byte_array,rec_len)	
		time.Sleep(10*time.Milliseconds)
	}
}

//TEST OK
func broadcastLastStateUpdate(){
	number_of_broadcast := 2
	for{
		byte_array,err := json.Marshal(last_state_update)
		for i:=0; i < number_of_broadcast; i++{
			_,err := stateSocketSend.Write(byte_array)
			if err != nil {fmt.Println("Error with Marshal i broadcastLastStateUpdate()!!")}
		}
		time.Sleep(100*time.Milliseconds)	
	}	
}

//TEST OK
func broadcastOutsideOrder(outside_order outsideOrder){
	number_of_broadcast := 4
	byte_array, err := json.Marshal(outside_order)
	for i:=0; i < number_of_broadcast; i++{
		_,err := outsideOrderSocketSend.Write(byte_array)
		if err != nil {fmt.Println("Error with Marshal i broadcastOutsideOrder()!!")}
	}
}


//Test OK
func clearByteArray(byte_array []byte, len int) []byte {
	var clear uint8
	clear = 0
	for i:=0;i < len; i++{
		byte_array[i] = clear
	}
	return byte_array	
}

//Returner lokal IP adresse. Endre navn mtp navnekonvensjon og rydd litt opp ( hvis den trengs i det hele tatt)
func getLocalIp(port_number int) net.IP {	
	//Generating broadcast address
	localListenPort 	:= 7600
	baddr, _ := net.ResolveUDPAddr("udp4", "255.255.255.255:"+strconv.Itoa(port_number))

	//Generating localaddress
	tempConn, _ := net.DialUDP("udp4", nil, baddr)
	defer tempConn.Close()
	tempAddr := tempConn.LocalAddr()
	laddr, _ := net.ResolveUDPAddr("udp4", tempAddr.String())
	laddr.Port = localListenPort
	return laddr.IP
}