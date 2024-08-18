package processor

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"sync"

	"github.com/atadzan/bv-manager-bot/messages"
)

const (
	successResponseSample = "[info] Available formats for a9LDPn-MO4I:"
	storageFilePath       = "proxies.csv"
)

type Processor interface {
	ListProxies() (msg string)
	CheckProxies() (msg string)
	UpdateProxies(inputText string) (msg string)
	ClearProxyList() (msg string)
}

type eventProcessor struct {
	proxies []messages.Proxy
}

func New() Processor {
	return &eventProcessor{
		proxies: readFromFile(),
	}
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
	var wg = &sync.WaitGroup{}
	var mx = &sync.Mutex{}
	for _, proxy := range p.proxies {
		wg.Add(1)
		go func() {
			cmd := exec.Command("yt-dlp", "-F", fmt.Sprint("https://www.youtube.com/watch?v=a9LDPn-MO4I"), "--proxy", proxy.URL, "--socket-timeout", "1")
			output, err := cmd.CombinedOutput()
			if err != nil {
				log.Printf("Error while executign command: %s. Error: %v. Output: %s", cmd.String(), err, string(output))
			}
			if strings.Contains(string(output), successResponseSample) {
				mx.Lock()
				activeProxies = append(activeProxies, fmt.Sprintf("%s. Region: %s", proxy.URL, proxy.CountryCode))
				mx.Unlock()
			}
			defer func() {
				wg.Done()
			}()
		}()
	}

	wg.Wait()
	if len(activeProxies) != 0 {
		msg = fmt.Sprintf("We have %d active proxies from %d.\n%s", len(activeProxies), len(p.proxies), strings.Join(activeProxies, "\n"))
	} else {
		msg = fmt.Sprintf("We don't have any active proxies 😢.")
	}
	return
}

func (p *eventProcessor) UpdateProxies(inputText string) (msg string) {
	var proxies []messages.Proxy
	if err := json.Unmarshal([]byte(inputText), &proxies); err != nil {
		msg = fmt.Sprintf("Can't unmarshal input message. Input message: %s. Error: %v", inputText, err)
		return
	}
	p.proxies = append(p.proxies, proxies...)

	if err := saveToFile(p.proxies); err != nil {
		log.Printf("Can't update file containing proxy list. Error: %v\n", err)
	}
	return messages.ProxiesSuccessfullyUpdated
}

func (p *eventProcessor) ClearProxyList() (msg string) {
	p.proxies = []messages.Proxy{}
	removeStorageFile()
	return messages.ProxyListSuccessfullyCleared
}
