FROM golang:1.21 as builder

# Copy the present working directory to our source directory in Docker.
# Change the current directory in Docker to our source directory.
COPY . /app
WORKDIR /app

# Build our application as a static build.
# The mount options add the build cache to Docker to speed up multiple builds.
RUN --mount=type=cache,target=/root/.cache/go-build \
	--mount=type=cache,target=/go/pkg \
	go build -ldflags '-s -w -extldflags "-static"' -o /app/choomba ./cmd/choomba

# Download the static build of Litestream directly into the path & make it executable.
# This is done in the builder and copied as the chmod doubles the size.
# TODO: Add compatibility for multiarch setups
ADD https://github.com/benbjohnson/litestream/releases/download/v0.3.13/litestream-v0.3.13-linux-amd64-static.tar.gz /tmp/litestream.tar.gz
RUN tar -C /usr/local/bin -xzf /tmp/litestream.tar.gz

# This starts our final image; based on alpine to make it small.
FROM alpine:3

# You can optionally set the replica URL directly in the Dockerfile.
# ENV REPLICA_URL=s3://BUCKETNAME/db

# Copy executable & Litestream from builder.
COPY --from=builder /app/choomba /app/choomba
COPY --from=builder /usr/local/bin/litestream /usr/local/bin/litestream

RUN apk add bash

# Create data directory (although this will likely be mounted too)
RUN mkdir -p /data

# Copy Litestream configuration file & startup script.
COPY etc/litestream.yml /etc/litestream.yml
COPY scripts/run.sh /scripts/run.sh

CMD [ "/scripts/run.sh" ]
