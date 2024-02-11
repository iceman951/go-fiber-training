FROM alpine:3.18
ADD ./bin/calcal /
ENV TZ=Asia/Bangkok
CMD ["/calcal"]
