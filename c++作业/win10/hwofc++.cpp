#include"pch.h"
#include<iostream>
#include<vector>
#include<string>

#include<fstream>
#include"employee.h"
using namespace std;
template<typename datatype>
void quicksort(vector<datatype*> &vec, int low, int high)//必须传引用,否则出错,因为vector是一个类对象
{
	if (low < high)
	{
		int l = low;
		int r = high;
		datatype *key = vec[l];//记录key值
		while (l < r)
		{
			while (l < r&&*vec[r] > *key)//从右往左遍历,找到第一个小于key的元素
				--r;
			vec[l] = vec[r];
			while (l < r&&*key > *vec[l])//从左往右遍历,找到第一个大于key值的元素
				++l;
			vec[r] = vec[l];
		}
		vec[l] = key;//其实此时l=r
		quicksort(vec, low, l - 1);
		quicksort(vec, r + 1, high);
	}
}



int main() {
	/*
	vector<employee*> v;
	int l = 10;
	for (int i = 0; i < l; i++)
	{	
		technician *t = new technician();
	
		t->accumPay = rand() % 100 + 1.5;
		t->SetName("hello" + to_string(rand() % 10));
		v.push_back(t);
	}
	quicksort(v, 0, v.size() - 1);
	for (int i = 0; i < v.size(); i++)
	{
		v[i]->UpperTran();
		cout << v[i]->GetaccumPay() <<v[i]->GetName()<< endl;
	}*/

	manager m1;
	technician t1, t2, t3, t4, t5, t6;
	salesmanager sm1;
	salesman s1;
	string namestr;
	int size = 9, i;
	int tmp;

	vector<employee *>vchar;
	vchar.push_back(&m1);
	vchar.push_back(&t1);
	vchar.push_back(&t2);
	vchar.push_back(&t3);
	vchar.push_back(&t4);
	vchar.push_back(&t5);
	vchar.push_back(&t6);
	vchar.push_back(&sm1);
	vchar.push_back(&s1);

	for (i = 0; i < size; i++) {
		cout << "请输入下一个雇员的姓名:";
		cin >> namestr;
		vchar[i]->SetName(namestr);
		vchar[i]->promote();
	}
	cout << "请输入兼职技术人员" << t1.GetName() << "本月的工作时间：";
	cin >> tmp;
	t1.SetworkHours(tmp);
	cout << "请输入兼职技术人员" << t2.GetName() << "本月的工作时间：";
	cin >> tmp;
	t2.SetworkHours(tmp);
	cout << "请输入兼职技术人员" << t3.GetName() << "本月的工作时间：";
	cin >> tmp;
	t3.SetworkHours(tmp);
	cout << "请输入兼职技术人员" << t4.GetName() << "本月的工作时间：";
	cin >> tmp;
	t4.SetworkHours(tmp);
	cout << "请输入兼职技术人员" << t5.GetName() << "本月的工作时间：";
	cin >> tmp;
	t5.SetworkHours(tmp);
	cout << "请输入兼职技术人员" << t6.GetName() << "本月的工作时间：";
	cin >> tmp;
	t6.SetworkHours(tmp);

	cout << "请输入销售经理" << sm1.GetName() << "所管辖部门本月的销售总额：";
	float s2;
	cin >> s2;
	sm1.Setsales(s2);

	cout << "请输入推销员" << s1.GetName() << "本月的销售额：";
	cin >> s2;
	s1.Setsales(s2);

	for (i = 0; i < size; i++) {
		vchar[i]->pay();
	}

	//转换name 为大写
	for (int i = 0; i < vchar.size(); i++)
	{
		vchar[i]->UpperTran();
	}


	quicksort(vchar, 0, vchar.size()-1);
	cout << "输出所有人员收入排序" << endl;

	for (i = 0; i < size; i++) {
		cout << vchar[i]->GetName() << "编号" << vchar[i]->GetindividualEmpNo() << "级别为" << vchar[i]->Getgrade() << "级，本月工资" << vchar[i]->GetaccumPay() << endl;
	}

	cout << "输出兼职技术人员收入排序" << endl;

	for (i = 0; i < size; i++) {
		if (vchar[i]->num == 2)
			cout << vchar[i]->GetName() << "编号" << vchar[i]->GetindividualEmpNo() << "级别为" << vchar[i]->Getgrade() << "级，本月工资" << vchar[i]->GetaccumPay() << endl;
	}

	ofstream tfile("./hwc.txt");
	for (i = 0; i < size; i++) {
		tfile << vchar[i]->GetName() << "编号" << vchar[i]->GetindividualEmpNo() << "级别为" << vchar[i]->Getgrade() << "级，本月工资" << vchar[i]->GetaccumPay() << endl;
	}
	tfile.close();
	return 0;
}
