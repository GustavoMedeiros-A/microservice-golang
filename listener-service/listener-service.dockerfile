FROM alpine:latest

RUN mkdir /app

COPY listenerApp /app

# COPY --from=builder /app/authApp /app 

CMD [ "/app/listenerApp" ]