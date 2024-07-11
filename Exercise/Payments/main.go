package main

import (
	"fmt"
	"log"
	// "io/utils"
	"os"
	"strconv"
	// "strings"
)


// "Interface for overloaidng functions "
type PaymentProcess interface{
	pay(float64) (string ,*bool)
}


//  "Converst []string slice to float64 number"
func byteToFloat(b []byte) float64{

	// fmt.Printf("%T",string(b))
	bal,error := strconv.ParseFloat(string(b),64)
	if error!= nil{
		fmt.Println("Some issue while reading balance :",error)
		logError("Some issue while reading balance")

		os.Exit(1)
	}
	return bal
}


// "Read the balance and return it in float64 format"
func getBalance() (float64){

	// "ioutil.ReadFile() ka function and os.ReadFile() same hai"
	content, error := os.ReadFile("balance.txt")

	if error != nil {
		fmt.Println("Some issue while reading balance : ",error)
		logError("Some issue while reading balance")
		os.Exit(1)
	}
	return byteToFloat(content)
	
}

//  "function to update balance from account"
func updateBalance(m float64) *bool{
	status := false
	bal , openErr := os.OpenFile("balance.txt" , os.O_RDWR , 0666)
	if openErr != nil {
		fmt.Println("issue while upating balance : ",openErr)
		logError("issue while upating balance")
		os.Exit(1)
	}
	// fmt.Printf("%T",bal)
	data := make([]byte , 1000)
	// defer bal.Close()
	dataToUpdate , readErr :=  bal.Read(data)
	if readErr != nil {
		fmt.Println("Some issue while reading balance : ",readErr)
		logError("Some issue while reading balance")
		os.Exit(1)
	}
	// fmt.Println(dataToUpdate)
	// fmt.Printf("Read %d bytes:\n%s\n", dataToUpdate, data[:dataToUpdate])
	
	// fmt.Printf("%v , %T",string(data) , data)

	toUpdate,convertErr := strconv.ParseFloat(string(data[:dataToUpdate]),64)
	if convertErr!= nil{
		fmt.Println("Server Down :",convertErr)
		logError("Some issue while converting")
		os.Exit(1)
	}

	if toUpdate < m {
		fmt.Println("Out of balance")
		logError("Out of balance")
		os.Exit(1)
	}
	toUpdate = toUpdate - m 

	// "Convert float to string"
	str := strconv.FormatFloat(toUpdate, 'f', -1, 64)

	
	writeErr:= os.WriteFile("balance.txt", []byte(str), 0666)
	if convertErr!= nil{
		fmt.Println("Server Down :",writeErr)
		logError("issue while upating balance")
		os.Exit(1)
	}else{
		status = true 
	}
	fmt.Println("Payment Done! of amout :",m)
	fmt.Println("Accont balance :",getBalance())
	return &status
}

// "Struct jo Gpay ka data store karega"
type gpay struct{
	name string
	upiID string
}

func (g gpay)pay(f float64) (string,*bool){
	status := updateBalance(f)
	log := fmt.Sprint(f,"/- amount send via Gpay, paid by ",g.name)
	return log,status
}

// "Struct jo credit card  ka data store karega"
type creditCard struct{
	name string
	cardNumber string
	cvv string
	expiryDate string
}

func (c creditCard)pay(f float64) (string,*bool){
	status := updateBalance(f)
	log := fmt.Sprint(f,"/- amount send via creditCard, paid by ",c.name)
	return log,status
}


// "Method which wites to log after payment"
func writeLogs(p PaymentProcess, amount float64){
	logFile , logErr := os.OpenFile("data.log" , os.O_RDWR|os.O_CREATE | os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Log file Error : ",logErr)
		os.Exit(1)
	}
	defer logFile.Close()
	// str := strconv.FormatFloat(getBalance(), 'f', -1, 64)
	logger := log.New(logFile, "", log.LstdFlags)
	// logger.Println("Balance in account : "+str)
	
	message, status := p.pay(amount)
	if *status!= true{
		fmt.Println("Payment Fail")
		os.Exit(1)
	}
	fmt.Println("Log saved")
	fmt.Println(message,"\n")
	// logger := log.New(logFile, "", log.LstdFlags)
	logger.Println(message)
	// logError(message)
}

// Function to log errors to data.log
func logError(msg string) {
	logFile, logErr := os.OpenFile("data.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Log file Error:", logErr)
		return
	}
	defer logFile.Close()

	logger := log.New(logFile, "", log.LstdFlags)
	logger.Println("ERROR:", msg)
}

func main() {
	fmt.Println("Accont balance :",getBalance(),"\n")

	logFile , logErr := os.OpenFile("data.log" , os.O_RDWR|os.O_CREATE | os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Log file Error : ",logErr)
		os.Exit(1)
	}
	defer logFile.Close()
	str := strconv.FormatFloat(getBalance(), 'f', -1, 64)
	logger := log.New(logFile, "", log.LstdFlags)
	logger.Println("Balance in account : "+str)
	
	satyam := gpay{
		name : "satyam",
		upiID: "satyam@okaxis",
	}
	// satyam.pay(95)
	// writeLogs(satyam,10)


	Jaduu := creditCard{
		name: "Jadu",
		cardNumber: "1234567898",
		expiryDate: "07/26",
		cvv: "235",
	}
	// Jaduu.pay(10)
	// writeLogs(Jaduu,20)

	var money float64
	var name int
	// Prompt user for input
	fmt.Println("Choose payment method : \n1-> Gpay\n2-> CreditCard ")
	fmt.Scan(&name)
	  // Multiple values in case statement 
	switch name { 
		case 1: 
			fmt.Print("Enter payment amount : ") 
			fmt.Scan(&money)
			writeLogs(satyam,money)

		case 2: 
			fmt.Print("Enter payment amount : ") 
			fmt.Scan(&money)
			writeLogs(Jaduu,money)

		default : 
			os.Exit(1)
	}   

}


