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
