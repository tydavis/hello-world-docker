package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

func main() {
	hi := []string{"Hello", "Hola", "Bonjour",
		"Ciao", "こんにちは", "안녕하세요",
		"Cześć", "Olá", "Здравствуйте",
		"Chào bạn", "您好", "Hallo"}

	fmt.Printf("Saying hello in up to %d languages! \n", len(hi))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	tick := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-tick.C:
			r := rand.Intn(len(hi))
			fmt.Println(hi[r])
		case <-c:
			// Got a Ctrl-C
			fmt.Println("Exiting!")
			os.Exit(0)
		default:
			time.Sleep(300 * time.Millisecond)
		}
	}
}
