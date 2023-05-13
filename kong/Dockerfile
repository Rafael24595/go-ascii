FROM kong:latest
USER root
# Example for Python
# PYTHONWARNINGS=ignore is needed to build gevent on Python 3.9
RUN apt-get update
RUN apt-get install -y python3
RUN apt-get install -y python3-pip
RUN apt-get install -y python3-dev
RUN apt-get install -y musl-dev
RUN apt-get install -y libffi-dev
RUN apt-get install -y gcc
RUN apt-get install -y g++
RUN apt-get install -y file
RUN apt-get install -y make
RUN pip3 install kong-pdk
WORKDIR /opt/kong-python-pdk
COPY ./kong-python-pdk .
# reset back the defaults
USER kong
ENTRYPOINT ["/docker-entrypoint.sh"]
EXPOSE 8000 8443 8001 8444
STOPSIGNAL SIGQUIT
HEALTHCHECK --interval=10s --timeout=10s --retries=10 CMD kong health
CMD ["kong", "docker-start"]