package naming

import (
	"strings"
	"sync"
)

var acronymsMapRWMutex sync.RWMutex
var acronymsMap = map[string]string{
	"tcp":   "TCP",
	"http":  "HTTP",
	"udp":   "UDP",
	"id":    "ID",
	"ssl":   "SSL",
	"tls":   "TLS",
	"cpu":   "CPU",
	"dob":   "DOB",
	"ttl":   "TTL",
	"sso":   "SSO",
	"https": "HTTPS",
	"ip":    "IP",
	"xss":   "XSS",
	"os":    "OS",
	"sip":   "SIP",
	"xml":   "XML",
	"json":  "JSON",
	"html":  "HTML",
	"xhtml": "XHTML",
	"xsl":   "XSL",
	"xslt":  "XSLT",
	"yaml":  "YAML",
	"toml":  "TOML",
	"wlan":  "WLAN",
	"wifi":  "WIFI",
	"vm":    "VM",
	"jvm":   "JVM",
	"ui":    "UI",
	"uri":   "URI",
	"url":   "URL",
	"sla":   "SLA",
	"scp":   "SCP",
	"smtp":  "SMTP",
	"soa":   "SOA",
	"oa":    "OA",
	"svg":   "SVG",
	"png":   "PNG",
	"jpg":   "JPG",
	"jpeg":  "JPEG",
	"pdf":   "PDF",
	"io":    "IO",
}

// AddAcronym adds acronyms for Pascal and Camel cases
func AddAcronym(acronyms ...string) {
	acronymsMapRWMutex.Lock()
	for _, acronym := range acronyms {
		acronymsMap[strings.ToLower(acronym)] = acronym
	}
	acronymsMapRWMutex.Unlock()
}

// RemoveAcronym removes acronyms for Pascal and Camel cases
func RemoveAcronym(acronyms ...string) {
	acronymsMapRWMutex.Lock()
	for _, acronym := range acronyms {
		delete(acronymsMap, strings.ToLower(acronym))
	}
	acronymsMapRWMutex.Unlock()
}
