package main

import "pome-email/email"

type serverT struct{}
var serverIns email.EmailServiceServer = &serverT{}
