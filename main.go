package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("valores.txt")
	scanner := bufio.NewScanner(f)

	var somaTotal, somaMarcados float64

	if len(os.Args) != 1 && os.Args[1] == "ajuda" {
		fmt.Println("Coloque os valores no arquivo 'valores.txt' para somá-los.")
		fmt.Println("Se alguma valor começar com '#', será mostrado separadamente uma outra soma subtraindo-os.")
		os.Exit(0)
	} else {
		fmt.Print("Rode 'somador ajuda' para um guia completo.\n\n")
	}

	for scanner.Scan() {
		var v float64
		if text := scanner.Text(); text != "" && !strings.HasPrefix(text, "#") {
			v = parseFloat64(text)

			somaTotal += v
		} else if strings.HasPrefix(text, "#") {
			v = parseFloat64(strings.TrimPrefix(text, "#"))
			somaMarcados += v
			somaTotal += v
		}
	}

	fmt.Printf("A soma total é %.2f\n", somaTotal)

	if somaMarcados != 0 {
		fmt.Printf("A soma sem os valores marcados é %.2f", somaTotal - somaMarcados)
	}
}

func parseFloat64(text string) (v float64) {
	v, err := strconv.ParseFloat(text, 64)
	if err != nil {
		panic(err)
	}

	return
}
