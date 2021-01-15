package main

import(
	//"fmt"
	"github.com/fsayed/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func playList(player Player, songs []string) {
	for _, song := range songs {
		player.Play(song)
	}
	player.Stop()
	//NOTE: Type is a struct but do not include {}
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		for _, song := range songs {
			recorder.Record(song)
		}
	}

}

func main() {
	//player := gadget.TapePlayer{}
	
	//Now, TapeRecorder type also has a Play() and Stop() methods.
	//player := gadget.TapeRecorder{}
	mixTape := []string{"November Rain", "Wheat Kings", "Shot through the heart"}

	//The problem is that even though TapePlayer and TapeRecoder have the same functions
	//We can't pass type Recoder as type Player. They are different. We want to solve this
	//so that doesn't matter what the type is, as long as it has the function it should be able to
	//call that function.
	//We solve this problem using interfaces
	//An interface is a set of methods that a certain value is expected to have
	playList(gadget.TapeRecorder{}, mixTape)
	playList(gadget.TapePlayer{}, mixTape)


	
}
