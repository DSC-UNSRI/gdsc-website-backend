FROM scratch
WORKDIR /app
COPY ./app .
COPY .env .
CMD [ "./app" ]
EXPOSE 8000