package types

type StonkzMetric struct {
	Timestamp int64   `json:"timestamp"`
	Price     float64 `json:"price"`
}
type StonkzMetricsResponse struct {
	StonkzMetrics []StonkzMetric `json:"stonkz_metrics"`
	StonkzName    string         `json:"stonkz_name"`
}
