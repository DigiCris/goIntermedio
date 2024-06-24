/*
1)clono repositorio vacio goIntermedio
git clone repositorioVacioLLamado_goIntermedio
2) Creo carpeta donde voy a trabajar y entro
mkdir goIntermedio/goPOO
cd goIntermedio/goPOO
3) La establezco como modulo
go mod init goIntermedio/goPOO
4) creo un modulo dentro de la carpeta principal (sms/sms.go)
mkdir sms
cd sms
touch sms.go => poner codigo
cd ..
5) Creo archivo principal main.go, compilo y ejecuto
touch main.go => poner codigo
go build -o main .
./main
6) en main importo como nombreModulo/CarpetaModuloEspecifico
import "goIntermedio/goPOO/sms"
avanzado => desde el 14
*/

package main

//import "goIntermedio/goPOO/sms" // modulo/carpetaEspecifica
import "fmt"

type INotificationFactory interface {
	SendNotification()
	GetSender() Isender
}
type Isender interface {
	GetSenderMethod()	string
	GetSenderChannel()	string
}

/////////////////////sms/////////////////////////////
type SMSNotification struct {
} // relacionado con: INotificationFactory

func (SMSNotification) SendNotification() {
	fmt.Println("soy sms")
}

type SMSNotificationSender struct {
}// relacionado con: Isender

func (SMSNotificationSender) GetSenderMethod() string {
	return "Method: sms"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "sms channel"
}

func (SMSNotification) GetSender() Isender {
	return SMSNotificationSender{}
}
/////////////////////////////////////////////////////////

//////////////////////////Email////////////////////////
type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("soy email")
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Method: email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "email channel"
}

func (EmailNotification) GetSender() Isender {
	return EmailNotificationSender{}
}
/////////////////////////////////////////////////////////

// getNotificationFactory Si le paso email crea un tipo EmailNotification, si le paso sms un SmsNotification
func getNotificationFactory(notificationType string) (INotificationFactory,error) {
	switch(notificationType){
	case "email":
		return &EmailNotification{}, nil
	case "sms":
		return &SMSNotification{}, nil
	}
	return nil, fmt.Errorf("No Notification Type")
}

// sendNotification recibe INOtificationFactory y lo manda para el metodo SendNotification que en el caso del sms imprime que es un sms y del email que es un email
func sendNotification(f INotificationFactory) {
	f.SendNotification();
}

/* getMethod recibe INOtificationFactory que puede ser email o sms. 
redirge al GetSender correspondiente de cada uno.
que devuelve un objeto EmailNotificationSender o smsNotificationSender
y llama a la función getSenderMethod de las clases de los ...NotificationSender
Eso devuelve "Method: email" o "Method: sms" dependiendo del caso
imprimo eso que se devolvió
*/
func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod());
}


func main() {
	// creo instancia del objeto
	smsFactory, _ := getNotificationFactory("sms")
	emailFactory, _ := getNotificationFactory("email")
	// Imprimo en panta que tipo es con polimorfismo
	sendNotification(smsFactory)
	sendNotification(emailFactory)
	// dice cual fue el metodo utilizado en cada caso
	getMethod(smsFactory)
	getMethod(emailFactory)
}
