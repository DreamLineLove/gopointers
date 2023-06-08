# Welcome to GoPointers!
<img src="https://www.golinuxcloud.com/wp-content/uploads/goher2-1-218x300.jpg" width="150px" />
In this repository, I explore the use of **pointers** in the Go programming language.

# Requirements
- You should have the standard go compiler and toolchain installed.
- You should have access to the *go* command

# Usage
- First, clone the repository to your local machine:
```go
    $ git clone https://github.com/DreamLineLove/gopointers.git
```

- Navigate to ```cd gopointers/```

- Use this command to run with default setings:
```go
    $ go run ./cmd/pointers/
```

- Use this command to get a list of available command line flag options:
```go
    $ go run ./cmd/pointers/ -help
```

## Flags
### -basicsString
Sets the value of the basicsString variable in basics() function
```go
    $ go run ./cmd/pointers/ -basicsString=something
```

### -functionsString
Sets the value of the functionsString variable in functions() function
```go
    $ go run ./cmd/pointers/ -functionsString=something
```

### -methodsString
Sets the value of the methodsString variable in methods() function
```go
    $ go run ./cmd/pointers/ -methodsString=something
```

