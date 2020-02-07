package Player

type Player struct {
	fname           string
	lname           string
	plateAppearance int
	atBat           int
	single          int
	double          int
	triple          int
	homerun         int
	walks           int
	hitByPitch      int
}

func (P Player) BattingAVG() int {
	return (P.single + P.double + P.triple + P.homerun) / P.atBat
}

func (P Player) Slugging() int {
	return (P.single + (P.double * 2) + (P.triple * 3) + (P.homerun * 4)) / P.atBat
}

func (P Player) OnBase() int {
	return (P.single + P.double + P.triple + P.homerun + P.hitByPitch + P.walks) / P.atBat
}
