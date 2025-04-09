package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	message := "Message périodique"
	var wg sync.WaitGroup
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	// Channel pour gérer la concurrence (overkill mais permet d'explorer les possibilités de Go dès mtn)
	syncChan := make(chan struct{}, 1) //limité à 1 action
	syncChan <- struct{}{}

	wg.Add(2)

	// goroutine => réception
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() { //Scanner.Scan permet la lecture bloquante
			<-syncChan // concurrence : attendre

			message = scanner.Text()
			fmt.Fprintln(os.Stderr, "nouveau message reçu:", message)

			syncChan <- struct{}{} // concurrence : libérez
		}
	}()

	// Goroutine => émission
	go func() {
		defer wg.Done()
		for range ticker.C {
			<-syncChan // concurrence : attentdre

			fmt.Println(message)

			syncChan <- struct{}{} // concurrence : libérez
		}
	}()

	wg.Wait()
}
