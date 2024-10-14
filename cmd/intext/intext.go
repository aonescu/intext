package main

func main() {
    engine := IntextEngine{
        agent: Agent{},
        description: UniversalDescription{
            intent: UserIntent{intent: "example"},
            description: "example description",
        },
    }
    engine.Start()
    engine.Process()
}