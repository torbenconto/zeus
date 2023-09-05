package zeus

import (
	"strings"

	zeus_lexer "github.com/torbenconto/zeus/lexer"
)

type word struct {
	text   string
	weight int
}

// SummarizeText summarizes a given text.
func SummarizeText(text string, noOfLines int) string {

	lexer := zeus_lexer.NewLexer(text)

	ppPara := preProcess(text)
	sentWeights := []word{}
	wordWeights := make(map[string]int)

	for _, word := range strings.Split(ppPara, " ") {
		word = strings.TrimSpace(word)
		wordWeights[word] = strings.Count(ppPara, word)
	}

	for {
		token := lexer.NextToken()
		if token.Type == zeus_lexer.Unknown {
			break
		}

		s := token.Value
		s = strings.TrimSpace(s)
		if s != "" {
			sentWeights = append(sentWeights, word{text: s, weight: getSentWeight(s, wordWeights)})
		}
	}
	sentWeights = sortSent(sentWeights)

	if len(sentWeights) < noOfLines {
		noOfLines = len(sentWeights)
	}

	summary := ""
	for i := 0; i < noOfLines; i++ {
		summary += sentWeights[i].text + ". "
	}

	return strings.TrimSpace(summary)
}

func preProcess(text string) string {

	lexer := zeus_lexer.NewLexer(text)

	stopWords := []string{"a", "about", "above", "after", "again", "against", "all", "am", "an", "and", "any", "are", "as", "at", "be", "because", "been", "before", "being", "below", "between", "both", "but", "by", "could", "did", "do", "does", "doing", "down", "during", "each", "few", "for", "from", "further", "had", "has", "have", "having", "he", "he'd", "he'll", "he's", "her", "here", "here's", "hers", "herself", "him", "himself", "his", "how", "how's", "i", "i'd", "i'll", "i'm", "i've", "if", "in", "into", "is", "it", "it's", "its", "itself", "let's", "me", "more", "most", "my", "myself", "nor", "of", "on", "once", "only", "or", "other", "ought", "our", "ours", "ourselves", "out", "over", "own", "same", "she", "she'd", "she'll", "she's", "should", "so", "some", "such", "than", "that", "that's", "the", "their", "theirs", "them", "themselves", "then", "there", "there's", "these", "they", "they'd", "they'll", "they're", "they've", "this", "those", "through", "to", "too", "under", "until", "up", "very", "was", "we", "we'd", "we'll", "we're", "we've", "were", "what", "what's", "when", "when's", "where", "where's", "which", "while", "who", "who's", "whom", "why", "why's", "with", "would", "you", "you'd", "you'll", "you're", "you've", "your", "yours", "yourself", "yourselves"}

	var tokens []string

	for {
		token := lexer.NextToken()
		if token.Type == zeus_lexer.Unknown {
			break
		}

		// Consider each token.
		tokenValue := strings.TrimSpace(token.Value)
		if tokenValue != "" {
			tokens = append(tokens, tokenValue)
		}
	}

	// Remove stop words.
	var filteredTokens []string
	for _, token := range tokens {
		// Check if the token is not a stop word.
		isStopWord := false
		for _, stopWord := range stopWords {
			if token == stopWord {
				isStopWord = true
				break
			}
		}
		if !isStopWord {
			filteredTokens = append(filteredTokens, token)
		}
	}

	return strings.Join(filteredTokens, " ")
}

func getSentWeight(sent string, wordWeights map[string]int) int {
	weight := 0
	biggest := 0
	for _, word := range strings.Split(sent, " ") {
		w, ok := wordWeights[word]
		if ok {
			weight += w
			if w > biggest {
				biggest = w
			}
		}
	}
	return weight / biggest
}

func sortSent(sents []word) []word {
	for i := 0; i < len(sents); i++ {
		for j := 0; j < len(sents)-1; j++ {
			if sents[j].weight < sents[j+1].weight {
				temp := sents[j]
				sents[j] = sents[j+1]
				sents[j+1] = temp
			}
		}
	}
	return sents
}
