package storage

import (
	"fmt"
	"log"
	"net/smtp"
)

func bytesInUse(username string) int64 {
	return 0
}

const sender = "notifications@example.com"
const password = "correcthorsebatterystaple"
const hostname = "smtp.example.com"

const template = `Warning: you are using %d bytes of storage,
%d%% of your quota.`

func CheckQuota(username string)  {
	used := bytesInUse(username)
	const quota = 1000000000
	precent := 100 * used / quota
	if precent < 90 {
		return  // ok
	}
	msg := fmt.Sprintf(template, used, precent)
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+"587", auth, sender, []string{username}, []byte(msg))
	if err != nil {
		log.Printf("smt.SendMail(%s) failed: %s", username, err)
	}
}
