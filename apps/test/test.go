package main

import (
	"fmt"

	"log"

	"github.com/dustinevan/go-smaz"
)

func main() {
	s := "beep beeep beeeeep. This is a test of the smaz system, this is only a test"
	fmt.Println(s, "is", len([]byte(s)), "bytes")
	cs := smaz.Compress([]byte(s))
	fmt.Println(len([]byte(cs)), "bytes compressed")

	url := "http://tags.rd.linksynergy.com/imp?eID=1&type=1&nID=80&duration=50&mID=6178&v=2.26.8&aID=3821785&cts=08%2F11%2F2017+06%3A27%3A24&sID=090f6560-7e88-11e7-9f80-d19a58d5bcae&url=http%3A%2F%2Famp.rd.linksynergy.com%2F%3Fmerchant%3Djuice6178%26nID%3D80%26width%3D300%26height%3D250%26strategy%3Dretargeting%26cb%3D599168%2520%26redirecturl%3Dhttp%253A%252F%252Finsight.adsrvr.org%252Ftrack%252Fclk%253Fimp%253Dcc0ab5ae-e9b6-4082-b72e-33528ea2b0af%2526ag%253D4kzliob%2526sfe%253Dbc513b4%2526sig%253DRTzDKCLRYrQnI5y6hqDgI51wlPyg53paYhJemK5ba-w.%2526crid%253Dy74eegic%2526cf%253D94364%2526fq%253D14%2526td_s%253Dwww.minq.com%2526rcats%253Dusw%2526mcat%253D%2526mste%253D%2526mfld%253D2%2526mssi%253D%2526mfsi%253D86eq54bpwq%2526sv%253Dcasale%2526uhow%253D126%2526agsa%253D%2526rgco%253DUnited%252520States%2526rgre%253DTennessee%2526rgme%253D557%2526rgci%253DCrossville%2526rgz%253D38571%2526dt%253DMobile%2526osf%253DAndroid%2526os%253DAndroid70%2526br%253DChrome%2526svpid%253D183455%2526rlangs%253D01%2526mlang%253D%2526did%253D%2526rcxt%253DOther%2526tmpc%253D%2526vrtd%253D%2526osi%253D%2526osv%253D%2526daid%253D%2526dnr%253D0%2526vpb%253D%2526svsc%253D%2526dur%253DCigKDWNoYXJnZS1hbGwtMTkiFwjt__________8BEgpkcmF3YnJpZGdlChoKB251czRpNW4QpJEOIgsItdXZTBIEbm9uZRCkkQ4.%2526crrelr%253D%2526npt%253D%2526svscid%253D192296%2526mk%253DMotorola%2526mdl%253DXT1585%2526testid%253Dwith-iplookup%2526r%253D"
	fmt.Println(len([]byte(url)), "bytes")
	curl := smaz.Compress([]byte(url))
	fmt.Println(len([]byte(curl)), len(string(curl)), "bytes compressed")

	match, err := smaz.Decompress(curl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(match))
	fmt.Println( string(match) == url)
}
