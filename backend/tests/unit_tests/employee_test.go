package unit_tests

import (
	"testing"
	"time"
	"fmt"
	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
	"github.com/chadawan9913/SE-Lab3-B6409913/entity"
	
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

	t.Run (`employee fail`, func(t *testing.T) {
		employee := entity.Employee{
			FirstName : "chadawan",
			LastName:	"sattaso",
			Email: "chadawan@gmail.com",
			Phone: "098450627290",
			Date: time.Now().AddDate(-1,0,0),
			Vital: "aaaaaaaaa",
		}
		ok, err := govalidator.ValidatorStruct(employee)
		g.expect(ok).NotTo(gomega.BeTrue())
		g.expect(err).NotTo(gomega.BeNil())
		g.expect(err.Err()).To(Equal(fmt.sprintf("Phone: %s dose not as stringlength(10|10)",employee.Phone)))
	})

	
}