import  yagmail

args = {
    "user":"xx @qq.com",
    "password":" x",
    "host":"smtp.qq.com",
    "port":"465"
} 

reciver = ["wang115byby@163.com","3453667697@qq.com"]
email = yagmail.SMTP(**args)

email.send(to=reciver, subject="hello",contents="i am van")
