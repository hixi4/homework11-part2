package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// Відкрити файл з контактами
	file, err := os.Open("numbers.txt")
	if err != nil {
		fmt.Println("Помилка відкриття файлу:", err)
		return
	}
	defer file.Close()

	// Створити регулярний вислів для пошуку телефонних номерів
	// Формати: (123) 456-7890, 123-456-7890, 123.456.7890, 1234567890
	regexFull := regexp.MustCompile(`\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4}`)
	// Створити регулярний вислів для пошуку 10-значних номерів
	regexTenDigits := regexp.MustCompile(`\d{10}`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		
		// Знайти всі відповідності у поточній стрічці для обох регулярних виразів
		fullMatches := regexFull.FindAllString(line, -1)
		tenDigitMatches := regexTenDigits.FindAllString(line, -1)
		
		for _, match := range fullMatches {
			fmt.Println("Знайдений номер (повний формат):", match)
		}
		for _, match := range tenDigitMatches {
			fmt.Println("Знайдений номер (10 цифр):", match)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Помилка читання файлу:", err)
	}
}
