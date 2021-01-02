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

## Concurrency
In a nutshell:
- `Go routines` are the basic unit of concurrnecy
- Anon functions can be used to kick off concurrent processes
  - Because the only way to start a goroutine is to put `go` in front of a function call, we often use anonymous functions when we want to start a goroutine.
      - Anonymous functions can be executed when they're declared!
      ```go
      go func () {
          // do stuff
      }()
      ```
      - Anon functions maintain access to the lexical scope they are defined in - all the vars that are avaailable at the point when you declare the anon func are also available in the body of the function.
- Channels help organize and control the communication between different processes allowing us to avoid a race condition bug
- Check for race conditions with `go test --race` **might not always work**

Sage Thoughts:
- Make it work (tests pass), make it right (refactor), make it fast (optimize)
- *Premature Optimization is the root of all evil!*
- In effect, one cannot make it fast, without first making it work and making it right.

Clarity on `chan struct{}` semantics
- In the event we don't care *what* is sent on a channel, just *that* something is sent, we opt to utilize these symantics
- This is advantageous because this won't allocate anything. AKA this is the smallest data type available form memory.

**ALWAYS MAKE CHAN**
- When you declare something with var, it initializes it to its "zero" value. For a string "", int 0, etc...
- For channels, the "zero" value is `nil`, and if you try to send to it with `<-` it will block forever because you cannot send to `nil` channels

Select
- Recall that one can wait for vals to be sent to a chan with `myVar := <-ch` (reciever expression). This is a *blocking* call, as you're waiting for a value.
- Select allows youto wait on *multiple* channels. The first one to send a value "wins and the code undernewath the `case` is executed.
- Sometimes you'll want to include `time.After` in one of your cases to prevent your system from blocking forever.

`httptest`
- This is a great way to create test servers so you can have reliable, controllable tests.
- Uses the same interfaces as the "real" `net/http` servers which is consistent and less for you to learn.
# What I don't get

`t.Helper()` designates the calling function a test helper. When pringting file and line info, that function will be skipped. Helper may be called simultaneously from multiple goroutines.

`t.Fatal()` is used to end a program so we don't panic if subsequent steps continue. Is there something like this in the standard go lib?

In the associated [article](https://blog.cleancoder.com/uncle-bob/2014/05/14/TheLittleMocker.html), I don't get what Uncle Bob means by "reflection" in the following context: `Ironically, some mocking systems depend strongly on reflection, and are therefore very slow.`