FROM alpine:3.11

ENV GO111MODULE=on

RUN apk add --no-cache git make musl-dev openssh

# Install Go
RUN apk add --no-cache go

# Install Python and pip
RUN apk add --no-cache python3 py3-pip

# Install dotnet
RUN apk add bash icu-libs krb5-libs libgcc libintl libssl1.1 libstdc++ zlib
RUN apk add libgdiplus --repository https://dl-3.alpinelinux.org/alpine/edge/testing/

RUN wget https://download.visualstudio.microsoft.com/download/pr/08b6d245-401d-4b11-8e75-f2db47b5f166/5809c92b864453f3f666b8a9ce82f826/dotnet-sdk-5.0.400-linux-musl-x64.tar.gz -P /tmp/
RUN tar -zxvf /tmp/dotnet-sdk-5.0.400-linux-musl-x64.tar.gz -C /tmp/
RUN cp /tmp/dotnet /usr/local/bin/

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
ADD src/bitbucket/auth.go /go/src/bitbucket/
ADD src/bitbucket/events.go /go/src/bitbucket/

WORKDIR /go/src

RUN go mod init maceio
RUN go mod tidy
RUN go build -o maceio
RUN mv maceio ../bin/

CMD ["/go/bin/maceio"]
