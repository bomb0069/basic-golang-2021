// เมื่อก่อนใช้ชื่อ package เดียวกัน แต่ตอนนี้ให้แยกกันเพื่อชัดเจน
package demo_test

import (
	"demo"
	"testing"
)

func TestHello(t *testing.T) {
	// := คือ Short Assignment
	/*
		r := demo.Greeting()
	*/

	// prepare var before
	var r string
	r = demo.Greeting()
	if r != "Hello World" {
		t.Errorf("Error with %v", r)
	}

}
