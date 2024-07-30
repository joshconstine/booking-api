package sb

import (
	"booking-api/config"

	"github.com/nedpals/supabase-go"
)

const BaseAuthURL = "https://bvutljexdthpferjwsgz.supabase.co/auth/v1/recover"

// https://<project_ref>.supabase.co/rest/v1/
const ResetPasswordEndpoint = "auth/v1/recover"

var ClientInstance *supabase.Client

func CreateAuthClient(env config.EnvVars) error {
	sbHost := env.SUPABASE_URL
	sbSecret := env.SUPABASE_SECRET
	ClientInstance = supabase.CreateClient(sbHost, sbSecret)
	return nil
}
