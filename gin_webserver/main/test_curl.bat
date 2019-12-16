#curl -d "user=manu&password=123" http://localhost:8080/loginForm

>curl -X POST http://localhost:8000/uploadone -F "file=@C:\Users\13298\Desktop\WIFI.txt" -H "Content-Type: multipart/form-data"

curl -X POST http://localhost:8000/upload -F "upload[]=@C:\Users\13298\Desktop\WIFI.txt" -H "Content-Type: multipart/form-data"
#curl -X POST http://localhost:8000/upload -F "upload[]=@‪C:\Users\13298\Desktop\WIFI.txt" -H "Content-Type: multipart/form-data"