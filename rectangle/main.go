package main

import (
	"encoding/json"
	"fmt"
)

type Rectangle struct {
	Width  float32 `json:"width" desc:"Chieu dai"`
	Height float32 `json:"height" desc:"Chieu rong"`
}

// Calculate area
//   - Formula: Width * Height
//
// @return float32
func (r *Rectangle) Area() float32 {
	return r.Width * r.Height
}

// Calculate perimeter
//   - Formula: (Width + Height) * 2
//
// @return float32
func (r *Rectangle) Perimeter() float32 {
	return (r.Width + r.Height) * 2
}

func main() {
	rectangle := Rectangle{
		Width:  10,
		Height: 20,
	}

	area := rectangle.Area()
	perimeter := rectangle.Perimeter()
	fmt.Println("area:", area)
	fmt.Println("perimeter:", perimeter)

	data, err := json.Marshal(rectangle)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("data tag:", string(data))
}
