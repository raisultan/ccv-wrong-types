package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"credit-card-validator/internal/bank"
	"credit-card-validator/internal/luhn"
)

func main() {
	banks, err := bank.LoadFromFile("banks.txt")
	if err != nil {
		fmt.Printf("Ошибка загрузки банков: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Загружено банков: %d\n", len(banks))

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nВведите номер карты (или 'exit' для выхода): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("До свидания!")
			break
		}

		if len(input) < 13 || len(input) > 19 {
			fmt.Println("Неверная длина номера карты")
			continue
		}

		if luhn.Validate(input) {
			fmt.Println("✅ Номер карты валиден (алгоритм Луна)")
		} else {
			fmt.Println("❌ Номер карты невалиден")
			continue
		}

		foundBank := bank.FindBank(banks, input)
		if foundBank != nil {
			fmt.Printf("🏦 Банк: %s\n", foundBank.Name)
		} else {
			fmt.Println("🏦 Банк не определён")
		}
	}
}
