package elev_controller

import(
	"time"
	"fmt"
	."driver"
)

var number_of_floors = 4
var current_floor int
var next_floor int
var direction int 
var door_open = false
var init_status  = false
var reset_timer = false
// trenger kanskje en inbetween_floors tilstandsvariabel


func elevControllerMain(initChan chan bool, nextFloorChan chan int, currentFloorChan chan int, timerChan chan int, dirChan chan int ){
	for{
		select{
			case initialize := <- initChan:
				init_status = initialize
				if init_status == true{
					direction = -1					//bør randomizes
					dirChan <- direction 			//-1 = retning DOWN. Lage enum?
					D_set_motor_direction(-1)
				}
			
			case door_timer_is_done := <- timerChan:
				if door_timer_is_done == 1{
					timerIsDone()
				}

			case floor_reached := <- currentFloorChan:
				current_floor = floor_reached
				floorReached(timerChan)
				dirChan <- direction

			case go_to_next_floor := <- nextFloorChan:
				next_floor = go_to_next_floor
				goToFloor(timerChan)
				dirChan <- direction
			
			default:
				time.Sleep(30*time.Millisecond)	//sover i 0.03 sekunder for å frigjøre CPU
		}
	}
}

func goToFloor(timerChan chan int){
	fmt.Println("Reached function: goToFloor")
	if next_floor == -1{
		direction = 0
		D_set_motor_direction(direction)
	
	}else if next_floor == current_floor{
		floorReached(timerChan)
		// må muligens ha funksjon eller channel som sier i fra til brain at en bestilling er fullført
	}else if !door_open{
		if next_floor > current_floor{
			direction = 1
			D_set_motor_direction(direction)

		}else if next_floor < current_floor{
			direction = -1
			D_set_motor_direction(direction)
		}
	}
}

func floorReached(timerChan chan int){
	//håndter at man har nådd en etasje. Sjekk om etasjen er den man skal nå, hvis ja ->  stop og åpne døren etc.
	fmt.Println("Reached function: floorReached")
	D_set_floor_indicator(current_floor)
	if init_status{
		direction = 0
		D_set_motor_direction(direction)
	}
	if current_floor == next_floor{
		fmt.Println("Stopping at floor #",current_floor)
		D_set_motor_direction(0)
		door_open = true
		timerChan <- -1					// resetter timer. Må kanskje ha en egen resetTimerChan
		D_set_door_lamp(1)
		if current_floor == number_of_floors{
			direction = -1				//avhengig av elev_brain er det ikke sikkert det er nødvendig å endre retning
		}else if current_floor == 1{
			direction = 1				//avhengig av elev_brain er det ikke sikkert det er nødvendig å endre retning
		}
	}


}

func timerIsDone(){
	if door_open{
		door_open = false
		D_set_door_lamp(-1)
	}
}
