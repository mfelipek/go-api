FROM golang:1.9 as builder
LABEL go-api-builder=true

# replace shell with bash so we can source files
RUN rm /bin/sh && ln -s /bin/bash /bin/sh

RUN apt-get update \
    && apt-get -y install \
    libssl-dev \
    build-essential \
    curl

ENV NVM_DIR /usr/local/nvm
ENV NODE_VERSION 6.9.2
    
# install nvm
# https://github.com/creationix/nvm#install-script
RUN curl --silent -o- https://raw.githubusercontent.com/creationix/nvm/v0.32.0/install.sh | bash

# install node and npm
RUN source $NVM_DIR/nvm.sh \
    && nvm install $NODE_VERSION \
    && nvm alias default $NODE_VERSION \
    && nvm use default

# add node and npm to path so the commands are available
ENV NODE_PATH $NVM_DIR/v$NODE_VERSION/lib/node_modules
ENV PATH $NVM_DIR/versions/node/v$NODE_VERSION/bin:$PATH

# confirm installation
RUN node -v
RUN npm -v

# Copy package.json
ADD ./package.json /go/src/go-api/

# install npm libraries
WORKDIR "/go/src/go-api"

# latest version of node
RUN npm install -g n
RUN n stable

# application packages
RUN npm install --save react react-dom
RUN npm install --save react-router
RUN npm install --save react-router-dom
RUN npm install --save redux react-redux redux-thunk redux-logger
RUN npm install --save cross-fetch babel-polyfill
RUN npm install --save-dev webpack -g
RUN npm install --save-dev babel-loader@8.0.0-beta.0 @babel/core @babel/preset-env @babel/preset-react
RUN npm install --save-dev css-loader@0.23.1 postcss-loader@0.9.1 react-hot-loader@3.0.0-beta.6 style-loader@0.13.1

# list package.json content
RUN cat ./package.json

# Copy whole app
ADD ./ /go/src/go-api

RUN webpack --config webpack.config.js

WORKDIR "/go/src"

RUN go version

# Use "go get" to fetch and build the project, does not install (-d flag)
RUN go get -d go-api

# make the executable
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app-go-api ./go-api

#####################################################

FROM alpine:latest
LABEL go-api-builder=false

RUN apk update
RUN apk upgrade
RUN apk --no-cache add ca-certificates
# Change TimeZone
RUN apk add --update tzdata
ENV TZ=America/Sao_Paulo
# Clean APK cache
RUN rm -rf /var/cache/apk/*

WORKDIR /root/

# copy app executable froum builder stage
COPY --from=builder /go/src/app-go-api .

# copy static content
COPY --from=builder /go/src/go-api/templates ./go-api/templates
COPY --from=builder /go/src/go-api/static ./go-api/static
COPY --from=builder /go/src/go-api/keys ./go-api/keys

# Run the app command by default when the container starts.
ENTRYPOINT /root/app-go-api

# Document that the service listens on port 8080.
EXPOSE 8080