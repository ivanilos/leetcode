func predictPartyVictory(senate string) string {
    banned := make([]bool, len(senate))

    toBanRadiant := 0
    toBanDire := 0

    for true {
        bannedInRound := 0
        for i, c := range senate {
            if !banned[i] && c == 'R' {
                if toBanRadiant > 0 {
                    bannedInRound++
                    banned[i] = true
                    toBanRadiant--
                } else {
                    toBanDire++
                }
            } else if !banned[i] && c == 'D' {
                if toBanDire > 0 {
                    bannedInRound++
                    banned[i] = true
                    toBanDire--
                } else {
                    toBanRadiant++
                }

            }
        }

        if bannedInRound == 0 {
            break
        }
    }

    for i := 0; i < len(senate); i++ {
        if !banned[i] && senate[i] == 'R' {
            return "Radiant"
        } else if !banned[i] && senate[i] == 'D' {
            return "Dire"
        }
    }

    panic("Should not reach here")
}
