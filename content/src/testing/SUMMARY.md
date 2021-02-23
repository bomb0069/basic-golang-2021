# Testing in Golang

- [Testing](./testing.md)
- [Naming and Convention](./naming.md)
- [Benchmarking](./benchmark.md)
- [Profiling](./profile.md)
- [Sub Testing and Sub Benchmarking](./sub.md)

## การรัน Test

การรัน Test ใช้ Command

```cmd
$go test    // มันจะหาไฟล์ Testให้
```

ต้องการรัน Test ทั้งหมด ใน Sub-Folder

```cmd
$go test ./...

ok    demo    (cached)              // ขึ้นคำว่า cached แปลว่า เคยรันแล้ว จะรันใหม่ก็ต่อเมื่อ Production Code เปลี่ยน
?     demo/cmd   [no test files]    // แจ้งว่าไม่มี File test
```

เราสามารถเช็ค Coverage ของเทสที่เราเขียนได้โดยการรัน

```cmd
$go test -cover

PASS
coverage: 10.0% of statements
ok    demo    0.339s
```

หรือเราสามารถให้เก็บเป็น File ได้

```cmd
$go test -coverprofile=coverage.out

PASS
coverage: 10.0% of statements
ok      demo    0.228s
```

ซึ่งจะได้ผลลัพท์เก็บใน File coverage.out

```text
mode: set
demo/demo.go:3.24,5.2 1 1
```

หรือถ้าเราอยากเห็น Report สวยงาม แถมบอกด้วยว่าบรรทัดไหนผ่านการเทสไหม ก็สามารถใช้

```cmd
go tool cover -html=coverage.out
```

มันจะเปิด Browser ขึ้นมาพร้อม coverage report เลย
