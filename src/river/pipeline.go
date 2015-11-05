
package river

// This file is a central location for all channel definitions
// The functions which actually read/write from these channels are in
// different files

// Stage 1: Get input from the prompt.
var Input = make(chan string, 10)
// Stage 2: Parse and execute commands. Results are sent to stage 3 or stage 4.
var FilteredCommands = make(chan string, 10)
// Stage 3: Sanitize/manipulate input depending on enabled commands. Results are sent to the server.
var SendToServer = make(chan string, 10)
// Stage 4: Display results to the client
var Display = make(chan ConsoleMsg, 10)
