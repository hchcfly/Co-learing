package main

type Flyer interface {
	fly()
}
type Swimmer interface {
	swim()
}

type FLyFish interface {
	Flyer
	Swimmer
}

type Fish struct {
	
}

func (fish Fish) fly() {

}

func (fish Fish) swim() {
	
}


func main() {
	var ff FLyFish
	ff = Fish{}
	ff.fly()
	ff.swim()
}