package icloud

import (
	"fmt"
)

type loginRequest struct {
	AppleID       string `json:"apple_id"`
	Password      string `json:"password"`
	ExtendedLogin bool   `json:"extended_login"`
}

type ValidateResponse struct {
	Error

	DsInfo                       *DsInfo      `json:"dsInfo"`
	HasMinimumDeviceForPhotosWeb bool         `json:"hasMinimumDeviceForPhotosWeb"`
	ICDPEnabled                  bool         `json:"iCDPEnabled"`
	Webservices                  *Webservices `json:"webservices"`
	PcsEnabled                   bool         `json:"pcsEnabled"`
	ConfigBag                    *ConfigBag   `json:"configBag"`
	HsaTrustedBrowser            bool         `json:"hsaTrustedBrowser"`
	AppsOrder                    []string     `json:"appsOrder"`
	Version                      int          `json:"version"`
	IsExtendedLogin              bool         `json:"isExtendedLogin"`
	PcsServiceIdentitiesIncluded bool         `json:"pcsServiceIdentitiesIncluded"`
	HsaChallengeRequired         bool         `json:"hsaChallengeRequired"`
	RequestInfo                  *RequestInfo `json:"requestInfo"`
	PcsDeleted                   bool         `json:"pcsDeleted"`
	ICloudInfo                   *ICloudInfo  `json:"iCloudInfo"`
	Apps                         *Apps        `json:"apps"`
}

type Error struct {
	InitSuccess bool   `json:"success"`
	InitError   string `json:"error"`

	Status       int         `json:"status"`
	Errorcode    interface{} `json:"errorcode"`
	Message      string      `json:"message"`
	Authmismatch bool        `json:"authmismatch"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%#v", e)
}

type AppleIDEntity struct {
	IsPrimary bool   `json:"isPrimary"`
	Type      string `json:"type"`
	Value     string `json:"value"`
}

type DsInfo struct {
	LastName                  string           `json:"lastName"`
	ICDPEnabled               bool             `json:"iCDPEnabled"`
	Dsid                      string           `json:"dsid"`
	HsaEnabled                bool             `json:"hsaEnabled"`
	IroncadeMigrated          bool             `json:"ironcadeMigrated"`
	Locale                    string           `json:"locale"`
	BrZoneConsolidated        bool             `json:"brZoneConsolidated"`
	IsManagedAppleID          bool             `json:"isManagedAppleID"`
	GilliganInvited           string           `json:"gilligan-invited"`
	AppleIDAliases            []interface{}    `json:"appleIdAliases"`
	HsaVersion                int              `json:"hsaVersion"`
	IsPaidDeveloper           bool             `json:"isPaidDeveloper"`
	CountryCode               string           `json:"countryCode"`
	NotificationID            string           `json:"notificationId"`
	PrimaryEmailVerified      bool             `json:"primaryEmailVerified"`
	ADsID                     string           `json:"aDsID"`
	Locked                    bool             `json:"locked"`
	HasICloudQualifyingDevice bool             `json:"hasICloudQualifyingDevice"`
	PrimaryEmail              string           `json:"primaryEmail"`
	AppleIDEntries            []*AppleIDEntity `json:"appleIdEntries"`
	GilliganEnabled           string           `json:"gilligan-enabled"`
	FullName                  string           `json:"fullName"`
	LanguageCode              string           `json:"languageCode"`
	AppleID                   string           `json:"appleId"`
	FirstName                 string           `json:"firstName"`
	ICloudAppleIDAlias        string           `json:"iCloudAppleIdAlias"`
	NotesMigrated             bool             `json:"notesMigrated"`
	HasPaymentInfo            bool             `json:"hasPaymentInfo"`
	PcsDeleted                bool             `json:"pcsDeleted"`
	AppleIDAlias              string           `json:"appleIdAlias"`
	BrMigrated                bool             `json:"brMigrated"`
	StatusCode                int              `json:"statusCode"`
}

type Webservices struct {
	Reminders struct {
		URL    string `json:"url"`
		Status string `json:"status"`
	} `json:"reminders"`
	Calendar struct {
		URL    string `json:"url"`
		Status string `json:"status"`
	} `json:"calendar"`
	Docws struct {
		PcsRequired bool   `json:"pcsRequired"`
		URL         string `json:"url"`
		Status      string `json:"status"`
	} `json:"docws"`
	Settings struct {
		URL    string `json:"url"`
		Status string `json:"status"`
	} `json:"settings"`
	Ubiquity struct {
		URL    string `json:"url"`
		Status string `json:"status"`
	} `json:"ubiquity"`
	Streams struct {
		URL    string `json:"url"`
		Status string `json:"status"`
	} `json:"streams"`
	Keyvalue struct {
		URL    string `json:"url"`
		Status string `json:"status"`
	} `json:"keyvalue"`
	Ckdatabasews struct {
		PcsRequired bool   `json:"pcsRequired"`
		URL         string `json:"url"`
		Status      string `json:"status"`
	} `json:"ckdatabasews"`
	Photosupload struct {
		PcsRequired bool   `json:"pcsRequired"`
		URL         string `json:"url"`
		Status      string `json:"status"`
	} `json:"photosupload"`
	Archivews struct {
		URL    string `json:"url"`
		Status string `json:"status"`
	} `json:"archivews"`
	Photos struct {
		PcsRequired bool   `json:"pcsRequired"`
		UploadURL   string `json:"uploadUrl"`
		URL         string `json:"url"`
		Status      string `json:"status"`
	} `json:"photos"`
	Push struct {
		URL    string `json:"url"`
		Status string `json:"status"`
	} `json:"push"`
	Drivews struct {
		PcsRequired bool   `json:"pcsRequired"`
		URL         string `json:"url"`
		Status      string `json:"status"`
	} `json:"drivews"`
	Uploadimagews struct {
		URL    string `json:"url"`
		Status string `json:"status"`
	} `json:"uploadimagews"`
	Iwmb struct {
		URL    string `json:"url"`
		Status string `json:"status"`
	} `json:"iwmb"`
	Cksharews struct {
		URL    string `json:"url"`
		Status string `json:"status"`
	} `json:"cksharews"`
	Iworkexportws struct {
		URL    string `json:"url"`
		Status string `json:"status"`
	} `json:"iworkexportws"`
	Geows struct {
		URL    string `json:"url"`
		Status string `json:"status"`
	} `json:"geows"`
	Findme struct {
		URL    string `json:"url"`
		Status string `json:"status"`
	} `json:"findme"`
	Ckdeviceservice struct {
		URL string `json:"url"`
	} `json:"ckdeviceservice"`
	Iworkthumbnailws struct {
		URL    string `json:"url"`
		Status string `json:"status"`
	} `json:"iworkthumbnailws"`
	Account struct {
		ICloudEnv struct {
			ShortID string `json:"shortId"`
		} `json:"iCloudEnv"`
		URL    string `json:"url"`
		Status string `json:"status"`
	} `json:"account"`
	Fmf struct {
		URL    string `json:"url"`
		Status string `json:"status"`
	} `json:"fmf"`
	Contacts struct {
		URL    string `json:"url"`
		Status string `json:"status"`
	} `json:"contacts"`
}

type ConfigBag struct {
	Urls struct {
		AccountCreateUI     string `json:"accountCreateUI"`
		AccountLoginUI      string `json:"accountLoginUI"`
		AccountLogin        string `json:"accountLogin"`
		AccountRepairUI     string `json:"accountRepairUI"`
		DownloadICloudTerms string `json:"downloadICloudTerms"`
		RepairDone          string `json:"repairDone"`
		VettingURLForEmail  string `json:"vettingUrlForEmail"`
		AccountCreate       string `json:"accountCreate"`
		GetICloudTerms      string `json:"getICloudTerms"`
		VettingURLForPhone  string `json:"vettingUrlForPhone"`
	} `json:"urls"`
	AccountCreateEnabled string `json:"accountCreateEnabled"`
}

type RequestInfo struct {
	Country         string `json:"country"`
	TimeZone        string `json:"timeZone"`
	IsAppleInternal bool   `json:"isAppleInternal"`
	Region          string `json:"region"`
}

type Apps struct {
	Calendar struct {
	} `json:"calendar"`
	Reminders struct {
	} `json:"reminders"`
	Keynote struct {
		IsQualifiedForBeta bool `json:"isQualifiedForBeta"`
	} `json:"keynote"`
	Settings struct {
		CanLaunchWithOneFactor bool `json:"canLaunchWithOneFactor"`
	} `json:"settings"`
	Mail struct {
	} `json:"mail"`
	Numbers struct {
		IsQualifiedForBeta bool `json:"isQualifiedForBeta"`
	} `json:"numbers"`
	Photos struct {
	} `json:"photos"`
	Pages struct {
		IsQualifiedForBeta bool `json:"isQualifiedForBeta"`
	} `json:"pages"`
	Find struct {
		CanLaunchWithOneFactor bool `json:"canLaunchWithOneFactor"`
	} `json:"find"`
	Notes2 struct {
	} `json:"notes2"`
	Iclouddrive struct {
	} `json:"iclouddrive"`
	Newspublisher struct {
		IsHidden bool `json:"isHidden"`
	} `json:"newspublisher"`
	Fmf struct {
	} `json:"fmf"`
	Contacts struct {
	} `json:"contacts"`
}

type ICloudInfo struct {
	SafariBookmarksHasMigratedToCloudKit bool `json:"SafariBookmarksHasMigratedToCloudKit"`
}
