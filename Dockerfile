FROM golang:1.16-alpine AS Build

# Prepare dir
RUN mkdir /build
COPY . /build
WORKDIR /build

# Install package
RUN go install ./main.go && go mod tidy

# Build
RUN go build -o mgmt main.go

FROM alpine:3.13 AS Prod

# Prepare dir
RUN mkdir /app
COPY ./initial.env /app/.env

# Copy binary
COPY --from=Build /build/mgmt /app

# Change dir
WORKDIR /app

# Grant permission
RUN chmod +x ./mgmt

# Entrypoint
ENTRYPOINT [ "./mgmt" ]

