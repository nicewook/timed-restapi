package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var ready bool

func startHandler(w http.ResponseWriter, r *http.Request) {
	ready = true
	// tick := time.NewTicker(1 * time.Second)
	go func() {
	Loop:
		for {
			select {
			// case t := <-tick.C:
			case <-time.After(10 * time.Second):
				// fmt.Println("tick!", t)
				fmt.Println("tick!")
				if !ready {
					break Loop
				}
			}
		}
		fmt.Println("finish goroutine")

	}()
}
func shootHandler(w http.ResponseWriter, r *http.Request) {
	if !ready {
		fmt.Println("not ready")
		return
	}
	ready = false
	fmt.Println("ready! shoot!")
}

func main() {
	http.HandleFunc("/start", startHandler)
	http.HandleFunc("/shoot", shootHandler)
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatalf("err: %v", err)
	}
}
