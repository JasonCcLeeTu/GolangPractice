package control
import(
	"go_Code/exercise01/model"
	"strconv"
	
)

type CustomerService struct{

	 CustomerList []model.Customer
	 CustomerQuantity int 
     

}


func (customerservice *CustomerService) Add( customer model.Customer){

  
      customerservice.CustomerQuantity++
	  customer.Id = customerservice.CustomerQuantity
	  customerservice.CustomerList= append(customerservice.CustomerList,customer)

}


func (customerservice *CustomerService) Delete(id int) bool {
      
	index:=customerservice.checkIndex(id)
    if  index ==-1{
      return false
	}
	 customerservice.CustomerList=append(customerservice.CustomerList[0:index],customerservice.CustomerList[index+1:]...)
      return true
	
}

func (customerservice *CustomerService) Update(index int,customerMap map[string]interface{}){

         for key,val:=range customerMap{
			 
			 if  val !=""  {
				 switch key {
				 	 case "Name":
					    customerservice.CustomerList[index].Name =val.(string)
				     case "Gender":
						customerservice.CustomerList[index].Gender =val.(string)
					 case "Age":
						intval ,_ :=strconv.Atoi(val.(string))
						customerservice.CustomerList[index].Age =intval
					 case "Phone":
						customerservice.CustomerList[index].Phone =val.(string)
					 case "Email":
						customerservice.CustomerList[index].Email =val.(string)
				 }
				
				
			 }
		 }

}

func (customerservice *CustomerService) UpdateShow(id int) (*model.Customer,int){
	    index:=customerservice.checkIndex(id)
	   if  index==-1{
		   return nil,-1
	   }

	   return &customerservice.CustomerList[index],index
}

func (customerservice * CustomerService) checkIndex(id int) int {
         
	   for index,val:=  range customerservice.CustomerList{
		   if val.Id == id{
			   return index
		   }
		  
	   }
	   return -1


}


func NewService() *CustomerService{
	  customerservice:=&CustomerService{}
      new:= model.NewCustomer(1,"張和","male",22,"0912345674","fwefw@gmail.com")

	  customerservice.CustomerList = append(customerservice.CustomerList,new)
	  customerservice.CustomerQuantity = 1
	return   customerservice
}



