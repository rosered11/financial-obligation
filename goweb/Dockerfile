FROM golang:1.16-buster

COPY . .

RUN apt-get update 

# RUN echo "====== Starting to Download Fabric ========" \
#     && curl -sSL https://bit.ly/2ysbOFE -o bootstrap.sh \
#     && chmod 755 ./bootstrap.sh \
#     && bash ./bootstrap.sh -ds 2.2.3 1.5.0

RUN echo "======= Copying the binaries to /usr/local/bin====" \
    && cp ./bin/*    /usr/local/bin

WORKDIR ./src

RUN go mod tidy

RUN go build

CMD ["./financial-webapi"]