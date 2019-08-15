FROM alpine:latest
RUN apk update && apk add bash
COPY stockbot .
RUN chmod +x stockbot

CMD ["./stockbot"]

