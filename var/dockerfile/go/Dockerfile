FROM golang

WORKDIR /srv

ENV GOPATH /srv

# Install Docker
RUN apt-get update
RUN apt-get install -y curl
RUN curl -s https://get.docker.io/ubuntu/ | sh
# This allow our container to call host docker and starts other containers
RUN echo 'DOCKER_OPTS="-H :2375 -H unix:///var/run/docker.sock"' >> /etc/default/docker

# Install Docker Compose
RUN curl -L https://github.com/docker/compose/releases/download/1.2.0/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
RUN chmod +x /usr/local/bin/docker-compose

EXPOSE 2375

CMD /etc/init.d/docker start && sleep 1 && tail -F /var/log/upstart/docker.log
