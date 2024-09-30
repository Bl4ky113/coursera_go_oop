package main

import (
    "errors";
    "fmt";
    "bufio";
    "strings";
    "os";
)

type Animal interface {
    Eat()
    Move()
    Speak()
}

type Cow struct {
    food string
    movement string
    sound string
}

type Snake struct {
    food string
    movement string
    sound string
}

type Bird struct {
    food string
    movement string
    sound string
}

func (c Cow) Eat () { fmt.Printf("I like %s\n", c.food) }
func (c Cow) Speak () { fmt.Printf("I go %s\n", c.sound) }
func (c Cow) Move () { fmt.Printf("I %s around\n", c.movement) }

func (s Snake) Eat () { fmt.Printf("I like %s\n", s.food) }
func (s Snake) Speak () { fmt.Printf("I go %s\n", s.sound) }
func (s Snake) Move () { fmt.Printf("I %s around\n", s.movement) }

func (b Bird) Eat () { fmt.Printf("I like %s\n", b.food) }
func (b Bird) Speak () { fmt.Printf("I go %s\n", b.sound) }
func (b Bird) Move () { fmt.Printf("I %s around\n", b.movement) }

const ( // Usr commands
    newanimal = iota
    query
)

const ( // animal types
    cow = iota
    snake
    bird
)

const ( // animal actions
    food = iota
    sound
    movement
)

var animalBaseMap map[string]Animal = map[string]Animal{
    "cow": Cow{food: "grass", movement: "walk", sound: "moo"},
    "bird": Bird{food: "worms", movement: "fly", sound: "peep"},
    "snake": Snake{food: "mice", movement: "slither", sound: "hss"},
}

var animalListMap map[string]int = make(map[string]int)

func main () {
    usrReader := bufio.NewReader(os.Stdin)

    userMenu(usrReader)
    return
}

func userMenu (reader *bufio.Reader) {
    startMenuText()

    for {
        fmt.Printf(">")
        usrInput, _ := reader.ReadString('\n')

        if strings.ToLower(usrInput) == "exit\n" {
            break
        }
        
        splittedInput := strings.Split(strings.TrimSpace(usrInput), " ")
        commandType, err := handleUsrInput(&splittedInput);
        if err != nil {
            continue
        }

        switch commandType {
        case newanimal:
            animalData := splittedInput[1:len(splittedInput)]
            if 
                err := handleNewAnimal(&animalData)
                err != nil {
                continue
            }

            printAnimalList()

        case query:
            queryData := splittedInput[1:len(splittedInput)]
            if
                err := handleQueryAnimal(&queryData)
                err != nil {
                continue
            }
            
        default:
            continue
        }
    }
    
    return
}

func startMenuText () {
    fmt.Printf("Create an Animal by using:\n\n")
    fmt.Printf(">newanimal <animal_name> <animal_type>\n\n")
    
    fmt.Printf("Available Animal Types:\n")
    for key, _ := range animalBaseMap {
        fmt.Printf("\t- %s\n", key)
    }

    fmt.Printf("\n")
    fmt.Printf("%s\n", strings.Repeat("==", 25))

    fmt.Printf("Query or Search for a Breated Animal using:\n\n")
    fmt.Printf(">query <animal_name> <animal_action>\n\n")

    fmt.Printf("Available Animal Actinos:\n")
    fmt.Printf("\t- food\n\t- movement\n\t- sound\n\n")

    fmt.Printf("Exit by entering 'exit'\n")
    return
}

func handleUsrInput (usrInput *[]string) (int, error) {
    if len((*usrInput)) != 3 {
        fmt.Printf("ERROR, NEEDED 3 ARGUMENTS NOT %d\n", len((*usrInput)))
        return -1, errors.New("incorrect length of arguments in user input")
    }

    switch strings.ToLower((*usrInput)[0]) {
    case "newanimal":
        return newanimal, nil
    case "query":
        return query, nil
    default:
        fmt.Printf("ERROR, INCORRECT COMMAND\n")
        return -1, errors.New("incorrect command given in the user input")
    }
}

func handleNewAnimal (animalData *[]string) error {
    var animalTypeNum int

    if len((*animalData)) != 2 {
        fmt.Printf("ERROR, NEEDED 2 VALUES NOT %d\n", len((*animalData)))
        return errors.New("incorrect length of animal data")
    }

    animalType, ok := animalBaseMap[strings.ToLower((*animalData)[1])]
    if !ok {
        fmt.Printf("ERROR, %s IS NOT AN ANIMAL TYPE\n", (*animalData)[1])
        return errors.New("incorrect animal type")
    }

    switch animalType.(type) {
    case Cow:
        animalTypeNum = cow
    case Bird:
        animalTypeNum = bird
    case Snake:
        animalTypeNum = snake
    default:
        fmt.Printf("ERROR, %s IS NOT AN ANIMAL TYPE\n", (*animalData)[1])
        return errors.New("incorrect animal type")
    }

    return saveNewAnimal((*animalData)[0], animalTypeNum)
}

func saveNewAnimal (animalName string, animalType int) error {
    if 
        animalUnique := checkNewAnimalUnique(strings.ToLower(animalName));
        !animalUnique {
        fmt.Printf("ERROR, %s IS ALREADY AN CREATED ANIMAL\n", animalName)
        return errors.New("non-unique animal name")
    }

    animalListMap[animalName] = animalType

    return nil
}

func checkNewAnimalUnique (animalName string) bool {
    _, ok := animalListMap[animalName]
    if ok {
        return false
    }

    return true
}

func printAnimalList () {
    fmt.Printf("Animals Created:\n")
    for name, aType := range animalListMap {
        fmt.Printf("\t- %s ", name)

        switch aType {
        case cow:
            fmt.Printf("Cow")
        case snake:
            fmt.Printf("Snake")
        case bird:
            fmt.Printf("Bird")
        }

        fmt.Printf("\n")
    }
    return
}

func handleQueryAnimal (queryData *[]string) error {
    var actionType int

    animalType, ok := animalListMap[strings.ToLower((*queryData)[0])]
    if !ok {
        fmt.Printf("ERROR, %s IS NOT SAVED IN THE ANIMAL LIST\n", (*queryData)[0])
        return errors.New("query animal not found in list")
    }

    switch (strings.ToLower((*queryData)[1])) {
    case "eat":
        actionType = food
    case "sound":
        actionType = sound
    case "movement":
        actionType = movement
    default:
        fmt.Printf("ERROR, %s IS NOT AN ANIMAL ACTION\n", (*queryData)[1])
        return errors.New("not defined action")
    }

    return doAnimalAction((*queryData)[0], animalType, actionType)
}

func doAnimalAction (animalName string, animalTypeNum int, actionType int) error {
    var animalType Animal

    switch animalTypeNum {
    case cow:
        animalType = animalBaseMap["cow"]
    case bird:
        animalType = animalBaseMap["bird"]
    case snake:
        animalType = animalBaseMap["snake"]
    default:
        fmt.Printf("ERROR, ANIMAL TYPE NOT DEFINED")
        return errors.New("type passed is not an defined animal type")
    }

    fmt.Printf("I'm %s and ", animalName)

    switch actionType {
    case food:
        animalType.Eat()
    case sound:
        animalType.Speak()
    case movement:
        animalType.Move()
    default:
        fmt.Printf("ERROR, ANIMAL ACTION NOT DEFINED")
        return errors.New("action passed is not an defined animal action")
    }

    return nil
}
