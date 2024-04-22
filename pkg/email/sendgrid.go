package email

import (
	"booking-api/config"
	"fmt"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// A new function to send an email
func SendEmail(senderName, senderEmail, recipientName, recipientEmail, emailSubject, plainTextContent, htmlContent string) {

	env, err := config.LoadConfig("../")
	if err != nil {
		log.Fatal(err)
	}

	from := mail.NewEmail(senderName, senderEmail)
	subject := emailSubject
	to := mail.NewEmail(recipientName, recipientEmail)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(env.SEND_GRID_KEY)

	response, err := client.Send(message)

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func SendEmailTemplate(senderName, senderEmail, recipientName, recipientEmail, emailSubject, templateID string, dynamicData map[string]interface{}) error {

	env, err := config.LoadConfig("../")
	if err != nil {
		log.Fatal(err)
		return err
	}

	from := mail.NewEmail(senderName, senderEmail)
	subject := emailSubject
	to := mail.NewEmail(recipientName, recipientEmail)

	message := mail.NewV3Mail()
	message.SetFrom(from)
	message.SetTemplateID(templateID)

	p := mail.NewPersonalization()
	p.AddTos(to)
	p.Subject = subject
	p.DynamicTemplateData = dynamicData

	message.AddPersonalizations(p)

	response, err := sendgrid.NewSendClient(env.SEND_GRID_KEY).Send(message)
	if err != nil {
		log.Println(err)
		return err
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

	return nil
}
