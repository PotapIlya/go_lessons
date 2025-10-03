package main

import "fmt"

type animal interface {
	speaker
}

type speaker interface {
	speak() string
}

func echo(a animal) {
	fmt.Println(a.speak())
}

type cat struct{}

func (c *cat) speak() string {
	return "Мяу"
}

type dog struct{}

func (d *dog) speak() string {
	return "Гав"
}

func main() {
	var cat animal = &cat{}
	echo(cat)

	var dog animal = &dog{}
	echo(dog)
}

/*
type Logger interface {
    Log(message string)
}

type ConsoleLogger struct{}

func (ConsoleLogger) Log(message string) {
    fmt.Println("Console:", message)
}

type FileLogger struct {
    File *os.File
}

func (f FileLogger) Log(message string) {
    f.File.WriteString(message + "\n")
}

func DoWork(l Logger) {
    l.Log("Начинаем работу...")
    l.Log("Работа завершена!")
}
*/

/*
type Payment interface {
    Pay(amount float64) error
}

type PayPal struct{}
func (PayPal) Pay(amount float64) error {
    fmt.Println("Оплата через PayPal:", amount)
    return nil
}

type CreditCard struct{}
func (CreditCard) Pay(amount float64) error {
    fmt.Println("Оплата картой:", amount)
    return nil
}

func Checkout(p Payment, amount float64) {
    _ = p.Pay(amount)
}
*/
