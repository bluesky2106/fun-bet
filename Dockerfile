FROM golang:1.15.0-alpine3.12

# install important packages
RUN apk add --no-cache ca-certificates git bash

EXPOSE 8080

ENV GO111MODULE=on
ENV GIN_MODE=release

# Arguments
ARG REPO_DIR=$GOPATH/src/github.com/bluesky2106/fun-bet

# Create app directory in image filesystem 
RUN mkdir -p ${REPO_DIR}

WORKDIR $REPO_DIR

COPY . .

RUN CGO_ENABLED=0 go build -o server && \
    chmod u+x server

CMD ["./server"]