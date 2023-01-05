FROM golang:1.18-buster

WORKDIR /

COPY ./logistics-backend /logistics-backend
COPY ./resources /resources

EXPOSE 80

# USER nonroot:nonroot

ENTRYPOINT ["/oms-backend"]
