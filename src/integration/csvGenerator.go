package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func add_products() {
	records := [][]string{
		{"sku", "product", "category", "quantity", "price"},
	}

	fruits := []string{"orange", "avocado", "grape", "strawberry", "lemon", "apple", "cocoa", "kiwi", "grapefruit", "papaya", "tamarindo", "blueberry", "lemoncreep", "straw fruit", "beattle fruit", "harris fruit", "lime", "bread fruit", "passion fruit", "abacaxi", "pineaple", "great breeze fruit", "increadible trip fruit", "strange fruit", "howle fruit", "tryme fruit", "today no fruit", "tomorrow fruit", "yesterdays fruit", "never fruit", "ever fruit", "forever fruit"}
	products := [][]string{
		{"yogurt", "yogurt", "5.99"},
		{"dry", "fruits", "2.99"},
		{"soap", "cleaning", "1.99"},
		{"soda 365ml", "drinks", "4.99"},
		{"soda 700ml", "drinks", "4.99"},
		{"soda 1L", "drinks", "4.99"},
		{"soda 1,5L", "drinks", "4.99"},
		{"soda 2L", "drinks", "4.99"},
		{"soap", "cleaning", "1.99"},
		{"soda diet 365ml", "drinks", "4.99"},
		{"soda diet 700ml", "drinks", "4.99"},
		{"soda diet 1L", "drinks", "4.99"},
		{"soda diet 1,5L", "drinks", "4.99"},
		{"soda diet 2L", "drinks", "4.99"},
		{"juice 300ml", "drinks", "10.00"},
		{"juice 600ml", "drinks", "10.00"},
		{"juice 1L", "drinks", "10.00"},
		{"juice 2L", "drinks", "10.00"},
		{"juice diet 300ml", "drinks", "10.00"},
		{"juice diet 600ml", "drinks", "10.00"},
		{"juice diet 1L", "drinks", "10.00"},
		{"juice diet 2L", "drinks", "10.00"},
		{"whey protein 900g", "suplements", "40.00"},
		{"whey protein 2kg", "suplements", "100.00"},
		{"whey protein 5kg", "suplements", "180.00"},
		{"whey protein isolado 900g", "suplements", "40.00"},
		{"whey protein isolado 2kg", "suplements", "100.00"},
		{"whey protein isolado 5kg", "suplements", "180.00"},
		{"whey protein hisdrolisado 900g", "suplements", "500.00"},
		{"whey protein hisdrolisado 2kg", "suplements", "1000.00"},
		{"whey protein hisdrolisado 5kg", "suplements", "1500.00"},
		{"ice cream", "ice cream", "5.00"},
		{"5L ice cream", "ice cream", "25.99"},
		{"absolute 300ml", "vodka", "70.00"},
		{"absolute 600ml", "vodka", "70.00"},
		{"absolute 700ml", "vodka", "70.00"},
		{"absolute 1L", "vodka", "70.00"},
		{"absolute 1,5L", "vodka", "70.00"},
		{"absolute 5L", "vodka", "70.00"},
		{"smirnoff ice", "drink", "9.99"},
		{"absolute zero 300ml", "vodka", "70.00"},
		{"absolute zero 600ml", "vodka", "70.00"},
		{"absolute zero 700ml", "vodka", "70.00"},
		{"absolute zero 1L", "vodka", "70.00"},
		{"absolute zero 1,5L", "vodka", "70.00"},
		{"absolute zero 5L", "vodka", "70.00"},
		{"smirnoff ice", "drink", "9.99"},
		{"51 ice", "drink", "9.99"},
		{"skol beats", "drink", "9.99"},
		{"absolute light 300ml", "vodka", "70.00"},
		{"absolute light 600ml", "vodka", "70.00"},
		{"absolute light 700ml", "vodka", "70.00"},
		{"absolute light 1L", "vodka", "70.00"},
		{"absolute light 1,5L", "vodka", "70.00"},
		{"absolute light 5L", "vodka", "70.00"},
		{"absolute condensed 300ml", "vodka", "70.00"},
		{"absolute condensed 600ml", "vodka", "70.00"},
		{"absolute condensed 700ml", "vodka", "70.00"},
		{"absolute condensed 1L", "vodka", "70.00"},
		{"absolute condensed 1,5L", "vodka", "70.00"},
		{"absolute condensed 5L", "vodka", "70.00"},
		{"absolute milky 300ml", "vodka", "70.00"},
		{"absolute milky 600ml", "vodka", "70.00"},
		{"absolute milky 700ml", "vodka", "70.00"},
		{"absolute milky 1L", "vodka", "70.00"},
		{"absolute milky 1,5L", "vodka", "70.00"},
		{"absolute milky 5L", "vodka", "70.00"},
		{"absolute new years edition 300ml", "vodka", "70.00"},
		{"absolute new years edition 600ml", "vodka", "70.00"},
		{"absolute new years edition 700ml", "vodka", "70.00"},
		{"absolute new years edition 1L", "vodka", "70.00"},
		{"absolute new years edition 1,5L", "vodka", "70.00"},
		{"absolute new years edition 5L", "vodka", "70.00"},
	}
	n := 1
	for i := 0; i < len(fruits); i++ {
		for j := 0; j < len(products); j++ {
			if j == 0 {
				product := []string{strconv.Itoa(100 + n), fruits[i], "fruits", "50", "3"}
				records = append(records, product)
				n++
			}
			product := []string{strconv.Itoa(100 + n), products[j][0] + " " + fruits[i], products[j][1], "70", products[j][2]}
			records = append(records, product)
			n++
		}
	}

	f, err := os.Create("csv_files/products.csv")
	defer f.Close()

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}
