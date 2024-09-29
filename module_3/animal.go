package main 

import (
    "bufio";
    "fmt";
    "strings";
    "os";
)

type Animal struct {
    food string
    movement string
    sound string
}

var animalMap map[string]Animal = map[string]Animal{
    "cow": Animal{food: "grass", movement: "walk", sound: "moo"},
    "bird": Animal{food: "worms", movement: "fly", sound: "peep"},
    "snake": Animal{food: "mice", movement: "slither", sound: "hss"},
}

func (animal Animal) Eat () {
    fmt.Printf("I like %s\n", animal.food)
    return
}

func (animal Animal) Sound () {
    fmt.Printf("I go %s\n", animal.sound)
    return
}

func (animal Animal) Move () {
    fmt.Printf("I %s around\n", animal.movement)
}

func main () { 
    reader := bufio.NewReader(os.Stdin)

    fmt.Printf("Enter an animal and action from the list:\n")
    fmt.Printf("\tAnimal list:\n")
    for key, _ := range animalMap {
        fmt.Printf("\t- %s\n", key)
    }

    fmt.Printf("\tAction list:\n")
    fmt.Printf("\t- food\n\t- movement\n\t- sound\n")

    fmt.Printf("\nExit by entering 'exit'\n")

    for {
        fmt.Printf(">")
        usrInput, _ := reader.ReadString('\n')

        if strings.ToLower(usrInput) == "exit\n" {
            break
        }
        
        splittedInput := strings.Split(strings.TrimSpace(usrInput), " ")
        
        if len(splittedInput) != 2 {
            continue
        }


        animalObj, ok := animalMap[splittedInput[0]]
        if !ok {
            fmt.Printf("Enter a valid animal\n")
            continue
        }

        switch strings.TrimSpace(splittedInput[1]) {
        case "food":
            animalObj.Eat()
        case "movement":
            animalObj.Move()
        case "sound":
            animalObj.Sound()
        default:
            fmt.Printf("Enter a valid action\n")
        }
    }
    return
}
