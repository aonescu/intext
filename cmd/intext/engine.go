package main

import (
	"os"
	"log"
	"fmt"
)


type UniversalDescription struct {
    intent      UserIntent
    description string
}


type IntextEngine struct {
    agent       Agent
    description UniversalDescription
}

// Start the engine
func (e *IntextEngine) Start() {
    fmt.Println("Engine started")
}

// Action represents an action the agent should take
type Action struct {
    name string
}

// Read the original intext prompt from file
func getOriginalIntext() string {
    data, err := os.ReadFile("/Users/alex/intext/cmd/intext/original-prompt.txt")
    if err != nil {
        log.Fatalf("failed to read file: %v", err)
    }
    return string(data)
}

// Placeholder for prompting the agent
func promptAgent(description string) {
    fmt.Println("Prompting agent with description:", description)
}

// Combine original intext with new information
func combineIntext(original string, newText string) string {
    return original + " + " + newText
}

// Match user descriptions with available actions
func matchDescriptions(description string, descriptionList []UniversalDescription) string {
    for _, d := range descriptionList {
        if d.description == description {
            // Combine original and new intext
            originalIntext := getOriginalIntext()
            combined := combineIntext(originalIntext, d.description)
            return combined
        }
    }
    return "No match found"
}

// Process user intent and decide on an action for the agent
func (e *IntextEngine) Process() {
    // Match description and generate action
    action := matchDescriptions(e.description.description, []UniversalDescription{e.description})
    e.agent.action = action
    
    fmt.Println("Agent action set to:", e.agent.action)
    // Optionally, prompt the agent
    promptAgent(e.agent.action)
}