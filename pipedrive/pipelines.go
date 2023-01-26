package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

// PipelinesService handles pipelines related
// methods of the Pipedrive API.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Pipelines
type PipelinesService service

// Pipeline represents a Pipedrive pipeline.
type Pipeline struct {
	ID              int    `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	URLTitle        string `json:"url_title,omitempty"`
	OrderNr         int    `json:"order_nr,omitempty"`
	Active          bool   `json:"active,omitempty"`
	DealProbability bool   `json:"deal_probability,omitempty"`
	AddTime         string `json:"add_time,omitempty"`
	UpdateTime      string `json:"update_time,omitempty"`
	Selected        bool   `json:"selected,omitempty"`
}

func (p Pipeline) String() string {
	return Stringify(p)
}

// PipelinesResponse represents multiple pipelines response.
type PipelinesResponse struct {
	Success        bool           `json:"success,omitempty"`
	Data           []Pipeline     `json:"data,omitempty"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// PipelineResponse represents single pipeline response.
type PipelineResponse struct {
	Success        bool           `json:"success,omitempty"`
	Data           Pipeline       `json:"data,omitempty"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// PipelineDealsConversionRateResponse represents conversion response.
type PipelineDealsConversionRateResponse struct {
	Success bool `json:"success,omitempty"`
	Data    struct {
		StageConversions []struct {
			FromStageID    int     `json:"from_stage_id,omitempty"`
			ToStageID      int     `json:"to_stage_id,omitempty"`
			ConversionRate float64 `json:"conversion_rate,omitempty"`
		} `json:"stage_conversions,omitempty"`
		WonConversion  float64 `json:"won_conversion,omitempty"`
		LostConversion float64 `json:"lost_conversion,omitempty"`
	} `json:"data,omitempty"`
}

// PipelineDealsMovementResponse represents movement response.
type PipelineDealsMovementResponse struct {
	Success bool `json:"success,omitempty"`
	Data    struct {
		MovementsBetweenStages struct {
			Count int `json:"count,omitempty"`
		} `json:"movements_between_stages,omitempty"`
		NewDeals struct {
			Count   int   `json:"count,omitempty"`
			DealIds []int `json:"deal_ids,omitempty"`
			Values  struct {
				EUR float64 `json:"EUR,omitempty"`
			} `json:"values,omitempty"`
			FormattedValues struct {
				EUR string `json:"EUR,omitempty"`
			} `json:"formatted_values,omitempty"`
		} `json:"new_deals,omitempty"`
		DealsLeftOpen struct {
			Count   int   `json:"count,omitempty"`
			DealIds []int `json:"deal_ids,omitempty"`
			Values  struct {
				EUR float64 `json:"EUR,omitempty"`
			} `json:"values,omitempty"`
			FormattedValues struct {
				EUR string `json:"EUR,omitempty"`
			} `json:"formatted_values,omitempty"`
		} `json:"deals_left_open,omitempty"`
		WonDeals struct {
			Count   int   `json:"count,omitempty"`
			DealIds []int `json:"deal_ids,omitempty"`
			Values  struct {
				EUR int `json:"EUR,omitempty"`
			} `json:"values,omitempty"`
			FormattedValues struct {
				EUR string `json:"EUR,omitempty"`
			} `json:"formatted_values,omitempty"`
		} `json:"won_deals,omitempty"`
		LostDeals struct {
			Count   int   `json:"count,omitempty"`
			DealIds []int `json:"deal_ids,omitempty"`
			Values  struct {
				EUR int `json:"EUR,omitempty"`
			} `json:"values,omitempty"`
			FormattedValues struct {
				EUR string `json:"EUR,omitempty"`
			} `json:"formatted_values,omitempty"`
		} `json:"lost_deals,omitempty"`
		AverageAgeInDays struct {
			AcrossAllStages float64 `json:"across_all_stages,omitempty"`
			ByStages        []struct {
				StageID int `json:"stage_id,omitempty"`
				Value   int `json:"value,omitempty"`
			} `json:"by_stages,omitempty"`
		} `json:"average_age_in_days,omitempty"`
	} `json:"data,omitempty"`
}

// List returns data about all pipelines.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Pipelines/get_pipelines
func (s *PipelinesService) List(ctx context.Context) (*PipelinesResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/pipelines", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *PipelinesResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// GetByID returns data about a specific pipeline.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Pipelines/get_pipelines_id
func (s *PipelinesService) GetByID(ctx context.Context, id int) (*PipelineResponse, *Response, error) {
	uri := fmt.Sprintf("/pipelines/%v", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *PipelineResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// GetDeals returns deal in a specific pipeline.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Pipelines/get_pipelines_id_deals
func (s *PipelinesService) GetDeals(ctx context.Context, id int) (*PipelinesResponse, *Response, error) {
	uri := fmt.Sprintf("/pipelines/%v/deals", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *PipelinesResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// GetDealsConversionRate returns deals conversion rate in a specific pipeline.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Pipelines/get_pipelines_id_conversion_statistics
func (s *PipelinesService) GetDealsConversionRate(ctx context.Context, id int, startDate Timestamp, endDate Timestamp) (*PipelineDealsConversionRateResponse, *Response, error) {
	uri := fmt.Sprintf("/pipelines/%v/conversion_statistics", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, struct {
		StartDate string `url:"start_date,omitempty"`
		EndDate   string `url:"end_date,omitempty"`
	}{
		startDate.Format(),
		endDate.Format(),
	}, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *PipelineDealsConversionRateResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// GetDealsMovement returns deals movement in a specific pipeline.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Pipelines/get_pipelines_id_movement_statistics
func (s *PipelinesService) GetDealsMovement(ctx context.Context, id int, startDate Timestamp, endDate Timestamp) (*PipelineDealsMovementResponse, *Response, error) {
	uri := fmt.Sprintf("/pipelines/%v/movement_statistics", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, struct {
		StartDate string `url:"start_date,omitempty"`
		EndDate   string `url:"end_date,omitempty"`
	}{
		startDate.Format(),
		endDate.Format(),
	}, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *PipelineDealsMovementResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// PipelineCreateOptions specifices the optional parameters to the
// PipelineCreateOptions.Create method.
type PipelineCreateOptions struct {
	Name            string          `url:"name,omitempty"`
	DealProbability DealProbability `url:"deal_probability,omitempty"`
	OrderNr         int             `url:"order_nr,omitempty"`
	Active          ActiveFlag      `url:"active,omitempty"`
}

// Create a new pipeline.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Pipelines/post_pipelines
func (s *PipelinesService) Create(ctx context.Context, opt *PipelineCreateOptions) (*PipelineResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/pipelines", nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *PipelineResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// PipelineUpdateOptions specifices the optional parameters to the
// PipelinesService.Update method.
type PipelineUpdateOptions struct {
	Name            string          `url:"name,omitempty"`
	DealProbability DealProbability `url:"deal_probability,omitempty"`
	OrderNr         int             `url:"order_nr,omitempty"`
	Active          ActiveFlag      `url:"active,omitempty"`
}

// Update a specific pipeline.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Pipelines/put_pipelines_id
func (s *PipelinesService) Update(ctx context.Context, id int, opt *PipelineUpdateOptions) (*PipelineResponse, *Response, error) {
	uri := fmt.Sprintf("/pipelines/%v", id)
	req, err := s.client.NewRequest(http.MethodPost, uri, nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *PipelineResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Delete marks a specific pipeline as deleted.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Pipelines/delete_pipelines_id
func (s *PipelinesService) Delete(ctx context.Context, id int) (*Response, error) {
	uri := fmt.Sprintf("/pipelines/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
