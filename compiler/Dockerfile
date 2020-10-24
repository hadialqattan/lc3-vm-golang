FROM ubuntu:xenial

RUN apt-get update && apt-get install -y curl unzip flex wish make libncurses5-dev

RUN curl http://highered.mheducation.com/sites/dl/free/0072467509/104652/lc3tools_v12.zip > lc3tools_v12.zip

RUN unzip lc3tools_v12.zip

RUN cd lc3tools && ./configure --installdir /opt && make && make install

ENTRYPOINT ["/lc3tools/lc3as"] 
