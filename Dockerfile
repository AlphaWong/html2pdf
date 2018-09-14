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

FROM debian:latest
ENV MAX_SIZE=20
COPY --from=build /go/bin/app /
RUN mkdir /pdf
RUN apt-get update && apt-get install wget build-essential unzip -y && \
  wget -P /tmp/temp/ https://downloads.wkhtmltopdf.org/0.12/0.12.5/wkhtmltox_0.12.5-1.stretch_amd64.deb && \
  wget -P /tmp/temp/ https://noto-website.storage.googleapis.com/pkgs/Noto-unhinted.zip && \
  dpkg -i /tmp/temp/wkhtmltox_0.12.5-1.stretch_amd64.deb || true && \
  apt install -f -y && \
  dpkg -i /tmp/temp/wkhtmltox_0.12.5-1.stretch_amd64.deb && \
  mkdir -p /usr/share/fonts/Noto-unhinted && \
  unzip /tmp/temp/Noto-unhinted.zip -d /usr/share/fonts/Noto-unhinted/ && \
  fc-cache -fv
RUN wkhtmltopdf --version
EXPOSE 8000
CMD ["/app"]
