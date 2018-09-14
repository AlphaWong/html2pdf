[![pipeline status](https://gitlab.com/AlphaWong/html2pdf/badges/master/pipeline.svg)](https://gitlab.com/AlphaWong/html2pdf/pipelines)
[![coverage report](https://gitlab.com/AlphaWong/html2pdf/badges/master/coverage.svg)](https://alphawong.gitlab.io/html2pdf/coverage.html)

# Objective
Offer an API to convert html to pdf via POST method

# How
It wraps `wkhtmltopdf` for the html convert to pdf. The reason we do not use `weasyprint` is that weasyprint do not offer a suitable converted product. However, `wkhtmltopdf` always return a suitable result. Also, `chrome headless mode` do not have the feature we need such as footer and header.

# Run
```sh
./reload.sh
```

# Screenshot
## wkhtmltopdf
![alt wkhtmltopdf](https://i.imgur.com/nrH8RTV.png)
## weasyprint
![alt wkhtmltopdf](https://i.imgur.com/uEOf6eb.png)

# Postman collection
https://www.getpostman.com/collections/0e61ae04d5f54cb17a5a

# Curl
```
curl -X POST \
  http://127.0.0.1:8000/convert \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H 'Postman-Token: 18fad9e8-402b-435c-bd9a-173f3e200cc0' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -F file=@/home/alpha/Downloads/people.html \
  -F '--footer-left=[page]'
```

# Pre-condition
1. wkhtmltopdf for converting binary
2. Noto-unhinted.zip for i18n issue

# Issue
1. Noto-unhinted can resolve the cjk and thai character display issue.

# Reference
1. https://qiita.com/nju33/items/b80d92a4257edeb4b9a1
2. https://developers.google.com/web/updates/2017/04/headless-chrome#create_a_pdf_dom

# Credit
1. alan.tang@lalamove.com
1. christopher.chiu@lalamove.com
1. desmond.ho@lalamove.com
1. jack.tang@lalamove.com
1. simon.tse@lalamove.com
1. wachiu.siu@lalamove.com
