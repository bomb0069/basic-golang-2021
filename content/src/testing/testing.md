# Create and Run Test

## Create Unit Test

ลักษณะของ File Test ใน Go

- อยู่ที่เดียวกับ Production Code
- ใช้ชื่อ Package เดียวกัน ตามด้วย _test
- Function Test ต้องขึ้นต้นด้วยคำว่า "Test" เสมอ
- ภาษา Go แบบ Pure ๆ ต้องเขียนเช็คเงื่อนไขเอง

ตย. demo_test.go

```go
package demo_test   // ใช้ชื่อ Package เดียวกับ Production Code ตามด้วย _test
                    // ปกติถ้าชื่อ package ไม่ตรงกับ go.mod ตัว Compile go ก็จะด่ายกเว้นเป็นกรณีของ test

import (
    "demo"          // ต้องทำการ Import จาก Production Code เพื่อจำลองว่าถูกเรียกจาก นอก Package ของ Production
    "testing"       
)

func TestHello(t *testing.T) {      // Function Test ขึ้นต้นด้วย Test เสมอ
                                    // Test Framework  จะส่ง testing.T มาให้เราใช้
    var r string                    // ประกาศตัวแปลของ Go ซึ่งแาจจะใช้ r := demo.Greeting() ได้
    r = demo.Greeting()
    if r != "Hello World" {             // ต้องตรวจสอบ Result เองว่า Expected กับ Actual เท่ากันไหม
        t.Errorf("Error with %v", r)    // ใช้ t ที่ได้มาจัด Format ของ Error เวลา Test ไม่ผ่าน
    }
}
```

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

## Reference

- [Go Testing Package](https://golang.org/pkg/testing/)
