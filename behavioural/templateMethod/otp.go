package main

import "fmt"

/* https://golangbyexample.com/template-method-design-pattern-golang/
模板方法设计模式是一种行为设计模式，允许您为特定操作定义模板或算法。让我们用一个例子来了解模板设计模式。
考虑一次性密码或OTP（one time password）的例子。有不同类型的OTP可以触发，例如OTP可以是短信OTP或电子邮件OTP。但无论它是短信OTP
还是电子邮件OTP，OTP过程的整个步骤都是一样的。步骤是
	生成一个随机的n位数字。
	将此号码保存在缓存中，以便日后验证。
	准备内容
	发送通知
	发布指标
即使在将来，让我们假设引入了推送通知OTP，但它仍然会通过上述步骤。
在这种情况下，当特定操作的步骤相同，但操作的步骤可以由不同的实现者以不同的方式实现时，这将成为模板方法设计模式的候选者。
我们定义了一个模板或算法，它由固定数量的方法组成。操作的实现者覆盖了模板的方法。

现在查看下面的代码示例。
	iOtp表示一个接口，该接口定义了任何otp类型都应该实现的方法集
	短信和电子邮件是iOtp接口的实现者
	otp是定义模板方法genAndSendOTP（）的结构。otp嵌入iOtp接口。
重要信息：iOtp接口和otp结构的组合在go中提供了抽象类的功能
*/

type iOtp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
	publishMetric()
}

type Otp struct {
	iOtp iOtp
}

func (o *Otp) genAndSendOTP(otpLength int) error {
	otp := o.iOtp.genRandomOTP(otpLength)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	if err := o.iOtp.sendNotification(message); err != nil {
		return err
	}
	o.iOtp.publishMetric()
	return nil
}

type sms struct {
	Otp
}

func (s *sms) genRandomOTP(l int) string {
	randomOTP := "123"
	fmt.Printf("sms: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *sms) saveOTPCache(otp string) {
	fmt.Printf("sms: saving otp: %s to cache\n", otp)
}

func (s *sms) getMessage(otp string) string {
	return "sms OTP for login is " + otp
}

func (s *sms) sendNotification(message string) error {
	fmt.Printf("sms: sending sms: %s\n", message)
	return nil
}

func (s *sms) publishMetric() {
	fmt.Println("sms: publishing metrics")
}

type email struct {
	Otp
}

func (s *email) genRandomOTP(l int) string {
	randomOTP := "234"
	fmt.Printf("email: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *email) saveOTPCache(otp string) {
	fmt.Printf("email: saving otp: %s to cache\n", otp)
}

func (s *email) getMessage(otp string) string {
	return "email OTP for login is " + otp
}

func (s *email) sendNotification(message string) error {
	fmt.Printf("email: sending sms: %s\n", message)
	return nil
}

func (s *email) publishMetric() {
	fmt.Println("email: publishing metrics")
}

func main() {
	smsOTP := &sms{}
	o := Otp{iOtp: smsOTP}
	o.genAndSendOTP(4)

	emailOTP := &email{}
	o = Otp{iOtp: emailOTP}
	o.genAndSendOTP(4)
}

/* Output:
sms: generating random otp 123
sms: saving otp: 123 to cache
sms: sending sms: sms OTP for login is 123
sms: publishing metrics
email: generating random otp 234
email: saving otp: 234 to cache
email: sending sms: email OTP for login is 234
email: publishing metrics
*/
