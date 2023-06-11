package console

import (
	"testing"

	"github.com/stretchr/testify/assert"

	consolemocks "github.com/goravel/framework/contracts/console/mocks"
	"github.com/goravel/framework/support/file"
)

func TestObserverMakeCommand(t *testing.T) {
	observerMakeCommand := &ObserverMakeCommand{}
	mockContext := &consolemocks.Context{}
	mockContext.On("Argument", 0).Return("").Once()
	assert.Nil(t, observerMakeCommand.Handle(mockContext))
	assert.False(t, file.Exists("app/observers/user_observer.go"))

	mockContext.On("Argument", 0).Return("UserObserver").Once()
	assert.Nil(t, observerMakeCommand.Handle(mockContext))
	assert.True(t, file.Exists("app/observers/user_observer.go"))
	assert.True(t, file.Remove("app"))

	mockContext.AssertExpectations(t)
}
