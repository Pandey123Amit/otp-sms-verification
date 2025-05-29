package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func envACCOUNTSID() string {
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		log.Fatal("Error lodgin in .env")
	}
	return os.Getenv("TWILID_ACCOUNT_SID")
}
func envACCOUNTAUTH() string {
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		log.Fatal("Error lodgin in .env")
	}
	return os.Getenv("TWILIO_AUTHTOKEN")
}
func envACCOUNTSERVICE() string {
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		log.Fatal("Error lodgin in .env")
	}
	return os.Getenv("TWILID_SERVICE_ID")
}
