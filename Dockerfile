FROM golang:1-alpine

WORKDIR /go/src/github.com/polarpoint/kubernetes-admission-controller
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /polarpoint-kubernetes-admission-controller ./cmd/kubernetes-admission-controller/
CMD ["/polarpoint-kubernetes-admission-controller"]
