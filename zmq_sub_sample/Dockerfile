FROM golang

RUN mkdir -p /workspace
WORKDIR /workspace

# apt-get
RUN apt-get update
RUN apt-get install -y libczmq-dev libzmq3-dev

# source code
COPY . .

RUN go get

CMD go run ./main.go
