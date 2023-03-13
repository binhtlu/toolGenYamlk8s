##### Use binary **go-gen** in project folder to generate template file

### Files generate includes

* `init_config.sh`<br>
* `deploy-image.sh`<br>
* `Dockerfile`<br>
* `Dockerfile-<date_time>.bak`<br>
* `.dockerignore`<br>
* `deployment_file.yaml`<br>
* `deployment_file.yaml-<date-time>.bak`<br>
* `service_file.yaml`<br>
* `path-to-config/app.conf.example`<br>

### Usage of go-gen
**params**
```
-filepath string
    file absolute path config
-image string
    Set name of image to deploy on k8s (default "nginx")
-name string
    Set name of deployment (default "test-deployment")
-ns string
    Set name of namespace (default "default")
-prefix string
    add prefix to template configs
-replica string
    Set number of replica (default "1")
-secret string
    secret for pull image (default "regcred")
-servicename string
    Set name of Service (default "test-service")
```

#### Example: Full options
```
go-gen -filepath=<abs-path-to-confix> -prefix=<prefix> -image=<image-name> -name=<deployment-name> -ns=<name-namespace> -replica=<number-of-replica> -servicename=<service-name>
```
