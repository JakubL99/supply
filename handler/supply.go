package handler

import (
	"context"
	"strconv"
	pb "supply/proto"

	"github.com/micro/micro/v3/service/logger"
)

type Handler struct {
}

type Price struct {
	Price string `bson:"price"`
}

func MarshalPrice(price *pb.Price) *Price {
	logger.Info("TU nie ma błedu")
	return &Price{
		Price: price.Price,
	}
}

func UnmarshalExpense(expense Price) pb.Expense {
	return pb.Expense{
		Expense: expense.Price,
	}
}

func (h *Handler) Calculate(ctx context.Context, req *pb.Price, rsp *pb.Expense) error {
	price := MarshalPrice(req)
	var expense float64
	logger.Info("Succes MarshalPrice")

	floatPrice, err := strconv.ParseFloat(price.Price, 64)
	if err != nil {
		logger.Error("Error strconv string to Float:  ", err)
	}

	logger.Info("Succes PasrseFloat")

	if floatPrice < 100.00 {
		expense = 25.00
		str := strconv.FormatFloat(expense, 'f', 2, 64)
		var p Price
		p.Price = str
		rp := UnmarshalExpense(p)
		rsp.Expense = rp.Expense

	} else if floatPrice < 200.00 {
		expense = 15.00
		str := strconv.FormatFloat(expense, 'f', 2, 64)
		var p Price
		p.Price = str
		rp := UnmarshalExpense(p)
		rsp.Expense = rp.Expense

	} else {
		expense = 00.00
		str := strconv.FormatFloat(expense, 'f', 2, 64)
		var p Price
		p.Price = str
		rp := UnmarshalExpense(p)
		rsp.Expense = rp.Expense
	}
	logger.Info("Success Calculate Price Supply")
	return nil
}
