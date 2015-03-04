package main

import(
 	//"fmt"
 	."driver"
 	"time"
)



func main(){
	D_init()
	D_set_motor_direction(-1)


	for{

		if D_get_button_signal(BUTTON_COMMAND,0) == 1{
			D_set_button_lamp(BUTTON_COMMAND,0,1)
		}
		if D_get_button_signal(BUTTON_COMMAND,1) == 1{
			D_set_button_lamp(BUTTON_COMMAND,1,1)
		}
		if D_get_button_signal(BUTTON_COMMAND,2) == 1{
			D_set_button_lamp(BUTTON_COMMAND,2,1)
		}
		if D_get_button_signal(BUTTON_COMMAND,3) == 1{
			D_set_button_lamp(BUTTON_COMMAND,3,1)
		}

		time.Sleep(50*time.Millisecond)
	}
	

}


