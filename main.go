package main

func main() {
	addN := func(m int) func(int) int {
		return func(n int) int {
			return m + n
		}
	}

	addFive := addN(5)
	result := addN(6)
	println(addFive)
	println(result)
}
