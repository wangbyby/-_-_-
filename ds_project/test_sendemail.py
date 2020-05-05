import  yagmail

args = {
    "user":"xx @qq.com",
    "password":" x",
    "host":"smtp.qq.com",
    "port":"465"
} 

reciver = ["wwww@163.com","www@qq.com"]
email = yagmail.SMTP(**args)

email.send(to=reciver, subject="hello",contents="i am van")
