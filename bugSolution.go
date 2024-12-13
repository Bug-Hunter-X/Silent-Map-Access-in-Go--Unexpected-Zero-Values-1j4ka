```go
func main() {
    var m map[string]int
    if val, ok := m["foo"]; ok {
        fmt.Println(val) // This will only be printed if "foo" exists
    } else {
        fmt.Println("Key not found") // Handle the missing key gracefully
    }
}
```