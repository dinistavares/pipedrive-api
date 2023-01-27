package pipedrive

import (
  "context"
  "fmt"
  "net/http"
)

// PersonsService handles activities related
// methods of the Pipedrive API.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons
type PersonsService service

// Person represents a Pipedrive person.
type Person struct {
  ID        int `json:"id,omitempty"`
  CompanyID int `json:"company_id,omitempty"`
  OwnerID   struct {
    ID         int    `json:"id,omitempty"`
    Name       string `json:"name,omitempty"`
    Email      string `json:"email,omitempty"`
    HasPic     bool   `json:"has_pic,omitempty"`
    PicHash    string `json:"pic_hash,omitempty"`
    ActiveFlag bool   `json:"active_flag,omitempty"`
    Value      int    `json:"value,omitempty"`
  } `json:"owner_id,omitempty"`
  OrgID                       interface{} `json:"org_id,omitempty"`
  Name                        string      `json:"name,omitempty"`
  FirstName                   string      `json:"first_name,omitempty"`
  LastName                    string      `json:"last_name,omitempty"`
  OpenDealsCount              int         `json:"open_deals_count,omitempty"`
  RelatedOpenDealsCount       int         `json:"related_open_deals_count,omitempty"`
  ClosedDealsCount            int         `json:"closed_deals_count,omitempty"`
  RelatedClosedDealsCount     int         `json:"related_closed_deals_count,omitempty"`
  ParticipantOpenDealsCount   int         `json:"participant_open_deals_count,omitempty"`
  ParticipantClosedDealsCount int         `json:"participant_closed_deals_count,omitempty"`
  EmailMessagesCount          int         `json:"email_messages_count,omitempty"`
  ActivitiesCount             int         `json:"activities_count,omitempty"`
  DoneActivitiesCount         int         `json:"done_activities_count,omitempty"`
  UndoneActivitiesCount       int         `json:"undone_activities_count,omitempty"`
  ReferenceActivitiesCount    int         `json:"reference_activities_count,omitempty"`
  FilesCount                  int         `json:"files_count,omitempty"`
  NotesCount                  int         `json:"notes_count,omitempty"`
  FollowersCount              int         `json:"followers_count,omitempty"`
  WonDealsCount               int         `json:"won_deals_count,omitempty"`
  RelatedWonDealsCount        int         `json:"related_won_deals_count,omitempty"`
  LostDealsCount              int         `json:"lost_deals_count,omitempty"`
  RelatedLostDealsCount       int         `json:"related_lost_deals_count,omitempty"`
  ActiveFlag                  bool        `json:"active_flag,omitempty"`
  Phone                       []struct {
    Value   string `json:"value,omitempty"`
    Primary bool   `json:"primary,omitempty"`
  } `json:"phone,omitempty"`
  Email []struct {
    Value   string `json:"value,omitempty"`
    Primary bool   `json:"primary,omitempty"`
  } `json:"email,omitempty"`
  FirstChar                       string      `json:"first_char,omitempty"`
  UpdateTime                      string      `json:"update_time,omitempty"`
  AddTime                         string      `json:"add_time,omitempty"`
  VisibleTo                       string      `json:"visible_to,omitempty"`
  PictureID                       interface{} `json:"picture_id,omitempty"`
  NextActivityDate                interface{} `json:"next_activity_date,omitempty"`
  NextActivityTime                interface{} `json:"next_activity_time,omitempty"`
  NextActivityID                  interface{} `json:"next_activity_id,omitempty"`
  LastActivityID                  int         `json:"last_activity_id,omitempty"`
  LastActivityDate                string      `json:"last_activity_date,omitempty"`
  TimelineLastActivityTime        interface{} `json:"timeline_last_activity_time,omitempty"`
  TimelineLastActivityTimeByOwner interface{} `json:"timeline_last_activity_time_by_owner,omitempty"`
  LastIncomingMailTime            interface{} `json:"last_incoming_mail_time,omitempty"`
  LastOutgoingMailTime            interface{} `json:"last_outgoing_mail_time,omitempty"`
  OrgName                         string 		`json:"org_name,omitempty"`
  OwnerName                       string      `json:"owner_name,omitempty"`
  CcEmail                         string      `json:"cc_email,omitempty"`
}

func (p Person) String() string {
  return Stringify(p)
}

// PersonsResponse represents multiple persons response.
type PersonsResponse struct {
  Success        bool           `json:"success,omitempty"`
  Data           []Person       `json:"data,omitempty"`
  AdditionalData AdditionalData `json:"additional_data,omitempty"`
}


// PersonsResponse represents multiple persons response.
type PersonsSearchResponse struct {
  Success        bool           `json:"success,omitempty"`
  Data    struct {
    Items   []struct {
      Item   Person `json:"item,omitempty"`
    } `json:"items,omitempty"`
  } `json:"data,omitempty"`
  AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// PersonResponse represents single person response.
type PersonResponse struct {
  Success        bool           `json:"success,omitempty"`
  Data           Person         `json:"data,omitempty"`
  AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// PersonDealsResponse represents deals associated to a person.
type PersonDealsResponse struct {
  Success        bool           `json:"success,omitempty"`
  Data           []Deal         `json:"data,omitempty"`
  AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// PersonActivitesResponse represents deals associated to a person.
type PersonActivitesResponse struct {
  Success        bool           `json:"success,omitempty"`
  Data           []Activity     `json:"data,omitempty"`
  AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// PersonAddFollowerResponse represents add follower response.
type PersonAddFollowerResponse struct {
  Success bool `json:"success,omitempty"`
  Data    struct {
    UserID   int    `json:"user_id,omitempty"`
    ID       int    `json:"id,omitempty"`
    PersonID int    `json:"person_id,omitempty"`
    AddTime  string `json:"add_time,omitempty"`
  } `json:"data,omitempty"`
}

// PersonFindOptions specifices the optional parameters to the
// PersonsService.Create method.
type PersonFindOptions struct {
  Term      	  string    `url:"term,omitempty"`
  SearchByEmail uint      `url:"search_by_email,omitempty"`
}

// PersonFindOptions specifices the optional parameters to the
// PersonsService.Create method.
type PersonSearchOptions struct {
  Term      	  string    `url:"term,omitempty"`
  Fields 		  string    `url:"fields,omitempty"`
}

// Get a specific person.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/put_persons_id
func (s *PersonsService) Get(ctx context.Context, id int) (*PersonResponse, *Response, error) {
  uri := fmt.Sprintf("/persons/%v", id)
  req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

  if err != nil {
    return nil, nil, err
  }

  var record *PersonResponse

  resp, err := s.client.Do(ctx, req, &record)

  if err != nil {
    return nil, resp, err
  }

  return record, resp, nil
}

// List deals associated to a person
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/Persons#getPersonDeals
func (s *PersonsService) ListDeals(ctx context.Context, id int) (*PersonDealsResponse, *Response, error) {
  uri := fmt.Sprintf("/persons/%v/deals", id)
  req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

  if err != nil {
    return nil, nil, err
  }

  var record *PersonDealsResponse

  resp, err := s.client.Do(ctx, req, &record)

  if err != nil {
    return nil, resp, err
  }

  return record, resp, nil
}

// List deals associated to a person
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/Persons#getPersonDeals
func (s *PersonsService) ListActivities(ctx context.Context, id int) (*PersonActivitesResponse, *Response, error) {
  uri := fmt.Sprintf("/persons/%v/activities", id)
  req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

  if err != nil {
    return nil, nil, err
  }

  var record *PersonActivitesResponse

  resp, err := s.client.Do(ctx, req, &record)

  if err != nil {
    return nil, resp, err
  }

  return record, resp, nil
}

// List all persons.
//
// Pipedrive API docs:https://developers.pipedrive.com/docs/api/v1/#!/Persons/get_persons_find
func (s *PersonsService) Find(ctx context.Context, opt *PersonFindOptions) (*PersonsResponse, *Response, error) {
  req, err := s.client.NewRequest(http.MethodGet, "/persons/find", struct {
    Term      		string    `url:"term,omitempty"`
    SearchByEmail   uint      `url:"search_by_email,omitempty"`
  }{
    opt.Term,
    opt.SearchByEmail,
  }, nil)

  if err != nil {
    return nil, nil, err
  }

  var record *PersonsResponse

  resp, err := s.client.Do(ctx, req, &record)

  if err != nil {
    return nil, resp, err
  }

  return record, resp, nil
}

// List all persons.
//
// Pipedrive API docs:https://developers.pipedrive.com/docs/api/v1/#!/Persons/get_persons
func (s *PersonsService) Search(ctx context.Context, opt *PersonSearchOptions) (*PersonsSearchResponse, *Response, error) {
  req, err := s.client.NewRequest(http.MethodGet, "/persons/search", struct {
    Term      string    `url:"term,omitempty"`
    Fields    string    `url:"fields,omitempty"`
  }{
    opt.Term,
    opt.Fields,
  }, nil)

  if err != nil {
    return nil, nil, err
  }

  var record *PersonsSearchResponse

  resp, err := s.client.Do(ctx, req, &record)

  if err != nil {
    return nil, resp, err
  }

  return record, resp, nil
}

// List all persons.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/get_persons
func (s *PersonsService) List(ctx context.Context) (*PersonsResponse, *Response, error) {
  req, err := s.client.NewRequest(http.MethodGet, "/persons", nil, nil)

  if err != nil {
    return nil, nil, err
  }

  var record *PersonsResponse

  resp, err := s.client.Do(ctx, req, &record)

  if err != nil {
    return nil, resp, err
  }

  return record, resp, nil
}

// AddFollower adds a follower to person.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/post_persons_id_followers
func (s *PersonsService) AddFollower(ctx context.Context, id int, userID int) (*PersonAddFollowerResponse, *Response, error) {
  uri := fmt.Sprintf("/persons/%v/followers", id)
  req, err := s.client.NewRequest(http.MethodPost, uri, nil, struct {
    UserID int `json:"user_id,omitempty"`
  }{
    userID,
  })

  if err != nil {
    return nil, nil, err
  }

  var record *PersonAddFollowerResponse

  resp, err := s.client.Do(ctx, req, &record)

  if err != nil {
    return nil, resp, err
  }

  return record, resp, nil
}

// PersonCreateOptions specifices the optional parameters to the
// PersonsService.Create method.
type PersonCreateOptions struct {
  Name      string    `json:"name,omitempty"`
  OwnerID   uint      `json:"owner_id,omitempty"`
  OrgID     uint      `json:"org_id,omitempty"`
  Email     string    `json:"email,omitempty"`
  Phone     string    `json:"phone,omitempty"`
  VisibleTo VisibleTo `json:"visible_to,omitempty"`
  AddTime   Timestamp `json:"add_time,omitempty"`
}

// Create a new person.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/post_persons
func (s *PersonsService) Create(ctx context.Context, opt *PersonCreateOptions) (*PersonResponse, *Response, error) {
  req, err := s.client.NewRequest(http.MethodPost, "/persons", nil, struct {
    Name      string    `json:"name,omitempty"`
    OwnerID   uint      `json:"owner_id,omitempty"`
    OrgID     uint      `json:"org_id,omitempty"`
    Email     string    `json:"email,omitempty"`
    Phone     string    `json:"phone,omitempty"`
    VisibleTo VisibleTo `json:"visible_to,omitempty"`
    AddTime   string    `json:"add_time,omitempty"`
  }{
    opt.Name,
    opt.OwnerID,
    opt.OrgID,
    opt.Email,
    opt.Phone,
    opt.VisibleTo,
    opt.AddTime.FormatFull(),
  })

  if err != nil {
    return nil, nil, err
  }

  var record *PersonResponse

  resp, err := s.client.Do(ctx, req, &record)

  if err != nil {
    return nil, resp, err
  }

  return record, resp, nil
}

// PersonUpdateOptions specifices the optional parameters to the
// PersonUpdateOptions.Update method.
type PersonUpdateOptions struct {
  Name      string    `json:"name,omitempty,omitempty"`
  OwnerID   uint      `json:"owner_id,omitempty,omitempty"`
  OrgID     uint      `json:"org_id,omitempty,omitempty"`
  Email     string    `json:"email,omitempty,omitempty"`
  Phone     string    `json:"phone,omitempty,omitempty"`
  VisibleTo VisibleTo `json:"visible_to,omitempty,omitempty"`
}

// Update a specific person.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/put_persons_id
func (s *PersonsService) Update(ctx context.Context, id int, opt *PersonUpdateOptions) (*PersonResponse, *Response, error) {
  uri := fmt.Sprintf("/persons/%v", id)
  req, err := s.client.NewRequest(http.MethodPut, uri, nil, opt)

  if err != nil {
    return nil, nil, err
  }

  var record *PersonResponse

  resp, err := s.client.Do(ctx, req, &record)

  if err != nil {
    return nil, resp, err
  }

  return record, resp, nil
}

// Merge selected persons.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/put_persons_id_merge
func (s *PersonsService) Merge(ctx context.Context, id int, mergeWithID int) (*PersonResponse, *Response, error) {
  uri := fmt.Sprintf("/persons/%v/merge", id)
  req, err := s.client.NewRequest(http.MethodPut, uri, nil, struct {
    MergeWithID int `json:"merge_with_id,omitempty"`
  }{
    mergeWithID,
  })

  if err != nil {
    return nil, nil, err
  }

  var record *PersonResponse

  resp, err := s.client.Do(ctx, req, &record)

  if err != nil {
    return nil, resp, err
  }

  return record, resp, nil
}

// DeleteFollower removes follower from person.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/delete_persons_id_followers_follower_id
func (s *PersonsService) DeleteFollower(ctx context.Context, id int, followerID int) (*Response, error) {
  uri := fmt.Sprintf("/persons/%v/followers/%v", id, followerID)
  req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

  if err != nil {
    return nil, err
  }

  return s.client.Do(ctx, req, nil)
}

// Delete marks person as deleted.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/delete_persons_id
func (s *PersonsService) Delete(ctx context.Context, id int) (*Response, error) {
  uri := fmt.Sprintf("/persons/%v", id)
  req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

  if err != nil {
    return nil, err
  }

  return s.client.Do(ctx, req, nil)
}

// DeletePicture deletes person picture.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/delete_persons_id_picture
func (s *PersonsService) DeletePicture(ctx context.Context, id int) (*Response, error) {
  uri := fmt.Sprintf("/persons/%v/picture", id)
  req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

  if err != nil {
    return nil, err
  }

  return s.client.Do(ctx, req, nil)
}

// DeleteMultiple marks multiple persons as deleted.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/delete_persons
func (s *PersonsService) DeleteMultiple(ctx context.Context, ids []int) (*Response, error) {
  req, err := s.client.NewRequest(http.MethodDelete, "/persons", &DeleteMultipleOptions{
    Ids: arrayToString(ids, ","),
  }, nil)

  if err != nil {
    return nil, err
  }

  return s.client.Do(ctx, req, nil)
}
