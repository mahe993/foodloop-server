FROM postgres:16.1-alpine

ENV GID=1000
ENV UID=1000

RUN addgroup -g $GID owner
RUN adduser -S -u $UID owner owner

USER owner