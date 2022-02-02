package main

/*
#cgo linux CFLAGS: -fplugin=./golang.so

void echo() {
  printf("link: https://github.com/makalek/gogo");
}

*/
import "C"

func main() {
	C.echo()
	return
}