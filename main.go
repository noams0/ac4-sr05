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

	wg.Add(2)

	// Goroutine pour lire l'entrée standard (asynchrone)
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			mutex.Lock()
			// Décommenter pour tester l'atomicité :
			// fmt.Fprintln(os.Stderr, ".")
			// time.Sleep(10 * time.Second)

			message = scanner.Text()
			fmt.Fprintln(os.Stderr, "nouveau message reçu:", message)
			mutex.Unlock()
		}
	}()

	// Goroutine pour afficher périodiquement le message
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			mutex.Lock()
			// Décommenter pour tester l'atomicité :
			// fmt.Fprintln(os.Stderr, ".")
			// time.Sleep(10 * time.Second)
			fmt.Println(message)
			mutex.Unlock()
		}
	}()

	//Overkill mais le programme ne s'arrêtera que quand les deux go routines s'arrêteront
	wg.Wait()
}
