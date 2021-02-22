# Slice

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
