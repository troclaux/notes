# Operating Systems

> a system that intermediates between computer hardware and software

- functions:
  - process management
  - memory management
  - file system management
  - driver management
    - example: printer drivers
  - security and protection
  - graphical interface

## basic concepts

- kernel
  - the core of the OS
  - one of the first programs loaded during startup

types of kernels and other classifications:
- monolithic
- microkernel: a smaller-sized kernel
- hybrid: a middle ground between monolithic and microkernel
- exokernel systems

- layered systems
- virtual machines
- client-server model

- process: an instance of a program in execution
- multitasking system: the ability to run multiple processes simultaneously
- journaling: records the operations to be performed on the file system

- print spooler

## process creation, termination, and hierarchy

events that trigger process creation:
- system initialization
- a system call by a running process to create a new process
- user request to create a new process
  - example: double-clicking to open a word document
- starting a batch job

### process states

- running: using the processor (in execution)
- ready: temporarily halted, but queued to use the processor
- blocked (or waiting): unable to execute until some event occurs

- thread: the smallest unit of execution within a process
  - a process can contain multiple threads
  - threads can share resources, such as memory and file handles

- race conditions: occur when the behavior/result of software depends on the timing of access or the interaction of multiple threads or processes
- critical section/region:
  - part of the code that accesses shared resources
  - cannot be executed by more than one thread or process at a time
- mutual exclusion: a parallel programming property that states:
  - only one process should access resources during critical sections
- mutex:
  - ensures that only one process has access to resources during critical sections
- semaphore: a variable used to control access to a resource by multiple threads
  - used to prevent issues in critical sections
- Scheduler: determines the order in which processes gain access to the CPU
- Throughput: rate of process completion
- Preemption: the ability of the OS to switch which process has access to the CPU
- Starvation: occurs when a process is unable to proceed with execution due to not receiving computational resources
- Preemptive multitasking system

categories of scheduling algorithms:
- batch
  - slow response
  - uses schedulers with low or no preemption
  - used in mainframe computers
- interactive
  - user interaction requires preemption
  - used in personal computers
- real-time
  - unnecessary preemption
  - used in systems that run processes only when necessary
    - example: traffic control systems

## Scheduling Algorithms

- First Come First Serve (FCFS)
- Shortest Job First (SJF)
- Shortest Remaining Time Next (SRT)
- Round Robin

- Priority scheduling
- Multi-level queue scheduling
- Guaranteed scheduling
- Lottery scheduling

## Deadlock

> Occurs when a process depends on an event from another process and vice versa

There are 4 conditions for a deadlock to occur, called Coffman's conditions:
1. Mutual exclusion
  - Each resource is either being used by one process or is available
2. Hold and wait condition
  - Processes holding previously allocated resources can request new ones
3. No preemption
  - Allocated resources cannot be forcibly taken away from a process
4. Circular wait condition
  - Processes form a 'cycle' of resource waiting

Possible solutions:
- Ostrich algorithm
  - Ignores the problem
  - Simplest
  - Used by popular OSs (Windows, Linux)
- Banker's algorithm
  - Proposed by Edsger Dijkstra
  - Used in systems with a fixed number of resources
- Detection and recovery
  - The system monitors resource requests and releases
  - Whenever a resource is released, the resource graph is updated, and a check for cycles is performed
  - If a cycle is detected, one of the processes in the cycle is terminated
  - If this does not resolve the deadlock, another process is terminated until the deadlock ends
- Deadlock prevention
- Deadlock avoidance

# Memory Management

- Memory leak: occurs when a program fails to release memory that is no longer needed, gradually reducing the available memory for the system
  - Can occur in two situations:
    - Without pointers: memory blocks are allocated, but there are no pointers pointing to the desired memory area

- Swapping: a technique that transfers entire processes from main memory (RAM) to secondary memory (like SSD)
  - Used only when data is not present in RAM
- Memory bus: allows data and instruction transfer between CPU and memory

## Virtual Memory

- Allows processes to use secondary memory instead of main memory (RAM)
- Allows the system to run larger applications on machines with limited RAM
- Divides physical and virtual memory into small units called 'pages'

- Page: virtual memory
  - Usually has a fixed size of 4 KB
- Frame: physical memory (RAM)

Types of virtual memory:
- Paging
- Segmentation

- TLB (Translation Lookaside Buffer)
- Dirty bit: indicates whether a memory page has been modified since it was loaded into memory
- Thrashing: a phenomenon where the OS spends more time swapping pages between RAM and secondary memory than executing program instructions

### Segmentation Replacement Algorithms

- First-fit: allocates to the first space in memory that is larger or equal in size
- Next-fit: saves the location of the last free space used. The next search will start from the saved location
- Best-fit
- Quick-fit
- Circular-fit

### Page Replacement Algorithms

- Optimal
- NRU (Not Recently Used)
- LRU (Least Recently Used)
- FIFO (First In First Out)
- Clock

# Input/Output Management

TODO

# File Systems

TODO