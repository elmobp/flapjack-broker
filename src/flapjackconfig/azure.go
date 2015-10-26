package flapjackconfig

type AzureAlarm struct {
	Context struct {
		Condition struct {
			MetricName      string `json:"metricName"`
			MetricUnit      string `json:"metricUnit"`
			MetricValue     string `json:"metricValue"`
			Operator        string `json:"operator"`
			Threshold       string `json:"threshold"`
			TimeAggregation string `json:"timeAggregation"`
			WindowSize      string `json:"windowSize"`
		} `json:"condition"`
		ConditionType     string `json:"conditionType"`
		Description       string `json:"description"`
		ID                string `json:"id"`
		Name              string `json:"name"`
		PortalLink        string `json:"portalLink"`
		ResourceGroupName string `json:"resourceGroupName"`
		ResourceID        string `json:"resourceId"`
		ResourceName      string `json:"resourceName"`
		ResourceType      string `json:"resourceType"`
		SubscriptionID    string `json:"subscriptionId"`
		Timestamp         string `json:"timestamp"`
	} `json:"context"`
	Properties struct{} `json:"properties"`
	Status     string   `json:"status"`
}


