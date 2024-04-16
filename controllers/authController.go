package controllers

import (
	"booking-api/config"
	"booking-api/data/request"
	validate "booking-api/pkg/kit"
	"booking-api/pkg/sb"
	"booking-api/services"
	auth "booking-api/view/auth"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/nedpals/supabase-go"
)

const (
	sessionUserKey        = "user"
	sessionAccessTokenKey = "accessToken"
)

type AuthController struct {
	UserService services.UserService
	sb          *supabase.Client
}

func NewAuthController(userService services.UserService, sb *supabase.Client) *AuthController {
	return &AuthController{
		UserService: userService,
		sb:          sb,
	}
}

func (ac *AuthController) HandleAccountSetupCreate(w http.ResponseWriter, r *http.Request) error {
	params := auth.AccountSetupParams{
		Username: r.FormValue("username"),
	}
	var errors auth.AccountSetupErrors
	ok := validate.New(&params, validate.Fields{
		"Username": validate.Rules(validate.Min(2), validate.Max(50)),
	}).Validate(&errors)
	if !ok {
		return render(r, w, auth.AccountSetupForm(params, errors))
	}
	user := GetAuthenticatedUser(r)
	account := request.CreateUserRequest{
		UserID:   user.User.UserID,
		Username: params.Username,
	}
	if err := ac.UserService.CreateUser(&account); err != nil {
		return err
	}
	return hxRedirect(w, r, "/")
}

func (ac *AuthController) HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
	return render(r, w, auth.Login())
}

func (ac *AuthController) HandleSignupIndex(w http.ResponseWriter, r *http.Request) error {
	return render(r, w, auth.Signup())
}

func HandleAccountSetupIndex(w http.ResponseWriter, r *http.Request) error {
	return render(r, w, auth.AccountSetup())
}

func (ac *AuthController) HandleLoginWithGoogle(w http.ResponseWriter, r *http.Request) error {
	resp, err := sb.ClientInstance.Auth.SignInWithProvider(supabase.ProviderSignInOptions{
		Provider:   "google",
		RedirectTo: "http://localhost:8080/auth/callback",
	})
	if err != nil {
		return err
	}
	http.Redirect(w, r, resp.URL, http.StatusSeeOther)
	return nil
}

func (ac *AuthController) HandleLoginCreate(w http.ResponseWriter, r *http.Request) error {
	credentials := supabase.UserCredentials{
		Email: r.FormValue("email"),
	}
	err := sb.ClientInstance.Auth.SendMagicLink(r.Context(), credentials.Email)
	if err != nil && err.Error() != "" {
		slog.Error("login error", "err", err)
		return render(r, w, auth.LoginForm(credentials, auth.LoginErrors{
			InvalidCredentials: err.Error(),
		}))
	}
	return render(r, w, auth.MagicLinkSuccess(credentials.Email))
}

func (ac *AuthController) HandleAuthCallback(w http.ResponseWriter, r *http.Request) error {
	accessToken := r.URL.Query().Get("access_token")
	if len(accessToken) == 0 {
		return render(r, w, auth.CallbackScript())
	}
	if err := ac.setAuthSession(w, r, accessToken); err != nil {
		return err
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}

func (ac *AuthController) HandleLogoutCreate(w http.ResponseWriter, r *http.Request) error {
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	session, _ := store.Get(r, sessionUserKey)
	session.Values[sessionAccessTokenKey] = ""
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
	return nil
}

func (ac *AuthController) setAuthSession(w http.ResponseWriter, r *http.Request, accessToken string) error {
	// load config
	env, err := config.LoadConfig(".")
	if err != nil {
		return err
	}
	store := sessions.NewCookieStore([]byte(env.SESSION_SECRET))
	session, _ := store.Get(r, sessionUserKey)
	session.Values[sessionAccessTokenKey] = accessToken
	return session.Save(r, w)
}
