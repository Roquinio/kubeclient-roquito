# Kubernetes client wrotes in Golang
## Usage 
<br>

### &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Help pannel
```
roquito -h 
```
***

### &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Overview of the cluster

```
roquito
```

```
---------------------------------------------------
  ____                            _   _
 |  _ \    ___     __ _   _   _  (_) | |_    ___
 | |_) |  / _ \   / _` | | | | | | | | __|  / _ \
 |  _ <  | (_) | | (_| | | |_| | | | | |_  | (_) |
 |_| \_\  \___/   \__, |  \__,_| |_|  \__|  \___/
                     |_|
---------------------------------------------------
Kubernetes client wrotes in Golang
---------------------------------------------------
fr1-srv-master-k8s-01
        RAM: Healty
        Disk : Healty
        CPU : Healty
        Node Status : Ready
fr1-srv-worker-k8s-02
        RAM: Healty
        Disk : Healty
        CPU : Healty
        Node Status : Ready
fr1-srv-worker-k8s-01
        RAM: Healty
        Disk : Healty
        CPU : Healty
        Node Status : Ready
fr1-srv-worker-k8s-03
        RAM: Healty
        Disk : Healty
        CPU : Healty
        Node Status : Ready
```

<br>

***

## How to install 

### Binaries already compiled


#### &nbsp; Requirement
- Curl 

#### &nbsp; Command

```
curl ftp.baptisteroques.fr/public/roquito/install-roquito.sh | bash
```

### Compiled by yourself 


#### &nbsp; Requirement 
- Git
- Golang 1.17.2 or latest version

#### &nbsp; Command

``` 
git clone https://github.com/Roquinio/kubeclient-roquito.git

cd kubeclient-roquito/

go build
./roquito
```
