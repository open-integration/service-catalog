FROM alpine:3.9

RUN apk update && apk add --no-cache ca-certificates && apk upgrade

ARG NAME
ARG DIR

COPY ./$DIR/dist/$NAME .

RUN mv $NAME service

EXPOSE 80

ENV PORT=80

CMD "/service"