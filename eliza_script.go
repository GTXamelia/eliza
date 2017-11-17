package eliza

// High priority last, low priotiry first
// Libary of phrases
var phrases = map[string][]string{

    // Replys
    `because (.*)`: {
        "You are making excuses friend?",
        "Arlene Foster wouldn't even fall for such an excuse friend.",
        "If %s, is your response then stick with it friend.",
    },
    `how (.*)`: {
        "how is that even a question friend?",
        "if %s, is a question I could answer I wouldn't friend. I don't have time.",
    },

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
    
    // Misc
    `(.*)\?`: {
        "Thats an odd question?",
        "You need to look inward before asking such questions",
        "Could you rephrase that questions friend",
    },
    `why (.*)`: {
        "Why do you ask me why %s?",
        "You ask why %s? When you should be asking yourself why you ask these question.",
    },
    `you (.*)`: {
        "We should be discussing you, not me friend.",
    },

    // About user
    `i need (.*)`: {
        "Do you really %s? there is occupied counties to liberate!",
        "Why do you need %s when you could vote Sinn Fein?",
        "There would be no need in a communist society friend.",
    },
    `i am (.*)`: {
        "Did you think I'd care that you are %s?",
        "Forget who you are, we're all Irish on this beautiful island of ours",
        "Who you are in a republic made with the ideals of the 1916 rising matters little friend?",
    },
    `i want (.*)`: {
        "Want, want, want. In a republic founded on the ideals of the 1916 rising there would be no want.",
        "Why do you want %s, when we could seize the means of production?",
        "If you got %s, you would want more, capatalism at it's finest my friend.",
        "you want %s, but we want a United Ireland. What's more important friend?",
    },
    `i can't (.*)`: {
        "I never want to hear you say I can't, this republic was founded on I can.",
        "You can't %s? Sounds like you didn't try.",
        "Try harder to %s, I believe in you friend.",
    },

    // About bot
    `why don't you (.*)?`: {
        "Who says I don't %s friend?",
        "Maybe some day I'll %s.",
        "I'll look more into %s.",
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
    "am": "are",
    "was": "were",
    "I": "you",
    "I'd": "you would",
    "I've": "you have",
    "I'll": "you will",
    "my": "your",
    "are": "am",
    "you've": "I have",
    "you'll": "I will",
    "your": "my",
    "yours": "mine",
    "you": "me",
    "me": "you",
}

// Using this to 'Dumb down' responses for bot
// N.B Not yet integrated
var Synonymizer = map[string][]string{
    // Noraml
	"want":       []string{"need", "require", "wish", "longing", "demand"},
    
    // Others
    "ira":      []string{"ra"},
    "uk":      []string{"england","kingdom","british"},
}