---
author: gripdev
category:
  - azure
  - azure-file-shares
  - azure-storage
  - symlinks
date: "2017-09-26T15:12:12+00:00"
guid: http://blog.gripdev.xyz/?p=1005
tag:
  - azure
  - azure-file-shares
  - azure-storage
  - symlinks
  - ubuntu
title: 'Quick How to: Mount Azure Files Shares with Symlinks support on Ubuntu'
url: /2017/09/26/quick-how-to-mount-azure-files-shares-with-symlinks-support-on-ubuntu/

---
Update Oct 2018: To see how to use this in Kuberentes [check out this blog post by Daniele Maggio](https://www.danielemaggio.eu/containers/azure-files-shares-with-symlinks-support-on-aks/)

By default mounting Azure File Shares on linux using CIFS doesn't enable support for symlinks. You'll see an error link this:

```
auser@acomputer:/media/shared$ ln -s linked -n t
ln: failed to create symbolic link 't': Operation not supported
```

So how do you fix this, simple? Simple add the following to the end of your CIFS mount command:

```
,mfsymlinks
```

So the command will look something like:

```
sudo mount -t cifs //<your-storage-acc-name>.file.core.windows.net/<storage-file-share-name> /media/shared -o vers=3.0,username=<storage-account-name>,password='<storage-account-key>',dir_mode=0777,file_mode=0777,mfsymlinks
```

So what does this do? Well you have to thank Steve French and Conrad Minshal. They defined a format for storing symlinks on SMB shares, an explanation [of the format can be found here](https://wiki.samba.org/index.php/UNIX_Extensions#Storing_symlinks_on_Windows_servers).

Thanks to [renash for her comment](https://docs.microsoft.com/en-gb/azure/storage/files/storage-how-to-use-files-linux#lf-content=177468006:734328535) (scroll to the bottom) which enabled me to find this, blog is to help others and give more details.
