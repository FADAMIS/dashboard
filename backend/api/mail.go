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

func SendRegisterConfirm(receiverEmail string, name string, surname string, campName string, date string) {
	godotenv.Load()
	senderEmail := os.Getenv("EMAIL_ADDRESS")
	senderPassword := os.Getenv("EMAIL_PASS")
	senderSmtpHost := os.Getenv("EMAIL_SMTP_HOST")
	senderSmtpPort, _ := strconv.Atoi(os.Getenv("EMAIL_SMTP_PORT"))

	body := "Potvrzení registrace účastníka " + name + " " + surname + " na kempu " + campName + ", který se bude konat " + date + ". Tato zpráva byla automaticky generována"

	m := gomail.NewMessage()
	m.SetHeader("From", senderEmail)
	m.SetHeader("To", receiverEmail)
	m.SetHeader("Subject", "Potvrzení registrace na kempech TechDays/HackDays")
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(senderSmtpHost, senderSmtpPort, senderEmail, senderPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: false, ServerName: senderSmtpHost}
	err := d.DialAndSend(m)

	// try again after ten minutes
	if err != nil {
		go func() {
			time.Sleep(time.Minute * 10)
			SendRegisterConfirm(receiverEmail, name, surname, campName, date)
		}()
	}
}

func SendParticipantList(receiverEmail string, participants []entities.Participant, campName string) {
	filename := excelizeParticipants(participants, campName)

	godotenv.Load()
	senderEmail := os.Getenv("EMAIL_ADDRESS")
	senderPassword := os.Getenv("EMAIL_PASS")
	senderSmtpHost := os.Getenv("EMAIL_SMTP_HOST")
	senderSmtpPort, _ := strconv.Atoi(os.Getenv("EMAIL_SMTP_PORT"))

	body := "Dobrý den,\nzde je zaslána tabulka s účastníky kempu" + campName + " určena pro tvorbu pohledávek. Tato zpráva byla automaticky generována."

	m := gomail.NewMessage()
	m.SetHeader("From", senderEmail)
	m.SetHeader("To", receiverEmail)
	m.SetHeader("Subject", "Tabulka účastníků "+campName+" pro tvorbu pohledávek")
	m.SetBody("text/plain", body)
	m.Attach(filename)

	d := gomail.NewDialer(senderSmtpHost, senderSmtpPort, senderEmail, senderPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: false, ServerName: senderSmtpHost}
	err := d.DialAndSend(m)

	// try again after ten minutes
	if err != nil {
		go func() {
			time.Sleep(time.Minute * 10)
			SendParticipantList(receiverEmail, participants, campName)
		}()
	}
}

func excelizeParticipants(participants []entities.Participant, campName string) string {
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

	filename := "./backup/" + strings.ToLower(campName) + "_" + time.Now().Format("02_January_2006") + ".xlsx"

	f.SaveAs(filename)

	return filename
}
