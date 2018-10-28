FROM gitlab.xunlei.cn/docker/xluser-centos:v1.0.1

MAINTAINER  zhangkunming <zhangkunming@xunlei.com>

ADD ./bin /smart-edu-server/bin
ADD ./conf /smart-edu-server/conf

WORKDIR /smart-edu-server/bin

CMD ["./smart-edu-server"]
