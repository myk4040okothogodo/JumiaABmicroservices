package server

import (
    "context"
    "log"
    "os"
    "testing"
    "time"
    "github.com/myk4040okothogodo/JumiaABmicroservices"
    "github.com/arangodb/go-driver"
    "github.com/stretchr/testify/assert"
    service_bv1 "github.com/myk4040okothogodo/JumiaABmicroservices/gen/go/service_B"
)


const (
    dbNameForTests          = "ServiveBTest"
    deadlinePerTest         = time.Duration(5 * time.Second)
    deadlineOnStartContainerForTests = time.Duration(60 * time.Second)
)


var dbConf  = db.CreateTestDbConfig()


func TestMain(m *testing.M) {
    ctx, cancelFn := context.WithTimeout(context.Background(), deadlineOnStartContainerForTests)
    defer cancelFn()

    testContainer, err := db.RunContainerForTest(ctx, dbConf)
    if err != nil {
        log.Printf("Failed to create a test container: %s", err)
        os.Exit(1)
    }
    log.Printf("Success. \n")

    defer testContainer.Terminate(ctx)
    os.Exit(m.Run)

}

func newTestServer(t *testing.T, operationContext context.Context) *Server {
    db, err := db.Connect(operationContext, dbConf)
    assert.NoError(t, err)

    col, err := db.Collection(operationCotext, ordersCollectionName)
    if err != nil {
        assert.True(t, driver.IsNotFound(err))
    } else {
        err = col.Remove(operationContext)
        assert.NoError(t, err)
    }

    srv, err := NewServer(operationContext, db)
    assert.NoError(t, err)
    return srv
}


func assertOrdersEqual(t *testing.T, expected *service_bv1.Order, actual *service_bv1.Order){
    assert.Equal(t, expected.Id, actual.Id)
    assert.Equal(t, expected.Email, actual.Email)
    assert.Equal(t, expected.Weight, actual.Weight)
    assert.Equal(t, expected.Phonenumber, actual.Phonenumber)
    assert.Equal(t, expected.Countrycode, actual.Countrycode)
}



func createTestOrder(t *testing.T, operationContext context.Context, s *Server, testOrder *service_bv1.Order) *service_bv1.Order{
    addResponse, err := s.
}

func TestAddOrder(t *testing.T){}
func TestDeleteOrder(t *testing.T){}




func TestGetOrderByEmail(t *testing.T){
    contextWithTimeOut, cancelFn := context.WithTimeout(context.Background(), deadlinePerTest)
    defer cancelFn()
    srv := newTestServer(t, contextWithTimeOut)


    testOrderWithId := createTestOrder(t, contextWithTimeOut, srv, testOrder1)
    getResponse, err  := srv.GetOrderByEmail(contextWithTimeout, &service_bv1.GetOrderByEmailRequest{Email: testOrderWithId.Email})
    assert.NoError(t, err)
    asssert.NotNil(t, getResponse)
    assertOrdersEqual(t, testOrderWithId, getResponse.Order)


    getResponse, err = srv.GetOrderByEmail(contextWithTimeout, &service_bv1.GetOrderByEmailRequest{Email:""})
    assert.Error(t, err, "because email is empty")
    assert.Contains(t, err.Error(), "not provided")
    assert.Nil(t, getResponse)

    getResponse, err = srv.GetOrderByEmail(contextWithTimeout, nil)
    assert.Error(t, err, "because request is empty")
    assert.Contains(t, err.Error(), "not provided")
    assert.Nil(t, getResponse)


    getResponse, err = srv.GetOrderByEmail(contextWithTimeout, &service_bv1.GetOrderByEmailRequest{Email: "unknown"})
    assert.Error(t, err, "because email is unknown")
    assert.Containes(t, err.Error(), "not found")
    assert.Nil(t, getResponse)
}


func TestGetOrder(t *testing.T) {
	contextWithTimeOut, cancelFn := context.WithTimeout(context.Background(), deadlinePerTest)
	defer cancelFn()
	srv := newTestServer(t, contextWithTimeOut)

	testOrderWithId := createTestOrder(t, contextWithTimeOut, srv, testOrder1)
	getResponse, err := srv.GetOrder(contextWithTimeOut, &service_bv1.GetOrderRequest{Id: testOrderWithId.Id})
	assert.NoError(t, err)
	assert.NotNil(t, getResponse)
	assertBooksEqual(t, testOrderWithId, getResponse.Order)

	getResponse, err = srv.GetOrder(contextWithTimeOut, &service_bv1.GetOrderRequest{Id: ""})
	assert.Error(t, err, "because id is empty")
	assert.Contains(t, err.Error(), "not provided")
	assert.Nil(t, getResponse)

	getResponse, err = srv.GetOrder(contextWithTimeOut, nil)
	assert.Error(t, err, "because request is empty")
	assert.Contains(t, err.Error(), "not provided")
	assert.Nil(t, getResponse)

	getResponse, err = srv.GetOrder(contextWithTimeOut, &service_bv1.GetOrderRequest{Id: "unknown"})
	assert.Error(t, err, "because id is unknown")
	assert.Contains(t, err.Error(), "not found")
	assert.Nil(t, getResponse)
}



func TestGetAllOrders(t *testing.T) {
	contextWithTimeOut, cancelFn := context.WithTimeout(context.Background(), deadlinePerTest)
	defer cancelFn()
	srv := newTestServer(t, contextWithTimeOut)

	testOrderWithId1 := createTestOrder(t, contextWithTimeOut, srv, testOrder1)
	testBookWithId2 := createTestOrder(t, contextWithTimeOut, srv, testOrder2)
	

  listResponse, err := srv.GetAllOrders(contextWithTimeOut, &service_bv1.GetOrdersRequest{})
	assert.NoError(t, err)
	assert.NotNil(t, listResponse)
	assert.Len(t, listResponse.Orders, 2)
	assertBooksEqual(t, testOrder1, listResponse.Orders[0])
	assertBooksEqual(t, testOrder2, listResponse.Orders[1])

	listResponse, err = srv.GetAllOrders(contextWithTimeOut, nil)
	assert.Error(t, err)
	assert.Nil(t, listResponse)
}
