# Fundamentals

Complete the online book https://quii.gitbook.io/learn-go-with-tests/go-fundamentals
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


# What I don't get

`t.Helper()` designates the calling function a test helper. When pringting file and line info, that function will be skipped. Helper may be called simultaneously from multiple goroutines.