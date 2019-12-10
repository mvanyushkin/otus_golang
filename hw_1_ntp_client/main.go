package main

import (
	"errors"
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func main() {
	fmt.Println("The App has started.")
	currentTime, err := getNtpTime()
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Unable to get the value, occurred an exception: %v", err.Error()))
		os.Exit(-1)
		return
	}
	fmt.Println(currentTime)
	os.Exit(0)
}

func getNtpTime() (time.Time, error) {
	for index := 0; index < 4; index++ {
		currentHost := ntpHostsList[index]
		fmt.Println("We are asking the host %v", currentHost)
		response, err := ntp.Time(currentHost)
		if err != nil {
			fmt.Fprintln(os.Stderr, "The host %v doesn't work, has happened an exception %v", currentHost, err.Error())
			continue
		}

		return response, nil
	}

	return time.Time{}, errors.New("Unable to get the exact time cause there are no available NTP servers")
}

var ntpHostsList = []string{
	"0.beevik-ntp.pool.ntp.org",
	"0.ru.pool.ntp.org",
	"1.ru.pool.ntp.org",
	"2.ru.pool.ntp.org",
	"3.ru.pool.ntp.org",
}
