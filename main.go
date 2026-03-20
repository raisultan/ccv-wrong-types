package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Bank struct {
	Name    string
	BinFrom string
	BinTo   string
}

func loadBankData(path string) ([]Bank, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var banks []Bank
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}

		name := strings.TrimSpace(parts[0])
		binFrom := strings.TrimSpace(parts[1])
		binTo := strings.TrimSpace(parts[2])

		banks = append(banks, Bank{
			Name:    name,
			BinFrom: binFrom,
			BinTo:   binTo,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return banks, nil
}

func extractBIN(cardNumber string) (string, error) {
	if len(cardNumber) < 6 {
		return "", fmt.Errorf("card number too short")
	}
	return cardNumber[:6], nil
}

func identifyBank(banks []Bank, cardNumber string) *Bank {
	bin, err := extractBIN(cardNumber)
	if err != nil {
		return nil
	}

	for i := range banks {
		if bin >= banks[i].BinFrom && bin <= banks[i].BinTo {
			return &banks[i]
		}
	}

	return nil
}

func validateLuhn(number string) bool {
	if len(number) == 0 {
		return false
	}

	sum := 0
	isSecond := false

	for i := len(number) - 1; i >= 0; i-- {
		digit := int(number[i] - '0')

		if digit < 0 || digit > 9 {
			return false
		}

		if isSecond {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		isSecond = !isSecond
	}

	return sum%10 == 0
}

func main() {
	banks, err := loadBankData("banks.txt")
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

		if validateLuhn(input) {
			fmt.Println("✅ Номер карты валиден (алгоритм Луна)")
		} else {
			fmt.Println("❌ Номер карты невалиден")
			continue
		}

		foundBank := identifyBank(banks, input)
		if foundBank != nil {
			fmt.Printf("🏦 Банк: %s\n", foundBank.Name)
		} else {
			fmt.Println("🏦 Банк не определён")
		}
	}
}
