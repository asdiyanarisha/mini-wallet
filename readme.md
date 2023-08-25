## Mini Wallet

A service mini wallet build it with golang. and store it in json file

### Requirement
- Golang 1.9
- Docker

### How To Run

to running this application, you must build a application image first with docker:
```shell
docker build -t mini-wallet:0.1.0 --rm .
```

after image already created. next step is running a docker container with command below:
```shell
docker run -p 8000:8080 --rm -d --name mini-wallet mini-wallet:0.1.0
```
if steps in above is success, this app can be served in a postman or browser when you go to ``http://localhost:8000``