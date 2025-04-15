# Hello go

Example app for OpenShift.

## Import ubi9 go 1.23

```
oc -n openshift tag registry.redhat.io/ubi9/go-toolset:1.23 golang:1.23-ubi9
oc -n openshift import-image golang:1.23-ubi9
```
