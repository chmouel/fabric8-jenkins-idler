package condition

import (
	"fmt"
	"strings"

	"github.com/fabric8-services/fabric8-jenkins-idler/internal/model"
	"github.com/fabric8-services/fabric8-jenkins-idler/internal/util"
	"github.com/sirupsen/logrus"
)

// Condition defines a single Eval method which returns true or false.
type Condition interface {
	// Return true if the condition is true for a given object.
	Eval(object interface{}) (bool, error)
}

// Conditions defines map of Condition instances by their names
type Conditions struct {
	conditions map[string]Condition
}

// NewConditions create a new instance of Conditions.
func NewConditions() Conditions {
	return Conditions{
		conditions: make(map[string]Condition),
	}
}

// Eval evaluates a list of Conditions for a given object. It returns false if
// any of the conditions evaluates to false, otherwise true.
func (c *Conditions) Eval(o interface{}) (bool, util.MultiError) {
	errors := util.MultiError{}

	u, ok := o.(model.User)
	if !ok {
		errors.Collect(fmt.Errorf("%T is not of type User", o))
		return false, errors
	}

	log := logrus.WithFields(logrus.Fields{
		"id":        u.ID,
		"name":      u.Name,
		"component": "condition",
	})

	result := true
	condStates := make(map[string]bool)
	for name, ci := range c.conditions {
		r, err := ci.Eval(o)
		if err != nil {
			log.Error(err)
			errors.Collect(err)
		}
		if !r {
			result = false
		}
		condStates[name] = r
	}

	log.Infof("conditions/idle: %t = %s", result, c.conditionMapToString(condStates))
	return result, errors
}

// Add adds a condition with its name to the this Conditions instance.
func (c *Conditions) Add(name string, condition Condition) {
	c.conditions[name] = condition
}

func (c *Conditions) conditionMapToString(conditions map[string]bool) string {
	var result []string
	for key, value := range conditions {
		result = append(result, fmt.Sprintf("%s=%t", key, value))
	}
	return strings.Join(result, " | ")
}
