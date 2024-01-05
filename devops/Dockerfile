FROM golang:1.21

# Copyright
LABEL "maintainer"="Andriy Kandzyuba"
LABEL "org"="openexchange.io"
LABEL "license"="Apache 2.0"

# Add new group and user
RUN addgroup otel && useradd otel -g otel -d /home/otel

# Set new user as default
USER otel:otel

# Set working directory
WORKDIR /home/otel

# Download Go modules
COPY ../go.mod ../go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
COPY ../*.go ./
COPY ../pkg ./pkg

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o ./openexchange-otel-go

# Epose the port to listen on
EXPOSE 8080

# Start the application
CMD [ "/home/otel/openexchange-otel-go" ]