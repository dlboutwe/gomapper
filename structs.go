package main

import (
	"fmt"
  "os"
  "time"
  "net"
  "strings"
  "github.com/tatsushid/go-fastping"
)

type Probe struct{
  id int
  name, inputIP string
  status bool
}

// Converts string to IPAddr for later code useage
func (p *Probe) IPAddressConstructor() (*net.IPAddr, error) {
  r, err := net.ResolveIPAddr("ip4:icmp", p.inputIP)
  if err != nil {
    //Need to add error handling here.
      fmt.Println(err)
      os.Exit(1)
  }
  return r, err
}

func (p *Probe) Mongo_I(r time.Duration) bool{
  //convert time.Duration to string
  d := r.String()
  //truncate last 2 characters from rtt raw data
  d = strings.TrimSuffix(d, "ms")
  moaddr := MOAddr{"vps.rebirtharmitage.com:21701", "gomap", "test"}
  movalue := MOValue{"key", "value", "TEST004", d}
  a := mongo_insert(moaddr, movalue)
  return a
}

func (p *Probe) Mongo_O() []MOGValue {
	moaddr := MOAddr{"vps.rebirtharmitage.com:21701", "gomap", "test"}
  movalue := MOValue{"key", "value", "TEST004", ""}
	a := mongo_export(moaddr, movalue)
  return a
}

// Acutal ping to destination ip
func (p *Probe) Pinger() {
  dest, err := p.IPAddressConstructor()
  pi := fastping.NewPinger()
	for i := 0; i < 10; i++{
		pi.AddIPAddr(dest)
		pi.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
      p.Mongo_I(rtt)
		}    
		err = pi.Run()
		if err != nil {
				fmt.Println(err)
		}
	} 
	a := p.Mongo_O()
	fmt.Println(a)
}



