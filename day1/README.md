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

เนื่องจาก Go มีข้อกำหนด เรื่อง Access Modifier ของแต่ละ Function โดยที่ Function ใดที่ต้องการให้ Access จากภายนอก Package ก็จะต้องขึ้นต้นด้วยตัวพิมพ์ใหญ่ เช่น ในกรณีนี้ Greeting สามารถเข้าถึง/เรียกใช้ได้จาก Package อื่น ๆ ถ้าเปลี่ยนเป็น greeting ก็จะเรียกใช้ไม่ได้

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

## Testing with Go

### Create Unit Test

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

## Go Type

### Alias TYPE

เราสามารถสร้าง Type ของตัวแปลใน Go ใหม่ได้ผ่าน  คำสั่ง type

```go
    type MONTH int  // สร้าง type ใหม่ที่ชื่อว่า MONTH โดยอ้างอิงจาก int

    var january MONTH // สร้างตัวแปลชื่อ january เป็นประเภท MONTH
```

### Example of Error

เนื่องจาก Go เป็นภาษาที่ return type ของ function สามารถทำได้มากกว่า 1 ตัวแปร เช่น เราอาจจะอยากสร้าง function ที่ return int และ string ไปพร้อมกัน ก็ได้ และสิ่งที่ Go ไม่มีคือ exception เพราะฉะนั้น ถ้า function ของเราต้องรองรับกรณีที่ทำได้สำเร็จจะ return แบบนึง ถ้า error จะ return แบบนึงก็สามารถที่จะสร้าง functiom ที่ return หลาย ๆ ตัวแปลได้

```go
package demo

import (
    "errors"
    "fmt"
)

// เป็น function ที่รับ int และ return int กับ error
func someShort(x int) (int, error) {

    err := errors.New("Normal error")

    // ถ้าเราเขียนแบบ short return เวลาจะส่งค่ากลับก็ใช้ return ตามด้วยลำดับของตัวแปลลตามรูปแบบที่เรากำหนดด้านบน 
    return 0, err
}

// เป็น function ที่ return int กับ error เหมือนกัน 
// ต่างกันตรงที่กำหนดชื่อตัวแปรไว้เลยว่า
// ค่าที่เป็น int จะ return ผ่านตัวแปร result และ error return ผ่าน err
func some(x int) (result int, err error) {
    
    // เราสามารถ set ค่าให้ตัวแปร err ได้
    err = errors.New("Normal error")

    // ค่าตัวแปร result, err จะถูก return กลับไป
    // ค่า result จะมีค่าเป็น 0 ตาม default ของภาษา
    return
}

func main() {

    // การรับค่า 2 ตัวแปรจาก function
    a, err := some(5)

    // การเลือกรับเฉพาะตัวแปร b จาก result โดยไม่สน err
    // ในเมื่อการประกาศตัวแปรทุกตัวใน go ต้องถูกเอาไปใช้ไม่งั้นจะ compile error ตัวไหนที่ไม่ได้ใช้จึงต้องใส่เป็น _ ไว้ก่อน 
    b, _ := some(5)
    
    // การเลือกรับเฉพาะตัวแปร err จาก result โดยไม่สน ตัวแปรแรก
    _, err2 := some(5)

}
```

คำถามคือ ถ้าเราเขียน  function ที่ return error มาด้วย เราควรจะเอา error ไว้ตำแหน่งไหน พี่ปุ๋ยเลยให้ดู ตัวอย่างจากพวก Library ที่อยู่ใน golang website ก็จะเห็นว่าเขาไว้เป็นตัวสุดท้ายเสมอเลย เพราะฉะนั้นเราจะเอาภาษาอะไรมาพัฒนา เราจำเป็นต้องแนวคิด แนวการเขียนของภาษานั้น ๆ ด้วย

## Collection Type in Go

ใน Go มี Type เป็น Collection อยู่ 3 ประเภท

- Array :   น่าจะรู้กันอยู่ ไม่พูดเยอะ
- Slice :   Array ที่ไม่ต้องกำหนดขนาด
- Map :     Collection แบบ Key, Value

### Array in Go

```go
// ประกาศตัวแปรที่ชื่อว่า numbers เป็น Array ของ int ที่มีขนาดเท่ากับ 5
var numbers [5]int

// การอ้างถึง Array ทำได้ใน []
numbers[0] = 1
numbers[1] = 2

// การประกาศ Array ของ String โดยสร้างจากค่าตั้งต้นเลย
var colors = [2]string{"Red", "Blue"}
```

#### รูปแบบของ For Loop

ใน Go มี Loop แค่แบบเดียวคือ Loop For เพราะฉะนั้นเวลาจะใช้ต้องไปประยุกต์เอาเอง ว่าจะใช้งานแบบไหน เช่นตัวอย่างของการเอา For Loop ไปใช้กับข้อมูลใน Array

```go
// Loop จาก Len ของ Array แล้วก็อ้างถึงข้อมูลตามตำแหน่งในแต่ละรอบ
for i := 0; i < len(colors); i++ {
    fmt.Println(colors[i])
}

// ใช้ Loop ร่วมกับ  range เพื่อแกะข้อมูลออกมาทีละตัว
for i, v := range colors {
    // range จะ return ค่ากลับมา 2 ตัวคือ Index และ Value  ถ้าไม่อยากได้ก็สามารถใช้  _ ได้
    fmt.Println(i, " ", v)
}

// infinite loop
for {

}

// for 5 loop อันนี้เอาไปประยกต์ใช้กับ while loop ได้
var i = 0
for i < 5 {
    i++
    log.Printf("Logging %v", i)
}

2021/01/XX 16:41:11 Logging 1
2021/01/XX 16:41:11 Logging 2
2021/01/XX 16:41:11 Logging 3
2021/01/XX 16:41:11 Logging 4
2021/01/XX 16:41:11 Logging 5
```

### Slice in Go

คือ Loop ที่ไม่ต้องกำหนดขนาดตอนตั้งต้น ที่เหลือเหมือน Array

```go
// Array
var colors = [5]string{"Red", "Blue", "Green", "Yellow", "Black"}

// Slice
var colorSlice = []string{"Red", "Blue", "Green", "Yellow", "Black"}
```

เราสามารถสร้าง Slice ได้ด้วยคำสั่ง make พร้อมระบุขนาดเบื้องต้นได้

```go
// Make
var colorMake = make([]string, 5)
```

ลอง Print ออกมาดู

```go
fmt.Println("colors =     ", colors)
fmt.Println("colorSlice =     ", colorSlice)
fmt.Println("colorMake =     ", colorMake)

colors =         [Red Blue Green Yellow Black]
colorSlice =     [Red Blue Green Yellow Black]
colorMake =      [    ]
```

เราสามารถตัดข้อมูลจาก Array มาเป็น Slice อีกตัวได้โดยระบุตำแหน่ง

```go
//  เอาข้อมูลจากตำแหน่งที่ 0 ถึงตัวก่อนตำแหน่งที่ 2 
a := colors[0:2]
fmt.Println("a = ", a)

[result]
a =  [Red Blue]
```

เราสามารถเพิ่มสมาชิกเข้าไปใน Slice ได้ผ่านคำสั่ง append

```go
a = append(a, "XXX")
fmt.Println("a = ", a)

[result]
a =  [Red Blue XXX]
```

```go
a = append(a, "YYY")
fmt.Println("a = ", a)

[result]
a =  [Red Blue XXX YYY]
```

คำถามที่สำคัญที่คนเขียน Go ต้องรู้คือ ถ้าเราแก้ไขข้อมูลใน Slice a (ซึ่งเกิดจาก Array colors) จะมีผลต่อข้อมูลใน Array colors ไหมนะลองแก้ตัวแรกดู

```go
a[0] = "New"
fmt.Println("colors,a = ", colors, a)

[result]
colors,a =  [New Blue XXX YYY Black] [New Blue XXX YYY]
```

จะเห็นได้ว่า Slice a ใช้ที่เก็บเดียวกันกับ Array colors เนื่องจากใช้ pointer เดียวกัน เพราะฉะนั้น

- ตอนที่เรา append XXX เข้าไปต่อท้าย a ข้อมูลก็จะไปทับตัวที่อยู่ใน colors ด้วย

```text
จากที่

        |------|------|-------|--------|-------|
        | Red  | Blue | Green | Yellow | Black |
        |------|------|-------|--------|-------|
        ^                                      ^
        |               colors                 |
        
        ^             ^
        |      a      |

พอเราสั่ง append a ด้วย XXX ก็เลยกลายเป็น 

        |------|------|-------|--------|-------|
        | Red  | Blue |  XXX  | Yellow | Black |
        |------|------|-------|--------|-------|
        ^                                      ^
        |               colors                 |
        
        ^                     ^
        |          a          |

เพราะมันไปต่อจาก a เดิม ซึ่งแปลว่า colors ก็ถูกแก้ด้วย
```

เราสามารถเรียงลำดับของที่อยู่ใน Slice ได้ด้วย

```go
var intSlice = []int{11, 3, 4, 1, 5, 6, 2}
sort.Ints(intSlice)   // เอา Slice มาเรียงลำดับตามตัวเลข
fmt.Println(intSlice)

var stringSlice = []string{"wat", "sivarat", "bomb", "karan"}
sort.Strings(stringSlice)   // เอา Slice มาเรียงลำดับตามตัวอักษร
fmt.Println(stringSlice)

[result]
[1 2 3 4 5 6 11]
[bomb karan sivarat wat]
```

### Map in Go

Collection Type สุดท้ายใน Go คือ Map ซึ่งมันก็คือการเก็บข้อมูลเป็นคู่แบบ Key, Value ที่เราคุ้นเคยเนี่ยแหละ

```go
numbers := make(map[string]int)
numbers["one"] = 1
fmt.Println("Numbers =", numbers)
```

ในกรณีที่เราต้องการอ่านค่าจาก Key ก็ทำได้

```go
n1 := numbers["one"]
fmt.Println("n1 =", n1)

[result]
n1 = 1
```

กรณีไม่มีมันจะเอา Zero Value (คืออันเดียวกับ Default Value) มาใส่ให้

```go
n2 := numbers["two"]
fmt.Println("n2 =", n2)

[result]
n2 = 0
```

ซึ่งถ้าเราต้องการตรวจสอบก่อนว่ามีค่าหรือเปล่าก็สามารถเช็ค error ของ map ได้

```go
if e2, found := numbers["two"]; found { // ถ้าเจอ ตัวแปร found จะเป็น true
    fmt.Println("Found = ", e2)
} else {
    fmt.Println("Not Found")
}

[result]
Not Found
```

## Error vs Defer

เนื่องจากใน Go ไม่มี exception เพราะฉะนั้นในกรณีที่เกิดปัญหาตอน runtime แล้วต้องการ Handle case ที่เกิดขึ้น เช่น ต้องปิด connection Database ถึงแม้ว่าจะเกิด Run Time Error ก็ตาม เราสามารถทำได้โดยการสร้าง Function ที่จะทำงานทุกครั้ง เมื่อ Function หลักของเราทำงานเสร็จ แล้วเช็คว่าเกิด panic ขึ้นได้ ผ่าน recover function

```go
func main() {

    deferAtEnd()

    deferWhenError()

}

func deferAtEnd() {
    fmt.Println("deferAtEnd - Start")
    // ประกาศว่าให้ เรียก panicHandler หลังจบการทำงานของ function ส่วนใหญ่เขียนไว้เพื่อ Close Resource ที่ร้องไว้ เมื่อทำงานเสร็จ
    defer panicHandler()
    fmt.Println("deferAtEnd - After defer panicHandler")

    fmt.Println("deferAtEnd - doSomething")

    fmt.Println("deferAtEnd - Start")
}

func deferWhenError() {
    fmt.Println("deferWhenError - Start")
    defer panicHandler()
    fmt.Println("deferWhenError - After defer panicHandler")

    b, err := ioutil.ReadFile("try_panic.go")
    fmt.Println("deferWhenError - After readfile")
    if err != nil {
        fmt.Println("deferWhenError - before panic")
        // call function panic เพื่อบอกว่าเกิด error ณ runtime
        panic("some error")
        fmt.Println("deferWhenError - after panic")
    }
    // จาก code จะเกิด panic และไม่เข้ามาทำงานในส่วนนี้
    fmt.Println(string(b))
    fmt.Println("deferAtEnd - After doSomething")

    fmt.Println("deferWhenError - Start")
}

// function ที่จะ Handle ทุกครั้งที่ deferAtEnd และ deferWithError ทำงานเสร็จ
func panicHandler() {
    // อ่าน error จาก function recover
    err := recover()
    // เช็คว่ามี panic ไหมโดยดุจาก err ถ้าไม่เกิด panic ก็ไม่เข้าใน if
    if err == "some error" {
        fmt.Println("panicHandler - Try to recover from panic")
    }
    fmt.Println("panicHandler - Do it at the end")    
}
```

เพราะฉะนั้น

```text
Error is for application error, Defer for run time and need to recover state
```

## Struct

เป็น Type ลักษณะ Object ที่สามารถกำหนดได้เองว่าภายในจะประกอบไปด้วย attribute อะไรบ้าง

### Simple struct

```go
type user struct {
    id   int
    name string
}

func main() {
    u1 := user{1, "bomb0069"}        // Go จะใส่ให้ตามลำดัับที่เรากำหนดไว้ กรณนี้ 1->id, bomb0069->name
    u2 := user{id: 2, name: "karan"} // ระบุชื่อเลยว่าจะเอาเข้า attribute ไหน
    u1.name = "XXX"                  // assign ค่าไปที่ attribute ได้
    fmt.Printf("%v", u1)
    fmt.Printf("%+v", u2)
}
```

หรือเราจะสร้าง function เพื่อ new object ก็ได้

```go
func New(id int, name string) User {
    return User{1, "bomb0069"}
}

func main() {
    u1 := New(1, "bomb0069")
    fmt.Printf("%v", u1)
}
```

### Struct with Struct Composition

```go
type User struct {
    Id   int
    UserForNew   // User มีการ Embeded เอา UserForNew  เข้ามาด้วย
}

type User2 struct {
    Id          int
    UserForNew  UserForNew  // User2 ประกอบไปด้วย id กับ UserForNew
}

type UserForNew struct {
    Name string
}

func New(id int, name string) User {
    return User{1, UserForNew{"bomb0069"}}
}

func main() {
    u1 := User{1, UserForNew{"bomb0069"}}
    u2 := User2{2, UserForNew{"karan"}}
    u3 := NewUser("bomb0069")
    u1.Name = "XXX"   // กรณีมีการ Embeded เราสามารถเรียกใช้ Attribute ในของที่เรา Embeded มาได้โดยตรง 
    fmt.Printf("%+v\n", u1)
    fmt.Printf("%+v\n", u1.Name) // อ้างถึงตรง
    fmt.Printf("%+v\n", u1.UserForNew.Name) // อ้่างถึงผ่าน UserForNew
    fmt.Printf("%+v\n", u2.UserForNew.Name)
    fmt.Printf("%+v\n", u3)

    fmt.Printf("%+v\n", u3.Name)
}

[result]
{Id:1 UserForNew:{Name:XXX}}  // ถึงจะเป็น Embeded เวลา print ก็ยังเห็นเหมือนมี UserForNew ครอบอยู่
XXX
XXX
karan
{Id:1 UserForNew:{Name:bomb0069}}
bomb0069
```

### Method in Struct

เนื่องจาก Go ไม่ใช่ Object Oriented เพราะถึงมี Struct แต่ก็มีแค่ Attribute ไม่มี Behavior หรือ Method แต่เราสามารสร้าง function ไปผูกกับ struct เพื่อทำหน้าที่คล้าย method ของ OOP ได้

```go
type User struct {
    Id int
    Name string
}

// ประกาศว่า function นี้มี receiver type (หมายถึง struct ที่จะเอาไปผูกแล้วส่งมาเมื่อถูกเรียก) คือ User
func (user User) printName() {
    fmt.Println("User := ", user.Name)
}

func main() {
    u := User{2, "karan"}
    // สามารถเอา u ที่เกิดจาก User มาเรียกใช้ printName ได้
    u.printName()
}
```

### Method of Struct and Overwrite Method in Go (for embeded)

```go
type User struct {
    Name string
}

type User2 struct {
    Id int
    User
}

func (user User) printName() {
    fmt.Println("User := ", user.Name)
}

func (user User2) printName() {
    fmt.Println("User 2 := ", user.Name)  // User2 สามารถเรียกใช้ Name ใน User ได้โดยตรงเพราะเป็น Embeded
}

func main() {
    u2 := User2{2, User{"karan"}}
    u2.printName()

}

[result]
User 2 :=  karan
```

### Method of Struct with Pointer

Method ก่อนหน้านี้เป็นแค่ action ที่ทำผ่าน struct เท่านั้นไม่สามารถ เปลี่ยนแปลงค่าที่อยู่ใน Struct ได้ ลองดู Code ข้างล่างนี้

```go
type User struct {
    Id   int
    Name string
}

func (user User) sleep() {
    user.Name = "XXXXX"   // มีการ Set ค่า Name เป็น XXXXX
}

func main() {
    u := User{2, "karan"}
    u.sleep()              // sleep เปลี่ยนค่า Name เป็น XXXXX
    fmt.Println(u.Name)    // print ค่า Name ออกมา
}

[result]
karan                      // ผลออกมาค่าไม่เปลี่ยน ยังเป็น karan เหมือนตอน New
```

เนื่องจาก receiver type ของ function sleep ส่ง User มาเป็น value ทำให้การเปลี่ยนแปลงภายในไม่มีผลกับภายนอก แก้โดยเพิ่ม * ไปตรง receiver type และ & ไปตอน new object (อันหลังผมไม่ใส่ก็ยังใช้ได้ แต่พี่ปุ๋ยบอกว่าให้ใส่เพื่อเราจะได้รู้ว่าเราจะเอาไปใช้แบบไหน)

```go
type User struct {
    Id   int
    Name string
}

func (user *User) sleep() {  // เพิ่ม * หน้า User เพื่อบอกว่าเป็น pointer
    user.Name = "XXXXX"
}

func main() {
    u := User{2, "karan"}     // สร้าง user แบบปกติ
    u.sleep()
    fmt.Println("User := " + u.Name)

    u2 := &User{2, "karan"}   // สร้าง user แบบมี & ด้านหน้า
    u2.sleep()
    fmt.Println("User2 :=" + u2.Name)
}

[result]
User := XXXXX   // ผลออกเหมือนกัน
User2 := XXXXX
```

## Backlog of Content

List ของ git log จากทีี่ลองทำตามใน Class

```text
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
