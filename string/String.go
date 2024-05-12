package main
import(
"strings"
"testing"
)

//ReverseWords takes a string as input and returns the string with each word Reversed
func
ReverseWords(input string) string {
words :=
strings.Fields(input)
reversed :=
make ([]string,len(words))
for i,word :=range words {
reversedWord := ""
for _ ,char := range word {
reversedWord = string (char) + reversedWord
}
reversed[i] =reversedWord
}
return
string.join(reversed," ")
}
//TestReverseWords test the Reversed Function
func
TestReverseWords(t*testing.T){
testCases :=[] struct {
input
string
expectedOutput string
}{
{"Hello, World!","olleH ,!dlroW"},
{"Go Programming","oG gnimmargorP"} ,
{"I love Go !","I evol !oG"},
}

for _, tc := range testCases {
output :=
ReverseWords(tc.input)
if output != tc. expectedOutput {

t.Errorf("ReverseWords(%s) %s, expected %s", tc.input,
output,
tc.expectedOutput)
}
}
}
