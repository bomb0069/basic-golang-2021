# Struct

เป็น Type ลักษณะ Object ที่สามารถกำหนดได้เองว่าภายในจะประกอบไปด้วย attribute อะไรบ้าง

## Simple struct

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

## Struct with Struct Composition

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

## Method in Struct

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

## Method of Struct and Overwrite Method in Go (for embeded)

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

## Method of Struct with Pointer

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
