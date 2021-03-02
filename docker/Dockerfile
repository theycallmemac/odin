FROM ubuntu:18.04
LABEL author="Kyle McCann <ylmcc@redbrick.dcu.ie>"

WORKDIR /home/kyle/testing
COPY . /odin-testing



# Get some prerequisites
RUN apt-get update
RUN DEBIAN_FRONTEND="noninteractive" apt-get install -y tzdata wget build-essential sudo apt-utils systemd

# Install Go
RUN wget -c https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
ENV PATH=$PATH:/usr/local/go/bin
RUN go version

# Build Odin
ENV ODIN_EXEC_ENV True
ENV ODIN_MONGODB "mongodb://localhost:27017"
RUN ls /odin-testing
RUN sudo make /odin-testing

VOLUME ["/home/kyle/testing/"]
CMD ["/etc/init.d/odin", "start"]
