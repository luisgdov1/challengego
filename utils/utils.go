package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/challengego/db"
)

func ReadDataCSV(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

// FORMATO MES/DIA
func ClassifiedData(data_csv [][]string, erro error) db.RESUMEN {
	if erro != nil {
		fmt.Println("ERROR EN LOS DATOS")
	}
	var balance_total = 0.00
	var debit_balance = 0.00
	var credit_balance = 0.00
	var debit_ops = 0
	var credit_ops = 0
	var diccionario = make(map[string]int)
	for _, sublista := range data_csv {
		month_number := strings.Split(string(sublista[1]), "/")
		month_name := GetMonthKey(month_number)
		total_ops, exists := diccionario[month_name]
		balance := ParseFloatString(string(sublista[2]))
		if exists {
			diccionario[month_name] = total_ops + 1
		} else {
			diccionario[month_name] = 1
		}

		balance_total = balance_total + balance
		if balance > 0 {
			debit_ops++
			debit_balance = debit_balance + balance
		} else {
			credit_ops++
			credit_balance = credit_balance + balance
		}
	}

	transactions := make([]db.TRANSACTIONS_RESUMEN, 0)
	for key, value := range diccionario {
		transaction := db.TRANSACTIONS_RESUMEN{
			Month:               key,
			Number_transactions: value,
		}
		transactions = append(transactions, transaction)
	}
	resumen_object := db.RESUMEN{
		Total_balance:          float32(balance_total),
		Total_transaction:      debit_ops + credit_ops,
		Average_debit:          float32(debit_balance) / float32(debit_ops),
		Average_credit:         float32(credit_balance) / float32(credit_ops),
		Transactions_per_month: transactions,
	}
	return resumen_object
}

func ParseFloatString(number_string string) float64 {
	float_value, erro := strconv.ParseFloat(number_string, 64)
	if erro != nil {
		panic(erro)
	}
	return float_value
}

func GetMonthKey(number_string []string) string {
	str_aux := number_string[0]
	int_month, erro := strconv.Atoi(str_aux)
	if erro != nil {
		panic(erro)
	}
	name_month := time.Month(int_month).String()
	return name_month
}
