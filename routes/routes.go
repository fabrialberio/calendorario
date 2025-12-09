package routes

import "net/http"

// Common path parameter names used for multiple routes.
const (
	KeyID      = "id"
	KeyDate    = "date"
	KeyBackURL = "back_url"
)

const (
	routeAdminBase     = "/admin"
	routeSecretaryBase = "/secretary"
	routeTeacherBase   = "/teacher"
)

const (
	RouteAdmin               = routeAdminBase + "/"
	RouteAdminCalendar       = routeAdminBase + "/calendar"
	RouteAdminSubjects       = routeAdminBase + "/subjects"
	RouteAdminTimetableClass = routeAdminBase + "/timetableclass"

	RouteClass    = "/class"
	RouteLoadDate = "/loaddate"
	RouteLoadTerm = "/loadterm"
	RouteLogin    = "/login"
	RouteLogout   = "/logout"
	RouteMonth    = "/month"
	RouteSubject  = "/subject"
	RouteTerm     = "/term"
	RouteVacation = "/vacation"
	RouteWeek     = "/week"

	RouteSecretary = routeSecretaryBase + "/"

	RouteTeacher = routeTeacherBase + "/"
)

func RedirectToBackURL(w http.ResponseWriter, r *http.Request) {
	backURL := r.FormValue(KeyBackURL)
	http.Redirect(w, r, backURL, http.StatusSeeOther)
}
