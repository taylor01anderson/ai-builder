package services

type AgentTask struct {
    Name string
    Prompt string
}

func RunWorkflow(tasks []AgentTask) []string {

    var results []string

    for _, task := range tasks {

        result, _ := GenerateResponse(task.Prompt)

        results = append(results, result)
    }

    return results
}