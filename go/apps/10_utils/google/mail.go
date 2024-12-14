package google

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"

	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

// GmailClient はGmail APIクライアントを表します
type GmailClient struct {
	Service *gmail.Service
}

// NewGmailClient は新しいGmail APIクライアントを作成します
// client: 認証済みのHTTPクライアント
func NewGmailClient(client *http.Client) (*GmailClient, error) {
	srv, err := gmail.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
		return nil, err
	}

	return &GmailClient{Service: srv}, nil
}

// SendEmail はGmail APIを使用してメールを送信します
// from: 送信者のメールアドレス
// to: 受信者のメールアドレスのスライス
// cc: CCのメールアドレスのスライス
// bcc: BCCのメールアドレスのスライス
// subject: メールの件名
// body: メールの本文
// isHTML: 本文がHTMLかどうかを示すブール値
// attachments: 添付ファイルのパスのスライス
func (client *GmailClient) SendEmail(from string, to []string, cc []string, bcc []string, subject string, body string, isHTML bool, attachments []string) error {
	var message gmail.Message

	// 受信者、CC、BCC、件名を設定
	emailTo := "To: " + formatEmailList(to) + "\r\n"
	emailCc := ""
	if len(cc) > 0 {
		emailCc = "Cc: " + formatEmailList(cc) + "\r\n"
	}
	emailBcc := ""
	if len(bcc) > 0 {
		emailBcc = "Bcc: " + formatEmailList(bcc) + "\r\n"
	}
	subject = "Subject: " + subject + "\r\n"
	mime := "MIME-version: 1.0;\n"

	var msg bytes.Buffer
	msg.WriteString(emailTo)
	msg.WriteString(emailCc)
	msg.WriteString(emailBcc)
	msg.WriteString(subject)
	msg.WriteString(mime)

	// 添付ファイルがある場合の処理
	if len(attachments) > 0 {
		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)
		mime = fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n\n", writer.Boundary())
		msg.WriteString(mime)

		// メール本文を追加
		part, err := writer.CreatePart(textproto.MIMEHeader{
			"Content-Type": {"text/plain; charset=UTF-8"},
		})
		if err != nil {
			return fmt.Errorf("unable to create email part: %v", err)
		}
		part.Write([]byte(body))

		// 添付ファイルを追加
		for _, attachment := range attachments {
			file, err := os.Open(attachment)
			if err != nil {
				return fmt.Errorf("unable to open attachment: %v", err)
			}
			defer file.Close()

			part, err := writer.CreateFormFile("attachment", filepath.Base(attachment))
			if err != nil {
				return fmt.Errorf("unable to create attachment part: %v", err)
			}
			_, err = io.Copy(part, file)
			if err != nil {
				return fmt.Errorf("unable to copy attachment: %v", err)
			}
		}

		writer.Close()
		message.Raw = base64.URLEncoding.EncodeToString(buf.Bytes())
	} else {
		// HTMLメールかプレーンテキストメールかを設定
		if isHTML {
			mime += "Content-Type: text/html; charset=\"UTF-8\";\n\n"
		} else {
			mime += "Content-Type: text/plain; charset=\"UTF-8\";\n\n"
		}
		msg.WriteString(mime)
		msg.WriteString("\n")
		msg.WriteString(body)
		message.Raw = base64.URLEncoding.EncodeToString(msg.Bytes())
	}

	// メールを送信
	_, err := client.Service.Users.Messages.Send("me", &message).Do()
	if err != nil {
		return fmt.Errorf("unable to send email: %v", err)
	}

	fmt.Println("Email sent successfully.")
	return nil
}

// formatEmailList はメールアドレスのリストをカンマ区切りの文字列にフォーマットします
// emails: メールアドレスのスライス
func formatEmailList(emails []string) string {
	return strings.Join(emails, ", ")
}

// GetEmails はGmail APIを使用してメールを取得します
// query: 検索クエリ
// maxResults: 最大取得件数
func (client *GmailClient) GetEmails(query string, maxResults int64) ([]*gmail.Message, error) {
	user := "me"
	req := client.Service.Users.Messages.List(user).Q(query).MaxResults(maxResults)
	res, err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve emails: %v", err)
	}

	var messages []*gmail.Message
	for _, m := range res.Messages {
		msg, err := client.Service.Users.Messages.Get(user, m.Id).Do()
		if err != nil {
			return nil, fmt.Errorf("unable to retrieve email: %v", err)
		}
		messages = append(messages, msg)
	}

	return messages, nil
}

// DeleteEmail はGmail APIを使用してメールを削除します
// messageID: 削除するメールのID
func (client *GmailClient) DeleteEmail(messageID string) error {
	user := "me"
	err := client.Service.Users.Messages.Delete(user, messageID).Do()
	if err != nil {
		return fmt.Errorf("unable to delete email: %v", err)
	}

	fmt.Println("Email deleted successfully.")
	return nil
}
