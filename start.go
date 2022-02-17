package main
//(2, 3, E) LFRFF
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

// The grid that the rover moves within
var Grid_M int
var Grid_N int

// Struct to store a location and direction
type Location struct {
	x int
	y int
	bearing string
}

// Struct to store location and movement data of a rover
type Rover struct {
	pos Location
	moveSet string
	lost bool
}

func main() {
	fmt.Print("Enter Grid Size (width, height): ")
	reader := bufio.NewReader(os.Stdin)
	
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	
	SetGrid(input)
	
	fmt.Print("Enter Rover location and move set ((x, y, Bearing) LFRFF): ")
	// ReadString will block until the delimiter is entered
	input2, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	
	slice := []string{}
	slice = GetRoverData(input2, slice)
	
	var loc Location
	loc = SetLocation(slice)
	
	var rover Rover
	if(CheckRoverPos(loc) != true){
		fmt.Println("Rover started outside of the grid.\nGrid(",Grid_M,", ",Grid_N,")\nRover(",loc.x,", ",loc.y,")")
	} else {
		rover.pos = loc
		rover.moveSet = slice[3]
		rover.lost = false
		rover = MoveRover(rover)
		if (rover.lost == true) {
			fmt.Println("(",rover.pos,") LOST")
		} else {
			fmt.Println("(",rover.pos,")")
		}
	}
}

func SetGrid(input string){
	// remove the unwanted characters from the string
	input = strings.Replace(input, ",", "", -1)
	
	slice1 := []string{}
	slice1 = strings.Fields(input)
	
	if (len(slice1) != 2) {
		fmt.Println("Error Incompatible grid entered")
		os.Exit(3)
	}
	
	mStr := slice1[0]
	nStr := slice1[1]
	
	m, err := strconv.Atoi(mStr)
	if err != nil {
		fmt.Println("An error occured while converting string to int, Exiting", err)
		os.Exit(3)
	}
	
	n, err := strconv.Atoi(nStr)
	if err != nil {
		fmt.Println("An error occured while converting string to int, Exiting", err)
		os.Exit(3)
	}
	
	Grid_M = m
	Grid_N = n
}

func GetRoverData(input string, slice []string) []string {
	// remove the unwanted characters from the string
	input = strings.Replace(input, ",", "", -1)
	input = strings.Replace(input, "(", "", -1)
	input = strings.Replace(input, ")", "", -1)
	//fmt.Println(input)
	
	slice = strings.Fields(input)
	return slice
}

func SetLocation(slice []string) Location {
	xStr := slice[0]
	yStr := slice[1]
	
	x, err := strconv.Atoi(xStr)
	if err != nil {
		fmt.Println("An error occured while converting string to int, Exiting", err)
		os.Exit(3)
	}
	
	y, err := strconv.Atoi(yStr)
	if err != nil {
		fmt.Println("An error occured while converting string to int, Exiting", err)
		os.Exit(3)
	}
	
	var loc Location
	loc.x = x
	loc.y = y
	loc.bearing = slice[2]
	
	return loc
}

func CheckRoverPos(loc Location) bool {
	if(loc.x <= Grid_M) && (loc.x >= 0) {
		if (loc.y <= Grid_N) && (loc.y >= 0){
			return true
		}
	} 
	return false
}

func MoveRover(rover Rover) Rover {	
	for i, c := range rover.moveSet {
		//fmt.Println(rover.pos)
		if rover.lost == true{
			return rover
		} else if (strings.ToUpper(string(c)) == "F") {
			rover = MoveForwards(rover)
		} else if (strings.ToUpper(string(c)) == "L") {
			rover = TurnLeft(rover)
		} else if (strings.ToUpper(string(c)) == "R") {
			rover = TurnRight(rover)
		} else {
			fmt.Println("Error Undefined move at number:",i," move:",c)
			return rover
		}	
	}
	//fmt.Println(rover.pos)
	return rover
}

func MoveForwards(rover Rover) Rover {
	var current_location Location
	current_location = rover.pos
	bearing := strings.ToUpper(rover.pos.bearing)
	if bearing == "N"{
		current_location.y += 1
	} else if bearing == "E"{
		current_location.x += 1
	} else if bearing == "S"{
		current_location.y -= 1
	} else if bearing == "W"{
		current_location.x -= 1
	} else {
		fmt.Println("Error Undefined Bearing",bearing)
		return rover
	}
	if (CheckRoverPos(current_location) == true) {
			rover.pos = current_location
		} else {
			rover.lost = true
		}
	return rover
}

func TurnRight(rover Rover) Rover {
	bearing := strings.ToUpper(rover.pos.bearing)
	if bearing == "N"{
		rover.pos.bearing = "E"
	} else if bearing == "E"{
		rover.pos.bearing = "S"
	} else if bearing == "S"{
		rover.pos.bearing = "W"
	} else if bearing == "W"{
		rover.pos.bearing = "N"
	} else {
		fmt.Println("Error Undefined Bearing",bearing)
		return rover
	}
	return rover
}

func TurnLeft(rover Rover) Rover {
	bearing := strings.ToUpper(rover.pos.bearing)
	if bearing == "N"{
		rover.pos.bearing = "W"
	} else if bearing == "E"{
		rover.pos.bearing = "N"
	} else if bearing == "S"{
		rover.pos.bearing = "E"
	} else if bearing == "W"{
		rover.pos.bearing = "S"
	} else {
		fmt.Println("Error Undefined Bearing",bearing)
		return rover
	}
	return rover
}