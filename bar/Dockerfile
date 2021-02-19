# syntax = docker/dockerfile:1.2
FROM 605210442326.dkr.ecr.us-east-1.amazonaws.com/hinge-builder:v1.14 as builder

ADD . /src

WORKDIR /src

RUN go build -o /go/bin/app

FROM 605210442326.dkr.ecr.us-east-1.amazonaws.com/hinge-runner:v1.14_3.12.2
COPY --from=builder /go/bin/app /app

EXPOSE 8080

CMD ["/app"]
