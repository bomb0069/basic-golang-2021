# Go Simple Project

## Create demo.go

สร้างไฟล์ demo.go (ชื่อเดียวกับ module) ปกติคนที่เขียน go จะมีการจัดการ module และจะสร้าง file ที่ชื่อเดียวกับ module ขึ้นมา เพื่อเป็นจุดเริ่มต้นของ module นั้น ซึ่งไม่จำเป็นต้องมีก็ได้ โดย File นี้ปกติจะไม่ถูกใช้ในการรัน

```go
package demo    // กำหนดชื่อ package ให้เหมือนชื่อ module ที่ initial มา

func Greeting() string {  // ประกาศ Function ชื่อ Greeting โดยมี Return Type เป็น String
    return "Hello World"
}
```

## ข้อกำหนดเรื่อง Function Name

เนื่องจาก Go มีข้อกำหนด เรื่อง Access Modifier ของแต่ละ Function โดยที่ Function ใดที่ต้องการให้ Access จากภายนอก Package ก็จะต้องขึ้นต้นด้วยตัวพิมพ์ใหญ่ เช่น ในกรณีนี้ Greeting สามารถเข้าถึง/เรียกใช้ได้จาก Package อื่น ๆ ถ้าเปลี่ยนเป็น greeting ก็จะเรียกใช้ไม่ได้

## Create main.go

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

## Simple Project Structure

```tree
demo
|- go.mod   // Package dependency for go project
|- demo.go  // Default go file for each module
|- cmd
    |- main.go   // runable file with main function
```
