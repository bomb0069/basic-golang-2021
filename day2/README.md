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

เวลาที่ Package มันเริ่มใหญ่ขึ้น เราก็สามารถแยกไปเป็น Repo ใหม่ได้เลยนะ

### Clean Architecture and Hexagonal architecture

มันแยกหน้าที่ได้ เห็นชัด แต่ต้องระวัง เดี๋ยวจะไปเจอ Pattern แถมจะไปเจอพวก cli สำหรับการ generate code ต้องระวัง คำถามคือแต่ละเรื่องจำเป็นต้องมีไหม

## Building REST APIs

### net/http

```go
func main() {
  http.HandleFunc("/", Response)
  http.HandleFunc("/users", UserHandler)
  http.ListenAndServe(":8080", nil)
}

func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

#### Func return func

```go
func Aaa(int a) func(int) int {
    // Func Aaa จะ Return เป็น Function ที่รับ int และ return int
}
```

#### Marshall/Unmarshall vs Encoder/Decoder

มี 2 วิธีในการ แปลง json ต้องไปลองเองว่าอะไรดี กว่า ดีสุดก็เขียนแล้วรันเทียบดู

```go
json.NewEncoder(w).Encode(u)
```

#### ยิง Performance Test

```cmd
wrk -c 100 -t 5 -d http://localhost:8080
```

### Echo

```go

func main() {
  e := echo.New()
  e.GET("/", hello)
  e.GET("/users", getUser)
  e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
  return c.JSON(http.StatusOK, "Hello world")
}
```

### Gin

```go
func hello(c *gin.Context) {
  c.JSON(http.StatusOK, "Hello world")
}

func main() {
  r := gin.Default()
  r.GET("/", hello)
  r.GET("/users", getUsers)
  r.Run(":8080")
}
```

```cmd
% go run . 
go: downloading github.com/gin-gonic/gin v1.6.3
go: downloading github.com/go-playground/validator/v10 v10.2.0
go: downloading github.com/golang/protobuf v1.3.3
go: downloading golang.org/x/sys v0.0.0-20200116001909-b77594299b42
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:  export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.hello (3 handlers)
[GIN-debug] GET    /users                    --> main.getUsers (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
```

#### using env:  export GIN_MODE=release

#### Close log mode

### [Fiber](https://github.com/gofiber/fiber)

เร็วกว่ามาก ข้างในใช้ [fasthttp](https://github.com/valyala/fasthttp)

```go
func hello(c *fiber.Ctx) error {
  return c.JSON("Hello world")
}

func main() {
  app := fiber.New()
  app.Get("/", hello)
  app.Get("/users", getUsers)
  app.Listen(":8080")
}
```

### สรุป

วิธีการคล้ายกัน แต่ต้องดูว่าเราชอบแบบไหน และ community มีคนใช้กันเยอะไหม แต่ดีที่สุดคือ POC ในหลาย ๆ ด้าน และเลือกให้เหมาะกับงาน

## [Workshop](https://github.com/up1/workshop-golang-20210116/tree/main/workshop-gin-mongodb)

### Structure

main -> user -> db -> mongo

### การทำงานของ Router และ Middlewares

### การทำ Resource และ Inject ไปใช้ใน Repository

- การทำ Inject DB Resource

- การทำ Resource ตรงกลางให้เขามาเรียกใช้ แล้วคืนตอนหลัง เช่นไปใช้ MongoBetween

### การแยก Implementation

ระหว่าง User ที่ใช้ Route ของ GIN อยู่ ไปใช้ อย่างอื่น ก็ต้องตัด Dependency ผ่าน Interface ที่เราต้องการ

## Link
https://github.com/swaggo/gin-swagger
https://echo.labstack.com/cookbook/graceful-shutdown
https://github.com/wg/wrk
https://github.com/tsliwowicz/go-wrk
https://echo.labstack.com/
https://github.com/gin-gonic/gin
https://github.com/gofiber/fiber
https://github.com/valyala/fasthttp/pulse
https://github.com/coinbase/mongobetween
https://echo.labstack.com/guide/routing
