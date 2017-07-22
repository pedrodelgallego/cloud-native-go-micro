FROM alpine

EXPOSE 8080

copy server /app/server
RUN chmod +x /app/server

ENV CONSUL_HTTP_ADDR localhost:8500

ENTRYPOINT /app/server
