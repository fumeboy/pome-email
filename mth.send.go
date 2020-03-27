package main

import (
	"context"
	"fmt"
	"github.com/go-gomail/gomail"
	"pome-email/email"
)

func (this *serverT) Send(ctx context.Context, r*email.SendRequest)(resp*email.PomeAsyncResp, err error){
	fmt.Println("method: Add")
	gm := gomail.NewMessage()
	gm.SetHeader("From", config.Smtp.User)
	gm.SetHeader("To", r.To...)
	gm.SetHeader("Subject", r.Subject)
	gm.SetBody(r.ContentType, r.Body)

	d := gomail.NewDialer(config.Smtp.Host, config.Smtp.Port, config.Smtp.User, config.Smtp.Password)
	err = d.DialAndSend(gm)

	return nil, err
}
