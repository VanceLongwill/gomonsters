package main

import "testing"

func TestAddMonsterToCity(t *testing.T) {
	monsters := NewMonsterCollection()
	unlimitedCity := &City{Name: "a", maxMonsters: -1, Monsters: monsters}
	firstMonster := NewMonster(0)
	if destroyed, err := unlimitedCity.AddMonster(firstMonster); destroyed || err != nil {
		t.Errorf("City with unlimited monster capacity should not be destroyed or full")
	}

	monsters = NewMonsterCollection()
	city := &City{Name: "a", maxMonsters: 2, Monsters: monsters}
	secondMonster := NewMonster(1)

	if destroyed, err := city.AddMonster(firstMonster); destroyed || err != nil {
		t.Errorf("City should not be destroyed when monsters are under capacity")
	}

	if destroyed, err := city.AddMonster(secondMonster); !destroyed || err != nil {
		t.Errorf("City should be destroyed when monsters reach capacity")
	}

	if monsters.Length() != 2 {
		t.Errorf("City should add the correct number of monsters")
	}
}

func TestDestroyCity(t *testing.T) {
	city := NewCity("asd", 2)
	city.Destroy()
	if !city.Destroyed {
		t.Errorf("City.Destroy() should destroy city")
	}

	if err := city.Destroy(); err != ErrCityDestroyed {
		t.Errorf("Should throw error when trying to destroy a destroyed city")
	}
}
