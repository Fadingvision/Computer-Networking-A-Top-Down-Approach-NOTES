package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
)

// func main() {
// 	// Set up authentication information.
// 	auth := smtp.PlainAuth("", "fuyigekua@foxmail.com", "password", "imap.exmail.qq.com")

// 	// Connect to the server, authenticate, set the sender and recipient,
// 	// and send the email all in one step.
// 	to := []string{"chenxingyu01@xiaoduotech.com"}
// 	msg := []byte("To: chenxingyu01@xiaoduotech.com\r\n" +
// 		"Subject: smtp client! \r\n" +
// 		"\r\n" +
// 		"This is the email body.\r\n")
// 	err := smtp.SendMail("imap.exmail.qq.com:993", auth, "fuyigekua@foxmail.com", to, msg)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// with ssl
func main() {
	from := mail.Address{"", "fuyigekua@foxmail.com"}
	to := mail.Address{"", "chenxingyu01@xiaoduotech.com"}
	subj := "This is the email subject"
	body := "This is an example body.\n With two lines."

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	// servername := "imap.exmail.qq.com:993"
	servername := "smtp.exmail.qq.com:465"

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", "fuyigekua@foxmail.com", "password", host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()

}
