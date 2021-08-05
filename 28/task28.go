package main

import "fmt"

type EnglishSpeaker struct {
}

type EnglishListener interface {
	listenToEnglishSpeech()
}

type RussianListener struct {
}

type EnglishToRussianAdapter struct {
	rus *RussianListener
}

type BritishListener struct {
}

func (e *EnglishSpeaker) SpeakEnglish(listener EnglishListener) {
	fmt.Println("English speaker speaks english.")
	listener.listenToEnglishSpeech()
}

func (b *BritishListener) listenToEnglishSpeech() {
	fmt.Println("British listener understands English speech.")
}

func (r *RussianListener) listenToRussianSpeech() {
	fmt.Println("Russian listener understands Russian speech.")
}

func (EtoR *EnglishToRussianAdapter) listenToEnglishSpeech() {
	fmt.Println("Adapter hears English and translates English to Russian.")
	EtoR.rus.listenToRussianSpeech()
}

func main() {
	EnglishSpeaker := &EnglishSpeaker{}
	BritishListener := &BritishListener{}

	EnglishSpeaker.SpeakEnglish(BritishListener)
	fmt.Printf("---\n")

	RussianListener := &RussianListener{}
	EnglishToRussianAdapter := &EnglishToRussianAdapter{
		rus: RussianListener,
	}

	EnglishSpeaker.SpeakEnglish(EnglishToRussianAdapter)
}
