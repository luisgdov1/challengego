package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"

	"github.com/challengego/db"
)

func Prepare_email(name string, email string, data db.RESUMEN) {
	SENDGRID_USER := os.Getenv("USER_SENDGRID")
	SENDGRID_PASSWORD := os.Getenv("PASSWORD_SENDGRID")
	gmail_auth := smtp.PlainAuth("", SENDGRID_USER, SENDGRID_PASSWORD, "smtp.sendgrid.net")
	context := map[string]interface{}{
		"Name":                name,
		"Balance":             data.Total_balance,
		"Promedio_Debito":     data.Average_debit,
		"Prmedio_Credito":     data.Average_credit,
		"Total_transacciones": data.Total_transaction,
		"Operaciones":         data.Transactions_per_month,
	}
	var body bytes.Buffer
	t := template.Must(template.ParseFiles("templates/free-simple-card.html"))
	t.Execute(&body, context)

	message := bytes.NewBufferString("From: soporteskydelight@gmail.com\r\n" +
		"To: " + email + "\r\n" +
		"Subject: Reporte stori\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		body.String() + "\r\n")

	if erro := smtp.SendMail("smtp.sendgrid.net:587", gmail_auth, "soporteskydelight@gmail.com", []string{email}, message.Bytes()); erro != nil {
		fmt.Println("Error en el envio del email")
		fmt.Println(erro)
		return
	}
	fmt.Println("ENVIAMOS EL CORREO CON EXITO")
}
