package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path"
)

func Log_me_daddy() {
	my_log, err := syslog.New(syslog.LOG_SYSLOG, "not happy")
	if err != nil {
		log.Fatalln("this is bad")
	}
	log.Println("got here")
	log.SetOutput(my_log)
	log.Print("now that's good 🥹")
}

func Log_me_custom() {

	ME_OL_LOG_FILE := path.Join(os.TempDir(), "loggergang.log")
	f, err := os.OpenFile(ME_OL_LOG_FILE, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("GETTAATTTYA HEYA")
	}
	defer f.Close()
	me_ol_logger := log.New(f, "south to all my boss bois", log.LstdFlags)
	me_ol_logger.Println("come and rescue meeee take me out the club")
}
