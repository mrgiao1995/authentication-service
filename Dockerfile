# This Dockerfile are use for build and publish app in a server like k8s, etc
FROM golang:1.20.5-alpine3.18 as build-stage

# If using priate repo
# ARG PAT
# ENV GOPRIVATE=github.com/<org>

# Port of app
ENV PORT=3000

WORKDIR /app

COPY go.mod .
COPY go.sum .

# If needed git
# RUN apk update && \
#     apk add git
# If using priate repo
# RUN git config --global url."https://pat:$PAT@github.com/<org>".insteadOf "https://github.com/<org>"

# Install dependencies
RUN go mod download

# Copy source
COPY . . 

# If main are not in root
# WORKDIR <path_to_main> 

# Build
RUN go build -o ./app .

# Create image to run
FROM scratch

# Copy builded program
# Change the source path to correct source if needed COPY --from=build-stage /app/<path_to_main>/app /app
COPY --from=build-stage /app/app /app
# Add certificates
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE $PORT

CMD ["./app"]