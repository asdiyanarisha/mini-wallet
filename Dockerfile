# Build Stage
FROM golang:1.19-bullseye AS BuildStage
ENV TZ=Asia/Jakarta
ENV DEBIAN_FRONTEND noninteractive

ENV GO111MODULE=on

COPY . /app/
WORKDIR /app/

EXPOSE 8080

RUN CGO_ENABLED=0 go build -v .


# Deploy Stage
FROM debian:bullseye-slim AS DeployStage

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
 && rm -rf /var/lib/apt/lists/*

ENV TZ=Asia/Jakarta
ENV DEBIAN_FRONTEND noninteractive

COPY --from=BuildStage ./app/mini-wallet ./mini-wallet
COPY --from=BuildStage ./app/data ./data

EXPOSE 8080

CMD ["./mini-wallet"]
