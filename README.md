# experimental-kpackaging 

a small experiment to test test the use of Kpack-cli 




## Setup 

Create a new cluster
```
make cluster 
```

Configure the docker registry secrets using the steps described [here](https://github.com/buildpacks-community/kpack/blob/main/docs/tutorial.md) 

## Install kpack controller
```
make Install
```

### Create the required CRDs
```
make setup 
```

## Run the example 
```
go run main.go 
```

## check the build status 
if you have the [kpack-cli](https://github.com/vmware-tanzu/kpack-cli/releases) installed
```
kpack build logs kp-test -n default
```



