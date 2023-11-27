fROM golang:1.21.4

RUN apt update
RUN apt -y upgrade
RUN apt install -y make
RUN apt install -y protobuf-compiler

WORKDIR /home/app

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN make build

CMD ["bin/server"]