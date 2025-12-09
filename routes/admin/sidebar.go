package admin

import (
	"calendorario/pkg/templates"
	"calendorario/routes"
)

func SidebarRows(selectedIndex int) []templates.SidebarRow {
	rows := []templates.SidebarRow{
		{FAIcon: "fa-calendar-days", Label: "Calendario", Href: routes.RouteAdminCalendar},
		{FAIcon: "fa-calendar-week", Label: "Orario", Href: routes.RouteAdminTimetableClass},
		{FAIcon: "fa-book", Label: "Materie", Href: routes.RouteAdminSubjects},
		{FAIcon: "fa-arrow-right-from-bracket", Label: "Esci", Href: routes.RouteLogout},
	}

	for i, _ := range rows {
		rows[i].Selected = selectedIndex == i
	}

	return rows
}
