# docker build -t tink-rebooter .
# docker run -it --rm --privileged -v /:/host --pid host tink-rebooter
FROM golang:1.16 as builder

WORKDIR /code
COPY go.mod go.sum /code/
RUN go mod download

COPY . /code
RUN make build

FROM busybox

COPY --from=builder /code/bin/tink-reboot-linux-amd64 /tink-reboot
COPY entrypoint.sh entrypoint.sh

ENTRYPOINT [ "/entrypoint.sh" ]
