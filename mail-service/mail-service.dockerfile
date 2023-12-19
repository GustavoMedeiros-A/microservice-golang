FROM alpine:latest

RUN mkdir /app

# COPY the template and the app to the docker image
COPY mailerApp /app
COPY templates /templates

CMD [ "/app/mailerApp" ]