FROM alpine
ADD usetemp /usetemp
ENTRYPOINT [ "/usetemp" ]
