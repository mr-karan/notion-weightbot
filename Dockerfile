# Dockerfile
ARG ARCH
ARG PLATFORM
FROM ${ARCH}/golang:1.15 AS builder
ENV CGO_ENABLED=0
WORKDIR /builder
COPY . .
RUN make build

FROM --platform=${PLATFORM} alpine
WORKDIR /app
COPY --from=builder /builder/weightbot .
COPY --from=builder /builder/weight_tracker_sample.csv weight_tracker.csv
CMD ["./weightbot"]
