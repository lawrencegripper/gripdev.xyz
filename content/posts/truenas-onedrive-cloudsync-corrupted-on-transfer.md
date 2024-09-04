---
author: gripdev
category:
  - uncategorized
date: "2021-06-07T08:12:33+00:00"
guid: https://blog.gripdev.xyz/?p=1491
title: TrueNas OneDrive Cloudsync corrupted on transfer
url: /2021/06/07/truenas-onedrive-cloudsync-corrupted-on-transfer/

---
This is a quick one, if you get the following error or similar:

```
2021/06/07 00:15:35 ERROR : Attempt 3/3 failed with 1 errors and: corrupted on transfer: sizes differ 189118 vs 130560
2021/06/07 00:15:35 Failed to copy: corrupted on transfer: sizes differ 189118 vs 130560
```

These track back to an issue with OneDrives metadata generation altering the size of the file. You can see details on this issue here: https://github.com/rclone/rclone/issues/399

To resolve this you need to add `--ignore-size` to the `rclone` config that TrueNas creates.

While the UI doesn't expose a `extra-args` field, it is present in the underlying database. This post guides you through how to add additional args: https://www.truenas.com/community/threads/cloud-sync-task-add-extra-rclone-args-to-specify-azure-archive-access-tier.85526/

For this OneDrive error the following works: (Assuming you only have 1 cloudsync task ID == 1)

```
$ sqlite3 /data/freenas-v1.db
```

```
update tasks_cloudsync set args = "--ignore-size" where id = 1;
```

You can double check the change with the following

```
sqlite> .headers on
sqlite> select * from tasks_cloudsync;

```

Then it's just a case of re-running the task in the UI and 🎉

[![](/wp-content/uploads/2021/06/image.png)](/wp-content/uploads/2021/06/image.png)
