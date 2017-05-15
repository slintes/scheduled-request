FROM alpine:3.5

RUN apk --no-cache add ca-certificates

COPY ./scheduled-request-linux /go/bin/scheduled-request

ENV PATH="/go/bin:$PATH"
WORKDIR /go/bin

CMD scheduled-request -interval=${INTERVAL} -url=${URL} -method=${METHOD} -contentType=${CONTENTTYPE} -body=${BODY}