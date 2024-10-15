package main

func main() {
    userIntent := UserIntent{rawIntent: "Test Intent"}
    description := UniversalDescription{
        intent:      userIntent,
        description: "Sample Description",
    }

    engine := IntextEngine{
        agent:       Agent{},
        description: description,
    }

    engine.Start()
    engine.Process()
}