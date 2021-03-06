package driver


/*
#cgo LDFLAGS: -lcomedi -lm
#include "io.h"
*/
import "C"


func Io_init() bool {
	p := int(C.io_init())
	return bool(p != 0)
}

func Io_set_bit(channel int) {
	C.io_set_bit(C.int(channel))
}

func Io_clear_bit(channel int) {
	C.io_clear_bit(C.int(channel))
}

func Io_write_analog(channel int, value int){
	C.io_write_analog(C.int(channel), C.int(value))
}

func Io_read_bit(channel int) int {
	return int(C.io_read_bit(C.int(channel)))
}

func Io_read_analog(channel int) int {
	return int(C.io_read_analog(C.int(channel)))
}

