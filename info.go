package main

type Info struct {
	Status  string `json:"status"`
	Details struct {
		DiskSpace struct {
			Status  string `json:"status"`
			Details struct {
				Total     int64 `json:"total"`
				Free      int64 `json:"free"`
				Threshold int   `json:"threshold"`
			} `json:"details"`
		} `json:"diskSpace"`
		Jms struct {
			Status  string `json:"status"`
			Details struct {
				Provider string `json:"provider"`
			} `json:"details"`
		} `json:"jms"`
		Db struct {
			Status  string `json:"status"`
			Details struct {
				Database string `json:"database"`
				Hello    int    `json:"hello"`
			} `json:"details"`
		} `json:"db"`
		RefreshScope struct {
			Status string `json:"status"`
		} `json:"refreshScope"`
		DiscoveryComposite struct {
			Status  string `json:"status"`
			Details struct {
				DiscoveryClient struct {
					Status  string `json:"status"`
					Details struct {
						Services []string `json:"services"`
					} `json:"details"`
				} `json:"discoveryClient"`
				Eureka struct {
					Description string `json:"description"`
					Status      string `json:"status"`
					Details     struct {
						Applications struct {
							TIMEOFFSERVICE             int `json:"TIMEOFF-SERVICE"`
							LOOKUPSERVICE              int `json:"LOOKUP-SERVICE"`
							EMAILCONSUMER              int `json:"EMAIL-CONSUMER"`
							ONBOARDINGSERVICE          int `json:"ONBOARDING-SERVICE"`
							COMPENSATIONSERVICE        int `json:"COMPENSATION-SERVICE"`
							DATASYNCSERVICE            int `json:"DATASYNC-SERVICE"`
							DISPATCHERSERVICE          int `json:"DISPATCHER-SERVICE"`
							PERSONSERVICE              int `json:"PERSON-SERVICE"`
							NOTIFICATIONSERVICE        int `json:"NOTIFICATION-SERVICE"`
							CLIENTSERVICEPROXYHR       int `json:"CLIENT-SERVICE-PROXY-HR"`
							SERVICEDISCOVERY           int `json:"SERVICE-DISCOVERY"`
							REPORTSERVICE              int `json:"REPORT-SERVICE"`
							EVENTLOGGERSERVICE         int `json:"EVENTLOGGER-SERVICE"`
							DOCUMENTSERVICE            int `json:"DOCUMENT-SERVICE"`
							USERSERVICE                int `json:"USER-SERVICE"`
							TRANSACTIONRECORDERSERVICE int `json:"TRANSACTIONRECORDER-SERVICE"`
							EMAILSERVICE               int `json:"EMAIL-SERVICE"`
							Jwtms                      int `json:"JWTMS"`
							Ppadapter                  int `json:"PPADAPTER"`
						} `json:"applications"`
					} `json:"details"`
				} `json:"eureka"`
			} `json:"details"`
		} `json:"discoveryComposite"`
		Hystrix struct {
			Status string `json:"status"`
		} `json:"hystrix"`
	} `json:"details"`
}
