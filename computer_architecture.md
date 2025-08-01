[os](./os.md)

# computer architecture

> design and organization of a computer's core components

- CPU (Central Processing Unit): executes instructions
  - Control Unit: directs operations of the processor
  - Arithmetic Logic Unit (ALU): Performs arithmetic and logic operations
  - registers: Small, fast storage locations inside the CPU for temporary data
- Memory (RAM, cache, registers): stores data and programs
  - RAM (Random Access Memory): volatile memory used for storing data and instructions currently being used
  - Cache: Small, fast memory close to the CPU for frequently accessed data
  - Storage (e.g. SSD/HDD): non-volatile memory for long-term data storage
- Input/Output systems: allow communication between computer and the outside world
  - e.g. keyboard, mouse, display, printer, network card

## memory

- primary: stores the OS, running programs and the data those programs use during computer usage
  - directly accessible by the CPU
  - fast and temporary
  - e.g. RAM, cache, registers

- secondary: non-volatile memory used for long-term data storage
  - non-volatile: data is preserved when power is off
  - slower (compared to RAM)
  - cheaper (per GB)
  - e.g. HDD, SSD, USB flash drives

- how they work together:
  - when running a program:
    - the program is loaded from secondary memory into RAM
    - the CPU accesses the data from RAM, but it first pulls frequently used data into cache and ultimately into registers for immediate use
    - frequently used data may even go into cache for faster access

### memory hierarchy

the closer to the cpu, the faster and more expensive the memory

speed and cost: NVM (SSD/HDD) < DRAM (RAM) < SRAM (Cache) < Registers

### types of memory technology

- DRAM (Dynamic RAM)
  - used in main memory (RAM sticks in computer)
  - volatile: loses data when power is off
  - slower than SRAM, but cheaper
  - needs refresh: stores data using capacitors that leak charge, so it must be refreshed thousands of times per second
- SRAM (Static RAM)
  - used in CPU cache (L1, L2, L3)
  - volatile
  - faster than DRAM
- NVM (Non-Volatile Memory)
  - used in SSDs, USB drives, BIOS/firmware
  - non-volatile
  - slower than DRAM/SRAM, but persistent

speed comparison: NVM is slower than DRAM is slower than SRAM

## system buses

> transfer data between components

## instruction cycle (fetch-decode-execute)

The CPU executes programs by repeating this cycle:

1. fetch: get the instruction from memory
1. decode: interpret the instruction
1. execute: perform the operation

## Instruction Set Architecture (ISA)

> defines the machine language understood by a CPU

examples: x86, ARM, RISC-V

- defines:
  - supported instructions (e.g. add, load, jump)
  - data types
  - registers
  - memory addressing modes

## performance factors

- Clock speed (GHz): Determines how many cycles per second the CPU can perform
- CPI (Cycles Per Instruction): A measure of how efficiently instructions are executed
- Pipelining: Overlaps instruction stages to improve throughput
- Parallelism: Uses multiple cores or threads for simultaneous operations

## von neumann architecture

> computer design model that forms the foundation for most modern computers

- key features
  - single memory for data and instructions
    - both program instructions and data are stores in the same memory space
    - this allows the CPU to fetch and execute instruction sequentially
  - Central Processing Unit (CPU)
    - Arithmetic Logic Unit (ALU)
    - Control Unit (CU)
  - input/output system
  - sequential execution of instructions
  - program stored in memory
- innovations
  - programs can be written, stored and modified electronically
    - no need to rewire circuits for each task
    - previously, computers were hardwired to perform specific tasks
  - a computer can load and run different programs without hardware changes
  - programming languages and compilers became possible
  - software can be debugged, improved and reused

![von neumann architecture](./images/von_neumann_architecture.png)

- Memory: where data and instructions are stored, it allows for the storage and retrieval of information
- Control Unit: manages the operations of the computer, it directs the flow of data between the CPU and other components
- Arithmetic Logic Unit (ALU): performs arithmetic and logical operations
  - responsible for calculations and decision-making processes
- Input: This refers to the devices or methods through which data is entered into the computer system
- Output: This refers to the devices or methods through which data is presented to the user or other systems
- Central Processing Unit (CPU): central component that carries out the instructions of a computer program
  - contains the ALU and Control Unit
- Accumulator: This is a register in the CPU that stores intermediate results of arithmetic and logic operations

---

- firmware: software that is stored in hardware and tells it how to work
  - e.g. BIOS or UEFI (starts boot process in computers)

