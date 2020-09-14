package main

import "ct/say"
import "ct/say/chinese"
import "ct/say/iso"
import "fmt"
import "log"
import "errors"

func PrintTest() {
	fmt.Println("fmt.Println in main package")
	log.Println("log.Println in main package")
}

func greeting(who string) (string, error) {
	if who == "" {
		return "", errors.New("missing who")
	}
	message := fmt.Sprintf("Hi %v", who)
	return message, nil
}

func main() {
	ch.Hello()
	chinese.Hi()
	iso.Hi()
	PrintTest()
	ErrorTest()

	log.Println("call Exception Handler")
	err:=ExceptionHandler()
	log.Println(err)

	log.Println("call Exception Test")
	ExceptionTest()
}

func ErrorTest() {
	message, error := greeting("smite")
	if error != nil {
		fmt.Println("greetng error")
	} else {
		fmt.Println(message)
	}
}

func ExceptionTest() {
	message, error := greeting("")
	if error != nil {
		panic("greetng error")
	} else {
		fmt.Println(message)
	}
	log.Println("End Exception Test")
}

func ExceptionHandler() (err error) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("panic recover! p: %v\n", p)
			err=errors.New("after handler error")
		}
	}()
	ExceptionTest()
	log.Println("End Exception Handler")
	return nil
}
