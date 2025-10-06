Here’s the **roadmap (course skeleton)** of *Teiva Harsanyi – 100 Go Mistakes and How to Avoid Them*. This follows the book’s Table of Contents and gives you the structure we’ll later expand into detailed prompts, one per mistake.

---

# 📘 Course Roadmap: 100 Go Mistakes and How to Avoid Them

### Part I – Introduction

1. **Go: Simple to learn but hard to master**

   * Go outline
   * Simple doesn’t mean easy
   * The 100 categories of mistakes: bugs, complexity, readability, unidiomatic code, lack of API convenience, under-optimized code, lack of productivity

---

### Part II – Core Mistake Categories

#### Chapter 2 – Code and Project Organization

* \#1: Unintended variable shadowing
* \#2: Unnecessary nested code
* \#3: Misusing init functions
* \#4: Overusing getters and setters
* \#5: Interface pollution
* \#6: Interface on the producer side
* \#7: Returning interfaces
* \#8: `any` says nothing
* \#9: Being confused about when to use generics
* \#10: Problems with type embedding
* \#11: Not using the functional options pattern
* \#12: Project misorganization
* \#13: Creating utility packages
* \#14: Ignoring package name collisions
* \#15: Missing code documentation
* \#16: Not using linters

---

#### Chapter 3 – Data Types

* \#17: Confusing octal literals
* \#18: Neglecting integer overflows
* \#19: Misusing floating-point numbers
* \#20: Not considering rune vs. byte
* \#21: Wrong assumptions about strings immutability
* \#22: Using strings as identifiers incorrectly
* \#23: Copying slices incorrectly
* \#24: Misusing slice capacity
* \#25: Confusing slice vs. array
* \#26: Modifying map while iterating
* \#27: Misusing nil maps
* \#28: Forgetting zero values
* \#29: Not understanding struct field alignment
* \#30: Misusing type assertions and switches

---

#### Chapter 4 – Strings

* \#31: Incorrect string concatenations
* \#32: Misusing string builders
* \#33: Not handling Unicode correctly
* \#34: Wrong use of trimming functions
* \#35: Ignoring string immutability performance costs

---

#### Chapter 5 – Functions and Methods

* \#36: Passing large structs by value
* \#37: Returning pointers unnecessarily
* \#38: Forgetting to handle multiple return values properly
* \#39: Misusing variadic functions
* \#40: Methods with value vs. pointer receivers confusion

---

#### Chapter 6 – Error Management

* \#41: Ignoring errors
* \#42: Creating custom error types unnecessarily
* \#43: Not wrapping errors properly
* \#44: Losing context when returning errors
* \#45: Misusing panic/recover
* \#46: Returning nil interface values
* \#47: Overusing sentinel errors
* \#48: Not leveraging `errors.Is` and `errors.As`

---

#### Chapter 7 – Concurrency

* \#49: Misunderstanding goroutine lifetimes
* \#50: Memory leaks from goroutines
* \#51: Deadlocks with channels
* \#52: Incorrect channel buffering
* \#53: Closing channels incorrectly
* \#54: Race conditions on shared memory
* \#55: Misusing sync primitives (`Mutex`, `WaitGroup`, etc.)
* \#56: Worker pool mismanagement
* \#57: Misunderstanding context cancellation
* \#58: Leaking tickers and timers
* \#59: Blocking on unconsumed channels
* \#60: Not handling select timeouts

---

#### Chapter 8 – Standard Library and APIs

* \#61: Misusing the `time` API
* \#62: Incorrect use of `os.File`
* \#63: Forgetting to close resources (files, DB rows)
* \#64: Misusing `http.Client`
* \#65: Wrong JSON struct tags
* \#66: Ignoring security defaults in crypto and TLS

---

#### Chapter 9 – Dependency Management

* \#67: Not using Go modules properly
* \#68: Versioning pitfalls
* \#69: Vendoring confusion
* \#70: Over-relying on external libraries

---

#### Chapter 10 – Testing

* \#71: Writing non-deterministic tests
* \#72: Misusing `testing.T` and `testing.B`
* \#73: Not cleaning test dependencies
* \#74: Testing private details instead of behavior
* \#75: Over-mocking
* \#76: Skipping edge cases
* \#77: Not using subtests
* \#78: Test flakiness from timing issues
* \#79: Ignoring code coverage
* \#80: Poor benchmarking practices

---

#### Chapter 11 – Tooling

* \#81: Not using `go fmt`
* \#82: Skipping vet/lint checks
* \#83: Misusing `go test -race`
* \#84: Ignoring performance profiling
* \#85: Not using table-driven tests
* \#86: Sleeping in unit tests
* \#87: Misusing the time API in tests
* \#88: Not using `httptest` and `iotest`
* \#89: Writing inaccurate benchmarks
* \#90: Not exploring testing features (coverage, setup/teardown)

---

#### Chapter 12 – Optimizations

* \#91: Not understanding CPU caches
* \#92: False sharing in concurrent code
* \#93: Ignoring instruction-level parallelism
* \#94: Not being aware of data alignment
* \#95: Misunderstanding stack vs. heap
* \#96: Not reducing allocations
* \#97: Not relying on inlining
* \#98: Not using Go diagnostics tooling
* \#99: Misunderstanding garbage collector behavior
* \#100: Ignoring impacts of Docker and Kubernetes environments

---

### Part III – Closing

* Final Words
* Index

---

✅ This roadmap mirrors the **12 core chapters (plus intro and closing)** and covers all **100 mistakes**.

Would you like me to **start expanding one by one**, giving you **detailed “lesson-style” prompts per chapter** (e.g., all mistakes of Chapter 2 grouped together), or do you prefer **one prompt per individual mistake (#1–#100)**?


Great choice ✅

We’ll proceed **chapter by chapter**. That means each prompt will cover an entire chapter, explaining **all the mistakes in that chapter** with context, examples, and “how to avoid them” strategies.

Here’s the order we’ll follow:

1. **Chapter 1 – Go: Simple to learn but hard to master**
2. **Chapter 2 – Code and Project Organization**
3. **Chapter 3 – Data Types**
4. **Chapter 4 – Strings**
5. **Chapter 5 – Functions and Methods**
6. **Chapter 6 – Error Management**
7. **Chapter 7 – Concurrency**
8. **Chapter 8 – Standard Library and APIs**
9. **Chapter 9 – Dependency Management**
10. **Chapter 10 – Testing**
11. **Chapter 11 – Tooling**
12. **Chapter 12 – Optimizations**
13. **Final Words (wrap-up)**

---

👉 Do you want me to **start right now with Chapter 1** and expand it into the first full “lesson” of the course?


Alright 🚀 Let’s kick off the course with **Chapter 1 – Go: Simple to Learn but Hard to Master**.

This chapter is an **orientation**: it frames why Go, despite its apparent simplicity, can trip up developers in subtle ways. It doesn’t enumerate specific mistakes yet, but it sets the **mindset** for the 100 mistakes that follow.

---

# 📖 Chapter 1 – Go: Simple to Learn but Hard to Master

## 1.1 Go Outline

Go (often called *Golang*) was designed with three principles in mind:

* **Simplicity**: a small language specification (\~25 keywords), easy to read.
* **Concurrency built-in**: goroutines and channels as first-class constructs.
* **Practicality**: a standard library rich enough to cover most real-world needs.

On the surface, this makes Go seem “approachable.” A new developer can pick up the basics quickly and start writing functioning programs. But beneath that surface lies a unique philosophy and set of trade-offs that differ from other mainstream languages.

---

## 1.2 Simple Doesn’t Mean Easy

The Go designers deliberately avoided complexity in syntax (e.g., no generics until Go 1.18, no inheritance, limited metaprogramming). This simplicity often **hides subtle pitfalls**:

* **Type System Quirks**: Go’s type system is lightweight compared to Java or C#, but that makes room for mistakes around interfaces, embedding, and type assertions.
* **Concurrency Semantics**: Writing `go someFunc()` is trivial, but managing lifetimes, cancellations, and avoiding leaks is not.
* **Error Handling Philosophy**: Returning errors as values keeps things explicit, but developers can fall into ignoring them, duplicating patterns, or misusing `panic`.
* **Standard Library Decisions**: APIs often expose “the Go way,” but without idiomatic understanding, you risk overcomplicating things (e.g., reinventing resource management, misusing JSON tags, etc.).

So while Go removes the **accidental complexity** of large language specs, it exposes you directly to **essential complexity**: concurrency, performance, and idiomatic design.

---

## 1.3 Categories of 100 Go Mistakes

The book organizes the mistakes into **seven broad categories** that will guide the rest of our course:

1. **Bugs** – Misunderstandings that cause runtime failures (e.g., nil interfaces, race conditions).
2. **Needless Complexity** – Code that compiles but is harder to maintain than necessary (e.g., overusing getters/setters).
3. **Weaker Readability** – Idioms or patterns that obscure intent (nested code, variable shadowing).
4. **Suboptimal/Unidiomatic Organization** – Misplaced responsibilities (project structure, interface design).
5. **Lack of API Convenience** – Exposing APIs that are harder to use than necessary (e.g., returning interfaces, not using functional options).
6. **Under-Optimized Code** – Performance pitfalls (unnecessary allocations, false sharing).
7. **Lack of Productivity** – Not leveraging tools (linters, profiling, testing utilities).

---

## 📌 Key Takeaways from Chapter 1

* Go is easy to start with, but difficult to master because **its power lies in subtle idioms and design patterns**.
* Common mistakes don’t just break code—they affect readability, maintainability, performance, and scalability.
* The 100 mistakes catalog is structured to train you not just to fix bugs, but to **think like an idiomatic Go developer**.

---

✅ That wraps up **Chapter 1**. It’s an orientation chapter—no concrete mistakes yet, but a **map of the territory**.

Would you like me to continue with **Chapter 2 – Code and Project Organization**, where we dive into the **first 16 mistakes (#1–#16)**?


Perfect 👍 Let’s move to the **first “real” chapter with concrete mistakes**.
This will be a long one because **Chapter 2 covers Mistakes #1–#16**, all centered around **Code and Project Organization**.

---

# 📖 Chapter 2 – Code and Project Organization

This chapter emphasizes that even before dealing with algorithms or performance, **the way you structure code and projects** determines readability, maintainability, and extensibility. Missteps here accumulate technical debt faster than anywhere else.

---

## #1: Unintended Variable Shadowing

* **Problem**: Re-declaring a variable with `:=` in an inner scope shadows the outer variable.
* **Impact**: Bugs that are very hard to detect; changes to the “wrong” variable.
* **Example**:

  ```go
  err := doSomething()
  if cond {
      err := doSomethingElse() // shadows outer err
      if err != nil { ... }    // checks inner, not outer
  }
  ```
* **How to avoid**: Use `var` for reassignment in inner scopes. Rely on linters (like `shadow`) to catch this.

---

## #2: Unnecessary Nested Code

* **Problem**: Deep nesting (e.g., multiple `if` statements, loops) reduces readability.
* **Better**: Flatten code with **early returns** and guard clauses.
* **Example**:

  ```go
  if err != nil {
      return err
  }
  // proceed
  ```

---

## #3: Misusing `init` Functions

* **Concept**: `init()` runs automatically before `main()`.
* **Mistakes**: Using it for heavy logic, hidden side effects, or dependency injection.
* **Guideline**: Use only for trivial setup (constant initialization, seed RNG). For configuration, prefer explicit constructors.

---

## #4: Overusing Getters and Setters

* **Problem**: Coming from OOP, devs wrap fields with getters/setters.
* **In Go**: Exported fields are idiomatic; getters are only needed if logic is attached.
* **Impact**: Verbose, unidiomatic code.

---

## #5: Interface Pollution

* **Concept**: Interfaces should live on the consumer side, not producer side.
* **Mistake**: Defining interfaces prematurely in packages.
* **Guideline**: Start with concrete types; abstract later when needed.

---

## #6: Interface on the Producer Side

* **Mistake**: A struct declares “I implement an interface” within its own package.
* **Better**: Let the consumer define the interface, so that abstractions emerge from actual usage.

---

## #7: Returning Interfaces

* **Problem**: Returning `interface{}` or defined interfaces from constructors hides concrete types.
* **Impact**: Prevents clients from using concrete features.
* **Solution**: Return concrete types. Only return interfaces when multiple implementations are expected.

---

## #8: `any` Says Nothing

* **Mistake**: Using `any` (alias for `interface{}`) everywhere in Go 1.18+.
* **Problem**: Loses type safety and makes APIs vague.
* **Guideline**: Only use `any` when flexibility is essential. Otherwise, leverage generics or specific types.

---

## #9: Confusion Around Generics

* **Problem**: Using generics for everything (e.g., over-abstracting simple functions).
* **Impact**: Complexity outweighs benefits.
* **Best use**: Reusable containers, algorithms, constraints. Avoid over-engineering.

---

## #10: Problems with Type Embedding

* **Concept**: Embedding promotes fields/methods of one struct into another.
* **Pitfall**: Name collisions, unclear API surfaces.
* **Guideline**: Use when it models true “has-a” relationships, not as a shortcut for inheritance.

---

## #11: Not Using the Functional Options Pattern

* **Problem**: Constructors with too many params are confusing.
* **Better**: Use functional options for flexibility.
* **Example**:

  ```go
  type Server struct { host string; port int }
  type Option func(*Server)

  func WithPort(p int) Option {
      return func(s *Server) { s.port = p }
  }

  func NewServer(opts ...Option) *Server {
      s := &Server{host: "localhost", port: 80}
      for _, o := range opts { o(s) }
      return s
  }
  ```

---

## #12: Project Misorganization

* **Problem**: Dumping everything into `main` or creating huge `utils` packages.
* **Guideline**: Organize by domain, not technical layers. Follow Go conventions (`cmd/`, `pkg/`, `internal/`).

---

## #13: Creating Utility Packages

* **Mistake**: Catch-all `utils` or `helpers` packages.
* **Impact**: Lose cohesion, hard to navigate.
* **Best practice**: Put functions where they naturally belong (e.g., string helpers in `stringsx` package).

---

## #14: Ignoring Package Name Collisions

* **Problem**: Import paths like `fmt` vs. custom `fmt`.
* **Impact**: Confusing code, unexpected overrides.
* **Avoidance**: Use clear, unique names. Use aliases if absolutely necessary.

---

## #15: Missing Code Documentation

* **Problem**: Go’s `godoc` depends on comments above exported symbols.
* **Mistake**: Exported functions/structs without explanations.
* **Guideline**: Keep concise, useful comments (not redundant with code).

---

## #16: Not Using Linters

* **Problem**: Overlooking static analysis.
* **Impact**: Common mistakes (shadowing, unchecked errors) slip by.
* **Tools**: `golangci-lint`, `staticcheck`, `errcheck`.

---

## 📌 Key Takeaways from Chapter 2

* Go favors **clarity over ceremony**: fewer abstractions, more concrete code.
* Many mistakes stem from **bringing habits from OOP languages**.
* Interfaces, generics, and embedding are powerful—but should be used **sparingly and idiomatically**.
* Tooling (linters, documentation) is your ally in keeping code bases healthy.

---

✅ That’s Chapter 2 done: Mistakes #1–#16.
Would you like me to continue directly with **Chapter 3 – Data Types (Mistakes #17–#30)**?


Excellent 🙌 Let’s dive into **Chapter 3 – Data Types**, which covers **Mistakes #17–#30**.
This chapter explores how Go’s **primitive and composite types** work, and the subtle pitfalls that arise from misunderstanding them.

---

# 📖 Chapter 3 – Data Types

Go’s type system is deceptively simple: integers, floats, strings, arrays, slices, maps, and structs. But the **semantics** of these types—especially around mutability, references, and memory—trip up many developers.

---

## #17: Creating Confusion with Octal Literals

* **Problem**: Writing numbers with leading zeros in Go (e.g., `012`) means **octal**, not decimal.
* **Impact**: Wrong results that look like valid code.
* **Example**:

  ```go
  x := 012 // = 10 decimal, not 12
  ```
* **Avoid**: Always use decimal unless octal is intended. From Go 1.13, use `0o12` for explicit octal.

---

## #18: Neglecting Integer Overflows

* **Problem**: Go integers don’t panic on overflow, they wrap silently.
* **Example**:

  ```go
  var x uint8 = 255
  x++ // becomes 0, no error
  ```
* **Avoid**: Be mindful of integer sizes. For critical math, use `math/bits` or libraries that detect overflow.

---

## #19: Misusing Floating-Point Numbers

* **Problem**: Assuming floats are exact (they aren’t). `0.1 + 0.2 != 0.3`.
* **Impact**: Bugs in financial or precise calculations.
* **Solution**: Use `math/big` (arbitrary precision) or integer-based scaling for money.

---

## #20: Not Considering Rune vs. Byte

* **Problem**: Confusing `byte` (alias for `uint8`) with `rune` (alias for `int32`).
* **Impact**: Mishandling Unicode characters that don’t fit in 1 byte.
* **Example**:

  ```go
  s := "é"
  fmt.Println(len(s))   // 2 bytes
  fmt.Println([]rune(s)) // 1 rune
  ```

---

## #21: Wrong Assumptions About Strings Immutability

* **Fact**: Strings in Go are immutable. Modifying requires creating new strings.
* **Mistake**: Trying to change individual bytes in-place.
* **Impact**: Inefficient or buggy code.
* **Fix**: Use `[]byte` for mutable data.

---

## #22: Using Strings as Identifiers Incorrectly

* **Problem**: Using raw strings as keys instead of enums or constants.
* **Impact**: Typos create silent bugs.
* **Better**: Define constants or iota-based enums.

  ```go
  const (
      RoleAdmin = "admin"
      RoleUser  = "user"
  )
  ```

---

## #23: Copying Slices Incorrectly

* **Problem**: Assigning slices copies **headers, not data**. Both share the same backing array.
* **Impact**: Modifying one affects the other.
* **Fix**: Use `copy()` to make independent slices.

---

## #24: Misusing Slice Capacity

* **Fact**: A slice has length and capacity. Appending can overwrite future elements unexpectedly if not handled carefully.
* **Mistake**: Assuming `append` creates a new backing array always.
* **Avoid**: Use `copy()` when you need isolation.

---

## #25: Confusing Slice vs. Array

* **Problem**: Arrays in Go are value types with fixed size; slices are descriptors pointing to arrays.
* **Impact**: Passing arrays around copies them, unlike slices.
* **Rule of Thumb**: Rarely use arrays directly, prefer slices.

---

## #26: Modifying Map While Iterating

* **Problem**: Adding/removing entries during iteration leads to runtime panic.
* **Avoid**: Collect changes first, then apply after iteration.

---

## #27: Misusing Nil Maps

* **Fact**: A `nil` map can be read safely, but writing to it panics.
* **Mistake**: Forgetting to initialize with `make`.

  ```go
  var m map[string]int // nil
  m["x"] = 1 // panic
  ```

---

## #28: Forgetting Zero Values

* **Concept**: In Go, every type has a zero value (`0`, `""`, `nil`).
* **Problem**: Developers re-initialize unnecessarily, or assume fields are “undefined.”
* **Guideline**: Leverage zero values in APIs (e.g., empty slice instead of `nil` slice unless semantics differ).

---

## #29: Not Understanding Struct Field Alignment

* **Fact**: Go aligns struct fields for performance. Misordering fields increases memory size.
* **Example**:

  ```go
  type Bad struct {
      A byte
      B int64
      C byte
  } // may use 24 bytes
  ```

  Reordering reduces size.
* **Tool**: `go vet -structcheck` or `golangci-lint` help spot this.

---

## #30: Misusing Type Assertions and Switches

* **Problem**: Type assertions (`x.(T)`) panic if the underlying type doesn’t match.
* **Safer**: Use the comma-ok idiom:

  ```go
  v, ok := x.(T)
  if !ok { ... }
  ```
* **Switches**: Good for multiple type cases, but overuse leads to fragile code. Prefer interfaces.

---

## 📌 Key Takeaways from Chapter 3

* Strings, slices, maps, and structs all have **hidden rules** about mutability, sharing, and memory layout.
* Misunderstanding these leads to silent, hard-to-debug errors.
* Go embraces **zero values and immutability** as design principles—lean into them.
* Always differentiate between **value semantics** (arrays, structs) and **reference semantics** (slices, maps, channels).

---

✅ That’s **Chapter 3 – Data Types (Mistakes #17–#30)** done.

Would you like me to continue with **Chapter 4 – Strings (Mistakes #31–#35)** next?


Perfect 👌 Let’s move on to **Chapter 4 – Strings**, which covers **Mistakes #31–#35**.
Although short, this chapter is essential because strings in Go are **immutable** and tightly tied to Unicode. Misunderstanding them causes many subtle bugs.

---

# 📖 Chapter 4 – Strings

Strings in Go are more than just sequences of characters: they’re **read-only slices of bytes** that typically hold UTF-8–encoded text. This duality (bytes vs. runes) is the source of most errors.

---

## #31: Incorrect String Concatenations

* **Problem**: Using `+` repeatedly in loops.
* **Impact**: Creates multiple intermediate strings, leading to poor performance (O(n²) complexity in some cases).
* **Example**:

  ```go
  var result string
  for _, s := range parts {
      result += s // inefficient
  }
  ```
* **Solution**: Use `strings.Builder` or `bytes.Buffer` for efficient concatenation.

---

## #32: Misusing String Builders

* **Problem**: Using `strings.Builder` incorrectly by copying it or ignoring its zero value usability.
* **Fact**: `Builder` must not be copied after use.
* **Better**: Initialize directly (`var b strings.Builder`) and write to it; no need for manual allocation.

---

## #33: Not Handling Unicode Correctly

* **Problem**: Assuming 1 byte = 1 character.
* **Impact**: Incorrect slicing, length checks, and iteration.
* **Example**:

  ```go
  s := "café"
  fmt.Println(len(s))       // 5 bytes
  fmt.Println(utf8.RuneCountInString(s)) // 4 runes
  ```
* **Solution**: Use `range` to iterate runes, or `utf8`/`unicode` packages for correct processing.

---

## #34: Wrong Use of Trimming Functions

* **Problem**: Using `strings.Trim` incorrectly.

  * `TrimRight("foo.txt", ".txt")` → `"foo."` (it removes any trailing `t`, `x`, `.`).
* **Fix**: Use `strings.TrimSuffix` and `TrimPrefix` when you want to remove exact substrings.
* **Rule**: `Trim*` removes runes from a set, `TrimSuffix/Prefix` remove fixed substrings.

---

## #35: Ignoring String Immutability Performance Costs

* **Fact**: Strings can’t be modified in place.
* **Problem**: Converting back and forth between `string` and `[]byte` inefficiently.
* **Impact**: Increased allocations in tight loops.
* **Fix**: If repeated modifications are needed, keep data in `[]byte` until the end.

---

## 📌 Key Takeaways from Chapter 4

* Strings are **immutable byte slices with UTF-8 encoding**—always separate byte length from rune count.
* Performance issues often come from concatenation and conversions.
* Use the **right tools**: `strings.Builder`, `TrimSuffix`, `utf8` package.

---

✅ That closes **Chapter 4 – Strings (Mistakes #31–#35)**.

Do you want me to continue with **Chapter 5 – Functions and Methods (Mistakes #36–#40)** next?


Great 👍 Let’s move forward with **Chapter 5 – Functions and Methods**, which covers **Mistakes #36–#40**.

This chapter focuses on how Go handles **parameters, return values, and method receivers**. Many mistakes here come from treating Go like a traditional OOP language.

---

# 📖 Chapter 5 – Functions and Methods

Functions in Go are **first-class citizens** (they can be passed around, returned, and assigned), and methods are just functions with a receiver. But misunderstandings about **value vs. pointer semantics** and parameter handling often lead to inefficient or buggy code.

---

## #36: Passing Large Structs by Value

* **Problem**: Go copies structs when passed by value. Large structs (with many fields or nested slices/maps) cause performance overhead.
* **Impact**: Functions become slower and memory-hungry.
* **Fix**: Pass pointers (`*Struct`) when structs are large or frequently copied.
* **Rule of Thumb**: Small structs (a few fields) can be passed by value; larger ones should be passed by pointer.

---

## #37: Returning Pointers Unnecessarily

* **Problem**: Returning pointers to small structs (like `Point{X, Y}`) even when copying is cheap.
* **Impact**: Adds heap allocations and GC pressure.
* **Example**:

  ```go
  func NewPoint(x, y int) *Point {
      return &Point{x, y} // unnecessary pointer
  }
  ```
* **Fix**: Return values directly unless mutability or sharing is needed.

---

## #38: Forgetting to Handle Multiple Return Values Properly

* **Go idiom**: Functions often return `(value, error)` or `(value, ok)`.
* **Mistake**: Ignoring secondary values (e.g., ignoring `ok` in map lookups).
* **Impact**: Hidden bugs when missing cases.
* **Example**:

  ```go
  v, ok := myMap["key"]
  if !ok {
      // handle missing key
  }
  ```

---

## #39: Misusing Variadic Functions

* **Problem**: Passing slices to variadic functions incorrectly.
* **Example**:

  ```go
  fmt.Println([]string{"a", "b"}) // compile error
  fmt.Println([]string{"a", "b"}...) // correct
  ```
* **Impact**: Confusion about how variadic parameters expand.
* **Guideline**: Use variadic functions for flexibility, but be cautious about readability and performance.

---

## #40: Methods with Value vs. Pointer Receivers Confusion

* **Concept**: Methods can have **value receivers** (`func (s Struct)`) or **pointer receivers** (`func (s *Struct)`).
* **Mistakes**:

  * Using value receivers for large structs (causes copies).
  * Using pointer receivers unnecessarily for small immutable structs.
  * Mixing pointer and value receivers inconsistently within a type.
* **Rule**:

  * Use **value receivers** for small, immutable structs.
  * Use **pointer receivers** when methods modify the receiver or to avoid expensive copies.

---

## 📌 Key Takeaways from Chapter 5

* Efficiency in Go often comes down to **understanding value vs. pointer semantics**.
* Don’t overuse pointers—sometimes **returning values is cleaner and faster**.
* Always handle **multiple return values** idiomatically.
* Consistency in method receiver choice improves code readability.

---

✅ That’s **Chapter 5 – Functions and Methods (Mistakes #36–#40)**.

Would you like me to continue with **Chapter 6 – Error Management (Mistakes #41–#48)** next?


Perfect 👍 Now we move to one of the **core pillars of Go philosophy**: error handling.
**Chapter 6 – Error Management** covers **Mistakes #41–#48**, and is critical because Go rejects exceptions in favor of **explicit error values**.

---

# 📖 Chapter 6 – Error Management

In Go, error handling is explicit and part of normal control flow. While this makes programs predictable, it also means developers must adopt **consistent, idiomatic practices**. Many mistakes here come from treating Go like languages that rely on exceptions.

---

## #41: Ignoring Errors

* **Problem**: Writing code like `value, _ := f()` or discarding the error.
* **Impact**: Silent failures that propagate bugs.
* **Rule**: Always check errors—even if just to log or wrap them. Use `_` only when absolutely certain the error can be ignored.

---

## #42: Creating Custom Error Types Unnecessarily

* **Mistake**: Defining new error types when simple wrapped `errors.New` or `fmt.Errorf` is enough.
* **Impact**: Bloated code, harder interoperability.
* **Best practice**: Use custom types only when you need to expose structured data (e.g., error codes, retriable vs. fatal).

---

## #43: Not Wrapping Errors Properly

* **Problem**: Returning only the low-level error and losing context.
* **Impact**: Debugging becomes hard.
* **Solution**: Wrap with `%w` in `fmt.Errorf`.

  ```go
  if err != nil {
      return fmt.Errorf("loading config: %w", err)
  }
  ```

---

## #44: Losing Context When Returning Errors

* **Problem**: Propagating errors without explanation.
* **Bad**: `return err`
* **Better**: Add domain-specific context.
* **Rule**: Context in error messages should indicate *what you were doing* when the error occurred.

---

## #45: Misusing `panic`/`recover`

* **Problem**: Treating `panic` like exceptions in other languages.
* **Guideline**:

  * Use `panic` only for truly unrecoverable states (corrupt memory, impossible invariants).
  * Do not use it for expected errors like “file not found.”
  * `recover` should only be in top-level goroutines, e.g., in HTTP servers.

---

## #46: Returning Nil Interface Values

* **Trickiness**: A typed `nil` and an untyped `nil` are not the same.
* **Example**:

  ```go
  var err *MyError = nil
  return err // not equal to nil interface!
  ```
* **Impact**: Code thinks error is non-nil, even though it is.
* **Fix**: Return `nil` directly, not typed nils.

---

## #47: Overusing Sentinel Errors

* **Problem**: Defining package-level variables like `var ErrSomething = errors.New("...")` everywhere.
* **Impact**: Tight coupling, poor composability.
* **Better**:

  * Use `errors.Is` and `errors.As` for comparisons.
  * Wrap errors and check behavior, not identity.

---

## #48: Not Leveraging `errors.Is` and `errors.As`

* **Problem**: Comparing errors with `==` instead of idiomatic helpers.
* **Correct**:

  ```go
  if errors.Is(err, io.EOF) { ... }
  if e := new(MyError); errors.As(err, &e) { ... }
  ```

---

## 📌 Key Takeaways from Chapter 6

* Go’s philosophy: **errors are values**, not exceptions.
* Always add context, wrap properly, and check errors explicitly.
* Reserve `panic` for unrecoverable situations.
* Use `errors.Is` and `errors.As` to build robust error handling.

---

✅ That completes **Chapter 6 – Error Management (Mistakes #41–#48)**.

Shall I continue with **Chapter 7 – Concurrency (Mistakes #49–#60)** next?


Awesome 🚀 Now we enter one of Go’s defining features — **concurrency**.
**Chapter 7 – Concurrency** covers **Mistakes #49–#60**, and it’s where many developers struggle, because goroutines and channels are deceptively simple to use but very easy to misuse.

---

# 📖 Chapter 7 – Concurrency

Go makes concurrency cheap and accessible (`go func()` launches a goroutine), but correctness requires discipline.
This chapter catalogues the **most frequent pitfalls** when working with goroutines, channels, sync primitives, and contexts.

---

## #49: Misunderstanding Goroutine Lifetimes

* **Problem**: Launching goroutines without knowing when they should stop.
* **Impact**: Zombie goroutines leak memory and CPU.
* **Fix**: Always tie goroutines to a **lifetime** (context, channel, or parent).

---

## #50: Memory Leaks from Goroutines

* **Problem**: Goroutines block forever on channels or waits.
* **Example**:

  ```go
  ch := make(chan int)
  go func() { ch <- 1 }() // if nobody reads, goroutine leaks
  ```
* **Fix**: Ensure goroutines have **exit paths**. Use buffered channels or `select` with `ctx.Done()`.

---

## #51: Deadlocks with Channels

* **Problem**: Cyclic dependencies where goroutines wait on each other.
* **Example**:

  * Goroutine A waits on B
  * Goroutine B waits on A
* **Avoid**: Keep channel communication **directional and structured**.

---

## #52: Incorrect Channel Buffering

* **Problem**: Choosing the wrong buffer size (0 vs. >0).
* **Impact**: Blocking when you don’t expect it, or unbounded growth with large buffers.
* **Guideline**: Start with unbuffered channels, add buffering only when necessary.

---

## #53: Closing Channels Incorrectly

* **Rule**: Only the **sender** should close a channel.
* **Mistake**: Closing from the receiver side, or writing after close → panic.
* **Guideline**: Close channels to signal **no more values**, not to indicate errors.

---

## #54: Race Conditions on Shared Memory

* **Problem**: Multiple goroutines accessing shared data without synchronization.
* **Impact**: Non-deterministic bugs, data corruption.
* **Tool**: Run `go test -race` to detect.
* **Fix**: Use `sync.Mutex`, `sync.RWMutex`, or channels for safe communication.

---

## #55: Misusing Sync Primitives

* **Examples**:

  * Forgetting `wg.Done()` in `sync.WaitGroup`.
  * Locking a `sync.Mutex` twice.
  * Deferring `Unlock` too late, creating deadlocks.
* **Guideline**: Keep synchronization minimal and consistent.

---

## #56: Worker Pool Mismanagement

* **Problem**: Spawning too many goroutines or not controlling pool size.
* **Fix**: Use bounded worker pools with buffered jobs channel and `sync.WaitGroup`.

---

## #57: Misunderstanding Context Cancellation

* **Problem**: Ignoring `context.Context`, or not propagating it.
* **Impact**: Goroutines continue after request ends.
* **Best practice**: Always pass `ctx` down call chains, check `ctx.Done()` in long-running tasks.

---

## #58: Leaking Tickers and Timers

* **Problem**: `time.NewTicker` and `time.After` keep resources alive if not stopped.
* **Fix**: Always `defer ticker.Stop()` or `timer.Stop()` when done.

---

## #59: Blocking on Unconsumed Channels

* **Problem**: Sending values to a channel that no one reads.
* **Impact**: Deadlocks or goroutine leaks.
* **Fix**: Ensure readers exist, or use buffered channels appropriately.

---

## #60: Not Handling Select Timeouts

* **Problem**: Blocking forever in a `select` without a timeout case.
* **Fix**: Always consider a timeout, e.g.:

  ```go
  select {
  case val := <-ch:
      fmt.Println(val)
  case <-time.After(time.Second):
      fmt.Println("timeout")
  }
  ```

---

## 📌 Key Takeaways from Chapter 7

* Goroutines are **cheap but not free**—always manage their lifetimes.
* Channels should be used deliberately: ownership, buffering, and closing rules matter.
* Race conditions and deadlocks are the silent killers of Go concurrency.
* `context.Context` is the idiomatic way to propagate cancellation.
* Tickers, timers, and channels require **explicit cleanup**.

---

✅ That wraps **Chapter 7 – Concurrency (Mistakes #49–#60)**.

Would you like me to continue with **Chapter 8 – Standard Library and APIs (Mistakes #61–#66)** next?


Perfect 👌 Now we continue with **Chapter 8 – Standard Library and APIs**, which covers **Mistakes #61–#66**.

This chapter shows how **misusing Go’s standard library** leads to bugs, inefficiencies, or subtle errors. The standard library is powerful, but you must understand its design decisions.

---

# 📖 Chapter 8 – Standard Library and APIs

Go’s stdlib covers everything from **time management** to **I/O and networking**. Misunderstandings often arise because APIs look simple but hide complex resource behaviors.

---

## #61: Misusing the `time` API

* **Problem**: Treating `time.Time` and `time.Duration` carelessly.
* **Common mistakes**:

  * Confusing **nanoseconds** vs. **seconds** in durations.
  * Using `time.After` repeatedly → memory leaks.
  * Comparing `time.Time` incorrectly (`==` instead of `.Equal`).
* **Fixes**:

  * Always check units (use `time.Second`, `time.Millisecond`).
  * Use `time.Since(start)` or `time.Until(deadline)` idioms.
  * Release timers/tickers with `Stop()`.

---

## #62: Incorrect Use of `os.File`

* **Problem**: Forgetting to close files, leading to descriptor leaks.
* **Pattern**:

  ```go
  f, err := os.Open("file.txt")
  if err != nil { return err }
  defer f.Close()
  ```
* **Extra mistake**: Deferring inside loops → too many open files.
* **Fix**: Close explicitly in loops, only defer for single open/close.

---

## #63: Forgetting to Close Resources (files, DB rows, etc.)

* **Resources**: `sql.Rows`, HTTP bodies, etc.
* **Mistake**: Assuming GC will clean up.
* **Impact**: Resource exhaustion under load.
* **Best practice**: Always `defer Close()` immediately after checking errors.

---

## #64: Misusing `http.Client`

* **Problem**: Creating a new `http.Client` for every request.
* **Impact**: Exhausts sockets, ignores connection pooling.
* **Fix**: Reuse a single `http.Client` instance across requests.
* **Extra**: Always close response bodies:

  ```go
  resp, err := client.Get(url)
  if err != nil { ... }
  defer resp.Body.Close()
  ```

---

## #65: Wrong JSON Struct Tags

* **Problem**: Forgetting `json:"fieldName"` leads to incorrect marshaling/unmarshaling.
* **Example**:

  ```go
  type User struct {
      Name string `json:"name"`
  }
  ```
* **Impact**: API contracts break silently.
* **Tip**: Use tools like `jsonlint` or `encoding/json` tests to verify.

---

## #66: Ignoring Security Defaults in Crypto and TLS

* **Problem**: Overriding defaults (e.g., disabling certificate checks, using weak ciphers).
* **Impact**: Security vulnerabilities.
* **Guideline**: Trust Go’s strong defaults unless you deeply understand crypto.
* **Bad**:

  ```go
  &tls.Config{InsecureSkipVerify: true} // NEVER
  ```
* **Good**: Use defaults or configure explicitly with up-to-date practices.

---

## 📌 Key Takeaways from Chapter 8

* Go’s stdlib is designed to be **safe and efficient by default**—problems come from overriding or ignoring idioms.
* Always **close resources** (files, DB connections, HTTP bodies).
* Reuse `http.Client`, avoid leaking timers and tickers.
* Respect security defaults: they exist for a reason.

---

✅ That’s **Chapter 8 – Standard Library and APIs (Mistakes #61–#66)**.

Would you like me to continue with **Chapter 9 – Dependency Management (Mistakes #67–#70)** next?


Great 👍 Let’s continue with **Chapter 9 – Dependency Management**, which covers **Mistakes #67–#70**.

This is a shorter chapter but very important: Go modules changed how dependencies are handled, and many projects still suffer from mismanagement.

---

# 📖 Chapter 9 – Dependency Management

Go modules (`go.mod`, `go.sum`) are the standard way to handle dependencies since Go 1.11+. Mistakes usually come from misunderstanding versioning, vendoring, or over-relying on external libraries.

---

## #67: Not Using Go Modules Properly

* **Problem**: Sticking to GOPATH or mixing vendoring with modules.
* **Impact**: Broken builds, inconsistent environments.
* **Fix**: Always use Go modules (`go mod init`, `go mod tidy`). Keep `go.mod` and `go.sum` committed to version control.

---

## #68: Versioning Pitfalls

* **Problem**: Forgetting semantic import versioning (e.g., `github.com/pkg/v2/...`).
* **Impact**: Wrong versions pulled silently.
* **Fix**: Respect Go’s semantic import rules. Use tags and `go get` properly.
* **Tip**: Pin exact versions in `go.mod` for reproducibility.

---

## #69: Vendoring Confusion

* **Problem**: Copying dependencies into `vendor/` without understanding how Go resolves them.
* **Impact**: Hidden mismatches between vendored code and `go.mod`.
* **Fix**: Use `go mod vendor` only when absolutely needed (e.g., corporate firewalls). Otherwise, rely on module proxy.

---

## #70: Over-Relying on External Libraries

* **Problem**: Importing libraries for trivial features (e.g., UUIDs, string helpers).
* **Impact**: Extra attack surface, harder upgrades, bloated code.
* **Guideline**: Prefer stdlib and small, well-maintained libraries. Add external dependencies only when they solve non-trivial problems.

---

## 📌 Key Takeaways from Chapter 9

* Always adopt **Go modules**—they are the standard.
* Respect **semantic import versioning** rules.
* Use vendoring only when required.
* Limit dependencies: every external package introduces risk.

---

✅ That wraps up **Chapter 9 – Dependency Management (Mistakes #67–#70)**.

Shall I continue with **Chapter 10 – Testing (Mistakes #71–#80)** next?


Perfect 👌 Let’s continue with **Chapter 10 – Testing**, which covers **Mistakes #71–#80**.

Testing is first-class in Go (`testing` package, `go test` command). But even with great tooling, developers often fall into traps that make tests flaky, incomplete, or misleading.

---

# 📖 Chapter 10 – Testing

Go promotes **simple, table-driven, deterministic tests**. The mistakes in this chapter come from treating testing as an afterthought or misusing the built-in tools.

---

## #71: Writing Non-Deterministic Tests

* **Problem**: Tests that depend on timing, random seeds, or external services.
* **Impact**: Flaky tests that sometimes pass, sometimes fail.
* **Fix**:

  * Mock dependencies.
  * Control randomness with fixed seeds.
  * Avoid real network/file access in unit tests.

---

## #72: Misusing `testing.T` and `testing.B`

* **Mistakes**:

  * Ignoring `t.Helper()` in helper functions → wrong line numbers in failures.
  * Using `t.Fatal` in goroutines → test doesn’t fail properly.
  * Mismanaging benchmarks with `testing.B`.
* **Guideline**: Keep test helpers clear and benchmarks isolated.

---

## #73: Not Cleaning Test Dependencies

* **Problem**: Tests leaving temp files, DB rows, or goroutines running.
* **Impact**: Interference between tests.
* **Fix**: Always clean up with `defer` or `t.Cleanup()`.

---

## #74: Testing Private Details Instead of Behavior

* **Problem**: Over-testing internal functions instead of exposed behavior.
* **Impact**: Fragile tests that break on refactors.
* **Guideline**: Test public APIs; private details may be tested indirectly through exported functionality.

---

## #75: Over-Mocking

* **Problem**: Mocking too much, making tests reflect mocks instead of reality.
* **Impact**: Passing tests but broken production code.
* **Fix**: Use mocks only where external dependencies can’t be controlled (e.g., network). Favor integration tests where possible.

---

## #76: Skipping Edge Cases

* **Problem**: Testing only “happy paths.”
* **Impact**: Bugs surface in production under error conditions.
* **Fix**: Always test with:

  * Empty slices/maps.
  * Zero values.
  * Boundary values.
  * Errors from dependencies.

---

## #77: Not Using Subtests

* **Problem**: Writing repetitive tests instead of structured subtests.
* **Better**:

  ```go
  for _, tc := range cases {
      t.Run(tc.name, func(t *testing.T) {
          got := MyFunc(tc.input)
          if got != tc.want { t.Errorf(...) }
      })
  }
  ```

---

## #78: Test Flakiness from Timing Issues

* **Problem**: Using `time.Sleep` in tests.
* **Impact**: Slow, flaky tests.
* **Fix**: Synchronize with channels, conditions, or mocks instead of sleeping.

---

## #79: Ignoring Code Coverage

* **Problem**: Not measuring test coverage → blind spots.
* **Fix**: Run `go test -cover` and use tools like `go tool cover -html=cover.out`.
* **Note**: Aim for meaningful coverage, not just a high percentage.

---

## #80: Poor Benchmarking Practices

* **Mistakes**:

  * Forgetting to reset timer in benchmarks (`b.ResetTimer()`).
  * Allocating inside loops instead of outside.
  * Ignoring compiler optimizations (dead code elimination).
* **Fix**: Benchmark critical code paths realistically.

---

## 📌 Key Takeaways from Chapter 10

* Go testing is powerful, but **discipline is required**: avoid randomness, timing issues, and unclean states.
* Focus tests on **behavior, not implementation details**.
* Use subtests and table-driven style for clarity.
* Coverage and benchmarks are part of healthy test suites.

---

✅ That’s **Chapter 10 – Testing (Mistakes #71–#80)**.

Would you like me to continue with **Chapter 11 – Tooling (Mistakes #81–#90)** next?


Great 🙌 Now let’s move to **Chapter 11 – Tooling**, which covers **Mistakes #81–#90**.

Go’s tooling (`go fmt`, `go vet`, linters, profiling tools, etc.) is one of the strongest parts of the ecosystem. This chapter shows how neglecting them hurts productivity and code quality.

---

# 📖 Chapter 11 – Tooling

Go ships with a built-in toolchain for formatting, analysis, and performance. Misusing or ignoring these tools often leads to inconsistent, buggy, or inefficient code.

---

## #81: Not Formatting Code with `go fmt`

* **Problem**: Writing unformatted code or enforcing style manually.
* **Fix**: Always run `go fmt` (or let your IDE do it).
* **Benefit**: Eliminates debates on style; all Go code looks consistent.

---

## #82: Ignoring `go vet`

* **Problem**: Skipping static analysis warnings.
* **Impact**: Missed bugs (e.g., misused `Printf` verbs, unreachable code).
* **Fix**: Run `go vet ./...` regularly; integrate it into CI pipelines.

---

## #83: Not Using Linters

* **Problem**: Relying only on `go vet`.
* **Better**: Use `golangci-lint` (aggregates many linters).
* **Value**: Catches stylistic issues, complexity, dead code, unsafe constructs.

---

## #84: Misusing `go build` and `go install`

* **Mistake**: Confusing `go build` (builds binary locally) and `go install` (installs binary into `$GOBIN`).
* **Impact**: Deployment confusion.
* **Fix**: Use `go build` for testing binaries, `go install` for reusable commands.

---

## #85: Ignoring `go mod tidy`

* **Problem**: Accumulating unused dependencies in `go.mod`.
* **Fix**: Run `go mod tidy` to clean.
* **Benefit**: Keeps dependency graph minimal and clean.

---

## #86: Misusing `go generate`

* **Problem**: Treating `go generate` as a build tool.
* **Correct**: It’s a **convenience hook** for code generation, not part of compilation.
* **Best practice**: Document required generators in comments for reproducibility.

---

## #87: Not Profiling with `pprof`

* **Problem**: Optimizing blindly.
* **Fix**: Use `net/http/pprof` or `go test -bench -cpuprofile` to analyze.
* **Outcome**: Real insights into bottlenecks.

---

## #88: Not Using Race Detector

* **Problem**: Skipping `go test -race`.
* **Impact**: Concurrency bugs remain hidden.
* **Fix**: Run it in CI pipelines, especially for concurrent code.

---

## #89: Poor Logging Practices

* **Mistake**: Using `fmt.Println` for logging everywhere.
* **Impact**: No levels, no timestamps, hard to debug.
* **Fix**: Use `log` package or structured logging (`zap`, `zerolog`).

---

## #90: Ignoring Tooling Automation

* **Problem**: Running tools manually instead of CI/CD integration.
* **Fix**: Automate `fmt`, `vet`, linters, and tests in pipelines.
* **Benefit**: Consistent enforcement across teams.

---

## 📌 Key Takeaways from Chapter 11

* Tooling is **integrated into Go’s DNA**—not using it is wasting free productivity.
* `fmt`, `vet`, linters, and race detectors catch problems early.
* Profiling with `pprof` helps optimize based on data, not guesses.
* Automation ensures consistency across teams and environments.

---

✅ That’s **Chapter 11 – Tooling (Mistakes #81–#90)**.

Would you like me to continue with the **final chapter: Chapter 12 – Optimizations (Mistakes #91–#100)** next?


Perfect 🚀 Let’s wrap up with **Chapter 12 – Optimizations**, which covers the last set: **Mistakes #91–#100**.

Optimization is where many Go developers trip: either **premature optimization** (wasting effort) or **missing critical performance issues** (allocations, false sharing, bad memory patterns).

---

# 📖 Chapter 12 – Optimizations

Go is fast by default, but performance-sensitive applications demand care. This chapter explains common pitfalls that affect CPU, memory, and scalability.

---

## #91: Premature Optimization

* **Problem**: Micro-optimizing before measuring.
* **Impact**: Wasted time, unreadable code.
* **Rule**: First, write correct and idiomatic code → then measure → then optimize.

---

## #92: Unnecessary Memory Allocations

* **Problem**: Creating new slices/maps inside loops.
* **Example**:

  ```go
  for i := 0; i < n; i++ {
      s := make([]int, 0) // allocates every iteration
  }
  ```
* **Fix**: Pre-allocate (`make([]int, 0, n)`) or reuse buffers.

---

## #93: Copying Data Unnecessarily

* **Problem**: Passing large slices or structs by value.
* **Impact**: More GC pressure, cache misses.
* **Fix**: Pass slices by reference (already references), use pointers where beneficial.

---

## #94: Poor Slice Management

* **Problem**: Growing slices inefficiently (`append` in hot loops).
* **Fix**: Pre-size slices when possible.
* **Tip**: Use `copy` for merging slices efficiently.

---

## #95: Ignoring Escape Analysis

* **Problem**: Variables unexpectedly allocated on the heap instead of the stack.
* **Impact**: More GC overhead.
* **Tool**: `go build -gcflags="-m"` shows escape analysis.
* **Fix**: Adjust pointer usage to keep objects stack-allocated.

---

## #96: Not Using Sync Pools for Reusable Objects

* **Problem**: Repeatedly allocating expensive objects (buffers, structs).
* **Fix**: Use `sync.Pool` to recycle.
* **Best use**: High-throughput systems with short-lived objects.

---

## #97: False Sharing Between Goroutines

* **Problem**: Two goroutines modifying fields on the same cache line.
* **Impact**: CPU cache contention → slower performance.
* **Fix**: Use padding (`[64]byte`) or struct alignment to separate hot fields.

---

## #98: Inefficient String and Byte Conversions

* **Problem**: Repeated `[]byte(string)` and `string([]byte)` conversions.
* **Impact**: Allocations.
* **Fix**: Minimize conversions; keep data in the right type until needed.

---

## #99: Not Leveraging Inlining and Compiler Optimizations

* **Problem**: Writing code that prevents inlining (too complex functions).
* **Fix**: Keep hot-path functions small; let Go’s compiler inline automatically.
* **Tool**: `go build -gcflags="-m"` to see inlining decisions.

---

## #100: Neglecting GC and Profiling in Production

* **Problem**: Optimizing locally but ignoring live workload behavior.
* **Fix**: Always profile in production-like environments with `pprof`.
* **Watch**: GC pauses, heap growth, goroutine count.

---

## 📌 Key Takeaways from Chapter 12

* **Don’t guess—measure.** Profiling is the foundation of optimization.
* Avoid unnecessary allocations, copying, and conversions.
* Use Go’s memory model (escape analysis, sync.Pool) to your advantage.
* Cache-friendly data structures and pre-sizing slices matter at scale.
* Optimizations must always balance **performance and readability**.

---

✅ That’s the **final chapter: Chapter 12 – Optimizations (Mistakes #91–#100)**.

We’ve now completed the **entire course**: all 12 chapters, all 100 mistakes, explained chapter by chapter. 🎉

Would you like me to now generate a **condensed executive summary of all 100 mistakes** (like a cheat sheet / quick reference) so you can keep it as a compact study guide?


Perfect 🙌 Here’s a **condensed executive summary / cheat sheet** of all **100 Go mistakes**, organized by chapter.
This is the **quick reference version**: short, actionable, and easy to revisit when coding.

---

# 📑 Executive Summary — 100 Go Mistakes

---

## **Chapter 1 – Go: Simple to Learn but Hard to Master**

* Go is simple, but idioms are subtle. Mistakes often come from misunderstanding concurrency, interfaces, and error handling.

---

## **Chapter 2 – Code & Project Organization (Mistakes #1–16)**

1. Mixing `init` misuse.
2. Ignoring package boundaries.
3. Overusing globals.
4. Poor naming.
5. Cyclic dependencies.
6. Misusing `internal` packages.
7. Interfaces in wrong packages.
8. Overusing interfaces.
9. Returning interfaces instead of structs.
10. Inconsistent visibility (`exported` vs private).
11. Using getters/setters unnecessarily.
12. Nesting deeply instead of composition.
13. Not separating domain vs infrastructure.
14. Overcomplicated project structure.
15. Mixing business and transport layers.
16. Ignoring Go idioms for package layout.

---

## **Chapter 3 – Data Types (Mistakes #17–30)**

17. Zero values misunderstood.
18. Nil slices/maps confusion.
19. Over-allocating slices.
20. Misusing arrays.
21. Forgetting slice capacity.
22. Sharing underlying arrays unexpectedly.
23. Appending incorrectly.
24. Slice vs array parameter confusion.
25. Nil vs empty slice differences.
26. Map iteration order assumption.
27. Reading from nil maps.
28. Writing to nil maps.
29. Using pointers to slices/maps unnecessarily.
30. Using value receivers with large structs.

---

## **Chapter 4 – Strings (Mistakes #31–35)**

31. Inefficient string concatenation (`+` in loops).
32. Misusing `strings.Builder`.
33. Confusing bytes vs runes.
34. Wrong use of `Trim` vs `TrimSuffix/Prefix`.
35. Ignoring string immutability performance costs.

---

## **Chapter 5 – Functions & Methods (Mistakes #36–40)**

36. Passing large structs by value.
37. Returning pointers unnecessarily.
38. Ignoring multiple return values.
39. Misusing variadic functions.
40. Inconsistent value vs pointer receivers.

---

## **Chapter 6 – Error Management (Mistakes #41–48)**

41. Ignoring errors (`_, _ :=`).
42. Overusing custom error types.
43. Not wrapping errors with `%w`.
44. Returning errors without context.
45. Misusing `panic/recover`.
46. Returning typed nil interfaces.
47. Overusing sentinel errors.
48. Not using `errors.Is` / `errors.As`.

---

## **Chapter 7 – Concurrency (Mistakes #49–60)**

49. Not managing goroutine lifetimes.
50. Goroutine leaks.
51. Deadlocks on channels.
52. Wrong channel buffering.
53. Closing channels incorrectly.
54. Race conditions.
55. Misusing sync primitives (`Mutex`, `WaitGroup`).
56. Worker pool mismanagement.
57. Ignoring `context.Context`.
58. Leaking timers/tickers.
59. Sending to unconsumed channels.
60. Not handling select timeouts.

---

## **Chapter 8 – Standard Library & APIs (Mistakes #61–66)**

61. Misusing `time` API (units, comparison, leaks).
62. Forgetting to close files.
63. Not closing DB rows/HTTP bodies.
64. Re-creating `http.Client`.
65. Wrong JSON struct tags.
66. Ignoring TLS/crypto defaults.

---

## **Chapter 9 – Dependency Management (Mistakes #67–70)**

67. Not using Go modules properly.
68. Ignoring semantic versioning.
69. Vendoring confusion.
70. Over-relying on external libraries.

---

## **Chapter 10 – Testing (Mistakes #71–80)**

71. Non-deterministic tests.
72. Misusing `testing.T/B`.
73. Not cleaning up tests.
74. Testing private details.
75. Over-mocking.
76. Skipping edge cases.
77. Not using subtests.
78. Flaky tests with `time.Sleep`.
79. Ignoring coverage.
80. Poor benchmarking.

---

## **Chapter 11 – Tooling (Mistakes #81–90)**

81. Not running `go fmt`.
82. Ignoring `go vet`.
83. Not using linters (`golangci-lint`).
84. Misusing `go build` vs `go install`.
85. Ignoring `go mod tidy`.
86. Misusing `go generate`.
87. Not profiling with `pprof`.
88. Skipping race detector.
89. Poor logging (`fmt.Println` instead of logger).
90. No automation in CI/CD.

---

## **Chapter 12 – Optimizations (Mistakes #91–100)**

91. Premature optimization.
92. Extra allocations in loops.
93. Copying data unnecessarily.
94. Inefficient slice growth.
95. Ignoring escape analysis.
96. Not using `sync.Pool` where useful.
97. False sharing between goroutines.
98. Repeated string/byte conversions.
99. Blocking inlining.
100. Not profiling GC/production workloads.

---

# 📝 Final Takeaways

* **Idiomatic Go matters**: mistakes aren’t just bugs, they reduce readability, maintainability, and performance.
* **Concurrency & error handling** are the trickiest areas.
* **Tooling & profiling** are your safety nets—use them always.
* **Optimization comes last**: measure first, optimize only what matters.

---

✅ You now have the **full course (chapter-by-chapter)** + this **executive cheat sheet** for fast recall.

Do you want me to also generate a **PDF/EPUB version of the cheat sheet** so you can keep it offline alongside the full course notes?
