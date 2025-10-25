package email

import (
	"fmt"
	"strings"
)

type ActionEmailTemplateData struct {
	ReceiverName        string
	Title               string
	Content             []string
	ListItems           []string
	BottomContent       []string
	PrimaryActionText   string
	PrimaryActionURL    string
	SecondaryActionText string
	SecondaryActionURL  string
}

func GenerateEmailTemplate(data ActionEmailTemplateData) string {
	var sb strings.Builder

	for _, p := range data.Content {
		sb.WriteString(fmt.Sprintf(`<p style="margin: 0 0 16px 0; color: #444; line-height: 1.6; font-size: 14px;">%s</p>`, p))
	}
	paragraphs := sb.String()
	sb.Reset()

	if len(data.ListItems) > 0 {
		sb.WriteString("<ul style='margin: 16px 0; padding-left: 20px;'>")
		for _, item := range data.ListItems {
			sb.WriteString(fmt.Sprintf(`<li style="margin-bottom: 8px; color: #444; line-height: 1.5; font-size: 14px;">%s</li>`, item))
		}
		sb.WriteString("</ul>")
	}
	listHTML := sb.String()
	sb.Reset()

	for _, p := range data.BottomContent {
		sb.WriteString(fmt.Sprintf(`<p style="margin: 0 0 12px 0; color: #666; font-size: 13px; line-height: 1.5;">%s</p>`, p))
	}
	bottomParagraphs := sb.String()
	sb.Reset()

	if data.SecondaryActionText != "" && data.SecondaryActionURL != "" {
		sb.WriteString(fmt.Sprintf(`
			<a href="%s" style="color: #333333; text-decoration: none; border: 1.5px solid #333333; padding: 8px 20px; border-radius: 6px; font-size: 0.875rem; display: inline-block; font-weight: 500;">%s</a>`,
			data.SecondaryActionURL, data.SecondaryActionText))
	}
	if data.PrimaryActionText != "" && data.PrimaryActionURL != "" {
		sb.WriteString(fmt.Sprintf(`
			<a href="%s" style="color: #ffffff; text-decoration: none; background-color: #000000; padding: 8px 20px; border-radius: 6px; margin-left: 12px; font-size: 0.875rem; display: inline-block; font-weight: 500;">%s</a>`,
			data.PrimaryActionURL, data.PrimaryActionText))
	}
	buttons := sb.String()

	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link href="https://fonts.googleapis.com/css2?family=Space+Grotesk:wght@400;500;600;700&display=swap" rel="stylesheet">
  <title>Email</title>
</head>
<body style="margin: 0; padding: 0; font-family: 'Space Grotesk', sans-serif; background-color: #f5f5f5;">
  <table width="100%%" cellpadding="0" cellspacing="0" style="background-color: #f8f8f8;">
    <tr><td align="center">
      <table style="max-width: 600px; width: 100%%; margin: 0; background-color: #ffffff;">
        <!-- Thin black header strip -->
        <tr>
          <td style="background-color: #000000; height: 4px;"></td>
        </tr>
        <!-- Logo section -->
        <tr>
          <td style="padding: 24px 40px; border-bottom: 1px solid #e5e5e5;">
            <h1 style="color: #000; margin: 0; font-size: 14px; font-weight: 600; letter-spacing: 0.5px;">move-pilot</h1>
          </td>
        </tr>
        <!-- Email content -->
        <tr><td style="padding: 32px 40px;">
          <p style="font-size: 13px; color: #666; margin: 0 0 4px 0;">Hi %s,</p>
          <h2 style="font-weight: 500; font-size: 20px; color: #000; margin: 0 0 24px 0;">%s</h2>
          %s
          %s
          %s
          <table style="margin: 32px 0;">
            <tr><td>%s</td></tr>
          </table>
          <p style="color: #666; margin: 24px 0 0 0; font-size: 13px;">Best regards,</p>
          <p style="color: #666; margin: 4px 0 0 0; font-size: 13px; font-weight: 500;">The move-pilot Team</p>
        </td></tr>
        <!-- Footer -->
        <tr>
          <td style="padding: 20px 40px; border-top: 1px solid #e5e5e5; text-align: center;">
            <p style="color: #999; margin: 0; font-size: 11px;">© 2025 move-pilotot. All rights reserved.</p>
          </td>
        </tr>
        <!-- Thin black footer strip -->
        <tr>
          <td style="background-color: #000000; height: 4px;"></td>
        </tr>
      </table>
    </td></tr>
  </table>
</body>
</html>
`, data.ReceiverName, data.Title, paragraphs, listHTML, bottomParagraphs, buttons)
}

func GeneratePlainTextEmail(data ActionEmailTemplateData) string {
	var plainTextBuilder strings.Builder
	plainTextBuilder.WriteString(fmt.Sprintf("Hi %s,\n\n", data.ReceiverName))
	plainTextBuilder.WriteString(fmt.Sprintf("%s\n\n", data.Title))

	for _, p := range data.Content {
		plainTextBuilder.WriteString(fmt.Sprintf("%s\n", p))
	}

	if len(data.ListItems) > 0 {
		plainTextBuilder.WriteString("\n")
		for _, item := range data.ListItems {
			plainTextBuilder.WriteString(fmt.Sprintf("- %s\n", item))
		}
	}

	if len(data.BottomContent) > 0 {
		plainTextBuilder.WriteString("\n")
		for _, p := range data.BottomContent {
			plainTextBuilder.WriteString(fmt.Sprintf("%s\n", p))
		}
	}

	if data.PrimaryActionText != "" && data.PrimaryActionURL != "" {
		plainTextBuilder.WriteString(fmt.Sprintf("\n%s: %s\n", data.PrimaryActionText, data.PrimaryActionURL))
	}

	if data.SecondaryActionText != "" && data.SecondaryActionURL != "" {
		plainTextBuilder.WriteString(fmt.Sprintf("%s: %s\n", data.SecondaryActionText, data.SecondaryActionURL))
	}

	plainTextBuilder.WriteString("\nThanks,\nThe move-pilotot Team\n\n© 202move-pilotilot. All rights reserved.")

	return plainTextBuilder.String()
}
