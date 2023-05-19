FROM golang:1.19.1-alpine
RUN mkdir test
WORKDIR /test
RUN echo "hello wrold" > sample
COPY ./ /test
CMD ["ls"]
CMD ["go", "test", "-v"]

    