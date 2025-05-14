# Operating Systems

> software that manages computer hardware and software resources and provides essential services for computer programs and users

- functions:
  - hardware abstraction
    - provides standardized interface between complex hardware and software applications
    - allows programs to interact with hardware without knowing its specific technical details
    - enables software to run across different hardware configurations with minimal modifications
  - process management
    - scheduling CPU time
    - handles the creation, scheduling, and termination of computer processes
    - implements multitasking, allowing multiple programs to run seemingly simultaneously
  - memory management/allocation
    - implements virtual memory techniques to extend available memory using disk space
    - allocates and tracks memory usage for different programs
  - file system management
    - organizes and tracks files and directories
    - manages file permissions and access controls
    - provides methods for creating, reading, writing, and deleting files
  - driver management
    - example: printer drivers
  - security and protection
  - user interface (CLI and GUI)
  - networking

> [!NOTE]
> complete linux = kernel + GNU system utilities and libraries + other management scripts + installation scripts

## basic concepts

- kernel
  - the core of the OS
  - one of the first programs loaded during startup
  - acts as a bridge between hardware and user-level applications, controlling the allocation of system resources
  - responsible for:
    - hardware abstraction
    - process scheduling and management
    - memory allocation and management
    - file system operations (read/write)
    - device I/O operations
    - system calls handling
    - interrupt handling
    - inter-process communication (IPC)
- GNU system utilities
  - cp, mv, rm, mkdir, grep, sed, awk, sort, uniq, chmod, ssh, etc
- GNU libraries
  - glibc (GNU Clibrary)
    - string manipulation
    - memory management (`malloc`, `free`)
    - math operations
    - file I/O
    - networking
  - libstdc++ (GNU standard C++ library)
  - libm (GNU math library)
  - libutil (GNU utility library)
  - librt (GNU real-time library)
  - etc

- types of kernels and other classifications:
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

## program creation

1. os creates a new process
1. os uses virtual memory manager to allocate different regions of memory for the process
1. the stack is automatically created by the os
1. the heap is also set up by the os
1. execution starts

- stack: special region of memory that
  - automatically manages memory
  - is used for function calls (e.g. return addresses, function parameters, local variables)
  - Last-In, First-Out
- heap: another special region of memory
  - manages memory manually (you have to allocate and free)
  - used for dynamic memory allocation at runtime

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
- mutex: ensures that only one process has access to resources during critical sections
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

### process vs thread

- process: program under execution
  - can contain multiple threads
  - isolated memory
- thread: part of the process
  - takes less time to create and terminate in comparison to processes
  - threads within the same process can share memory
    - threads **cannot** share memory with threads from other processes

## Scheduling Algorithms

- First Come First Serve (FCFS)
- Shortest Job First (SJF)
- Shortest Remaining Time Next (SRT)
- Round Robin: preemptive algorithm that assigns each process/thread a fixed time slot in a rotating order

- Priority scheduling
- Multi-level queue scheduling
- Guaranteed scheduling
- Lottery scheduling

## Deadlock

> occurs when a process depends on an event from another process and vice versa

- there are 4 conditions for a deadlock to occur, called Coffman's conditions:
  - mutual exclusion: each resource is either being used by one process or is available
  - hold and wait condition: processes holding previously allocated resources can request new ones
  - no preemption: allocated resources cannot be forcibly taken away from a process
  - circular wait condition: processes form a 'cycle' of resource waiting

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

## Memory Management

- Memory leak: occurs when a program fails to release memory that is no longer needed, gradually reducing the available memory for the system
  - Can occur in two situations:
    - Without pointers: memory blocks are allocated, but there are no pointers pointing to the desired memory area

- Swapping: a technique that transfers entire processes from main memory (RAM) to secondary memory (like SSD)
  - Used only when data is not present in RAM
- Memory bus: allows data and instruction transfer between CPU and memory

### Virtual Memory

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

#### Segmentation Replacement Algorithms

- First-fit: allocates to the first space in memory that is larger or equal in size
- Next-fit: saves the location of the last free space used. The next search will start from the saved location
- Best-fit
- Quick-fit
- Circular-fit

#### Page Replacement Algorithms

- Optimal
- NRU (Not Recently Used)
- LRU (Least Recently Used)
- FIFO (First In First Out)
- Clock

### stack and heap memory

[stack and heap memory](https://courses.grainger.illinois.edu/cs225/fa2022/resources/stack-heap/)

- when a program is running, it takes up memory

- each program has its own memory layout, separated from other programs
- allocation and deallocation of memory is done automatically

- stack: stores local variables
- heap: dynamic memory for programmer to allocate
- data: stores global variables, separated as initialized and uninitialized
- text: stores the code being executed

```
| high address |
| stack |
| free memory |
| heap |
| uninitialized data |
| initialized data |
| text |
```

## Input/Output Management

TODO

## File Systems

TODO
