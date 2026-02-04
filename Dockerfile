# Dockerfile
# Stage 1: Build the Go application
# Use a specific Go version for reproducibility. Alpine is used for a smaller builder image.
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go module files and download dependencies.
# This leverages Docker's layer caching, so dependencies are only re-downloaded if go.mod or go.sum change.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application as a static binary.
# CGO_ENABLED=0 creates a static binary without any C dependencies.
# -ldflags "-w -s" strips debug information, reducing the binary size.
# The output is placed in the root directory for easy copying to the next stage.
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /bpip-core cmd/bpip-core/main.go

# Stage 2: Create the final minimal image
# Use a distroless static image for the smallest possible footprint and attack surface.
# It contains only the application, its runtime dependencies, and basic OS files (like CA certificates).
FROM gcr.io/distroless/static-debian11 AS final

# Add OCI standard labels to link the image to its source and provide metadata.
LABEL org.opencontainers.image.source=https://github.com/dresbach-records/bpip-core
LABEL org.opencontainers.image.description="BPIP Core — Economic Participation Protocol Engine"
LABEL org.opencontainers.image.licenses=Apache-2.0

# Copy the static binary from the builder stage
COPY --from=builder /bpip-core /

# Copy the default configuration file.
# The application will use this unless a configuration is mounted over it.
COPY config.yml /

# Distroless images run as a non-root user by default ('nonroot').
# Explicitly setting the user is a security best practice.
USER nonroot:nonroot

# Set the entrypoint for the container. This is the command that will be executed when the container starts.
ENTRYPOINT ["/bpip-core"]
