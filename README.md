# Welcome to GoPointers!
<img src="https://camo.githubusercontent.com/94761affed6454156a526a0fcab454ed4a432d9472087a9d330598a38ffe56cd/68747470733a2f2f7261772e6769746875622e636f6d2f676f6c616e672d73616d706c65732f676f706865722d766563746f722f6d61737465722f676f706865722e706e67" width="175px" />

In this repository, I explore the use of **pointers** in the Go programming language.

# Requirements
- You should have the standard go compiler and toolchain <a href="https://go.dev/learn/" target="_blank">installed</a>
- You should have access to the *go* command
- You should have set up *GOPATH* and *GOROOT* environment variables

# Usage
- First, clone the repository to your local machine:
```
    $ git clone https://github.com/DreamLineLove/gopointers.git
```

- Navigate to ```$ cd gopointers```

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

