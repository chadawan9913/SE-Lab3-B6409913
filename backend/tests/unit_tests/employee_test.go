package tests

import (
	"testing"
	"time"
	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"github.com/chadawan9913/SE-Lab3-B6409913/backend/entity"
	
)

func TestEmployeeValidation(t *testing.T){
	g := gomega.NewGomegaWithT(t)

	t.Run (`employee success`, func(t *testing.T) {
		employee := entity.Employee{
			FirstName : "chadawan",
			LastName:	"sattaso",
			Email: "chadawan@gmail.com",
			Phone: "0984506272",
			Date: time.Now().AddDate(-1,0,0),
			Vital: "aaaaaaaaa",
		}
		ok, err := govalidator.ValidatorStruct(employee)
		g.expect(ok).To(gomega.BeTrue())
		g.expect(err).To(gomega.BeNil())

	})

	
}