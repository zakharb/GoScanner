
![logo](logo.png)

## GO Scanner  
### Simple GO Port Scanner  


![](https://img.shields.io/badge/version-1.0-red)
![](https://img.shields.io/badge/go-1.18-red)

## Content  
[Important info](#important_info)  
[Install](#install)  
[Run](#install)  
[Usage](#usage)  


<a name="important_info"/>

## Important info  
</a>  

> GO Scanner is utility to fast port scanning    
> Using goroutines for multiprocessing    

<a name="install"/>  

## Install  
</a>  

- Install GO and run  
```
go mod init goscanner
go mod tidy 
go build goscanner.go  
```

<a name="run"/>  

## Run  
</a>  

- Start with some help  
```
└─$ ./goscanner -h       
usage: goscanner [-h|--help] --ip "<value>" --port "<value>"

                 Start scanner

Arguments:

  -h  --help  Print help information
      --ip    Target IP address
      --port  Ports to scan, example: 21 / 80,443 / 1-1024
```

<a name="run"/>  

## Run  
</a>  

- Scan one port  
```
└─$ ./goscanner --ip 127.0.0.1 --port 80        
```

- Scan some ports  
```
└─$ ./goscanner --ip 127.0.0.1 --port 80,443       
```

- Scan port range  
```
└─$ ./goscanner --ip 127.0.0.1 --port 1-1024       
```