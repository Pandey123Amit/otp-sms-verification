# 📲 OTP Verification API using Twilio (Go/Gin)

A lightweight and secure OTP (One-Time Password) verification API built in **Go**, using **Gin** as the web framework and **Twilio** for SMS delivery.

## 🚀 Features

- Send OTP to a user's phone number via SMS
- Verify OTP for user authentication
- Input validation using `go-playground/validator`
- Error handling and structured JSON responses

## 🛠️ Tech Stack

- Go (Golang)
- Gin (Web Framework)
- Twilio (SMS API)
- Go Validator (Field validation)
- Delve (for debugging, optional)

## 📦 Installation

```bash
git clone https://github.com/Pandey123Amit/otp-sms-verification.git
cd otp-sms-verification
go mod tidy


TWILIO_ACCOUNT_SID=your_account_sid
TWILIO_AUTH_TOKEN=your_auth_token
TWILIO_PHONE_NUMBER=+1234567890


send Otp

{
  "phoneNumber": "+918986605695"
}


verifyotp
{
  "user": {
    "phoneNumber": "+918986605695"
  },
  "code": "018315"
}
