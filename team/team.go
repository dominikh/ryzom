package team

// The formula for calculating how much XP you get in a team of size T,
// killing mobs LVL levels higher than your highest used skill:
//
// D = floor(LVL / 21) * 90;
// X = 1000 + (LVL * 100) + (D * (LVL - 20));
// exp = [2 / (T + 1)] * C * X
//
// T = number of team members
// C = difficulty coefficient (looked up from a table)
// X = basic experience (calculated for positive level differences, looked up for negative ones)
//
//
// The formula for calculating the required base experience to get
// `exp` experience in a team of size T:
//
// X = exp/((2 / (T + 1)) * C)
//
// To get the required level difference, we calculate the base
// experience for every level difference from 0 to 100 and pick the
// first that is >= X.
//
//
// Note that the above formulas imply that only the level difference
// matters, in a constant fashion. Your level itself plays no role.
// However, higher levels need more XP for levelling up.

import (
	"errors"
	"math"
	"github.com/dominikh/ryzom/mobs"
	"github.com/dominikh/ryzom/xp"
)

type Team struct {
	Size         int
	HighestSkill int
}

type TrainingUnit struct {
	Team       Team
	Creature   mobs.Creature
	Levels     int
	StartingXP int
	MinKills   int
	MaxKills   int
	Capped     bool
}

func (unit TrainingUnit) ProgressTeamSize(start, n int) ([]int) {
	result := make([]int, 0, n)
	for i := start; i < start + n; i++ {
		unit.Team.Size = i
		result = append(result, unit.Team.CalculateXP(unit.Creature))
	}

	return result
}

func (unit TrainingUnit) ProgressHighestSkill(start, n int) ([]int) {
	result := make([]int, 0, n)
	for i := start; i < start + n; i++ {
		unit.Team.HighestSkill = i
		result = append(result, unit.Team.CalculateXP(unit.Creature))
	}

	return result
}


func (team Team) CalculateXP(creature mobs.Creature) int {
	exp := team.CalculateXPUncapped(creature)
	if exp > 3000 {
		return 3000
	}

	return exp
}

func (team Team) CalculateXPUncapped(creature mobs.Creature) int {
	c := creature.Coefficient()
	x := float64(xp.BaseXP(creature.Level.Level - team.HighestSkill))
	exp := (2.0 / (float64(team.Size) + 1)) * c * x

	return int(math.Ceil(exp))
}

func (team Team) FindLevelDifference(desiredXP int, coefficient float64) (int, error) {
	requiredX := int(math.Ceil((float64(desiredXP) / ((2 / (float64(team.Size) + 1)) * coefficient))))

	for l := -49; l <= 100; l++ {
		if xp.BaseXP(l) >= requiredX {
			return l, nil
		}
	}

	return 0, errors.New("No suitable level found")
}

func (team Team) NewTrainingPlan(minimumXP int, mob mobs.Mob, length int) (units []TrainingUnit) {
	units = make([]TrainingUnit, 0, length)
	requiredLevelDifference, err := team.FindLevelDifference(minimumXP, mob.Coefficient())

	if err != nil {
		return
	}

	for i := 0; i < length; i++ {
		roundedLevel, err := mobs.RoundLevel(mob, team.HighestSkill+requiredLevelDifference)
		if err != nil {
			break
		}

		creature := mobs.Creature{mob, roundedLevel}

		uncappedXP := team.CalculateXPUncapped(creature)
		cappedXP := team.CalculateXP(creature)

		minKills := int(math.Ceil(float64(xp.XPToNextLevel(team.HighestSkill)) / float64(cappedXP)))
		advancedTeam := team
		advancedTeam.HighestSkill = team.HighestSkill + (roundedLevel.Level - team.HighestSkill - requiredLevelDifference + 1)
		if advancedTeam.HighestSkill > 250 {
			advancedTeam.HighestSkill = 250
		}
		advancedCappedXP := advancedTeam.CalculateXP(creature)
		maxKills := int(math.Ceil(float64(xp.XPToNextLevel(advancedTeam.HighestSkill)) / float64(advancedCappedXP)))

		if minKills > maxKills {
			minKills, maxKills = maxKills, minKills
		}

		units = append(units, TrainingUnit{team, creature,
			roundedLevel.Level - team.HighestSkill - requiredLevelDifference + 1, uncappedXP, minKills, maxKills, uncappedXP > 3000})

		team.HighestSkill = roundedLevel.Level - requiredLevelDifference + 1
		if team.HighestSkill >= 250 {
			break
		}
	}

	return
}
