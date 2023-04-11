package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner1 := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite um número A: ")
	scanner1.Scan()
	a := scanner1.Text()
	scanner2 := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite um número B: 20")
	scanner2.Scan()
	b := scanner2.Text()
	fmt.Println()

}
func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Não é possivel a divisão por zero")
	}
	return a / b, nil
}
