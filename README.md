# Executtor

> A tiny execution wrapper for programs that run safely and shutdown gracefully (if possible)

### Install Executtor

```bash
$ go get github.com/drgomesp/executtor`
```

### Import Executtor

```go
import x "github.com/drgomesp/execcutor"
```

### Basic Example

```go
func main() {
    // Run returns an error which you can capture and handle
    _ = x.Run(func(ctx context.Context, args ...string) error {
        // your code goes here
        return nil
    })
}
```

### More examples

See [examples/](https://github.com/drgomesp/execcutor/blob/master/examples/main.go)

### License

[MIT](https://github.com/drgomesp/execcutor/blob/master/LICENSE)

---

Built by [Daniel Ribeiro](https://github.com/drgomesp).
