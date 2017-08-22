package main

import (
	"fmt"

	"log"

	"github.com/dustinevan/go-smaz"
)

func main() {
	url := "http://tags.rd.linksynergy.com/imp?eID=1&type=1&nID=20&duration=76&mID=711&v=2.26.8&aID=3923512&cts=08%2F11%2F2017+04%3A27%3A20&sID=09e5f54e-7e88-11e7-969c-d5952bae77f6&url=http%3A%2F%2Famp.rd.linksynergy.com%2F%3Fmerchant%3Dspartan_race%26nID%3D20%26width%3D300%26height%3D600%26strategy%3Dretargeting%26cb%3D5164167568320125335%26redirecturl%3Dhttp%3A%2F%2Fpixel.mathtag.com%2Fclick%2Fimg%253Fmt_aid%253D5164167568320125335%2526mt_id%253D2715666%2526mt_adid%253D103779%2526mt_sid%253D252854%2526mt_exid%253D39%2526mt_inapp%253D0%2526mt_uuid%253Df9b758cd-9429-4200-b2e7-f49cc2649fee%2526mt_lp%253Dhttp%25253A%2F%2Fwww.spartan.com%2Fen%25253Frd_eid%25253Da8e24cb6-323e-11e6-9b41-a840dc54da1e%252526rdmid%25253D711%252526rdadid%25253D2267671%2526redirect%253D%26nuid%3Df9b758cd-9429-4200-b2e7-f49cc2649fee%26ip%3D68.205.115.66%26referrer%3Dhttp%253A%2F%2Fwww.diynetwork.com%2Fhow-to%2Frooms-and-spaces%2Fkitchen%2Fkitchens-on-a-budget-our-10-favorites-from-rate-my-space-pictures%26exchange%3Dsas%26site%3Dwww.diynetwork.com"
	fmt.Println(len([]byte(url)), "bytes")
	curl := smaz.Compress([]byte(url))
	fmt.Println(len([]byte(curl)), len(string(curl)), "bytes compressed")

	match, err := smaz.Decompress(curl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println( string(match) == url)
}
