#include "pch.h"
#include<iostream>
#include<cstring>
#include<vector>
#include"employee.h"
using namespace std;
int employee::employeeNo = 1000;
void employee::UpperTran() {
	for (int i = 0; i < s.size(); i++)
	{
		for (int i = 0; i < s.size(); i++) {
			if (s[i] >= 'a'&&s[i] <= 'z') {
				s[i] -= 32;
			}
		}
	}
} // 转换大写
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
	s = names;
}
string employee::GetName()
{
	return s;
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
	if (     com.GetaccumPay()> accumPay  ) return false;
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