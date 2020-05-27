# Installation

There are two steps to installing Odin:
1. [Odin Engine, MongoDB and Odin CLI](https://github.com/theycallmemac/odin/blob/master/INSTALL.md#1-odin-engine-mongodb-and-odin-cli)
2. [Odin Observability Dashboard](https://github.com/theycallmemac/odin/blob/master/INSTALL.md#2-odin-observability-dashboard)

After completing both steps you will be ready to consult [the documentation](https://github.com/theycallmemac/odin/blob/final-year-project/docs/documentation/Odin-User-Manual.pdf) to start using Odin!


## 1. Odin Engine, MongoDB and Odin CLI

First off, before building the project, we must first clone the repository via HTTPS with the following command:

```
git clone https://github.com/theycallmemac/odin.git 
```

To build the project, we can consult the Makefile in this directory. This file will automate the installation of the:
- Odin Engine
- Odin CLI
- MongoDB instance

Along with this, the Odin Engine will be run as a systemd service and the Odin CLI will be universally accessible from the `/bin` directory.

To utilise this automation we must run the makefile as the root user as so with the make command. This will:
- build the Odin Engine
- build the Odin CLI
- move the `odin-engine/config/odin-config.yml` file to the root user home directory
- move the generated CLI and Engine binary to the `/bin` directory
- move the `odin-engine/init/odin.service` file to `/lib/systemd/system` so it can be run as a systemd service
- install a locally accessible MongoDB
- creates the odin group, which users must be a member of to use the system.

We can verify all components were successfully install with the following series of commands:
```
systemctl status odin

systemctl status mongod

odin --help
```

It is advised the first two commands are run through root or with sudo. These commands will allow systemd to report the current status of the Odin Engine and the local MongoDB instance. The final command will verify you have a working install of the Odin CLI tool.

It’s also advisable to set the following two environment variables for the Odin Engine to use:

```
export ODIN_EXEC_ENV=True
export ODIN_MONGODB="mongodb://localhost:27017"
```

## 2. Odin Observability Dashboard

To set up the Odin Observability Dashboard, just run the following command in both the `odin-dashboard/client` and `odin-dashboard/server` directory to install dependencies:

```
npm install
```

In the `odin-dashboard/client` directory, you can run: 

`npm install -g @angular/cli@latest`

This will install the latest version of the Angular CLI tool. To make the user interface accessible at `http://localhost:4200`, just run:

```
ng serve
```

In the `odin-dashboard/server` directory just run:

```
npm start
```

This will start up the backend server for the dashabord. This server will start listing on port 3000 and will be accessible at http://localhost:3000. 

---
