package bank

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

func LoadFromFile(path string) ([]Bank, error) {
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

func FindBank(banks []Bank, cardNumber string) *Bank {
	if len(cardNumber) < 6 {
		return nil
	}

	prefix := cardNumber[:6]

	for i := range banks {
		if prefix >= banks[i].BinFrom && prefix <= banks[i].BinTo {
			return &banks[i]
		}
	}

	return nil
}
