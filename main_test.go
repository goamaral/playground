package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, postgres

func TestGORM(t *testing.T) {
	id := "fcc03183-58e9-4caa-8e3f-c822ba77e740"
	trigger := Trigger{Base: Base{ID: id}}
	if err := DB.Create(&trigger).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if trigger.ID != id {
		t.Errorf("Failed, id does not match")
	}
	if err := DB.Where(Base{ID: trigger.ID}).FirstOrCreate(&trigger).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if trigger.ID != id {
		t.Errorf("Failed, id does not match")
	}

	trigger = Trigger{Base: Base{ID: id}}
	if err := DB.Where(Base{ID: trigger.ID}).FirstOrCreate(&trigger).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if trigger.ID != id {
		t.Errorf("Failed, id does not match")
	}

	triggers := []Trigger{}
	if err := DB.Model(Trigger{}).Find(&triggers).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if len(triggers) != 1 {
		t.Errorf("Failed, expected 1 trigger, got %d", len(triggers))
	}
	if triggers[0].ID != id {
		t.Errorf("Failed, id does not match")
	}
}
