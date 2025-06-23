
# gorecon - Go-Based Web Recon Scanner

gorecon is a lightweight, fast, and extensible web reconnaissance tool built in Go. It performs network and HTTP probing on target domains to assist with security testing, bug bounty hunting, and red team recon.

---

## 🚀 Features

- 🔎 Port probing on common web service ports (80, 443, 8080, etc.)
- 🌐 DNS CNAME lookup for subdomain analysis
- 🧠 Foundation for subdomain takeover detection (coming soon)
- ⚡ Highly concurrent and fast (future enhancement)
- 📄 JSON output (planned)
- 🔧 Easily extendable with additional scanning modules

---

## 🧰 Technologies Used

| Task             | Go Package                   |
|------------------|------------------------------|
| HTTP requests    | `net/http`                   |
| Port scanning    | `net.DialTimeout`            |
| DNS lookups      | `net.LookupCNAME`            |
| CLI flags        | `flag`                      |
| Concurrency      | Goroutines and Channels (planned) |

---

## 🏁 Getting Started

### Prerequisites

- Go 1.20+ installed

### Clone and Build

```bash
git clone https://github.com/iamritikbhardwaj/gorecon.git
cd gorecon
go build -o gorecon main.go
```

### Usage

Create a text file with target domains or IP addresses, one per line. For example:

```text
example.com
google.com
unclaimed.github.io
```

Run gorecon:

```bash
./gorecon -targets domains.txt
```

Sample output:

[DNS] example.com -> CNAME: example.com.
[PORT SCAN] Scanning example.com...
[OPEN] example.com:80
[OPEN] example.com:443

---

## ⚙️ CLI Flags

| Flag       | Description                            | Default    |
| ---------- | -------------------------------------- | ---------- |
| `-targets` | Path to file with target domains/IPs   | *Required* |
| `-timeout` | Timeout for network requests (seconds) | 5          |
| `-v`       | Enable verbose logging                 | false      |

---

## 📚 Future Roadmap

* ✅ DNS CNAME Lookup
* ✅ Port scanning on common web ports
* 🔲 Concurrency for faster scans
* 🔲 HTTP status and header enumeration
* 🔲 Subdomain takeover detection
* 🔲 Directory brute forcing with wordlists
* 🔲 JSON output and report generation
* 🔲 Webhooks and dashboard integration

---

## 📝 License

MIT License

---

## ✍️ Author

Ritik Singh – Go enthusiast and security engineer in training

---

## 📚 Related Resources

* [TryHackMe](https://tryhackme.com)
* [Hack The Box](https://hackthebox.eu)
* [OWASP Top 10](https://owasp.org/www-project-top-ten/)

---