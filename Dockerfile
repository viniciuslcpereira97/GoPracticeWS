# Imagem GO utilizada
FROM iron/go:dev

MAINTAINER Vinicius Luiz <viniciuslcp97@gmail.com>

# Diretório atual
WORKDIR /app

# Exporta uma variáveis de ambiente que serão utilizadas
ENV SRC_DIR=/go/src/GoPracticeWS
ENV WS_CONFIG_DIR=$SRC_DIR/config 

# Adiciona arquivos do diretório do host à imagem Docker
ADD . $SRC_DIR

# Entra no diretório SRC - constrói o arquivo binário da aplicação em GO - copia o binário para o Working dir
RUN cd $SRC_DIR; go build -o GoPracticeWS; cp GoPracticeWS /app/

# Executa o binário ao montar a imagem
ENTRYPOINT ["./GoPracticeWS"]
