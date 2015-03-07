package main

import(
	"fmt"
	 ."driver"
	 "timer"
	 "time"
)

//var Number_of_floors = 4 //for å enkelt endre antall etasjer. Må også endres i driver.

func main() {

	timerChan := make(chan int)
	initChan := make(chan bool)	//for å kalle initialiseringsrutiner i elev_controller og elev_brain
	nextFloorChan := make(chan int)	//brukes av elev_brain til å kommunisere med/til elev_controller
	currentFloorChan := make(chan int)	//bør kanskje være global variable, da nettverk, brain og controller trenger current_floor
	directionChan := make(chan int)	//må finne ut om controller trenger å endre på dir, eller om bare brain trenger denne informasjone




	







	// i := D_init()
	// fmt.Println(i)
	// D_set_motor_direction(0)


	// last_floor:=0
	// D_set_motor_direction(1)
	// fmt.Println("Motor er satt opp")
	// for{
	// 	if(D_get_floor_sensor_signal() == 2 && last_floor != 2){
	// 		D_set_motor_direction(-1)
	// 		fmt.Println("Motor er satt ned")
	// 		D_set_floor_indicator(2)
	// 		last_floor = 2
	// 		}
	// 	if(D_get_floor_sensor_signal() == 0 && last_floor != 0){
	// 		D_set_motor_direction(1)
	// 		fmt.Println("Motor er satt opp")
	// 		D_set_floor_indicator(0)
	// 		last_floor = 0
	// 		}
	
	// }

}
	

