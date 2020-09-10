FROM golang:1.13
RUN mkdir /backend-project
ADD . /backend-project
WORKDIR /backend-project
RUN go install
RUN go build -o main .
CMD ["/backend-project/main"]