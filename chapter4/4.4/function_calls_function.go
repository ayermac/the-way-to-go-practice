package main

var g string

func main() {
	g = "G"
	print(g)
	f1()
}

func f1() {
	g := "O"
	print(g)
	f2()
}

func f2() {
	print(g)
}

// G0G