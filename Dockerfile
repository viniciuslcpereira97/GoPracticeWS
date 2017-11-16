FROM iron/go:dev

MAINTAINER Vinicius Luiz <viniciuslcp97@gmail.com>

WORKDIR /app

ENV SRC_DIR=/go/src/GoPracticeWS

ENV WS_CONFIG_DIR=$SRC_DIR/config 

ADD . $SRC_DIR

RUN cd $SRC_DIR; go build -o GoPracticeWS; cp GoPracticeWS /app/

ENTRYPOINT ["./GoPracticeWS"]