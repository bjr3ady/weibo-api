package definitions

//City is the weibo city struct
type City struct {
	CityCode string
	CityName string
}

//Location is the weibo code to location struct
type Location struct {
	LocationCode string
	LocationName string
}

//Country is the weibo country struct
type Country struct {
	CountryCode string
	CountryName string
}

//Province is the weibo province struct
type Province struct {
	ProvinceCode string
	ProvinceName string
}

//Timezone is the timezone struct
type Timezone struct {
	ZoneCode  int
	ZoneValue string
}