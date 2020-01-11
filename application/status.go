package application

import (
	"encoding/json"
	"errors"
	"github.com/bjr3ady/weibo-api/constant"
	"github.com/bjr3ady/weibo-api/definitions"
	"strconv"
	"strings"
)

//GetHomeTimeline query weibo user's home timeline
func GetHomeTimeline(accessToken string, count int) ([]definitions.Status, error) {
	response, err := SimpleGet(constant.BaseURI, constant.HomeTimeline, "?access_token=", accessToken, "&count=", strconv.Itoa(count+1))
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var statusesResult definitions.TimelineResult
	if err := json.Unmarshal(response.OriginalBytes, &statusesResult); err != nil {
		return nil, err
	}
	var statuses []definitions.Status
	for i := range statusesResult.Statuses {
		sts := &statusesResult.Statuses[i]
		if sts.RetweetedStatus != nil {
			if err = sts.ShowRetweeted(); err != nil {
				return nil, err
			}
		}
		if sts.User != nil {
			if err = sts.ShowUser(); err != nil {
				return nil, err
			}
		}
		statuses = append(statuses, *sts)
	}
	return statuses, nil
}

//GetUserTimeline query weibo user's timeline
func GetUserTimeline(accessToken string) ([]definitions.Status, error) {
	response, err := SimpleGet(constant.BaseURI, constant.UserTimeline, "?access_token=", accessToken)
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var statusesResult definitions.TimelineResult
	if err := json.Unmarshal(response.OriginalBytes, &statusesResult); err != nil {
		return nil, err
	}
	var statuses []definitions.Status
	for i := range statusesResult.Statuses {
		sts := &statusesResult.Statuses[i]
		if sts.RetweetedStatus != nil {
			if err = sts.ShowRetweeted(); err != nil {
				return nil, err
			}
		}
		if sts.User != nil {
			if err = sts.ShowUser(); err != nil {
				return nil, err
			}
		}
		statuses = append(statuses, *sts)
	}
	return statuses, nil
}

//GetBilateralTimeline query weibo bilateral users' timeline
func GetBilateralTimeline(accessToken string, count int) ([]definitions.Status, error) {
	response, err := SimpleGet(constant.BaseURI, constant.BilateralTimeline, "?access_token=", accessToken, "&count=", strconv.Itoa(count))
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var statusesResult definitions.TimelineResult
	if err := json.Unmarshal(response.OriginalBytes, &statusesResult); err != nil {
		return nil, err
	}
	var statuses []definitions.Status
	for i := range statusesResult.Statuses {
		sts := &statusesResult.Statuses[i]
		if sts.RetweetedStatus != nil {
			if err = sts.ShowRetweeted(); err != nil {
				return nil, err
			}
			if sts.User != nil {
				if err = sts.ShowUser(); err != nil {
					return nil, err
				}
			}
		}
		statuses = append(statuses, *sts)
	}
	return statuses, nil
}

//GetReposts query weibo reposts
func GetReposts(accessToken string, id int) (*definitions.RepostResult, error) {
	response, err := SimpleGet(constant.BaseURI, constant.RepostTimeline, "?access_token=", accessToken, "&id=", strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var repostResult definitions.RepostResult
	if err := json.Unmarshal(response.OriginalBytes, &repostResult); err != nil {
		return nil, err
	}
	return &repostResult, nil
}

//ShowStatus query weibo status by id
func ShowStatus(accessToken string, statusID int) (*definitions.Status, error) {
	response, err := SimpleGet(constant.BaseURI, constant.StatusShow, "?access_token=", accessToken, "&id=", strconv.Itoa(statusID))
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var status definitions.Status
	if err := json.Unmarshal(response.OriginalBytes, &status); err != nil {
		return nil, err
	}
	if status.RetweetedStatus != nil {
		if err := status.ShowRetweeted(); err != nil {
			return nil, err
		}
	}
	return &status, nil
}

//ShowStatusCommentsAndRepostsCount query weibo status' comments and reposts count
func ShowStatusCommentsAndRepostsCount(accessToken string, statusIDs []int) ([]definitions.StautsCountResult, error) {
	var idsStrArr []string
	for _, id := range statusIDs {
		idsStrArr = append(idsStrArr, strconv.Itoa(id))
	}
	idsStr := strings.Join(idsStrArr, ",")
	response, err := SimpleGet(constant.BaseURI, constant.Count, "?access_token=", accessToken, "&ids=", idsStr)
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, err
	}
	var countResults []definitions.StautsCountResult
	if err := json.Unmarshal(response.OriginalBytes, &countResults); err != nil {
		return nil, err
	}
	return countResults, nil
}

//ShareStatus post status to the weibo
func ShareStatus(accessToken, statusStr string) (*definitions.Status, error) {
	response, err := SimplePost(constant.BaseURI+constant.Share, map[string]string{"access_token": accessToken, "status": statusStr})
	if err != nil {
		return nil, err
	}
	if !response.NoError {
		return nil, errors.New(response.Error)
	}
	var status definitions.Status
	if err := json.Unmarshal(response.OriginalBytes, &status); err != nil {
		return nil, err
	}
	return &status, nil
}
