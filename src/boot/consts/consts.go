package consts

const (
	PORT             = "8088"
	APPLICATION_JSON = "application/json"
	APPLICATION_XML  = "application/xml"

	//response header
	HEADER_CONTENTTYPE = "Content-Type"
	HEADER_ACCEPT      = "Accept"

	//Api functions
	SHIPPEDPROJECTS          = "projects"
	SHIPPEDPROJECTS_SERVICES = "projects/%s/services"
	SHIPPEDPROJECTS_ENVS     = "projects/%s/envs"
	SHIPPED_BUILDS_PACKS     = "buildpacks"
	SHIPPED_DEPENDENCIES     = "projects/%s/dependencies"
	SHIPPEDPROJECT_BUILDS    = "projects/%s/builds"
)
