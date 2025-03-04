package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Khoahnhn/go-kafka-elastichsearch/elasticsearch"
	"github.com/Khoahnhn/go-kafka-elastichsearch/internal/constants"
	"github.com/Khoahnhn/go-kafka-elastichsearch/pkg/database"
	"github.com/olivere/elastic/v7"
)

func CreateUserRepository(user User) (User, error) {
	if err := database.DB.Create(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func GetUsersRepository(offset, pageSize int) ([]User, int64, error) {
	var users []User
	var total int64

	if err := database.DB.Model(&User{}).
		Where("deleted_at IS NULL").
		Count(&total).
		Error; err != nil {
		return nil, 0, err
	}

	if err := database.DB.Limit(pageSize).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func GetUserByIDRepository(id string) (User, error) {
	var user User
	if err := database.DB.First(&user, "id = ?", id).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func UpdateUserRepository(user User) (User, error) {
	if err := database.DB.Save(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func DeleteUserRepository(id string) error {
	if err := database.DB.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func SearchUserRepository(query string, filter map[string]string) ([]User, error) {
	client, err := elasticsearch.GetElasticClient()
	if err != nil {
		return nil, err
	}

	esQuery := elastic.NewBoolQuery().
		Should(
			elastic.NewMatchQuery("name", query),
			elastic.NewTermQuery("email", query),
		)

	//// Wildcard Query (tìm kiếm với ký tự đại diện)
	//if wildcard, exists := filter["wildcard"]; exists {
	//	esQuery.Should(
	//		elastic.NewWildcardQuery("name", "*"+wildcard+"*"),
	//		elastic.NewWildcardQuery("email", "*"+wildcard+"*"),
	//	)
	//}
	//
	//// Filter Query (lọc dữ liệu theo điều kiện)
	//if email, exists := filter["email"]; exists {
	//	esQuery.Filter(elastic.NewTermQuery("email", email))
	//}
	//if createdAfter, exists := filter["created_after"]; exists {
	//	esQuery.Filter(elastic.NewRangeQuery("created_at").Gte(createdAfter))
	//}

	searchResult, err := client.Search().
		Index(constants.IndexUser).
		Query(esQuery).
		Do(context.Background())

	if err != nil {
		return nil, err
	}

	// Nếu không có kết quả nào được tìm thấy
	if searchResult.Hits.TotalHits.Value == 0 {
		return nil, fmt.Errorf("không tìm thấy kết quả cho truy vấn: %s", query)
	}

	var users []User
	for _, hit := range searchResult.Hits.Hits {
		var user User
		//err = json.Unmarshal(hit.Source, &user)
		//if err != nil {
		//	fmt.Println("Lỗi Unmarshal:", err)
		//	continue
		//}

		if err := json.Unmarshal(hit.Source, &user); err == nil {
			users = append(users, user)
		}

		users = append(users, user)
	}
	return users, nil
}
