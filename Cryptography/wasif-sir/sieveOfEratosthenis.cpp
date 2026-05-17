#include <bits/stdc++.h>
using namespace std;
int main() {
  int prime[1000] = {0};
  int N = 500;
  for (int i = 0; i <= N; i++)
    prime[i] = i;
  prime[0] = prime[1] = 0;
  for (int i = 2; i * i <= N; i++) {
    if (prime[i] != 0) {
      for (int j = i * i; j <= N; j += i) {
        prime[j] = 0;
      }
    }
  }
  for (int i = 2; i <= N; i++) {
    if (prime[i] != 0) {
      cout << prime[i] << " ";
    }
  }
  cout << endl;
  return 0;
}
