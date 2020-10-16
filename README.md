# algorithm
常用的算法

### 时间复杂度样例
```go
package demo

// O(1)
func o1() {
	x := 0
	y := 0
	temp := x
	x = y
	y = temp
}

// O(n)
func On() {
	n := 100
	x := 0
	for i := 1; i <= n; i++ {
		x++
	}
}

// O(logN)
func LogN() {
	i := 1
	n := 100
	for i < n {
		i = i * 2
	}
}

// O(nlogN)
func OnlogN() {
	n := 100
	for i := 1; i <= n; i++ {
		x := 1
		for x < n {
			x = x*2
		}
	}
}

// O(n^2)
func ON2() {
	n := 100
	x := 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			x++
		}
	}
}
```
