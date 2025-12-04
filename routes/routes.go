package routes

// Common path parameter names used for multiple routes.
const (
	KeyID   = "id"
	KeyDate = "date"
)

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

	RouteClass    = "/class"
	RouteLoadTerm = "/loadterm"
	RouteLogin    = "/login"
	RouteLogout   = "/logout"
	RouteMonth    = "/month"
	RouteTerm     = "/term"
	RouteVacation = "/vacation"
	RouteWeek     = "/week"

	RouteSecretary = routeSecretaryBase + "/"

	RouteTeacher = routeTeacherBase + "/"
)
