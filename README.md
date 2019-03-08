[![Go Report Card](https://goreportcard.com/badge/github.com/anselb/agile-helper)](https://goreportcard.com/report/github.com/anselb/agile-helper)

# Agile Helper

A Go tool to help you do agile development

## What does Agile Helper do?
Agile Helper will (eventually) help you do these two things:
- Take your Trello or GitHub Project board tasks and place them on your calendar within certain specifications such as:
   - Avoiding other events already on your calendar
   - What time you start and end work
   - If you work on the weekends
- Calculate the total number of points committed in at the start of a sprint, the number of points completed at the end of a sprint, and your average points completed per sprint.

## How to Get Started
- Clone the repo locally and make sure you `go get` all of the required imports at the top of `server.go`.
- Log in to https://developers.trello.com/ and create a key and a token at https://trello.com/app-key.
- Copy the key and token to their respective variables in `sample.env`. Then, delete "sample" so that the file is `.env`.
- Set the variables on line 60 and 61 to their corresponding parts on your Trello board.
   - `boardID` is the board id and can be found in the url when you are on your Trello board.
   - `listName` is the name of your list that you would like to read the cards from.
- The cards should named in the format, number - task (Ex. "4 - set up auth"), and the numbers should be less than 10.
- Run the command `go run server.go` in your terminal, and check the sqlite database to see if everything worked.
