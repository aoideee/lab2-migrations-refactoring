package handlers

import (
	"net/http"
)

//Required routes (Home, About, Contact)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello~ My name is Tysha Daniels, and this is my lab on building a Go application. On to more interesting things,\nfor the semester project I decided to go with the Community Library Management System.\nAnd my reasoning is quite simple. I just really like books, so I was naturally drawn to this topic.\nBy choosing something I enjoy, I am hoping that I will stay motivated and engaged throughout the semester.\n"))
}
	
func About(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I am 19 years old, and am currently in my 2nd semester of my third year of the Bachelor's in IT program here at the University of Belize.\nI was born and raised here in Belmopan, and my favourite pastimes are either reading or sleeping :)\n"))
}

func Contact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You can reach me via email at 2023158020@ub.edu.bz.\nAnd my GitHub username is aoideee.\n"))
}

//Creative route (Quote)

func Quote(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Man is not worried by real problems so much as by his imagined anxieties about real problems. -Epictetus\n"))
}