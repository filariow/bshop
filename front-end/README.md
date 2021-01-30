# front-end

## Project setup
```
yarn install
```

### Compiles and hot-reloads for development


```
# generate certificates
pushd certs
openssl req -nodes -new -x509 -keyout server.key -out server.cert
popd

# run the application
yarn run serve
```

### Compiles and minifies for production
```
yarn run build
```

### Run your tests
```
yarn run test
```

### Lints and fixes files
```
yarn run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).
