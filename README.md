# Drone Plugin FaasCLI

The FaasCLI plugin can be used to build and publish functions to the container registry. The following pipeline configuration uses the FaasCLI plugin to build and publish.

```yaml
- name: build_deploy
  image: juli3nk/drone-faascli
  settings:
    yaml_file: stack.yml
    templates:
    - go
    - golang-middleware
    build: true
    push: true
    registry: registry.example.com
    registry_username: user
    registry_password: pass
```

## Parameter Reference

`yaml_file`
    yaml file describing function(s)

`templates`
    array of OpenFaaS templates to pull

`build`
    enable to build OpenFaaS function containers

`build_args`
    pass custom arguments to docker build

`build_args_from_env`
    pass the envvars as custom arguments to docker build

`push`
    enable to push OpenFaaS functions to remote registry

`docker_dns`
    set dns servers for the container

`registry`
    authenticates to registry

`registry_username`
    authenticates with username

`registry_password`
    authenticates with password

`registry_auth`
    auth token for the registry

`registry_insecure=false`
    enable insecure communication to registry

`debug`
    launch the docker daemon in verbose debug mode
