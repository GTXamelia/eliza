# Eliza bot created in the 'GO' programming language

### Description
This Eliza bot was created in the 'GO' programming language. This Eliza bot has the personality of an Irish politician Gerry Adams. This bot is served using a HTML webpage and using jQuery as a 'bridge' to connect it with the Eliza bot.

#### N.B This bot was created as a satirical view on Gerry Adams and the Sinn Féin party.
#### Personality was chosen to make user interaction more entertaining.


### Instructions
- Clone repository.
- Open CMD or Terminal
- CD into cloned repository
- CD into 'main' folder inside cloned repository
- Using CMD run 'go run elizaMain.go' or 'go build elizaMain.go'
  - Run the excutable 'elizaMain.exe'
- Go to address '127.0.0.1:8080' or 'localhost:8080'

### Operation (Step-by-Step)
> - User connects to '127.0.0.1:8080' or 'localhost:8080'
>- elizaMain.go takes http request and directs user
>- User is shown a HTML webpage which graps settings from styles.css and jquery/javascript from scripts.js
>- elizaMain.go grabs user input from scripts.js when user clicks enter
>- elizaMain.go passes input to elizaBrain.go which processes the input trading information with elizaVocab.go
>- Once input has finished, elizaBrain.go returnes formated input
>- When elizaMain.go has received the formated input it is received as output
>- Ouput is then passed to scripts.js and is ouput to the HTML page by appending to the page
>- The user is then prompted by a sound, picture, bots name and the response.
>- The user can then respond to the bot to continue the conversation

### File Structure
- main 
  - elizaMain.go
- webpage
  - assets
    - image
      - background.jpg
      - profile.png
      - user.jpg
    - sound
      - gerrySong.mp3
      - responseAlert.mp3
  - index.html
  - script.js
  - style.css
- LICENCE
- README.md
- elizaBrain.go
- elizaVocab.go

