package report

import (
	"fmt"
	"strings"

	"github.com/projectdiscovery/gologger"
)

// taskOrder is optional — it defines print order
var taskOrder = []string{
	"Task 1 - Linux Checks",
	"Task 2 - Service Exploits",
	"Task 3 - Weak File Permissions - Readable /etc/shadow",
	"Task 4 - Weak File Permissions - Writable /etc/shadow",
	"Task 5 - Weak File Permissions - Writable /etc/passwd",
	"Task 6 - Sudo - Shell Escape Sequences",
	"Task 7 - Sudo - Environment Variables",
	"Task 8 - Cron Jobs - File Permissions",
	"Task 9 - Cron Jobs - PATH Environment Variable",
	"Task 10 - Cron Jobs - Wildcards",
	"Task 11 - SUID / SGID Executables - Known Exploits",
	"Task 12 - SUID / SGID Executables - Shared Object Injection",
	"Task 13 - SUID / SGID Executables - Environment Variables",
	"Task 14 - SUID / SGID Executables - Abusing Shell Features (#1)",
	"Task 15 - SUID / SGID Executables - Abusing Shell Features (#2)",
	"Task 16 - Passwords & Keys - History Files",
	"Task 17 - Passwords & Keys - Config Files",
	"Task 18 - Passwords & Keys - SSH Keys",
	"Task 19 - NFS",
	"Task 20 - Kernel Exploits",
}

// print summary of all tasks
func PrintSummary() {
	fmt.Println()
	fmt.Println(" [ 📋 ] Linux Privilege Escalation Checklist")

	for _, full := range taskOrder {
		left := full
		right := ""
		if idx := strings.Index(full, " - "); idx != -1 {
			left = full[:idx]
			right = full[idx+3:]
		}

		status, exists := taskResults[full]
		var mark string
		switch {
		case exists && status == Pass:
			mark = " [ ✅ ] "
		case exists && status == Manual:
			mark = " [ ✏️  ] "
		default:
			mark = " [ ❌ ] "
		}

		if right != "" {
			fmt.Printf("%s  %-10s - %s\n", mark, left, right)
		} else {
			fmt.Printf("%s  %-10s - %s\n", mark, left, full)
		}
	}
	fmt.Println()
}

// generic labeled output helper
func PrintField(label, name string, value interface{}) {
	gologger.Print().Label(label).Msgf("%-20s: %v", name, value)
}
