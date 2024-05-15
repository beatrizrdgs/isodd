# Is Odd

`IsOdd` is a GoLang package that provides a simple function to determine whether a given integer is odd or not.



## Installation

To install `IsOdd`, you can use go get:

```bash
go get -u github.com/beatrizrdgs/isodd
```



## Usage

Simply invoke the `IsOdd` function to check if the integer is odd.

```go
package main

import (
    "fmt"
    "github.com/beatrizrdgs/isodd"
)

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    fmt.Println("Checking if numbers are odd:")
    for _, number := range numbers {
        if isodd.IsOdd(number) {
            fmt.Printf("%d is odd\n", number)
        } else {
            fmt.Printf("%d is not odd\n", number)
        }
    }
}
```



## Documentation

The package contains only one function:

### func [IsOdd](isodd.go)
`func IsOdd(n int) bool`

IsOdd takes an integer n as input and returns true if the number is odd, otherwise it returns false.



## Dependencies

This package depends on the `IsEven` function from the [iseven](https://github.com/beatrizrdgs/iseven) package.



## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.