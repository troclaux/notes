
# compilers

> program that translates source code into machine code that can be executed directly by the computer's processor

- machine code: binary representation of instructions that can be executed by the CPU
- machine code != assembly code
  - assembly code is a human-readable representation of machine code
  - machine code is binary code that the CPU can execute directly

## compiled vs interpreted

- compiled languages: source code is transformed into machine code (binary code) before it is executed
  - e.g. C, C++, Golang, Rust
  - compiled binary is specific to the architecture and OS it was compiled for
  - a machine doesn't need the compiler to run the compiled executable
  - faster execution time
- interpreted languages: source code is executed line-by-line by an interpreter at runtime
  - e.g. JavaScript, Python, Ruby
  - as long as the OS has an interpreter, the code can be run on any platform
  - a machine needs the interpreter to run the source code
  - slower execution time

## compile time vs runtime

> [!IMPORTANT]
> compile time != runtime
> compile time: period when the source code is converted into machine code
> runtime: period when the program is actually running after it has been compiled

- interpreted languages typically skip full compilation to machine code, but may compile to bytecode before execution (e.g. `.py` => `.pyc` files)
- compile time happens before runtime for compiled languages

## compilation process

> process that translates source code written in high level programming languages into machine code

for example, this is how a C program is compiled and executed:

1. preprocessor (hello.c => hello.i)
1. compiler (hello.i => hello.s)
1. assembler (hello.s => hello.o)
1. linker (hello.o => hello)
1. loader (hello)
1. execution

> [!NOTE]
> the conversion of file extensions above is specific to C
> the assembler receives assembly code (input) and converts it into object code (output)
> the linker receives object files (input) and combines them into an executable file (output)
> the loader loads the executable file into memory and prepares it for execution

- object code: partially compiled machine code but not yet linked into a complete executable

- other languages' compiler/interpreter will translate to different types of files
  - for example, in python: file.py => bytecode => file.pyc
    - bytecode: low-level, platform-independent representation of your source code that the Python interpreter can execute
      - it is not machine code, but an intermediate step between source code and actual execution

1. preprocessing: focuses on textual replacements and macro expansion
- this is done by the preprocessor
- generates expanded code
2. compilation: source code is compiled into assembly code
- this is done by the compiler
3. assembly: assembly code is assembled into machine code
- this is done by the assembler
- machine code: binary representation of the code
4. linking: machine code is linked with libraries and other object files
- this is done by the linker
- object files: files that contain machine code that has been generated by a compiler or assembler
  - usually named with a .o or .obj extension
5. loading: executable file is loaded into memory by the operating system
- also sets up necessary memory protection and access controls
- access controls: restricts access of programs to computer resources
  - examples of access controls:
    - memory protection
    - file system access
    - network access
6. execution: CPU executes the machine code instructions in memory
- fetches instructions from memory
- decodes instructions
- executes instructions according to the Instruction Set Architecture (ISA)

### compiler phases (language-agnostic)

these are abstract steps that most compilers follow internally when transforming code, regardless of language:

1. lexical analysis: converts source code into tokens (e.g. keywords, operators, identifiers)
- tokenization: break source code into meaningful chunks called tokens
- error detection: catch errors like unrecognized symbols or invalid characters
- example:
  - input: `x = 2 + 3;`
  - output: tokens: `IDENTIFIER(x), ASSIGN(=), CONSTANT(2), PLUS(+), CONSTANT(3), SEMICOLON(;)`
2. syntax analysis (parsing): builds an AST using grammar rules (*)
- Abstract Syntax Tree (AST): tree representation of the structure of source code written in a programming language
- AST construction: combines tokens into tree nodes according to the programming languages grammar
3. semantic analysis: verify the AST for semantic correctness (e.g. type checking, scope resolution, error handling, etc)
4. intermediate code generation: translates the AST into an IR (*) that is simpler and closer to machine code
- IR (Intermediate Representation): a simplified, structured form of code used internally by compilers for analysis, optimization and code generation
  - e.g. python bytecode (`.pyc` files), java bytecode (used in java, kotlin, scala), LLVM IR (used in C, C++, rust, swift)
5. optimization: improves IR for performance (reduce runtime or memory usage)
6. code generation: converts IR to machine code/assembly code for the target architecture
7. assembly and linking: produce final executable
8. loader and execution: done by OS when program runs
- load executable into memory
- set up the runtime environment
- begin execution at the entry point (e.g. `main` function)

## field order optimization

- memory alignment
  - most processors access memory more efficiently if data is aligned to certain boundaries
  - e.g.: `int32` is typically aligned to a 4-byte boundary, and an `int64` is aligned to an 8-byte boundary
  - this **doesn't** apply to classes in Java and Python
    - because python objects are implemented as dictionaries internally
    - because JVM reorders fields in optimization
  - golang gives you more direct control over memory layout
- padding
  - if fields in a struct are not naturally aligned due to their order, the compiler inserts padding bytes between fields to ensure proper alignment
- struct size
  - the total size of a struct includes its fields and any padding added for alignment

example of bad field order

```go
type Example1 struct {
  a int8   // 1 byte
  b int64  // 8 bytes
  c int8   // 1 byte
}
```

- memory layout explanation:
  - `a`: 1 byte
  - padding: 7 bytes (after `a` to align `b` to an 8 byte boundary)
  - `b`: 8 bytes
  - `c`: 1 byte
  - padding: 7 bytes (after `c` to align the struct size to the largest alignment, which is 8 bytes)

example of optimized field order

```go
type Example2 struct {
  b int64  // 8 bytes
  a int8   // 1 byte
  c int8   // 1 byte
}
```

- `b`: 8 bytes
- `a`: 1 byte
- `c`: 1 byte
- padding: 6 bytes (added after c to align the struct size to the largest alignment, which is 8 bytes)

> [!IMPORTANT]
> general rule of thumb to minimize padding
> group fields of the same size together
> order fields from largest to smallest in terms of size (e.g. `int64`, then `int32`, then `int16`, and finally `int8`)

---

- Abstract Syntax Tree (AST): tree representation of the structure of source code written in a programming language
- transpiled languages: typescript
  - source code is converted into another language
  - e.g. typescript => JavaScript, sass => css
  - transpilation typically converts between languages at a similar abstraction level (e.g., TypeScript → JavaScript), unlike compilation to machine code

