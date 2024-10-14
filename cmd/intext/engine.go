package main

import (
	"os"
	"log"
)

type UniversalDescription struct {
    intent UserIntent
    description string
}

type IntextEngine struct {
    agent Agent
    description UniversalDescription
}

func (e *IntextEngine) Start() {
    // Start the engine
}

type Action struct {
    name string
}

func getOriginalIntext() string {
    data, err := os.ReadFile("/Users/alex/intext/cmd/intext/original-prompt.txt")
    if err != nil {
        log.Fatalf("failed to read file: %v", err)
    }
    return string(data)
}

// Generic function to prompt the agent
func promptAgent() {}

// Original intext looks at the starting one how it proposes to match with the new one.
func combineIntext() string {
    // o := getOriginalIntext()
    return "New intext"
}

func matchDescriptions(description string, descriptionList []UniversalDescription) string {
    // Get parsed user intent in to a string that can be used to match from the available actions that llm agents expose 
    for _, d := range descriptionList {
        if d.description == description {
            desc := combineIntext()
            return desc
        }
    }
    return "No match found"
}

func (e *IntextEngine) Process() {
    // Process the user intent and generate the action that the agent should take
    action := matchDescriptions(e.description.description, []UniversalDescription{e.description})
    e.agent.action = action
}