#pragma once
#include "pch.h"
using namespace std;
class employee {
protected:
	
	char name[20];
	string s;
	int individualEmpNo;
	int grade;
	static int employeeNo;
	//float accumPay;
public:
	int num;
	
	float accumPay;
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
	void UpperTran();
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