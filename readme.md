# Activityserve-example is a simple go app that creates and actor and posts "hello" to the world

This file just illustrates the usage philosophy of activityserve. More complicated examples (might) come later.

For now you will just have to set up your actor using `MakeActor` (which creates the storage backend and saves your actor to disk) and then use that actor to perform the familiar ActivityPub actions (for example `actor.follow("https://fosstodon.org/users/qwazix")` or `actor.CreateNote("Hello World", "")`)