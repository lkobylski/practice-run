# Text.com practice-run 🚀

## Task:
The goal is to create a simple server in the Go language, similar to Slack, with rooms/channels, participants, and an option to send a text message.

As for the client, you can use the developer console built into the web browsers.

There is no need to create your own client solution.
The service must allow clients to connect via the WebSocket interface and support a few commands like create a room, join a room, leave a room, and send a message in a room - it should be broadcast to other room participants.
You’ll get a draft of the service.
There are a few bugs or performance deficiencies.
Cooperate with your buddy, and try to find and correct these defects.

Please remember to upload your solution to a Git server (Github or Gitlab) at the end of the practice run before the summary meeting.

At the summary meeting, please present your application (live demo), describe the source code in detail, the challenges you’ve faced, and ideas you’d like to implement if you had more time.

You can expect some questions about the code.

## Requirements:
* with rooms/channels
* participants
* option to send a text message
* WebSocket
* commands: create a room, join a room, leave a room
* send a message in a room
Commands:

* /room new-nice-room
* /join new-nice-room
* /leave new-nice-room

general room?
