FROM alpine
RUN apk --no-cache add curl
ADD  supply /supply
ENTRYPOINT [ "/supply" ]