# Go Type

## Alias TYPE

เราสามารถสร้าง Type ของตัวแปลใน Go ใหม่ได้ผ่าน  คำสั่ง type

```go
    type MONTH int  // สร้าง type ใหม่ที่ชื่อว่า MONTH โดยอ้างอิงจาก int

    var january MONTH // สร้างตัวแปลชื่อ january เป็นประเภท MONTH
```

## Example of Error

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
