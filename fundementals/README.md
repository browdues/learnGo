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

# What I don't get

`t.Helper()` designates the calling function a test helper. When pringting file and line info, that function will be skipped. Helper may be called simultaneously from multiple goroutines.