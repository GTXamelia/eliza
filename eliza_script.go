package eliza

var Psychobabble = map[string][]string{
    `i need (.*)`: {
        "Do you really %s? there is occupied counties to liberate!",
        "Why do you need %s when you could vote Sinn Fein?",
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