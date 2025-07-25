# XCalate
A privilege escalation tools based on TryHackMe.

Run simple check (not thorough) for possible vector of privilege escalation.

In the meantime it checks for 
- Linux
    - Possible Service Exploits
    - Weak File Permissions
        - /etc/shadow - Readable & Writable
        - /etc/passwd - Writable
    - Sudo - Shell Escape Sequences
    - SUID Executable
    - Cron Jobs
        - Overwrite custom cronjob
        - Check user's path inside the file
        - Usage of wildcards in cronjob

- Window
    - To be continued...

- ISSUE
    - cron only checks if custom script start with "/" THEN wildcard process, need to remove the condition so that every script will go through the process.
    - cron somehow unable to locate overwrite.sh