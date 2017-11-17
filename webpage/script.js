// Input and output variables
const inputText = $("#userInput");
const ouputText = $("#botOutput");

// Jquery keypress for user input
inputText.keypress(function(e){

    // While key 13 which is 'Enter' is not pressed keep taking in characters
    // Withouth this if statement the program will run on each character input
    if(e.keyCode != 13){
        return;
    }

    // Prevents page refresh
    e.preventDefault();

    // Variable Setup
    const userInput = inputText.val();
    inputText.val("");
    const userPrompt = {"input" : userInput }
    ouputText.append("<p style=\"color:white\">" + "<b>User: </b>" + userInput + "</p>");
    
    // When go can comunicate with webpage '.done' will run otherwise '.fail' will
    $.get("/elizaPrompt", userPrompt)
        .done(response => {
            // if the webpage can connect to the GO eliza files then the user will be prompted with a response from eliza
            const output = "<p style=\"color:white\">" + "<b>Bot: </b>" + response + "</p>";

            // The bots response with HTML elements are output to the webpage with an alert
            setTimeout(function(){ouputText.append(output), new Audio('assets/sound/responseAlert.mp3').play();}, 1000);
            
        }).fail(() => {
            // If the webpage cannot retrieve the GO eliza files the user will be informed
            const output = "<p style=\"color:white\"><b>ALERT: </b>Bot is offline</p>";

            // An alert is ouput with HTML elements informing the user
            setTimeout(function(){ouputText.append(output), new Audio('assets/sound/responseAlert.mp3').play();}, 1000);
        });
});