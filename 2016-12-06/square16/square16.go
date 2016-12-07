package square16

import "fmt"

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
		if IsSolution2(s) {
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

func IsSolution2(s Square4x4) bool {
	if perimeter2(s, White) == 16 && perimeter2(s, Black) == 16 {
		return true
	}
	return false
}

func perimeter2(s Square4x4, colour Field) int {
	ret := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if s[i][j] == colour {
				ret += edge(s, i, j, i-1, j) + edge(s, i, j, i, j-1) + edge(s, i, j, i+1, j) + edge(s, i, j, i, j+1)
			}
		}
	}
	return ret
}

func edge(s Square4x4, fieldX, fieldY, neighbourX, neighbourY int) int {
	if neighbourX < 0 || neighbourX > 3 || neighbourY < 0 || neighbourY > 3 {
		return 1
	}
	if s[fieldX][fieldY] != s[neighbourX][neighbourY] {
		return 1
	}
	return 0
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
