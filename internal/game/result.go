package game

func (g *State) IsApproved(minimum int) bool {
	totalPoints := len(g.Questions) * 10

	if totalPoints == 0 {
		return false
	}

	return g.Points*100 >= totalPoints*minimum
}
