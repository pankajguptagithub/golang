package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct{ sync.Mutex }
type Philo struct {
	leftCS  *ChopS
	rightCS *ChopS
}

var host_permission = make(chan int, 2)

func hostControl(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	host_permission <- 1
	<-host_permission
}

func (p Philo) eat(wg *sync.WaitGroup, i int) {

	for i := 0; i < 3; i++ { // Each philosopher can eat max 3 times
		fmt.Printf("starting to eat %d\n", i)

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("Philosopher %d eating\n", i)
		time.Sleep(2000 * time.Millisecond)

		p.leftCS.Unlock()
		p.rightCS.Unlock()
		fmt.Printf("finishing eating %d\n", i)

	}
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	fmt.Println("Dinning Phillosophers Conference - Hungry maniacs")
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {

		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]}
	}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(&wg, i)
	}
	wg.Wait()
}
