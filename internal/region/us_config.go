package region

import (
	"github.com/wso2/integration-platform-tools/pkg/api/platformgit"
)

// US Region Configurations
var DEFAULT_ENV_CONFIG = BuildEnvConfig(
	"https://console.choreo.dev",
	"https://apis.choreo.dev",
	"https://sts.choreo.dev",
	"https://choreocontrolplane.choreo.dev",
	"https://app.choreo.dev",
	"https://apis.choreo.dev",
	"https://apis.choreo.dev",
	"/projects/1.0.0/graphql",
	RegionOverrides{
		AsgardeoClientId: "aVKhTSUMu_QfEwmCtrcuWoLy92oa",
		STSClientID:      "Fu48ZbRFQsvzGe3ZVkGt2W5K1yga",
		STSTokenUrl:      "https://sts.choreo.dev/oauth2/token",
		AsgardeoScope:    "openid+email+profile",
		STSScopes:        "urn:choreosystem:componentutils:component_manage urn:choreosystem:componentutils:component_file_view urn:choreosystem:configmanagement:config_create urn:choreosystem:configmanagement:config_delete urn:choreosystem:configmanagement:config_manage urn:choreosystem:configmanagement:config_view urn:choreosystem:componentsmanagement:component_create urn:choreosystem:choreodevopsportalapi:component_manage urn:choreosystem:choreodevopsportalapi:deployment_manage urn:choreosystem:choreodevopsportalapi:deployment_view urn:choreosystem:componentsmanagement:component_logs_view urn:choreosystem:componentsmanagement:component_init_view choreo:project_view urn:choreosystem:organizationapi:org_manage urn:choreosystem:componentsmanagement:component_file_view urn:choreosystem:componentsmanagement:component_manage choreo:domain_manage choreo:domain_view choreo:url_mapping_manage choreo:url_mapping_approve choreo:url_mapping_view",
		GhApp: platformgit.GHAppConfig{
			AppUrl:      "https://github.com/apps/wso2-cloud-app",
			InstallUrl:  "https://github.com/apps/wso2-cloud-app/installations/new",
			AuthUrl:     "https://github.com/login/oauth/authorize",
			ClientId:    "Iv1.804167a242012c66",
			RedirectUrl: "https://console.choreo.dev/ghapp",
		},
		BillingUrl:           "https://subscriptions.wso2.com",
		SysApiPrefixUrl:      "https://5659b6b7-1063-41ed-8e39-d91857699255-systemapis",
		DevopsApiBaseUrl:     "https://apis.choreo.dev/devops/1.0.0",
		ApimPublisherBaseUrl: "https://sts.choreo.dev/api/am/publisher/v2",
		TokenUrl:             "https://api.asgardeo.io/t/a/oauth2/token",
		PlatformHostname:     "customdns.e1-us-east-azure",
		DevantConfig: DevantConfig{
			ConsoleUrl:       "https://console.devant.dev",
			AsgardeoClientId: "09YlJuqQZdFNRDC0sx3DHHDnZvIa",
		},
	},
)

var ENV_CONFIG_STAGE = BuildEnvConfig(
	"https://console.st.choreo.dev",
	"https://apis.st.choreo.dev",
	"https://sts.st.choreo.dev",
	"https://choreocontrolplane.st.choreo.dev",
	"https://app.st.choreo.dev",
	"https://apis.st.choreo.dev",
	"https://apis.st.choreo.dev",
	"/projects/1.0.0/graphql",
	RegionOverrides{
		AsgardeoClientId: "NoOBydRztff7iENCq0LM2uuRs2ca",
		STSClientID:      "ZNPW_kOxRfzgd7vzkhLYIaLwRgMa",
		STSTokenUrl:      "https://sts.st.choreo.dev:443/oauth2/token",
		AsgardeoScope:    "openid+email+profile",
		STSScopes:        "urn:choreosystem:componentutils:component_manage urn:choreosystem:componentutils:component_file_view urn:choreosystem:configmanagement:config_create urn:choreosystem:configmanagement:config_delete urn:choreosystem:configmanagement:config_manage urn:choreosystem:configmanagement:config_view urn:choreosystem:componentsmanagement:component_create urn:choreosystem:choreodevopsportalapi:component_manage urn:choreosystem:choreodevopsportalapi:deployment_manage urn:choreosystem:choreodevopsportalapi:deployment_view urn:choreosystem:componentsmanagement:component_logs_view urn:choreosystem:componentsmanagement:component_init_view choreo:project_view urn:choreosystem:organizationapi:org_manage urn:choreosystem:componentsmanagement:component_file_view urn:choreosystem:componentsmanagement:component_manage choreo:domain_manage choreo:domain_view choreo:url_mapping_manage choreo:url_mapping_approve choreo:url_mapping_view",
		GhApp: platformgit.GHAppConfig{
			AppUrl:      "https://github.com/apps/wso2-cloud-app-stage",
			InstallUrl:  "https://github.com/apps/wso2-cloud-app-stage/installations/new",
			AuthUrl:     "https://github.com/login/oauth/authorize",
			ClientId:    "Iv1.20fd2645fc8a5aab",
			RedirectUrl: "https://console.st.choreo.dev/ghapp",
		},
		BillingUrl:           "https://subscriptions.st.wso2.com",
		SysApiPrefixUrl:      "https://c1b2cfb2-e965-4d28-b36d-b34f162ecc30-systemapis",
		DevopsApiBaseUrl:     "https://apis.st.choreo.dev/devops/1.0.0",
		ApimPublisherBaseUrl: "https://sts.st.choreo.dev/api/am/publisher/v2",
		TokenUrl:             "https://stage.api.asgardeo.io/t/a/oauth2/token",
		PlatformHostname:     "customdns.e1-us-east-azure.st",
		DevantConfig: DevantConfig{
			ConsoleUrl:       "https://preview-st.devant.dev",
			AsgardeoClientId: "fO22Kjf5AIZSGRO4R3kYUgTadyYa",
		},
	},
)

var ENV_CONFIG_DEV = BuildEnvConfig(
	"https://consolev2.preview-dv.choreo.dev",
	"https://apis.preview-dv.choreo.dev",
	"https://sts.preview-dv.choreo.dev",
	"https://choreocontrolplane.preview-dv.choreo.dev",
	"https://app.preview-dv.choreo.dev",
	"https://apis.preview-dv.choreo.dev",
	"https://apis.preview-dv.choreo.dev",
	"/projects/1.0.0/graphql",
	RegionOverrides{
		AsgardeoClientId: "_eEveWFdTSJPaui7DmCuU5DUrUEa",
		STSClientID:      "LYtnNTW5NFrISSB4fMiXcnFTIasa",
		STSTokenUrl:      "https://sts.preview-dv.choreo.dev:443/oauth2/token",
		AsgardeoScope:    "openid+email+profile",
		STSScopes:        "urn:choreocontrolplane:componentutils:component_manage urn:choreocontrolplane:componentutils:component_file_view urn:choreocontrolplane:configmanagement:config_create urn:choreocontrolplane:configmanagement:config_delete urn:choreocontrolplane:configmanagement:config_manage urn:choreocontrolplane:configmanagement:config_view urn:choreocontrolplane:componentsmanagement:component_create urn:choreocontrolplane:choreodevopsportalapi:component_manage urn:choreocontrolplane:choreodevopsportalapi:deployment_manage urn:choreocontrolplane:choreodevopsportalapi:deployment_view urn:choreocontrolplane:componentsmanagement:component_logs_view urn:choreocontrolplane:componentsmanagement:component_init_view choreo:project_view urn:choreocontrolplane:organizationapi:org_manage urn:choreocontrolplane:componentsmanagement:component_file_view urn:choreocontrolplane:componentsmanagement:component_manage choreo:domain_manage choreo:domain_view choreo:url_mapping_manage choreo:url_mapping_approve choreo:url_mapping_view",
		GhApp: platformgit.GHAppConfig{
			AppUrl:      "https://github.com/apps/wso2-cloud-app-dev",
			InstallUrl:  "https://github.com/apps/wso2-cloud-app-dev/installations/new",
			AuthUrl:     "https://github.com/login/oauth/authorize",
			ClientId:    "Iv1.f6cf2cd585148ee7",
			RedirectUrl: "https://consolev2.preview-dv.choreo.dev/ghapp",
		},
		BillingUrl:           "https://subscriptions.dv.wso2.com",
		SysApiPrefixUrl:      "https://783c6c4d-8b9b-4190-b70a-e717ab1ee739-systemapis",
		DevopsApiBaseUrl:     "https://apis.preview-dv.choreo.dev/devops/1.0.0",
		ApimPublisherBaseUrl: "https://sts.preview-dv.choreo.dev/api/am/publisher/v2",
		TokenUrl:             "https://dev.api.asgardeo.io/t/a/oauth2/token",
		PlatformHostname:     "customdns.e1-us-east-azure.preview-dv",
		DevantConfig: DevantConfig{
			ConsoleUrl:       "https://preview-dv.devant.dev",
			AsgardeoClientId: "zL9kF4GCPiN2veO8judQvwlqLb8a",
		},
	},
)

func init() {
	InitUSRegion(&ENV_CONFIG_DEV, &ENV_CONFIG_STAGE, &DEFAULT_ENV_CONFIG)
}
