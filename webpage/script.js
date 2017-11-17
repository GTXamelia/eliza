const form = $("#userInput");
const list = $("#botOutput");

form.keypress(function(e){

    // While key 13 which is 'Enter' is not pressed keep taking in characters
    // Withouth this if statement the program will run on each character input
    if(e.keyCode != 13){
        return;
    }

    // Prevents page refresh
    e.preventDefault();

    const userInput = form.val();
    form.val("");
    
    list.append("<p>" + "<b>User: </b>" + userInput + "</p>");
    
    const queryParams = {"input" : userInput }
    
    $.get("/chat", queryParams)
        .done(resp => {

            const newItem = "<p><b>Bot: </b>" + resp + "</p>";
            setTimeout(function(){list.append(newItem), new Audio('assets/sound/responseAlert.mp3').play();}, 1000);
            
        }).fail(() => {

            const newItem = "<p><b>Bot is offline</b> </p>";
            list.append(newItem);
        });
});