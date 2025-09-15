package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
)

type String struct { // структура, если строка
	Input string
	Type  string
	Len   string
}

func (s *String) Scan(input any) { // input any (любой тип, но внутри указать ожидаемый)
	str := input.(string) // указываем как одидаемый - строка
	fields := strings.Fields(str)
	s.Input = str
	s.Type = "string"
	s.Len = "Это строка длиной " + strconv.Itoa(len(fields)) + " символов"
}

type Integer struct { // структура, если число
	Input  int
	Type   string
	Square string
}

func (i *Integer) Scan(input any) {
	in := input.(int)
	sq := math.Pow(float64(in), 2) // квадрат числа
	i.Input = in
	i.Type = "int"
	i.Square = "Это число и его квадрат: " + strconv.FormatFloat(sq, 'f', 2, 64)
}

type Bool struct { // структура, если булева
	Input  bool
	Type   string
	Square string
}

func (b *Bool) Scan(input any) {
	bl := input.(bool)
	b.Input = bl
	b.Type = "bool"
	b.Square = "Булево значение: " + strconv.FormatBool(bl)
}

type Float struct { // структура, если число с плав.точ.
	Input  float64
	Type   string
	Square string
}

func (f *Float) Scan(input any) {
	fl := input.(float64)
	f.Input = fl
	f.Type = "float64"
	f.Square = "Неизвестный тип"

}

func main() {
	var i Integer // переменная нужного типа
	var s String
	var b Bool
	var f Float
	for {
		pp.Print("Ввод: ")
		scanner := bufio.NewScanner(os.Stdin)
		if ok := scanner.Scan(); !ok {
			pp.Println("Ошибка ввода")
		}
		text := scanner.Text()
		if text == "exit" {
			break
		}

		if num, err := strconv.Atoi(text); err == nil { // если num переведет текст в число - значит число,
			i.Scan(num) // и если err = nil
			pp.Println(i)

		} else if fl, err := strconv.ParseFloat(text, 64); err == nil {
			f.Scan(fl)
			pp.Println(f)
		} else if bl, err := strconv.ParseBool(text); err == nil {
			b.Scan(bl)
			pp.Println(b)
		} else {
			s.Scan(text)
			pp.Println(s)
		}
	}
}
