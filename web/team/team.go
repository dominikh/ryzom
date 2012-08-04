package main

import (
	"fmt"
	"github.com/dominikh/ryzom/mobs"
	"github.com/dominikh/ryzom/team"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type ProgressiveUnit struct {
	team.TrainingUnit
	LevelProgression string
	SizeProgression  string
}

type Plan struct {
	Mob   mobs.Mob
	Units []ProgressiveUnit
}

type Query struct {
	TeamSize     int
	HighestSkill int
	MinimumXP    int
	Mobs         string
	Plans        []Plan
}

func formatChartData(data []int) string {
	dataStrings := make([]string, 0, len(data))
	for _, value := range data {
		dataStrings = append(dataStrings, fmt.Sprintf("%d", value))
	}

	return strings.Join(dataStrings, ",")
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("team.html")

	teamSize, err1 := strconv.Atoi(r.FormValue("team_size"))
	highestSkill, err2 := strconv.Atoi(r.FormValue("highest_skill"))
	minimumXP, err3 := strconv.Atoi(r.FormValue("minimum_xp"))
	mobList := strings.ToLower(r.FormValue("mobs"))

	query := Query{teamSize, highestSkill, minimumXP, mobList, make([]Plan, 0)}

	log.Printf("New request: team size = '%s', highest skill = '%s', minimum XP = '%s', mobs = '%s'", r.FormValue("team_size"), r.FormValue("highest_skill"), r.FormValue("minimum_xp"), r.FormValue("mobs"))

	if err1 == nil && err2 == nil && err3 == nil {
		desiredMobs := strings.Split(mobList, " ")

		ourTeam := team.Team{teamSize, highestSkill}
		for _, mob := range desiredMobs {
			cMob := mobs.StringToMob(mob)
			units := ourTeam.NewTrainingPlan(minimumXP, cMob, 3)
			progressive_units := make([]ProgressiveUnit, 0, len(units))
			for _, unit := range units {
				progressive_units = append(progressive_units, ProgressiveUnit{unit, formatChartData(unit.ProgressHighestSkill(ourTeam.HighestSkill, 10)), formatChartData((unit.ProgressTeamSize(1, 9)))})
			}
			query.Plans = append(query.Plans, Plan{cMob, progressive_units})
		}
	}

	t.Execute(w, query)
}

func main() {
	http.HandleFunc("/", handler)

	port := os.Getenv("GO_HTTP_PORT")
	if len(port) == 0 {
		port = "8080"
	}

	port = ":" + port
	log.Println("Server Listening on", port)
	http.ListenAndServe(":8080", nil)
}
