FROM ubuntu:18.04
LABEL author="Kyle McCann <ylmcc@redbrick.dcu.ie>"

WORKDIR /odin-testing
COPY . /odin-testing

RUN pwd & ls -l 

RUN  cp /odin-testing/odin-engine/config/odin-config.yml /root/
RUN ls /root 

# Get some prerequisites
RUN apt-get update
RUN DEBIAN_FRONTEND="noninteractive" apt-get install -y tzdata wget build-essential sudo apt-utils net-tools vi

# Install Go
RUN wget -c https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
ENV PATH=$PATH:/usr/local/go/bin
RUN go version


# Build Odin
RUN cd /odin-testing/odin-engine && make install




#Build Odin Cli
RUN cd /odin-testing/odin-cli && make install
RUN /bin/bash -c 'if grep -q "odin" /etc/group; then echo "odin group already exists!"; else groupadd odin && echo "odin group created!"; fi'


# Env variables
ENV ODIN_EXEC_ENV True
ENV ODIN_MONGODB "mongodb://0.0.0.0:27017/"



VOLUME ["/home/kyle/odin/"]
CMD ["/bin/odin-engine", "-id", "master-node", "raft0"]