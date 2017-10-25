FROM bash:4
COPY fsw /fsw
EXPOSE 8080
ENTRYPOINT ["/fsw"]

