FROM iron/go:dev

MAINTAINER Vinicius Luiz <viniciuslcp97@gmail.com>

WORKDIR /app

ENV SRC_DIR=/go/src/practice-ws

ENV WS_CONFIG_DIR=$SRC_DIR/config 

ADD . $SRC_DIR

RUN cd $SRC_DIR; go build -o practice-ws; cp practice-ws /app/

ENTRYPOINT ["./practice-ws"]