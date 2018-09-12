#!/bin/sh 

echo "http://dl-cdn.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories && \
echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories && \
echo "http://dl-cdn.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories && \
apk add --no-cache wkhtmltopdf font-noto curl fontconfig && \
curl -O https://noto-website.storage.googleapis.com/pkgs/Noto-unhinted.zip && \
mkdir -p /usr/share/fonts/Noto-unhinted && \
unzip Noto-unhinted.zip -d /usr/share/fonts/Noto-unhinted/ && \
rm Noto-unhinted.zip && \
fc-cache -fv 
