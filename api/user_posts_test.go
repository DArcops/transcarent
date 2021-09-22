package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var (
	expectedOutput = gin.H{
		"id": 1,
		"user_info": map[string]interface{}{
			"name":     "Leanne Graham",
			"email":    "Sincere@april.biz",
			"username": "Bret",
		},
		"posts": []interface{}{
			map[string]interface{}{
				"id":    1,
				"title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
				"body":  "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
			},
			map[string]interface{}{
				"id":    2,
				"title": "qui est esse",
				"body":  "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla",
			},
			map[string]interface{}{
				"id":    3,
				"title": "ea molestias quasi exercitationem repellat qui ipsa sit aut",
				"body":  "et iusto sed quo iure\nvoluptatem occaecati omnis eligendi aut ad\nvoluptatem doloribus vel accusantium quis pariatur\nmolestiae porro eius odio et labore et velit aut",
			},
			map[string]interface{}{
				"id":    4,
				"title": "eum et est occaecati",
				"body":  "ullam et saepe reiciendis voluptatem adipisci\nsit amet autem assumenda provident rerum culpa\nquis hic commodi nesciunt rem tenetur doloremque ipsam iure\nquis sunt voluptatem rerum illo velit",
			},
			map[string]interface{}{
				"id":    5,
				"title": "nesciunt quas odio",
				"body":  "repudiandae veniam quaerat sunt sed\nalias aut fugiat sit autem sed est\nvoluptatem omnis possimus esse voluptatibus quis\nest aut tenetur dolor neque",
			},
			map[string]interface{}{
				"id":    6,
				"title": "dolorem eum magni eos aperiam quia",
				"body":  "ut aspernatur corporis harum nihil quis provident sequi\nmollitia nobis aliquid molestiae\nperspiciatis et ea nemo ab reprehenderit accusantium quas\nvoluptate dolores velit et doloremque molestiae",
			},
			map[string]interface{}{
				"id":    7,
				"title": "magnam facilis autem",
				"body":  "dolore placeat quibusdam ea quo vitae\nmagni quis enim qui quis quo nemo aut saepe\nquidem repellat excepturi ut quia\nsunt ut sequi eos ea sed quas",
			},
			map[string]interface{}{
				"id":    8,
				"title": "dolorem dolore est ipsam",
				"body":  "dignissimos aperiam dolorem qui eum\nfacilis quibusdam animi sint suscipit qui sint possimus cum\nquaerat magni maiores excepturi\nipsam ut commodi dolor voluptatum modi aut vitae",
			},
			map[string]interface{}{
				"id":    9,
				"title": "nesciunt iure omnis dolorem tempora et accusantium",
				"body":  "consectetur animi nesciunt iure dolore\nenim quia ad\nveniam autem ut quam aut nobis\net est aut quod aut provident voluptas autem voluptas",
			},
			map[string]interface{}{
				"id":    10,
				"title": "optio molestias id quia eum",
				"body":  "quo et expedita modi cum officia vel magni\ndoloribus qui repudiandae\nvero nisi sit\nquos veniam quod sed accusamus veritatis error",
			},
		},
	}
)

func TestGetUserPosts(t *testing.T) {
	//Create a test request.
	wr := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(wr)
	ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: "1"})

	GetUserPosts(ctx)

	assert.Equal(t, 200, wr.Code)

	var got gin.H
	err := json.Unmarshal(wr.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	//Compare the response with correct answer. Field by field
	assert.EqualValues(t, expectedOutput["id"], got["id"])

	assert.Equal(t, expectedOutput["user_info"], got["user_info"])

	assert.Equal(t, len(expectedOutput["posts"].([]interface{})), len(got["posts"].([]interface{})))

	expectedPosts := expectedOutput["posts"].([]interface{})
	gotPosts := got["posts"].([]interface{})

	//Compare all attributes from each post.
	for i := range gotPosts {
		formatedExpectedPost := expectedPosts[i].(map[string]interface{})
		formatedGotPost := gotPosts[i].(map[string]interface{})

		assert.EqualValues(t, formatedExpectedPost["id"], formatedGotPost["id"])
		assert.EqualValues(t, formatedExpectedPost["body"], formatedGotPost["body"])
		assert.EqualValues(t, formatedExpectedPost["title"], formatedGotPost["title"])
	}

	//Simulate an error ocurred in external data source, to achhieve that use a mock response.
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	//Let's suppose an internal error from the requested API was triggered.
	httpmock.RegisterResponder("GET", "https://jsonplaceholder.typicode.com/posts",
		httpmock.NewStringResponder(500, ``))

	wrError := httptest.NewRecorder()
	ctxError, _ := gin.CreateTestContext(wrError)
	ctxError.Params = append(ctxError.Params, gin.Param{Key: "id", Value: "1"})

	GetUserPosts(ctxError)

	assert.Equal(t, http.StatusInternalServerError, wrError.Code)
}
