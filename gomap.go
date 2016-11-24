package main

import (
	"fmt"
  "os"
  "time"
  "net"
  "github.com/tatsushid/go-fastping"
)

func main(){
  icmp_test()
}


func icmp_test(){
  
  p := fastping.NewPinger()
  ra, err := net.ResolveIPAddr("ip4:icmp", "8.8.8.8")
  if err != nil {
      fmt.Println(err)
      os.Exit(1)
  }
  
	for i := 0; i < 100; i++{
		p.AddIPAddr(ra)
		p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
				fmt.Println(rtt)
				//fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
		}
		
		p.OnIdle = func() {
				//fmt.Println("finish")
		}
		err = p.Run()
		if err != nil {
				fmt.Println(err)
		}
	}

}