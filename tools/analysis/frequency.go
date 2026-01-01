package analysis

func PrintCharacterFrequencies(input string) {
	frequencyMap  := map[string]int{}

	for i := 0; i < len(input); i++ {
		currentCharacter := string(input[i])

		if _, containsKey := frequencyMap[currentCharacter]; containsKey {
			frequencyMap[currentCharacter] += 1
		} else {
			frequencyMap[currentCharacter] = 1
		}
	}
}