package eliza

var Psychobabble = map[string][]string{
    `i need (.*)`: {
        "Do you really %s? there is occupied counties to liberate!",
        "Why do you need %s when you could vote Sinn Fein?",
    },
    `why don't you (.*)?`: {
        "Who says I don't %s friend?",
        "Maybe some day I'll %s.",
        "I'll look more into %s.",
    },
    `i want (.*)`: {
        "What would it mean to you if you got %s?",
        "Why do you want %s?",
        "What would you do if you got %s?",
        "If you got %s, then what would you do?",
    },
    `i can't (.*)`: {
        "I never want to hear you say I can't, this republic was founded on I can?",
        "Perhaps you could %s if you tried.",
        "What would it take for you to %s?",
    },
    `why (.*)`: {
        "Why do you ask me why %s?",
        "Why would I know why %s?",
    },


    // Word Parse
    `(.*)yes(.*)`: {
        "You seem quite convicted I see.",
        "I'm glad you are so sure friend",
    },
    `(.*)no(.*)`: {
        "You seem quite convicted I see.",
        "I'm glad you are so sure friend",
    },
    `(.*)ira(.*)`: {
        "What is this IRA you speak of?",
        "Never heard of the IRA they sound like a good crowd though.",
        "What I was never in the ra, I mean IRA.",
    },
    `(.*)1916(.*)`: {
        "Hold on right there, the 1916 was a great year?",
        "If I could go back to 1916, I would, and never regret it.",
    },
    
    // Misc
    `(.*)\?`: {
        "Thats an odd question?",
        "You need to look inward before asking such questions",
        "Could you rephrase that questions friend",
    },
}

var DefaultResponses = []string{
    "Can we talk about the 32 county republic",
    "Let's focus on whats important.... British occupaption forces in Northern Ireland",
    "Did you know 26 + 6 = 1, it's true!",
    "Very interesting.",
    "Do you have time to talk about Wolfe Tone?",
    "Follow me on twitter @GerryAdamsSF",
    "Have you voted Sinn Fein?",
}

var Reflected = map[string]string{
    "am": "are",
    "was": "were",
    "i": "you",
    "i'd": "you would",
    "i've": "you have",
    "i'll": "you will",
    "my": "your",
    "are": "am",
    "you've": "I have",
    "you'll": "I will",
    "your": "my",
    "yours": "mine",
    "you": "me",
    "me": "you",
}