**Concurrency** is composition of independent execution computations, which may or may not run in parallel.
**Parallelism** is the ability to execute multiple computations simultaneously.

***Concurrency enables parallelism***

What is a process?
- An instance of a running program is a process
- Process provides environment for a program to execute
- Process has at least one thread - (main thread)
- Process can have multiple threads

What is a thread?
- Threads are the smallest unit of execution that CPU accepts
- Each thread has its own stack
- Threads share the same address space (memory)
- Threads can run independent of each other
- OS scheduler makes scheduling decision at thread level, not process level
- Threads can run concurrently or in parallel

Can we divide our application into Processes and Threads and achieve Concurrency?
- Yes! (But there are limitations...)
- Context Switches are expensive
  - Process Context - heavier then thread context (expensive!)
    - Process state
    - CPM scheduling info
    - Memory management info
    - Accounting info
    - I/O status info
  - Thread Context - lighter than process context (relatively cheap)
    - Program counter
    - CPU registers
    - Stack

    -> Cpu has to spend time (time not executing your program) copying the context of the current executing thread into memory and restoring the context of the next chosen thread

Can we scale the number of threads per process?
 - Not much, actually. For two reasons:
   - Fixed stack size & C10K problem
 - If we scale too high, we hit C10k problem
   - Whats that?
     - The CPU allocates a process a time slice for execution on a CPU core
     - The time slice is equally divided among the threads
     - [  Scheduler Period  ][ Number of Threads ][ Thread Time Slice]
     - [        10ms        ][        2          ][     5ms/thread   ]  
     - [        10ms        ][        5          ][     2ms/thread   ]
   - As time increases, time slice decreases until the CPU begins to spend more time context switching than solving the problem. This manifests itself in a system becoming less responsive

Fixed stack size. 
- Threads are allocated a fixed stack size. 
- `ulimit -a` // reveal stack size
- On my 2018 MB Pro - 8MB is stack size

Why is concurrency Hard?
 - Threads share memory
 - When multiple threads access memory, the result can be indeterministic.
   - Aka data race

We need Memmory access synchroniztion to solve this!
 - Locking can reduce parallelism
 - Inappropriate locks can lead to deadlocks

Go's Concurrency Tool Set
 - **goroutines** - concurrently executing functions
 - **channels** - allow communications between goroutines
 - **select** - is used to multiplex the channels (aka combine muliple signals into one channel)
 - **sync package** - provides classical synchronization tolls like mutex, conditional variables, and others

Goroutines
 - can think of them as **user space threads managed by go runtime**
 - goroutines are extremely lightweight. They start with 2KB of stack, which grows and shrinks as required
 - Low CPU overhead - 3 instructions per function call
 - Can create 100s of 1000s in same address space
 - Data is communicated between the goroutines using channels so sharing of memory can be avoided.
 - Context switching between goroutines is much **cheaper** than thread context switching
 - Go runtime can be more selective in what is persisted for retrieval, how it is persisted, and when the persisting needs to occur
 - Go runtime creates worker OS threads
 - Goroutines runs in the context of OS thread
  
Can you make main() wait for goroutine to execute before checking value of data?
- This is where the WaitGroup packge helps
- go statements fork the go routine, and after it is scheduled and work is complete, it joins back with main routine
- If main doesn't wait for the goroutine, then its very possible that the program will finish before the goroutine has completed.
- In order to create a join point, we use sync wiat group to deterministically block the main routine.