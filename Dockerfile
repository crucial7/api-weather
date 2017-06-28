FROM golang:1.8

RUN curl https://glide.sh/get | sh

COPY . /go/src/github.com/crucial7/api-weather
WORKDIR /go/src/github.com/crucial7/api-weather

RUN chmod +x build.sh
RUN ./build.sh

CMD ["./index"]
