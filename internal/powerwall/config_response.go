package powerwall

import "time"

type ConfigResponse struct {
	Vin    string `json:"vin,omitempty"`
	Meters []struct {
		Location   string `json:"location,omitempty"`
		Type       string `json:"type,omitempty"`
		Cts        []bool `json:"cts,omitempty"`
		Inverted   []bool `json:"inverted,omitempty"`
		Connection struct {
			IPAddress       string `json:"ip_address,omitempty"`
			Port            int    `json:"port,omitempty"`
			ShortID         string `json:"short_id,omitempty"`
			DeviceSerial    string `json:"device_serial,omitempty"`
			NeurioConnected bool   `json:"neurio_connected,omitempty"`
			HTTPSConf       struct {
				ClientCert          string `json:"client_cert,omitempty"`
				ClientKey           string `json:"client_key,omitempty"`
				ServerCaCert        string `json:"server_ca_cert,omitempty"`
				ServerName          string `json:"server_name,omitempty"`
				MaxIdleConnsPerHost int    `json:"max_idle_conns_per_host,omitempty"`
			} `json:"https_conf,omitempty"`
		} `json:"connection,omitempty"`
		RealPowerScaleFactor float64 `json:"real_power_scale_factor,omitempty"`
	} `json:"meters,omitempty"`
	BatteryBlocks []struct {
		Vin                      string `json:"vin,omitempty"`
		MinSoe                   int    `json:"min_soe,omitempty"`
		MaxSoe                   int    `json:"max_soe,omitempty"`
		Type                     string `json:"type,omitempty"`
		CanConnection            string `json:"can_connection,omitempty"`
		PviPowerStatus           string `json:"pvi_power_status,omitempty"`
		EnableInverterSolarMeter bool   `json:"enable_inverter_solar_meter,omitempty"`
	} `json:"battery_blocks,omitempty"`
	IslandContactorController struct {
		Type       string `json:"type,omitempty"`
		Connection struct {
			CanConnection string `json:"can_connection,omitempty"`
		} `json:"connection,omitempty"`
		NumberOfPhases int `json:"number_of_phases,omitempty"`
	} `json:"island_contactor_controller,omitempty"`
	IslandConfig struct {
	} `json:"island_config,omitempty"`
	Dio struct {
	} `json:"dio,omitempty"`
	SiteInfo struct {
		CustomerPreferredExportRule string  `json:"customer_preferred_export_rule,omitempty"`
		BatteryCommissionDate       string  `json:"battery_commission_date,omitempty"`
		BackupReservePercent        int     `json:"backup_reserve_percent,omitempty"`
		MaxSiteMeterPowerAc         int     `json:"max_site_meter_power_ac,omitempty"`
		MinSiteMeterPowerAc         int     `json:"min_site_meter_power_ac,omitempty"`
		NominalSystemEnergyAc       int     `json:"nominal_system_energy_ac,omitempty"`
		NominalSystemPowerAc        float64 `json:"nominal_system_power_ac,omitempty"`
		TariffContent               struct {
			Code         string `json:"code,omitempty"`
			Name         string `json:"name,omitempty"`
			Utility      string `json:"utility,omitempty"`
			DailyCharges []struct {
				Amount int    `json:"amount,omitempty"`
				Name   string `json:"name,omitempty"`
			} `json:"daily_charges,omitempty"`
			DemandCharges struct {
				All struct {
					All int `json:"ALL,omitempty"`
				} `json:"ALL,omitempty"`
				Summer struct {
				} `json:"Summer,omitempty"`
				Winter struct {
				} `json:"Winter,omitempty"`
			} `json:"demand_charges,omitempty"`
			EnergyCharges struct {
				All struct {
					All int `json:"ALL,omitempty"`
				} `json:"ALL,omitempty"`
				Summer struct {
					OffPeak float64 `json:"OFF_PEAK,omitempty"`
					OnPeak  float64 `json:"ON_PEAK,omitempty"`
				} `json:"Summer,omitempty"`
				Winter struct {
				} `json:"Winter,omitempty"`
			} `json:"energy_charges,omitempty"`
			Seasons struct {
				Summer struct {
					FromDay    int `json:"fromDay,omitempty"`
					ToDay      int `json:"toDay,omitempty"`
					FromMonth  int `json:"fromMonth,omitempty"`
					ToMonth    int `json:"toMonth,omitempty"`
					TouPeriods struct {
						OffPeak []struct {
							FromDayOfWeek int `json:"fromDayOfWeek,omitempty"`
							ToDayOfWeek   int `json:"toDayOfWeek,omitempty"`
							FromHour      int `json:"fromHour,omitempty"`
							FromMinute    int `json:"fromMinute,omitempty"`
							ToHour        int `json:"toHour,omitempty"`
							ToMinute      int `json:"toMinute,omitempty"`
						} `json:"OFF_PEAK,omitempty"`
						OnPeak []struct {
							FromDayOfWeek int `json:"fromDayOfWeek,omitempty"`
							ToDayOfWeek   int `json:"toDayOfWeek,omitempty"`
							FromHour      int `json:"fromHour,omitempty"`
							FromMinute    int `json:"fromMinute,omitempty"`
							ToHour        int `json:"toHour,omitempty"`
							ToMinute      int `json:"toMinute,omitempty"`
						} `json:"ON_PEAK,omitempty"`
					} `json:"tou_periods,omitempty"`
				} `json:"Summer,omitempty"`
				Winter struct {
					FromDay    int `json:"fromDay,omitempty"`
					ToDay      int `json:"toDay,omitempty"`
					FromMonth  int `json:"fromMonth,omitempty"`
					ToMonth    int `json:"toMonth,omitempty"`
					TouPeriods struct {
					} `json:"tou_periods,omitempty"`
				} `json:"Winter,omitempty"`
			} `json:"seasons,omitempty"`
			SellTariff struct {
				Name         string `json:"name,omitempty"`
				Utility      string `json:"utility,omitempty"`
				DailyCharges []struct {
					Amount int    `json:"amount,omitempty"`
					Name   string `json:"name,omitempty"`
				} `json:"daily_charges,omitempty"`
				DemandCharges struct {
					All struct {
						All int `json:"ALL,omitempty"`
					} `json:"ALL,omitempty"`
					Summer struct {
					} `json:"Summer,omitempty"`
					Winter struct {
					} `json:"Winter,omitempty"`
				} `json:"demand_charges,omitempty"`
				EnergyCharges struct {
					All struct {
						All int `json:"ALL,omitempty"`
					} `json:"ALL,omitempty"`
					Summer struct {
						OffPeak float64 `json:"OFF_PEAK,omitempty"`
						OnPeak  float64 `json:"ON_PEAK,omitempty"`
					} `json:"Summer,omitempty"`
					Winter struct {
					} `json:"Winter,omitempty"`
				} `json:"energy_charges,omitempty"`
				Seasons struct {
					Summer struct {
						FromDay    int `json:"fromDay,omitempty"`
						ToDay      int `json:"toDay,omitempty"`
						FromMonth  int `json:"fromMonth,omitempty"`
						ToMonth    int `json:"toMonth,omitempty"`
						TouPeriods struct {
							OffPeak []struct {
								FromDayOfWeek int `json:"fromDayOfWeek,omitempty"`
								ToDayOfWeek   int `json:"toDayOfWeek,omitempty"`
								FromHour      int `json:"fromHour,omitempty"`
								FromMinute    int `json:"fromMinute,omitempty"`
								ToHour        int `json:"toHour,omitempty"`
								ToMinute      int `json:"toMinute,omitempty"`
							} `json:"OFF_PEAK,omitempty"`
							OnPeak []struct {
								FromDayOfWeek int `json:"fromDayOfWeek,omitempty"`
								ToDayOfWeek   int `json:"toDayOfWeek,omitempty"`
								FromHour      int `json:"fromHour,omitempty"`
								FromMinute    int `json:"fromMinute,omitempty"`
								ToHour        int `json:"toHour,omitempty"`
								ToMinute      int `json:"toMinute,omitempty"`
							} `json:"ON_PEAK,omitempty"`
						} `json:"tou_periods,omitempty"`
					} `json:"Summer,omitempty"`
					Winter struct {
						FromDay    int `json:"fromDay,omitempty"`
						ToDay      int `json:"toDay,omitempty"`
						FromMonth  int `json:"fromMonth,omitempty"`
						ToMonth    int `json:"toMonth,omitempty"`
						TouPeriods struct {
						} `json:"tou_periods,omitempty"`
					} `json:"Winter,omitempty"`
				} `json:"seasons,omitempty"`
			} `json:"sell_tariff,omitempty"`
		} `json:"tariff_content,omitempty"`
		GridCode          string `json:"grid_code,omitempty"`
		GridCodeOverrides []struct {
			Name  string  `json:"name,omitempty"`
			Value float64 `json:"value,omitempty"`
		} `json:"grid_code_overrides,omitempty"`
		Country         string `json:"country,omitempty"`
		State           string `json:"state,omitempty"`
		Distributor     string `json:"distributor,omitempty"`
		Utility         string `json:"utility,omitempty"`
		Retailer        string `json:"retailer,omitempty"`
		Region          string `json:"region,omitempty"`
		SiteName        string `json:"site_name,omitempty"`
		Timezone        string `json:"timezone,omitempty"`
		ITCCliff        int    `json:"ITC_cliff,omitempty"`
		PanelMaxCurrent int    `json:"panel_max_current,omitempty"`
	} `json:"site_info,omitempty"`
	Strategy struct {
		Control          string `json:"control,omitempty"`
		ForecastMethod   string `json:"forecast_method,omitempty"`
		PvForecastMethod string `json:"pv_forecast_method,omitempty"`
		TOUMode          string `json:"TOU_mode,omitempty"`
	} `json:"strategy,omitempty"`
	DefaultRealMode             string `json:"default_real_mode,omitempty"`
	EnableInverterMeterReadings bool   `json:"enable_inverter_meter_readings,omitempty"`
	FreqShiftLoadShed           struct {
		Soe    int     `json:"soe,omitempty"`
		DeltaF float64 `json:"delta_f,omitempty"`
	} `json:"freq_shift_load_shed,omitempty"`
	FreqSupportParameters struct {
		TargetSoe               int `json:"target_soe,omitempty"`
		LowSoeLimit             int `json:"low_soe_limit,omitempty"`
		HighSoeLimit            int `json:"high_soe_limit,omitempty"`
		OffsetGain              int `json:"offset_gain,omitempty"`
		MaxChargePowerOffset    int `json:"max_charge_power_offset,omitempty"`
		MaxDischargePowerOffset int `json:"max_discharge_power_offset,omitempty"`
	} `json:"freq_support_parameters,omitempty"`
	Solar struct {
		Brand            string `json:"brand,omitempty"`
		Model            string `json:"model,omitempty"`
		PowerRatingWatts int    `json:"power_rating_watts,omitempty"`
	} `json:"solar,omitempty"`
	Solars []struct {
		Brand            string `json:"brand,omitempty"`
		Model            string `json:"model,omitempty"`
		PowerRatingWatts int    `json:"power_rating_watts,omitempty"`
	} `json:"solars,omitempty"`
	Logging struct {
		DatapumpLogRateMs int `json:"datapump_log_rate_ms,omitempty"`
	} `json:"logging,omitempty"`
	Installer struct {
		Email                  string `json:"email,omitempty"`
		CustomerID             string `json:"customer_id,omitempty"`
		Company                string `json:"company,omitempty"`
		Phone                  string `json:"phone,omitempty"`
		SolarInstallationType  string `json:"solar_installation_type,omitempty"`
		RunSitemaster          bool   `json:"run_sitemaster,omitempty"`
		VerifiedConfig         bool   `json:"verified_config,omitempty"`
		Location               string `json:"location,omitempty"`
		Mounting               string `json:"mounting,omitempty"`
		Wiring                 string `json:"wiring,omitempty"`
		BackupConfiguration    string `json:"backup_configuration,omitempty"`
		SolarInstallation      string `json:"solar_installation,omitempty"`
		HasPowerlineToEthernet bool   `json:"has_powerline_to_ethernet,omitempty"`
		HasCellularModemToWifi bool   `json:"has_cellular_modem_to_wifi,omitempty"`
	} `json:"installer,omitempty"`
	Customer struct {
		Registered bool `json:"registered,omitempty"`
	} `json:"customer,omitempty"`
	IndustrialNetworks struct {
	} `json:"industrial_networks,omitempty"`
	Credentials []struct {
		AccessRoles []string `json:"access_roles,omitempty"`
		Username    string   `json:"username,omitempty"`
		Password    string   `json:"password,omitempty"`
	} `json:"credentials,omitempty"`
	AutoMeterUpdate bool `json:"auto_meter_update,omitempty"`
	BridgeInverter  struct {
		Enabled bool      `json:"enabled,omitempty"`
		EndTime time.Time `json:"end_time,omitempty"`
	} `json:"bridge_inverter,omitempty"`
	ClientProtocols struct {
		Dnp3                  bool `json:"dnp3,omitempty"`
		M2M                   bool `json:"m2m,omitempty"`
		Mtls                  bool `json:"mtls,omitempty"`
		Modbus                bool `json:"modbus,omitempty"`
		AllowPublicIP         bool `json:"allow_public_ip,omitempty"`
		AllowInternalNetworks bool `json:"allow_internal_networks,omitempty"`
		ChargingKiosk         bool `json:"charging_kiosk,omitempty"`
	} `json:"client_protocols,omitempty"`
	TestTimers struct {
	} `json:"test_timers,omitempty"`
	SolarPowerwall bool `json:"solar_powerwall,omitempty"`
}
