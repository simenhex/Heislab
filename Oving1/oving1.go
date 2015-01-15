package main


import(
	"fmt"
	"runtime"
	"time"
)

var i int=0

func thread0(){
	for counter:=0; counter<1000000; counter++{
		i++
	}

}

func thread1(){
	for counter:=0; counter<1000000; counter++{
		i--
	}

}

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	go thread0()
	go thread1()

	time.Sleep(100*time.Millisecond)
	fmt.Println(i)
}
