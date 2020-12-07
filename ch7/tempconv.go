package main

import (
	"flag"
	"fmt"
)

func main()  {

}

type celsiusFlag struct {
	Celsius
}

type Celsius float64
type Fahrenheit float64

func CToF(c Celsius) Fahrenheit {return Fahrenheit(c*9/5 + 32)}
func FToC(f Fahrenheit) Celsius {return Celsius((f-32)*5/9)}

func (f *celsiusFlag) Set(s string) error {
	var uint string
	var value float64
	fmt.Scanf(s, "%f%s", &value, &uint) // 从输入 s中解析一个浮点值value和一个字符串unit
	switch uint {
	case "C", "oC" :
		f.Celsius = Celsius(value)
		return nil
	case "F", "oF" :
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}

	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) * Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
