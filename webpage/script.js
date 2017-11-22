// Input and output variables
const inputText = $("#userInput");
const ouputText = $("#botOutput");

var singSong = new Audio('assets/sound/gerrySong.mp3');

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
    ouputText.append("<img src=\"/assets/image/user.jpg\" class=\"profileImage\">" + "<p class=\"reply\">" + "<b class=\"userName\">User: </b>" + userInput + "</p>");

    if(userInput.indexOf("sing") > -1){
        setTimeout(function(){singSong.play()}, 1000);
    }
    if(userInput.indexOf("stop") > -1){
        setTimeout(function(){singSong.pause()}, 1000);
        singSong.currentTime = 0;
    };
    
    // When go can comunicate with webpage '.done' will run otherwise '.fail' will
    $.get("/elizaPrompt", userPrompt)
        .done(response => {
            // if the webpage can connect to the GO eliza files then the user will be prompted with a response from eliza
            const output = "<img src=\"/assets/image/profile.png\" class=\"profileImage\">" + "<p class=\"reply\">" + "<b class=\"userName\">Gerry Adams: </b>" + response + "</p>";

            // The bots response with HTML elements are output to the webpage with an alert
            setTimeout(function(){ouputText.append(output), new Audio('assets/sound/responseAlert.mp3').play(),document.getElementById('userInput').scrollIntoView();}, 1000);
            
        }).fail(() => {
            // If the webpage cannot retrieve the GO eliza files the user will be informed
            const output = "<img src=\"/assets/image/profile.png\" class=\"profileImage\">" + "<p class=\"reply\">" + "<b class=\"userName\">Gerry Adams: </b>" + "I am away directing Sinn Féin towards a united Ireland." + "</p>";

            // An alert is ouput with HTML elements informing the user
            setTimeout(function(){ouputText.append(output), new Audio('/assets/sound/responseAlert.mp3').play(),document.getElementById('userInput').scrollIntoView();}, 1000);
        });
});