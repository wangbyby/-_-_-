#include<iostream>
#include<vector>
#include "sort.cpp"
#include "emp.cpp"
#include<fstream>
using namespace std;
//version only for ubantu
int main(){
    

    manager m1;
	technician t1,t2,t3,t4,t5,t6;
	salesmanager sm1;
	salesman s1;
	char namestr[20];
	int size = 9, i, j;
	int w[6];
	employee * temp;

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
        cin >> w[1];
        t1.SetworkHours(w[1]);
		cout << "请输入兼职技术人员" << t2.GetName() << "本月的工作时间：";
		cin >> w[2];
		t2.SetworkHours(w[2]);
		cout << "请输入兼职技术人员" << t3.GetName() << "本月的工作时间：";
		cin >> w[3];
		t3.SetworkHours(w[3]);
		cout << "请输入兼职技术人员" << t4.GetName() << "本月的工作时间：";
		cin >> w[4];
		t4.SetworkHours(w[4]);
		cout << "请输入兼职技术人员" << t5.GetName() << "本月的工作时间：";
		cin >> w[5];
		t5.SetworkHours(w[5]);
		cout << "请输入兼职技术人员" << t6.GetName() << "本月的工作时间：";
		cin >> w[6];
		t6.SetworkHours(w[6]);

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
    

	quicksort(vchar , 0, vchar.size() - 1);
	cout << "输出所有人员收入排序" << endl;

	for (i = 0; i < size; i++) {
		cout << vchar[i]->GetName() << "编号" << vchar[i]->GetindividualEmpNo() << "级别为" << vchar[i]->Getgrade() << "级，本月工资" << vchar[i]->GetaccumPay() << endl;
	}

	cout << "输出兼职技术人员收入排序" << endl;

	for (i = 0; i < size; i++) {
		if (vchar[i]->num == 2)
			cout << vchar[i]->GetName() << "编号" << vchar[i]->GetindividualEmpNo() << "级别为" << vchar[i]->Getgrade() << "级，本月工资" << vchar[i]->GetaccumPay() << endl;
	}

	ofstream tfile("./hwc.txt", ios::trunc);
	for (i = 0; i < size; i++) {
		tfile << vchar[i]->GetName() << "编号" << vchar[i]->GetindividualEmpNo() << "级别为" << vchar[i]->Getgrade() << "级，本月工资" << vchar[i]->GetaccumPay() << endl;
	}
	tfile.close();

    return 0;
}