FROM gitlab.xunlei.cn/docker/xluser-centos:v1.0.1

MAINTAINER  zhangkunming <zhangkunming@xunlei.com>

ADD ./bin /github.com/KenmyZhang/smart-edu-server/bin
ADD ./conf /github.com/KenmyZhang/smart-edu-server/conf

WORKDIR /github.com/KenmyZhang/smart-edu-server/bin

CMD ["./smart-edu-server"]
