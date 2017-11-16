const form = $("#userInput");
const list = $("#botOutput");

form.keypress(function(event){
    if(event.keyCode != 13){
        return;
    }

    event.preventDefault();

    const userInput = form.val();
    form.val("");
    
    list.append("<p>" + "<b>User: </b>" + userInput + "</p>");
    
    const queryParams = {"input" : userInput }
    
    $.get("/chat", queryParams)
        .done(function(resp){

            const newItem = "<p><b>Bot: </b>" + resp + "</p>";

            setTimeout(function(){list.append(newItem), new Audio('assets/sound/responseAlert.mp3').play();}, 1000);
            
        }).fail(function(){
            const newItem = "<p><b>Bot: </b> </p>";
            list.append(newItem);
        });
});