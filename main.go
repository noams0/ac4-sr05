package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	message := "Message périodique"
	var wg sync.WaitGroup
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	// Channel pour gérer la concurrence (overkill mais permet d'explorer les possibilités de Go dès mtn)
	syncChan := make(chan struct{}, 1)
	syncChan <- struct{}{}

	wg.Add(2)

	// goroutine => réception
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			<-syncChan // concurrence : attendre
			mutex.Lock()

			// décommenter => tester l'atomicité :
			fmt.Fprintln(os.Stderr, ".")
			time.Sleep(5 * time.Second)

			message = scanner.Text()
			fmt.Fprintln(os.Stderr, "nouveau message reçu:", message)

			mutex.Unlock()
			syncChan <- struct{}{} // concurrence : libérez
		}
	}()

	// Goroutine => émission
	go func() {
		defer wg.Done()
		for range ticker.C {
			<-syncChan // concurrence : attentdre
			mutex.Lock()

			// décommenter => tester l'atomicité :
			fmt.Fprintln(os.Stderr, "..")
			time.Sleep(5 * time.Second)
			fmt.Println(message)

			mutex.Unlock()
			syncChan <- struct{}{} // concurrence : libérez
		}
	}()

	wg.Wait()
}
