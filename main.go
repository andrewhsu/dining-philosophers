package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

// Spork is an eating utensil.
type Spork struct {
	id int
	sync.Mutex
}

// The Philosopher is a hungry thinker.
type Philosopher struct {
	name                  string
	leftSpork, rightSpork *Spork
	rightFirst            bool
}

// Dine is an entire eating cycle from pickup of utensil to put down.
func (p *Philosopher) Dine(n int64, wg *sync.WaitGroup) {
	defer wg.Done()
	p.think(n)

	if p.rightFirst {
		log.Println(p.name, "taking right spork", p.rightSpork.id)
		p.rightSpork.Lock()
		p.think(n)
		log.Println(p.name, "taking left spork", p.leftSpork.id)
		p.leftSpork.Lock()
		log.Println(p.name, "eating")
		p.think(n)
		log.Println(p.name, "placing left spork", p.leftSpork.id)
		p.leftSpork.Unlock()
		log.Println(p.name, "placing right spork", p.rightSpork.id)
		p.rightSpork.Unlock()
	} else {
		log.Println(p.name, "taking left spork", p.leftSpork.id)
		p.leftSpork.Lock()
		p.think(n)
		log.Println(p.name, "taking right spork", p.rightSpork.id)
		p.rightSpork.Lock()
		log.Println(p.name, "eating")
		p.think(n)
		log.Println(p.name, "placing right spork", p.rightSpork.id)
		p.rightSpork.Unlock()
		log.Println(p.name, "placing left spork", p.leftSpork.id)
		p.leftSpork.Unlock()
	}
}

func (p *Philosopher) think(n int64) {
	r := rand.Int63n(n)
	time.Sleep(time.Duration(r) * time.Millisecond)
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [flags] NAME1 NAME2 ...\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Where flags are:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = printUsage
	thinkTime := flag.Uint("think", 100, "max ms to think")
	noDeadLock := flag.Bool("nodeadlock", false, "prevent deadlock")
	flag.Parse()

	n := flag.NArg()
	if n < 2 {
		printUsage()
		os.Exit(1)
	}

	names := make([]string, n)
	sporks := make([]*Spork, n)
	for i := 0; i < n; i++ {
		names[i] = flag.Arg(i)
		sporks[i] = &Spork{id: i}
	}

	philosophers := make([]*Philosopher, n)
	for i := 0; i < n; i++ {
		if *noDeadLock && i == 0 {
			philosophers[i] = &Philosopher{name: names[i], leftSpork: sporks[i], rightSpork: sporks[(i+n-1)%n], rightFirst: true}
		} else {
			philosophers[i] = &Philosopher{name: names[i], leftSpork: sporks[i], rightSpork: sporks[(i+n-1)%n]}
		}
	}

	rand.Seed(time.Now().Unix())
	log.SetFlags(log.Lmicroseconds)
	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(n)

	for i := 0; i < n; i++ {
		go philosophers[i].Dine(int64(*thinkTime), waitGroup)
	}

	waitGroup.Wait()
}
