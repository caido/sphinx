package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net"
	"net/http"
	"sort"
	"strings"

	"gopkg.in/yaml.v1"
)

// Request contains any info necessary to ratelimit a request
type Request map[string]interface{}

// InSlice tests whether or not a string exists in a slice of strings
func InSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// ReMarshal parses interface{} into concrete types
func ReMarshal(config interface{}, target interface{}) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, target)
}

// HTTPToSphinxRequest converts an http.Request to a Request
func HTTPToSphinxRequest(r *http.Request) Request {
	return map[string]interface{}{
		"path":       r.URL.Path,
		"headers":    r.Header,
		"remoteaddr": realipRemoteAddr(r),
		"method":     r.Method,
	}
}

// Hash hashes a string based on the given salt
func Hash(str, salt string) string {
	if salt == "" {
		return str
	}
	hash := hmac.New(sha256.New, []byte(salt))
	hash.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

// SortedKeys returns a sorted slice of map keys
func SortedKeys(obj map[string]interface{}) []string {
	// use make so as to prevent re-allocation
	keys := make([]string, len(obj))
	i := 0
	for k := range obj {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}

// realipRemoteAddr returns the real IP of the user
func realipRemoteAddr(r *http.Request) string {
	ips := r.Header.Get("X-Forwarded-For")
	splitIps := strings.Split(ips, ",")

	if len(splitIps) > 0 {
		netIP := net.ParseIP(splitIps[0])
		if netIP != nil {
			return netIP.String()
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}

	netIP := net.ParseIP(ip)
	if netIP != nil {
		ip := netIP.String()
		if ip == "::1" {
			return "127.0.0.1"
		}
		return ip
	}

	return ""
}
