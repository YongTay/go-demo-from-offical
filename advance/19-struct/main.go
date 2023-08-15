package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) Introduce() {
	fmt.Println("Hi, my name is", p.Name, ", I'm", p.Age, "years old")
}

// 官方没有支持函数重载功能, 需要尽量通过合适的函数参数来实现，也可以自行通过反射实现
// func (p *Person) Introduce(name string) {
// 	p.Name = name
// 	fmt.Println("Hi, my name is", p.Name)
// }

func (p *Person) NewInstance(args ...any) *Person {
	value := reflect.ValueOf(p)
	if len(args) == 2 {
		var name string
		var age int
		// 尝试将第一个参数转换为string类型，否则转换为int类型
		val, ok := args[0].(string)
		if ok {
			name = val
			age = args[1].(int)
		} else {
			age = args[0].(int)
			name = args[1].(string)
		}
		ret := value.MethodByName("NewInstanceOfNameAndAge").Call([]reflect.Value{reflect.ValueOf(name), reflect.ValueOf(age)})
		return ret[0].Interface().(*Person)
	} else {
		if len(args) == 1 {
			// 尝试将第一个参数转换为string类型，否则转换为int类型
			val, ok := args[0].(string)
			if ok {
				ret := value.MethodByName("NewInstanceOfName").Call([]reflect.Value{reflect.ValueOf(val)})
				return ret[0].Interface().(*Person)
			} else {
				ret := value.MethodByName("NewInstanceOfAge").Call([]reflect.Value{reflect.ValueOf(val)})
				return ret[0].Interface().(*Person)
			}
		}
	}
	return &Person{}
}

func (p *Person) NewInstanceOfAge(age int) *Person {
	return &Person{Age: age}
}
func (p *Person) NewInstanceOfName(name string) *Person {
	return &Person{Name: name}
}
func (p *Person) NewInstanceOfNameAndAge(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}

// 不会出现重载现象
func Introduce() {
	fmt.Println("Hi, my name is anonymous")
}

// func Introduce(s string) {
// 	fmt.Println("Hi, my name is anonymous")
// }

func main() {
	// p := Person{"John", 20}
	// p.Introduce()
	// Hi, my name is John

	// Introduce()
	// Hi, my name is anonymous

	// p2 := p.NewInstanceOfAge(22)
	// p2.Introduce()

	p3 := Person{"Tom", 23}
	p4 := p3.NewInstance("Jerry")
	p4.Introduce()
}
