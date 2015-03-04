package elev_handler
import "time"


func DoorTimer(timerChan chan int)  {

	select {
		case <- time.After(3 * time.Second):
			timerChan <- 1

	}
}


