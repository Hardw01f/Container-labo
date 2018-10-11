#practice01

#env
Mac book
GOOS=darwin,GOARCH=amd64

Run go build using crosscomplie for MacOS

copy binary to local from container

run binary on local


## How to 

```
$ docker build -t practice01 .
$ docker run practice01
$ docker cp CONTAINER_ID:/go/src/github.com/Laughingkitten/practice01/hello ./

$ ls
$ ./hello
```
