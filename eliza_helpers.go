package eliza

import (
    "math/rand"
)

func ElizaIntro() string {
    return randChoice(Introductions)
}

func randChoice(list []string) string {
    randIndex := rand.Intn(len(list))
    return list[randIndex]
}