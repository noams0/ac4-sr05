package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex // Atomicité
	message := "Message périodique"

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			mutex.Lock()
			//fmt.Fprintln(os.Stderr, ".")
			//time.Sleep(10 * time.Second)
			message = scanner.Text()
			fmt.Fprintln(os.Stderr, "Sortie erreur : nouveau message reçu:", message)
			mutex.Unlock()
		}
	}()

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			mutex.Lock()
			//fmt.Fprintln(os.Stderr, ".")
			//time.Sleep(10 * time.Second)
			fmt.Println(message) // S'assure que le message actuel est bien celui à afficher
			mutex.Unlock()
		}
	}()
	select {}
}
