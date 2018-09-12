# /bin/bash

docker build . -t pdf-service && docker run -it --rm -p 8000:8000 pdf-service