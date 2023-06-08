# Welcome to GoPointers!
In this repository, I explore the use of **pointers** in Go programming language.

# Prerequisites
- You should have the standard go compiler and toolchain installed.
- You should have access to the *go* command

# Usage
- First, clone the repository to your local machine:
```go
    $ git clone https://github.com/DreamLineLove/gopointers.git
```

- Navigate to */gopointers*

- Use this command to run the application with default setings:
```go
    $ go run ./cmd/pointers/
```

- Use this command to get command line flag options:
```go
    $ go run ./cmd/pointers/ -help
```

## Flags
#### -basicsStrings
sets the value of the basicsString variable in basics() function
```go
    $ go run ./cmd/pointers/ -basicsString=something
```

- *-functionsString* sets the value of the functionsString variable in functions() function
```go
    $ go run ./cmd/pointers/ -functionsString=something
```

- *-methodsString* sets the value of the methodsString variable in methods() function
```go
    $ go run ./cmd/pointers/ -methodsString=something
```

