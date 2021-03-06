package main

import (
	"fmt"
	"math"
	"runtime"
)

// 定数宣言
const Pi = 3.14

func add(x int, y int) int {
	return x + y
}

func swap(str1, str2 string) (string, string) {
	return str2, str1
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func test001(x int) int {
	var i int = 1
	// 関数内ではvar宣言の代わりに := の代入文を使って
	// 暗黙的な型宣言ができる
	k := i + 2
	return k
}

func test002(x float64) int {
	var y int = int(x)
	// シンプルな型変換記載
	z := int(x)
	return y + z
}

func test003() {
	// 定数を出力
	fmt.Println(Pi)
}

func test004() {
	// For文
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func test005() {
	// while文はないのでForのみ
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	// 無限ループの場合は条件なしで書く
	//for {
	//}
}

func test006() {
	// if文
	value := 2
	if value == 2 {
		fmt.Println(value)
	} else {
		fmt.Println("value invalid")
	}
}

func test007() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

func pow(x, n, lim float64) float64 {
	// if文の中に評価前ステートメントを書くことができる
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func test008() {
	// switch
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
}

func test009() {
	// defer
	// 関数からリターンするまで実行を遅延させる
	defer fmt.Println("world")
	fmt.Println("hello")
}

func test010() {
	// defer が複数ある場合はスタックされる
	// スタックなので9から0までカウントされる
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func test011() {
	// ポインタ
	i, j := 42, 2000
	p := &i
	// iを出力するので42
	fmt.Println(*p)
	*p = 21
	// iが書き換わるので21
	fmt.Println(i)

	p = &j
	*p = *p / 20
	// jが書き換わるので100
	fmt.Println(j)
}

type Vertex struct {
	X int
	Y int
}

func test012() {
	// 構造体
	fmt.Println(Vertex{1, 2})
}

func test013() {
	// 構造体のアクセス
	v := Vertex{1, 2}
	fmt.Println(v.X)
	fmt.Println(v.Y)
}

func test014() {
	// 構造体には.でアクセスする
	// ポインタは(*p).Xでアクセスしてもp.Xでアクセスも可
	v := Vertex{2, 3}
	p := &v
	(*p).X = 10
	p.Y = 20
	fmt.Println(v)
}

func test015() {
	// 配列
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a)
}

func test016() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func test017() {
	// Slices
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)
}

func test018() {
	// Slicesは参照を使用している
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	x := names[1:3]
	fmt.Println(x)
	// PaulがXXXに書き換えられる
	x[0] = "XXX"
	fmt.Println(names)
}

func test019() {
	// スライスリテラル
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
}

func main() {
	fmt.Printf("start\n")
	test019()
	//fmt.Println(add(1, 2))
	//a, b := swap("a", "b")
	//fmt.Println(a, b)
	//fmt.Println(split(17))
}
