package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Question:
    Implement the dining philosopherâ€™s problem with the following constraints/modifications.

        There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

        Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)

        The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).

        In order to eat, a philosopher must get permission from a host which executes in its own goroutine.

        The host allows no more than 2 philosophers to eat concurrently.

        Each philosopher is numbered, 1 through 5.

        When a philosopher starts eating (after it has obtained necessary locks) it prints â€œstarting to eat <number>â€ on a line by itself, where <number> is the number of the philosopher.

        When a philosopher finishes eating (before it has released its locks) it prints â€œfinishing eating <number>â€ on a line by itself, where <number> is the number of the philosopher.

Reference:
    [Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines)
    [A pattern for limiting the number of goroutines in execution.](https://medium.com/@zufolo/a-pattern-for-limiting-the-number-of-goroutines-in-execution-56e13b226e72)
    [justforfunc #9: The Context Package](https://youtu.be/LSzR0VEraWw)
*/

const numberOfPhilosophers int = 5
const numberOfChopsticks int = numberOfPhilosophers
const numberOfTimesToEat int = 3
const numberOfJobsRunningConcurrently int = 2

func main() {

	var CSticks []*ChopS = make([]*ChopS, numberOfChopsticks)
	for idxCS := 0; idxCS < numberOfChopsticks; idxCS++ {
		CSticks[idxCS] = new(ChopS)
	}

	var philos []*Philo = make([]*Philo, numberOfPhilosophers)
	for idxP := 0; idxP < numberOfPhilosophers; idxP++ {
		philos[idxP] = &Philo{CSticks[idxP], CSticks[(idxP+1)%numberOfChopsticks], idxP + 1}
	}

	var chOkayToExitMain chan bool = make(chan bool)
	// Start the Host Controller so that the Philosophers can begin eating.
	go fxHostController(philos, chOkayToExitMain)

	// Exit the main function once fxHostController signals the end of processing.
	<-chOkayToExitMain
}

/*
 * Description:
 *		ChopS is a structure that has information about the Chopsticks
 */
type ChopS struct {
	sync.Mutex
}

/*
 * Description:
 *		Philo is a structure that has information about the Philosophers
 */
type Philo struct {
	leftCS, rightCS *ChopS
	number          int
}

/*
 * Description:
 *		Philo::eat obtains a lock for 2 chopsticks and eats 3 times.
 */
func (p *Philo) eat(goRouting int) {

	for eatCtr := 0; eatCtr < numberOfTimesToEat; eatCtr++ {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("\t starting to eat:", p.number, "\t round:", eatCtr+1, "\t goRouting:", goRouting)
		time.Sleep(100 * time.Millisecond)
		fmt.Println("\t finishing eating:", p.number, "\t round:", eatCtr+1, "\t goRouting:", goRouting)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
}

/*
 * Description:
 *		fxHostWorker informs the Philosopher to start eating.
 */
func fxHostWorker(startEatingChannel chan int, goRouting int, finishedEating chan bool, addressOfPhilos []*Philo) {
	for true {
		select {
		case philosopherIndex := <-startEatingChannel:
			fmt.Println("* Permission to eat has been given to Philosopher:", philosopherIndex+1, "\t[ goRouting:", goRouting, "]")
			(addressOfPhilos)[philosopherIndex].eat(goRouting)
			fmt.Println("* Philosopher", philosopherIndex+1, "has finished eating. \t[ goRouting:", goRouting, "]")
			finishedEating <- true
		}
	}
}

/*
 * Description:
 *		fxHostController allows no more than 2 philosophers to eat concurrently.
 */
func fxHostController(addressOfPhilos []*Philo, doExit chan bool) {

	var chStartEating chan int = make(chan int)
	var chFinishedEating chan bool = make(chan bool)

	// Start the Host Workers - they limit the concurrently running tasks.
	for i := 0; i < numberOfJobsRunningConcurrently; i++ {
		go fxHostWorker(chStartEating, i+1, chFinishedEating, addressOfPhilos)
	}

	// Ask the Philosophers to start eating.
	for j := 0; j < numberOfPhilosophers; j++ {
		go func(j int) {
			chStartEating <- j
		}(j)
	}

	// Wait for all the Philosophers to finish eating.
	for c := 0; c < numberOfPhilosophers; c++ {
		<-chFinishedEating
	}

	// Signal that all the Philosophers have finished eating.
	doExit <- true
}