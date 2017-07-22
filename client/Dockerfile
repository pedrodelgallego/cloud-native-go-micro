FROM alpine

EXPOSE 8080
ENV CONSUL_HTTP_ADDR localhost:8500

COPY client /app/client
RUN chmod +x /app/client


ENTRYPOINT /app/client
