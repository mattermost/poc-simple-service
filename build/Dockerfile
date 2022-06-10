# This dockerfile is used to build Mattermost poc-simple-service
# A multi stage build, with golang used as a builder
# and gcr.io/distroless/static as runner
ARG GO_IMAGE=golang:1.18.3@sha256:5417b4917fa7ed3ad2678a3ce6378a00c95bfd430c2ffa39936fce55130b5f2c
# hadolint ignore=DL3006
FROM ${GO_IMAGE} as builder

#GO_BUILD_PLATFORMS holds the platforms that we will build the docker image against
ARG GO_BUILD_PLATFORMS=linux-amd64

# Setup directories structure and compile
COPY . /src
WORKDIR /src
RUN make go-build

# Shrink final image since we only need the rtcd binary
# and use distroless container image as runner for security
FROM gcr.io/distroless/static@sha256:d6fa9db9548b5772860fecddb11d84f9ebd7e0321c0cb3c02870402680cc315f as runner
COPY --from=builder /src/poc-simple-service /opt/service/bin/poc-simple-service

# We should refrain from running as privileged user
# Run as UID for nobody
USER 65534

WORKDIR /opt/service/bin
ENTRYPOINT ["./poc-simple-service"]