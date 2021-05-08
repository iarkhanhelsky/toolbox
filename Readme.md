# toolbox

Set of handy tools for occasional use. 

## lsprofiles

List essential information about your [provisioning profiles](https://developer.apple.com/documentation/appstoreconnectapi/profiles) 

```
$ lsprofiles --help
Usage of lsprofiles:
  -a string
    	Filter by Application ID
  -appid-filter string
    	Filter by Application ID
  -d	Print full information for each profile
  -p	Print provisioning profile plist
  -path string
    	Directory path or *.mobileprovision file (default "/Users/dm/Library/MobileDevice/Provisioning Profiles")
  -print-details
    	Print full information for each profile
  -print-plist
    	Print provisioning profile plist
  -u string
    	Filter by UUID
  -uuid-filter string
    	Filter by UUID
  -v	Show version and exit
```

### Example
```
$ lsprofiles
Created    Env Team ID    Name                                     UUID                                 File
2021-04-09     XX99Y9Z999 *                                        73f2206a-afe5-11eb-85d8-a348b1982457 73f2206a-afe5-11eb-85d8-a348b1982457.mobileprovision
2021-04-15   P XX99Y9Z999 com.org.App1                             79ea2738-afe5-11eb-ae90-53a47f3f01eb 79ea2738-afe5-11eb-ae90-53a47f3f01eb.mobileprovision
2021-04-15   D XX99Y9Z999 com.org.App1                             85d97cec-afe5-11eb-ba99-abef47a8fd4f 85d97cec-afe5-11eb-ba99-abef47a8fd4f.mobileprovision
2021-03-24   D XX99Y9Z999 com.org.App2                             a9aa5b8c-afe5-11eb-a8e6-53de6cdc54d5 a9aa5b8c-afe5-11eb-a8e6-53de6cdc54d5.mobileprovision
```



## lshosts

Print `~/.ssh/config` in compacted way for easier greping. 

It treats each `Host` entry in ssh-config as separate host. And prints essential information about it.  

### Example
```
$ lshosts
*
vm-super-1                     8.8.8.8
vm-box-1                       neo@8.8.8.1 -p 2221
vm-box-2                       trinity@8.8.8.1 -p 2222
vm-box-3                       morpheus@8.8.8.1 -p 2223
```
