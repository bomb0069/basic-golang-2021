# Error and Defer

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
