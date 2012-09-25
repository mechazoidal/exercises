package main

import (
  "code.google.com/p/go-tour/wc";
  "strings";
)

func WordCount(s string) map[string]int {
  counts := make(map[string]int)
  tokens :=int strings.Fields(s)
  for _, s := range tokens {
    counts[s] += 1
  }
  return counts
}

func main() {
  wc.Test(WordCount)
}
