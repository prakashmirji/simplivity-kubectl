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
    go build -o svt-hosts ./hosts.go
```

Step 2:
move the kubectl-svt to /usr/local/bin

```
    chmod +x kubectl-svt
    sudo mv kubectl-svt /usr/local/bin
```

Then move the binaries you built in the preivous step to the /usr/local/bin or $PATH

```
    sudo mv svt-vms svt-hosts /usr/local/bin
```

### How to test
Note: Assuming you are having access to kubernetes cluster and kubectl client before running below command

Run below kubectl command

To get list of VMs managed by simplivity cluster

```
    kubectl svt vms
```

You should output like below ( sample )
```
List of VMs
----------------------------------------
Get list of all VMs
----------------------------------------
initial-user-cluster-6c76d6db95-l4mbk

initial-user-cluster-6c76d6db95-hp9m5

initial-user-cluster-6c76d6db95-hh5mt

initial-user-cluster-0-4qt9k-7957447795-qh29v

gke-on-prem-osimage-1.14.7-gke.40-20200205-54cf7b4525

gke-on-prem-admin-workstation-2

gke-on-prem-admin-appliance-vsphere-1.2.2-gke.2

gke-f5-bigip-15.1

gke-admin-node-6d9f987c5b-fdt8w

gke-admin-node-6d9f987c5b-2ggrd

gke-admin-master-2d6w8

UbuntuDebug
```


To get list of Hosts managed by simplivity cluster

```
    kubectl svt hosts
```

You should output like below ( sample)
```
List of Hosts
----------------------------------------
Get all Hosts without params 
----------------------------------------
10.3.4.62

10.3.4.61
```

### Contribuitions
Contributions and feedback is welcome

If you have feedback, please email me at prakash-r.mirji@hpe.com or create an issues in git repo



