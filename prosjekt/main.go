package main

import(
	"fmt"
	 //."driver"
	 ."elev_handler"
	 "time"
)
func main() {

	timerChan := make(chan int)
	fmt.Println("Starter timer ved tidspunkt: ", time.Now(),"\n")
	go DoorTimer(timerChan)
	for{
		select{
			case DoorTimerIsDone := <- timerChan:
				if DoorTimerIsDone == 1{
					fmt.Println("Timer er ferdig ved tidspunkt: ", time.Now(),"\n")

				}
			default:	//Husk å ha en default i en select/case struktur
				time.Sleep(1*time.Second)	//Viktig å ha en sleep i en for/select struktur for å gi andre Go rutiner CPU aksess
		}
	}













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
	

