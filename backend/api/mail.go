package api

import (
	"crypto/tls"
	"os"
	"reflect"
	"strings"
	"time"

	"strconv"

	"github.com/FADAMIS/dashboard/entities"
	"github.com/joho/godotenv"
	"github.com/xuri/excelize/v2"
	"gopkg.in/gomail.v2"
)

func SendRegisterConfirmation(receiverEmail string, name string, surname string, campName string, date int64) {
	godotenv.Load()
	senderEmail := os.Getenv("SENDER_ADDRESS")
	senderPassword := os.Getenv("SENDER_PASS")
	senderSmtpHost := os.Getenv("SMTP_HOST")
	senderSmtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	dateString := time.Unix(date, 0).Format("02. 01. 2006")

	body := "Potvrzení registrace účastníka " + name + " " + surname + " na kempu " + campName + ", který se bude konat " + dateString + ". Tato zpráva byla automaticky generována"

	m := gomail.NewMessage()
	m.SetHeader("From", senderEmail)
	m.SetHeader("To", receiverEmail)
	m.SetHeader("Cc", senderEmail)
	m.SetHeader("Subject", "Potvrzení registrace na kempech TechDays/HackDays")
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(senderSmtpHost, senderSmtpPort, senderEmail, senderPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: false, ServerName: senderSmtpHost}
	err := d.DialAndSend(m)

	// try again after ten minutes
	if err != nil {
		go func() {
			time.Sleep(time.Minute * 10)
			SendRegisterConfirmation(receiverEmail, name, surname, campName, date)
		}()
	}
}

func SendParticipantList(camp entities.Camp) {
	filename := excelizeParticipants(camp)

	godotenv.Load()
	receiverEmail := os.Getenv("RECEIVER_ADDRESS")
	senderEmail := os.Getenv("SENDER_ADDRESS")
	senderPassword := os.Getenv("SENDER_PASS")
	senderSmtpHost := os.Getenv("SMTP_HOST")
	senderSmtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	campDate := time.Unix(camp.Date, 0).Format("02. 01. 2006")

	body := "Dobrý den,\nzde je tabulka s účastníky kempu" + camp.Name + " " + campDate + ". Tato zpráva byla vygenerována automaticky.\nS pozdravem\nTým kempů"

	m := gomail.NewMessage()
	m.SetHeader("From", senderEmail)
	m.SetHeader("To", receiverEmail)
	m.SetHeader("Cc", senderEmail)
	m.SetHeader("Subject", camp.Name+" "+campDate+" - prezence")
	m.SetBody("text/plain", body)
	m.Attach(filename)

	d := gomail.NewDialer(senderSmtpHost, senderSmtpPort, senderEmail, senderPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: false, ServerName: senderSmtpHost}
	err := d.DialAndSend(m)

	// try again after ten minutes
	if err != nil {
		go func() {
			time.Sleep(time.Minute * 10)
			SendParticipantList(camp)
		}()
	}
}

func excelizeParticipants(camp entities.Camp) string {
	participants := camp.Participants

	f := excelize.NewFile()
	defer f.Close()

	style, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	})

	f.SetCellStr("Sheet1", "A1", "Jméno")
	f.SetCellStr("Sheet1", "B1", "Příjmení")
	f.SetCellStr("Sheet1", "C1", "E-Mail")
	f.SetCellStr("Sheet1", "D1", "Telefon")

	f.SetRowStyle("Sheet1", 1, 1, style)

	for y, p := range participants {
		v := reflect.ValueOf(p)
		// values in participant struct
		values := make([]string, v.NumField())
		for i := 0; i < v.NumField(); i++ {
			// append only if value is string
			if v.Field(i).Kind() == reflect.String {
				values[i] = v.Field(i).String()
			}
		}

		// put values in spreadsheet
		for x := 0; x < len(values); x++ {
			cell, _ := excelize.CoordinatesToCellName(x, y+2)
			f.SetCellStr("Sheet1", cell, values[x])
		}
	}

	date := time.Unix(camp.Date, 0).Format("02_January_2006")
	filename := "./backup/" + strings.ToLower(camp.Name) + "_" + date + ".xlsx"

	f.SaveAs(filename)

	return filename
}
