package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Determine the total number of repetitions of words in a paragraph.

var paragraph string = "The Sun is the star at the center of the Solar System. It is a nearly perfect sphere of hot plasma, with internal convective motion that generates a magnetic field via a dynamo process. It is, by far, the most important source of energy for life on Earth."

// count_repeated(paragraph) => 11

// "the" is seen 5 times (i.e. 4 repetitions)
// "is" is seen 3 times (i.e. 2 repetitions)
// "of" is seen 3 times (i.e. 2 repetitions)
// "it" is seen 2 times (i.e. 1 repetition)
// "a" is seen 3 times (i.e. 2 repetitions)

// hence number of repetitions is 4+2+2+1+2 = 11

func main () {
 reg, err := regexp.Compile("[^a-z]+")
 if err != nil {
   fmt.Printf("Error %v", err)
   os.Exit(1)
 }
 tokens := strings.Split(paragraph, " ")
 wordsmap := make(map[string]int)
 reps := 0
 for _, word := range tokens {
   //sanitize
   word = reg.ReplaceAllString(strings.ToLower(word), "")
   //add to map
   if occ, ok := wordsmap[word]; !ok {
     wordsmap[word] = 1
   } else {
     wordsmap[word] = occ+1
     reps += 1
   }
 }

 for k, v := range wordsmap {
   fmt.Printf("\"%s\" is seen %d times\n", k,v)
 }
 fmt.Printf("Total repetitions %d of %d words out of total %d words\n", reps, len(wordsmap), len(tokens))
}


