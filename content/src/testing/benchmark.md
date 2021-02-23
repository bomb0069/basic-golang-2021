# Benchmarking

## การรัน Benchmark เพื่อเทียบประสิทธิภาพของFunction

เราสามารถเปรียบเทียบความเร็วในการทำงานของ App ได้ด้วยการเขียน Benchmark แล้วรัน

ตย. เขียน Benchmark ใน test file demo_test.go

```go
func BenchmarkHello(b *testing.B) {
    for i := 0; i < b.N; i++ {
        demo.Greeting()
    }
}

func BenchmarkHello2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        demo.Greeting2()
    }
}
```

แล้วรันผ่านคำสั่ง

```cmd
$go test -bench .
goos: darwin            // บอกว่าใช้  OS อะไร darwin = macOS
goarch: amd64           // บอก CPU
pkg: demo
BenchmarkHello-4        1000000000             0.347 ns/op      // บอกว่ารันใช้เวลาเท่าไหร่ต่อ 1 Operation
BenchmarkHello2-4       1000000000             0.366 ns/op
PASS
ok      demo    1.288s
```

## Reference

- [How to write benchmarks in Go](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)
- [มาใช้งาน Benchmark ในภาษา Go กัน](https://www.somkiat.cc/benchmark-in-golang/)
- [Package test - import "cmd/go/internal/test" and search this with GOMAXPROCS](https://golang.org/pkg/cmd/go/internal/test/)
- [Package runtime - import "runtime" and search this with GOMAXPROCS](https://golang.org/pkg/runtime/)
