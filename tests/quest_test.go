package tests

import (
	gameFixtures "elie-api/modules/fixtures/gamification"
	gameController "elie-api/modules/gamification/application/controllers"
	gamificationModels "elie-api/modules/gamification/models"
	"encoding/json"
	"fmt"
	"github.com/go-playground/assert/v2"
	"net/http"
	"testing"
)

func TestIShouldGetQuestsWithCorrectTag(t *testing.T) {

	c, w = CreateGinTestContext()
	tag := gameFixtures.CreateTag()

	c.Request, _ = http.NewRequest("GET", fmt.Sprintf("/api/v1/quests?tag=%s", tag.Name), nil)
	gameController.GetQuests(c)

	assert.Equal(t, w.Code, http.StatusOK)

	var retrievedQuests []gamificationModels.Quest
	err := json.NewDecoder(w.Body).Decode(&retrievedQuests)
	if err != nil {
		t.Errorf("Error while decoding response : %v", err)
	}

	for _, retrievedQuest := range retrievedQuests {
		if retrievedQuest.TagId == tag.Id {
			assert.Equal(t, retrievedQuest.TagId, tag.Id)
			assert.Equal(t, retrievedQuest.Tag.Id, tag.Id)
			assert.Equal(t, retrievedQuest.Tag.Name, tag.Name)
			break
		}
	}
}