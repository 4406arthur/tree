
# mlass-job-watcher

This repository implements a simple controller for watching Pods status own by specfic Job.


## Details

The controller uses [client-go library](https://github.com/kubernetes/client-go/tree/master/tools/cache) extensively.
The details of interaction points of the sample controller with various mechanisms from this library are
explained [here](docs/controller-client-go.md).

## Fetch dependencies

Like the rest of Kubernetes, ontroller has used
[godep](https://github.com/tools/godep) and `$GOPATH` for years and is
now adopting go 1.11 modules.  There are thus two alternative ways to
go about fetching this demo and its dependencies.


### When using go 1.11 modules

When using go 1.11 modules (`GO111MODULE=on`), issue the following
commands --- starting in whatever working directory you like.

```sh
git clone https://github.com/kubernetes/sample-controller.git
cd sample-controller
```

Note, however, that if you intend to
[generate code](#changes-to-the-types) then you will also need the
code-generator repo to exist in an old-style location.  One easy way
to do this is to use the command `go mod vendor` to create and
populate the `vendor` directory.

### A Note on kubernetes/kubernetes

If you are developing Kubernetes according to
https://github.com/kubernetes/community/blob/master/contributors/guide/github-workflow.md
then you already have a copy of this demo in
`kubernetes/staging/src/k8s.io/sample-controller` and its dependencies
--- including the code generator --- are in usable locations
(valid for all Go versions).
