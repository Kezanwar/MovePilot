package email

import (
	"errors"
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const (
	admin_email    = "admin@affiliate-lead-gen.com"
	no_reply_email = "noreply@affiliate-lead-gen.com"
	admin_name     = "move-pilot Admin"
	no_reply_name  = "move-pilotot"
)

type Client struct {
	client *sendgrid.Client
	apiKey string
}

type EmailMessage struct {
	From        *mail.Email
	To          *mail.Email
	Subject     string
	PlainText   string
	HTMLContent string
}

type SendOptions struct {
	ToEmail      string
	ToName       string
	Subject      string
	TemplateData ActionEmailTemplateData
}

func NewClient() (*Client, error) {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	if apiKey == "" {
		return nil, errors.New("SENDGRID_API_KEY environment variable is not set")
	}

	return &Client{
		client: sendgrid.NewSendClient(apiKey),
		apiKey: apiKey,
	}, nil
}

func (c *Client) Send(options SendOptions) error {
	if options.ToEmail == "" {
		return errors.New("recipient email is required")
	}
	if options.Subject == "" {
		return errors.New("subject is required")
	}
	if options.TemplateData.ReceiverName == "" {
		options.TemplateData.ReceiverName = options.ToName
		if options.TemplateData.ReceiverName == "" {
			options.TemplateData.ReceiverName = "User"
		}
	}
	if options.TemplateData.Title == "" {
		return errors.New("template title is required")
	}
	if len(options.TemplateData.Content) == 0 {
		return errors.New("template content is required")
	}

	// Generate HTML content from template
	htmlContent := GenerateEmailTemplate(options.TemplateData)

	// Generate plain text version
	plainText := GeneratePlainTextEmail(options.TemplateData)

	// Create email message
	from := mail.NewEmail(no_reply_name, no_reply_email)
	to := mail.NewEmail(options.ToName, options.ToEmail)

	m := mail.NewSingleEmail(
		from,
		options.Subject,
		to,
		plainText,
		htmlContent,
	)

	response, err := c.client.Send(m)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	if response.StatusCode >= 400 {
		return fmt.Errorf("email send failed with status code %d: %s", response.StatusCode, response.Body)
	}

	fmt.Println(response.Body)

	return nil
}

type OTPEmailData struct {
	ToEmail string
	ToName  string
	OTPCode string
}

func (c *Client) SendOTP(data OTPEmailData) error {
	if data.ToEmail == "" {
		return errors.New("recipient email is required")
	}
	if data.OTPCode == "" {
		return errors.New("OTP code is required")
	}

	// Create the OTP content with the code prominently displayed
	content := []string{
		"Your verification code is:",
	}

	// The OTP code will be displayed in a highlighted box in the HTML version
	otpCodeDisplay := fmt.Sprintf(`
		<div style="background-color: #f5f5f5; border: 2px solid #333333; padding: 20px; text-align: center; margin: 20px 0; border-radius: 8px;">
			<h1 style="font-size: 32px; letter-spacing: 5px; margin: 0; color: #333333; font-weight: 700;">%s</h1>
		</div>
	`, data.OTPCode)

	// Add the OTP code display to content for HTML rendering
	content = append(content, otpCodeDisplay)

	options := SendOptions{
		ToEmail: data.ToEmail,
		ToName:  data.ToName,
		Subject: "Your move-pilotot Verification Code",
		TemplateData: ActionEmailTemplateData{
			ReceiverName: data.ToName,
			Title:        "Confirm Your Email Address",
			Content:      content,
		},
	}

	return c.Send(options)
}
