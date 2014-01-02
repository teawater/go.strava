package strava

import (
	"encoding/json"
)

type CurrentAthleteService struct {
	client *Client
}

func NewCurrentAthleteService(client *Client) *CurrentAthleteService {
	return &CurrentAthleteService{client}
}

/*********************************************************/

type CurrentAthleteGetCall struct {
	service *CurrentAthleteService
}

func (s *CurrentAthleteService) Get() *CurrentAthleteGetCall {
	return &CurrentAthleteGetCall{
		service: s,
	}
}

func (c *CurrentAthleteGetCall) Do() (*AthleteDetailed, error) {
	data, err := c.service.client.run("GET", "/athlete", nil)
	if err != nil {
		return nil, err
	}

	var athlete AthleteDetailed
	err = json.Unmarshal(data, &athlete)
	if err != nil {
		return nil, err
	}

	athlete.postProcessDetailed()

	return &athlete, nil
}

/*********************************************************/

type CurrentAthleteListActivitiesCall struct {
	service *CurrentAthleteService
	ops     map[string]interface{}
}

func (s *CurrentAthleteService) ListActivities() *CurrentAthleteListActivitiesCall {
	return &CurrentAthleteListActivitiesCall{
		service: s,
		ops:     make(map[string]interface{}),
	}
}

func (c *CurrentAthleteListActivitiesCall) Before(before int) *CurrentAthleteListActivitiesCall {
	c.ops["before"] = before
	return c
}

func (c *CurrentAthleteListActivitiesCall) After(after int) *CurrentAthleteListActivitiesCall {
	c.ops["after"] = after
	return c
}

func (c *CurrentAthleteListActivitiesCall) Page(page int) *CurrentAthleteListActivitiesCall {
	c.ops["page"] = page
	return c
}

func (c *CurrentAthleteListActivitiesCall) PerPage(perPage int) *CurrentAthleteListActivitiesCall {
	c.ops["per_page"] = perPage
	return c
}

func (c *CurrentAthleteListActivitiesCall) Do() ([]*ActivitySummary, error) {
	data, err := c.service.client.run("GET", "/athlete/activities", c.ops)
	if err != nil {
		return nil, err
	}

	activities := make([]*ActivitySummary, 0)
	err = json.Unmarshal(data, &activities)
	if err != nil {
		return nil, err
	}

	for _, a := range activities {
		a.postProcessSummary()
	}

	return activities, nil
}

/*********************************************************/

type CurrentAthleteListFriendsCall struct {
	service *CurrentAthleteService
	ops     map[string]interface{}
}

func (s *CurrentAthleteService) ListFriends() *CurrentAthleteListFriendsCall {
	return &CurrentAthleteListFriendsCall{
		service: s,
		ops:     make(map[string]interface{}),
	}
}

func (c *CurrentAthleteListFriendsCall) Page(page int) *CurrentAthleteListFriendsCall {
	c.ops["page"] = page
	return c
}

func (c *CurrentAthleteListFriendsCall) PerPage(perPage int) *CurrentAthleteListFriendsCall {
	c.ops["per_page"] = perPage
	return c
}

func (c *CurrentAthleteListFriendsCall) Do() ([]*AthleteSummary, error) {
	data, err := c.service.client.run("GET", "/athlete/friends", c.ops)
	if err != nil {
		return nil, err
	}

	friends := make([]*AthleteSummary, 0)
	err = json.Unmarshal(data, &friends)
	if err != nil {
		return nil, err
	}

	for _, a := range friends {
		a.postProcessSummary()
	}

	return friends, nil
}

/*********************************************************/

type CurrentAthleteListFollowersCall struct {
	service *CurrentAthleteService
	ops     map[string]interface{}
}

func (s *CurrentAthleteService) ListFollowers() *CurrentAthleteListFollowersCall {
	return &CurrentAthleteListFollowersCall{
		service: s,
		ops:     make(map[string]interface{}),
	}
}

func (c *CurrentAthleteListFollowersCall) Page(page int) *CurrentAthleteListFollowersCall {
	c.ops["page"] = page
	return c
}

func (c *CurrentAthleteListFollowersCall) PerPage(perPage int) *CurrentAthleteListFollowersCall {
	c.ops["per_page"] = perPage
	return c
}

func (c *CurrentAthleteListFollowersCall) Do() ([]*AthleteSummary, error) {
	data, err := c.service.client.run("GET", "/athlete/followers", c.ops)
	if err != nil {
		return nil, err
	}

	followers := make([]*AthleteSummary, 0)
	err = json.Unmarshal(data, &followers)
	if err != nil {
		return nil, err
	}

	for _, a := range followers {
		a.postProcessSummary()
	}

	return followers, nil
}

/*********************************************************/

type CurrentAthleteListClubsCall struct {
	service *CurrentAthleteService
}

func (s *CurrentAthleteService) ListClubs() *CurrentAthleteListClubsCall {
	return &CurrentAthleteListClubsCall{
		service: s,
	}
}

func (c *CurrentAthleteListClubsCall) Do() ([]*ClubSummary, error) {
	data, err := c.service.client.run("GET", "/athlete/clubs", nil)
	if err != nil {
		return nil, err
	}

	clubs := make([]*ClubSummary, 0)
	err = json.Unmarshal(data, &clubs)
	if err != nil {
		return nil, err
	}

	for _, c := range clubs {
		c.postProcessSummary()
	}

	return clubs, nil
}
