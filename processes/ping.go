package main
//import "github.com/sparrc/go-ping"
pinger.SetPrivileged(true)

func main(){
go get github.com/sparrc/go-ping

}

/*

timeout := flag.Duration("t", time.Second*2, "")
ch := make(chan *ping.Statistics)
for i := 0; i < 3; i++ {
go myPing(XX, *timeout, ch)
}

for i := 0; i < 3; i++ {
stats := <- ch
fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n\n",
stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
}

func myPing(XXX, toValue time.Duration, ch chan *ping.Statistics) {
	pinger, _ := ping.NewPinger("www.google.com")
	pinger.Timeout = toValue
	pinger.Run()
	ch <- pinger.Statistics()
}

 */