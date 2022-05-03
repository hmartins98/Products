FROM hugomartins98/techbrothers:grpc

ENV HOME /root

WORKDIR ${HOME}

RUN mkdir -p -m 0600 ~/.ssh && \
    ssh-keyscan -H github.com >> ~/.ssh/known_hosts
COPY ./.ssh/id_ed25519 ./.ssh/id_ed25519
RUN chmod 0600 ./.ssh/id_ed25519

WORKDIR /go/src

COPY ./github.com/hmartins98/Products ./Products

WORKDIR /go/src/Products

RUN go mod init Products
RUN go mod tidy
RUN go build -o main .

EXPOSE 5000

CMD ["./main"]