package eliza

import (
    "fmt"
    "math/rand"
    "regexp"
    "strings"
)

func ReplyTo(statement string) string {
    statement = preprocess(statement)

    for pattern, responses := range Psychobabble {
        re := regexp.MustCompile(pattern)
        matches := re.FindStringSubmatch(statement)

        if len(matches) > 0 {
            var fragment string
            if len(matches) > 1 {
                fragment = reflect(matches[1])
            }

            response := randChoice(responses)
            if strings.Contains(response, "%s") {
                response = fmt.Sprintf(response, fragment)
            }
            return response
        }
    }

    return randChoice(DefaultResponses)
}

func preprocess(statement string) string {
    statement = strings.TrimRight(statement, "\n.!")
    statement = strings.ToLower(statement)
    return statement
}

func reflect(fragment string) string {
    words := strings.Split(fragment, " ")
    for i, word := range words {
        if Reflected, ok := Reflected[word]; ok {
            words[i] = Reflected
        }
    }
    return strings.Join(words, " ")
}

func randChoice(list []string) string {
    randIndex := rand.Intn(len(list))
    return list[randIndex]
}