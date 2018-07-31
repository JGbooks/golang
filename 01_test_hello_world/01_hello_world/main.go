package main
/*
#include <stdio.h>
void TestHello()
{
	printf("cgo :HELLO WORLD\r\n");
	return;
}
*/
import "C"
import (
	"log"
	"time"
)

func main() {
	//	C.TestHello()
	log.Println("windows 32 hello world!")
	time.Sleep(10 * time.Second)
}
