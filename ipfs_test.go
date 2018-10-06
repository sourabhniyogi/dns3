package dns3

import (
	"sync"
	"testing"
	"time"

	"github.com/miekg/dns"
)

func AnotherHelloServer(w ResponseWriter, req *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(req)

	m.Extra = make([]dns.RR, 1)
	m.Extra[0] = &dns.TXT{Hdr: RR_Header{Name: m.Question[0].Name, Rrtype: TypeTXT, Class: ClassINET, Ttl: 0}, Txt: []string{"Hello example"}}
	w.WriteMsg(m)
}

func TestServingListenAndServe(t *testing.T) {
	HandleFunc("example.com.", AnotherHelloServer)
	defer HandleRemove("example.com.")

	waitLock := sync.Mutex{}
	server := &dns.Server{Addr: ":0", Net: "udp", ReadTimeout: time.Hour, WriteTimeout: time.Hour, NotifyStartedFunc: waitLock.Unlock}
	waitLock.Lock()

	go func() {
		server.ListenAndServe()
	}()
	waitLock.Lock()

	c, m := new(dns.Client), new(dns.Msg)
	m.SetQuestion("example.com.", dns.TypeTXT)
	addr := server.PacketConn.LocalAddr().String() // Get address via the PacketConn that gets set.
	r, _, err := c.Exchange(m, addr)
	if err != nil {
		t.Fatal("failed to exchange example.com", err)
	}
	txt := r.Extra[0].(*dns.TXT).Txt[0]
	if txt != "Hello example" {
		t.Error("unexpected result for example.com", txt, "!= Hello example")
	}
	server.Shutdown()
}

/*
func TestRegisterDomain(t *testing.T) {
	domain := "eth.hacker"
	RegisterDomain(domain)
}

func TestSubmitZone(t *testing.T) {
	domain := "eth.hacker"
	zoneHash := "QmXkTBPtuJ1pTYRQ1U4AsSgAy1vE7r1EaMSAJ4pKMkZj89"
	SubmitZone(domain, zoneHash)
}
*/
