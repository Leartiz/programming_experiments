FROM golang:1.22.3-alpine3.19 AS build

LABEL description="Test Deploy n23"
LABEL version=v1.00

# ?
ENV GOOS linux

# ?
ENV CGO_ENABLED 0

RUN apk update --no-cache

# The --no-cache option allows to not cache the index locally,
#   which is useful for keeping containers small.
RUN apk add --no-cache make vim

WORKDIR /n23
COPY . .

RUN rm Makefile.windows || return 0
RUN rm main.exe || return 0

RUN mv Makefile.linux Makefile
RUN make build

FROM alpine

RUN apk update --no-cache
RUN apk add --no-cache make vim

WORKDIR /n23
COPY --from=build /n23/main /n23/main
COPY --from=build /n23/Makefile /n23/Makefile

EXPOSE 54545/tcp

CMD make only-run