// Copyright 2019 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
package main

import (
	"github.com/newrelic/nri-prometheus/internal/cmd/scraper"
	"github.com/newrelic/nri-prometheus/internal/integration"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		logrus.WithError(err).Fatal("while loading configuration")
	}

	logrus.Infof("Starting New Relic's Prometheus OpenMetrics Integration version %s", integration.Version)
	logrus.Debugf("Config: %#v", cfg)

	err = scraper.Run(cfg)
	if err != nil {
		logrus.WithError(err).Fatal("error occurred while running scraper")
	}
}
