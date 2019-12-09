package main

import (
	"errors"
	"log"
	"time"
)
import "github.com/beevik/ntp"

func main() {
	currentTime, _ := getNtpTime();
	log.Print(currentTime)
}

var ntpHostsList = []string{
	"0.beevik-ntp.pool.ntp.org",
	"0.ru.pool.ntp.org",
	"1.ru.pool.ntp.org",
	"2.ru.pool.ntp.org",
	"3.ru.pool.ntp.org",
}

func getNtpTime() (time.Time, error) {
	for index := 0; index < 4; index++ {
		address := ntpHostsList[index];
		response, err := ntp.Time(address);
		if err != nil {
			continue;
		}

		return response, nil;
	}

	return time.Time{}, errors.New("Unable to get the exact time cause there are no available NTP servers")
}
