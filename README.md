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
Add the *maceio.yaml* file to the application repository. The directory hierarchy should look like this:
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

# Configuration file
Some examples of the *maceio.yaml* file can be found below.

## Terraform
```
---
tests:
  - name: tf_validate
    cmd: terraform validate
  - name: tf_plan
    cmd: terraform plan
```

We do not think you should add the apply parameter from terraform to the tests, but remember to add the flag **-auto-approve** if you decide to do it

## Rspec
```
---
tests:
  - name: "Build test"
    cmd: "rspec spec hello_world.rb"
```

## Pytest
```
---
tests:
  - name: "Pytest local"
    cmd: pytest
```

## Custom

```
---
tests:
  - name: "Install mathutils"
    cmd: pip3 install mathutils
  - name: "Pytest local"
    cmd: pytest
```

The *cmd* specified in the configuration file will be executed in the root directory of the application repository.
