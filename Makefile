.ONESHELL:
all: mongodb engine cli group

engine:
	cd odin-engine
	make install
	cp config/odin-config.yml $${HOME}
	cp init/odin.service /lib/systemd/system
	systemctl daemon-reload
	systemctl start odin
	systemctl status odin

mongodb:
	apt update
	apt install gnupg
	wget -qO - https://www.mongodb.org/static/pgp/server-4.2.asc | sudo apt-key add -
	echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu bionic/mongodb-org/4.2 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-4.2.list
	apt update
	apt install -y mongodb-org
	systemctl daemon-reload
	systemctl start mongod

cli:
	cd odin-cli
	make install

group:
	/bin/bash -c 'if grep -q "odin" /etc/group; then echo "odin group already exists!"; else groupadd odin && echo "odin group created!"; fi'
