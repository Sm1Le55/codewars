package kata

import "fmt"
import "strings"
import "strconv"
import "errors"

const RESULT_TEMPLATE = "%v:W=%v;D=%v;L=%v;Scored=%v;Conceded=%v;Points=%v"
const NOPLAY_TEMPLATE = "%v:This team didn't play!"
const ERROR_FLOAT_TEMPLATE = "Error(float number):%v"

type Match struct {
	FirstTeam string
	SecondTeam string
	FirstTeamPoints int
	SecondTeamPoints int
}

func NbaCup(resultSheet, toFind string) (result string) {	
	if toFind=="" {
		return
	}

	matches , err := parseMatches(resultSheet)
	if err != nil {
		return fmt.Sprintf(ERROR_FLOAT_TEMPLATE,err)
	}
	
	var cntW,cntD,cntL,scored,conceded,points int

	for _, val := range matches {
		if val.FirstTeam == toFind {
			scored += val.FirstTeamPoints
			conceded += val.SecondTeamPoints

			if val.FirstTeamPoints > val.SecondTeamPoints {
				cntW++
				points += 3
			} else if val.FirstTeamPoints == val.SecondTeamPoints {
				cntD++
				points += 1
			} else {
				cntL++				
			}
		} else if val.SecondTeam == toFind {
			scored += val.SecondTeamPoints
			conceded += val.FirstTeamPoints

			if val.SecondTeamPoints > val.FirstTeamPoints {
				cntW++
				points += 3
			} else if val.FirstTeamPoints == val.SecondTeamPoints {
				cntD++
				points += 1
			} else {
				cntL++				
			}
		}
	}
	if cntW+cntD+cntL == 0 {
		result = fmt.Sprintf(NOPLAY_TEMPLATE,toFind)
	} else {
		result = fmt.Sprintf(RESULT_TEMPLATE,toFind,cntW,cntD,cntL,scored,conceded,points)
	}	
	return 
}

func parseMatches(resultSheet string) ([]Match, error) {
	var result = []Match{}
	matches := strings.Split(resultSheet,",")
	
	for _, buf := range matches {	
		var match Match = Match{
			FirstTeam: "",
			SecondTeam: "",
			FirstTeamPoints: -1,
			SecondTeamPoints: -1,
		}

		for _, val := range strings.Split(buf," ") {
			i, err := strconv.ParseInt(val,10,64)
			if err == nil {
				if match.FirstTeamPoints == -1 {
					match.FirstTeamPoints = int(i)
				} else {
					match.SecondTeamPoints = int(i)
				}
			} else {
				_, err := strconv.ParseFloat(val,64)
				if err == nil {
					return nil, errors.New(buf)
				} else {
					if match.FirstTeamPoints == -1 {
						match.FirstTeam += (" " + val)
					} else {
						match.SecondTeam += (" " + val)
					}
				}
			}
		}
		match.FirstTeam = strings.TrimSpace(match.FirstTeam)
		match.SecondTeam = strings.TrimSpace(match.SecondTeam)
		result = append(result, match)	
	}
	return result,nil
}