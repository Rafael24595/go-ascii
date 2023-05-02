FROM ubuntu
RUN apt-get update
RUN apt-get install -y nano
RUN apt-get install -y lsof
RUN apt-get install -y curl
RUN apt-get install -y git
RUN apt-get install -y golang
RUN mkdir -p /go-ascii
RUN git clone https://github.com/Rafael24595/go-ascii.git ./go-ascii
WORKDIR go-ascii
RUN git checkout dev
ENTRYPOINT git pull && go run main.go