package sendgridClient

import (
	"fmt"
	"log"
	"os"
	"spiel/notification-center/models"
	"strings"

	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmailToFounders(newUser models.User) {
	notifyBillsMessage := mail.NewPersonalization()

	var subject, plainTextContent string
	firstAndLast := fmt.Sprintf("%s %s", newUser.FirstName, newUser.LastName)

	from := mail.NewEmail("Spiel", "info@tryspiel.com")

	bills := []*mail.Email{
		mail.NewEmail("Draco Bill", "oscar@tryspiel.com"),
		mail.NewEmail("Fastlane Bill", "alex@tryspiel.com"),
		mail.NewEmail("Dirty Bill", "don@tryspiel.com"),
	}

	notifyBillsMessage.AddTos(bills...)

	if strings.Contains(newUser.Email, "ycombinator") || strings.Contains(newUser.Email, "tryspiel") {
		subject = fmt.Sprintf("🚀🚀🚀GUYS SOMEONE FROM YC JUST JOINED OUR PLATFORM🚀🚀🚀 %s", newUser.Email)
		plainTextContent = `🚀🚀🚀🚀🚀🚀🚀🚀🚀🚀🚀🚀🚀🚀 ` +
			firstAndLast +
			` has joined! WOO HOO! - heres their email: ` +
			newUser.Email +
			`🦄🦄🦄🦄🦄🦄🦄🦄🦄🦄☁️☁️☁️☁️☁️☁️☁️☁️☁️⛰️⛰️⛰️⛰️⛰️⛰️⛰️⛰️⛰️⛰️⛰️⛰️⛰️⛰️⛰️`
	} else {
		subject = fmt.Sprintf("We have a new user!")
		plainTextContent = `Heres the new user: ` +
			firstAndLast +
			`, and here's their email: ` +
			newUser.Email
	}

	personalization := mail.NewPersonalization()
	personalization.AddTos(bills...)

	message := mail.NewV3Mail()
	message.AddPersonalizations(personalization)
	message.SetFrom(from)
	message.Subject = subject
	message.AddContent(&mail.Content{
		Type:  "text/plain",
		Value: plainTextContent,
	})

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func SendEmailPromptToWebUser(webUser, spieler models.User) {
	from := mail.NewEmail("Spiel", "info@tryspiel.com")
	subject := fmt.Sprintf("Congratulations %s!", webUser.FirstName)

	firstAndLast := fmt.Sprintf("%s %s", webUser.FirstName, webUser.LastName)

	to := mail.NewEmail(firstAndLast, webUser.ID)
	plainTextContent := `You have taken the first step in joining Spiel, ` +
		`and receiving all the benefits we’ve worked hard to provide for you, and all our users! ` +
		spieler.FirstName + " " + spieler.LastName + `, ` + spieler.Title +
		` at ` + spieler.Company + `,  has received your question, ` +
		`download our app and sign up so you can see ` +
		spieler.FirstName + `'s personalized video answer specifically for you.`

	htmlContent := generateHTMLForWebQuestion(webUser, spieler)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(response.StatusCode)
		log.Println(response.Body)
		log.Println(response.Headers)
	}
}
