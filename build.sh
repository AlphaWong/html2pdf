#!/bin/sh 

apt-get -qq update
apt-get -qq install wget build-essential unzip
wget -q -P /tmp/temp/ https://downloads.wkhtmltopdf.org/0.12/0.12.5/wkhtmltox_0.12.5-1.stretch_amd64.deb
wget -q -P /tmp/temp/ https://noto-website.storage.googleapis.com/pkgs/Noto-unhinted.zip
dpkg -i /tmp/temp/wkhtmltox_0.12.5-1.stretch_amd64.deb > /dev/null
apt -qq install -f -y
dpkg -i /tmp/temp/wkhtmltox_0.12.5-1.stretch_amd64.deb > /dev/null
mkdir -p /usr/share/fonts/Noto-unhinted
unzip /tmp/temp/Noto-unhinted.zip -d /usr/share/fonts/Noto-unhinted/
fc-cache -fv