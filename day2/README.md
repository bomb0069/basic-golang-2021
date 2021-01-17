# Day 2

## Review Homework

### File Permission

0644
0755

### Folder Structure for Entity

user.go and entity/user.go are you confuse which file are you change

### ทิศทางของ Code func ใน struc

func (cli Cli) AddUser

func (cli *Cli) PrintUser

### Golang ไม่มี Case Type มีแต่ Convert Type

```go
var users []byte
users = []byte("[]")  // convert string to byte array
```

### Hardcode vs Initial Parameter

รอดูพี่ปุ๋ยพาทำ

```go
var storage := Store (Path: "static/db.json")
```

### เอา Interface มากั้น การทำงานของ Service ที่เป็น Store

[Example Code](https://github.com/up1/workshop-golang-20210116/tree/main/workshop-cli)

```go
func main() {
  // Initial dependencies
  store := cli.FileStore{Filename: "demo.json"}
  userService := cli.UserService{Store: &store}

  // Parser input from command line
  command := os.Args[1]
  switch command {
  case "add":
    name := os.Args[2]
    age := os.Args[3]
    iage, _ := strconv.Atoi(age)
    userService.AddNew(name, iage)
  case "list":
    res := userService.ListAll()
    fmt.Println(res)
  }
}
```

### Demo Package for SubPackage/SubModule

```cmd
mkdir demo_package
cd demo_package

go mod init sample

touch sample.go
```

create sub packag section 1

```cmd
mkdir section1
cd section1

touch user.go
```

**** อย่าทำนะ  import "../sample/section1" อันนี้แปลว่าไม่เข้าใจ

**** การทำงานข้าม Package จำเป็นต้องเอา Interface มากั้นไหม ดูความเหมาะสม Standard ส่วนใหญ่เขาเอามากั้น แต่ถ้าเราใช้ภายในกันเองก็อาจจะไม่จำเป็นเพราะ เรื่องความซับซ้อน

### Update Version of Library

อยู่ใน go.mod ไม่ควรแก้ไขเอง
ลบใน go.mod เลยก็ได้ เดี๋ยวมันก็เอาของใหม่มา

เวลาที่เราเอา import package ออกจาก Code มันจะยังไม่ clear ออกจาก go.mod ต้องเคลียร์เอง

```cmd
go mod tidy // clear package  ที่ไม่ได้ใช้งาน
```

อ่านเพิ่ม
[How to Write Go Code](https://golang.org/doc/code.html)

ที่อยู่ของ Library ที่เราเอามาลงไว้ที่เครื่องจะอยู่ใน

${GOPATH}/pkg/mod

### GoLang is test by Default นะครับ

คนเขียนภาษา Go ยังเขียน Test เลย เราเป็นใคร เอา Go มาใช้แล้วไม่เขียน Test

## Go project structure

### Flat

handler.go ทำงานหลังจาก route.go เพื่อ handle กรณีที่ส่งต่อไป package อื่น ๆ

route.go ก็ไม่ควรยาว ๆ แล้วให้มีแยก route ย่อยออกไป

### Layer

แต่ถ้าเยอะแล้วกระโดดไปมาหรือเปล่า แล้วก็มาต้องการ IDE แบบที่ Link ไปมาได้เพราะมันเยอะมาก หายากมาก เปิดไฟล์เต็มไปหมด

### Modular

แยกตามกลุ่ม Business เลย แยก Sub Module เป็นเรื่อง ๆ ตามกลุ่มการทำงาน เวลาที่เราต้องแก้อะไร เราก็จะสามารเข้าไปหามันได้เลย
