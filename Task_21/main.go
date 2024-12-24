/* Реализовать паттерн «адаптер» на любом примере. */
package main

import "fmt"

type Target interface {
	GetLengthInCentimeters() float64
}

type Adaptee struct{}

func (a *Adaptee) GetLengthInInches() float64 {
	return 10.0
}

type Adapter struct {
	adaptee *Adaptee
}

func (a *Adapter) GetLengthInCentimeters() float64 {
	inches := a.adaptee.GetLengthInInches()
	return inches * 2.54
}

type Client struct {
	target Target
}

func (c *Client) PrintLength() {
	length := c.target.GetLengthInCentimeters()
	fmt.Printf("Длина в сантиметрах: %.2f см\n", length)
}

func main() {
	adaptee := &Adaptee{}
	adapter := &Adapter{adaptee: adaptee}
	client := &Client{target: adapter}
	client.PrintLength()
}
