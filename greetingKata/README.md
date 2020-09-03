# Greeting Kata
- Write a method greet that accepts a string input and returns a string output. greet("Bob") should return "Hello, Bob."
- Handle null or empty names with "Hello, my friend."
- When the name is in all caps shout back at them with "HELLO BOB!"
- Oh No, our greet function was working perfectly according to the specification and now Product wants us to greet groups of people as well. Modify greet to accept an array of one or two names. when it is two use and to separate them.  "Hello, Jill and Jane."
- Now we need to handle arbitrary sized list of names. For lists of size > 2 use a coma delimited list "Hello, Amy, Brian, and Charlotte."
- When an individual name in the list is in all caps they get special emphasis and are placed at the end. ["Amy", "BRIAN", "Bob", "MAX"] should return "Hello, Amy and Bob. AND HELLO BRIAN AND MAX!"

## What I learned

- The last test made my simple implementation turn ugly faster than plastic surgery on MJ
- [A string can't be nil in go](https://stackoverflow.com/questions/36001477/why-cant-a-string-be-nil-in-go#36006313)
- [You can't overload functions in go. You can work around this with variadic functions.](https://golangbyexample.com/function-method-overloading-golang/)
- There is a thing in the universe of computer science known as a variadic function
- If I use `watch`, I can watch my tests. E.g. `watch go test -v`

## What I still don't understand

- `go test -cover` tells me code coverage is ~96%. How do I tell what staements are NOT covered?