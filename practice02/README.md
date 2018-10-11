#practice00

build Golang server on container

## How to 

```
$ docker build -t practice02 .
$ docker run -d -p 8080:8080 practice02
```

ブラウザで
http://localhost:8080
http://localhost:8080/hey
にアクセス
