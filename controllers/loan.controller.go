package controllers

import (
	"math"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jpbmdev/payment-api/errorsResponse"
	"github.com/jpbmdev/payment-api/models"
	"github.com/jpbmdev/payment-api/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// -----------------------------------------------
// -- Loan controller
// -----------------------------------------------
type LoanController interface {
	CreateLoan(ctx *gin.Context)
	GetLoans(ctx *gin.Context)
	GetLoanById(ctx *gin.Context)
	GetPaymentsByLoanId(ctx *gin.Context)
	AddPaymentToLoan(ctx *gin.Context)
}

type loanController struct {
	service             services.LoanService
	userService         services.UserService
	targetSchemaService services.TargetSchemaService
	decisionTreeService services.DecisionTreeService
	paymentService      services.PaymentService
}

//Function to crete new loan controller
func NewLoanController() LoanController {
	return &loanController{
		service:             services.NewLoanService(),
		userService:         services.NewUserService(),
		targetSchemaService: services.NewTargetSchemaService(),
		decisionTreeService: services.NewDecisionTreeService(),
		paymentService:      services.NewPaymentService(),
	}
}

// AddPaymentToLoan godoc
// @Summary Add Payment To Loan
// @Schemes
// @Description Add Payment To Loan
// @Tags loan
// @Accept json
// @Produce json
// @Param id  path string true "ID"
// @Param AddPaymentToLoanDto body models.AddPaymentToLoanDto true "payload"
// @Success 200 {object}  models.Payment
// @Failure 400 {object}  models.FailedOperation
// @Failure 404 {object}  models.FailedOperation
// @Failure 500 {object}  models.FailedOperation
// @Router /loan/{id}/payment [put]
func (c *loanController) AddPaymentToLoan(ctx *gin.Context) {

	var addPaymentToLoanDto models.AddPaymentToLoanDto

	//Get request body
	err := ctx.ShouldBindJSON(&addPaymentToLoanDto)
	if err != nil {
		errorsResponse.Error400(ctx, err.Error())
		return
	}

	//check if date is in ISO 8601
	paymentDate, err := time.Parse("2006-01-02", addPaymentToLoanDto.Date)
	if err != nil {
		errorsResponse.Error400(ctx, err.Error())
		return
	}

	//Check if the id passed is a mongoID
	loanId, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		errorsResponse.Error400(ctx, "Invalid Loan ID")
		return
	}

	//Find the loan
	var loan models.Loan
	err = c.service.FindLoanById(loanId, &loan)
	//Handle errors
	if err != nil {
		errorsResponse.Error404(ctx, "Loan does not exists")
		return
	}

	//Check if the loan has debt
	if loan.Debt == 0 {
		errorsResponse.Error400(ctx, "The loan has no debt")
		return
	}

	//Check if the amount to pay is valid
	if loan.Debt < addPaymentToLoanDto.Amount {
		errorsResponse.Error400(ctx, "The amount is greater than the debt")
		return
	}

	//Find the index of the month of the payment in the loan history
	index := -1
	for i := range loan.LoanHistory {
		if paymentDate.After(loan.LoanHistory[i].MonthStart.AddDate(0, 0, -1)) && paymentDate.Before(loan.LoanHistory[i].MonthEnd.AddDate(0, 0, 1)) {
			index = i
			break
		}
	}

	if index == -1 {
		errorsResponse.Error400(ctx, "This date is not in range of the loan")
		return
	}

	if loan.LoanHistory[index].PaymentId != primitive.NilObjectID {
		errorsResponse.Error400(ctx, "This month is already paid")
		return
	}

	//Create the payment
	payment := models.Payment{
		ID:     primitive.NewObjectID(),
		LoanId: loan.ID,
		Date:   paymentDate,
		Amount: addPaymentToLoanDto.Amount,
	}

	//Update Debt
	loan.Debt -= addPaymentToLoanDto.Amount
	//Round to two decimals
	loan.Debt = math.Round((loan.Debt)*100) / 100
	//Update loan history
	loan.LoanHistory[index].PaymentId = payment.ID
	for i := index; i < len(loan.LoanHistory); i++ {
		loan.LoanHistory[i].Accumulated += addPaymentToLoanDto.Amount
		//Round to two decimals
		loan.LoanHistory[i].Accumulated = math.Round((loan.LoanHistory[i].Accumulated)*100) / 100
	}

	//The two operations below should be made inside a transaction
	//But for time and simplicity i am not going to do it

	//Add the payment to the db
	err = c.paymentService.CreatePayment(payment)
	if err != nil {
		errorsResponse.Error500(ctx, err.Error())
		return
	}

	//Update the loan in the db
	err = c.service.UpdateLoanPayment(loan.ID, loan.Debt, loan.LoanHistory)
	if err != nil {
		errorsResponse.Error500(ctx, err.Error())
		return
	}

	//Create Success Response
	ctx.JSON(http.StatusOK, payment)
}

// GetLoans godoc
// @Summary Get Loans
// @Schemes
// @Description Get loans, you can pass a start date, and a end date and the endpoint will find all loans STARTED in that range, if no params are passed this will return all loans, this endpoint supports a very simple pagination where you can select the page and the pageSize
// @Tags loan
// @Produce json
// @Param   from      query     string     false  "string valid"
// @Param   to      query     string     false  "string valid"
// @Param   pageSize      query     int     false  "int valid"
// @Param   page     query     int     false  "int valid"
// @Success 200 {object}  models.Loans
// @Failure 500 {object}  models.FailedOperation
// @Router /loan [get]
func (c *loanController) GetLoans(ctx *gin.Context) {
	var fromDate time.Time
	var toDate time.Time

	//If from is passed transform it to a date
	if ctx.Query("from") != "" {
		var err error
		fromDate, err = time.Parse("2006-01-02", ctx.Query("from"))
		if err != nil {
			errorsResponse.Error400(ctx, err.Error())
			return
		}
	}

	//If to is passed transform it to a date
	if ctx.Query("to") != "" {
		var err error
		toDate, err = time.Parse("2006-01-02", ctx.Query("to"))
		if err != nil {
			errorsResponse.Error400(ctx, err.Error())
			return
		}
	}

	//If from and to are passed and are valid check that from < to
	if ctx.Query("from") != "" && ctx.Query("to") != "" && !fromDate.Before(toDate) {
		errorsResponse.Error400(ctx, "From should be before to")
		return
	}

	//Get loans form db
	loans, err := c.service.FindLoansByDate(fromDate, toDate, ctx.Keys["pageSize"].(int), ctx.Keys["page"].(int))

	if err != nil {
		errorsResponse.Error500(ctx, err.Error())
		return
	}

	//Create Success Response
	ctx.JSON(http.StatusOK, loans)
}

// GetLoanById godoc
// @Summary Get Loan by Id
// @Schemes
// @Description Get loan by Id
// @Tags loan
// @Param id  path string true "ID"
// @Success 200 {object}  models.Loan
// @Failure 400 {object}  models.FailedOperation
// @Failure 404 {object}  models.FailedOperation
// @Failure 500 {object}  models.FailedOperation
// @Router /loan/{id} [get]
func (c *loanController) GetLoanById(ctx *gin.Context) {
	//Check if the id passed is a mongoID
	loanId, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		errorsResponse.Error400(ctx, "Invalid Loan ID")
		return
	}

	//Find the loan
	var loan models.Loan
	err = c.service.FindLoanById(loanId, &loan)
	//Handle errors
	if err != nil {
		errorsResponse.Error404(ctx, "Loan does not exists")
		return
	}

	//Create Success Response
	ctx.JSON(http.StatusOK, loan)
}

// GetPaymentsByLoanId godoc
// @Summary Get Payments by loan Id
// @Schemes
// @Description Get Payments by loan Id
// @Tags loan
// @Param id  path string true "ID"
// @Success 200 {object}  models.Payments
// @Failure 400 {object}  models.FailedOperation
// @Failure 500 {object}  models.FailedOperation
// @Router /loan/{id}/payment [get]
func (c *loanController) GetPaymentsByLoanId(ctx *gin.Context) {
	//Check if the id passed is a mongoID
	loanId, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		errorsResponse.Error400(ctx, "Invalid Loan ID")
		return
	}

	loans, err := c.paymentService.GetPaymentsByLoanId(loanId)
	//Handle errors
	if err != nil {
		errorsResponse.Error500(ctx, err.Error())
		return
	}

	//Create Success Response
	ctx.JSON(http.StatusOK, loans)
}

// CreateLoan godoc
// @Summary Create Loan
// @Schemes
// @Description Create loan, Full detail on the readme
// @Tags loan
// @Accept json
// @Produce json
// @Param CreateLoanDto body models.CreateLoanDto true "payload"
// @Success 201 {object}  models.Loan
// @Failure 400 {object}  models.FailedOperation
// @Failure 404 {object}  models.FailedOperation
// @Failure 500 {object}  models.FailedOperation
// @Router /loan [post]
func (c *loanController) CreateLoan(ctx *gin.Context) {
	var createLoanDto models.CreateLoanDto

	//Get request body
	err := ctx.ShouldBindJSON(&createLoanDto)
	if err != nil {
		errorsResponse.Error400(ctx, err.Error())
		return
	}

	//check if date is in ISO 8601
	startDateTime, err := time.Parse("2006-01-02", createLoanDto.StartDate)
	if err != nil {
		errorsResponse.Error400(ctx, err.Error())
		return
	}

	var user models.User
	var latestTargetSchema models.TargetSchema

	var userError, targetSchemaError error

	//Get the user and the target schema in parallel for performance
	//Used waitgroup instad of channels for simplicity and time
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		//check if user exists
		userError = c.userService.FindUserById(createLoanDto.UserId, &user)
	}()
	go func() {
		defer wg.Done()
		//get Latest target schema
		targetSchemaError = c.targetSchemaService.FindLatestTargetSchema(&latestTargetSchema)
	}()
	wg.Wait()

	//Handle errors
	if userError != nil {
		errorsResponse.Error404(ctx, "User does not exists")
		return
	}
	if targetSchemaError != nil {
		errorsResponse.Error400(ctx, targetSchemaError.Error())
		return
	}

	//Get loans started year before this loan year
	lastYearLoans, err := c.service.FindLastYearLoans(user.ID, startDateTime)
	if err != nil {
		errorsResponse.Error500(ctx, err.Error())
		return
	}

	//Calculate the number of loans and total amount
	numLoans := len(lastYearLoans)
	totalAmount := 0.0
	for i := range lastYearLoans {
		totalAmount += lastYearLoans[i].Amount
	}

	//This function create and test the decision tree with the targetSchema
	target, err := c.decisionTreeService.CreateAndExecuteDecisionTree(
		latestTargetSchema.DesicionTree,
		float64(numLoans),
		float64(totalAmount),
	)

	if err != nil {
		errorsResponse.Error500(ctx, err.Error())
		return
	}

	//Find the target rate and max
	var targetParam models.TargetParams
	for i := range latestTargetSchema.Targets {
		if latestTargetSchema.Targets[i].Name == target {
			targetParam = latestTargetSchema.Targets[i]
		}
	}

	if createLoanDto.Amount > float64(targetParam.Max) {
		errorsResponse.Error400(
			ctx,
			"The max amount of your target "+targetParam.Name+" is: "+strconv.Itoa(targetParam.Max),
		)
		return
	}

	//Calculate the quota of each month
	quota := c.service.CalculateQuota(float64(createLoanDto.Term), targetParam.Rate, createLoanDto.Amount)
	//Round to two decimals
	quota = math.Round((quota)*100) / 100

	monthStarDateTime := startDateTime
	loanHistory := []models.LoanHistory{}
	monthDebt := 0.0

	//Create the loan history to know the debt on each month
	for i := 0; i < createLoanDto.Term; i++ {
		//Round to two decimals
		monthDebt = math.Round((monthDebt+quota)*100) / 100
		loanHistoryItem := models.LoanHistory{}
		loanHistoryItem.Accumulated = 0
		loanHistoryItem.MonthDebt = monthDebt
		loanHistoryItem.MonthStart = monthStarDateTime
		loanHistoryItem.PaymentId = primitive.NilObjectID
		loanHistoryItem.MonthEnd = monthStarDateTime.AddDate(0, 1, -1)
		monthStarDateTime = monthStarDateTime.AddDate(0, 1, 0)
		loanHistory = append(loanHistory, loanHistoryItem)
	}

	loanToCreate := models.Loan{
		ID:             primitive.NewObjectID(),
		Amount:         createLoanDto.Amount,
		Term:           createLoanDto.Term,
		Rate:           targetParam.Rate,
		UserId:         createLoanDto.UserId,
		TargetSchemaId: latestTargetSchema.ID,
		TargetName:     targetParam.Name,
		StartDate:      startDateTime,
		EndDate:        monthStarDateTime.AddDate(0, 0, -1),
		Quota:          quota,
		Debt:           monthDebt,
		LoanHistory:    loanHistory,
	}

	//Create the loan on the database
	err = c.service.CreateLoan(loanToCreate)
	if err != nil {
		errorsResponse.Error500(ctx, err.Error())
		return
	}

	//Create Success Response
	ctx.JSON(http.StatusCreated, loanToCreate)
}
