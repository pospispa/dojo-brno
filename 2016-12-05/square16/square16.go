package square16

import (
	"fmt"
	"os"
)

type Point struct {
	X, Y int
}
type Solution struct {
	Black, White []Point
}

func IsSolution(s Solution) bool {
	if perimeter(s.White) == 16 && perimeter(s.Black) == 16 {
		return true
	}
	return false
}

func perimeter(p []Point) int {
	if len(p) == 0 {
		return 0
	}
	var r int
	for i := range p[1:] {
		r += distance(p[i+1], p[i])
	}
	r += distance(p[len(p)-1], p[0])
	return r
}

func distance(p1, p2 Point) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func GetAllSolutions() {
	var duplicate bool
	var valSolRot1, valSolRot2, valSolRot3, invValSol, invValSolRot1, invValSolRot2, invValSolRot3 Square4x4
	allSols := make(map[Square4x4]Empty)
	allCombinations := Generate()
	for i := range allCombinations {
		s := Array16toSquare4x4(i)
		if !TwoObjectsOnly(s) {
			continue
		}
		if IsBlackOrWhiteEncircled(s) {
			continue
		}
		potSol := Square4x4toSolution(s)
		if IsSolution(potSol) {
			duplicate = false
			for validSolution := range allSols {
				invValSol = ColorInversion(validSolution)
				invValSolRot1 = Rotation(invValSol)
				invValSolRot2 = Rotation(Rotation(invValSol))
				invValSolRot3 = Rotation(Rotation(Rotation(invValSol)))
				valSolRot1 = Rotation(validSolution)
				valSolRot2 = Rotation(Rotation(validSolution))
				valSolRot3 = Rotation(Rotation(Rotation(validSolution)))
				if IsSquare4x4Equal(s, invValSol) || IsSquare4x4Equal(s, invValSolRot1) || IsSquare4x4Equal(s, invValSolRot2) || IsSquare4x4Equal(s, invValSolRot3) || IsSquare4x4Equal(s, valSolRot1) || IsSquare4x4Equal(s, valSolRot2) || IsSquare4x4Equal(s, valSolRot3) {
					duplicate = true
				}
				if IsSquare4x4Equal(s, SymmetryX(validSolution)) || IsSquare4x4Equal(s, SymmetryY(validSolution)) {
					duplicate = true
				}
				if IsSquare4x4Equal(s, SymmetryX(invValSol)) || IsSquare4x4Equal(s, SymmetryY(invValSol)) {
					duplicate = true
				}
				if IsSquare4x4Equal(s, SymmetryX(invValSolRot1)) || IsSquare4x4Equal(s, SymmetryY(invValSolRot1)) {
					duplicate = true
				}
				if IsSquare4x4Equal(s, SymmetryX(invValSolRot2)) || IsSquare4x4Equal(s, SymmetryY(invValSolRot2)) {
					duplicate = true
				}
				if IsSquare4x4Equal(s, SymmetryX(invValSolRot3)) || IsSquare4x4Equal(s, SymmetryY(invValSolRot3)) {
					duplicate = true
				}
				if IsSquare4x4Equal(s, SymmetryX(valSolRot1)) || IsSquare4x4Equal(s, SymmetryY(valSolRot1)) {
					duplicate = true
				}
				if IsSquare4x4Equal(s, SymmetryX(valSolRot2)) || IsSquare4x4Equal(s, SymmetryY(valSolRot2)) {
					duplicate = true
				}
				if IsSquare4x4Equal(s, SymmetryX(valSolRot3)) || IsSquare4x4Equal(s, SymmetryY(valSolRot3)) {
					duplicate = true
				}
			}
			if duplicate {
				continue
			}
			allSols[s] = Empty{}
			fmt.Println("")
			fmt.Println("---------")
			for i := 3; i > -1; i-- {
				for j := 0; j < 4; j++ {
					if s[j][i] == Black {
						fmt.Printf("|X")
					} else {
						fmt.Printf("| ")
					}
				}
				fmt.Println("|")
				fmt.Println("---------")
			}
		}
	}
}

func ColorInversion(s Square4x4) Square4x4 {
	var ret Square4x4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if s[i][j] == Black {
				ret[i][j] = White
			} else {
				ret[i][j] = Black
			}
		}
	}
	return ret
}

func Rotation(s Square4x4) Square4x4 {
	var ret Square4x4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			ret[i][j] = s[3-j][i]
		}
	}
	return ret
}

func SymmetryY(s Square4x4) Square4x4 {
	var ret Square4x4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			ret[i][j] = s[i][3-j]
		}
	}
	return ret
}

func SymmetryX(s Square4x4) Square4x4 {
	var ret Square4x4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			ret[i][j] = s[3-i][j]
		}
	}
	return ret
}

func IsSquare4x4Equal(s1, s2 Square4x4) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if s1[i][j] != s2[i][j] {
				return false
			}
		}
	}
	return true
}

func IsBlackOrWhiteEncircled(s Square4x4) bool {
	colour := s[0][0]
	for i := 0; i < 4; i++ {
		if s[i][0] != colour {
			return false
		}
	}
	for i := 0; i < 4; i++ {
		if s[i][3] != colour {
			return false
		}
	}
	for i := 0; i < 4; i++ {
		if s[0][i] != colour {
			return false
		}
	}
	for i := 0; i < 4; i++ {
		if s[3][i] != colour {
			return false
		}
	}
	return true
}

type Empty struct {
}

type Field int

const (
	White Field = 0 + iota
	Black
)

type Array16 [16]Field

func increment(arr Array16) Array16 {
	ret := arr
	for i := 15; i >= 0; i-- {
		if arr[i] == White {
			ret[i] = Black
			break
		} else {
			ret[i] = White
		}
	}
	return ret
}

func Generate() map[Array16]Empty {
	ret := make(map[Array16]Empty)
	var item Array16
	for i := 0; i < 16; i++ {
		item[i] = White
	}
	allCombinations := powInt(2, 16)
	for i := 0; i < allCombinations; i++ {
		ret[item] = Empty{}
		item = increment(item)
	}
	return ret
}

func powInt(base, power int) int {
	if power < 0 {
		return -1
	}
	if power == 0 {
		return 1
	}
	ret := 1
	for i := 0; i < power; i++ {
		ret *= base
	}
	return ret
}

type Square4x4 [4][4]Field

func Array16toSquare4x4(arr Array16) Square4x4 {
	var ret Square4x4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			ret[i][j] = arr[4*i+j]
		}
	}
	return ret
}

type Coordinate struct {
	x, y int
}

func getAllConnected(colour Field, coord Coordinate, s Square4x4, found map[Coordinate]Empty) map[Coordinate]Empty {
	if _, ok := found[coord]; ok {
		return found
	}
	if coord.x < 0 || coord.x > 3 || coord.y < 0 || coord.y > 3 {
		return found
	}
	ret := found
	if s[coord.x][coord.y] == colour {
		ret[coord] = Empty{}
		newCoord := coord
		newCoord.x = coord.x - 1
		ret = getAllConnected(colour, newCoord, s, ret)
		newCoord = coord
		newCoord.x = coord.x + 1
		ret = getAllConnected(colour, newCoord, s, ret)
		newCoord = coord
		newCoord.y = coord.y - 1
		ret = getAllConnected(colour, newCoord, s, ret)
		newCoord = coord
		newCoord.y = coord.y + 1
		ret = getAllConnected(colour, newCoord, s, ret)
	}
	return ret
}

func getConnected(x, y int, s Square4x4) map[Coordinate]Empty {
	ret := make(map[Coordinate]Empty)
	if x < 0 || x > 3 || y < 0 || y > 3 {
		return ret
	}
	colour := s[x][y]
	coord := Coordinate{x, y}
	return getAllConnected(colour, coord, s, ret)
}

func TwoObjectsOnly(s Square4x4) bool {
	var coord Coordinate
	found := false
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if s[i][j] == Black {
				found = true
				coord = Coordinate{i, j}
				break
			}
		}
	}
	if !found {
		return false
	}
	connected := getConnected(coord.x, coord.y, s)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if s[i][j] == Black {
				crd := Coordinate{i, j}
				if _, ok := connected[crd]; !ok {
					return false
				}
			}
		}
	}
	found = false
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if s[i][j] == White {
				found = true
				coord = Coordinate{i, j}
				break
			}
		}
	}
	if !found {
		return false
	}
	connected = getConnected(coord.x, coord.y, s)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if s[i][j] == White {
				crd := Coordinate{i, j}
				if _, ok := connected[crd]; !ok {
					return false
				}
			}
		}
	}
	return true
}

func Square4x4toSolution(s Square4x4) Solution {
	var ret, ret2 Solution
	var p, pAdded Point
	var index int
	ret.Black = make([]Point, 25)
	ret.White = make([]Point, 25)
	allPoints := make(map[Point]Empty)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if s[i][j] != Black {
				continue
			}
			p.X = i
			p.Y = j
			pAdded = p
			allPoints[pAdded] = Empty{}
			pAdded = p
			pAdded.Y += 1
			allPoints[pAdded] = Empty{}
			pAdded = p
			pAdded.X += 1
			allPoints[pAdded] = Empty{}
			pAdded = p
			pAdded.X += 1
			pAdded.Y += 1
			allPoints[pAdded] = Empty{}
		}
	}
	outerPoints := removeInternal(allPoints, s, Black)
	firstPoint := getLeftBottomMost(outerPoints)
	index = 0
	ret.Black[index] = firstPoint
	taken := make(map[Point]Empty)
	taken[firstPoint] = Empty{}
	currPoint := getNextPoint(firstPoint, taken, outerPoints, s, Black)
	for firstPoint.X != currPoint.X || firstPoint.Y != currPoint.Y {
		index += 1
		ret.Black[index] = currPoint
		taken[currPoint] = Empty{}
		currPoint = getNextPoint(currPoint, taken, outerPoints, s, Black)
		if index == 1 {
			delete(taken, firstPoint)
		}
	}
	index += 1
	ret.Black[index] = currPoint
	ret2.Black = make([]Point, index+1)
	for i := 0; i < index+1; i++ {
		ret2.Black[i] = ret.Black[i]
	}

	allPoints = make(map[Point]Empty)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if s[i][j] != White {
				continue
			}
			p.X = i
			p.Y = j
			pAdded = p
			allPoints[pAdded] = Empty{}
			pAdded = p
			pAdded.Y += 1
			allPoints[pAdded] = Empty{}
			pAdded = p
			pAdded.X += 1
			allPoints[pAdded] = Empty{}
			pAdded = p
			pAdded.X += 1
			pAdded.Y += 1
			allPoints[pAdded] = Empty{}
		}
	}
	outerPoints = removeInternal(allPoints, s, White)
	firstPoint = getLeftBottomMost(outerPoints)
	index = 0
	ret.White[index] = firstPoint
	taken = make(map[Point]Empty)
	taken[firstPoint] = Empty{}
	currPoint = getNextPoint(firstPoint, taken, outerPoints, s, White)
	for firstPoint.X != currPoint.X || firstPoint.Y != currPoint.Y {
		index += 1
		ret.White[index] = currPoint
		taken[currPoint] = Empty{}
		currPoint = getNextPoint(currPoint, taken, outerPoints, s, White)
		if index == 1 {
			delete(taken, firstPoint)
		}
	}
	index += 1
	ret.White[index] = currPoint
	ret2.White = make([]Point, index+1)
	for i := 0; i < index+1; i++ {
		ret2.White[i] = ret.White[i]
	}

	return ret2
}

func getNextPointXY(p Point, taken, all map[Point]Empty, s Square4x4, colour, opposite Field, X, Y int) Point {
	ret := p
	ret.X += X
	if _, found := taken[ret]; !found {
		if _, belongs := all[ret]; belongs {
			return ret
		}
	}
	ret = p
	ret.Y += Y
	if _, found := taken[ret]; !found {
		if _, belongs := all[ret]; belongs {
			return ret
		}
	}
	fmt.Println("")
	fmt.Println("getNextPointXY: no neighbour found in all points. This should never happen")
	fmt.Println("")
	fmt.Printf("Point: (%v)", p)
	fmt.Println("")
	fmt.Printf("Taken: (%v)", taken)
	fmt.Println("")
	fmt.Printf("All: (%v)", all)
	os.Exit(1)
	return ret
}

func getNextPointY(p Point, taken, all map[Point]Empty, s Square4x4, colour, opposite Field) Point {
	ret := p
	ret.Y += -1
	if _, found := taken[ret]; !found {
		if _, belongs := all[ret]; belongs {
			return ret
		}
	}
	ret = p
	ret.Y += 1
	if _, found := taken[ret]; !found {
		if _, belongs := all[ret]; belongs {
			return ret
		}
	}
	fmt.Println("")
	fmt.Println("getNextPointXY: no neighbour found in all points. This should never happen")
	fmt.Println("")
	fmt.Printf("Point: (%v)", p)
	fmt.Println("")
	fmt.Printf("Taken: (%v)", taken)
	fmt.Println("")
	fmt.Printf("All: (%v)", all)
	os.Exit(1)
	return ret
}

func getNextPointX(p Point, taken, all map[Point]Empty, s Square4x4, colour, opposite Field) Point {
	ret := p
	ret.X += -1
	if _, found := taken[ret]; !found {
		if _, belongs := all[ret]; belongs {
			return ret
		}
	}
	ret = p
	ret.X += 1
	if _, found := taken[ret]; !found {
		if _, belongs := all[ret]; belongs {
			return ret
		}
	}
	fmt.Println("")
	fmt.Println("getNextPointXY: no neighbour found in all points. This should never happen")
	fmt.Println("")
	fmt.Printf("Point: (%v)", p)
	fmt.Println("")
	fmt.Printf("Taken: (%v)", taken)
	fmt.Println("")
	fmt.Printf("All: (%v)", all)
	os.Exit(1)
	return ret
}

func getNextPoint(p Point, taken, all map[Point]Empty, s Square4x4, colour Field) Point {
	var opposite Field
	if colour == Black {
		opposite = White
	} else {
		opposite = Black
	}
	if p.X == 0 && p.Y == 0 {
		return getNextPointXY(p, taken, all, s, colour, opposite, 1, 1)
	}
	if p.X == 0 && p.Y == 4 {
		return getNextPointXY(p, taken, all, s, colour, opposite, 1, -1)
	}
	if p.X == 4 && p.Y == 0 {
		return getNextPointXY(p, taken, all, s, colour, opposite, -1, 1)
	}
	if p.X == 4 && p.Y == 4 {
		return getNextPointXY(p, taken, all, s, colour, opposite, -1, -1)
	}
	if p.X == 0 {
		if s[p.X][p.Y-1] == s[p.X][p.Y] {
			return getNextPointY(p, taken, all, s, colour, opposite)
		}
		if s[p.X][p.Y-1] == colour && s[p.X][p.Y] == opposite {
			return getNextPointXY(p, taken, all, s, colour, opposite, 1, -1)
		}
		return getNextPointXY(p, taken, all, s, colour, opposite, 1, 1)
	}
	if p.X == 4 {
		if s[p.X-1][p.Y-1] == s[p.X-1][p.Y] {
			return getNextPointY(p, taken, all, s, colour, opposite)
		}
		if s[p.X-1][p.Y-1] == colour && s[p.X-1][p.Y] == opposite {
			return getNextPointXY(p, taken, all, s, colour, opposite, -1, -1)
		}
		return getNextPointXY(p, taken, all, s, colour, opposite, -1, 1)
	}
	if p.Y == 0 {
		if s[p.X-1][p.Y] == s[p.X][p.Y] {
			return getNextPointX(p, taken, all, s, colour, opposite)
		}
		if s[p.X-1][p.Y] == colour && s[p.X][p.Y] == opposite {
			return getNextPointXY(p, taken, all, s, colour, opposite, -1, 1)
		}
		return getNextPointXY(p, taken, all, s, colour, opposite, 1, 1)
	}
	if p.Y == 4 {
		if s[p.X-1][p.Y-1] == s[p.X][p.Y-1] {
			return getNextPointX(p, taken, all, s, colour, opposite)
		}
		if s[p.X-1][p.Y-1] == colour && s[p.X][p.Y-1] == opposite {
			return getNextPointXY(p, taken, all, s, colour, opposite, -1, -1)
		}
		return getNextPointXY(p, taken, all, s, colour, opposite, 1, -1)
	}
	if s[p.X-1][p.Y-1] == s[p.X][p.Y-1] && s[p.X-1][p.Y] == s[p.X][p.Y] {
		return getNextPointX(p, taken, all, s, colour, opposite)
	}
	if s[p.X-1][p.Y-1] == s[p.X-1][p.Y] && s[p.X][p.Y-1] == s[p.X][p.Y] {
		return getNextPointY(p, taken, all, s, colour, opposite)
	}
	if s[p.X-1][p.Y-1] == s[p.X-1][p.Y] && s[p.X-1][p.Y-1] == s[p.X][p.Y-1] && s[p.X-1][p.Y-1] != s[p.X][p.Y] {
		return getNextPointXY(p, taken, all, s, colour, opposite, 1, 1)
	}
	if s[p.X-1][p.Y-1] == s[p.X-1][p.Y] && s[p.X-1][p.Y-1] == s[p.X][p.Y] && s[p.X-1][p.Y-1] != s[p.X][p.Y-1] {
		return getNextPointXY(p, taken, all, s, colour, opposite, 1, -1)
	}
	if s[p.X-1][p.Y-1] == s[p.X][p.Y-1] && s[p.X-1][p.Y-1] == s[p.X][p.Y] && s[p.X-1][p.Y-1] != s[p.X-1][p.Y] {
		return getNextPointXY(p, taken, all, s, colour, opposite, -1, 1)
	}
	if s[p.X-1][p.Y] == s[p.X][p.Y-1] && s[p.X-1][p.Y] == s[p.X][p.Y] && s[p.X-1][p.Y] != s[p.X-1][p.Y-1] {
		return getNextPointXY(p, taken, all, s, colour, opposite, -1, -1)
	}

	fmt.Println("")
	fmt.Println("getNextPoint: no neighbour found in all points. This should never happen")
	fmt.Println("")
	fmt.Printf("Point: (%v)", p)
	fmt.Println("")
	fmt.Printf("Taken: (%v)", taken)
	fmt.Println("")
	fmt.Printf("All: (%v)", all)
	os.Exit(1)
	ret := p
	return ret
}

func getLeftBottomMost(ps map[Point]Empty) Point {
	var ret Point
	first := true
	for p := range ps {
		if first {
			ret = p
			first = false
		}
		if p.Y < ret.Y {
			ret = p
		} else if p.Y == ret.Y && p.X < ret.X {
			ret = p
		}
	}
	return ret
}

func removeInternal(all map[Point]Empty, s Square4x4, c Field) map[Point]Empty {
	ret := make(map[Point]Empty)
	for p := range all {
		if p.X == 0 || p.X == 4 || p.Y == 0 || p.Y == 4 {
			ret[p] = Empty{}
			continue
		}
		if s[p.X-1][p.Y-1] != c || s[p.X-1][p.Y] != c || s[p.X][p.Y-1] != c || s[p.X][p.Y] != c {
			ret[p] = Empty{}
			continue
		}
	}
	return ret
}
