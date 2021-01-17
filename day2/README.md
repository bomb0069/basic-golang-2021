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
