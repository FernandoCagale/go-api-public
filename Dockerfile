FROM scratch

ADD bin/main /main

ENTRYPOINT ["./main"]