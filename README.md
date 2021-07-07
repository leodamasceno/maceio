# Maceio
Golang application that listens for webhook events coming from Github, runs tests previously defined in a configuration file and returns the output to a pull request.


# What does it do?
It runs a list of tests defined in a YAML file, located in the application's repository. Maceio currently supports the following test types:
- **Custom**
- **Pytest**
- **RSpec**
- **Rake**
- **Terraform**

The custom type can run any command present on Alpine. For instance, this would be the right check if you need to run **gem** or **pip3**. For more information, check the *Dockerfile* in this repository.

# Configuring your repository
### Add the configuration file
Add the *maceio.yaml* file to your repository, the structure should look like the one below:
```
├── app
│   ├── css
│   │   ├── **/*.css
│   ├── favicon.ico
│   ├── images
│   ├── js
│   │   ├── **/*.js
│   └── partials/template
├── maceio.yaml
├── README.md
├── package.json
└── .gitignore
```

### Add the webhook
- **Payload URL:** https://*YOUR_URL*/webhook
- **Content type:** application/json
- **Secret:** Secret created with the command ```ruby -rsecurerandom -e 'puts SecureRandom.hex(20)'``` 
- **Which events would you like to trigger this webhook:** Select *Pull requests* and *Pushes*

### Personal access tokens
Follow [this](https://docs.github.com/en/github/authenticating-to-github/keeping-your-account-and-data-secure/creating-a-personal-access-token) guide and select the scopes *repo* and *admin:repo_hook*.


# Configuration file
Some examples of the *maceio.yaml* file can be found below.

### Terraform
```
---
tests:
  - name: tf_validate
    cmd: terraform validate
  - name: tf_plan
    cmd: terraform plan
```

We do not think you should add the apply parameter from terraform to the tests, but remember to add the flag **-auto-approve** if you decide to do it

### Rspec
```
---
tests:
  - name: "Build test"
    cmd: "rspec spec hello_world.rb"
```

### Pytest
```
---
tests:
  - name: "Pytest local"
    cmd: pytest
```

### Custom

```
---
tests:
  - name: "Install mathutils"
    cmd: pip3 install mathutils
  - name: "Pytest local"
    cmd: pytest
```

The *cmd* specified in the configuration file will be executed in the root directory of the application repository.

# Running it
You can choose one of the two options below to run the application.

### From Github
The version you will find in Github may be unstable. It is not recommended for production.

Clone the repository:
```
git clone https://github.com/leodamasceno/maceio.git
cd maceio
```

Build the image locally:
```
docker build -t "bazer:0.1" .
```

Run it:
```
docker run --name bazer -p 8080:8080 \
-e GIT_TOKEN="*YOUR_GIT_TOKEN*" \
-e GIT_SECRET="*YOUR_GIT_SECRET*" bazer:0.1
```
### From DockerHub
Run it:
```
docker run --name bazer -p 8080:8080 \
-e GIT_TOKEN="*YOUR_GIT_TOKEN*" \
-e GIT_SECRET="*YOUR_GIT_SECRET*" leodamasceno/maceio:latest
```
