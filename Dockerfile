FROM alpine:3.11

ENV GO111MODULE=on

RUN apk add --no-cache git make musl-dev openssh

# Install Go
RUN apk add --no-cache go

# Install Python and pip
RUN apk add --no-cache python3 py3-pip

# Install Pytest
RUN pip3 install -U pytest

# Install Ruby, RSpec and Rake
RUN apk add --no-cache ruby
RUN gem install rspec rake

# Install Terraform
RUN wget https://releases.hashicorp.com/terraform/0.12.21/terraform_0.12.21_linux_amd64.zip
RUN unzip terraform_0.12.21_linux_amd64.zip && rm terraform_0.12.21_linux_amd64.zip
RUN mv terraform /usr/bin/terraform

# Install Kubeval
RUN wget https://github.com/instrumenta/kubeval/releases/download/v0.16.1/kubeval-linux-amd64.tar.gz
RUN tar xf kubeval-linux-amd64.tar.gz && rm kubeval-linux-amd64.tar.gz
RUN cp kubeval /usr/local/bin

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

ADD src/main.go /go/src/

WORKDIR /go/src

RUN go mod init maceio
RUN go mod tidy
RUN go build -o maceio
RUN mv maceio ../bin/

CMD ["/go/bin/maceio"]
