package consultations

import (
	"github.com/bredbrains/tthk-api-go/models/enums"
	"github.com/bredbrains/tthk-api-go/utils"
)

const (
	GeneralSubjects       = utils.SchoolBaseUrl + "/oppetoo/opetajate-konsultatsioonid/uldainete-konsultatsioonid/"
	Transport             = utils.SchoolBaseUrl + "/oppetoo/opetajate-konsultatsioonid/transporditehnika-valdkonna-konsultatsioonid/"
	Mechanics             = utils.SchoolBaseUrl + "/oppetoo/opetajate-konsultatsioonid/mehaanika-ja-metallitootluse-valdkonna-konsultatsioonid/"
	Energy                = utils.SchoolBaseUrl + "/oppetoo/opetajate-konsultatsioonid/mehhatroonka-osakonna-konsultatsiooid/"
	InformationTechnology = utils.SchoolBaseUrl + "/infotehnoloogia-valdkonna-konsultatsioonid/"
	Logistics             = utils.SchoolBaseUrl + "/logistika-valdkonna-konsultatsioonid/"
	TextileSales          = utils.SchoolBaseUrl + "/oppetoo/opetajate-konsultatsioonid/tekstiili-ja-kaubanduse-valdkonna-konsultatsioonid/"
)

var DepartmentByLink = map[string]int{
	GeneralSubjects:       enums.GeneralSubjects,
	Transport:             enums.Transport,
	Mechanics:             enums.Mechanics,
	Energy:                enums.Energy,
	InformationTechnology: enums.InformationTechnology,
	Logistics:             enums.Logistics,
	TextileSales:          enums.TextileSales,
}
