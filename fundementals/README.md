# Fundamentals

Complete the online book https://quii.gitbook.io/learn-go-with-tests/
# WIL (What I Learned)

Benchmarks can be easily written
- go test -bench=.
- utilize B.n

Examples
- There is an actual notion of examples
- If what is written in `// Output` comment doesn't match output, Example fails

`range` lets you iterate over an array. Everytime it is called it returns 2 vals; index and val

`_` blank identifier

`make` creates a slice with a fixed starting capactiy

Arrays
- size encoded in its type. I.E. can't pass [4]int into a func that expects [5]int

Slices are variable size. Use `append` to push

Interfaces can be satisfied minimally by structs with many more parameters.
-> ties into `parametric polymorphism` declaring interfaces so you can define functions that can be used by different types

*POINTERS*
- struct pointers are automatically dereferenced
- Sometimes you won't need to accept a pointer in the struct method. However, it is convention to keep these arguments aligned. So if a method exists (like deposit) that requires a pointer, other struct methods should also include a pointer (i.e balance)

- go copies vals when you pass them to methods/functions, so if you want the function to mutate the object passed to it, pass with a pointer reference.

*NIL*
- Pointers can be nil
- When a function returns a pointer to something, you need to check if its nil or you might raise a runtime exception. The compiler won't help you here
- Useful when you want to describe a value that may be missing

There are many linters available for go. Some verify that you've checked for errs in the right place.
```
go get -u github.com/kisielk/errcheck
errcheck .
```

Useful for adding more domain specific meaning to values
Can let you implement interfaces

Reference Types - An interesting property of maps is that you can modify them without passing them as a pointer. This is because `map` is a reference ype, meaing it holds a reference to the undrelying data structure. The underlying structure is a `hash table` Maps being a reference is good b/c no matter how big they get there will only be one copy.

## Errors
- can be made into constants and passed around (see `maps`)
- can be generated with a wrapper that implements the error interface

**Use Mocks (and tests in general) to test *behavior*, rather than implementation! -> Think in terms of how it behaves!**

# What I don't get

`t.Helper()` designates the calling function a test helper. When pringting file and line info, that function will be skipped. Helper may be called simultaneously from multiple goroutines.

`t.Fatal()` is used to end a program so we don't panic if subsequent steps continue. Is there something like this in the standard go lib?

In the associated [article](https://blog.cleancoder.com/uncle-bob/2014/05/14/TheLittleMocker.html), I don't get what Uncle Bob means by "reflection" in the following context: `Ironically, some mocking systems depend strongly on reflection, and are therefore very slow.`