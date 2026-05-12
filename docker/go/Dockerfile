FROM golang:1.25

# Set the Current Working Directory inside the container
WORKDIR /app
ENV GOPROXY=direct
ENV GOSUMDB=off
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /main ./cmd
CMD ["/main"]
