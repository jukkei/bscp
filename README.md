# bscp - Backward scp

This is my first project using Go. The meaning of this project is to study Go!

## Problem description

* Make ssh-connection to a remote server
* Find an interesting file you want to copy to your local machine
* When at remote server, usually it's not possible to copy file to local machine (no ssh-server on local machine)
* You need to exit the ssh-session or open another terminal and use `scp` from local machine
* Problems:
  * You just have `ssh username@remoteIp` in your bash history
  * You don't remember in what folder the file was
  * You don't remember the `scp` syntax

Don't worry, we have the solution for this.

## Solution: bscp

`bscp` terminal application is a program that is supposed to be used at remote server (ssh). It simply creates a fully working `scp` command that you can paste to your local machine.

`bscp` gets the information from the following environment variables: `SSH_CONNECTION`, `USER` and `PWD`.

## Example

```bash
# on remote server
pi@raspberrypi:~ $ ls
Projects Downloads image.jpeg
pi@raspberrypi:~ $ bscp image.jpeg
scp -P 22 pi@192.168.1.21:/home/username/image.jpeg .
pi@raspberrypi:~ $
```

Copy-paste and

```bash
# at local machine
ubuntu@vm:~ $ scp -P 22 pi@192.168.1.21:/home/username/image.jpeg .
image.jpeg                           100% 2082   131.2KB/s   00:00
ubuntu@vm:~ $
```