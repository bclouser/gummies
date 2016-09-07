/*
	The Thought for this file is to connect to any external
	modules of code written in any language. 

	It will talk to a single messaging entrypoint (maybe using zmq) which any
	thread of execution in any language can be listening. Only one application should be
	setup to respond to a single message and ignore all others.

	The called application can do whatever it desires and return whatever it wants.
	The response will be sent all the way back up the stack to whatever interface
	originally sent the request, whether web, socket, mqtt, etc.

	This application will have a finely tuned timeout to ensure that the called 
	application doesn't hang this program. We should also ensure nothing malicious is returned.
*/