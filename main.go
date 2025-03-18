package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

// deux gardes => deux go routines

func main() {
	var mutex sync.Mutex
	message := "Message périodique"
	inputChan := make(chan string)

	// Lecture asynchrone de l'entrée standard
	go readInput(inputChan)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			mutex.Lock()
			fmt.Println(message) // IMPORTANT : affiché sur stdout pour être récupéré par un autre programme
			mutex.Unlock()

		case msg := <-inputChan:
			mutex.Lock()
			message = msg
			fmt.Println(message)                                       // Envoie sur stdout pour que l'autre programme puisse le recevoir
			fmt.Fprintln(os.Stderr, "Nouveau message reçu :", message) // Indication sur stderr
			mutex.Unlock()
		}
	}
}

// Fonction pour lire l'entrée standard de manière asynchrone
func readInput(ch chan<- string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() { //tant qu'il scanne
		ch <- scanner.Text() // Envoie la ligne lue au canal
	}
}
