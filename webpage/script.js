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

    const userInput = inputText.val();
    inputText.val("");
    
    ouputText.append("<p>" + "<b>User: </b>" + userInput + "</p>");
    
    const queryParams = {"input" : userInput }
    
    // When go can comunicate with webpage '.done' will run otherwise '.fail' will
    $.get("/elizaPrompt", queryParams)
        .done(resp => {
            
            const output = "<p><b>Bot: </b>" + resp + "</p>";
            setTimeout(function(){ouputText.append(output), new Audio('assets/sound/responseAlert.mp3').play();}, 1000);
            
        }).fail(() => {
            const output = "<p><b>Bot is offline</b> </p>";
            setTimeout(function(){ouputText.append(output), new Audio('assets/sound/responseAlert.mp3').play();}, 1000);
        });
});