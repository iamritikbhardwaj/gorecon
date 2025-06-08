
# 🛠️ gorecon - Go-Based Web Recon Scanner

**gorecon** is a lightweight, fast, and extensible web reconnaissance tool built in Go. It performs network and HTTP probing on target domains to assist with security testing, bug bounty hunting, and red team recon.

---

## 🚀 Features

- 🔎 Port probing on common web service ports (`80`, `443`, `8080`, etc.)
- 🌐 HTTP(S) status & header enumeration
- 🧠 Subdomain takeover detection (CNAME analysis)
- 📂 Directory fuzzing using wordlists
- ⚡ Highly concurrent and fast
- 📄 JSON output (optional)
- 🔧 Easily extendable (webhooks, dashboards, etc.)

---

## 🧪 Example Usage

```bash
./gorecon -targets domains.txt
````

Provide a list of domains or IPs to scan for:


example.com
testsite.io
192.168.1.1

---

## 📦 Project Structure

gorecon/
├── main.go
├── scan/
│   ├── http.go       # HTTP request and analysis logic
│   ├── ports.go      # TCP port scanner
│   └── dns.go        # Subdomain/CNAME resolver
├── utils/
│   └── logger.go     # Logging utility
├── wordlists/
│   └── common.txt    # Directory brute-force list
└── README.md

---

## 🔧 Technologies Used

| Task          | Go Package                    |
| ------------- | ----------------------------- |
| HTTP requests | `net/http`, `httptrace`       |
| Port scanning | `net.Dial`, `net.DialTimeout` |
| Concurrency   | Goroutines, Channels          |
| CLI flags     | `flag` or `cobra`             |
| DNS lookups   | `net.LookupCNAME`             |

---

## 🧠 What You’ll Learn

* Go programming with real-world cybersecurity use cases
* Writing concurrent, fast, and maintainable code
* Web technologies: HTTP headers, status codes, DNS
* Recon techniques used in bug bounty and penetration testing

---

## 🌱 Future Extensions

* ✅ JSON output for scripting and integration
* ✅ Slack/Discord webhook support
* ✅ Web dashboard using Go templates
* ✅ Multi-threaded throttling and rate limiting
* ✅ Plugin system or API mode

---

## 🏁 Getting Started

```bash
go build -o gorecon main.go
./gorecon -targets domains.txt
```

Add your own wordlists in `wordlists/`, or use popular ones from [SecLists](https://github.com/danielmiessler/SecLists).

---

## 📸 Screenshots (Coming Soon)

---

## 📝 License

MIT License

---

## ✍️ Author

Ritik Singh – [@iamritikbhardwaj](https://github.com/iamritikbhardwaj)
Security engineer in training | Go enthusiast | Future OSCP 🛡️

---

## 📚 Related Resources

* [TryHackMe](https://tryhackme.com)
* [Hack The Box](https://hackthebox.com)
* [OWASP Top 10](https://owasp.org)
* [Go net/http package](https://pkg.go.dev/net/http)

