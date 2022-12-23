import json
import smtplib
from email.message import EmailMessage
from config.settings import settings
import datetime


def logging():
    time = datetime.datetime.now()
    return time


def TerminalLog(time=logging()):
    return "%s %s:%s:%s" % (time.date(), time.hour, time.minute, time.second)


def SendingMail(ch, method, properties, body):
    raw = json.loads(body.decode())
    target = raw.get("email")
    token = raw.get("token")
    timeStart = logging()
    print(f"{TerminalLog(timeStart)}\t\t\t\t{target}\t\t")
    Sender = settings.MAIL_NAME
    Password = settings.MAIL_PASSWORD
    To = target
    try:
        server = smtplib.SMTP('smtp.gmail.com', 587)
        server.starttls()
        server.login(
            Sender,
            Password,
        )
        subject = "Account Verification Token"
        body = f"""
            It's Your Token For Account Verification.
            Don't share this token to anyone.
            Your Token:
            {token}
        """
        em = EmailMessage()
        em["From"] = Sender
        em["To"] = To
        em["Subject"] = subject
        em.set_content(body)
        server.sendmail(Sender, To, em.as_string())
        server.close()
    except Exception as E:
        print(E)
