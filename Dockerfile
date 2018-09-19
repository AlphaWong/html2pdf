FROM golang:latest as build

ENV GO111MODULE=on
ENV MAX_SIZE=20

WORKDIR /go/src/github.com/AlphaWong/html2pdf
COPY . .

RUN go fmt ./...
RUN CGO_ENABLE=0 GOOS=linux \
  go build \
  -tags netgo \
  -installsuffix netgo,cgo \
  -v -a \
  -ldflags '-s -w -extldflags "-static"' \ 
  -o app

RUN apt-get -qq update && apt-get -qq install wget xz-utils
RUN wget -P /tmp/ https://github.com/upx/upx/releases/download/v3.95/upx-3.95-amd64_linux.tar.xz
RUN tar xvf /tmp/upx-3.95-amd64_linux.tar.xz -C /tmp
RUN mv /tmp/upx-3.95-amd64_linux/upx /go/bin

RUN upx --ultra-brute -qq app && \
  upx -t app && \
  mv ./app /go/bin/app

FROM gcr.io/google-appengine/debian9:latest
ENV MAX_SIZE=20
ENV LANG=en_US.UTF-8
ENV LC_ALL=en_US.UTF-8
COPY --from=build /go/bin/app /
COPY ./build.sh .
RUN mkdir /pdf
RUN sh ./build.sh
RUN wkhtmltopdf --version
EXPOSE 80
CMD ["/app"]
