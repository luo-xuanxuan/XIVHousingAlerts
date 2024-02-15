package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type HomeSize int

const (
	Small HomeSize = iota
	Medium
	Large
)

type PurchaseSystem int

const (
	unk0 PurchaseSystem = iota
	unk1
	unk2
	FreeCompany
	unk4
	Individual
	FirstCome
	Unrestricted
)

type LottoPhase int

const (
	phase0 LottoPhase = iota
	Missing
	Results
	Unavailable
)

type PlotData struct {
	WorldID                  int            `json:"world_id"`
	DistrictID               int            `json:"district_id"`
	Ward                     int            `json:"ward_number"`
	Plot                     int            `json:"plot_number"`
	Size                     HomeSize       `json:"size"`
	Price                    int            `json:"price"`
	LastUpdated              float64        `json:"last_updated_time"`
	FirstSeen                float64        `json:"first_seen_time"`
	EstimatedTimeOpenMinimum float64        `json:"est_time_open_min"`
	EstimatedTimeOpenMaximum float64        `json:"est_time_open_max"`
	PurchaseSystem           PurchaseSystem `json:"purchase_system"`
	LottoEntries             int            `json:"lotto_entries"`
	LottoPhase               LottoPhase     `json:"lotto_phase"`
	LottoPhaseEnd            int            `json:"lotto_phase_until"`
}

type DistrictData struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	NumOpenPlots   int        `json:"num_open_plots"`
	OldestPlotTime float64    `json:"oldest_plot_time"`
	OpenPlots      []PlotData `json:"open_plots"`
}

type WorldData struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	Districts      []DistrictData `json:"districts"`
	NumOpenPlots   int            `json:"num_open_plots"`
	OldestPlotTime float64        `json:"oldest_plot_time"`
}

func GetWorld(id int) (WorldData, error) {
	url := "https://paissadb.zhu.codes/worlds/" + strconv.Itoa(id)

	resp, err := http.Get(url)
	if err != nil {
		return WorldData{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return WorldData{}, err
	}

	var worldData WorldData
	if err := json.Unmarshal(body, &worldData); err != nil {
		return WorldData{}, err
	}

	return worldData, nil
}
