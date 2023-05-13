package utils

import (
	"bytes"
	"html/template"

	"github.com/challengego/db"
)

func prepare_email(name string, data db.RESUMEN) {
	
	context := map[string]interface{}{
		"Name":                name,
		"Balance":             data.Total_balance,
		"Promedio_Debito":     data.Average_debit,
		"Prmedio_Credito":     data.Average_credit,
		"Total_transacciones": data.Total_transaction,
		"Operaciones":         data.Transactions_per_month,
	}
	var body bytes.Buffer
	t := template.Must(template.ParseFiles("/templates/free-simple-card.html"))
	t.Execute(&body, context)
}
