# Dockerfile
ARG ARCH
FROM ${ARCH}/golang:1.15 AS builder
ENV CGO_ENABLED=0
WORKDIR /builder
COPY . .
RUN make build

FROM --platform=${ARCH} alpine
WORKDIR /app
COPY --from=builder /builder/weightbot .
COPY --from=builder /builder/weight_tracker_sample.csv weight_tracker.csv
CMD ["/bin/sh"]
