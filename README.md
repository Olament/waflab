WAFLab 🐾
====

WAFLab is a web-based platform for testing WAFs.

# Live Demo

https://waflab.org/

# Architecture

WAFLab contains 2 parts:

Name | Description | Language | Source code
----|------|----|----
Frontend | Web frontend UI for WAFLab | Javascript + React + Ant Design | https://github.com/microsoft/waflab/tree/master/web
Backend | RESTful API backend for WAFLab | Golang + Beego + MySQL | https://github.com/microsoft/waflab

# Installation

## Prerequisites

- [Go](https://golang.org/)
- NPM (shipped with [Node.js](https://nodejs.org/))
- Docker (e.g., [Docker Desktop](https://docs.docker.com/docker-for-windows/install/) on Windows 10)

## Server-side

Get the source code from Github via Git

```bash
git clone https://github.com/microsoft/waflab.git
```

### Set up the database

WAFLab use database to store generated testcases and test results.

Prepare a [Xorm ORM](https://gitea.com/xorm/xorm) supported database (MySQL is recommended), replace `root:123@tcp(localhost:3306)/` in [conf/app.conf](https://github.com/microsoft/waflab/blob/master/conf/app.conf) with your own connection string. WAFLab will create a database named `waflab` and necessary tables in it if not exist. All Xorm supported databases are listed [here](https://gitea.com/xorm/xorm#user-content-drivers-support).

### Setup Go backend

Run Server-backend (at port 7070 by default):

```bash
cd waflab
go run main.go
```
 
Use your own WAFBench image (optional)

If you want to add some customized behavior to the package sending process. You can use your own WAFBench image. Once your made the change to the WAFBench codebase, you need to build a new Docker image and upload it to DockerHub.

```bash
cd WAFBench
docker build . -t org/wafbench
docker login
docker push org/wafbench
```

Then, change the ```docker.io/waflab/wafbench``` within ```docker/master.go``` to the image you just build. Ex. ```docker.io/org/wafbench```


### Setup frontend web UI

#### Install the frontend dependencies with NPM (or Yarn if you like):

```bash
cd waflab/web
npm install
```

#### Run frontend (at port 7000 by default):

```bash
npm start
```

WAFLab web UI is now avaliable at: http://localhost:7000

#### Build frontend into static files and it will be served by Go server at port 7070 together with backend API:

```bash
npm build
```

## License

This project is licensed under the [MIT license](LICENSE).

If you have any issues or feature requests, please contact us. PR is welcomed.
- https://github.com/microsoft/waflab/issues

## Contributing

This project welcomes contributions and suggestions.  Most contributions require you to agree to a
Contributor License Agreement (CLA) declaring that you have the right to, and actually do, grant us
the rights to use your contribution. For details, visit https://cla.opensource.microsoft.com.

When you submit a pull request, a CLA bot will automatically determine whether you need to provide
a CLA and decorate the PR appropriately (e.g., status check, comment). Simply follow the instructions
provided by the bot. You will only need to do this once across all repos using our CLA.

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or
contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.

## Trademarks

This project may contain trademarks or logos for projects, products, or services. Authorized use of Microsoft
trademarks or logos is subject to and must follow
[Microsoft's Trademark & Brand Guidelines](https://www.microsoft.com/en-us/legal/intellectualproperty/trademarks/usage/general).
Use of Microsoft trademarks or logos in modified versions of this project must not cause confusion or imply Microsoft sponsorship.
Any use of third-party trademarks or logos are subject to those third-party's policies.
