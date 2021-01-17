# Day 1

## Initial Go Project

```cmd
$go mod init demo
```

จะได้ file ชื่อ go.mod มาใน folder

```go
module demo

go 1.15
```

## Go Simple Project

### Create demo.go

สร้างไฟล์ demo.go (ชื่อเดียวกับ module) ปกติคนที่เขียน go จะมีการจัดการ module และจะสร้าง file ที่ชื่อเดียวกับ module ขึ้นมา เพื่อเป็นจุดเริ่มต้นของ module นั้น ซึ่งไม่จำเป็นต้องมีก็ได้ โดย File นี้ปกติจะไม่ถูกใช้ในการรัน

```go
package demo    // กำหนดชื่อ package ให้เหมือนชื่อ module ที่ initial มา

func Greeting() string {  // ประกาศ Function ชื่อ Greeting โดยมี Return Type เป็น String
    return "Hello World"
}
```

### ข้อกำหนดเรื่อง Function Name

เนื่องจาก Golang มีข้อกำหนด เรื่อง Access Modifier ของแต่ละ Function โดยที่ Function ใดที่ต้องการให้ Access จากภายนอก Package ก็จะต้องขึ้นต้นด้วยตัวพิมพ์ใหญ่ เช่น ในกรณีนี้ Greeting สามารถเข้าถึง/เรียกใช้ได้จาก Package อื่น ๆ ถ้าเปลี่ยนเป็น greeting ก็จะเรียกใช้ไม่ได้

### Create main.go

สร้าง Folder cmd ขึ้นมาเพื่อรองรับการ execute จาก cmd และสร้าง file ชื่อว่า main.go ขึ้นมา

```go
package main    // package สำหรับ main.go จะเป็น main เสมอ

import "demo"   // ต้องการใช้ Func Greeting() จาก module demo จึงต้อง import package demo เข้ามา 

func main() {
    demo.Greeting()  // เรียกใช้ Greeting ที่อยู่ใน package demo
}
```

เนื่องจาก Package demo มีการกำหนด Function Greeting โดยที่มีตัวอักษรแรกเป็นตัวใหญ่ เลยสามารถเรียกใช้จาก Function main() ซึ่งอยู่ใน Package main  และเป็นคนละ Package  กับ Greeting ได้

ถ้าเราเปลี่ยน Greeting เป็น greeting ตอน Compile go ก็จะด่าเรา

### Simple Project Structure

```tree
demo
|- go.mod   // Package dependency for go project
|- demo.go  // Default go file for each module
|- cmd
    |- main.go   // runable file with main function
```

## Testing with Golang

### Create Unit Test

ลักษณะของ File Test ใน Golang

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

### การรัน Test

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

### การรัน Benchmark เพื่อเทียบประสิทธิภาพของFunction

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

## Backlog of Content

List ของ git log จากทีี่ลองทำตามใน Class

```
9263df9 [Added]: Alias TYPE
b2f5581 [Added]: Example of Error
0b17abc [Added]: Example of Error with fmt package
656fa8a [Added]: Array
7e2f6c0 [Added]: Another type of for loop
d56c0d5 [Added]: write info in for with log.Printf
bb19fa3 [Added]: Slice
9052fca [Added]: append top Slice
a2d51ef [Added]: append 2 into Slice from Array
e6dd00a [Added]: init slice and make a slice
116ec16 [Added]: slice address in memory via &
b8f609d [Added]: sort of slice
79b88c4 [Added]: map
908e63e [Added]: Recovery from panic
5f95111 [Added]: Error vs Defer
06dd6ce [Added]: Simple struct
0bbe1e9 [Added]: New Struct
ddee4aa [Added]: Struct with Struct Comnposition
d3f2ede [Added]: Method in Struct
37f04c5 [Added]: Method of Struct and Overwrite Method in Go (for embeded)
5b26d5c [Added]: Method of Struct with Pointer
820e592 [Added]: Pointer with & from new object
174a278 [Added]: Interface of All
e00f4ce [Added]: Interface of Primitive Type Int
6b22fcd [Added]: Interface of Stringer
f876801 [Removed]: Stringer Interface
8b92209 [Added]: JSON Example and Exercise for add and list user
32731ca [Reorganized]: Folder Structure and move all of workshop into day1
4f7d6f4 [Init]: New task module for homework
73ee837 [Moved]: input.json into execise for initial data file
cef04a1 [Added]: Initial Cli with method NumberOfMember and AddMember
648726e [Added]: feature getMember with id
f5534c7 [Refactor]: Test to compare object instead of primitive
cd06b4a [Added]: Reviews Session of Day2
a97828c [Added]: Dependency Test of Store in UserService
9a078c5 [Added]: go module and sub module in go project
2f93cf4 [Added]: Version of Library
b72de03 [Added]: Project Structure
545a459 [Added]: Section2 SubPackage to call to Section1 SubPackage
32bbc90 [Added]: additional Go Project Structure
b40e03c [Added]: Content for Building RESTFul APIs
8310af8 [Added]: อธิบายการทำงานของ Router และ Middleware
```
