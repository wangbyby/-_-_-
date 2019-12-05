#include<vector>
#include<iostream>
#include<string>
#include<vector>
using namespace std;

class employee {
protected:
	string name;
	int individualEmpNo;
	int grade;
	static int employeeNo;
	float accumPay;
public:
	int num;
	employee();
	~employee();
	virtual void pay() = 0;
	virtual void promote(int increment = 0);
	void SetName(string);
	string GetName();
	int GetindividualEmpNo();
	int Getgrade();
	float GetaccumPay();
	bool operator > (employee  & com);
    void UpperTran(){
        for (int i = 0; i<name.size(); i++) {
		if (name[i] >= 'a'&&name[i] <= 'z') {
			name[i] -= 32;
		}
	}
    }
};

class technician :public employee
{
protected:
	float hourlyRate;
	int workHours;
public:
	technician();
	void SetworkHours(int wh);
	void pay();
	void promote(int);
	bool operator > (employee  & com);
    void SetIncome(float a) {
        this->accumPay = a;
    }
};

class salesman :virtual public employee
{
protected:
	float CommRate;
	float sales;
public:
	salesman();
	void Setsales(float sl);
	void pay();
	void promote(int);
	bool operator > (employee  & com);
};

class manager :virtual public employee
{
protected:
	float monthlyPay;
public:
	manager();
	void pay();
	void promote(int);
	bool operator > (employee  & com);
};

class salesmanager :public manager, public salesman
{
public:
	salesmanager();
	void pay();
	void promote(int);
	bool operator > (employee  & com);
};




int employee::employeeNo = 1000;

employee::employee()
{
	individualEmpNo = employeeNo++;
	grade = 1;
	accumPay = 0.0;
	num = 0;
}
employee::~employee() {}
void employee::promote(int increment)
{
	grade += increment;
}
void employee::SetName(string names)
{
	this->name = names;
}
string employee::GetName()
{
	return name;
}
int employee::GetindividualEmpNo()
{
	return individualEmpNo;
}
int employee::Getgrade()
{
	return grade;
}
float employee::GetaccumPay()
{
	return  accumPay;
}
bool employee::operator > (employee & com) {
	if ( com.GetaccumPay() >  accumPay) return false;
	else return true;
}

technician::technician()
{
	hourlyRate = 100;
	num = 2;
}
void technician::SetworkHours(int wh)
{
	workHours = wh;
}
void technician::pay()
{
	accumPay = hourlyRate * workHours;
}
void technician::promote(int)
{
	employee::promote(2);
}
bool technician::operator > (employee & com) {
	if (com.GetaccumPay() > accumPay) return false;
	else return true;
}

salesman::salesman()
{
	CommRate = 0.04;
	num = 4;
}
void salesman::Setsales(float sl)
{
	sales = sl;
}
void salesman::pay()
{
	accumPay = sales * CommRate;
}
void salesman::promote(int) {
	employee::promote(0);
}
bool salesman::operator > (employee & com) {
	if (com.GetaccumPay() > accumPay) return false;
	else return true;
}

manager::manager() {
	monthlyPay = 8000;
	num = 1;
}
void manager::pay() {
	accumPay = monthlyPay;
}
void manager::promote(int) {
	employee::promote(3);
}
bool manager::operator > (employee & com) {
	if (com.GetaccumPay() > accumPay) return false;
	else return true;
}

salesmanager::salesmanager() {
	monthlyPay = 5000;
	CommRate = 0.005;
	num = 3;
}
void salesmanager::pay() {
	accumPay = monthlyPay + CommRate * sales;
}
void salesmanager::promote(int) {
	employee::promote(2);
}
bool salesmanager::operator > (employee & com) {
	if (com.GetaccumPay() > accumPay) return false;
	else return true;
}