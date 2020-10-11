package dm_api

const applicationMethod = `func(svr *{{.Service}}Service){{.Method}}(header *protocol.RequestHeader, request *protocolx.{{.Method}}Request) (rsp interface{}, err error) {
	var (
		transactionContext, _          = factory.CreateTransactionContext(nil)
	)
	rsp = &protocolx.{{.Method}}Response{}
	if err=request.ValidateCommand();err!=nil{
		err = protocol.NewCustomMessage(2,err.Error())
	}
	if err = transactionContext.StartTransaction(); err != nil {
		log.Error(err)
		return nil, err
	}
	defer func() {
		transactionContext.RollbackTransaction()
	}()
{{.Logic}}	
	err = transactionContext.CommitTransaction()
	return
}`
