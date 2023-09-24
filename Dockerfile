FROM debian:stable-slim

COPY exhibition /bin/exhibition

ENV PORT=8000

# Install and update ca-certificates
RUN apt-get update && \
    apt-get install -y ca-certificates && \
    update-ca-certificates

CMD [ "/bin/exhibition" ]