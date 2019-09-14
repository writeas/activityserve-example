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

	// This should be run only once, it creates the actor
	actor, _ := activityserve.MakeActor("activityserve_test_actor_2", "This is an activityserve test actor", "Service")
	actor.Follow("https://fosstodon.org/users/qwazix/")
	actor.CreateNote("Hello World!")

	// this can be run any subsequent time
	actor, _ := activityserve.LoadActor("activityserve_test_actor_2")
	actor.CreateNote("I'm building #ActivityPub stuff", "")

	activityserve.Serve()
}
