FROM golang:alpine

# Set the working directory inside the container
WORKDIR /app/backend

# Définissez la variable CGO_ENABLED
ENV CGO_ENABLED=1

COPY go.mod go.sum ./

RUN go mod download

RUN apk add --no-cache gcc musl-dev

# Copy the necessary files from the host to the container
COPY . /app/backend

RUN go build -o main .

CMD ["./main"]