FROM debian:stable-slim

COPY simple-go-server /bin/simple-go-server

CMD [ "/bin/simple-go-server" ]
