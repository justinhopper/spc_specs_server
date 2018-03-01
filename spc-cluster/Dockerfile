FROM golang
ADD . /go/src/spc-app
RUN go install spc-app
ENTRYPOINT /go/bin/main
CMD ["./spc-app"]
