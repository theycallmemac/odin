#!/usr/bin/env python3
from smtplib import SMTP
from email.mime.text import MIMEText
from requests import Session
import pyodin

def sendMail(msg, fromMail, toMail):
    server = SMTP('smtp.gmail.com:587')

    server.ehlo()
    server.starttls()
    server.login('your.email@gmail.com', 'gmail app password')
    server.sendmail(fromMail, [toMail], msg.as_string())
    server.quit()

def getURL(odin):
    session = Session()
    URL = "https://en.wikipedia.org/wiki/Special:Random"
    response = session.get(URL)

    odin.watch("URL", response.url)

    msg = MIMEText(response.url)

    return msg

def main():
    odin = pyodin.Odin(config="wiki.yml")

    fromMail = "odin@localhost"
    toMail = "your.email@gmail.com"

    odin.watch("Sender", fromMail)
    odin.watch("Receiver", toMail)

    msg = getURL(odin)
    msg['Subject'] = "Your Daily Wikipedia Article"
    msg['From'] = fromMail
    msg['To'] = toMail

    sendMail(msg, fromMail, toMail)

if __name__ == "__main__":
    main()

