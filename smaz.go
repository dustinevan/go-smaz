// Package smaz is an implementation of the smaz library
// (https://github.com/antirez/smaz) for compressing small strings.
package smaz

import (
	"errors"
)

var (
	codeStrings = []string{
		"tags.rd.linksynergy.com","jp-tags.mediaforge.com","tags.mediaforge.com","https:","253D","http:",
		"2Famp.rd.linksynergy.com","26redirecturl","3Fmerchant","26strategy","2Fjp-amp.mediaforge.com",
		"duration","26height","prodID","2Fpixel.mathtag.com","js","252Finsight.adsrvr.org",
		"jp-tags.rd.linksynergy.com","26width","3Dhttp","3Dretargeting","?prodID","imp?eID","26nID","catID",
		"2526mt_adid","2526mt_inapp","252Ftg.socdm.com","reqid","2F","253Fmt_aid","253D0","xdom","2526mt_exid",
		"2526mt_uuid","2F11","2526redirect","type","26cb","2526mt_id","253A","2Fadclick.g.doubleclick.net",
		"2.26.8","2526mt_sid","252520States","253Dhttps","2526rlangs","2F2017+07","25253A","26referrer",
		"2526crrelr","26exchange","2F2017+20","253DOther","253DUnited","2526svscid","252Ftrack","2526svpid",
		"2526rcats","2526mlang","253DChrome","2Fclick","252F","3D300","http","mID","url","cts","nID","sID",
		"2526mfsi","2526mste","2526mfld","2526mssi","2526rgci","2526rgco","2526rgre","2526td_s","2526rgme",
		"2526mcat","2526uhow","2526agsa","2526crid","2526vrtd","2526tmpc","2526daid","2526rcxt","2526svsc",
		"aID","2526sig","3D250","?catID","2526mt_lp","2526testid","2526mt_3pck","2Famp.mediaforge.com",
		"252Fclk","2526rgz","2526osf","2526did","2526osi","2526dur","2526osv","2526dnr","2526sfe","2526vpb",
		"253Fimp","2526npt","253Dhttp","2526mdl","253Dwith-iplookup","2526ag","3A","2Fimg","2526fq","2526sv",
		"2526dt","2526os","2526br","2526cf","on_demandware.store","26nuid","https","253DWindows",
		"253Dno-iplookup","26site","3Dprospecting","3972","3D728","2F2017+12","252520Windows",
		"252Fadclick.g.doubleclick.net","253DMobileOptimizedWeb","252C0","252Fv1","252Frd","prod","3D20",
		"2526r","3Dhttps","253D127","08","253Den","253DMobile","3A28","3Dashley_furniture_retargeting",
		"252526utm_medium","2.21.6","3A29","252Fs","2526mk","252526utm_campaign","3D90","3D80",
		"3DINSERT_RANDOM_NUMBER_HERE","252526rdadid","pt","26ip","253DPC","253DWindows10","253DAndroid",
		"category","253D4","252526rdmid","2520","252526client","2F2017+06","25253Futm_source","act?eID",
		"cat","25253D","3A30","252520-","5744","3D49","2Fclicktrack.pubmatic.com","2FAdDisplayTrackerServlet",
		"253D1","25252F","3A27","postcode","252526num","252526sig","253Dcasale","2936","product","252526adurl",
		"3282","253DWindows7","253D2","252526rd_eid","3Dpacsun_prospecting","?pt","3Dretargeting-scaleout",
		"4600?reqid","2526adpt","253DSafari","5846","25253Fsa","252526ai","2Faclk","25253Ddisplay",
		"25252Ftag.ladsp.com","3D600","25253D1","253DAndroid70","4601?rmpb","4950","4008","np_banner_model_id",
		"253Dgoogle","imp?nt","253DInternetExplorer11","252C1","26click_encoding","suburb","253D01",
		"253Dgoogle.site-not-provided","home","np_banner_set_id","3538?rmpb","2728","253Dopenx","3320?rmpb",
		"253D15","5899?rmpb","252526utm_content","cart","20","3D160","2Fwww.autoline.com.br","25253FclickData",
		"3537?rmpb","253Fsa","253D3","2526client","253Furl","np_campaign_id","?orderNumber","3Dautoline-branding",
		"5847?rmpb","5940?rmpb","2526ai","3Drtenpo6875","253DApple","5939?rmpb","25253Drakuten_segmentado",
		"2Fbeacon-us-iad2.rubiconproject.com","state","4631","3Ddennis_kirk_search","2Fwww.hunterboots.com",
		"253D126","252Fads.simpli.fi","25253Dconvencional_branding-autoline_","bathrooms","2526adurl","2F2017+21",
		"?reqid","5619","3Dcas","25253Dl","2F2017+08","np_client_id","4337?jsonp_callback","5900","253Dnopx",
		"4876","25253Frd_eid","25253Drakuten","253DiOS","80","252Fc.dr.adingo.jp","253Ddynamic_display",
		"Lacoste_Sale","bedrooms","252Faclk","3Dfortis_edu_prospecting","253DFirefox","L1212-51,PF6949-51,L1264-51",
		"3Damiami4740","eng?eID","3Drtenpo6867","2Fbeacon-us-west.rubiconproject.com","3Dmobile-deep-forest",
		"3DA-GaboppkF4-GfewtthlNco5zW3iVKwg","sa_saclearance_saclearancewomen","3787","area","3D320","5061",
		"3Dashley_furniture_prospecting","2F2017+04","price","253DPubmatic_OpenRTB","252Fg","2Fwww.fortis.edu",
		"252FU","2Fjp-amp.rd.linksynergy.com","2526num","2FAdServer","253DIllinois","SALESCAT_ALL-BACKPACKS",
		"253DUnknown","253Dappnexus","253Dpubmatic","2305","253D136067","25252Fclk","253DAndroid60","252526utm_term",
		"252526utm_source","25253Dautoline_onedigital","253Ddesktop","49","253Dbrowser","4953?rmpb","252Fnxtck.com",
		"template-chaordic","np_height","5110?rmpb","click.rd.linksynergy.com","3Dautoli6596","253DiPhone",
		"253DSamsung","group1","253DNew","25253D_url","5022?rmpb","4952?rmpb","3Dpanasonic","253DFlorida",
		"2Fgoogle-bidout-d.openx.net","3Dspartan_race_mobile","25253DCL3G86_KnNQCFUujaAod36EIsw","2526utm_campaign",
		"2Fbeacon-eu-ams3.rubiconproject.com","252520Explorer","25252Fadclick.g.doubleclick.net","eID","4600",
		"orderNumber","2Fbeacon","3D50","25253Dretargeting","4951?rmpb","253DMassachusetts","3Dmbl-retargeting",
		"3DHunter6372","253D2376431","2526utm_medium","2.24.8","253DPennsylvania","3Dtoshiba","2938","252520York",
		"np_ad_id","np_width","5009?rmpb","2124","2Fwww.dailymail.co.uk","2526c3braK3c","252520Kingdom","undefined",
		"252520Carolina","4718?reqid","25252Frd.adingo.jp","3Dautoline-prospeccao","253D188942",
		"252Fnym1-ib.adnxs.com","2526referer","252C2619418","253Dl","252Fdragonage.wikia.com","253Fttd_r","3062?rmpb",
		"253DInternet","2873","25253Fm","6043?rmpb","6044?rmpb","3Dsearch","3Dstaples_canada_crm","253Dy29",
		"6313?rmpb","252Fwww.petsathome.com","3Dkaren_millen_uk","3063?rmpb","other","253Dadconductor","5369?rmpb",
		"4574","6381?rmpb","253D183295","3Dhomepage","3DINSERT_INSERT_CLICKURL_HERE_HERE","5003?rmpb","3325","5007",
		"253Futm_source","252Ftag.ladsp.com","6738?rmpb","253DTablet","25253DL","5609","3652",
		"3Dfortis_edu_retargeting","3064?rmpb","253DCanada","3043?rmpb","3Dretargeting-english","3Dcontextual",
		"3Dwww.dailymail.co.uk","3Dapn","253D6432","25253Fts","253D142878","4358","252Findex","253DCalifornia",
		"252C8180","2526turl","253Fsifi"}

	codes    = make([][]byte, len(codeStrings))
	codeTrie trieNode
)

func init() {
	for i, code := range codeStrings {
		codes[i] = []byte(code)
		codeTrie.put([]byte(code), byte(i))
	}
}

// A trieNode represents a logical vertex in the trie structure.
// The trie maps []byte -> byte.
type trieNode struct {
	branches [256]*trieNode
	val      byte
	terminal bool
}

// put inserts the mapping k -> v into the trie, overwriting any previous value.
// It returns true if the element was not previously in t.
func (n *trieNode) put(k []byte, v byte) bool {
	for _, c := range k {
		next := n.branches[int(c)]
		if next == nil {
			next = &trieNode{}
			n.branches[c] = next
		}
		n = next
	}
	n.val = v
	if n.terminal {
		return false
	}
	n.terminal = true
	return true
}

func flushVerb(out, verb []byte) []byte {
	// We can write a max of 255 continuous verbatim characters,
	// because the length of the continuous verbatim section is represented
	// by a single byte.
	var chunk []byte
	for len(verb) > 0 {
		if len(verb) < 255 {
			chunk, verb = verb, nil
		} else {
			chunk, verb = verb[:255], verb[255:]
		}
		if len(chunk) == 1 {
			// 254 is code for a single verbatim byte.
			out = append(out, 254)
		} else {
			// 255 is code for a verbatim string.
			// It is followed by a byte containing the length of the string.
			out = append(out, 255, byte(len(chunk)))
		}
		out = append(out, chunk...)
	}
	return out
}

// Compress compresses a byte slice and returns the compressed data.
func Compress(input []byte) []byte {
	out := make([]byte, 0, len(input)/2) // estimate output size
	var verb []byte

	for len(input) > 0 {
		prefixLen := 0
		var code byte
		n := &codeTrie
		for i, c := range input {
			next := n.branches[int(c)]
			if next == nil {
				break
			}
			n = next
			if n.terminal {
				prefixLen = i + 1
				code = n.val
			}
		}

		if prefixLen > 0 {
			input = input[prefixLen:]
			out = flushVerb(out, verb)
			verb = verb[:0]
			out = append(out, code)
		} else {
			verb = append(verb, input[0])
			input = input[1:]
		}
	}
	return flushVerb(out, verb)
}

// ErrDecompression is returned when decompressing invalid smaz-encoded data.
var ErrDecompression = errors.New("invalid or corrupted compressed data")

// Decompress decompresses a smaz-compressed byte slice and return a new slice
// with the decompressed data.
// err is nil if and only if decompression fails for any reason
// (e.g., corrupted data).
func Decompress(b []byte) ([]byte, error) {
	dec := make([]byte, 0, len(b)) // estimate initial size

	for len(b) > 0 {
		switch b[0] {
		case 254: // verbatim byte
			if len(b) < 2 {
				return nil, ErrDecompression
			}
			dec = append(dec, b[1])
			b = b[2:]
		case 255: // verbatim string
			if len(b) < 2 {
				return nil, ErrDecompression
			}
			n := int(b[1])
			if len(b) < n+2 {
				return nil, ErrDecompression
			}
			dec = append(dec, b[2:n+2]...)
			b = b[n+2:]
		default: // look up encoded value
			dec = append(dec, codes[int(b[0])]...)
			b = b[1:]
		}
	}

	return dec, nil
}
