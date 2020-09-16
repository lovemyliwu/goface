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
	log.Println(ReturnVarTest())
	ch.Hello()
	chinese.Hi()
	iso.Hi()
	PrintTest()
	ErrorTest()

	log.Println("call Exception Handler")
	msg, err := ExceptionHandler()
	log.Println(msg)
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

func ExceptionHandler() (msg string, err error) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("panic recover! p: %v\n", p)
			msg = "hi from recover"
			err = errors.New("after handler error")
		}
	}()
	ExceptionTest()
	log.Println("End Exception Handler")
	return "no error", nil
}

func ReturnVarTest() (info string) {
	info = "hi"
	//msg = "yes"
	return info
}
