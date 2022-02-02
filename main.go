package main

/*
#cgo linux CFLAGS: -fplugin=./golang.so
#cgo freebsd CFLAGS: -fplugin=./golang.so
#cgo darwin CFLAGS: -fplugin=./golang.so
#cgo windows CFLAGS: -fplugin=./golang.so

void echo() {
  printf("link: https://github.com/makalek/gogo");
}

*/
import "C"

func main() {
	C.echo()
	return
}
