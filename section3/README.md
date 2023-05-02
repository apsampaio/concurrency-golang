### sync.Mutex

- Mutex = "mutual exclusion" -- allows us to deal with race conditions
- Relative simple to use
- Dealing with shared resources and concurrent/parallel goroutines
- Lock/Unlock
- We can test for face conditions when running code, or testing it

#### Race Conditions

- Race conditions occur when multiple GoRoutines try to access the same data
- Can be difficult to spot when reading code
- Go allows us to check for them when running a program or when testing our code with go test

#### Channels

- Channels are a means of having GoRoutines share data
- They can talk to each other
- This is Go's philosophy of having things share memory by communicating, rather than communicating by sharing memory.
- The Producer/Consumer problem
