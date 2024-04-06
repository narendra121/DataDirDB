FROM golang:alpine

RUN mkdir /goapp

WORKDIR /goapp

COPY . /goapp

RUN go build  -o mygoapp  cmd/go-app/main.go 

COPY . .
EXPOSE 8080
CMD [ "./mygoapp","-env","docker" ]