# golang
Golang practice project

- This is a sample project to explore different topics of go programming language.

## Useful commands
```golang
    - go mod init           // Initializes a new module in the current directory
    - go run main.go        // compile and run the program in a single step, without the need to create an executable file
    - go help               // to find more commands (with info) / helps to read doc. Known as tools or toolset given by go
    - go env GOPATH         // to check current path where go is installed
    - go env                // go related env
    - GOOS="windows" go build // Build executable for windows/linux/Mac (darwin)
    
```

## Others
| new()                         | make() |
| ---                           | ---    |
| Allocate memory but no INIT   | Allocate memory and INIT |
| You will get a memory address | You will get a memory address |
| Zeroed storage                | Non-zeroed storage |

