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

FROM gcr.io/google-appengine/debian9:latest
ENV MAX_SIZE=20
COPY --from=build /go/bin/app /
COPY ./build.sh .
RUN mkdir /pdf
RUN sh ./build.sh
RUN wkhtmltopdf --version
EXPOSE 80
CMD ["/app"]
