package eliza

// High priority last, low priotiry first
// Libary of phrases
var phrases = map[string][]string{
    `i need (.*)`: {
        "You say you need %s. But really what you need is a united Ireland friend.",
        "In a country founded on the ideals of the 1916 rising you wouldn't need %s.",
        "Do you really need %s, you could pick upa hobby, like learning Gaeilge mo chara",
        "All you capatalist pigs are the same. %s when you could give, give back the country join a group like the IR... I mean sinn Féin youth.",
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
    `i can't (.*)`: {
        "If everyone in 1916 had your attitude we'd still be under a union jack.",
        "Maybe you could %s, but you'd need to but a bit of effort into it.",
        "You can %s, you just need inspiration, which you can find in 'My Little Book of Tweets' on amazon for only €8.00!",
    },
    `i am (.*)`: {
        "You are %s? In a free republic that matters very little friend.",
        "You say you are %s, but can you prove that?",
    },
    `are you (.*)`: {
        "Never ask if I am %s, because this is not an interview.",
        "You don't ask a woman her age, you don't ask a Gerry if he is %s. Them are the rules friend.",
        "Please don't ask if I am %s. I will never give a straight answer.",
    },
    `what (.*)`: {
        "Why are you asking me 'what %s'. This sounds likea questions you should ask a higher power my friend.",
        "'what %s' is a weird question to ask a politician, I think a better question is why are you asking one?",
        "I find google could answer this question give it a go on 'www.google.ie' not that .co.uk one.",
    },
    `how (.*)`: {
        "If you are asking me 'how %s' then you have come to the wrong man for the job if you need an semi automatic.... car I can help with that.",
        "Look inward and you may be able to find the answer. If not you could always vote Sinn Féin we will fix it before it's a problem.",
        "What is it you're really asking friend?",
        "I find google could answer this question give it a go on 'www.google.ie'. not on .co.uk.",
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
    `I think (.*)`: {
        "Do you think %s? Or is that just a guess?",
        "That's a fair thought so it is.",
        "Can you confirm this?",
    },
    `yes`: {
        "Conviction, I like it. We need more like you in the IR.... Sinn Féin youth, sign up today!",
        "You seem certain. People who are certain with themselfs make great shots with..... a basketball.",
    },
    `is it (.*)`: {
        "How do you know it is %s?",
        "is it %s?",
    },
    `it is (.*)`: {
        "Conviction a great trait!",
        "Although conviction is great aways have a plan in case it isin't %s.",
    },
    `can you (.*)`: {
        "Where did you hear I can %s?",
        "Did the IR... Sinn Féin member tell you I can %s?",
        "That's some tall tail. Who told you I can %s?",
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
    "sorry":    []string{"regretful","apologies","pardon","retract","atone","forgiveness"},
    "thanks":   []string{"regretful","apologies","pardon","retract","atone","forgiveness"},
    "hello":    []string{"hey","hi","greetings","salutations","howya"},
    
    // Words
    "ira":      []string{"ra"},
    "uk":       []string{"england","kingdom","british","union"},
    "gerry":    []string{"adams","madman"},
    "monarchy": []string{"royals","royal","queen","king","princess","prince","dutchess","duke","monarch"},
}