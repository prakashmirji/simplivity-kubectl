# simplivity-kubectl

This repo contains the source for building and running sample kubectl plugin for hpe simplivity.


### Pre-reqs

The plugin uses simplivity-go sdk package from https://github.com/HewlettPackard/simplivity-go


### How to build
This plugin is basic initial version. It currently supports two arguments to get the vms lists and hosts lists managed by simplivity cluster.
Note: Code can be improved and modified to suite the actual needs. Currently this repo show cases the plugin functionality.

Step 1:
checkout the source repo
Edit the vms.go and hosts.go to pass simplivity cluster credentials
    username
    password
    ip/hostname

```cd simplivity-kubectl
    go build -o svt-vms ./vms.go
    go build -o svt-hosts ./hosts.go```

Step 2:
move the kubectl-svt to /usr/local/bin
```chmod +x kubectl-svt
    sudo mv kubectl-svt /usr/local/bin```

Then move the binaries you built in the preivous step to the /usr/local/bin or $PATH
```sudo mv svt-vms svt-hosts /usr/local/bin```

### How to test
Note: Assuming you are having access to kubernetes cluster and kubectl client before running below command

Run below kubectl command
To get list of VMs managed by simplivity cluster
```kubectl svt vms```

To get list of Hosts managed by simplivity cluster
```kubectl svt hosts```

### Contribuitions
Contributions and feedback is welcome
If you have feedback, please email me at prakash-r.mirji@hpe.com or create an issues in git repo



