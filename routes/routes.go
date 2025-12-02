package routes

// Common id path parameter name used for all routes.
const KeyID = "id"

const (
	routeAdminBase     = "/admin"
	routeSecretaryBase = "/secretary"
	routeTeacherBase   = "/teacher"

	routeCalendar       = "/calendar"
	routeTimetableClass = "/timetableclass"
)

const (
	RouteAdmin               = routeAdminBase + "/"
	RouteAdminCalendar       = routeAdminBase + routeCalendar
	RouteAdminTimetableClass = routeAdminBase + routeTimetableClass
	RouteAdminLoadTerm       = routeAdminBase + "/loadterm"

	RouteLogin    = "/login"
	RouteLogout   = "/logout"
	RouteMonth    = "/month"
	RouteTerm     = "/term"
	RouteVacation = "/vacation"

	RouteSecretary = routeSecretaryBase + "/"

	RouteTeacher = routeTeacherBase + "/"
)
