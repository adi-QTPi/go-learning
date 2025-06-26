## Speed (why Go?)
- faster than interpretted languages (like python, js, python)
- 3 types of languages (slowest to fastest)
    * Interpretted -> Python, JS, Ruby, PHP
    * VM -> Java, C#
    * Compiled (natively) -> Rust, C, C++, Go
- then go should be the fastest because in the compiled language section, BUT.
Go slows down and becomes equivalent to C# level because of the Go-Runtime code which has to run everytime (this runtime code manages the memory).
- Thus go takes up much less memory.
- Go **runs faster** than interpretted languages and **compiles faster** than other compiled languages.

### Compilation?
- Compilation is the process of human readable code to machine language(binary)...
- `main.go` -> `$ go build` (run compiler) -> executable formed (some.exe);
- now you never need to run using go, u have the binary man!
### Interpretor?
- `main.py` wants `$ python3 main.py` to run it.
- if you want someone to use the product made using interpretted lang, u will have to essentially **provide him with** the source code!

While using compiled lang, just share the executable!\
-> distribution wise : products made using compiled languages are better.

## Memory management
- **Rust/C** -> manual allocation and management (Most efficient)
- **Go** 
    * garbage collected language(like Java) but nothing like JVM exists.
    * runtime code is added to it and garbage collection is taken care of.
    * bloated, extra code added always which manages the memory
- **Java** -> automated, done by java VM (or JVM) (java bytecode works within it) ; Takes a lot of memory (Least efficient)

## Go Runtime (maqsad)
- clean up unused memory.