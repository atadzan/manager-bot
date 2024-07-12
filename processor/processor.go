package processor

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/atadzan/bv-manager-bot/config"
)

const successResponseSample = "[info] Available formats for a9LDPn-MO4I:"

type Processor interface {
	ListProxies() (msg string)
	CheckProxies() (msg string)
}

type eventProcessor struct {
	proxies []config.Proxy
}

func New(proxies []config.Proxy) Processor {
	return &eventProcessor{proxies: proxies}
}

func (p *eventProcessor) ListProxies() (msg string) {
	if len(p.proxies) == 0 {
		msg = fmt.Sprintf("Proxy list is empty")
		return
	}
	msg = fmt.Sprintf("We have %d proxies.", len(p.proxies))
	for _, proxy := range p.proxies {
		msg = fmt.Sprintf("%s\nURL: %s. Region: %s", msg, proxy.URL, proxy.CountryCode)
	}
	return
}

func (p *eventProcessor) CheckProxies() (msg string) {
	var activeProxies []string
	if len(p.proxies) == 0 {
		msg = fmt.Sprintf("Proxy list is empty")
		return
	}
	for _, proxy := range p.proxies {
		cmd := exec.Command("yt-dlp", "-F", fmt.Sprint("https://www.youtube.com/watch?v=a9LDPn-MO4I"), "--proxy", proxy.URL)
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("Error while executign command: %s. Error: %v. Output: %s", cmd.String(), err, string(output))
			continue
		}
		if strings.Contains(string(output), successResponseSample) {
			activeProxies = append(activeProxies, fmt.Sprintf("%s. Region: %s", proxy.URL, proxy.CountryCode))
		}
	}
	if len(activeProxies) != 0 {
		msg = fmt.Sprintf("We have %d active proxies from %d.\n%s", len(activeProxies), len(p.proxies), strings.Join(activeProxies, "\n"))
	} else {
		msg = fmt.Sprintf("We don't have any active proxies 😢.")
	}
	return
}
