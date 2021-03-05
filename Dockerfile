FROM ubuntu:18.04
LABEL author="Kyle McCann <ylmcc@redbrick.dcu.ie>"

WORKDIR /tmp/odin
COPY . /odin-testing

RUN  cp /odin-testing/start.sh /usr/start.sh 
RUN  cp /odin-testing/odin-engine/config/odin-config.yml /root/


# Get some prerequisites
RUN apt-get update
RUN DEBIAN_FRONTEND="noninteractive" apt-get install -y tzdata wget build-essential sudo apt-utils net-tools netcat systemd

# Install Go
RUN wget -c https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
ENV PATH=$PATH:/usr/local/go/bin
RUN go version

RUN cp  /odin-testing/cert/ca.crt /usr/local/share/ca-certificates/
RUN chmod 644 /usr/local/share/ca-certificates/ca.crt && update-ca-certificates

# Build Odin
RUN cd /odin-testing/odin-engine && make install




#Build Odin Cli
RUN cd /odin-testing/odin-cli && make install
RUN /bin/bash -c 'if grep -q "odin" /etc/group; then echo "odin group already exists!"; else groupadd odin && echo "odin group created!"; fi'


# Env variables
ENV ODIN_EXEC_ENV True
ENV ODIN_MONGODB "mongodb://CHANGE_IP:27017/"



VOLUME ["/home/kyle/odin/"]
CMD ["/bin/odin"]