// Package smaz is an implementation of the smaz library
// (https://github.com/antirez/smaz) for compressing small strings.
package smaz

import (
	"errors"
)

var (
	codeStrings = []string{
		"linksynergy","mediaforge","com","https:","jp-tags","tags","253D","http:","26redirecturl","3Fmerchant",
		"26strategy","duration","26height","prodID","js","26width","3Dhttp","rd","3Dretargeting","?prodID",
		"imp?eID","26nID","catID","2526mt_adid","2526mt_inapp","reqid","2F","253Fmt_aid","253D0","xdom",
		"2526mt_exid","2526mt_uuid","2F11","2526redirect","type","26cb","252Finsight","2526mt_id","253A",
		"2526mt_sid","252520States","253Dhttps","2Famp","2526rlangs","2F2017+07","25253A","26referrer",
		"2526crrelr","2Fjp-amp","26exchange","2F2017+20","253DOther","253DUnited","2526svscid","252Ftrack",
		"2526rcats","2526svpid","2526mlang","253DChrome","2Fclick","252F","2Fpixel","mathtag","3D300","http",
		"mID","cts","url","nID","sID","2526td_s","2526mfsi","2526rgco","2526mfld","2526agsa","2526mste",
		"2526crid","2526uhow","2526rgme","2526mssi","2526mcat","2526rgci","2526rgre","2526daid","2526tmpc",
		"2526vrtd","2526rcxt","2526svsc","aID","2526sig","3D250","?catID","2526mt_lp","2526testid","2526mt_3pck",
		"252Fclk","2526rgz","2526osf","2526dnr","2526osi","2526did","2526osv","2526dur","2526sfe","2526vpb",
		"253Fimp","2526npt","253Dhttp","2526mdl","doubleclick","253Dwith-iplookup","2526ag","adsrvr","3A",
		"2Fimg","2526sv","2526dt","2526br","2526fq","2526os","2526cf","26nuid","2Fwww","https","253DWindows",
		"253Dno-iplookup","26site","3Dprospecting","252C0","3972","3D728","2F2017+12","252520Windows",
		"253DMobileOptimizedWeb","252Frd","252Ftg","252Fv1","prod","3D20","2526r","3Dhttps","253D127","08",
		"on_demandware","253Dwww","253Den","253DMobile","3A28","2Fadclick","3Dashley_furniture_retargeting",
		"252526utm_medium","socdm","3A29","252Fs","3D90","2526mk","252526utm_campaign","3D80",
		"3DINSERT_RANDOM_NUMBER_HERE","252526rdadid","pt","26ip","253DPC","253DWindows10","253DAndroid",
		"category","253D4","252526rdmid","2520","252526client","2F2017+06","25253Futm_source","act?eID",
		"cat","25253D","3A30","252520-","3D49","5744","2FAdDisplayTrackerServlet","253D1","25252F",
		"253Dgoogle","3A27","postcode","25253Ddisplay","252526num","252526sig","253Dcasale","2936",
		"product","org","252526adurl","3282","253DWindows7","253D2","252526rd_eid","3Dpacsun_prospecting",
		"?pt","3Dretargeting-scaleout","4600?reqid","26","2526adpt","253DSafari","5846","25253Fsa",
		"252526ai","2Faclk","3D600","25253D1","253DAndroid70","4601?rmpb","4950","4008","np_banner_model_id",
		"imp?nt","253DInternetExplorer11","252C1","26click_encoding","suburb","253D01","home",
		"np_banner_set_id","net","3538?rmpb","2728","252Fadclick","rubiconproject","253Dopenx",
		"3320?rmpb","253D15","5899?rmpb","20","252526utm_content","cart","3D160","25253FclickData",
		"3537?rmpb","253Fsa","253D3"}

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
