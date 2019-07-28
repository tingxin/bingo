FROM alpine:3.5
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY ./bin/. /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/bingo"]
