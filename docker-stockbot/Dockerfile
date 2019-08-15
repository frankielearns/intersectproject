FROM golang:alpine
COPY stockbot .
RUN chmod +x stockbot
COPY docker-entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/docker-entrypoint.sh
ENTRYPOINT ["docker-entrypoint.sh"]

CMD ["./stockbot"]

