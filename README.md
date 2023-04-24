Don't communicate by sharing memory, share memory by communicating

### Share Memory by Communicating

Or, in more deatil: Dont over-engineer things by using shared memory and complicated, error-prone synchronization primitives; instead, use message-passing between GoRoutines so variables and data can be used in the appropriate sequence.

### When to Use

- A golden rule for concurrency: if you don't need it, don't use it.

- Keep your application's complexity to an absolute minimum; it's easier to write, easier to understand, and easier to maintain.

### What Will be Covered

- We'll start by looking at the basic types in the sync package:
  mutexes (semaphores), and wait groups.

- We'll go through three of the classic computer science problems:
  - The Producer/Consumer problem
  - The Dining Philosopher problem
  - The Sleeping Barber problenm
