FROM scratch
MAINTAINER Ivan Pedrazas <ipedrazas@gmail.com>

ADD dockerfile-validator /

CMD ["/dockerfile-validator"]
