package robot

var move []Point = []Point{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

// Advance goes forward
func Advance() {
	Step1Robot.X += move[Step1Robot.Dir].X
	Step1Robot.Y += move[Step1Robot.Dir].Y
}

// Right turns right
func Right() {
	Step1Robot.Dir = (Step1Robot.Dir + 1) % 4
}

// Left turns left
func Left() {
	Step1Robot.Dir = (Step1Robot.Dir + 3) % 4
}

func StartRobot(cmd chan Command, act chan Action) {
	for c := range cmd {
		act <- Action(c)
	}
	close(act)
}

func isInRoom(x, y RU, rect Rect) bool {
	return x >= rect.Min.Easting && x <= rect.Max.Easting &&
		y >= rect.Min.Northing && y <= rect.Max.Northing
}

func goForward(robot *Step2Robot, rect Rect) {
	var newX, newY RU
	newX = robot.Easting + RU(move[robot.Dir].X)
	newY = robot.Northing + RU(move[robot.Dir].Y)
	if isInRoom(newX, newY, rect) {
		robot.Pos = Pos{newX, newY}
	}
}

func turn(robot *Step2Robot, alpha int) {
	robot.Dir = Dir(int(robot.Dir)+alpha+4) % 4
}

func Room(rect Rect, robot Step2Robot, act chan Action, rep chan Step2Robot) {
	for a := range act {
		switch a {
		case 'A':
			goForward(&robot, rect)
		case 'R':
			turn(&robot, 1)
		case 'L':
			turn(&robot, -1)
		}
	}
	rep <- robot
}

func StartRobot3(name string, script string, act chan Action3, log chan string) {
	act <- Action3{name, script}
	act <- Action3{name, "D"}
}

func Room3(rect Rect, robots []Step3Robot, act chan Action3, rep chan []Step3Robot, log chan string) {
	robotExistence := map[string]*Step3Robot{}
	isPlaced := map[Pos]bool{}
	for i := range robots {
		if robots[i].Name == "" {
			log <- "Robot has noname"
		}
		if robotExistence[robots[i].Name] != nil {
			log <- "Duplicate robot name"
		} else {
			robotExistence[robots[i].Name] = &robots[i]
		}
		if isPlaced[robots[i].Pos] {
			log <- "Robot placed at the same place"
		} else {
			isPlaced[robots[i].Pos] = true
		}
		if !isInRoom(robots[i].Easting, robots[i].Northing, rect) {
			log <- "Robot placed outside of the room"
		}
	}
	doneCount := 0
	for a := range act {
		robot := robotExistence[a.Name]
		script := a.Script
		if len(script) == 1 && script[0] == 'D' {
			doneCount++
			if doneCount == len(robots) {
				rep <- robots
				return
			}
		}
		if robot == nil {
			log <- "An action from unknown robot"
		}
		for _, action := range script {
			if robot != nil {
				if action == ' ' || action == 'D' {
				} else if action == 'A' {
					newX := robot.Easting + RU(move[robot.Dir].X)
					newY := robot.Northing + RU(move[robot.Dir].Y)
					if !isInRoom(newX, newY, rect) {
						log <- "A robot attempting to advance into a wall"
					} else {
						if isPlaced[Pos{newX, newY}] {
							log <- "A robot attempting to advance into another robot"
						} else {
							isPlaced[robot.Pos] = false
							isPlaced[Pos{newX, newY}] = true
							goForward(&robot.Step2Robot, rect)
						}
					}
				} else if action == 'R' {
					turn(&robot.Step2Robot, 1)
				} else if action == 'L' {
					turn(&robot.Step2Robot, -1)
				} else {
					log <- "An undefined command in a script"
					break
				}
			}
		}
	}
}
