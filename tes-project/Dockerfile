FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --no-cache --update add ca-certificates tzdata && \
    mkdir /app
    
WORKDIR /app

EXPOSE 8005

COPY engine /app

CMD /app/engine