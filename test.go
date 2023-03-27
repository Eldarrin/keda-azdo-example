func stripAgentVFromArray(array []string) []string {
	var result []string
	for _, item := range array {
		if !strings.HasPrefix(item, "Agent.Version") {
			result = append(result, item)
		}
	}
	return result
}

func getCanAgentDemandFulfilJob2(jr JobRequest, metadata *azurePipelinesMetadata) bool {
	countDemands := 0
	demandsInJob := stripAgentVFromArray(jr.Demands)
	demandsInScaler := stripAgentVFromArray(strings.Split(metadata.demands, ","))
	for _, demandInJob := range demandsInJob {
		for _, demandInScaler := range demandsInScaler {
			if demandInJob == demandInScaler {
				countDemands++
			}
		}
	}
	if metadata.requireAllDemands {
		return countDemands == len(demandsInJob) && countDemands == len(demandsInScaler)
	} else {
		return countDemands == len(demandsInJob)
	}
}
