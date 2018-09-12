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
  -o app \
  && mv ./app /go/bin/app

FROM alpine:latest
ENV MAX_SIZE=20
COPY --from=build /go/bin/app /
RUN mkdir /pdf
RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories && \
    echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories && \
    echo "http://dl-cdn.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories && \
    apk add --no-cache wkhtmltopdf font-noto curl fontconfig && \
    curl -O https://noto-website.storage.googleapis.com/pkgs/Noto-unhinted.zip && \
    mkdir -p /usr/share/fonts/Noto-unhinted && \
    unzip Noto-unhinted.zip -d /usr/share/fonts/Noto-unhinted/ && \
    rm Noto-unhinted.zip && \
    fc-cache -fv 
EXPOSE 8000
CMD ["/app"]
