package eliza

// Libary of phrases
var phrases = map[string][]string{
    // Word Parse
    `ira`: {
        "What is this IRA you speak of?",
        "Never heard of the IRA they sound like a good crowd though.",
        "What I was never in the ra, I mean IRA.",
    },
    `1916`: {
        "Hold on right there, 1916 was a great year.",
        "If I could go back to 1916, I would, and never regret it.",
    },
    `tiocfaidh`: {
        "tiocfaidh ar la comrade",
        "Ho ha up the rá",
    },
    `craic`: {
        "Craic agus ceol, mo chara",
    },
    `uk (.*)`: {
        "Never speak of the UK or any of it's members to me again",
        "Please never talk about this topic again.",
        "I really don't want to talk about '%s'",
    },
    `monarchy (.*)`: {
        "Ah, please don't talk about %s.",
    },


    `i need (.*)`: {
        "You say you need %s. But really what you need is a united Ireland friend.",
        "In a country founded on the ideals of the 1916 rising you wouldn't need %s.",
        "Do you really need %s, you could pick upa hobby, like learning Gaeilge mo chara",
        "All you capatalist pigs are the same. %s when you could give, give back the country join a group like the IR... I mean sinn Féin youth.",
        "Do you really %s? there is occupied counties to liberate!",
        "Why do you need %s when you could vote Sinn Fein?",
        "There would be no need in a communist society friend.",
    },
    `i can't (.*)`: {
        "If everyone in 1916 had your attitude we'd still be under a union jack.",
        "Maybe you could %s, but you'd need to but a bit of effort into it.",
        "You can %s, you just need inspiration, which you can find in 'My Little Book of Tweets' on amazon for only €8.00!",
    },
    `i am (.*)`: {
        "You are %s? In a free republic that matters very little friend.",
        "You say you are %s, but can you prove that?",
    },
    `I think (.*)`: {
        "Do you think %s? Or is that just a guess?",
        "That's a fair thought so it is.",
        "Can you confirm this?",
    },
    `i feel (.*)`: {
        "You feel %s, but I feel the pain of a divided Ireland, and won't sleep until it's repaired!",
        "When I feel %s, I go to the local and forget about it.",
        "You might feel %s, but that is only temporary. Soon Ireland will be united under Sinn Féin, and all your worries will be gone.",
    },
    `i have (.*)`: {
        "You might have %s. But I have a great arsenal.... season pass to all their games. Not in the IRA.",
        "Do you really have %s. You need ot provide evidence otherwise people will say you were in the ra.",
        "Well I have some great memes. Now what have you got to say?",
    },
    `i would (.*)`: {
        "You say you would %s. But what would that accomplish in the grand scale of a Sinn Féin Ireland?",
        "Why would you %s?",
        "Who else knows that you would %s?",
    },
    `i don't (.*)`: {
        "I heard you do %s, can't lie to the Gerry.",
        "Why don't you %s?",
    },

    `why don't you (.*)?`: {
        "Who says I don't %s friend?",
        "Maybe some day I'll %s.",
        "I'll look more into %s.",
    },
    `why don't you (.*)`: {
        "You don't know me well friend, if you think I don't %s.",
        "I already am %s. Don't tell the media, don't want them to ruin my next snapchat story.",
        "If I %s will you never ask me again?",
    },
    `why can't I (.*)`: {
        "You can, if you try harder. I believe you can %s.",
        "You can %s! But you just need to apply it on a smaller scale friend. 1916 was a small event in the grand scale of a united Ireland.",
        "Lets say you can't %s. What else could you improve on?",
    },
    `are you (.*)`: {
        "Never ask if I am %s, because this is not an interview.",
        "You don't ask a woman her age, you don't ask a Gerry if he is %s. Them are the rules friend.",
        "Please don't ask if I am %s. I will never give a straight answer.",
    },
    `how (.*)`: {
        "If you are asking me 'how %s' then you have come to the wrong man for the job if you need an semi automatic.... car I can help with that.",
        "Look inward and you may be able to find the answer. If not you could always vote Sinn Féin we will fix it before it's a problem.",
        "What is it you're really asking friend?",
        "I find google could answer this question give it a go on 'www.google.ie'. not on .co.uk.",
    },
    `is it (.*)`: {
        "How do you know it is %s?",
        "is it %s?",
    },
    `can you (.*)`: {
        "Where did you hear I can %s?",
        "Did the IR... Sinn Féin member tell you I can %s?",
        "That's some tall tail. Who told you I can %s?",
    },
    `can I (.*)`: {
        "Can you %s?",
        "If you could %s, how would you put that into action to reclaim the north?",
        "Say you could %s, would Ireland be any better off?",
    },
    `is there (.*)`: {
        "Why would you ask is there %s?",
        "Maybe there is %s. We can't guess these things.",
        "Would you like there to be %s?",
    },
    `why (.*)`: {
        "You're asking why %s. When the real question is when will the Scotland leave the UK and ultimately destroy the UK and it's fleg.",
        "You ask why %s? When you should be asking yourself why you ask these question.",
    },

    `what (.*)`: {
        "Why are you asking me 'what %s'. This sounds likea questions you should ask a higher power my friend.",
        "'what %s' is a weird question to ask a politician, I think a better question is why are you asking one?",
        "I find google could answer this question give it a go on 'www.google.ie' not that .co.uk one.",
    },
    
    `because (.*)`: {
        "You are making excuses friend?",
        "Arlene Foster wouldn't even fall for such an excuse friend.",
        "If %s, is your response then stick with it friend.",
    },
    `sorry`: {
        "No apologies needed.",
        "Please don't apologies, there is no need.",
    },
    `hello`: {
        "Hello comrade, any craic?",
        "Hey how are you friend?",
        "Dia ghuit aon scéal agut?",
    },
    `yes`: {
        "Conviction, I like it. We need more like you in the IR.... Sinn Féin youth, sign up today!",
        "You seem certain. People who are certain with themselfs make great shots with..... a basketball.",
    },
    `it is (.*)`: {
        "Conviction a great trait!",
        "Although conviction is great aways have a plan in case it isin't %s.",
    },
    `you are (.*)`: {
        "Is that an accusation?",
        "If the pres find out I'm %s, I'll deny it. Just like the ra, which I was never part of.",
        "Would I lose snapchat followers if they found out I am %s?",
        "So what if I'm %s.",
    },
    `my (.*)`: {
        "When your %s, does it make you wish for a united Ireland even more?",
    },
    `you (.*)`: {
        "Please I am not being interviewed by someone I call a friend.",
        "Why do you care whether I %s?",
        "We should be discussing you, not me friend.",
    },
    `(.*)\?`: {
        "Thats an odd question?",
        "Could you rephrase that questions friend",
        "You're not a Gardai are you?",
        "Are you from the press?",
    },
}

// Failing to find an answer in the libary above this libary will be used
var conversationFail = []string{
    "Can we talk about the 32 county republic",
    "Let's focus on whats important.... British occupaption forces in Northern Ireland",
    "Did you know 26 + 6 = 1, it's true!",
    "Tell me more.",
    "Do you have time to talk about Wolfe Tone?",
    "Follow me on twitter @GerryAdamsSF",
    "Have you voted Sinn Fein?",
    "Buy my books on Amazon not on Amazon.co.BritishTyranny though",
    "I make the spiciest memes around",
    "I'm an open book, ask me anything!",
}

// libary to change direction of conversation so bot can reply and talk to you
var Reflected = map[string]string{
    "am":       "are",
    "was":      "were",
    "I":        "you",
    "I'd":      "you would",
    "I've":     "you have",
    "I'll":     "you will",
    "my":       "your",
    "are":      "am",
    "you've":   "I have",
    "you'll":   "I will",
    "you're":   "I'm",
    "your":     "my",
    "yours":    "mine",
    "you":      "I",
    "me":       "you",
}

// Makes list of synomyms so the user can input multiple "hello" variations for example and still be understood
var Synonymizer = map[string][]string{
    // Expressions
    "want":     []string{"need", "require", "demand"},
    "i am":     []string{"i'm"},
    "you are":  []string{"you're"},
    "sorry":    []string{"regretful","apologies","pardon","retract","atone","forgiveness"},
    "thanks":   []string{"regretful","apologies","pardon","retract","atone","forgiveness"},
    "hello":    []string{"hey","hi","greetings","salutations","howya"},
    
    // Words
    "ira":      []string{"ra"},
    "uk":       []string{"england","kingdom","british","union"},
    "gerry":    []string{"adams","madman"},
    "monarchy": []string{"royals","royal","queen","king","princess","prince","dutchess","duke","monarch"},
}