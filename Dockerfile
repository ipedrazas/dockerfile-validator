FROM scratch
MAINTAINER Ivan Pedrazas <ipedrazas@gmail.com>

ADD docker-validator /

CMD ["/docker-validator"]
