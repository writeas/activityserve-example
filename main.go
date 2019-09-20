package main

import (
	"flag"

	"github.com/gologme/log"
	"github.com/writeas/activityserve"
)

var err error

func main() {

	debugFlag := flag.Bool("debug", false, "set to true to get debugging information in the console")
	flag.Parse()

	if *debugFlag == true {
		log.EnableLevel("info")
		log.EnableLevel("error")
	}

	activityserve.Setup("config.ini", *debugFlag)

	// This creates the actor if it doesn't exist.
	actor, _ := activityserve.GetActor("activityserve_test_actor_3", "This is an activityserve test actor", "Service")
	// actor.Follow("https://mastodon.social/users/qwazix")
	// actor.Follow("https://pe.ertu.be/accounts/qwazix")
	actor.Follow("https://pixelfed.social/users/qwazix")
	// actor.Follow("https://pleroma.site/users/qwazix")
	// actor.CreateNote("Hello World!", "")
	// let's boost @tzo's fox
	// actor.Announce("https://cybre.space/@tzo/102564367759300737")

	// this can be run any subsequent time
	// actor, _ := activityserve.LoadActor("activityserve_test_actor_2")
	// actor.CreateNote("I'm building #ActivityPub stuff", "")

	activityserve.Serve()
}
