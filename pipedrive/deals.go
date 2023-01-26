package pipedrive

import (
  "context"
  "fmt"
  "net/http"
)

// DealService handles deals related
// methods of the Pipedrive API.
//
// Pipedrive API dcos: https://developers.pipedrive.com/docs/api/v1/#!/Deals
type DealService service

// Deal represents a Pipedrive deal.
type Deal struct {
  ID                             int           `json:"id,omitempty"`
  CreatorUserID                  CreatorUserID `json:"creator_user_id,omitempty"`
  UserID                         UserID        `json:"user_id,omitempty"`
  PersonID                       PersonID      `json:"person_id,omitempty"`
  OrgID                          OrgID         `json:"org_id,omitempty"`
  StageID                        int           `json:"stage_id,omitempty"`
  Title                          string        `json:"title,omitempty"`
  Value                          int           `json:"value,omitempty"`
  Currency                       string        `json:"currency,omitempty"`
  AddTime                        string        `json:"add_time,omitempty"`
  UpdateTime                     string        `json:"update_time,omitempty"`
  StageChangeTime                string        `json:"stage_change_time,omitempty"`
  Active                         bool          `json:"active,omitempty"`
  Deleted                        bool          `json:"deleted,omitempty"`
  Status                         string        `json:"status,omitempty"`
  Probability                    int           `json:"probability,omitempty"`
  NextActivityDate               interface{}   `json:"next_activity_date,omitempty"`
  NextActivityTime               interface{}   `json:"next_activity_time,omitempty"`
  NextActivityID                 interface{}   `json:"next_activity_id,omitempty"`
  LastActivityID                 int           `json:"last_activity_id,omitempty"`
  LastActivityDate               string        `json:"last_activity_date,omitempty"`
  LostReason                     string        `json:"lost_reason,omitempty"`
  VisibleTo                      string        `json:"visible_to,omitempty"`
  CloseTime                      string        `json:"close_time,omitempty"`
  PipelineID                     int           `json:"pipeline_id,omitempty"`
  WonTime                        interface{}   `json:"won_time,omitempty"`
  FirstWonTime                   interface{}   `json:"first_won_time,omitempty"`
  LostTime                       string        `json:"lost_time,omitempty"`
  ProductsCount                  int           `json:"products_count,omitempty"`
  FilesCount                     int           `json:"files_count,omitempty"`
  NotesCount                     int           `json:"notes_count,omitempty"`
  FollowersCount                 int           `json:"followers_count,omitempty"`
  EmailMessagesCount             int           `json:"email_messages_count,omitempty"`
  ActivitiesCount                int           `json:"activities_count,omitempty"`
  DoneActivitiesCount            int           `json:"done_activities_count,omitempty"`
  UndoneActivitiesCount          int           `json:"undone_activities_count,omitempty"`
  ReferenceActivitiesCount       int           `json:"reference_activities_count,omitempty"`
  ParticipantsCount              int           `json:"participants_count,omitempty"`
  ExpectedCloseDate              string        `json:"expected_close_date,omitempty"`
  LastIncomingMailTime           interface{}   `json:"last_incoming_mail_time,omitempty"`
  LastOutgoingMailTime           interface{}   `json:"last_outgoing_mail_time,omitempty"`
  StageOrderNr                   int           `json:"stage_order_nr,omitempty"`
  PersonName                     string        `json:"person_name,omitempty"`
  OrgName                        string        `json:"org_name,omitempty"`
  NextActivitySubject            interface{}   `json:"next_activity_subject,omitempty"`
  NextActivityType               interface{}   `json:"next_activity_type,omitempty"`
  NextActivityDuration           interface{}   `json:"next_activity_duration,omitempty"`
  NextActivityNote               interface{}   `json:"next_activity_note,omitempty"`
  FormattedValue                 string        `json:"formatted_value,omitempty"`
  RottenTime                     interface{}   `json:"rotten_time,omitempty"`
  WeightedValue                  int           `json:"weighted_value,omitempty"`
  FormattedWeightedValue         string        `json:"formatted_weighted_value,omitempty"`
  OwnerName                      string        `json:"owner_name,omitempty"`
  CcEmail                        string        `json:"cc_email,omitempty"`
  OrgHidden                      bool          `json:"org_hidden,omitempty"`
  PersonHidden                   bool          `json:"person_hidden,omitempty"`
  Eight02Aa45Ecc05F31Fcebe8B706510389F56B7A041 interface{} `json:"802aa45ecc05f31fcebe8b706510389f56b7a041,omitempty"`
}

type  CreatorUserID struct {
  ID         int    `json:"id,omitempty"`
  Name       string `json:"name,omitempty"`
  Email      string `json:"email,omitempty"`
  HasPic     bool   `json:"has_pic,omitempty"`
  PicHash    string `json:"pic_hash,omitempty"`
  ActiveFlag bool   `json:"active_flag,omitempty"`
  Value      int    `json:"value,omitempty"`
}

type 	UserID struct {
  ID         int    `json:"id,omitempty"`
  Name       string `json:"name,omitempty"`
  Email      string `json:"email,omitempty"`
  HasPic     bool   `json:"has_pic,omitempty"`
  PicHash    string `json:"pic_hash,omitempty"`
  ActiveFlag bool   `json:"active_flag,omitempty"`
  Value      int    `json:"value,omitempty"`
}

type PersonID struct {
  Value int     `json:"value,omitempty"`
  Name  string  `json:"name,omitempty"`
  Email []Email `json:"email,omitempty"`
  Phone []Phone `json:"phone,omitempty"`
}

type Email struct {
  Value   string `json:"value,omitempty"`
  Primary bool   `json:"primary,omitempty"`
}

type Phone struct {
  Value   string `json:"value,omitempty"`
  Primary bool   `json:"primary,omitempty"`
}

type OrgID struct {
  Name        string      `json:"name,omitempty"`
  PeopleCount int         `json:"people_count,omitempty"`
  OwnerID     int         `json:"owner_id,omitempty"`
  Address     interface{} `json:"address,omitempty"`
  CcEmail     string      `json:"cc_email,omitempty"`
  Value       int         `json:"value,omitempty"`
  ActiveFlag  bool   `json:"active_flag,omitempty"`
}

func (d Deal) String() string {
  return Stringify(d)
}

// DealsResponse represents multiple deals response.
type DealsResponse struct {
  Success        bool           `json:"success,omitempty,omitempty"`
  Data           []Deal         `json:"data,omitempty,omitempty"`
  AdditionalData AdditionalData `json:"additional_data,omitempty,omitempty"`
}

// DealResponse represents single deal response.
type DealResponse struct {
  Success        bool           `json:"success,omitempty,omitempty"`
  Data           Deal           `json:"data,omitempty,omitempty"`
  AdditionalData AdditionalData `json:"additional_data,omitempty,omitempty"`
}

// ListUpdates about a deal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/get_deals_id_flow
func (s *DealService) ListUpdates(ctx context.Context, id int) (*DealsResponse, *Response, error) {
  uri := fmt.Sprintf("/deals/%v/flow", id)
  req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

  if err != nil {
    return nil, nil, err
  }

  var record *DealsResponse

  resp, err := s.client.Do(ctx, req, &record)

  if err != nil {
    return nil, resp, err
  }

  return record, resp, nil
}

// Find deals by name.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/get_deals_find
func (s *DealService) Find(ctx context.Context, term string) (*DealsResponse, *Response, error) {
  req, err := s.client.NewRequest(http.MethodGet, "/deals/find", &SearchOptions{
    Term: term,
  }, nil)

  if err != nil {
    return nil, nil, err
  }

  var record *DealsResponse

  resp, err := s.client.Do(ctx, req, &record)

  if err != nil {
    return nil, resp, err
  }

  return record, resp, nil
}

// List all deals.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/get_deals
func (s *DealService) List(ctx context.Context) (*DealsResponse, *Response, error) {
  req, err := s.client.NewRequest(http.MethodGet, "/deals", nil, nil)

  if err != nil {
    return nil, nil, err
  }

  var record *DealsResponse

  resp, err := s.client.Do(ctx, req, &record)

  if err != nil {
    return nil, resp, err
  }

  return record, resp, nil
}

// Duplicate a deal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/post_deals_id_duplicate
func (s *DealService) Duplicate(ctx context.Context, id int) (*DealResponse, *Response, error) {
  uri := fmt.Sprintf("/deals/%v/duplicate", id)
  req, err := s.client.NewRequest(http.MethodPost, uri, nil, nil)

  if err != nil {
    return nil, nil, err
  }

  var record *DealResponse

  resp, err := s.client.Do(ctx, req, &record)

  if err != nil {
    return nil, resp, err
  }

  return record, resp, nil
}

// DealsMergeOptions specifices the optional parameters to the
// DealService.Merge method.
type DealsMergeOptions struct {
  MergeWithID uint `url:"merge_with_id,omitempty,omitempty"`
}

// Merge two deals.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/put_deals_id_merge
func (s *DealService) Merge(ctx context.Context, id int, opt *DealsMergeOptions) (*Response, error) {
  uri := fmt.Sprintf("/deals/%v/merge", id)
  req, err := s.client.NewRequest(http.MethodPut, uri, nil, opt)

  if err != nil {
    return nil, err
  }

  return s.client.Do(ctx, req, nil)
}

type DealCreateOptions struct {
  StageID           int       `json:"stage_id,omitempty"`
  Probability       int       `json:"probability,omitempty"`
  PersonID          int       `json:"person_id,omitempty"`
  OrganizationID    int       `json:"org_id,omitempty"`
  PipelineID        int       `json:"pipeline_id,omitempty"`
  UserID            int       `json:"user_id,omitempty"`
  VisibleTo         string    `json:"visible_to,omitempty"`
  Value             string    `json:"value,omitempty"`
  Title             string    `json:"title,omitempty"`
  ExpectedCloseDate string    `json:"expected_close_date,omitempty"`
  Currency          string    `json:"currency,omitempty"`
  Status            string    `json:"status,omitempty"`
  LostReason        string    `json:"lost_reason,omitempty"`
  AddTime           string    `json:"add_time,omitempty"`
}

// Add a deal.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/Deals#addDeal
func (s *DealService) Add(ctx context.Context, opt *DealCreateOptions) (*DealResponse, *Response, error) {
  uri := fmt.Sprintf("/deals")
  req, err := s.client.NewRequest(http.MethodPost, uri, nil, opt)

  if err != nil {
    return nil, nil, err
  }

  var record *DealResponse

  resp, err := s.client.Do(ctx, req, &record)

  if err != nil {
    return nil, resp, err
  }

  return record, resp, nil
}

// DealsUpdateOptions specifices the optional parameters to the
// DealService.Update method.
type DealsUpdateOptions struct {
  Title          string `json:"title,omitempty,omitempty"`
  Value          string `json:"value,omitempty,omitempty"`
  Currency       string `json:"currency,omitempty,omitempty"`
  UserID         uint   `json:"user_id,omitempty,omitempty"`
  PersonID       uint   `json:"person_id,omitempty,omitempty"`
  OrganizationID uint   `json:"org_id,omitempty,omitempty"`
  StageID        uint   `json:"stage_id,omitempty,omitempty"`
  Status         string `json:"status,omitempty,omitempty"`
  LostReason     string `json:"lost_reason,omitempty,omitempty"`
  VisibleTo      uint   `json:"visible_to,omitempty,omitempty"`
}


// Update a deal.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/put_deals_id
func (s *DealService) Update(ctx context.Context, id int, opt *DealsUpdateOptions) (*Response, error) {
  uri := fmt.Sprintf("/deals/%v", id)
  req, err := s.client.NewRequest(http.MethodPut, uri, nil, opt)

  if err != nil {
    return nil, err
  }

  return s.client.Do(ctx, req, nil)
}

// DeleteFollower of a deal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals_id_followers_follower_id
func (s *DealService) DeleteFollower(ctx context.Context, id int, followerID int) (*Response, error) {
  uri := fmt.Sprintf("/deals/%v/followers/%v", id, followerID)
  req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

  if err != nil {
    return nil, err
  }

  return s.client.Do(ctx, req, nil)
}

// DeleteMultiple deletes deals in bulk.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals
func (s *DealService) DeleteMultiple(ctx context.Context, ids []int) (*Response, error) {
  req, err := s.client.NewRequest(http.MethodDelete, "/deals", &DeleteMultipleOptions{
    Ids: arrayToString(ids, ","),
  }, nil)

  if err != nil {
    return nil, err
  }

  return s.client.Do(ctx, req, nil)
}

// DeleteParticipant deletes participant in a deal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals_id_participants_deal_participant_id
func (s *DealService) DeleteParticipant(ctx context.Context, dealID int, participantID int) (*Response, error) {
  uri := fmt.Sprintf("/deals/%v/participants/%v", dealID, participantID)
  req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

  if err != nil {
    return nil, err
  }

  return s.client.Do(ctx, req, nil)
}

// Delete a deal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals_id
func (s *DealService) Delete(ctx context.Context, id int) (*Response, error) {
  uri := fmt.Sprintf("/deals/%v", id)
  req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

  if err != nil {
    return nil, err
  }

  return s.client.Do(ctx, req, nil)
}

// DeleteAttachedProduct deletes attached product.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals_id_products_product_attachment_id
func (s *DealService) DeleteAttachedProduct(ctx context.Context, dealID int, productAttachmentID int) (*Response, error) {
  uri := fmt.Sprintf("/deals/%v/products/%v", dealID, productAttachmentID)
  req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

  if err != nil {
    return nil, err
  }

  return s.client.Do(ctx, req, nil)
}
