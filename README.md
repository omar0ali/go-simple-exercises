# Golang Math Library

A simple arithmetic library implemented in Go, providing basic mathematical operations such as addition, subtraction, multiplication, and division.

## Usage

To perform calculations, use the following syntax:

```golang
go run . operation <numbers>...
```

## Example usage

```golang
go run . + 100 20
120
go run . - 100 20
80
go run . x 2 2
4
```

### Supported Operations:

Addition: + or add
Subtraction: - or sub
Multiplication: x or multi
Division: / or division

### For binary search:

For binary search, you can either provide space-separated values directly or use xargs to read from a file:

```bash
# Direct values
go run . bSearch <target> <sorted_values>

# Using xargs with a file
cat list.txt | xargs go run . bSearch <target>
```

### Main Function

```golang
func main() {
    var total int = 0
    if len(os.Args) == 1 {
        fmt.Println("No args provided.\nUsage: app.go operation <numbers>...\napp.go add 1 2 3")
    } else {
        switch os.Args[1] {
        case "add", "+":
            total = math.Add(os.Args[2:])
        case "sub", "-":
            total = math.Sub(os.Args[2:])
        case "multi", "x":
            total = math.Multiply(os.Args[2:])
        case "division", "/":
            total = math.Division(os.Args[2:])
        case "bSearch":
            ls := os.Args[2:]
            target := math.ConvertToInt(ls[0])
            values := ls[2:]
            fmt.Printf("%d %v\n", target, values)
            index := 0
            total, index = math.BinarySearch(target, values)
            fmt.Printf("Found at Index: %d\n", index)
        default:
            fmt.Println("Usage: app.go operation <numbers> ...\napp.go add 1 2 3")
        }
    }
    fmt.Printf("%v\n", total)
}
```
