# dining-philosophers

[![Go Report Card](https://goreportcard.com/badge/github.com/andrewhsu/dining-philosophers)](https://goreportcard.com/report/github.com/andrewhsu/dining-philosophers)

:spaghetti: This is a simulation of [dining philosophers](https://en.wikipedia.org/wiki/Dining_philosophers_problem) using idiomatic go. The program will have all the philosophers at the table pickup their sporks and eat once before exiting.

Provided you have a [working go installation](https://golang.org/doc/install), you can download and install the package like so:

```
$ go get github.com/andrewhsu/dining-philosophers
```

You need at least two arguments for the philosophers' names:

```
$ dining-philosophers
Usage: dining-philosophers [flags] NAME1 NAME2 ...
Where flags are:
  -nodeadlock
      prevent deadlock
  -think uint
      max ms to think (default 100)
```

Running with the default think time will usually illustrate the deadlock:

```
$ dining-philosophers Aristotle Confucius
13:27:25.373425 Aristotle taking left spork 0
13:27:25.404421 Confucius taking left spork 1
13:27:25.459111 Aristotle taking right spork 1
13:27:25.482604 Confucius taking right spork 0
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0xc8200620dc)
  /usr/local/go/src/runtime/sema.go:47 +0x26
sync.(*WaitGroup).Wait(0xc8200620d0)
  /usr/local/go/src/sync/waitgroup.go:127 +0xb4
main.main()
  /Users/andrewhsu/work/src/github.com/andrewhsu/dining-philosophers/main.go:106 +0x595

goroutine 17 [semacquire]:
sync.runtime_Semacquire(0xc8200620bc)
  /usr/local/go/src/runtime/sema.go:47 +0x26
sync.(*Mutex).Lock(0xc8200620b8)
  /usr/local/go/src/sync/mutex.go:83 +0x1c4
main.(*Philosopher).Dine(0xc820076030, 0x64, 0xc8200620d0)
  /Users/andrewhsu/work/src/github.com/andrewhsu/dining-philosophers/main.go:48 +0xef6
created by main.main
  /Users/andrewhsu/work/src/github.com/andrewhsu/dining-philosophers/main.go:103 +0x575

goroutine 18 [semacquire]:
sync.runtime_Semacquire(0xc8200620ac)
  /usr/local/go/src/runtime/sema.go:47 +0x26
sync.(*Mutex).Lock(0xc8200620a8)
  /usr/local/go/src/sync/mutex.go:83 +0x1c4
main.(*Philosopher).Dine(0xc820076060, 0x64, 0xc8200620d0)
  /Users/andrewhsu/work/src/github.com/andrewhsu/dining-philosophers/main.go:48 +0xef6
created by main.main
  /Users/andrewhsu/work/src/github.com/andrewhsu/dining-philosophers/main.go:103 +0x575
```

Passing the `-nodeadlock` flag will prevent deadlock by switching the order the first philosopher takes the spork:

```
$ dining-philosophers -nodeadlock Aristotle Confucius
13:30:03.244390 Confucius taking left spork 1
13:30:03.244751 Aristotle taking right spork 1
13:30:03.315417 Confucius taking right spork 0
13:30:03.315446 Confucius eating
13:30:03.406533 Confucius placing right spork 0
13:30:03.406551 Confucius placing left spork 1
13:30:03.492629 Aristotle taking left spork 0
13:30:03.492643 Aristotle eating
13:30:03.525176 Aristotle placing left spork 0
13:30:03.525201 Aristotle placing right spork 1
```
