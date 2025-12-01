package routes

const (
	routeTerm           = "/term"
	routeVacation       = "/vacation"
	routeCalendar       = "/calendar"
	routeTimetableClass = "/timetable_class"

	routeAdminBase     = "/admin"
	routeSecretaryBase = "/secretary"
	routeTeacherBase   = "/teacher"
)

const (
	RouteLogin  = "/login"
	RouteLogout = "/logout"

	RouteMonth = "/month"

	RouteAdmin               = routeAdminBase + "/"
	RouteAdminLoadTerm       = routeAdminBase + "/load_term"
	RouteAdminTerm           = routeAdminBase + routeTerm
	RouteAdminVacation       = routeAdminBase + routeVacation
	RouteAdminCalendar       = routeAdminBase + routeCalendar
	RouteAdminTimetableClass = routeAdminBase + routeTimetableClass

	RouteSecretary = routeSecretaryBase + "/"

	RouteTeacher = routeTeacherBase + "/"
)
