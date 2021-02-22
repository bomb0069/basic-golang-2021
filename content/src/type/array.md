# Array

```go
// ประกาศตัวแปรที่ชื่อว่า numbers เป็น Array ของ int ที่มีขนาดเท่ากับ 5
var numbers [5]int

// การอ้างถึง Array ทำได้ใน []
numbers[0] = 1
numbers[1] = 2

// การประกาศ Array ของ String โดยสร้างจากค่าตั้งต้นเลย
var colors = [2]string{"Red", "Blue"}
```

## รูปแบบของ For Loop

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
