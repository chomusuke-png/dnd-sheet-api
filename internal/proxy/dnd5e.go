package proxy

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Dnd5eClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewDnd5eClient(baseURL string) *Dnd5eClient {
	return &Dnd5eClient{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

func (client *Dnd5eClient) get(endpoint string, target interface{}) error {
	url := fmt.Sprintf("%s%s", client.baseURL, endpoint)

	response, err := client.httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch %s: %w", url, err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status %d from %s", response.StatusCode, url)
	}

	if err := json.NewDecoder(response.Body).Decode(target); err != nil {
		return fmt.Errorf("failed to decode response from %s: %w", url, err)
	}

	return nil
}

// --- Races ---

type RaceListResponse struct {
	Count   int         `json:"count"`
	Results []ListEntry `json:"results"`
}

type ListEntry struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

func (client *Dnd5eClient) GetRaces() (*RaceListResponse, error) {
	var response RaceListResponse
	if err := client.get("/races", &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (client *Dnd5eClient) GetRace(index string) (map[string]interface{}, error) {
	var response map[string]interface{}
	if err := client.get(fmt.Sprintf("/races/%s", index), &response); err != nil {
		return nil, err
	}
	return response, nil
}

// --- Classes ---

func (client *Dnd5eClient) GetClasses() (*RaceListResponse, error) {
	var response RaceListResponse
	if err := client.get("/classes", &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (client *Dnd5eClient) GetClass(index string) (map[string]interface{}, error) {
	var response map[string]interface{}
	if err := client.get(fmt.Sprintf("/classes/%s", index), &response); err != nil {
		return nil, err
	}
	return response, nil
}

// --- Spells ---

func (client *Dnd5eClient) GetSpells() (*RaceListResponse, error) {
	var response RaceListResponse
	if err := client.get("/spells", &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (client *Dnd5eClient) GetSpell(index string) (map[string]interface{}, error) {
	var response map[string]interface{}
	if err := client.get(fmt.Sprintf("/spells/%s", index), &response); err != nil {
		return nil, err
	}
	return response, nil
}

// --- Equipment ---

func (client *Dnd5eClient) GetEquipment() (*RaceListResponse, error) {
	var response RaceListResponse
	if err := client.get("/equipment", &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (client *Dnd5eClient) GetEquipmentItem(index string) (map[string]interface{}, error) {
	var response map[string]interface{}
	if err := client.get(fmt.Sprintf("/equipment/%s", index), &response); err != nil {
		return nil, err
	}
	return response, nil
}

// --- Structs procesados para el frontend ---

type AbilityBonus struct {
	AbilityIndex string `json:"ability_index"`
	Bonus        int    `json:"bonus"`
}

type RaceData struct {
	Index          string         `json:"index"`
	Name           string         `json:"name"`
	Speed          int            `json:"speed"`
	AbilityBonuses []AbilityBonus `json:"ability_bonuses"`
	Traits         []ListEntry    `json:"traits"`
}

type SavingThrow struct {
	Index string `json:"index"`
}

type ClassData struct {
	Index               string        `json:"index"`
	Name                string        `json:"name"`
	HitDie              int           `json:"hit_die"`
	SavingThrows        []SavingThrow `json:"saving_throws"`
	SpellcastingAbility string        `json:"spellcasting_ability"`
}

func (client *Dnd5eClient) GetRaceData(index string) (*RaceData, error) {
	var raw map[string]interface{}
	if err := client.get(fmt.Sprintf("/races/%s", index), &raw); err != nil {
		return nil, err
	}

	raceData := &RaceData{
		Index: index,
		Name:  raw["name"].(string),
		Speed: int(raw["speed"].(float64)),
	}

	// Parsea ability_bonuses
	if bonuses, ok := raw["ability_bonuses"].([]interface{}); ok {
		for _, bonus := range bonuses {
			bonusMap := bonus.(map[string]interface{})
			abilityScore := bonusMap["ability_score"].(map[string]interface{})
			raceData.AbilityBonuses = append(raceData.AbilityBonuses, AbilityBonus{
				AbilityIndex: abilityScore["index"].(string),
				Bonus:        int(bonusMap["bonus"].(float64)),
			})
		}
	}

	// Parsea traits
	if traits, ok := raw["traits"].([]interface{}); ok {
		for _, trait := range traits {
			traitMap := trait.(map[string]interface{})
			raceData.Traits = append(raceData.Traits, ListEntry{
				Index: traitMap["index"].(string),
				Name:  traitMap["name"].(string),
			})
		}
	}

	return raceData, nil
}

func (client *Dnd5eClient) GetClassData(index string) (*ClassData, error) {
	var raw map[string]interface{}
	if err := client.get(fmt.Sprintf("/classes/%s", index), &raw); err != nil {
		return nil, err
	}

	classData := &ClassData{
		Index:  index,
		Name:   raw["name"].(string),
		HitDie: int(raw["hit_die"].(float64)),
	}

	// Parsea saving_throws
	if savingThrows, ok := raw["saving_throws"].([]interface{}); ok {
		for _, savingThrow := range savingThrows {
			savingThrowMap := savingThrow.(map[string]interface{})
			classData.SavingThrows = append(classData.SavingThrows, SavingThrow{
				Index: savingThrowMap["index"].(string),
			})
		}
	}

	// Parsea spellcasting ability si existe
	if spellcasting, ok := raw["spellcasting"].(map[string]interface{}); ok {
		if spellcastingAbility, ok := spellcasting["spellcasting_ability"].(map[string]interface{}); ok {
			classData.SpellcastingAbility = spellcastingAbility["index"].(string)
		}
	}

	return classData, nil
}
