FROM alpine
ADD userApi-api /userApi-api
ENTRYPOINT [ "/userApi-api" ]
