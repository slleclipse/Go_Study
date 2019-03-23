package mp

import "fmt"

type Player interface {
	Play(source string)
}

func Play(source string, ptype string) {
	var p Player
	switch ptype {
		case "MP3":
			p = &MP3Player{}
		case "WAV":
			p = &WAVPlayer{}
		default:
			fmt.Println("Unsorported music type:", ptype)
	}
	p.Play(source)
}

