// Copyright (c) 2021 Proton Technologies AG
//
// This file is part of ProtonMail Bridge.
//
// ProtonMail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ProtonMail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with ProtonMail Bridge.  If not, see <https://www.gnu.org/licenses/>.

// Package settings provides access to persistent user settings.
package settings

import (
	"path/filepath"
)

// Keys of preferences in JSON file.
const (
	APIPortKey            = "UserPortApi"
	IMAPPortKey           = "UserPortImap"
	SMTPPortKey           = "UserPortSmtp"
	AllowProxyKey         = "AllowProxy"
	CacheEnabledKey       = "CacheEnabled"
	CacheCompressionKey   = "CacheCompression"
	CacheLocationKey      = "CacheLocation"
	CacheMinFreeAbsKey    = "CacheMinFreeAbs"
	CacheMinFreeRatKey    = "CacheMinFreeRat"
	CacheConcurrencyRead  = "CacheConcurrentRead"
	CacheConcurrencyWrite = "CacheConcurrentWrite"
	IMAPWorkers           = "ImapWorkers"
	FetchWorkers          = "FetchWorkers"
	AttachmentWorkers     = "AttachmentWorkers"
	CacheDir              = "CacheDir"
	TLSDir                = "TlsDir"
	CookieJar             = "CookieJar"
)

type Settings struct {
	*keyValueStore

	settingsDir string
}

func New(settingsPath string) *Settings {
	s := &Settings{
		keyValueStore: newKeyValueStore(settingsPath),
		settingsDir:   filepath.Dir(settingsPath),
	}

	s.setDefaultValues()

	return s
}

const (
	DefaultIMAPPort = "1143"
	DefaultSMTPPort = "1025"
	DefaultAPIPort  = "1042"
)

func (s *Settings) setDefaultValues() {
	s.setDefault(AllowProxyKey, "false")
	s.setDefault(CacheEnabledKey, "true")
	s.setDefault(CacheCompressionKey, "true")
	s.setDefault(CacheLocationKey, "")
	s.setDefault(CacheMinFreeAbsKey, "250000000")
	s.setDefault(CacheMinFreeRatKey, "")
	s.setDefault(CacheConcurrencyRead, "16")
	s.setDefault(CacheConcurrencyWrite, "16")
	s.setDefault(IMAPWorkers, "16")
	s.setDefault(FetchWorkers, "16")
	s.setDefault(AttachmentWorkers, "16")
	s.setDefault(APIPortKey, DefaultAPIPort)
	s.setDefault(IMAPPortKey, DefaultIMAPPort)
	s.setDefault(SMTPPortKey, DefaultSMTPPort)

	s.setDefault(CacheDir, filepath.Join(s.settingsDir, "cache"))
	s.setDefault(TLSDir, s.settingsDir)
	s.setDefault(CookieJar, filepath.Join(s.settingsDir, "cookies.json"))
}
