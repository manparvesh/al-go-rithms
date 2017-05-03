#include <bits/stdc++.h>
using namespace std;

void rand_ints() {
	int n = rand() % 100 + 1;
	int ar[n];
    for(int i=0;i<n;i++) ar[i] = rand() % 1000 + 1;

    for(int i=0;i<n;i++) cout << ar[i] << ",";
    cout << endl;

    if(next_permutation(ar, ar + n)){
        for(int i=0;i<n;i++) cout << ar[i] << ",";
        cout << endl;
    }
}

int main(){
    for(int i=0;i<3;i++) rand_ints();
    return 0;
}