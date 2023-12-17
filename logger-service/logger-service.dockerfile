FROM alpine:latest

RUN mkdir /app

COPY loggerServiceApp /app
# COPY --from=builder /app/loggerServiceApp /app 

CMD [ "/app/loggerServiceApp" ]