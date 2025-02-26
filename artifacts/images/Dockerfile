FROM golang:1.24 AS builder

WORKDIR /go/src/k8s.io/sample-controller
COPY ../../ .

RUN make build

FROM gcr.io/distroless/base-debian12
COPY --from=builder /go/src/k8s.io/sample-controller/_output/bin/sample-controller /
CMD ["/sample-controller"]
