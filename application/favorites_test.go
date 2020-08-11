package application

import (
	"testing"
)

func TestGetFavorites(t *testing.T) {
	favs, err := GetFavorites(TestToken, 1, 20)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(favs)
}

func TestGetFavoriteIDs(t *testing.T) {
	favs, err := GetFavoriteIDs(TestToken, 1, 20)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(favs)
}

func TestRemoveFavorite(t *testing.T) {
	sts, err := RemoveFavorite(TestToken, "4511668538965141")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(sts)
}
