# Objective
Offer a api to convert html to pdf via POST method

# Pre-condition
1. Install https://github.com/wkhtmltopdf/wkhtmltopdf
   1. sudo apt-get update && apt-get install wkhtmltopdf -y
2. docker build . -t pdf-service && docker run --rm -p 8000:8000 pdf-service