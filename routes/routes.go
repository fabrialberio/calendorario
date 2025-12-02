package routes

const (
	FlagCreate = "create"
	FlagUpdate = "update"
	FlagDelete = "delete"
)

const (
	routeAdminBase     = "/admin"
	routeSecretaryBase = "/secretary"
	routeTeacherBase   = "/teacher"

	routeCalendar       = "/calendar"
	routeTimetableClass = "/timetable_class"
)

const (
	RouteAdmin               = routeAdminBase + "/"
	RouteAdminCalendar       = routeAdminBase + routeCalendar
	RouteAdminTimetableClass = routeAdminBase + routeTimetableClass
	RouteAdminLoadTerm       = routeAdminBase + "/load_term"

	RouteLogin    = "/login"
	RouteLogout   = "/logout"
	RouteMonth 	  = "/month"
	RouteTerm     = "/term"
	RouteVacation = "/vacation"


	RouteSecretary = routeSecretaryBase + "/"

	RouteTeacher = routeTeacherBase + "/"
)
