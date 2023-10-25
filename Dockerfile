# syntax=docker/dockerfile:1

FROM golang:1.21 AS BUILD

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /spyd


FROM scratch
COPY --from=BUILD /spyd /spyd
COPY files files


CMD [ "/spyd" ]