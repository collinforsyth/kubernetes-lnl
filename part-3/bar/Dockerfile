# syntax = docker/dockerfile:1.2
FROM 605210442326.dkr.ecr.us-east-1.amazonaws.com/hinge-builder:v1.14 as builder

ADD . /src

WORKDIR /src

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /go/bin/app

FROM 605210442326.dkr.ecr.us-east-1.amazonaws.com/hinge-runner:v1.14_3.12.2
COPY --from=builder /go/bin/app /app

USER api

EXPOSE 8080

CMD ["/app"]
