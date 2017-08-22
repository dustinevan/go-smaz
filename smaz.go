// Package smaz is an implementation of the smaz library
// (https://github.com/antirez/smaz) for compressing small strings.
package smaz

import (
	"errors"
)

var (
	codeStrings = []string{
		"tags.rd.","linksynergy.com","jp-tags.","mediaforge.com","tags.","https://","http://","amp.rd.",
		"linksynergy.com","redirecturl","merchant","strategy","jp-amp.","duration","height","prodID",
		"pixel.mathtag.com","/js/","insight.adsrvr.org","width","retargeting","?prodID","imp?eID","nID",
		"catID","mt_adid","mt_inapp","tg.socdm.com","reqid","mt_aid","mt_exid","mt_uuid","redirect","type",
		"mt_id","adclick.g.doubleclick.net","mt_sid","States","https","rlangs","referrer","crrelr",
		"exchange","Other","United","svscid","track","svpid","rcats","mlang","Chrome","click","?catID",
		"testid","mt_3pck","amp.mediaforge.com","with-iplookup","on_demandware.store","Windows","no-iplookup",
		"prospecting","3D728","MobileOptimizedWeb","ashley_furniture_retargeting","utm_medium","campaign",
		"INSERT_RANDOM_NUMBER_HERE","Android","category","clicktrack.pubmatic.com","AdDisplayTrackerServlet",
		"postcode","product","pacsun_prospecting","scaleout","4600?reqid","Safari","display","tag.ladsp.com",
		"np_banner_model_id","google","InternetExplorer","click_encoding","suburb","google.site-not-provided",
		"home","np_banner_set_id","tm_content","cart","20","3D160","%25","252","253","53D","526","D%2","3D%","26r",
		"d%2","%3D","26m","%26","%2F","0%2","ID=","e%2","26s",".co","id%","ttp","htt","r%2","com","s%2","t%2",
		"ags","i%2","g%2","26d","6rg","%2","25","52","26","3D","53","D%","%3","2F","rg","d%","6r","ht","id","ID",
		"6m","co","0%","t%","at","D=","e%","om","6s","er",".c","tt","r%","tp","ag","%252","%253","253D","2526",
		"3D%2","53D%","D%25","d%25","526r","526m","e%25","http","r%25","id%2",".com","526s","s%25","t%25","i%25",
		"26rg","526d","252F","g%25","0%26","tags","ags.","26os","526o","26sv","com/","%253D","%2526","253D%",
		"53D%2","D%252","3D%25","2526r","2526m","d%253","id%25","2526s","r%253","2526d","526rg","%252F",
		"t%253","tags.","s%253","i%253","e%253","2526o","526os",".com/","526sv","link"}


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
