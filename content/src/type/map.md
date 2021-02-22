# Map

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
