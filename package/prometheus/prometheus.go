package prometheus

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/client_golang/prometheus"
)

var topics = []string{
	"article_view_count",
}

type PrometheusClient struct {
	Client v1.API
	Topics map[string]*prometheus.CounterVec
}

func NewPrometheusClient() *PrometheusClient {
	var clientTopics = make(map[string]*prometheus.CounterVec)
	connection, err := api.NewClient(api.Config{
		Address: "http://localhost:9090",
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}
	client := v1.NewAPI(connection)
	for _, topic := range topics {
		counter := prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: topic,
		}, []string{"article_id"})
		prometheus.Register(counter)
		clientTopics[topic] = counter
	}
	return &PrometheusClient{
		Client: client,
		Topics: clientTopics,
	}
}

func (p *PrometheusClient) Increment(topic string, articleID string) {
	p.Topics[topic].WithLabelValues(articleID).Inc()
}

func (p *PrometheusClient) GetStats(articleID string) (int, error) {
	result, _, err := p.Client.Query(context.Background(), articleID, time.Now())
	if err != nil {
		return 0, err
	}
	value, err := strconv.Atoi(result.String())
	if err != nil {
		return 0, err
	}
	return value, nil
}
