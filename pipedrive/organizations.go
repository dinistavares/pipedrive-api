package pipedrive

import (
  "context"
  "fmt"
  "net/http"
)

// OrganizationsService handles organization related
// methods of the Pipedrive API.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Organizations
type OrganizationsService service

// Organization represents a Pipedrive organization.
type Organization struct {
  ID        int `json:"id"`
  CompanyID int `json:"company_id"`
  OwnerID   struct {
    ID         int    `json:"id"`
    Name       string `json:"name"`
    Email      string `json:"email"`
    HasPic     bool   `json:"has_pic"`
    PicHash    string `json:"pic_hash"`
    ActiveFlag bool   `json:"active_flag"`
    Value      int    `json:"value"`
  } `json:"owner_id"`
  Name                            string      `json:"name"`
  OpenDealsCount                  int         `json:"open_deals_count"`
  RelatedOpenDealsCount           int         `json:"related_open_deals_count"`
  ClosedDealsCount                int         `json:"closed_deals_count"`
  RelatedClosedDealsCount         int         `json:"related_closed_deals_count"`
  EmailMessagesCount              int         `json:"email_messages_count"`
  PeopleCount                     int         `json:"people_count"`
  ActivitiesCount                 int         `json:"activities_count"`
  DoneActivitiesCount             int         `json:"done_activities_count"`
  UndoneActivitiesCount           int         `json:"undone_activities_count"`
  ReferenceActivitiesCount        int         `json:"reference_activities_count"`
  FilesCount                      int         `json:"files_count"`
  NotesCount                      int         `json:"notes_count"`
  FollowersCount                  int         `json:"followers_count"`
  WonDealsCount                   int         `json:"won_deals_count"`
  RelatedWonDealsCount            int         `json:"related_won_deals_count"`
  LostDealsCount                  int         `json:"lost_deals_count"`
  RelatedLostDealsCount           int         `json:"related_lost_deals_count"`
  ActiveFlag                      bool        `json:"active_flag"`
  CategoryID                      interface{} `json:"category_id"`
  PictureID                       interface{} `json:"picture_id"`
  CountryCode                     interface{} `json:"country_code"`
  FirstChar                       string      `json:"first_char"`
  UpdateTime                      string      `json:"update_time"`
  AddTime                         string      `json:"add_time"`
  VisibleTo                       string      `json:"visible_to"`
  NextActivityDate                string      `json:"next_activity_date"`
  NextActivityTime                interface{} `json:"next_activity_time"`
  NextActivityID                  int         `json:"next_activity_id"`
  LastActivityID                  int         `json:"last_activity_id"`
  LastActivityDate                string      `json:"last_activity_date"`
  TimelineLastActivityTime        interface{} `json:"timeline_last_activity_time"`
  TimelineLastActivityTimeByOwner interface{} `json:"timeline_last_activity_time_by_owner"`
  Address                         interface{} `json:"address"`
  AddressSubpremise               interface{} `json:"address_subpremise"`
  AddressStreetNumber             interface{} `json:"address_street_number"`
  AddressRoute                    interface{} `json:"address_route"`
  AddressSublocality              interface{} `json:"address_sublocality"`
  AddressLocality                 interface{} `json:"address_locality"`
  AddressAdminAreaLevel1          interface{} `json:"address_admin_area_level_1"`
  AddressAdminAreaLevel2          interface{} `json:"address_admin_area_level_2"`
  AddressCountry                  interface{} `json:"address_country"`
  AddressPostalCode               interface{} `json:"address_postal_code"`
  AddressFormattedAddress         interface{} `json:"address_formatted_address"`
  OwnerName                       string      `json:"owner_name"`
  CcEmail                         string      `json:"cc_email"`
}

func (o Organization) String() string {
  return Stringify(o)
}

// OrganizationsResponse represents multiple organizations response.
type OrganizationsResponse struct {
  Success        bool           `json:"success"`
  Data           []Organization `json:"data"`
  AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// OrganizationResponse represents single organization response.
type OrganizationResponse struct {
  Success        bool           `json:"success"`
  Data           Organization   `json:"data"`
  AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// OrganizationFindOptions specifices the optional parameters to the
// OrganizationsService.Create method.
type OrganizationFindOptions struct {
  Term      	  string    `url:"term"`
}

// OrganizationCreateOptions specifices the optional parameters to the
// OrganizationsService.Create method.
type OrganizationCreateOptions struct {
  Name      string    `json:"name"`
  OwnerID   uint      `json:"owner_id"`
  VisibleTo VisibleTo `json:"visible_to"`
  AddTime   Timestamp `json:"add_time"`
}

// Find all organizations.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Organizations/get_organizations_find
func (s *OrganizationsService) Find(ctx context.Context, opt *OrganizationFindOptions) (*OrganizationsResponse, *Response, error) {
  req, err := s.client.NewRequest(http.MethodGet, "/organizations/find", struct {
    Term      		string    `url:"term"`
  }{
    opt.Term,
  }, nil)

  if err != nil {
    return nil, nil, err
  }

  var record *OrganizationsResponse

  resp, err := s.client.Do(ctx, req, &record)

  if err != nil {
    return nil, resp, err
  }

  return record, resp, nil
}

// List all organizations.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Organizations/get_organizations
func (s *OrganizationsService) List(ctx context.Context) (*OrganizationsResponse, *Response, error) {
  req, err := s.client.NewRequest(http.MethodGet, "/organizations", nil, nil)

  if err != nil {
    return nil, nil, err
  }

  var record *OrganizationsResponse

  resp, err := s.client.Do(ctx, req, &record)

  if err != nil {
    return nil, resp, err
  }

  return record, resp, nil
}

// Create a new organization.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Organizations/post_organizations
func (s *OrganizationsService) Create(ctx context.Context, opt *OrganizationCreateOptions) (*OrganizationResponse, *Response, error) {
  req, err := s.client.NewRequest(http.MethodPost, "/organizations", nil, struct {
    Name      string    `json:"name"`
    OwnerID   uint      `json:"owner_id"`
    VisibleTo VisibleTo `json:"visible_to"`
    AddTime   string    `json:"add_time"`
  }{
    opt.Name,
    opt.OwnerID,
    opt.VisibleTo,
    opt.AddTime.FormatFull(),
  })

  if err != nil {
    return nil, nil, err
  }

  var record *OrganizationResponse

  resp, err := s.client.Do(ctx, req, &record)

  if err != nil {
    return nil, resp, err
  }

  return record, resp, nil
}


// Merge organizations.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/put_persons_id_merge
func (s *OrganizationsService) Merge(ctx context.Context, id int, mergeWithID int) (*OrganizationResponse, *Response, error) {
  uri := fmt.Sprintf("/organizations/%v/merge", id)
  req, err := s.client.NewRequest(http.MethodPut, uri, nil, struct {
    MergeWithID int `url:"merge_with_id"`
  }{
    mergeWithID,
  })

  if err != nil {
    return nil, nil, err
  }

  var record *OrganizationResponse

  resp, err := s.client.Do(ctx, req, &record)

  if err != nil {
    return nil, resp, err
  }

  return record, resp, nil
}

// DeleteFollower deletes a follower from an organization.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Organizations/delete_organizations_id_followers_follower_id
func (s *OrganizationsService) DeleteFollower(ctx context.Context, id int, followerID int) (*Response, error) {
  uri := fmt.Sprintf("/organizations/%v/followers/%v", id, followerID)
  req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

  if err != nil {
    return nil, err
  }

  return s.client.Do(ctx, req, nil)
}

// Delete marks an organization as deleted.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Organizations/delete_organizations_id
func (s *OrganizationsService) Delete(ctx context.Context, id int) (*Response, error) {
  uri := fmt.Sprintf("/organizations/%v", id)
  req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

  if err != nil {
    return nil, err
  }

  return s.client.Do(ctx, req, nil)
}

// DeleteMultiple deletes multiple organizations in bulk.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Organizations/delete_organizations
func (s *OrganizationsService) DeleteMultiple(ctx context.Context, ids []int) (*Response, error) {
  req, err := s.client.NewRequest(http.MethodDelete, "/organizations", &DeleteMultipleOptions{
    Ids: arrayToString(ids, ","),
  }, nil)

  if err != nil {
    return nil, err
  }

  return s.client.Do(ctx, req, nil)
}
