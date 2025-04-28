
# multiprocessor programming

> writing software that can run concurrently on multiple processors/cores to increase performance

- processes vs threads
  - process: independent program running in its own memory space
    - a process can contain multiple threads
    - doesn't share memory by default (they're isolated)
  - thread: the smallest unit of execution within a process
    - shares resources with other threads in the same process, like memory space
    - always belongs to a process

## parallelism and concurrency

- concurrency: ability of a program to manage multiple tasks at the same time
  - can happen on a single-core system with task switching
- parallelism: ability to execute multiple tasks at the same time
  - requires a multi-core system where multiple tasks can be executed simultaneously
- asynchronous: starting a task that runs on the background without blocking the main program flow
  - means that the program performs non-blocking tasks, usually involving I/O operations

## locks and synchronizations

- locks (mutexes): ensures that only one process has access to resources during critical sections
  - spinlock: type of lock used in multithreading environments
    - has this name because if a thread cannot acquire the lock, it will stay in a loop ("spin"), repeatedly checking until the lock becomes available
    - busy-waiting: technique where a thread repeatedly checks a condition in a loop, without doing any useful work
      - doesn't sleep or yield control, it just keeps checking: "waits while being busy"
    - best for short waits, where the lock is expected to be held for a very short time
    - wastes cpu cycles
- synchronization: when multiple threads access shared data, synchronization ensures that they don't interfere with each other

- atomic operations: operations that complete in a single step relative to other threads
  - e.g. `x++`, `compare-and-swap`
- semaphores: counters to control access to a shared resource pool
- race conditions: occur when two threads access shared data and try to change it at the same time
  - final result depends on the timing of the thread execution, which is unpredictable and dangerous
- mutual exclusion: a parallel programming property that states:
  - only one process should access resources during critical sections

## deadlock

> occurs when a process depends on an event from another process and vice versa

- there are 4 conditions for a deadlock to occur, called Coffman's conditions:
  - mutual exclusion: each resource is either being used by one process or is available
  - hold and wait condition: processes holding previously allocated resources can request new ones
  - no preemption: allocated resources cannot be forcibly taken away from a process
  - circular wait condition: processes form a 'cycle' of resource waiting
