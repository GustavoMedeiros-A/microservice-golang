FROM alpine:latest

RUN mkdir /app

COPY authApp /app

# COPY --from=builder /app/authApp /app 

CMD [ "/app/authApp" ]