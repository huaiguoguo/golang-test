package main

import "fmt"
import "math"

import "sort"

const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
	EB
)

func main() {
	var temp [3]bool
	fmt.Println(temp)
	fmt.Println(math.MaxInt32)
	// fmt.Println(B)
	// fmt.Println(KB)
	// fmt.Println(MB)
	// fmt.Println(GB)
	// fmt.Println(TB)
	// fmt.Println(PB)
	// fmt.Println(EB)

	if a := 1; a > 0 {
		fmt.Println("OK")
	}

	// LABEL1:
	// 	for i := 0; i < 10; i++ {
	// 		for {
	// 			fmt.Println(i)
	// 			continue LABEL1
	// 		}
	// 	}

	var test [2]int

	test2 := [...]int{1, 2, 3, 4, 5}

	fmt.Println(test)
	fmt.Println(test2)

	arr := [2][3][3]int{
		{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		},
		{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		},
	}

	fmt.Println(arr)

	a := [...]int{5, 2, 6, 3, 9}
	num := len(a)

	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] < a[j] {
				temp := a[j]
				a[j] = a[i]
				a[i] = temp
			}
		}
	}

	fmt.Println(a)

	sm := make([]map[int]string, 5)

	fmt.Println(sm)

	for i := range sm {
		sm[i] = make(map[int]string)
		sm[i][1] = "very good"
	}

	fmt.Println(sm)

	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e", 6: "f"}
	s := make([]int, len(m))
	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}

	fmt.Println(s)

	sort.Ints(s)
	fmt.Println(s)

	A(1, 2, 3, 4, 5)

	p := person{name: "jone", age: 19}

	fmt.Println(p)

	Ta := testa{testb: testb{Name: "B"}}

	// Tb := testb{Name: "B"}

	fmt.Println(Ta.Name, Ta.testc.Name)

	Ta.Print()

	var testz tz
	testz.Increase(100)

	fmt.Println(testz)

	c := make(chan bool)

	go func() {
		fmt.Println("go go go!")
		c <- true
	}()
	fmt.Println("...")
	fmt.Println(<-c)
	fmt.Println("hhhhhh")
}

type tz int

func (t *tz) Increase(num int) {
	*t += tz(num)
}

type testa struct {
	testb
}

type testb struct {
	testc
	Name string
}

type testc struct {
	Name string
}

func (a testc) Print() {
	fmt.Println("hahahah")

}

func A(a ...int) {
	fmt.Println(a)
}

type person struct {
	name string
	age  int
}
main.go|d48678537ecc5481d67aef9740ba6ca7188daba1|2159