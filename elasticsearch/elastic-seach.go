package elasticsearch

import (
	"context"
	"errors"
	"fmt"
	"github.com/Khoahnhn/go-kafka-elastichsearch/internal/constants"
	"github.com/Khoahnhn/go-kafka-elastichsearch/settings/env"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
	"path/filepath"
)

var esClient *elastic.Client

func InitElasticSearch() {
	var err error
	esURL := env.GetEnv("ELASTICSEARCH_HOST", "http://elasticsearch:9200")

	// Kết nối với ElasticSearch
	esClient, err = elastic.NewClient(elastic.SetURL(esURL), elastic.SetSniff(false))
	if err != nil {
		log.Fatalf("Lỗi kết nối Elasticsearch: %v", err)
	}

	// danh sách index cần tạo
	indices := []string{constants.IndexUser, constants.IndexProduct}

	for _, index := range indices {
		err := ensureIndex(index)
		if err != nil {
			log.Fatalf("Lỗi kiểm tra index %s: %v", index, err)
		}
	}
}

func ensureIndex(indexName string) error {
	if esClient == nil {
		return fmt.Errorf("elasticsearch client chưa được khởi tạo")
	}

	exists, err := esClient.IndexExists(indexName).Do(context.Background())
	if err != nil {
		return err
	}

	if !exists {
		// Đọc file mapping JSON từ thư mục mapping/
		mappingPath, _ := filepath.Abs(fmt.Sprintf("elasticsearch/mapping/%s_mapping.json", indexName))
		mapping, err := os.ReadFile(mappingPath)
		if err != nil {
			return fmt.Errorf("không thể đọc file mapping %s: %v", mappingPath, err)
		}

		// Tạo index với mapping
		_, err = esClient.CreateIndex(indexName).BodyString(string(mapping)).Do(context.Background())
		if err != nil {
			return fmt.Errorf("lỗi tạo index %s: %v", indexName, err)
		}

		log.Printf("Đã tạo index %s thành công!", indexName)
	} else {
		log.Printf("Index %s đã tồn tại.", indexName)
	}
	return nil
}

func GetElasticClient() (*elastic.Client, error) {
	if esClient == nil {
		return nil, errors.New("elasticsearch client chưa được khởi tạo")
	}
	return esClient, nil
}
